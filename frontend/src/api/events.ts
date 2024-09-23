import { axios } from "./axios";
import { CreateEventRequest, Event, UpdateEventRequest } from "./types";

const ENDPOINT = "/events";

export async function getEvents() {
  const { data } = await axios.get<Event[]>(ENDPOINT);
  return data;
}

export async function getEvent(id: string) {
  const { data } = await axios.get(`${ENDPOINT}/${id}`);
  return data;
}

export async function createEvent(event: CreateEventRequest) {
  const { data } = await axios.post(ENDPOINT, event);
  return data;
}

export async function updateEvent(event: UpdateEventRequest) {
  const { data } = await axios.put(`${ENDPOINT}`, event);
  return data;
}

export async function deleteEvent(id: string) {
  await axios.delete(`${ENDPOINT}/${id}`);
}

export async function getEventsWithPlannedEvents() {
  const { data } = await axios.get(`${ENDPOINT}-with-planned`);
  return data;
}
