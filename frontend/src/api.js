import axios from 'axios';
import { API_URL } from './config';

const api = axios.create({
    baseURL: API_URL,
    headers: {
        'Content-Type': 'application/json'
    }
});

// Request interceptor - her istekte çalışır
api.interceptors.request.use(
    config => {
        console.log('API İsteği:', config.method.toUpperCase(), config.url);
        return config;
    },
    error => {
        console.error('API İstek Hatası:', error);
        return Promise.reject(error);
    }
);

// Response interceptor - her yanıtta çalışır
api.interceptors.response.use(
    response => {
        console.log('API Yanıtı:', response.status, response.data);
        return response;
    },
    error => {
        console.error('API Yanıt Hatası:', error.response || error);
        return Promise.reject(error);
    }
);

export default api; 