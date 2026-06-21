<script lang="ts">
  import { onMount } from 'svelte';
  import { TrendingUp, AlertTriangle } from '@lucide/svelte';
  import Card from '$lib/components/ui/card.svelte';
  import CardContent from '$lib/components/ui/card-content.svelte';
  import Badge from '$lib/components/ui/badge.svelte';
  import Select from '$lib/components/ui/select.svelte';
  import SelectTrigger from '$lib/components/ui/select-trigger.svelte';
  import SelectValue from '$lib/components/ui/select-value.svelte';
  import SelectContent from '$lib/components/ui/select-content.svelte';
  import SelectItem from '$lib/components/ui/select-item.svelte';
  import { apiGet } from '$lib/api/client.js';
  import type { ActivityDTO } from '$lib/api/types.js';

  type SortBy = 'playedAt' | 'kda' | 'cs' | 'damage';
  type FilterBy = 'all' | 'win' | 'loss';
  type LoadState =
    | { status: 'idle' | 'loading' }
    | { status: 'error'; error: string }
    | { status: 'success'; data: ActivityDTO[] };

  let sortBy = $state<SortBy>('playedAt');
  let filterBy = $state<FilterBy>('all');
  let state = $state<LoadState>({ status: 'idle' });

  function kdaValue(a: ActivityDTO): number {
    if (a.deaths > 0) return (a.kills + a.assists) / a.deaths;
    return a.kills + a.assists;
  }

  function formatDateShort(iso: string): string {
    const d = new Date(iso);
    if (Number.isNaN(d.getTime())) return iso;
    return d.toLocaleString('en-US', {
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  }

  onMount(() => {
    let cancelled = false;
    state = { status: 'loading' };

    apiGet<ActivityDTO[]>('/v1/activities')
      .then((data) => {
        if (cancelled) return;
        state = { status: 'success', data };
      })
      .catch((err: unknown) => {
        if (cancelled) return;
        const msg = err instanceof Error ? err.message : 'Failed to load matches';
        state = { status: 'error', error: msg };
      });

    return () => { cancelled = true; };
  });

  let filteredSorted = $derived.by(() => {
    if (state.status !== 'success') return [];
    const filtered = state.data.filter((a) => filterBy === 'all' || a.result === filterBy);
    return [...filtered].sort((a, b) => {
      if (sortBy === 'playedAt') return new Date(b.playedAt).getTime() - new Date(a.playedAt).getTime();
      if (sortBy === 'kda') return kdaValue(b) - kdaValue(a);
      if (sortBy === 'cs') return b.cs - a.cs;
      if (sortBy === 'damage') return b.damage - a.damage;
      return 0;
    });
  });

  let totals = $derived.by(() => {
    if (state.status !== 'success') return { total: 0, wins: 0, losses: 0 };
    const wins = state.data.filter((m) => m.result === 'win').length;
    const losses = state.data.filter((m) => m.result === 'loss').length;
    return { total: state.data.length, wins, losses };
  });
</script>

<div class="space-y-6">
  <div class="mb-8">
    <h1 class="text-3xl font-semibold mb-2">Match Analysis</h1>
    <p class="text-muted-foreground">Match history powered by your Go API (more analytics coming soon)</p>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
    <Card>
      <CardContent class="pt-6">
        <p class="text-sm text-muted-foreground mb-1">Total Matches</p>
        <p class="text-3xl font-semibold">{totals.total}</p>
      </CardContent>
    </Card>
    <Card class="border-green-500/30 bg-green-500/5">
      <CardContent class="pt-6">
        <p class="text-sm text-muted-foreground mb-1">Wins</p>
        <p class="text-3xl font-semibold text-green-500">{totals.wins}</p>
      </CardContent>
    </Card>
    <Card class="border-red-500/30 bg-red-500/5">
      <CardContent class="pt-6">
        <p class="text-sm text-muted-foreground mb-1">Losses</p>
        <p class="text-3xl font-semibold text-red-500">{totals.losses}</p>
      </CardContent>
    </Card>
  </div>

  <div class="flex flex-col sm:flex-row sm:items-center gap-4">
    <Select bind:value={filterBy}>
      <SelectTrigger class="w-56 bg-secondary">
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="all">All Matches</SelectItem>
        <SelectItem value="win">Wins</SelectItem>
        <SelectItem value="loss">Losses</SelectItem>
      </SelectContent>
    </Select>

    <Select bind:value={sortBy}>
      <SelectTrigger class="w-56 bg-secondary">
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

  {#if state.status === 'loading' || state.status === 'idle'}
    <Card>
      <CardContent class="p-6">
        <p class="text-sm text-muted-foreground">Loading matches...</p>
      </CardContent>
    </Card>
  {/if}

  {#if state.status === 'error'}
    <Card>
      <CardContent class="p-6">
        <p class="font-medium">Couldn't load matches</p>
        <p class="text-sm text-muted-foreground mt-1">{state.error}</p>
      </CardContent>
    </Card>
  {/if}

  {#if state.status === 'success'}
    <div class="space-y-3">
      {#each filteredSorted as match (match.id)}
        {@const kda = kdaValue(match)}
        <a href="/match-analysis/{match.id}">
          <Card class="transition-all hover:scale-[1.01] cursor-pointer {match.result === 'win' ? 'border-l-4 border-l-green-500' : 'border-l-4 border-l-red-500'}">
            <CardContent class="p-5">
              <div class="flex flex-col md:flex-row md:items-center gap-4">
                <div class="flex items-center gap-3 w-full md:w-48">
                  <Badge
                    variant={match.result === 'win' ? 'default' : 'destructive'}
                    class="text-xs font-semibold {match.result === 'win' ? 'bg-green-500/10 text-green-500' : 'bg-red-500/10 text-red-500'}"
                  >
                    {match.result === 'win' ? 'WIN' : 'LOSS'}
                  </Badge>
                  {#if match.result === 'win'}
                    <TrendingUp class="w-4 h-4 text-green-500" />
                  {:else}
                    <AlertTriangle class="w-4 h-4 text-red-500" />
                  {/if}
                  <div class="min-w-0">
                    <p class="text-sm font-semibold truncate">{match.champion}</p>
                    <p class="text-xs text-muted-foreground truncate">{match.role} · {match.queueType}</p>
                  </div>
                </div>

                <div class="grid grid-cols-3 gap-6 flex-1">
                  <div>
                    <p class="text-xs text-muted-foreground mb-1">KDA</p>
                    <p class="text-lg font-semibold">{kda.toFixed(2)}</p>
                  </div>
                  <div>
                    <p class="text-xs text-muted-foreground mb-1">CS</p>
                    <p class="text-lg font-semibold">{match.cs}</p>
                  </div>
                  <div>
                    <p class="text-xs text-muted-foreground mb-1">Damage</p>
                    <p class="text-lg font-semibold">{(match.damage / 1000).toFixed(1)}k</p>
                  </div>
                </div>

                <div class="text-right">
                  <p class="text-xs text-muted-foreground">{formatDateShort(match.playedAt)}</p>
                </div>
              </div>
            </CardContent>
          </Card>
        </a>
      {/each}
    </div>
  {/if}
</div>
