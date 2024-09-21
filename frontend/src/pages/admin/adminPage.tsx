import PageContainer from "@/components/ui/pageContainer";
import SideBar from "./sidebar";
import { Outlet } from "react-router-dom";

export default function AdminPage() {
  return (
    <PageContainer>
      <div className="flex flex-1 gap-6 py-3">
        <SideBar />
        <div className="py-4">
          <Outlet />
        </div>
      </div>
    </PageContainer>
  );
}
