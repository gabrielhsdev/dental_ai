'use client';

import { useState } from 'react';
import { PatientInterface } from '@/common/commonInterfaces';
import { LOCAL_STORAGE_KEYS } from '@/common/constants';
import { useRouter } from 'next/navigation';

interface SelectedPatientState {
    patient: PatientInterface | null;
    isLoading: boolean;
    error: string | null;
}

export const useSelectedPatient = () => {
    const router = useRouter();
    const [state, setState] = useState<SelectedPatientState>({
        patient: null,
        isLoading: true,
        error: null,
    });

    const loadSelectedPatient = () => {
        try {
            const savedPatient = localStorage.getItem(LOCAL_STORAGE_KEYS.SELECTED_PATIENT);
            if (savedPatient) {
                try {
                    const parsed = JSON.parse(savedPatient);
                    setState({ patient: parsed, isLoading: false, error: null });
                } catch (err) {
                    setState({ patient: null, isLoading: false, error: 'Error parsing patient data' });
                }
            } else {
                setState((prev) => ({ ...prev, isLoading: false }));
            }
        } catch (error) {
            setState({ patient: null, isLoading: false, error: 'Error loading patient data' });
        }
    }

    const selectPatient = (patient: PatientInterface) => {
        localStorage.setItem(LOCAL_STORAGE_KEYS.SELECTED_PATIENT, JSON.stringify(patient));
        setState({ patient, isLoading: false, error: null });
        router.push('/dashboard/patient');
    };

    const clearSelectedPatient = () => {
        localStorage.removeItem(LOCAL_STORAGE_KEYS.SELECTED_PATIENT);
        setState({ patient: null, isLoading: false, error: null });
    };

    return {
        selectedPatient: state.patient,
        isLoading: state.isLoading,
        error: state.error,
        selectPatient,
        clearSelectedPatient,
        loadSelectedPatient,
    };
};
