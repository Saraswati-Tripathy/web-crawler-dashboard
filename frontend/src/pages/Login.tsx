import React, { useState } from 'react';
import { login } from '../services/api';
import { useNavigate } from 'react-router-dom';

export default function Login() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const res = await login({ email, password });
            localStorage.setItem('token', res.data.token);
            navigate('/dashboard');
        } catch (err) {
            alert('Login failed');
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <form onSubmit={handleLogin} className="bg-white p-6 rounded shadow-md w-full max-w-sm">
                <h2 className="text-xl font-bold mb-4">Login</h2>
                <input type="email" value={email} onChange={e => setEmail(e.target.value)} placeholder="Email" className="mb-2 w-full p-2 border" />
                <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Password" className="mb-4 w-full p-2 border" />
                <button type="submit" className="w-full bg-blue-500 text-white p-2 rounded">Login</button>
            </form>
        </div>
    );
}
