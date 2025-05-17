import { BaseErrorResponseInterface, BaseResponseInterface } from '@/common/commonInterfaces';
import axios from 'axios';

const buildHeaders = (token?: string): Record<string, string> => ({
    'Content-Type': 'application/json',
    ...(token ? { Authorization: `Bearer ${token}` } : {}),
});

// GET Request
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

// POST Request
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
