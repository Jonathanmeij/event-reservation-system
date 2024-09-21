import { useContext } from "react";
import { Button } from "./ui/button";
import { AuthContext } from "@/contexts/authContext";
import { Link } from "react-router-dom";
import { SearchInput } from "./ui/input";
import { ChevronDownIcon, ShieldEllipsis, Ticket, User } from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger
} from "@/components/ui/dropdown-menu";
import { InfoCircledIcon } from "@radix-ui/react-icons";

export default function Navbar() {
  const { isAuthenticated } = useContext(AuthContext);

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
          className="max-w-80"
        />
        <div className="flex items-center space-x-3">
          {isAuthenticated && <AccountDropdown />}
        </div>
      </div>
    </nav>
  );
}

function AccountDropdown() {
  const { userData } = useContext(AuthContext);

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button className="" size={"wideIcon"} variant={"ghost"}>
          <User className="size-5" />
          <ChevronDownIcon className="size-3.5" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end">
        <div className="flex items-center py-2">
          <div className="flex items-center justify-center ml-2 rounded-full size-8 bg-zinc-700 text-zinc-300">
            {userData?.firstName?.charAt(0)}
          </div>
          <div className="flex flex-col px-2">
            <span className="text-sm font-medium text-zinc-50">
              {userData?.firstName} {userData?.lastName}
            </span>
            <span className="text-xs tracking-wide text-zinc-300 opacity-60">
              {userData?.email}
            </span>
          </div>
        </div>
        <DropdownMenuSeparator />
        <DropdownMenuItem>
          <Ticket className="mr-2 size-4" />
          Tickets
        </DropdownMenuItem>
        <DropdownMenuItem>
          <User className="mr-2 size-4" />
          Account
        </DropdownMenuItem>
        <DropdownMenuItem>
          <InfoCircledIcon className="mr-2 size-4" />
          Informatie
        </DropdownMenuItem>
        <DropdownMenuSeparator />
        {userData?.role == "admin" && (
          <DropdownMenuItem asChild>
            <Link to="/admin" className="flex items-center gap-2">
              <ShieldEllipsis className="size-4" />
              Admin
            </Link>
          </DropdownMenuItem>
        )}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
