import React, { useState } from 'react';
import { handleLogin } from '../apis/login';
import { Link, useNavigate } from 'react-router-dom';
import { ToastContainer, toast } from 'react-toastify';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);

    // Clear any existing tokens before login
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');

    const result = await handleLogin(email, password);

    if (result.success) {
      const accessToken = result.data.access_token;
      const refreshToken = result.data.refresh_token;

      if (!accessToken) {
        toast.error('Login successful, but no access token received.');
        console.error('No access token in response:', result.data);
        setLoading(false);
        return;
      }

      // Store the new tokens
      localStorage.setItem('accessToken', accessToken);
      if (refreshToken) {
        localStorage.setItem('refreshToken', refreshToken);
      } else {
        console.log('No refresh token received');
      }

      toast.success(result.data.message || 'Logged in successfully!');
      console.log('Login Success - Access Token:', accessToken);
      setEmail('');
      setPassword('');
      navigate('/chat');
    } else {
      toast.error(result.error);
      console.error('Login Error:', result.error);
    }

    setLoading(false);
  };

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center">
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
        <h2 className="text-2xl font-bold text-center mb-6">Login</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-gray-700 mb-2" htmlFor="email">
              Email
            </label>
            <input
              type="email"
              id="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="w-full p-3 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Enter your email"
              required
            />
          </div>
          <div className="mb-6">
            <label className="block text-gray-700 mb-2" htmlFor="password">
              Password
            </label>
            <input
              type="password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full p-3 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="Enter your password"
              required
            />
          </div>
          <button
            type="submit"
            disabled={loading}
            className={`w-full p-3 rounded-lg text-white ${
              loading ? 'bg-blue-400 cursor-not-allowed' : 'bg-blue-600 hover:bg-blue-700'
            } transition duration-200`}
          >
            {loading ? 'Logging In...' : 'Login'}
          </button>
        </form>
        <div className="mt-4 text-center">
          <p className="text-gray-600">Don't have an account? <Link to="/signup" className="text-blue-600 hover:underline">Signup</Link></p>
        </div>
      </div>
      <ToastContainer position="top-right" autoClose={3000} />
    </div>
  );
};

export default Login;