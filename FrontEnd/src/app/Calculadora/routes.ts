
// @ts-nocheck
import { wrap } from 'svelte-spa-router/wrap'

const routes = { '/Inicial': wrap({ asyncComponent: () => import('./Inicial/Modul.svelte') }),};

export default routes;
