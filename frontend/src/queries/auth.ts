import { login } from "@/api/auth";
import { ErrorResponse } from "@/api/error";
import { LoginRequest, TokenResponse } from "@/api/types";
import { AxiosError } from "axios";
import { useMutation } from "react-query";

export default function useLogin(onSucces?: () => void) {
  return useMutation<TokenResponse, AxiosError<ErrorResponse>, LoginRequest>({
    mutationFn: (loginRequest: LoginRequest) => login(loginRequest),
    onSuccess: (res) => {
      localStorage.setItem("token", res.token);
      onSucces?.();
    }
  });
}
