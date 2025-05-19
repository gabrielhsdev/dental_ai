'use client';

import { useState } from 'react';
import toast from 'react-hot-toast';

import {
    requestCreatePatient,
    requestGetPatientById,
    requestGetPatientByUserId,
} from '@/services/patientService';

import {
    isErrorResponse,
    PatientInterface,
} from '@/common/commonInterfaces';
import { LOCAL_STORAGE_KEYS } from '@/common/constants';

interface PatientState {
    patients: PatientInterface[];
    selectedPatient: PatientInterface | null;
    isLoading: boolean;
    error: string | null;
}

export const usePatients = () => {
    const [state, setState] = useState<PatientState>({
        patients: [],
        selectedPatient: null,
        isLoading: false,
        error: null,
    });

    const setLoading = () => setState((prev) => ({ ...prev, isLoading: true, error: null }));

    const resetError = () => setState((prev) => ({ ...prev, error: null }));

    const fetchPatientsByUserId = async (token: string) => {
        if (!token) return;

        setLoading();
        try {
            const res = await requestGetPatientByUserId(token);
            if (isErrorResponse(res)) return handleError(res.message);

            setState((prev) => ({
                ...prev,
                patients: res.data,
                isLoading: false,
                error: null,
            }));
        } catch {
            handleError('Failed to fetch patients.');
        }
    };

    const fetchPatientById = async (patientId: string, token: string) => {
        if (!token) return;

        setLoading();
        try {
            const res = await requestGetPatientById(patientId, token);
            if (isErrorResponse(res)) return handleError(res.message);

            setState((prev) => ({
                ...prev,
                selectedPatient: res.data,
                isLoading: false,
                error: null,
            }));
        } catch {
            handleError('Failed to fetch patient details.');
        }
    };

    const createPatient = async (
        firstName: string,
        lastName: string,
        dateOfBirth: string,
        gender: string,
        phoneNumber: string,
        email: string,
        notes: string,
        token: string,
    ) => {
        if (!token) return;

        setLoading();
        try {
            const res = await requestCreatePatient(
                firstName,
                lastName,
                dateOfBirth,
                gender,
                phoneNumber,
                email,
                notes,
                token,
            );

            if (isErrorResponse(res)) return handleError(res.message);

            toast.success('Patient created successfully!');
            setState((prev) => ({
                ...prev,
                patients: [...prev.patients, res.data],
                isLoading: false,
                error: null,
            }));
        } catch {
            handleError('Failed to create patient.');
        }
    };

    const handleError = (message: string) => {
        toast.error(message);
        setState((prev) => ({
            ...prev,
            isLoading: false,
            error: message,
        }));
    };

    return {
        // State
        patients: state.patients,
        selectedPatient: state.selectedPatient,
        isLoading: state.isLoading,
        error: state.error,
        // Fetch functions
        fetchPatientsByUserId,
        fetchPatientById,
        createPatient,
        resetError,
    };
};
