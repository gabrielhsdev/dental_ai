import { BaseErrorResponseInterface, BaseResponseInterface, UserInterface } from "@/common/commonInterfaces";
import { postRequest } from "./baseRequests";
import { ENDPOINTS } from "@/common/constants";

export interface LoginData {
    token: string;
}

export async function requestLogin(
    email: string,
    password: string
): Promise<BaseResponseInterface<LoginData> | BaseErrorResponseInterface> {
    return await postRequest<LoginData>(`${ENDPOINTS.AUTH}/login`, {
        email,
        password,
    });
}

export async function requestMe(
    token: string
): Promise<BaseResponseInterface<UserInterface> | BaseErrorResponseInterface> {
    return await postRequest<UserInterface>(`${ENDPOINTS.AUTH}/me`, {}, token);
}

export async function requestRegister(
    userName: string,
    email: string,
    password: string,
    firstName: string,
    lastName: string,
): Promise<BaseResponseInterface<UserInterface> | BaseErrorResponseInterface> {
    return await postRequest<UserInterface>(`${ENDPOINTS.AUTH}/register`, {
        userName,
        email,
        password,
        firstName,
        lastName,
    });
}
