import { Event } from "@/api/types";
import { Card } from "@/components/ui/card";

export default function EventCard({ event }: { event: Event }) {
  return (
    <Card className="relative overflow-hidden group hover:cursor-pointer">
      <img
        src={event.imageUrl}
        alt={event.title}
        className="object-cover w-full h-48 transition-transform duration-300 ease-out transform group-hover:scale-105"
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
