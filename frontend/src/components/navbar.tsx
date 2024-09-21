import { useContext } from "react";
import { Button } from "./ui/button";
import { AuthContext } from "@/contexts/authContext";
import { Link } from "react-router-dom";
import { SearchInput } from "./ui/input";

export default function Navbar() {
  const { isAuthenticated, userData } = useContext(AuthContext);

  return (
    <nav className="w-full px-6 py-3 bg-zinc-100">
      <div className="flex items-center justify-between w-full mx-auto max-w-screen-2xl">
        <div className="flex items-center gap-8">
          <Link to={"/"} className="text-xl font-semibold tracking-tight">
            Brand
          </Link>
          <div className="flex gap-6">
            <Link to={"/"} className="text-sm text-zinc-500">
              Popular
            </Link>
            <Link to={"/"} className="text-sm text-zinc-500">
              New
            </Link>
          </div>
        </div>
        <SearchInput
          variant={"secondary"}
          placeholder="Search movie or genre"
          className="max-w-72"
        />
        <div className="flex items-center space-x-3">
          {!isAuthenticated && (
            <>
              <Button className="" size={"sm"} variant={"ghost"} asChild>
                <Link to="/login">Sign in</Link>
              </Button>
              <Button className="" size={"sm"} asChild>
                <Link to="/register">Sign up</Link>
              </Button>
            </>
          )}
          {isAuthenticated && userData?.role == "admin" && (
            <Button className="" size={"sm"} variant={"ghost"} asChild>
              <Link to="/admin">Admin</Link>
            </Button>
          )}
          {isAuthenticated && (
            <Button className="" size={"sm"} variant={"ghost"} asChild>
              <Link to="/account">account</Link>
            </Button>
          )}
        </div>
      </div>
    </nav>
  );
}
