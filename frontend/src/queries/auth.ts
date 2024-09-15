import { login } from "@/api/auth";
import { ErrorResponse } from "@/api/error";
import { LoginRequest } from "@/api/types";
import { AxiosError } from "axios";
import { useMutation } from "react-query";

export default function useLogin() {
  return useMutation<
    { token: string },
    AxiosError<ErrorResponse>,
    LoginRequest
  >({
    mutationFn: (loginRequest: LoginRequest) => login(loginRequest),
    onSuccess: (res) => {
      //retrieve jwt
      console.log(res);
    }
  });
}
