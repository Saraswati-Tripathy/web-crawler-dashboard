import axios from 'axios';

const API = axios.create({
    baseURL: 'http://localhost:8080/api',
});

API.interceptors.request.use(config => {
    const token = localStorage.getItem('token');
    if (token) config.headers.Authorization = `Bearer ${token}`;
    return config;
});

export const login = (data: { email: string; password: string }) =>
    API.post('/login', data);

export const register = (data: { username: string; email: string; password: string }) =>
    API.post('/register', data);

export const crawlURL = (data: string) =>
    API.post('/crawl', { url: data });

export const getResults = () =>
    API.get('/results');
