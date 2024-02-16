<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import { input } from '@/Calculadora/shared/store';

    export let operation: string;
    export let style = '';

    const dispatch = createEventDispatcher();

    function handleClick() {
        input.update(currentValue => `${currentValue}${operation}`);
        dispatch('buttonClick', operation);
    }
</script>

<button class="text-white font-bold m-1 h-20 w-20" on:click={handleClick} {style}>{operation} </button>
