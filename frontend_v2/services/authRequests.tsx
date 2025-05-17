import { BaseErrorResponseInterface, BaseResponseInterface, UserInterface } from "@/common/commonInterfaces";
import { getRequest, postRequest } from "./baseRequests";
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
