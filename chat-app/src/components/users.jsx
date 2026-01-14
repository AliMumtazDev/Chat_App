import React, { useState, useEffect } from 'react';
import { getMessageHistory } from '../apis/apislist/api';

const MapRecentUsers = ({ onUserSelect }) => {
  const [recentUsers, setRecentUsers] = useState([]);
  const token = localStorage.getItem('accessToken');

  useEffect(() => {
    const fetchRecentUsers = async () => {
      const result = await getMessageHistory(null, token); 
      if (result.success) {
        const uniqueUsers = [...new Map(result.data.map(item => [item.reciever_id, item])).values()];
        setRecentUsers(uniqueUsers);
      } else {
        console.error('Failed to fetch recent users:', result.error);
      }
    };
    fetchRecentUsers();
  }, [token]);

  return (
    <div className="p-4">
      <h3 className="text-lg font-semibold mb-2">Recent Users</h3>
      <ul>
        {recentUsers.map((user) => (
          <li
            key={user.reciever_id}
            onClick={() => onUserSelect({ id: user.reciever_id, username: user.username || `User ${user.reciever_id}` })}
            className="p-2 hover:bg-gray-200 cursor-pointer"
          >
            {user.username || `User ${user.reciever_id}`}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default MapRecentUsers;