<script>
    import Connect from "./components/Connect.svelte";
    import Browser from "./components/Browser.svelte";

    let connected = $state(false);
</script>

<main>
    {#if !connected}
        <Connect />
    {:else}
        <Browser />
    {/if}
</main>