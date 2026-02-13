import { createBrowserRouter } from "react-router";
import { RootLayout } from "./layouts/RootLayout";
import { PerformanceOverview } from "./pages/PerformanceOverview";
import { MetricDeepDive } from "./pages/MetricDeepDive";
import { MatchAnalysis } from "./pages/MatchAnalysis";
import { ActivityDetail } from "./pages/ActivityDetail";
import { Settings } from "./pages/Settings";

export const router = createBrowserRouter([
  {
    path: "/",
    Component: RootLayout,
    children: [
      { index: true, Component: PerformanceOverview },
      { path: "metrics", Component: MetricDeepDive },
      { path: "match-analysis", Component: MatchAnalysis },
      { path: "match-analysis/:id", Component: ActivityDetail },
      { path: "settings", Component: Settings },
    ],
  },
]);