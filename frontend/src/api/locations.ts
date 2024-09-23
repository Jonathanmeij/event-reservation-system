import { axios } from "./axios";
import { Location } from "./types";

const ENDPOINT = "/locations";

export async function getLocations() {
  const { data } = await axios.get<Location[]>(ENDPOINT);
  return data;
}

export async function getLocation(id: string) {
  const { data } = await axios.get(`${ENDPOINT}/${id}`);
  return data;
}
