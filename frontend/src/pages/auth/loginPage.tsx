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
import useLogin from "@/queries/auth";
import { LoginRequest } from "@/api/types";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { CircleX } from "lucide-react";

const formSchema = z.object({
  email: z.string().email(),
  password: z.string().min(6)
});

export default function LoginPage() {
  const navigate = useNavigate();
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: ""
    }
  });

  const { mutate, isLoading, error } = useLogin(onSucces);

  function onSucces() {
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
    <div className="flex min-h-screen w-full items-center justify-center bg-zinc-50">
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
                <div className="text-center text-sm text-red-500">
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
                      <Input placeholder="Email" {...field} />
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
                    <div className="flex w-full justify-between">
                      <FormLabel>Password</FormLabel>
                      <Link
                        to="#"
                        className="ml-auto inline-block text-sm underline"
                      >
                        Forgot your password?
                      </Link>
                    </div>
                    <FormControl>
                      <Input
                        placeholder="password"
                        type="password"
                        {...field}
                      />
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
          <div className="mt-4 text-center text-sm">
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
