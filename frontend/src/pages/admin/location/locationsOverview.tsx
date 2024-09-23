import { Skeleton } from "@/components/ui/skeleton";
import { useGetLocations } from "@/queries/location";
import { Pin } from "lucide-react";

export default function LocationsOverview() {
  const locations = useGetLocations();

  if (locations.data) {
    if (locations.data.length === 0) {
      return (
        <div className="flex flex-col items-center justify-center w-full h-full gap-1">
          <Pin className="mb-3 size-24 text-zinc-600" />
          <div className="text-3xl font-medium text-zinc-600">No Locations</div>
          <div className="text-sm text-zinc-400">
            Create an location to get started
          </div>
        </div>
      );
    }

    return (
      <div className="flex flex-col gap-3">
        <h1 className="pageTitle">Locations</h1>
        {locations.data.map((location) => (
          <div key={location.id}> {location.name} </div>
        ))}
      </div>
    );
  }

  if (locations.error) {
    return <div>Something went wrong</div>;
  }

  return <LocationsSkeleton />;
}

function LocationsSkeleton() {
  return (
    <div className="flex flex-col gap-3">
      <h1 className="pageTitle">Locations</h1>
      <div className="grid grid-cols-1 gap-3 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
        {[...Array(8)].map((_, index) => (
          <Skeleton key={index} className="" />
        ))}
      </div>
    </div>
  );
}
