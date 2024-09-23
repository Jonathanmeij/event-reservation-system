import { Event } from "@/api/types";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { useGetEventsWithPlannedEvents } from "@/queries/events";
import { ChevronDown, NotepadTextDashed } from "lucide-react";

export default function EventsOverview() {
  const events = useGetEventsWithPlannedEvents();

  if (events.data) {
    if (events.data.length === 0) {
      return (
        <div className="flex flex-col items-center justify-center w-full h-full gap-1">
          <NotepadTextDashed className="mb-3 size-24 text-zinc-600" />
          <div className="text-3xl font-medium text-zinc-600">No events</div>
          <div className="text-sm text-zinc-400">
            Create an event to get started
          </div>
        </div>
      );
    }

    return (
      <div className="flex flex-col gap-3">
        <h1 className="pageTitle">Events</h1>
        {events.data.map((event) => (
          <EventCard key={event.id} event={event} />
        ))}
      </div>
    );
  }

  if (events.error) {
    return <div>Something went wrong</div>;
  }

  return <EventsSkeleton />;
}

function EventCard({ event }: { event: Event }) {
  return (
    <button className="flex items-center gap-3 p-3 overflow-hidden bg-white shadow-sm rounded-2xl">
      <img
        src={event.images ? event.images[0] : ""}
        alt={event.title}
        className="object-cover w-16 h-16 rounded-lg"
      />
      <div className="flex items-center justify-between w-full">
        <div className="">
          <div className="">{event.title}</div>
          <div className="text-sm text-zinc-400">
            {event.Duration}m - {event.plannedEvents.length} planned
          </div>
        </div>
        <Button variant={"ghost"}>
          <ChevronDown className="size-6 text-zinc-600" />
        </Button>
      </div>
    </button>
  );
}

function EventsSkeleton() {
  return (
    <div className="flex flex-col gap-3 overflow-hidden">
      <h1 className="pageTitle">Events</h1>
      {[...Array(8)].map(() => (
        <Skeleton className="w-full h-32 rounded-lg" />
      ))}
    </div>
  );
}
