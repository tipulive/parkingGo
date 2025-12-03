<script>
    import { selectedPost } from '$lib/store/store.js';
    $inspect(selectedPost.subscribe.name);
</script>

{#if $selectedPost} <!-- This will not work -->
    <h2>{$selectedPost}</h2>
{/if}