// src/services/api.ts
import axios from 'axios';

const API_BASE = process.env.REACT_APP_API_BASE_URL;

if (!API_BASE) {
  throw new Error('REACT_APP_API_BASE_URL is not defined');
}

const api = axios.create({
  baseURL: API_BASE,
  withCredentials: true,
  headers: { 'Content-Type': 'application/json' }
});

// load any saved JWT
const token = localStorage.getItem('authToken');
if (token) {
  api.defaults.headers.common.Authorization = `Bearer ${token}`;
}

export default api;
