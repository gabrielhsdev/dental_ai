// ========== RESPONSE REQUESTS INTERFACES ==========
// ========== RESPONSE REQUESTS INTERFACES ==========
// ========== RESPONSE REQUESTS INTERFACES ==========
export interface BaseResponseInterface<T = unknown> {
    status: number;
    message: string;
    data: T;
    timestamp: string;
}

export interface BaseErrorResponseInterface {
    status: number;
    message: string;
    error: string;
    timestamp: string;
}

export function isErrorResponse(
    response: BaseResponseInterface | BaseErrorResponseInterface
): response is BaseErrorResponseInterface {
    return 'error' in response;
}

// ========== USER INTERFACES ==========
// ========== USER INTERFACES ==========
export interface UserInterface {
    id: string;
    username: string;
    email: string;
    password: string;
    firstName: string;
    lastName: string;
    createdAt: string;
    updatedAt: string;
}
