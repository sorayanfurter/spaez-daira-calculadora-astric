<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import { input } from '@/Calculadora/shared/store';
    export let value: number;

    const dispatch = createEventDispatcher();

    function handleClick() {
        input.update(currentValue => `${currentValue}${value}`);
        dispatch('buttonClick', value);
    }
</script>

<button style="background-color: blue;" class="text-white font-bold m-1 h-20 w-20" on:click={handleClick}>{value}</button>
