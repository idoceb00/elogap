export type MatchResult = 'win' | 'loss';

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
  duration: number;
  damage: number;
  vision: number;
  playedAt: string;
}

export interface MetricsSummary {
  playerId: string;
  range: string;
  totalMatches: number;
  winRate: number;
  avgKda: number;
  avgCsPerMin: number;
  avgVision: number;
  avgDamage: number;
  avgDurationSeconds: number;
}

export interface TrendPoint {
  date: string;
  winRate: number;
  kda: number;
  csPerMin: number;
}

export interface MetricsTrends {
  playerId: string;
  range: string;
  points: TrendPoint[];
}
