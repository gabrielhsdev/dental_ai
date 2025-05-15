export interface BaseStatus {
    message: string;
    timestamp: Date;
}

export interface SuccessStatus extends BaseStatus {
    success: true;
    data?: any;
}

export interface ErrorStatus extends BaseStatus {
    success: false;
    error: {
        code: string;
        details?: string;
    };
}

export type Status = SuccessStatus | ErrorStatus;

// Helper functions to create status objects
export const createSuccessStatus = (message: string, data?: any): SuccessStatus => ({
    success: true,
    message,
    data,
    timestamp: new Date(),
});

export const createErrorStatus = (message: string, code: string, details?: string): ErrorStatus => ({
    success: false,
    message,
    error: {
        code,
        details,
    },
    timestamp: new Date(),
});