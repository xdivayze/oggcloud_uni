import { Outlet } from "react-router-dom";

export default function Layout() {
    return(
        <div className="bg-green-ogg-1 min-h-screen min-w-screen">
            <Outlet />
        </div>
    )
}