import React, { useState } from 'react';
import { register } from '../services/api';
import { useNavigate } from 'react-router-dom';

export default function Register() {
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleRegister = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            await register({ username, email, password });
            navigate('/login');
        } catch (err) {
            alert('Registration failed');
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <form onSubmit={handleRegister} className="bg-white p-6 rounded shadow-md w-full max-w-sm">
                <h2 className="text-xl font-bold mb-4">Register</h2>
                <input type="text" value={username} onChange={e => setUsername(e.target.value)} placeholder="Username" className="mb-2 w-full p-2 border" />
                <input type="email" value={email} onChange={e => setEmail(e.target.value)} placeholder="Email" className="mb-2 w-full p-2 border" />
                <input type="password" value={password} onChange={e => setPassword(e.target.value)} placeholder="Password" className="mb-4 w-full p-2 border" />
                <button type="submit" className="w-full bg-green-500 text-white p-2 rounded">Register</button>
            </form>
        </div>
    );
}
