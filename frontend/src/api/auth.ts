import { axios } from "./axios";
import { LoginRequest, RegisterRequest, TokenResponse } from "./types";

const ENDPOINT = "/account";

export const login = async (loginRequest: LoginRequest) => {
  const { data } = await axios.post<TokenResponse>(
    `${ENDPOINT}/login`,
    loginRequest
  );
  return data;
};

export const register = async (registerRequest: RegisterRequest) => {
  const { data } = await axios.post<TokenResponse>(
    `${ENDPOINT}/register`,
    registerRequest
  );
  return data;
};
