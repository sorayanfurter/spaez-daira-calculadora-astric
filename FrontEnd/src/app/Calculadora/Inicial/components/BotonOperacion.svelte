<script lang="ts">
    export let operation: string;
    export let style = '';
</script>

<button class="text-white font-bold m-1 h-20 w-20" on:click {style}>{operation} </button>
