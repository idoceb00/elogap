<script lang="ts">
  import Card from '$lib/components/ui/card.svelte';
  import CardContent from '$lib/components/ui/card-content.svelte';

  let { title, description }: { title: string; description: string } = $props();
</script>

<Card class="border-primary/20 bg-primary/5">
  <CardContent class="p-6">
    <p class="font-medium">{title}</p>
    <p class="text-sm text-muted-foreground mt-1">{description}</p>
  </CardContent>
</Card>
