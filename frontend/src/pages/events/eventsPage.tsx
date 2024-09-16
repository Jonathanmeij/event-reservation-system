import { AuthContext } from "@/contexts/authContext";
import { useGetEvents } from "@/queries/events";
import { useContext } from "react";

export default function EventsPage() {
  const { data, isLoading, isError } = useGetEvents();
  const { userData, isAuthenticated } = useContext(AuthContext);

  return (
    <div>
      {isAuthenticated && (
        <div>
          <div>Welcome {userData?.firstName}</div>
        </div>
      )}
      {isLoading && <div>Loading...</div>}
      {isError && <div>Error</div>}
      {data && data.map((event) => <div key={event.id}>{event.title}</div>)}
    </div>
  );
}
