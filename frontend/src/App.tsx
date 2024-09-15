import { BrowserRouter, Route, Routes } from "react-router-dom";
import LoginPage from "./pages/auth/loginPage";
import RegisterPage from "./pages/auth/RegisterPage";
import EventsPage from "./pages/events/eventsPage";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/" element={<EventsPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
