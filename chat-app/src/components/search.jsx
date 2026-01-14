import React, { useState } from 'react';
import { getUsers } from '../apis/apislist/api';

const SearchUser = ({ onUserSelect }) => {
  const [searchQuery, setSearchQuery] = useState('');
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(false);
  const token = localStorage.getItem('accessToken');

  const handleSearch = async (e) => {
    const query = e.target.value;
    setSearchQuery(query);
    if (query.length > 2) {
      setLoading(true);
      try {
        const result = await getUsers(query, token);
        console.log('API Response:', result);
        if (result.data) {
          setUsers(result.data.users || []);
        } else {
          setUsers([]);
          console.error('Search failed:', result.error);
        }
      } catch (error) {
        console.error('Error fetching users:', error);
        setUsers([]);
      }
      setLoading(false);
    } else {
      setUsers([]);
    }
  };

  const handleSelect = (user) => {
    onUserSelect(user);
    setSearchQuery('');
    setUsers([]);
  };

  return (
    <div className="p-4">
      <input
        type="text"
        value={searchQuery}
        onChange={handleSearch}
        placeholder="Search users..."
        className="w-full p-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 transition-all duration-200 text-gray-700 placeholder-gray-400"
      />
      {loading && <p className="text-gray-500 text-sm mt-2">Loading...</p>}
      <ul className="mt-3 max-h-64 overflow-y-auto">
        {users.map((user) => (
          <li
            key={user.id}
            onClick={() => handleSelect(user)}
            className="p-3 hover:bg-blue-50 cursor-pointer border-b border-gray-200 transition-colors duration-150 text-gray-800"
          >
            {user.username}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default SearchUser;