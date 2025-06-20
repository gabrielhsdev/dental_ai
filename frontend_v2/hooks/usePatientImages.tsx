'use client';

import { useState } from 'react';
import toast from 'react-hot-toast';

import {
    requestCreatePatientImage,
    requestGetPatientImageById,
    requestGetPatientImagesByPatientId,
    requestGetPatientByUserId,
} from '@/services/patientImageService';

import {
    requestDiagnosticProcess,
    requestGetImageById,
    requestDiagnosticPredict,
} from '@/services/diagnosticsService';

import {
    isErrorResponse,
    PatientImageInterface,
} from '@/common/commonInterfaces';

interface PatientImageState {
    images: PatientImageInterface[];
    selectedImage: PatientImageInterface | null;
    isLoading: boolean;
    error: string | null;
}

export const usePatientImages = () => {
    const [state, setState] = useState<PatientImageState>({
        images: [],
        selectedImage: null,
        isLoading: false,
        error: null,
    });

    const setLoading = () =>
        setState((prev) => ({ ...prev, isLoading: true, error: null }));

    const resetError = () =>
        setState((prev) => ({ ...prev, error: null }));

    const handleError = (message: string) => {
        toast.error(message);
        setState((prev) => ({
            ...prev,
            isLoading: false,
            error: message,
        }));
    };

    const fetchImagesByPatientId = async (patientId: string, token: string) => {
        if (!token) return;
        setLoading();

        try {
            const res = await requestGetPatientImagesByPatientId(patientId, token);
            if (isErrorResponse(res)) return handleError(res.message);
            setState((prev) => ({
                ...prev,
                images: res.data,
                isLoading: false,
                error: null,
            }));
        } catch {
            handleError('Failed to fetch patient images.');
        }
    };

    const fetchImageById = async (imageId: string, token: string) => {
        if (!token) return;
        setLoading();

        try {
            const res = await requestGetPatientImageById(imageId, token);
            if (isErrorResponse(res)) return handleError(res.message);

            setState((prev) => ({
                ...prev,
                selectedImage: res.data,
                isLoading: false,
                error: null,
            }));
        } catch {
            handleError('Failed to fetch image.');
        }
    };

    const createPatientImage = async (
        patientId: string,
        imageData: string,
        fileType: string,
        description: string,
        inferenceData: string,
        token: string,
    ) => {
        if (!token) return;
        setLoading();

        try {
            const res = await requestCreatePatientImage(
                patientId,
                imageData,
                fileType,
                description,
                inferenceData,
                token,
            );

            if (isErrorResponse(res)) return handleError(res.message);

            toast.success('Image uploaded successfully!');
            setState((prev) => ({
                ...prev,
                images: [...prev.images, res.data],
                isLoading: false,
                error: null,
            }));
        } catch {
            handleError('Failed to upload image.');
        }
    };

    const fetchAllImagesByUser = async (userId: string, token: string) => {
        try {
            if (!token) return;
            setLoading();
            const res = await requestGetPatientByUserId(userId, token); // Assuming this fetches all images
            if (isErrorResponse(res)) return handleError(res.message);

            setState((prev) => ({
                ...prev,
                images: res.data,
                isLoading: false,
                error: null,
            }));
        } catch (error) {
            handleError('Failed to fetch all images by user.');
        }
    }

    // The below function is used to process the image using the diagnostic service
    // We will just return the image data for now, so we save what we want later on
    const processImage = async (file: File, token: string) => {
        setLoading();
        try {
            const res = await requestDiagnosticProcess(file, token);
            if (res.sucess == false) throw new Error('Failed to process image. Error in the model.');
            return res;
        } catch (error) {
            handleError('Failed to process image.');
        }
    };

    const predictImage = async (file: File, token: string) => {
        setLoading();
        try {
            const res = await requestDiagnosticPredict(file, token);
            if (res.sucess == false) throw new Error('Failed to predict image. Error in the model.');
            return res;
        } catch (error) {
            handleError('Failed to predict image.');
        }
    };

    // we will get smth like this: /api/results/8359209c-0be6-4af7-92d8-e5d5ad7ff13f.png
    // and we will make sure we pass only the id to the backend
    const getImageByPath = async (ImagePath: string, token: string) => {
        if (!token) return;
        setLoading();
        try {
            const imagePathArray = ImagePath.split('/');
            const imageId = imagePathArray[imagePathArray.length - 1];
            const res = await requestGetImageById(imageId, token);
            return res;
        } catch {
            handleError('Failed to fetch image.');
        }
    };

    return {
        images: state.images,
        selectedImage: state.selectedImage,
        isLoading: state.isLoading,
        error: state.error,
        fetchImagesByPatientId,
        fetchAllImagesByUser,
        fetchImageById,
        createPatientImage,
        resetError,
        // For our diagnostic service
        processImage,
        predictImage,
        getImageByPath,
    };
};
