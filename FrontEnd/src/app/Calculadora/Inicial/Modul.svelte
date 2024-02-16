<script lang="ts">
    import axios, { AxiosResponse } from 'axios';
    import type { HistoryEntry } from './components/Historial.svelte';
    import { onMount } from 'svelte';
    import Display from './components/Display.svelte';
    import BotonNumerico from './components/BotonNumerico.svelte';
    import BotonOperacion from './components/BotonOperacion.svelte';

    let input: string = '';
    let result: number = 0;
    let history: HistoryEntry[] = [];

    $: {
        input;
        result;
        history;
    }

    function handleNumericInput(value: number): void {
        input += value.toString();
    }

    function handleOperationInput(operation: string): void {
        if (input.length > 0) {
            //  const n1 = parseFloat(input.split(" ")[0]).toFixed(2);
            //  const n2 = parseFloat(input.split(" ")[2]).toFixed(2);
            input += ` ${operation} `;
        }
    }

    async function fetchHistory() {
        try {
            console.log('Fetching history...');
            const response: AxiosResponse<HistoryEntry[]> = await axios.get('http://localhost:9090/Historial');

            console.log('Response data:', response.data);
            history = response.data;
        } catch (error) {
            console.error('Error fetching history:', error);
        }
    }

    onMount(() => {
        fetchHistory();
    });

    async function handleButtonClick(event: { detail: string | number }): Promise<void> {
        const clickedValue = event.detail;

        if (clickedValue === '=') {
            if (input.length > 0) {
                try {
                    console.log('Fetching history...');
                    const response: AxiosResponse<{ resultado: number }> = await axios.post('http://localhost:3030/Calcular', {
                        // n1: parseFloat(input.split(" ")[0]),
                        // n2: parseFloat(input.split(" ")[2]),
                        operacion: input.split(' ')[1],
                    });
                    console.log('Response status:', response.status);
                    console.log('Response data:', response.data);
                    result = response.data.resultado;

                    // Reiniciar el input después de mostrar el resultado
                    input = '';

                    // Actualizar historial
                    fetchHistory();
                } catch (error) {
                    console.error('Error calculating:', error);
                }
            }
        } else if (clickedValue === 'C') {
            // Resetear input and result
            input = '';
            result = 0;
        } else if (typeof clickedValue === 'number') {
            handleNumericInput(clickedValue);
        } else if (typeof clickedValue === 'string') {
            handleOperationInput(clickedValue);
        }
    }
</script>

<div class="grid grid-cols-2 gap-4">
    <div class="col-span-1 p-10 flex items-center justify-center">
        <div class=" grid bg-white border-2 border-black">
            <div class="bg-gray-200 text-black font-bold text-center text-3xl p-10 h-30">Calculadora Basica</div>
            <!--Contenedor elementos calculadora-->
            <div class="p-10">
                <!-- Display -->
                <div class="mt-1 justify-center m-1">
                    <Display bind:input bind:result />
                </div>
                <!--Teclado-->
                <div class="grid grid-cols-4 grid-rows-5">
                    <!-- Fila 1 -->

                    <BotonOperacion operation="CE" on:buttonClick={handleButtonClick} style="background-color:orange" />
                    <BotonOperacion operation="C" on:buttonClick={handleButtonClick} style="background-color:orange" />
                    <BotonOperacion operation="+/-" on:buttonClick={handleButtonClick} style="background-color:red" />
                    <BotonOperacion operation="%" on:buttonClick={handleButtonClick} style="background-color:red" />

                    <!-- Fila 2 -->

                    <BotonNumerico value={7} on:buttonClick={handleButtonClick} />
                    <BotonNumerico value={8} on:buttonClick={handleButtonClick} />
                    <BotonNumerico value={9} on:buttonClick={handleButtonClick} />
                    <BotonOperacion operation="÷" on:buttonClick={handleButtonClick} style="background-color:red" />

                    <!-- Fila 3 -->

                    <BotonNumerico value={4} on:buttonClick={handleButtonClick} />
                    <BotonNumerico value={5} on:buttonClick={handleButtonClick} />
                    <BotonNumerico value={6} on:buttonClick={handleButtonClick} />
                    <BotonOperacion operation="x" on:buttonClick={handleButtonClick} style="background-color:red" />

                    <!-- Fila 4 -->

                    <BotonNumerico value={1} on:buttonClick={handleButtonClick} />
                    <BotonNumerico value={2} on:buttonClick={handleButtonClick} />
                    <BotonNumerico value={3} on:buttonClick={handleButtonClick} />
                    <BotonOperacion operation="-" on:buttonClick={handleButtonClick} style="background-color:red" />

                    <!-- Fila 5 -->

                    <BotonNumerico value={0} on:buttonClick={handleButtonClick} />
                    <BotonOperacion operation="." on:buttonClick={handleButtonClick} style="background-color:blue" />
                    <BotonOperacion operation="=" on:buttonClick={handleButtonClick} style="background-color:green" />
                    <BotonOperacion operation="+" on:buttonClick={handleButtonClick} style="background-color:red" />
                </div>
            </div>
        </div>
    </div>

    <!--Contenedor Historial-->
    <div class="col-span-1 p-10 flex justify-center">
        <!-- Historial -->
        <div class="bg-white border-2 border-black">
            <div class="bg-gray-200 p-10 text-black font-bold text-center text-3xl">Calculadora Basica</div>
            <h2 class=" font-bold mb-4 text-center text-3xl pt-8">Historial</h2>
            <div class="p-2">
                <table class="w-full bg-black text-white text-center items-center">
                    <thead>
                        <tr class="justify-center items-center">
                            <th class="border-b-2 border-gray-400 py-2 text-center">Fecha</th>
                            <th class="border-b-2 border-gray-400 py-2 text-center">Operación</th>
                            <th class="border-b-2 border-gray-400 py-2 text-center">Resultado</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#if history && history.length > 0}
                            {#each history as entry}
                                <tr class="border-b border-gray-300">
                                    <td class="py-2 text-center">{entry.fecha}</td>
                                    <td class="py-2 text-center">{entry.operacion}</td>
                                    <td class="py-2 text-center">{entry.resultado}</td>
                                </tr>
                            {/each}
                        {/if}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
