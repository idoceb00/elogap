<script lang="ts">
  import type { Snippet } from 'svelte';
  import { cn } from '$lib/utils.js';

  let {
    children,
    class: className
  }: {
    children?: Snippet;
    class?: string;
  } = $props();
</script>

<div class={cn('px-2 py-1.5 text-sm font-semibold', className)}>
  {@render children?.()}
</div>
