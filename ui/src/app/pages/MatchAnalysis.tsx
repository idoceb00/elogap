import { useEffect, useMemo, useState } from "react";
import { Link } from "react-router";
import { TrendingUp, AlertTriangle } from "lucide-react";
import { Card, CardContent } from "../components/ui/card";
import { Badge } from "../components/ui/badge";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../components/ui/select";
import { apiGet } from "../api/client";
import type { ActivityDTO } from "../api/types";

type SortBy = "playedAt" | "kda" | "cs" | "damage";
type FilterBy = "all" | "win" | "loss";

type LoadState =
  | { status: "idle" | "loading" }
  | { status: "error"; error: string }
  | { status: "success"; data: ActivityDTO[] };

function kdaValue(a: ActivityDTO): number {
  if (a.deaths > 0) return (a.kills + a.assists) / a.deaths;
  return a.kills + a.assists;
}

function formatDateShort(iso: string): string {
  const d = new Date(iso);
  if (Number.isNaN(d.getTime())) return iso;
  return d.toLocaleString("en-US", {
    month: "short",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit",
  });
}

export function MatchAnalysis() {
  const [sortBy, setSortBy] = useState<SortBy>("playedAt");
  const [filterBy, setFilterBy] = useState<FilterBy>("all");

  const [state, setState] = useState<LoadState>({ status: "idle" });

  useEffect(() => {
    let cancelled = false;
    setState({ status: "loading" });

    apiGet<ActivityDTO[]>("/v1/activities")
      .then((data) => {
        if (cancelled) return;
        setState({ status: "success", data });
      })
      .catch((err: unknown) => {
        if (cancelled) return;
        const msg = err instanceof Error ? err.message : "Failed to load matches";
        setState({ status: "error", error: msg });
      });

    return () => {
      cancelled = true;
    };
  }, []);

  const filteredSorted = useMemo(() => {
    if (state.status !== "success") return [];

    const filtered = state.data.filter((a) => {
      if (filterBy === "all") return true;
      return a.result === filterBy;
    });

    const sorted = [...filtered].sort((a, b) => {
      if (sortBy === "playedAt") {
        return new Date(b.playedAt).getTime() - new Date(a.playedAt).getTime();
      }
      if (sortBy === "kda") return kdaValue(b) - kdaValue(a);
      if (sortBy === "cs") return b.cs - a.cs;
      if (sortBy === "damage") return b.damage - a.damage;
      return 0;
    });

    return sorted;
  }, [state, filterBy, sortBy]);

  const totals = useMemo(() => {
    if (state.status !== "success") return { total: 0, wins: 0, losses: 0 };
    const wins = state.data.filter((m) => m.result === "win").length;
    const losses = state.data.filter((m) => m.result === "loss").length;
    return { total: state.data.length, wins, losses };
  }, [state]);

  return (
    <div className="space-y-6">
      {/* Page Header */}
      <div className="mb-8">
        <h1 className="text-3xl font-semibold mb-2">Match Analysis</h1>
        <p className="text-muted-foreground">
          Match history powered by your Go API (more analytics coming soon)
        </p>
      </div>

      {/* Summary Stats */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <Card>
          <CardContent className="pt-6">
            <p className="text-sm text-muted-foreground mb-1">Total Matches</p>
            <p className="text-3xl font-semibold">{totals.total}</p>
          </CardContent>
        </Card>
        <Card className="border-green-500/30 bg-green-500/5">
          <CardContent className="pt-6">
            <p className="text-sm text-muted-foreground mb-1">Wins</p>
            <p className="text-3xl font-semibold text-green-500">{totals.wins}</p>
          </CardContent>
        </Card>
        <Card className="border-red-500/30 bg-red-500/5">
          <CardContent className="pt-6">
            <p className="text-sm text-muted-foreground mb-1">Losses</p>
            <p className="text-3xl font-semibold text-red-500">{totals.losses}</p>
          </CardContent>
        </Card>
      </div>

      {/* Filters */}
      <div className="flex flex-col sm:flex-row sm:items-center gap-4">
        <Select value={filterBy} onValueChange={(v) => setFilterBy(v as FilterBy)}>
          <SelectTrigger className="w-56 bg-secondary">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="all">All Matches</SelectItem>
            <SelectItem value="win">Wins</SelectItem>
            <SelectItem value="loss">Losses</SelectItem>
          </SelectContent>
        </Select>

        <Select value={sortBy} onValueChange={(v) => setSortBy(v as SortBy)}>
          <SelectTrigger className="w-56 bg-secondary">
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="playedAt">Sort by Date</SelectItem>
            <SelectItem value="kda">Sort by KDA</SelectItem>
            <SelectItem value="cs">Sort by CS</SelectItem>
            <SelectItem value="damage">Sort by Damage</SelectItem>
          </SelectContent>
        </Select>
      </div>

      {/* Loading / Error */}
      {state.status === "loading" || state.status === "idle" ? (
        <Card>
          <CardContent className="p-6">
            <p className="text-sm text-muted-foreground">Loading matches…</p>
          </CardContent>
        </Card>
      ) : null}

      {state.status === "error" ? (
        <Card>
          <CardContent className="p-6">
            <p className="font-medium">Couldn’t load matches</p>
            <p className="text-sm text-muted-foreground mt-1">{state.error}</p>
          </CardContent>
        </Card>
      ) : null}

      {/* Match List */}
      {state.status === "success" ? (
        <div className="space-y-3">
          {filteredSorted.map((match) => (
            <MatchCard key={match.id} match={match} />
          ))}
        </div>
      ) : null}
    </div>
  );
}

function MatchCard({ match }: { match: ActivityDTO }) {
  const kda = kdaValue(match);

  return (
    <Link to={`/match-analysis/${match.id}`}>
      <Card
        className={`transition-all hover:scale-[1.01] cursor-pointer ${
          match.result === "win"
            ? "border-l-4 border-l-green-500"
            : "border-l-4 border-l-red-500"
        }`}
      >
        <CardContent className="p-5">
          <div className="flex flex-col md:flex-row md:items-center gap-4">
            {/* Result */}
            <div className="flex items-center gap-3 w-full md:w-48">
              <Badge
                variant={match.result === "win" ? "default" : "destructive"}
                className={`text-xs font-semibold ${
                  match.result === "win"
                    ? "bg-green-500/10 text-green-500"
                    : "bg-red-500/10 text-red-500"
                }`}
              >
                {match.result === "win" ? "WIN" : "LOSS"}
              </Badge>

              {match.result === "win" ? (
                <TrendingUp className="w-4 h-4 text-green-500" />
              ) : (
                <AlertTriangle className="w-4 h-4 text-red-500" />
              )}

              <div className="min-w-0">
                <p className="text-sm font-semibold truncate">{match.champion}</p>
                <p className="text-xs text-muted-foreground truncate">
                  {match.role} • {match.queueType}
                </p>
              </div>
            </div>

            {/* Core metrics */}
            <div className="grid grid-cols-3 gap-6 flex-1">
              <Metric label="KDA" value={kda.toFixed(2)} />
              <Metric label="CS" value={match.cs.toString()} />
              <Metric label="Damage" value={`${(match.damage / 1000).toFixed(1)}k`} />
            </div>

            {/* Date */}
            <div className="text-right">
              <p className="text-xs text-muted-foreground">{formatDateShort(match.playedAt)}</p>
            </div>
          </div>
        </CardContent>
      </Card>
    </Link>
  );
}

function Metric({ label, value }: { label: string; value: string }) {
  return (
    <div>
      <p className="text-xs text-muted-foreground mb-1">{label}</p>
      <p className="text-lg font-semibold">{value}</p>
    </div>
  );
}