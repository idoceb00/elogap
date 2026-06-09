import { useEffect, useState } from "react";
import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
} from "../components/ui/card";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  ResponsiveContainer,
} from "recharts";

import { getMetricsSummary, getMetricsTrends } from "../api/metrics";
import type { MetricsSummary, MetricsTrends } from "../api/types";

export function MetricDeepDive() {
  const [summary, setSummary] = useState<MetricsSummary | null>(null);
  const [trends, setTrends] = useState<MetricsTrends | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    async function load() {
      try {
        setLoading(true);

        const [s, t] = await Promise.all([
          getMetricsSummary("player_1", "30d"),
          getMetricsTrends("player_1", "30d"),
        ]);

        setSummary(s);
        setTrends(t);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Unknown error");
      } finally {
        setLoading(false);
      }
    }

    load();
  }, []);

  if (loading) return <div className="p-6">Loading metrics...</div>;
  if (error)
    return <div className="p-6 text-red-500">Error: {error}</div>;
  if (!summary || !trends) return null;

  const avgDurationMinutes = (
    summary.avgDurationSeconds / 60
  ).toFixed(1);

  return (
    <div className="p-6 space-y-6">
      <h1 className="text-3xl font-semibold">Metric Deep Dive</h1>

      {/* Summary Cards */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <Card>
          <CardContent className="pt-6">
            <p className="text-sm text-muted-foreground">
              Total Matches
            </p>
            <p className="text-2xl font-semibold">
              {summary.totalMatches}
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <p className="text-sm text-muted-foreground">
              Win Rate
            </p>
            <p className="text-2xl font-semibold">
              {summary.winRate}%
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <p className="text-sm text-muted-foreground">
              Avg KDA
            </p>
            <p className="text-2xl font-semibold">
              {summary.avgKda}
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <p className="text-sm text-muted-foreground">
              Avg CS/min
            </p>
            <p className="text-2xl font-semibold">
              {summary.avgCsPerMin}
            </p>
          </CardContent>
        </Card>
      </div>

      {/* Additional Stats */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <Card>
          <CardContent className="pt-6">
            <p className="text-sm text-muted-foreground">
              Avg Vision
            </p>
            <p className="text-xl font-semibold">
              {summary.avgVision}
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <p className="text-sm text-muted-foreground">
              Avg Damage
            </p>
            <p className="text-xl font-semibold">
              {summary.avgDamage}
            </p>
          </CardContent>
        </Card>

        <Card>
          <CardContent className="pt-6">
            <p className="text-sm text-muted-foreground">
              Avg Duration (min)
            </p>
            <p className="text-xl font-semibold">
              {avgDurationMinutes}
            </p>
          </CardContent>
        </Card>
      </div>

      {/* Trends Chart */}
      <Card>
        <CardHeader>
          <CardTitle>Performance Trend</CardTitle>
        </CardHeader>
        <CardContent>
          <ResponsiveContainer width="100%" height={300}>
            <LineChart data={trends.points}>
              <CartesianGrid strokeDasharray="3 3" />
              <XAxis dataKey="date" />
              <YAxis />
              <Tooltip />
              <Line
                type="monotone"
                dataKey="winRate"
                stroke="#8884d8"
                strokeWidth={2}
              />
              <Line
                type="monotone"
                dataKey="kda"
                stroke="#82ca9d"
                strokeWidth={2}
              />
            </LineChart>
          </ResponsiveContainer>
        </CardContent>
      </Card>
    </div>
  );
}