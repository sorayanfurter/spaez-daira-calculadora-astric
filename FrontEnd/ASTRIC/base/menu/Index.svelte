<script lang="ts">
    import Sider from './Sider.svelte';
    import Menu from './Menu.svelte';
    import { menuStore } from '../store'


    let menu: any[];
    let databses: any[];

    menuStore.subscribe((data) => {
        menu = data.items
        databses = data.databases
    })
</script>

<Sider {databses} />
<Menu {menu} />
