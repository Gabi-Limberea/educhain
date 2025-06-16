// src/services/api.ts
import axios from 'axios';

const API_BASE = process.env.REACT_APP_API_BASE_URL;

if (!API_BASE) {
  throw new Error('REACT_APP_API_BASE_URL is not defined');
}

const api = axios.create({
  baseURL: API_BASE,
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  }
});

// Add a request interceptor to always get the latest token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('authToken');
  // for debug
  // console.log('Token in interceptor:', token);
  
  // Ensure headers object exists
  config.headers = config.headers || {};
  
  // Set content type if not already set
  if (!config.headers['Content-Type']) {
    config.headers['Content-Type'] = 'application/json';
  }
  
  // Set authorization if token exists
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  // for debug
  // console.log('Request headers after setting token:', config.headers);
  return config;
});

export default api;
