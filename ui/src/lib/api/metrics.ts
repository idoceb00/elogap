import { apiGet } from './client.js';
import type { MetricsSummary, MetricsTrends } from './types.js';

export function getMetricsSummary(playerId: string, range: string) {
  return apiGet<MetricsSummary>(
    `/v1/metrics/summary?playerId=${playerId}&range=${range}`
  );
}

export function getMetricsTrends(playerId: string, range: string) {
  return apiGet<MetricsTrends>(
    `/v1/metrics/trends?playerId=${playerId}&range=${range}`
  );
}
