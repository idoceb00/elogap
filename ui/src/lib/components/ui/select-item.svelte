<script lang="ts">
  import type { Snippet } from 'svelte';
  import { cn } from '$lib/utils.js';

  let {
    value,
    children,
    class: className
  }: {
    value: string;
    children?: Snippet;
    class?: string;
  } = $props();
</script>

<div class={cn('relative flex w-full cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50', className)}>
  {@render children?.()}
</div>
