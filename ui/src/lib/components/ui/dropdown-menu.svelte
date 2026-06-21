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

  let open = $state(false);
</script>

<div class={cn('relative inline-block text-left', className)}>
  {@render children?.()}
</div>
