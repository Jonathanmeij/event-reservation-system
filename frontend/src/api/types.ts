// Code generated by tygo. DO NOT EDIT.

//////////
// source: types.go

/**
 * event
 */
export interface Event {
  id: number /* uint */;
  title: string;
  description: string;
  images: string[];
  cast: string[];
  directors: string[];
  genres: string[];
  Duration: number /* int */;
  createdAt: Date ;
  plannedEvents: PlannedEvent[];
}
export interface CreateEventRequest {
  title: string;
  description: string;
  imageUrl: string[];
  cast: string[];
  directors: string[];
  genres: string[];
  Duration: number /* int */;
  date: Date ;
}
export interface UpdateEventRequest {
  title: string;
  description: string;
  images: string[];
  cast: string[];
  directors: string[];
  genres: string[];
  duration: number /* int */;
}
export interface PlannedEvent {
  id: number /* uint */;
  eventId: number /* int */;
  locationId: number /* int */;
  date: Date ;
  location: Location;
}
export interface CreatePlannedEventRequest {
  eventId: number /* int */;
  locationId: number /* int */;
  date: Date ;
}
/**
 * Location
 */
export interface Location {
  id: number /* int */;
  name: string;
  amountOfPeople: number /* int */;
}
export interface CreateLocationRequest {
  name: string;
  amountOfPeople: number /* int */;
}
/**
 * users
 */
export interface User {
  id: number /* int */;
  firstName: string;
  lastName: string;
  email: string;
  role: 'admin' | 'user';
  createdAt: Date ;
}
export interface RegisterRequest {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}
export interface LoginRequest {
  email: string;
  password: string;
}
export interface TokenResponse {
  firstName: string;
  lastName: string;
  email: string;
  role: string;
  token: string;
}
