<script>
    import Connect from "./components/Connect.svelte";
    import Browser from "./components/Browser.svelte";
    import {EventsOn} from "../wailsjs/runtime/runtime.js";

    let connected = $state(false);

    EventsOn("connection-change", (connected) => {
        connected = connected;
    });
</script>

<main class="select-none">
    {#if !connected}
        <Connect onConnect={() => connected = true} />
    {:else}
        <Browser onDisconnect={() => connected = false} />
    {/if}
</main>