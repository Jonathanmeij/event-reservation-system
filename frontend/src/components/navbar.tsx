import { useContext } from "react";
import { Button } from "./ui/button";
import { AuthContext } from "@/contexts/authContext";
import { Link } from "react-router-dom";

export default function Navbar() {
  const { isAuthenticated } = useContext(AuthContext);

  return (
    <nav className="flex items-center justify-between px-6 py-3">
      <Link to={"/"} className="text-lg font-medium tracking-tight">
        Events
      </Link>
      <div className="flex items-center space-x-3">
        {isAuthenticated && (
          <>
            <Button className="" size={"sm"} variant={"ghost"} asChild>
              <Link to="/login">Sign in</Link>
            </Button>
            <Button className="" size={"sm"} asChild>
              <Link to="/register">Sign up</Link>
            </Button>
          </>
        )}
      </div>
    </nav>
  );
}
