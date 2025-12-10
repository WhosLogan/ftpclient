<script lang="ts">
    import {Button, Modal, P} from "flowbite-svelte";
    import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from "flowbite-svelte";
    import {ChangeDirectory, Close, ListDirectory} from "../../wailsjs/go/main/App";

    let {onDisconnect} = $props();
    let listings = $state([]);
    let modalBody = $state("");
    let modalOpen = $state(false);
    let modalTitle = $state("");

    async function refresh() {
        let contents = await ListDirectory();

        contents.sort((a, b) => {
            if (a.IsDir && !b.IsDir) return -1;
            if (!a.IsDir && b.IsDir) return 1;
            return a.Name.localeCompare(b.Name);
        });

        listings = contents;
    }

    async function changeDir(dir: string) {
        try {
            await ChangeDirectory(dir);
            await refresh();
        } catch (e) {
            modalTitle = "Error";
            modalBody = e;
            modalOpen = true;
        }
    }

    async function disconnect() {
        await Close();
        onDisconnect();
    }

    refresh();
</script>

<div class="relative h-screen">
    <div class="overflow-y-auto max-h-[calc(100%-80px)]">
        <Table striped={true}>
            <TableHead>
                <TableHeadCell>Name</TableHeadCell>
                <TableHeadCell>Directory</TableHeadCell>
                <TableHeadCell>Permissions</TableHeadCell>
                <TableHeadCell>Size</TableHeadCell>
                <TableHeadCell>
                    <span class="sr-only">Actions</span>
                </TableHeadCell>
            </TableHead>
            <TableBody>
                <TableBodyRow>
                    <TableBodyCell>..</TableBodyCell>
                    <TableBodyCell>True</TableBodyCell>
                    <TableBodyCell>N/A</TableBodyCell>
                    <TableBodyCell>N/A</TableBodyCell>
                    <TableBodyCell>
                        <button onclick={() => changeDir("..")}
                                class="text-primary-600 dark:text-primary-500 font-medium hover:underline">Open</button>
                    </TableBodyCell>
                </TableBodyRow>

                {#each listings as item}
                    <TableBodyRow>
                        <TableBodyCell>{item.Name}</TableBodyCell>
                        <TableBodyCell>{item.IsDir ? "True" : "False"}</TableBodyCell>
                        <TableBodyCell>{item.Perms}</TableBodyCell>
                        <TableBodyCell>{item.Size}</TableBodyCell>
                        <TableBodyCell>
                            {#if item.IsDir}
                                <button onclick={() => changeDir(item.Name)}
                                        class="text-primary-600 dark:text-primary-500 font-medium hover:underline">Open</button>
                            {:else}
                                <button class="text-primary-600 dark:text-primary-500 font-medium hover:underline">Download</button>
                            {/if}
                        </TableBodyCell>
                    </TableBodyRow>
                {/each}
            </TableBody>
        </Table>
    </div>

    <div class="absolute bottom-0">
        <div class="p-4">
            <Button onclick={disconnect}>Disconnect</Button>
            <Button onclick={refresh}>Refresh</Button>
        </div>
    </div>
</div>


<Modal title={modalTitle} form bind:open={modalOpen}>
    <P>{modalBody}</P>

    <Button type="submit">Ok</Button>
</Modal>