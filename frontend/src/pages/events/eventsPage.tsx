import { useGetEvents } from "@/queries/events";
import EventCard from "./eventCard";

export default function EventsPage() {
  const events = useGetEvents();

  if (events.data) {
    return (
      <div className="flex-1 p-6 bg-zinc-100">
        <div className="grid gap-4 mx-auto max-w-screen-2xl sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
          {events.data.map((event) => (
            <EventCard key={event.id} event={event} />
          ))}
        </div>

        <div className="wrapper"></div>
      </div>
    );
  }

  if (events.error) {
    return <div>Something went wrong</div>;
  }

  return <div>Loading...</div>;
}
