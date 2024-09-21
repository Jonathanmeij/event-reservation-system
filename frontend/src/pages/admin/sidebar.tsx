import { CalendarDays, MapPin, Tickets, Users } from "lucide-react";
import { Link, useLocation } from "react-router-dom";
import { Separator } from "@/components/ui/separator";

export default function SideBar() {
  return (
    <div className="h-full p-2 bg-white min-w-64 rounded-xl">
      <h2 className="px-3 mt-3 font-medium">Admin Panel</h2>
      <Separator className="my-3" />
      <div className="flex flex-col gap-3">
        <SidebarLink to={"/admin/events"}>
          <CalendarDays className="size-4" />
          Events
        </SidebarLink>
        <SidebarLink to={"/admin/locations"}>
          <MapPin className="size-4" />
          Locations
        </SidebarLink>

        <SidebarLink to={"/admin/tickets"}>
          <Tickets className="size-4" />
          Tickets
        </SidebarLink>
        <SidebarLink to={"/admin/users"}>
          <Users className="size-4" />
          Users
        </SidebarLink>
      </div>
    </div>
  );
}

function SidebarLink({
  to,
  children
}: {
  to: string;
  children: React.ReactNode;
}) {
  //get route
  const location = useLocation();
  const { pathname } = location;
  const isActive = pathname === to;

  return (
    <Link
      to={to}
      className={`flex h-9 items-center gap-2 rounded-md px-3 text-sm ${isActive ? "bg-orange-500 text-white" : "text-zinc-500 hover:bg-zinc-100"}`}
    >
      {children}
    </Link>
  );
}
