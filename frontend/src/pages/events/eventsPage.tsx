import { useGetEvents } from "@/queries/events";

export default function EventsPage() {
  const { data, isLoading, isError } = useGetEvents();

  return (
    <div>
      {isLoading && <div>Loading...</div>}
      {isError && <div>Error</div>}
      {data && data.map((event) => <div key={event.id}>{event.title}</div>)}
    </div>
  );
}
