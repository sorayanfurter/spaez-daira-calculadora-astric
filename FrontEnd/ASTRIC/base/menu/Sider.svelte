<script lang="ts">
    //import { menuAccion, database } from '../store';
    //import env from '@astric-env';
    //import Databases from './Databases.svelte';
    //import TopAppBar, { Row, Section, Title } from '@smui/top-app-bar';
    //import IconButton from '@smui/icon-button';
    export let databses: any = [];
</script>

<div class="flexy">
    <div class="top-app-bar-container">
        <!--
        <TopAppBar variant="static" dense>
            <Row>
                <Section>
                    <IconButton
                        on:click={() => menuAccion.open()}
                        class="material-icons">menu</IconButton
                    >
                    <Title>{$database.name}</Title>
                </Section>
                {#if env.menu.menu_database_enable}
                    <Section align="end">
                        <Databases {databses} />
                    </Section>
                {/if}
            </Row>
        </TopAppBar>
        -->
    </div>
</div>
