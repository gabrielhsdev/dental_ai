import { DiagnosticInterfaceProcessDetection } from "@/common/commonInterfaces";
import { postRequestWithoutResponse, getRequestWithoutResponse, getBlobRequest } from "./baseRequests";
import { ENDPOINTS } from "@/common/constants";

export async function requestDiagnosticProcess(
    file: File,
    token: string
): Promise<DiagnosticInterfaceProcessDetection> {
    try {
        const formData = new FormData();
        formData.append("file", file);

        return await postRequestWithoutResponse<DiagnosticInterfaceProcessDetection>(
            `${ENDPOINTS.DIAGNOSTICS}/process/detection`,
            formData,
            token
        );

    } catch (error) {
        console.error("Error requesting diagnostic process:", error);
        throw error;
    }
}

export async function requestGetImageById(
    imageId: string,
    token: string
): Promise<any> { // Not sure what this returns since on postman is just a image
    const url = `${ENDPOINTS.DIAGNOSTICS}/results/${imageId}`;
    try {
        return await getBlobRequest(url, token);
    } catch (error) {
        console.error("Error requesting diagnostic process by ID:", error);
        throw error;
    }
}
