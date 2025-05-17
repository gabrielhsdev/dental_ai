'use client';

import React, { createContext, useContext, useEffect, useState } from 'react';
import { useSession } from '@/hooks/useSession'; // adjust path as needed
import { useRouter } from 'next/router';

const SessionContext = createContext<ReturnType<typeof useSession> | null>(null);

export const SessionProvider = ({ children }: { children: React.ReactNode }) => {
    const session = useSession();

    useEffect(() => {
        const checkLogin = async () => {
            await session.isLoggedIn();
        };
        checkLogin();
    }, []);

    return (
        <SessionContext.Provider value={session}>
            {children}
        </SessionContext.Provider>
    );
};

export const useSessionContext = () => {
    const context = useContext(SessionContext);
    if (!context) {
        throw new Error('useSessionContext must be used within a SessionProvider');
    }
    return context;
};
