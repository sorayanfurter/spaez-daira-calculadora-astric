<script lang="ts">
    import '../../ASTRIC/start';
    import { loguin } from './store';
    import Menu from './menu/Index.svelte';
    import Routes from '../../src/routes.svelte';
    import Login from './login/login.svelte';
</script>

<div class="main">
    {#if $loguin.loguin}
        <Menu />
        <Routes />
    {:else}
        <Login />
    {/if}
</div>

<style>
    .main {
        margin: auto;
    }
</style>
