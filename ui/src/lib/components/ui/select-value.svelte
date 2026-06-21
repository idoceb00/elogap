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

<span class={cn('', className)}>
  {@render children?.()}
</span>
