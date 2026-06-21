<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/state';
  import { ArrowLeft, Swords, Target, Zap, Eye, Clock } from '@lucide/svelte';
  import Card from '$lib/components/ui/card.svelte';
  import CardContent from '$lib/components/ui/card-content.svelte';
  import Badge from '$lib/components/ui/badge.svelte';
  import Button from '$lib/components/ui/button.svelte';
  import { apiGet } from '$lib/api/client.js';
  import type { ActivityDTO } from '$lib/api/types.js';

  type LoadState =
    | { status: 'idle' | 'loading' }
    | { status: 'error'; error: string }
    | { status: 'success'; data: ActivityDTO };

  let state = $state<LoadState>({ status: 'idle' });

  function formatDurationMMSS(totalSeconds: number): string {
    const mins = Math.floor(totalSeconds / 60);
    const secs = totalSeconds % 60;
    return `${mins}:${secs.toString().padStart(2, '0')}`;
  }

  function formatPlayedAt(iso: string): string {
    const d = new Date(iso);
    if (Number.isNaN(d.getTime())) return iso;
    return d.toLocaleString('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    });
  }

  onMount(() => {
    const id = page.params.id;
    if (!id) {
      state = { status: 'error', error: 'Missing match id' };
      return;
    }

    let cancelled = false;
    state = { status: 'loading' };

    apiGet<ActivityDTO>(`/v1/activities/${id}`)
      .then((data) => {
        if (cancelled) return;
        state = { status: 'success', data };
      })
      .catch((err: unknown) => {
        if (cancelled) return;
        const msg = err instanceof Error ? err.message : 'Failed to load match';
        state = { status: 'error', error: msg };
      });

    return () => { cancelled = true; };
  });

  let match = $derived(state.status === 'success' ? state.data : null);
  let kda = $derived.by(() => {
    if (!match) return null;
    if (match.deaths > 0) return (match.kills + match.assists) / match.deaths;
    return match.kills + match.assists;
  });
  let csPerMin = $derived.by(() => {
    if (!match) return null;
    const minutes = match.duration / 60;
    if (minutes <= 0) return null;
    return match.cs / minutes;
  });
</script>

<div class="p-6 space-y-6 max-w-[1600px] mx-auto">
  <a href="/match-analysis">
    <Button variant="ghost" size="sm" class="mb-2">
      <ArrowLeft class="w-4 h-4 mr-2" />
      Back to Matches
    </Button>
  </a>

  {#if state.status === 'loading' || state.status === 'idle'}
    <Card>
      <CardContent class="p-6">
        <p class="text-sm text-muted-foreground">Loading match...</p>
      </CardContent>
    </Card>
  {:else if state.status === 'error'}
    <Card>
      <CardContent class="p-6">
        <p class="font-medium">Couldn't load match</p>
        <p class="text-sm text-muted-foreground mt-1">{state.error}</p>
      </CardContent>
    </Card>
  {:else if match}
    {@const durationLabel = formatDurationMMSS(match.duration)}
    {@const playedAtLabel = formatPlayedAt(match.playedAt)}

    <Card class="{match.result === 'win' ? 'border-l-4 border-l-green-500' : 'border-l-4 border-l-red-500'}">
      <CardContent class="p-6">
        <div class="flex flex-col md:flex-row md:items-center gap-4">
          <div class="flex items-center gap-4">
            <div class="w-20 h-20 rounded-xl bg-primary/10 flex items-center justify-center">
              <span class="text-2xl font-bold text-primary">{match.champion.slice(0, 2).toUpperCase()}</span>
            </div>
            <div>
              <div class="flex items-center gap-2 mb-1">
                <h1 class="text-2xl font-semibold">{match.champion}</h1>
                <Badge
                  variant={match.result === 'win' ? 'default' : 'destructive'}
                  class="{match.result === 'win' ? 'bg-green-500/10 text-green-500' : 'bg-red-500/10 text-red-500'}"
                >
                  {match.result === 'win' ? 'VICTORY' : 'DEFEAT'}
                </Badge>
              </div>
              <div class="flex items-center gap-3 text-sm text-muted-foreground">
                <span>{match.role}</span>
                <span>·</span>
                <span>{match.queueType}</span>
                <span>·</span>
                <span>{playedAtLabel}</span>
              </div>
            </div>
          </div>

          <div class="flex-1 grid grid-cols-2 md:grid-cols-4 gap-4 md:ml-auto">
            <div class="text-center">
              <p class="text-2xl font-semibold">{match.kills}/{match.deaths}/{match.assists}</p>
              <p class="text-xs text-muted-foreground">KDA: {kda ? kda.toFixed(2) : '—'}</p>
            </div>
            <div class="text-center">
              <p class="text-2xl font-semibold">{match.cs}</p>
              <p class="text-xs text-muted-foreground">{csPerMin ? `${csPerMin.toFixed(1)} CS/min` : '—'}</p>
            </div>
            <div class="text-center">
              <p class="text-2xl font-semibold">{(match.damage / 1000).toFixed(1)}k</p>
              <p class="text-xs text-muted-foreground">Damage</p>
            </div>
            <div class="text-center">
              <p class="text-2xl font-semibold">{durationLabel}</p>
              <p class="text-xs text-muted-foreground">Duration</p>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>

    <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
      <Card>
        <CardContent class="pt-6">
          <div class="flex items-center gap-2 mb-3">
            <Swords class="w-4 h-4 text-muted-foreground" />
            <span class="text-sm text-muted-foreground">K/D/A</span>
          </div>
          <p class="text-2xl font-semibold mb-1">{match.kills}/{match.deaths}/{match.assists}</p>
          <p class="text-xs text-muted-foreground">{kda ? `${kda.toFixed(2)} KDA` : '—'}</p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="pt-6">
          <div class="flex items-center gap-2 mb-3">
            <Target class="w-4 h-4 text-muted-foreground" />
            <span class="text-sm text-muted-foreground">CS/min</span>
          </div>
          <p class="text-2xl font-semibold mb-1">{csPerMin ? csPerMin.toFixed(1) : '—'}</p>
          <p class="text-xs text-muted-foreground">{match.cs} total CS</p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="pt-6">
          <div class="flex items-center gap-2 mb-3">
            <Zap class="w-4 h-4 text-muted-foreground" />
            <span class="text-sm text-muted-foreground">Damage</span>
          </div>
          <p class="text-2xl font-semibold mb-1">{(match.damage / 1000).toFixed(1)}k</p>
          <p class="text-xs text-muted-foreground">Total dealt</p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="pt-6">
          <div class="flex items-center gap-2 mb-3">
            <Eye class="w-4 h-4 text-muted-foreground" />
            <span class="text-sm text-muted-foreground">Vision</span>
          </div>
          <p class="text-2xl font-semibold mb-1">{match.vision}</p>
          <p class="text-xs text-muted-foreground">Vision score</p>
        </CardContent>
      </Card>
      <Card>
        <CardContent class="pt-6">
          <div class="flex items-center gap-2 mb-3">
            <Clock class="w-4 h-4 text-muted-foreground" />
            <span class="text-sm text-muted-foreground">Game Time</span>
          </div>
          <p class="text-2xl font-semibold mb-1">{durationLabel}</p>
          <p class="text-xs text-muted-foreground">{match.duration}s</p>
        </CardContent>
      </Card>
    </div>

    <Card class="border-primary/20 bg-primary/5">
      <CardContent class="p-6">
        <p class="font-medium">More match analytics coming soon</p>
        <p class="text-sm text-muted-foreground mt-1">
          Damage breakdown, gold timeline, comparisons and insights will appear here
          once the backend exposes those stats.
        </p>
      </CardContent>
    </Card>
  {/if}
</div>
