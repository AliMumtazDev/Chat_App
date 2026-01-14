import React, { useState, useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import { handleSendMessage, setupWebSocket } from '../apis/sendmsg';
import { jwtDecode } from 'jwt-decode';

const ChatSection = ({ selectedUser }) => {
    const [messages, setMessages] = useState([]);
    const [messageInput, setMessageInput] = useState('');
    const chatBoxRef = useRef(null);
    const wsRef = useRef(null);
    const token = localStorage.getItem('accessToken');
    const navigate = useNavigate();
    const [currentUserId, setCurrentUserId] = useState(null);

    // Extract current user ID from token
    useEffect(() => {
        if (token) {
            try {
                const decodedToken = jwtDecode(token);
                setCurrentUserId(decodedToken.user_id || decodedToken.sub); // Use user_id from your token
                console.log('Current User ID:', decodedToken.user_id);
            } catch (error) {
                console.error('Error decoding token:', error);
            }
        }
    }, [token]);

    useEffect(() => {
        if (!token || !selectedUser) return;

        const ws = setupWebSocket(token, (receivedMessage) => {
            console.log("WebSocket message received:", receivedMessage); // Debug log
            setMessages((prev) => [...prev, { ...receivedMessage, timestamp: new Date().toLocaleTimeString() }]);
        });
        wsRef.current = ws;

        return () => {
            if (wsRef.current) {
                wsRef.current.close();
            }
        };
    }, [token, selectedUser]);

    useEffect(() => {
        if (chatBoxRef.current) {
            chatBoxRef.current.scrollTop = chatBoxRef.current.scrollHeight;
        }
    }, [messages]);

    const handleSend = async () => {
        if (messageInput.trim() && selectedUser) {
            const result = await handleSendMessage(selectedUser.id, messageInput, token);
            if (result.success) {
                setMessages((prev) => [
                    ...prev,
                    { message: messageInput, sender_id: currentUserId, timestamp: new Date().toLocaleTimeString() }
                ]);
                setMessageInput('');
            } else {
                console.error('Send failed:', result.error);
            }
        }
    };

    const handleLogout = () => {
        localStorage.removeItem('accessToken');
        console.log('Logged out');
        navigate('/login');
    };

    // Determine if the message is from the current user
    const isSentByCurrentUser = (msg) => {
        return msg.sender_id === currentUserId;
    };

    return (
        <div className="flex-1 p-4 flex flex-col h-full">
            <div className="flex justify-between items-center mb-4">
                {selectedUser ? (
                    <h2 className="text-xl md:text-2xl font-bold text-gray-800">{selectedUser.username}</h2>
                ) : (
                    <h2 className="text-xl md:text-2xl font-bold text-gray-800">Select a user to chat</h2>
                )}
                <button
                    onClick={handleLogout}
                    className="p-2 bg-red-600 text-white rounded-lg hover:bg-red-700 transition-colors duration-200 shadow-md"
                >
                    Logout
                </button>
            </div>
            <div
                ref={chatBoxRef}
                className="flex-1 overflow-y-auto border border-gray-200 rounded-lg p-4 bg-gray-50 shadow-inner mb-4"
            >
                {messages.map((msg, index) => (
                    <div
                        key={index}
                        className={`mb-3 flex ${isSentByCurrentUser(msg) ? 'justify-end' : 'justify-start'}`}
                    >
                        <div
                            className={`max-w-xs md:max-w-md p-3 rounded-lg shadow-md ${
                                isSentByCurrentUser(msg)
                                    ? 'bg-blue-600 text-white'
                                    : 'bg-gray-200 text-gray-800'
                            }`}
                        >
                            <p>{msg.message || msg.content}</p>
                            <span className="text-xs text-gray-400 block mt-1">{msg.timestamp}</span>
                        </div>
                    </div>
                ))}
            </div>
            <div className="flex items-center space-x-2">
                <input
                    type="text"
                    value={messageInput}
                    onChange={(e) => setMessageInput(e.target.value)}
                    onKeyPress={(e) => e.key === 'Enter' && handleSend()}
                    className="flex-1 p-3 border border-gray-300 rounded-l-lg focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all duration-200 text-gray-700 placeholder-gray-400"
                    placeholder="Type a message..."
                />
                <button
                    onClick={handleSend}
                    className="p-3 bg-blue-600 text-white rounded-r-lg hover:bg-blue-700 transition-colors duration-200 shadow-md"
                >
                    Send
                </button>
            </div>
        </div>
    );
};

export default ChatSection;