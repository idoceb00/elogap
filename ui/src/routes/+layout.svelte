<script lang="ts">
  import '../app.css';
  import { page } from '$app/state';
  import { LayoutDashboard, Activity, TrendingUp, Settings, Search, Menu, X, Moon, Sun, User, Trophy } from '@lucide/svelte';
  import Button from '$lib/components/ui/button.svelte';
  import Input from '$lib/components/ui/input.svelte';

  let { children } = $props();

  let sidebarOpen = $state(true);
  let mobileSidebarOpen = $state(false);
  let theme = $state<'light' | 'dark'>('dark');

  const navigation = [
    { name: 'Overview', href: '/', icon: LayoutDashboard },
    { name: 'Metrics', href: '/metrics', icon: TrendingUp },
    { name: 'Match Analysis', href: '/match-analysis', icon: Activity },
    { name: 'Settings', href: '/settings', icon: Settings }
  ];

  function toggleTheme() {
    theme = theme === 'light' ? 'dark' : 'light';
    document.documentElement.classList.toggle('dark');
  }

  function isActive(href: string) {
    if (href === '/') return page.url.pathname === '/';
    return page.url.pathname.startsWith(href);
  }

  function toggleSidebar() {
    if (typeof window !== 'undefined' && window.innerWidth < 1024) {
      mobileSidebarOpen = !mobileSidebarOpen;
    } else {
      sidebarOpen = !sidebarOpen;
    }
  }
</script>

<div class="h-screen flex overflow-hidden bg-background dark">
  {#if mobileSidebarOpen}
    <button
      class="fixed inset-0 bg-black/50 z-40 lg:hidden"
      onclick={() => mobileSidebarOpen = false}
      aria-label="Close sidebar"
    ></button>
  {/if}

  <aside
    class="{sidebarOpen ? 'lg:w-64' : 'lg:w-0'} {mobileSidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'} fixed lg:relative inset-y-0 left-0 z-50 w-64 bg-sidebar border-r border-sidebar-border transition-all duration-300 flex-shrink-0 overflow-hidden"
  >
    <div class="h-full flex flex-col">
      <div class="h-16 flex items-center gap-3 px-6 border-b border-sidebar-border">
        <div class="w-8 h-8 rounded-lg bg-primary flex items-center justify-center">
          <Trophy class="w-5 h-5 text-primary-foreground" />
        </div>
        <span class="font-semibold text-lg text-sidebar-foreground">EloGap</span>
      </div>

      <nav class="flex-1 px-3 py-4 space-y-1">
        {#each navigation as item}
          {@const active = isActive(item.href)}
          <a
            href={item.href}
            class="flex items-center gap-3 px-3 py-2.5 rounded-lg transition-colors {active ? 'bg-sidebar-accent text-sidebar-accent-foreground' : 'text-sidebar-foreground hover:bg-sidebar-accent/50'}"
          >
            <item.icon class="w-5 h-5" />
            <span class="font-medium">{item.name}</span>
          </a>
        {/each}
      </nav>

      <div class="p-4 border-t border-sidebar-border">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-full bg-primary flex items-center justify-center">
            <User class="w-5 h-5 text-primary-foreground" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-sm font-medium text-sidebar-foreground truncate">Summoner</p>
            <p class="text-xs text-muted-foreground">Diamond IV</p>
          </div>
        </div>
      </div>
    </div>
  </aside>

  <div class="flex-1 flex flex-col min-w-0">
    <header class="h-16 bg-card border-b border-border flex items-center gap-4 px-6">
      <Button variant="ghost" size="sm" onclick={toggleSidebar} class="flex">
        {#if (sidebarOpen && typeof window !== 'undefined' && window.innerWidth >= 1024) || mobileSidebarOpen}
          <X class="w-5 h-5" />
        {:else}
          <Menu class="w-5 h-5" />
        {/if}
      </Button>

      <div class="hidden sm:flex flex-1 max-w-md">
        <div class="relative w-full">
          <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-muted-foreground" />
          <Input placeholder="Search champions, matches..." class="pl-9 bg-secondary border-border" />
        </div>
      </div>

      <Button variant="ghost" size="sm" onclick={toggleTheme} class="hidden sm:flex w-9 h-9 p-0">
        {#if theme === 'light'}
          <Moon class="w-4 h-4" />
        {:else}
          <Sun class="w-4 h-4" />
        {/if}
      </Button>

      <Button variant="ghost" size="sm" class="w-9 h-9 p-0 rounded-full">
        <div class="w-8 h-8 rounded-full bg-primary flex items-center justify-center">
          <User class="w-4 h-4 text-primary-foreground" />
        </div>
      </Button>
    </header>

    <main class="flex-1 overflow-auto bg-background">
      <div class="p-6 max-w-[1800px] mx-auto">
        {@render children()}
      </div>
    </main>
  </div>
</div>
