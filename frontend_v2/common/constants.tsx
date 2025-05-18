// Constants
export const API_BASE_URL = 'http://localhost';

export const ENDPOINTS = {
    AUTH: `${API_BASE_URL}/auth`,
    DB: `${API_BASE_URL}/db`,
    DIAGNOSTICS: `${API_BASE_URL}/diagnostics`,
} as const;

export const LOCAL_STORAGE_KEYS = {
    TOKEN: 'token'
}