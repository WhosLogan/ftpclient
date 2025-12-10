<script lang="ts">
    import { Button, Input, Checkbox, Modal, P } from 'flowbite-svelte';
    import {Connect} from "../../wailsjs/go/main/App";

    let anon = $state(false);
    let host = $state('');
    let username = $state('');
    let password = $state('');
    let error = $state('');
    let open = $state(false);
    let loading = $state(false);

    let {onConnect} = $props();

    async function connect() {
        loading = true;
        try {
            await Connect(host, username, password, anon);
            onConnect();
        } catch (e) {
            error = e;
            open = true;
        }
        loading = false;
    }
</script>

<div class="h-screen flex justify-center items-center">
    <div class="flex flex-col space-y-3 md:w-1/3">
        <h1 class="text-2xl text-center text-white">FTP Client</h1>
        <Input bind:value={host} placeholder="Host" />
        <Input bind:value={username} disabled={anon} placeholder="Username" />
        <Input bind:value={password} disabled={anon} placeholder="Password" type="password" />
        <Checkbox bind:checked={anon}>Use Anonymous</Checkbox>
        <Button disabled={loading} onclick={() => connect()}>Connect</Button>
    </div>
</div>

<Modal title="An error has occurred" form bind:open={open}>
    <P>{error}</P>

    {#snippet footer()}
        <Button type="submit">Ok</Button>
    {/snippet}
</Modal>