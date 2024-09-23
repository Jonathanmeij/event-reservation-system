import { getLocations } from "@/api/locations";
import { Location } from "@/api/types";
import { useQuery } from "react-query";

const QUERY_KEY = "locatio";

export function useGetLocations() {
  return useQuery<Location[], Error>({
    queryFn: getLocations,
    queryKey: QUERY_KEY
  });
}
