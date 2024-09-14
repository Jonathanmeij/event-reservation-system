import { useGetEvents } from "./pages/events/api/use-get-events";

function App() {
  const { data, isLoading } = useGetEvents();

  if (isLoading) {
    return (
      <>
        <div>Loading...</div>
      </>
    );
  }

  return (
    <>
      <div>{data?.map((event) => <div key={event.id}>{event.title}</div>)}</div>
    </>
  );
}

export default App;
