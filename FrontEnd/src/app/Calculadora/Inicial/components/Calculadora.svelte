<script lang="ts">
    import { http } from '@astric';
    import { onMount } from 'svelte';
    import type { HistoryEntry } from './Historial.svelte';
    import Display from './Display.svelte';
    import BotonNumerico from './BotonNumerico.svelte';
    import BotonOperacion from './BotonOperacion.svelte';

    let botones = [
        ['CE', 'C', '+/-', '%'],
        ['7', '8', '9', '÷'],
        ['4', '5', '6', 'x'],
        ['1', '2', '3', '-'],
        ['0', '.', '=', '+'],
    ];
    let input = '';
    let result = 0;
    let history: HistoryEntry[] = [];

    function getColor(operation: string) {
        // asignar colores a las operaciones
        switch (operation) {
            case 'CE':
                return 'orange';
            case 'C':
                return 'orange';
            case '+/-':
                return 'red';
            case '%':
                return 'red';
            case '÷':
                return 'red';
            case 'x':
                return 'red';
            case '-':
                return 'red';
            case '=':
                return 'green';
            case '+':
                return 'red';
            default:
                return 'blue';
        }
    }

    async function fetchHistory() {
        try {
            http.get('calculadora/Historial')
                .then((response: any) => {
                    if (response.datos !== null) {
                        history = response.datos;
                    }
                })
                .catch(error => {
                    console.error('Error fetching history:', error);
                });
        } catch (error) {
            console.error('Error fetching history:', error);
        }
    }

    onMount(() => {
        fetchHistory();
    });

    function handleButtonClick(value: string) {
        switch (value) {
            case 'CE':
                input = '';
                result = 0;
                break;
            case 'C':
                input = input.slice(0, -1);
                break;
            case '+/-':
                input = input.startsWith('-') ? input.slice(1) : `-${input}`;
                break;
            case '%':
                input = `${parseFloat(input) / 100}`;
                break;
            case '=':
                if (input !== undefined && typeof input === 'string' && input.length > 0) {
                    try {
                        http.post('calculadora/Calcular', {
                            n1: parseFloat(input.split(' ')[0]),
                            n2: parseFloat(input.split(' ')[2]),
                            operacion: input.split(' ')[1],
                        })
                            .then((response: any) => {
                                const datos = response.datos;
                                if (datos && datos.resultado !== undefined) {
                                    result = datos.resultado;

                                    if (datos.history !== undefined) {
                                        history = datos.history;
                                    }

                                    input = '';
                                } else {
                                    console.error('Invalid response structure:', response);
                                }
                            })
                            .catch((err: any) => {
                                console.log(err);
                            });
                    } catch (error) {
                        console.error('Error fetching history:', error);
                    }
                }
                break;
            case '+':
            case '-':
            case 'x':
            case '÷':
                input += ` ${value} `;
                break;
            default:
                input += value;
                break;
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
                    <Display {input} {result} />
                </div>
                <!--Teclado-->
                <div class="grid grid-cols-4 grid-rows-5">
                    {#each botones as fila}
                        {#each fila as boton}
                            {#if typeof boton === 'number'}
                                <BotonNumerico value={boton} on:click={() => handleButtonClick(boton)} />
                            {:else}
                                <BotonOperacion operation={boton} on:click={() => handleButtonClick(boton)} style={`background-color: ${getColor(boton)}`} />
                            {/if}
                        {/each}
                    {/each}
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
                            {#each history as dato}
                                <tr class="border-b border-gray-300">
                                    <td class="py-2 text-center">{dato.fecha}</td>
                                    <td class="py-2 text-center">{dato.operacion}</td>
                                    <td class="py-2 text-center">{dato.resultado}</td>
                                </tr>
                            {/each}
                        {/if}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>
