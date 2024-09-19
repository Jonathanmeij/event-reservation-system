import { Event } from "@/api/types";
import { Card } from "@/components/ui/card";

export default function EventCard({ event }: { event: Event }) {
  return (
    <Card className="relative overflow-hidden">
      <img
        src={event.imageUrl}
        alt={event.title}
        className="object-cover w-full h-48"
      />
      <div
        className="absolute inset-0 border backdrop-blur-3xl"
        style={{
          maskImage:
            "linear-gradient(0deg, rgba(0,0,0,1) 0%, rgba(0,212,255,0) 30%)"
        }}
      ></div>
      <div className="absolute bottom-0 w-full p-4 pt-8 bg-gradient-to-t from-black/50">
        <h3 className="text-xl font-medium text-white">{event.title}</h3>
      </div>
    </Card>
  );
}
