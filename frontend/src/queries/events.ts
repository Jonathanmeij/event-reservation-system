import { getEvents, getEventsWithPlannedEvents } from "@/api/events";
import { Event } from "@/api/types";
import { useQuery } from "react-query";

const QUERY_KEY = "events";

export function useGetEvents() {
  return useQuery<Event[], Error>({
    queryFn: getEvents,
    queryKey: QUERY_KEY
  });
}

export function useGetEventsWithPlannedEvents() {
  return useQuery<Event[], Error>({
    queryFn: getEventsWithPlannedEvents,
    queryKey: QUERY_KEY
  });
}
