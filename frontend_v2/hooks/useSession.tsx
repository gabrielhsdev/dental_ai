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
        } catch (err) {
            console.error('Login error:', err);
            toast.error('An unexpected error occurred during login.');
        }
    };

    const isLoggedIn = () => {
        const token = localStorage.getItem('token');
        if (token) {
            setSession((prev) => ({ ...prev, isLoading: true, error: null }));
            requestMe(token)
                .then((res) => {
                    if (isErrorResponse(res)) {
                        handleError(res.message);
                        return;
                    }
                    setSession({
                        user: res.data,
                        isLoading: false,
                        error: null,
                    });
                })
                .catch((err) => {
                    console.error('Error fetching user data:', err);
                    handleError('Failed to fetch user data.');
                });
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
        session
    };
};
