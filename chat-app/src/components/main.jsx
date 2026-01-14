import React, { useState } from 'react';
import SearchUser from './search';
import ChatSection from './chatbox';

const Main = () => {
  const [selectedUser, setSelectedUser] = useState(null);

  const handleUserSelect = (user) => {
    setSelectedUser(user);
  };

  return (
    <div className="flex flex-col md:flex-row h-screen bg-gradient-to-br from-gray-100 to-blue-50">
      <div className="w-full md:w-1/4 border-r border-gray-200 bg-white shadow-md">
        <SearchUser onUserSelect={handleUserSelect} />
      </div>
      <div className="w-full md:w-3/4 bg-white">
        <ChatSection selectedUser={selectedUser} />
      </div>
    </div>
  );
};

export default Main;