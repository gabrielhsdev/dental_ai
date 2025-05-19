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

// ========== PATIENT INTERFACES ==========
// ========== PATIENT INTERFACES ==========
export interface PatientInterface {
    id: string;
    userId: string;
    firstName: string;
    lastName: string;
    dateOfBirth: string;
    gender: string;
    phoneNumber: string;
    email: string;
    notes: string;
    createdAt: string;
    updatedAt: string;
}

// ========== PATIENT IMAGE INTERFACES ==========
// ========== PATIENT IMAGE INTERFACES ==========
export interface PatientImageInterface {
    id: string;
    patientId: string;
    imageData: string;
    fileType: string;
    description: string;
    uploadedAt: string;
    createdAt: string;
    updatedAt: string;
}

// ========== DIAGNOSTIC INTERFACES ==========
// ========== DIAGNOSTIC INTERFACES ==========
export interface DiagnosticInterfaceProcessDetection {
    model_id: string;
    result_image: string;
    sucess: boolean;
}