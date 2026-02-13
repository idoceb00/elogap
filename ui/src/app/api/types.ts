export type MatchResult = "win" | "loss";

export interface ActivityDTO {
  id: string;
  playerId: string;
  result: MatchResult;
  queueType: string;
  champion: string;
  role: string;
  kills: number;
  deaths: number;
  assists: number;
  cs: number;
  duration: number; // seconds
  damage: number;
  vision: number;
  playedAt: string; // ISO string
}

export interface Match {
  id: string;
  result: "win" | "loss";
  champion: string;
  role: "top" | "jungle" | "mid" | "adc" | "support";
  queueType: "Ranked Solo" | "Ranked Flex" | "Normal" | "ARAM";
  kills: number;
  deaths: number;
  assists: number;
  cs: number;
  damage: number;
  damageTaken: number;
  visionScore: number;
  gold: number;
  duration: number; // in minutes
  date: Date;
  gameLength: number; // in seconds
}

export interface PlayerStats {
  totalMatches: number;
  winRate: number;
  kda: number;
  csPerMin: number;
  currentStreak: { type: "win" | "loss"; count: number };
  formScore: number;
}

export interface TrendData {
  date: string;
  winRate: number;
  kda: number;
}

export interface Insight {
  type: "best" | "champion" | "streak";
  title: string;
  value: string;
  description?: string;
}
