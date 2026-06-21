<script lang="ts">
  import type { Snippet } from 'svelte';
  import { cn } from '$lib/utils.js';

  let {
    children,
    class: className,
    align = 'center'
  }: {
    children?: Snippet;
    class?: string;
    align?: 'start' | 'center' | 'end';
  } = $props();

  const alignClasses = {
    start: 'left-0',
    center: 'left-1/2 -translate-x-1/2',
    end: 'right-0'
  };
</script>

<div class={cn(
  'absolute z-50 mt-2 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md',
  alignClasses[align],
  className
)}>
  {@render children?.()}
</div>
