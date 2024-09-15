import { login, register } from "@/api/auth";
import { ErrorResponse } from "@/api/error";
import { LoginRequest, RegisterRequest, TokenResponse } from "@/api/types";
import { AxiosError } from "axios";
import { useMutation } from "react-query";

export function useLogin(onSucces?: () => void) {
  return useMutation<TokenResponse, AxiosError<ErrorResponse>, LoginRequest>({
    mutationFn: (loginRequest: LoginRequest) => login(loginRequest),
    onSuccess: (res) => {
      localStorage.setItem("token", res.token);
      onSucces?.();
    }
  });
}

export function useRegister(onSucces?: () => void) {
  return useMutation<TokenResponse, AxiosError<ErrorResponse>, RegisterRequest>(
    {
      mutationFn: (registerRequest: RegisterRequest) =>
        register(registerRequest),
      onSuccess: (res) => {
        localStorage.setItem("token", res.token);
        onSucces?.();
      }
    }
  );
}
