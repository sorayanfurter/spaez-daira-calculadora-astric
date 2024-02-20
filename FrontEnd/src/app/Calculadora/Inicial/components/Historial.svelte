<script context="module" lang="ts">
    export interface HistoryEntry {
        fecha: string;
        operacion: string;
        resultado: number;
    }
</script>

<script lang="ts">
    export let history: HistoryEntry[] = [];
</script>

<div>
    <h2>Historial</h2>
    <ul>
        {#each history as data}
            <li>{data.fecha} - {data.operacion} = {data.resultado}</li>
        {/each}
    </ul>
</div>
