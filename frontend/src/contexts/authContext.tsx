import { log } from "console";
import { jwtDecode } from "jwt-decode";
import { createContext, useEffect, useState } from "react";
import { set } from "react-hook-form";

interface AuthContextType {
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (token: string) => void;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextType>({
  isAuthenticated: false,
  isLoading: true,
  login: () => {},
  logout: () => {}
});

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const token = localStorage.getItem("token");

    if (token && !isTokenExpired(token)) {
      setIsAuthenticated(true);
      setIsLoading(false);
    }
  }, []);

  const login = (token: string) => {
    localStorage.setItem("token", token);
    setIsAuthenticated(true);
  };

  const logout = () => {
    localStorage.removeItem("token");
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, isLoading, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

function isTokenExpired(token: string | null): boolean {
  if (!token) {
    return true;
  }
  try {
    // Decode the token to get its payload
    const decoded = jwtDecode(token);

    // Get the current time in seconds since the Epoch
    const currentTime = Date.now() / 1000;

    // Compare the expiration time with the current time
    if (decoded.exp && decoded.exp < currentTime) {
      // Token has expired
      return true;
    } else {
      // Token is valid
      return false;
    }
  } catch (error) {
    // Handle errors (e.g., invalid token format)
    console.error("Invalid token:", error);
    return true; // Treat invalid token as expired
  }
}
