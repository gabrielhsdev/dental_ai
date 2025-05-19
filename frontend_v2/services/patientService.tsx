import { BaseErrorResponseInterface, BaseResponseInterface, PatientInterface } from "@/common/commonInterfaces";
import { ENDPOINTS } from "@/common/constants";
import { getRequest, postRequest } from "./baseRequests";

export async function requestCreatePatient(
    firstName: string,
    lastName: string,
    dateOfBirth: string,
    gender: string,
    phoneNumber: string,
    email: string,
    notes: string,
    token: string,
): Promise<BaseResponseInterface<PatientInterface> | BaseErrorResponseInterface> {
    return await postRequest<PatientInterface>(`${ENDPOINTS.DB}/patients/`, {
        firstName,
        lastName,
        dateOfBirth,
        gender,
        phoneNumber,
        email,
        notes,
    }, token);
}

export async function requestGetPatientByUserId(
    token: string,
): Promise<BaseResponseInterface<PatientInterface[]> | BaseErrorResponseInterface> {
    return await getRequest<PatientInterface[]>(`${ENDPOINTS.DB}/patients/user`, {}, token);
}

export async function requestGetPatientById(
    patientId: string,
    token: string,
): Promise<BaseResponseInterface<PatientInterface> | BaseErrorResponseInterface> {
    return await getRequest<PatientInterface>(`${ENDPOINTS.DB}/patients/${patientId}`, {}, token);
}
