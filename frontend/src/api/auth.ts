import { axios } from "./axios";
import { LoginRequest, RegisterRequest } from "./types";

const ENDPOINT = "/account";

export const login = async (loginRequest: LoginRequest) => {
  const { data } = await axios.post<{ token: string }>(
    `${ENDPOINT}/login`,
    loginRequest
  );
  return data;
};

export const register = async (registerRequest: RegisterRequest) => {
  const { data } = await axios.post(`${ENDPOINT}/register`, registerRequest);
  return data;
};
