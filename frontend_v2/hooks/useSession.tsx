'use client';

import { useRouter } from 'next/navigation';
import { useState } from 'react';
import toast from 'react-hot-toast';

import { requestLogin, requestMe, requestRegister } from '@/services/authService';
import { isErrorResponse, PatientInterface, UserInterface } from '@/common/commonInterfaces';
import { LOCAL_STORAGE_KEYS } from '@/common/constants';

interface SessionState {
    user: UserInterface | null;
    isLoading: boolean;
    error: string | null;
}

export const useSession = () => {
    const router = useRouter();
    const [session, setSession] = useState<SessionState>({
        user: null,
        isLoading: false,
        error: null,
    });

    const setLoading = () => setSession((prev) => ({ ...prev, isLoading: true, error: null }));
    const resetSession = () =>
        setSession({
            user: null,
            isLoading: false,
            error: null,
        });

    const handleError = (message: string) => {
        toast.error(message);
        setSession({ user: null, isLoading: false, error: message });
    };

    const handleLogin = async (email: string, password: string) => {
        setLoading();
        try {

            const loginRes = await requestLogin(email, password);
            if (isErrorResponse(loginRes)) return handleError(loginRes.message);

            const meRes = await requestMe(loginRes.data.token);
            if (isErrorResponse(meRes)) return handleError(meRes.message);

            localStorage.setItem(LOCAL_STORAGE_KEYS.TOKEN, loginRes.data.token);
            setSession({ user: meRes.data, isLoading: false, error: null });
            toast.success('Login successful!');
            router.push('/dashboard');
        } catch (error) {
            handleError('Login failed. Please try again.');
        }
    };

    const isLoggedIn = async () => {
        const token = localStorage.getItem(LOCAL_STORAGE_KEYS.TOKEN);

        if (!token) {
            resetSession();
            router.push('/');
            return;
        }

        try {
            setLoading();
            const res = await requestMe(token);

            if (isErrorResponse(res)) {
                handleError(res.message);
                router.push('/');
                return;
            }

            setSession({ user: res.data, isLoading: false, error: null });
        } catch {
            handleError('Failed to fetch user data.');
            router.push('/');
        }
    };

    const getToken = () => {
        const token = localStorage.getItem(LOCAL_STORAGE_KEYS.TOKEN);
        if (!token) {
            resetSession();
            router.push('/');
            return null;
        }
        return token;
    }

    const handleLogout = () => {
        localStorage.removeItem(LOCAL_STORAGE_KEYS.TOKEN);
        resetSession();
        toast.success('Logout successful!');
        router.push('/');
    };

    const handleRegister = async (
        userName: string,
        email: string,
        password: string,
        firstName: string,
        lastName: string
    ) => {
        setLoading();
        try {
            const registerRes = await requestRegister(userName, email, password, firstName, lastName);
            if (isErrorResponse(registerRes)) return handleError(registerRes.message);

            router.push('/');
        } catch (error) {
            handleError('Registration failed. Please try again.');
            return;
        }
    };


    return {
        handleLogin,
        handleLogout,
        handleRegister,
        isLoggedIn,
        getToken,
        session,
    };
};
