<script lang="ts">
  import { cn } from '$lib/utils.js';

  let {
    class: className
  }: {
    class?: string;
  } = $props();
</script>

<div class={cn('-mx-1 my-1 h-px bg-muted', className)}></div>
