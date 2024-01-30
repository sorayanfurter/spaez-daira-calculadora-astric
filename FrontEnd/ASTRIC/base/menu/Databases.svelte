<script lang="ts">
    import type { MenuComponentDev } from '@smui/menu';
    import { dbAccion } from '../store';
    import env from '@astric-env';
    import Menu from '@smui/menu';
    import List, { Item, Text } from '@smui/list';
    import Button, { Label } from '@smui/button';
    export let databses: any = [];
    let menu: MenuComponentDev;
</script>

<div style="min-width: 100px;">
    <Button on:click={() => menu.setOpen(true)}>
        <Label>{env.menu.menu_databases_name}</Label>
    </Button>
    <Menu bind:this={menu}>
        <List>
            {#each databses as sucursal}
                <Item
                    on:SMUI:action={() => {
                        dbAccion.set(sucursal.name, sucursal.db);
                    }}
                >
                    <Text>{sucursal.name}</Text>
                </Item>
            {/each}
        </List>
    </Menu>
</div>
