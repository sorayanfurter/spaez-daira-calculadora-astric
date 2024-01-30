<script lang="ts">
    //import { menuAccion, menuStore, loguin } from '../store';
    //import { push } from 'svelte-spa-router';
    //import Drawer, { Content, Header, Title, Subtitle, Scrim } from '@smui/drawer';
    //import List, { Item, Text } from '@smui/list';
    export let menu: any = [];
    //let open = false;

    //menuStore.subscribe((data: any) => (open = data.open));

    /*
    function setActive(value: string) {
        push(value);
        menuAccion.open();
    }*/
</script>

<div class="drawer-container">
    <!--
    <Drawer variant="modal" fixed={true} bind:open>
        <Header>
            <Title>{$loguin.userName}</Title>
            <Subtitle>{$loguin.roll}</Subtitle>
        </Header>
        <Content>
            <Item href="javascript:void(0)" on:click={() => setActive('/')}>
                <Text>HOME</Text>
            </Item>
            <List>
                {#each menu as m}
                    <Item href="javascript:void(0)" on:click={() => setActive(m.route)}>
                        <Text>{m.name}</Text>
                    </Item>
                {/each}
            </List>
        </Content>
    </Drawer>
    <Scrim on:click={() => menuAccion.open()} fixed={false} />
        -->
</div>
