import { CalendarDays, MapPin, Tickets, Users } from "lucide-react";
import { Link, useLocation } from "react-router-dom";

export default function SideBar({ className }: { className?: string }) {
  return (
    <div className="h-full p-6 bg-white min-w-64 rounded-xl">
      <h2 className="mb-4 font-medium">Admin Panel</h2>
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
      className={`flex items-center gap-2 ${isActive ? "text-blue-500" : "text-zinc-500"}`}
    >
      {children}
    </Link>
  );
}
