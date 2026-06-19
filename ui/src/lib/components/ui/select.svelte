<script lang="ts">
  import type { Snippet } from 'svelte';
  import { cn } from '$lib/utils.js';

  let {
    value = $bindable(''),
    children,
    class: className
  }: {
    value?: string;
    children?: Snippet;
    class?: string;
  } = $props();
</script>

<div class={cn('relative', className)}>
  {@render children?.()}
</div>
