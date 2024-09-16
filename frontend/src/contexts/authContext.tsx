import { TokenResponse } from "@/api/types";
import { jwtDecode } from "jwt-decode";
import { createContext, useEffect, useState } from "react";

interface AuthContextType {
  isAuthenticated: boolean;
  isLoading: boolean;
  userData: TokenResponse | null;
  login: (res: TokenResponse) => void;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextType>({
  isAuthenticated: false,
  isLoading: true,
  login: () => {},
  logout: () => {},
  userData: {
    firstName: "",
    lastName: "",
    email: "",
    role: "",
    token: ""
  }
});

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [userData, setUserData] = useState<TokenResponse>({
    firstName: "",
    lastName: "",
    email: "",
    role: "",
    token: ""
  });

  useEffect(() => {
    const storedUserData = localStorage.getItem("userData");

    if (storedUserData) {
      const userData = JSON.parse(storedUserData) as TokenResponse;
      const isExpired = isTokenExpired(userData.token);
      if (!isExpired) {
        setUserData(userData);
        setIsAuthenticated(true);
      }
    }
    setIsLoading(false);
  }, []);

  const login = (res: TokenResponse) => {
    localStorage.setItem("userData", JSON.stringify(res));
    setUserData(res);
    setIsAuthenticated(true);
  };

  const logout = () => {
    localStorage.removeItem("userData");
    setIsAuthenticated(false);
    setUserData({
      firstName: "",
      lastName: "",
      email: "",
      role: "",
      token: ""
    });
  };

  return (
    <AuthContext.Provider
      value={{ isAuthenticated, isLoading, userData, login, logout }}
    >
      {children}
    </AuthContext.Provider>
  );
};

function isTokenExpired(token: string | null): boolean {
  if (!token) {
    return true;
  }
  try {
    const decoded = jwtDecode(token);

    const currentTime = Date.now() / 1000;

    if (decoded.exp && decoded.exp < currentTime) {
      return true;
    } else {
      return false;
    }
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
  } catch (error) {
    return true;
  }
}
