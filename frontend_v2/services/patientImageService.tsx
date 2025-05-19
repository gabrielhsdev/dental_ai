import { BaseErrorResponseInterface, BaseResponseInterface, PatientImageInterface } from "@/common/commonInterfaces";
import { ENDPOINTS } from "@/common/constants";
import { getRequest, postRequest } from "./baseRequests";

export async function requestCreatePatientImage(
    patientId: string,
    imageData: string,
    fileType: string,
    description: string,
    token: string
): Promise<BaseResponseInterface<PatientImageInterface> | BaseErrorResponseInterface> {
    // Placehoplder for uploadedAt / createdAt and updatedAt
    console.log("Creating patient image with data:", {
        patientId,
        imageData,
        fileType,
        description,
    });
    return await postRequest<PatientImageInterface>(`${ENDPOINTS.DB}/patientsImages/`, {
        patientId,
        imageData,
        fileType,
        description,
        uploadedAt: new Date().toISOString(),
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
    }, token);
}

export async function requestGetPatientImageById(
    imageId: string,
    token: string
): Promise<BaseResponseInterface<PatientImageInterface> | BaseErrorResponseInterface> {
    return await getRequest<PatientImageInterface>(`${ENDPOINTS.DB}/patientsImages/${imageId}`, {}, token);
}

export async function requestGetPatientImagesByPatientId(
    patientId: string,
    token: string
): Promise<BaseResponseInterface<PatientImageInterface[]> | BaseErrorResponseInterface> {
    return await getRequest<PatientImageInterface[]>(`${ENDPOINTS.DB}/patientsImages/patient/${patientId}`, {}, token);
}


