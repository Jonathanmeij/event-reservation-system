import {
  BrowserRouter,
  Navigate,
  Outlet,
  Route,
  Routes
} from "react-router-dom";
import LoginPage from "./pages/auth/loginPage";
import RegisterPage from "./pages/auth/RegisterPage";
import EventsPage from "./pages/events/eventsPage";
import { useContext } from "react";
import { AuthContext } from "./contexts/authContext";
import Spinner from "./components/ui/spinner";
import AdminPage from "./pages/admin/adminPage";
import Navbar from "./components/navbar";

const PrivateRoute = ({ role }: { role?: string }) => {
  const { isAuthenticated, isLoading, userData } = useContext(AuthContext);

  if (isLoading) {
    return (
      <div className="flex items-center justify-center w-screen h-dvh">
        <Spinner />
      </div>
    );
  }

  const isAllowed = role ? userData?.role === role : true;

  if (!isAllowed) {
    return (
      <div>
        <div>Unauthorized</div>
      </div>
    );
  }

  return isAuthenticated ? <Outlet /> : <Navigate to="/login" replace />;
};

function App() {
  return (
    <div className="flex flex-col min-h-screen">
      <BrowserRouter>
        <Navbar />
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/" element={<EventsPage />} />
          <Route element={<PrivateRoute role="admin" />}>
            <Route path="/admin" element={<AdminPage />}>
              <Route path="events" element={<div>Events</div>} />
              <Route path="locations" element={<div>Locations</div>} />
              <Route path="tickets" element={<div>Tickets</div>} />
              <Route path="users" element={<div>Users</div>} />
            </Route>
          </Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
