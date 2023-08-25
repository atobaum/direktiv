import NamespaceLayout from "~/pages/namespace/Layout";
import OnboardingPage from "~/pages/OnboardingPage";
import { createBrowserRouter } from "react-router-dom";
import env from "~/config/env";
import { pages } from "./pages";

export const router = createBrowserRouter(
  [
    {
      path: "/",
      element: <OnboardingPage />,
    },
    {
      path: "/:namespace",
      element: <NamespaceLayout />,
      children: Object.values(pages).map((page) => page.route),
      errorElement: (
        <div className="flex h-screen">
          <h1 className="m-auto text-center text-2xl font-bold">
            😿
            <br />
            oh no, an error occurred
          </h1>
        </div>
      ),
    },
  ],
  {
    basename: env.VITE_BASE ?? undefined,
  }
);
