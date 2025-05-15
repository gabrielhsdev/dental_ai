import { ENDPOINTS, postRequest, requestLogin } from '@/services/requests';
import { useState, useEffect } from 'react';

interface User {
    id: string;
    email: string;
    name: string;
}

interface SessionState {
    user: User | null;
    isLoading: boolean;
    error: string | null;
}

export const useSession = () => {
    const [session, setSession] = useState<SessionState>({
        user: null,
        isLoading: true,
        error: null,
    });

    const handleLogin = async (email: string, password: string) => {
        try {
            const response = await requestLogin(email, password);
            console.log('Login response:', response); 
        } catch (error) {
            console.log("Test")
        }
    };

    return {
        handleLogin,
        session,
    };
};