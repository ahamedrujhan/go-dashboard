import RootLayout from "../layouts/RootLayout.tsx";
import Dashboard from "../pages/Dashboard.tsx";
import {createBrowserRouter, RouterProvider} from "react-router-dom";

function Routes() {

    const routes = [
        {
            path: "/",
            element: <RootLayout/>,
            children: [
                {
                    index : true,
                    element: <Dashboard/>
                },
                {
                    path: "*",
                    element: <Dashboard/>
                }
            ]
        }
    ]

    const router = createBrowserRouter([
        ...routes
    ])

    return (
        <RouterProvider router={router}/>
    )
}
export default Routes

