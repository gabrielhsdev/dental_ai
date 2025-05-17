'use client';

import { useRouter } from 'next/navigation'; // app router
import { useEffect, useState } from 'react';
import { isErrorResponse, UserInterface } from '@/common/commonInterfaces';
import { requestLogin, requestMe } from '@/services/authRequests';
import toast from 'react-hot-toast';

interface SessionState {
    user: UserInterface | null;
    isLoading: boolean;
    error: string | null;
}

export const useSession = () => {
    const router = useRouter(); // ✅ always called
    const [session, setSession] = useState<SessionState>({
        user: null,
        isLoading: false,
        error: null,
    });

    const handleLogin = async (email: string, password: string) => {
        setSession((prev) => ({ ...prev, isLoading: true, error: null }));

        try {
            const loginRes = await requestLogin(email, password);
            if (isErrorResponse(loginRes)) {
                handleError(loginRes.message);
                return;
            }

            const meRes = await requestMe(loginRes.data.token);
            if (isErrorResponse(meRes)) {
                handleError(meRes.message);
                return;
            }

            setSession({
                user: meRes.data,
                isLoading: false,
                error: null,
            });

            toast.success('Login successful!');
            localStorage.setItem('token', loginRes.data.token);

            // ✅ Router redirect after successful login
            router.push('/dashboard');
        } catch (err) {
            console.error('Login error:', err);
            toast.error('An unexpected error occurred during login.');
        }
    };

    const isLoggedIn = async () => {
        const token = localStorage.getItem('token');
        if (!token) return false;

        try {
            setSession((prev) => ({ ...prev, isLoading: true, error: null }));
            const res = await requestMe(token);

            if (isErrorResponse(res)) {
                handleError(res.message);
                return false;
            }

            setSession({
                user: res.data,
                isLoading: false,
                error: null,
            });

            router.push('/dashboard');
        } catch (err) {
            console.error('Error fetching user data:', err);
            handleError('Failed to fetch user data.');
            router.push('/');
        }
    };

    const handleLogout = () => {
        localStorage.removeItem('token');
        setSession({
            user: null,
            isLoading: false,
            error: null,
        });
        toast.success('Logout successful!');
        router.push('/');
    };

    const handleError = (message: string) => {
        toast.error(message);
        setSession({
            user: null,
            isLoading: false,
            error: message,
        });
    };

    return {
        handleLogin,
        handleLogout,
        isLoggedIn,
        session,
    };
};
