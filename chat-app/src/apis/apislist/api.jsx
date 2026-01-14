import axios from 'axios';
import { signupapi } from '../constants/apiendpoints';
import { loginapi } from '../constants/apiendpoints';
import { sendMessageApi } from '../constants/apiendpoints';
import { getUsersApi } from '../constants/apiendpoints';

export const signup = (username, email, password) => {
  return axios.post(`${signupapi}`, {
    username,
    email,
    password,
  });
};

export const login = (email, password) => {
  return axios.post(`${loginapi}`, {
    email,
    password,
  });
};

export const sendMessage = (reciever_id, content, token) => {
  const payload = {
    reciever_id: reciever_id,
    content: content,
  };
  console.log('Sending payload to /sendmessage:', payload);
  return axios.post(
    `${sendMessageApi}`,
    payload,
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );
};

export const getUsers = (searchQuery, token) => {
  return axios.get(`${getUsersApi}`, {
    params: { user: searchQuery },
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};
