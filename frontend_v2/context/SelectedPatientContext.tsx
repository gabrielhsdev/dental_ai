'use client';

import React, { createContext, useContext, useEffect } from 'react';
import { useSelectedPatient } from '@/hooks/useSelectedPatient';
import { ClipLoader } from 'react-spinners';

const SelectedPatientContext = createContext<ReturnType<typeof useSelectedPatient> | null>(null);

export const SelectedPatientProvider = ({ children }: { children: React.ReactNode }) => {
    const selectedPatient = useSelectedPatient();

    useEffect(() => {
        selectedPatient.loadSelectedPatient();
    }, []);

    if (selectedPatient.isLoading) {
        return (
            <div className="flex items-center justify-center h-screen">
                <ClipLoader color="#3B82F6" size={50} />
            </div>
        );
    }

    return (
        <SelectedPatientContext.Provider value={selectedPatient}>
            {children}
        </SelectedPatientContext.Provider>
    );
};

export const useSelectedPatientContext = () => {
    const context = useContext(SelectedPatientContext);
    if (!context) {
        throw new Error('useSelectedPatientContext must be used within a SelectedPatientProvider');
    }
    return context;
};
