import { createBrowserRouter } from "react-router";
import { consoleRoutes } from "./console/routes";
import { formRoutes } from "./forms/routes";

export const router = createBrowserRouter([
	...formRoutes,
	...consoleRoutes,
    {
        path: "*",
        Component: () => {
            return (<h1>Not found</h1>)
        },
    },
]);
