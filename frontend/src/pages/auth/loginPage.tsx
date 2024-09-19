import { Button } from "@/components/ui/button";
import {
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Link, useNavigate } from "react-router-dom";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage
} from "@/components/ui/form";
import { useLogin } from "@/queries/auth";
import { LoginRequest, TokenResponse } from "@/api/types";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { CircleX } from "lucide-react";
import { useContext } from "react";
import { AuthContext } from "@/contexts/authContext";

const formSchema = z.object({
  email: z.string().email(),
  password: z.string().min(6)
});

export default function LoginPage() {
  const navigate = useNavigate();
  const { login } = useContext(AuthContext);
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: ""
    }
  });

  const { mutate, isLoading, error } = useLogin(onSucces);

  function onSucces(res: TokenResponse) {
    login(res);
    navigate("/");
  }

  function onSubmit(values: z.infer<typeof formSchema>) {
    const loginRequest: LoginRequest = {
      email: values.email,
      password: values.password
    };

    mutate(loginRequest);
  }

  return (
    <div className="flex items-center justify-center flex-1 w-full bg-zinc-50">
      <div className="max-w-96">
        <CardHeader>
          <CardTitle className="text-2xl">Login</CardTitle>
          <CardDescription>
            Enter your email below to login to your account
          </CardDescription>
        </CardHeader>
        <CardContent>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="grid gap-4">
              {error && (
                <div className="text-sm text-center text-red-500">
                  <Alert variant="destructive">
                    <CircleX className="size-5" />
                    <AlertTitle>Error</AlertTitle>
                    <AlertDescription>
                      {error.response?.data.error}
                    </AlertDescription>
                  </Alert>
                </div>
              )}
              <FormField
                control={form.control}
                name="email"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Email</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="password"
                render={({ field }) => (
                  <FormItem>
                    <div className="flex justify-between w-full">
                      <FormLabel>Password</FormLabel>
                      <Link
                        to="#"
                        className="inline-block ml-auto text-sm underline"
                      >
                        Forgot your password?
                      </Link>
                    </div>
                    <FormControl>
                      <Input type="password" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <Button isLoading={isLoading} type="submit" className="w-full">
                Login
              </Button>
            </form>
          </Form>
          <div className="mt-4 text-sm text-center">
            Don&apos;t have an account?{" "}
            <Link to="/register" className="underline">
              Sign up
            </Link>
          </div>
        </CardContent>
      </div>
    </div>
  );
}
