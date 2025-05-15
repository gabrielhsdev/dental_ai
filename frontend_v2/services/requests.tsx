import axios from "axios";

const API_BASE_URL = 'http://localhost';

export const ENDPOINTS = {
    AUTH: `${API_BASE_URL}/auth`,
    DB: `${API_BASE_URL}/db`,
    DIAGNOSTICS: `${API_BASE_URL}/diagnostics`,
} as const;


export async function getRequest<T>(url: string, params?: Record<string, any>, token?: string) {
    try {
        const headers = token ? { Authorization: `Bearer ${token}` } : {};
        const response = await axios.get<T>(url, {
            params,
            headers
        });
        return response.data;
    } catch (error) {
        console.error('Error making GET request:', error);
        throw error;
    }
}

export async function postRequest<T>(url: string, data?: Record<string, any>, token?: string): Promise<T> {
    try {
        const headers: Record<string, string> = {
            'Content-Type': 'application/json',
            ...(token ? { Authorization: `Bearer ${token}` } : {})
        };

        const response = await fetch(url, {
            method: 'POST',
            headers,
            body: JSON.stringify(data)
        });

        const contentType = response.headers.get('content-type');
        const isJson = contentType && contentType.includes('application/json');
        const responseData = isJson ? await response.json() : await response.text();

        if (!response.ok) {
            console.error('Fetch Error Details:', {
                status: response.status,
                statusText: response.statusText,
                serverError: responseData,
                requestURL: url,
                requestData: data
            });
            throw new Error(`Request failed with status ${response.status}`);
        }

        return responseData as T;
    } catch (error) {
        console.error('Unexpected error:', error);
        throw error;
    }
}

// For GO REST API microservices Interfaces
export interface BaseResponse {
    status: number;
    message: string;
    data?: any;
    timestamp: string;
}

// LOGIN 
export interface LoginResponse extends BaseResponse {
    data: {
        token: string;
    }
}

export async function requestLogin(email: string, password: string): Promise<LoginResponse> {
    const response = await postRequest<LoginResponse>(ENDPOINTS.AUTH + '/login', {
        email: email,
        password: password
    });
    return response;
}

// ME
export interface MeResponse extends BaseResponse {
    data: {
        id: string;
        username: string;
        email: string;
        password: string;
        firstName: string;
        lastName: string;
        createdAt: string;
        updatedAt: string;
    }
}

export async function requestMe(token: string): Promise<MeResponse> {
    const response = await getRequest<MeResponse>(ENDPOINTS.AUTH + '/me', {}, token);
    return response;
}