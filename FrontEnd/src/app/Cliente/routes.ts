

// @ts-nocheck
import { wrap } from 'svelte-spa-router/wrap'

const routes =
{
    '/NuevoCliente': wrap({ asyncComponent: () => import('./NuevoCliente/Modul.svelte') }),
    '/Listar': wrap({ asyncComponent: () => import('./Listar/Modul.svelte') }),
};

export default routes;
