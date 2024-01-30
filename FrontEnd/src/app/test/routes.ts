
// @ts-nocheck
import { wrap } from 'svelte-spa-router/wrap'

const routes = { '/test': wrap({ asyncComponent: () => import('./test/Modul.svelte') }),};

export default routes;
