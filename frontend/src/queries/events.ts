import { getEvents } from "@/api/events";
import { Event } from "@/api/types";
import { useQuery } from "react-query";

const QUERY_KEY = "events";

export function useGetEvents() {
  return useQuery<Event[], Error>({
    queryFn: getEvents,
    queryKey: QUERY_KEY
  });
}
