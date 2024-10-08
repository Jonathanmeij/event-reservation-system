import { Event } from "@/api/types";
import { Card } from "@/components/ui/card";

export default function EventCard({ event }: { event: Event }) {
  return (
    <Card className="relative overflow-hidden group hover:cursor-pointer">
      <img
        src={event.images ? event.images[0] : ""}
        alt={event.title}
        className="object-cover w-full transition-transform duration-300 ease-out transform h-52 group-hover:scale-105"
      />
      <div className="absolute bottom-0 w-full p-4 pt-8 bg-gradient-to-t from-black/50">
        <span className="text-sm leading-none text-white/60">Upcoming</span>
        <h3 className="text-2xl font-medium leading-none text-white">
          {event.title}
        </h3>
      </div>
    </Card>
  );
}
