<script lang="ts">
  import { onMount } from 'svelte';
  import Card from '$lib/components/ui/card.svelte';
  import CardContent from '$lib/components/ui/card-content.svelte';
  import CardHeader from '$lib/components/ui/card-header.svelte';
  import CardTitle from '$lib/components/ui/card-title.svelte';
  import { getMetricsSummary, getMetricsTrends } from '$lib/api/metrics.js';
  import type { MetricsSummary, MetricsTrends } from '$lib/api/types.js';
  import { Chart, Svg, Axis, Line, Tooltip } from 'layerchart';

  let summary = $state<MetricsSummary | null>(null);
  let trends = $state<MetricsTrends | null>(null);
  let loading = $state(true);
  let error = $state<string | null>(null);

  onMount(async () => {
    try {
      loading = true;
      const [s, t] = await Promise.all([
        getMetricsSummary('player_1', '30d'),
        getMetricsTrends('player_1', '30d')
      ]);
      summary = s;
      trends = t;
    } catch (err) {
      error = err instanceof Error ? err.message : 'Unknown error';
    } finally {
      loading = false;
    }
  });

  let avgDurationMinutes = $derived(
    summary ? (summary.avgDurationSeconds / 60).toFixed(1) : '0'
  );
</script>

<div class="p-6 space-y-6">
  <h1 class="text-3xl font-semibold">Metric Deep Dive</h1>

  {#if loading}
    <div class="p-6">Loading metrics...</div>
  {:else if error}
    <div class="p-6 text-red-500">Error: {error}</div>
  {:else if summary && trends}
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
      <Card>
        <CardContent class="pt-6">
          <p class="text-sm text-muted-foreground">Total Matches</p>
          <p class="text-2xl font-semibold">{summary.totalMatches}</p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="pt-6">
          <p class="text-sm text-muted-foreground">Win Rate</p>
          <p class="text-2xl font-semibold">{summary.winRate}%</p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="pt-6">
          <p class="text-sm text-muted-foreground">Avg KDA</p>
          <p class="text-2xl font-semibold">{summary.avgKda}</p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="pt-6">
          <p class="text-sm text-muted-foreground">Avg CS/min</p>
          <p class="text-2xl font-semibold">{summary.avgCsPerMin}</p>
        </CardContent>
      </Card>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
      <Card>
        <CardContent class="pt-6">
          <p class="text-sm text-muted-foreground">Avg Vision</p>
          <p class="text-xl font-semibold">{summary.avgVision}</p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="pt-6">
          <p class="text-sm text-muted-foreground">Avg Damage</p>
          <p class="text-xl font-semibold">{summary.avgDamage}</p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="pt-6">
          <p class="text-sm text-muted-foreground">Avg Duration (min)</p>
          <p class="text-xl font-semibold">{avgDurationMinutes}</p>
        </CardContent>
      </Card>
    </div>

    <Card>
      <CardHeader>
        <CardTitle>Performance Trend</CardTitle>
      </CardHeader>
      <CardContent>
        <div class="h-[300px]">
          <Chart data={trends.points} x="date" y="winRate">
            {#snippet children({ context })}
              <Svg>
                <Axis placement="bottom" />
                <Axis placement="left" />
                <Line class="stroke-primary" stroke-width={2} />
                <Tooltip />
              </Svg>
            {/snippet}
          </Chart>
        </div>
      </CardContent>
    </Card>
  {/if}
</div>
