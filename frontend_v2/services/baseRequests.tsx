import { BaseErrorResponseInterface, BaseResponseInterface } from '@/common/commonInterfaces';
import axios from 'axios';

const buildHeaders = (token?: string): Record<string, string> => ({
    'Content-Type': 'application/json',
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
});

// GET Request - used for our GO microservices
export async function getRequest<T>(
    url: string,
    params?: Record<string, any>,
    token?: string
): Promise<BaseResponseInterface<T>> {
    try {
        const headers = token ? { Authorization: `Bearer ${token}` } : {};
        const response = await axios.get<BaseResponseInterface<T>>(url, {
            params,
            headers,
        });
        return response.data;
    } catch (error) {
        console.error('GET request error:', error);
        throw error;
    }
}

// POST Request - used for our GO microservices
export async function postRequest<T>(
    url: string,
    data?: Record<string, any>,
    token?: string
): Promise<BaseResponseInterface<T> | BaseErrorResponseInterface> {
    try {
        const response = await axios.post<BaseResponseInterface<T>>(url, data, {
            headers: buildHeaders(token),
        });
        return response.data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.error('POST request error:', error.response?.data);
            return error.response?.data as BaseErrorResponseInterface;
        }
        throw error;
    }
}

// POST Request without BaseResponseInterface or BaseErrorResponseInterface
export async function postRequestWithoutResponse<T>(
    url: string,
    data?: Record<string, any> | FormData,
    token?: string
): Promise<T> {
    try {
        const isFormData = data instanceof FormData;

        const response = await axios.post<T>(url, data, {
            headers: {
                ...buildHeadersv2(token),
                ...(isFormData ? {} : { 'Content-Type': 'application/json' }),
            },
        });

        return response.data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.error('POST request error:', error.response?.data);
            throw error;
        }
        throw error;
    }
}


// GET Request without BaseResponseInterface or BaseErrorResponseInterface
export async function getRequestWithoutResponse<T>(
    url: string,
    params?: Record<string, any>,
    token?: string
): Promise<T> {
    try {
        const response = await axios.get<T>(url, {
            params,
            headers: buildHeadersv2(token),
            responseType: 'blob', // ✅ CRITICAL FIX
        });
        return response.data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.error('GET request error:', error.response?.data);
            throw error;
        }
        throw error;
    }
}

export async function getBlobRequest(
    url: string,
    token?: string
): Promise<Blob> {
    try {
        const response = await axios.get<Blob>(url, {
            headers: buildHeadersv2(token),
            responseType: "blob", // 👈 This is key
        });
        return response.data;
    } catch (error) {
        if (axios.isAxiosError(error)) {
            console.error("Blob request error:", error.response?.data);
            throw error;
        }
        throw error;
    }
}

export function buildHeadersv2(token?: string): Record<string, string> {
    return {
        ...(token ? { Authorization: `Bearer ${token}` } : {}),
        // DO NOT include Content-Type here
    };
}
