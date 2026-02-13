import { useEffect, useMemo, useState } from "react";
import { useParams, Link } from "react-router";
import {
  ArrowLeft,
  Swords,
  Target,
  Zap,
  Eye,
  Clock,
} from "lucide-react";
import { Card, CardContent } from "../components/ui/card";
import { Badge } from "../components/ui/badge";
import { Button } from "../components/ui/button";
import { apiGet } from "../api/client";
import type { ActivityDTO } from "../api/types";

type LoadState =
  | { status: "idle" | "loading" }
  | { status: "error"; error: string }
  | { status: "success"; data: ActivityDTO };

function formatDurationMMSS(totalSeconds: number): string {
  const mins = Math.floor(totalSeconds / 60);
  const secs = totalSeconds % 60;
  return `${mins}:${secs.toString().padStart(2, "0")}`;
}

function formatPlayedAt(iso: string): string {
  const d = new Date(iso);
  if (Number.isNaN(d.getTime())) return iso;
  return d.toLocaleString("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}

export function ActivityDetail() {
  const { id } = useParams<{ id: string }>();

  const [state, setState] = useState<LoadState>({ status: "idle" });

  useEffect(() => {
    if (!id) {
      setState({ status: "error", error: "Missing match id" });
      return;
    }

    let cancelled = false;
    setState({ status: "loading" });

    apiGet<ActivityDTO>(`/v1/activities/${id}`)
      .then((data) => {
        if (cancelled) return;
        setState({ status: "success", data });
      })
      .catch((err: unknown) => {
        if (cancelled) return;
        const msg = err instanceof Error ? err.message : "Failed to load match";
        setState({ status: "error", error: msg });
      });

    return () => {
      cancelled = true;
    };
  }, [id]);

  const match = state.status === "success" ? state.data : null;

  const kda = useMemo(() => {
    if (!match) return null;
    if (match.deaths > 0) return (match.kills + match.assists) / match.deaths;
    return match.kills + match.assists;
  }, [match]);

  const csPerMin = useMemo(() => {
    if (!match) return null;
    const minutes = match.duration / 60;
    if (minutes <= 0) return null;
    return match.cs / minutes;
  }, [match]);

  if (state.status === "loading" || state.status === "idle") {
    return (
      <div className="p-6 space-y-4 max-w-[1600px] mx-auto">
        <Link to="/match-analysis">
          <Button variant="ghost" size="sm">
            <ArrowLeft className="w-4 h-4 mr-2" />
            Back to Matches
          </Button>
        </Link>
        <Card>
          <CardContent className="p-6">
            <p className="text-sm text-muted-foreground">Loading match…</p>
          </CardContent>
        </Card>
      </div>
    );
  }

  if (state.status === "error") {
    return (
      <div className="p-6 space-y-4 max-w-[1600px] mx-auto">
        <Link to="/match-analysis">
          <Button variant="ghost" size="sm">
            <ArrowLeft className="w-4 h-4 mr-2" />
            Back to Matches
          </Button>
        </Link>
        <Card>
          <CardContent className="p-6">
            <p className="font-medium">Couldn’t load match</p>
            <p className="text-sm text-muted-foreground mt-1">{state.error}</p>
          </CardContent>
        </Card>
      </div>
    );
  }

  // success
  const durationLabel = formatDurationMMSS(match!.duration);
  const playedAtLabel = formatPlayedAt(match!.playedAt);

  return (
    <div className="p-6 space-y-6 max-w-[1600px] mx-auto">
      {/* Back Button */}
      <Link to="/match-analysis">
        <Button variant="ghost" size="sm" className="mb-2">
          <ArrowLeft className="w-4 h-4 mr-2" />
          Back to Matches
        </Button>
      </Link>

      {/* Header */}
      <Card
        className={`${
          match!.result === "win"
            ? "border-l-4 border-l-green-500"
            : "border-l-4 border-l-red-500"
        }`}
      >
        <CardContent className="p-6">
          <div className="flex flex-col md:flex-row md:items-center gap-4">
            {/* Champion Avatar */}
            <div className="flex items-center gap-4">
              <div className="w-20 h-20 rounded-xl bg-primary/10 flex items-center justify-center">
                <span className="text-2xl font-bold text-primary">
                  {match!.champion.slice(0, 2).toUpperCase()}
                </span>
              </div>
              <div>
                <div className="flex items-center gap-2 mb-1">
                  <h1 className="text-2xl font-semibold">{match!.champion}</h1>
                  <Badge
                    variant={match!.result === "win" ? "default" : "destructive"}
                    className={`${
                      match!.result === "win"
                        ? "bg-green-500/10 text-green-500"
                        : "bg-red-500/10 text-red-500"
                    }`}
                  >
                    {match!.result === "win" ? "VICTORY" : "DEFEAT"}
                  </Badge>
                </div>
                <div className="flex items-center gap-3 text-sm text-muted-foreground">
                  <span>{match!.role}</span>
                  <span>•</span>
                  <span>{match!.queueType}</span>
                  <span>•</span>
                  <span>{playedAtLabel}</span>
                </div>
              </div>
            </div>

            {/* Quick Stats */}
            <div className="flex-1 grid grid-cols-2 md:grid-cols-4 gap-4 md:ml-auto">
              <div className="text-center">
                <p className="text-2xl font-semibold">
                  {match!.kills}/{match!.deaths}/{match!.assists}
                </p>
                <p className="text-xs text-muted-foreground">
                  KDA: {kda ? kda.toFixed(2) : "—"}
                </p>
              </div>
              <div className="text-center">
                <p className="text-2xl font-semibold">{match!.cs}</p>
                <p className="text-xs text-muted-foreground">
                  {csPerMin ? `${csPerMin.toFixed(1)} CS/min` : "—"}
                </p>
              </div>
              <div className="text-center">
                <p className="text-2xl font-semibold">
                  {(match!.damage / 1000).toFixed(1)}k
                </p>
                <p className="text-xs text-muted-foreground">Damage</p>
              </div>
              <div className="text-center">
                <p className="text-2xl font-semibold">{durationLabel}</p>
                <p className="text-xs text-muted-foreground">Duration</p>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Performance Grid (only what backend provides) */}
      <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
        <StatCard
          icon={Swords}
          label="K/D/A"
          value={`${match!.kills}/${match!.deaths}/${match!.assists}`}
          subtext={kda ? `${kda.toFixed(2)} KDA` : "—"}
        />
        <StatCard
          icon={Target}
          label="CS/min"
          value={csPerMin ? csPerMin.toFixed(1) : "—"}
          subtext={`${match!.cs} total CS`}
        />
        <StatCard
          icon={Zap}
          label="Damage"
          value={`${(match!.damage / 1000).toFixed(1)}k`}
          subtext="Total dealt"
        />
        <StatCard
          icon={Eye}
          label="Vision"
          value={match!.vision}
          subtext="Vision score"
        />
        <StatCard
          icon={Clock}
          label="Game Time"
          value={durationLabel}
          subtext={`${match!.duration}s`}
        />
      </div>

      {/* Coming Soon blocks for analytics-heavy parts */}
      <Card className="border-primary/20 bg-primary/5">
        <CardContent className="p-6">
          <p className="font-medium">More match analytics coming soon</p>
          <p className="text-sm text-muted-foreground mt-1">
            Damage breakdown, gold timeline, comparisons and insights will appear here
            once the backend exposes those stats.
          </p>
        </CardContent>
      </Card>
    </div>
  );
}

interface StatCardProps {
  icon: any;
  label: string;
  value: string | number;
  subtext: string;
}

function StatCard({ icon: Icon, label, value, subtext }: StatCardProps) {
  return (
    <Card>
      <CardContent className="pt-6">
        <div className="flex items-center gap-2 mb-3">
          <Icon className="w-4 h-4 text-muted-foreground" />
          <span className="text-sm text-muted-foreground">{label}</span>
        </div>
        <p className="text-2xl font-semibold mb-1">{value}</p>
        <p className="text-xs text-muted-foreground">{subtext}</p>
      </CardContent>
    </Card>
  );
}