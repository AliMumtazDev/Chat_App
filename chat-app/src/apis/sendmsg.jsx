import { sendMessage } from './apislist/api';

export const handleSendMessage = async (receiverId, content, token) => {
  try {
    const response = await sendMessage(receiverId, content, token);
    return { success: true, data: response.data };
  } catch (error) {
    if (error.response) {
      return { success: false, error: error.response.data.error || 'Failed to send message' };
    } else if (error.request) {
      return { success: false, error: 'No response from server. Please try again.' };
    } else {
      return { success: false, error: 'An unexpected error occurred.' };
    }
  }
};

let ws;

export const setupWebSocket = (token, onMessageReceived) => {
  if (ws && ws.readyState === WebSocket.OPEN) {
      console.log('WebSocket already connected');
      return ws;
  }

  const wsUrl = `ws://localhost:8004/protected/ws?token=${token}`;
  console.log('Connecting to WebSocket URL:', wsUrl);
  ws = new WebSocket(wsUrl);

  ws.onopen = () => {
      console.log('WebSocket connected');
      ws.send(JSON.stringify({ action: 'ping', message: 'hello from client' }));

      const pingInterval = setInterval(() => {
          if (ws.readyState === WebSocket.OPEN) {
              ws.send(JSON.stringify({ action: 'ping' }));
          }
      }, 30000);

      // Cleanup interval on close
      ws.onclose = () => clearInterval(pingInterval);
  };

  ws.onmessage = (event) => {
      try {
          const message = JSON.parse(event.data);
          console.log('Received WebSocket message:', message);
          onMessageReceived({ ...message, timestamp: new Date().toLocaleTimeString() });
      } catch (error) {
          console.error('Error parsing WebSocket message:', error);
      }
  };

  ws.onclose = (event) => {
      console.log('WebSocket disconnected:', event.code, event.reason);
      if (ws.readyState !== WebSocket.OPEN) {
          setTimeout(() => {
              console.log('Attempting to reconnect...');
              setupWebSocket(token, onMessageReceived);
          }, 3000);
      }
  };

  ws.onerror = (error) => {
      console.error('WebSocket error:', error);
  };

  return ws;
};