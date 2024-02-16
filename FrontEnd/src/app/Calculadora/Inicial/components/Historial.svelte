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
        {#each history as entry}
            <li>{entry.fecha} - {entry.operacion} = {entry.resultado}</li>
        {/each}
    </ul>
</div>
