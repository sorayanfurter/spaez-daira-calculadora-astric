https://axios-http.com/docs/intro
https://sveltematerialui.com/
https://imbrn.dev/v8n/#what-s-v8n 
https://github.com/jquense/yup
https://svelte-forms-lib-sapper-docs.vercel.app/yup
https://momentjs.com/
https://www.npmjs.com/package/svelte-chartjs
https://mathjs.org/docs/index.html

### Estructura de componentes

Para usar el CLI npm run cli 

- Nombre del componente
    - components
        - comp1
            - helpers
            - interfaces
            - schemas
            - validatos
            - Componente.svelte
        - comp2
            - helpers
            - interfaces
            - schemas
            - validatos
            - Componente.svelte
    - shared
    - main.svelte
    - routes.ts

### Archvos comunes

# main.svelte

<script lang="ts">
    import Router from 'svelte-spa-router';
    import routes from './routes';
    export let prefix: string;
</script>

<main><Router {prefix} {routes} /></main>

# routes.ts

import { wrap } from 'svelte-spa-router/wrap'

const routes = {};

export default routes;

- Ejemplo de ruta
'/http': wrap({ asyncComponent: () => import('./components/prueba/Http.svelte') }),

# store.ts

import { writable } from 'svelte/store';

export let store = writable()

### Importaciones

@astric          -- Utilierias del fram
@astric-validate -- Utilerias para validaciones 
@astric-env      -- Variables de entorno
@astric-store    -- Storage principal
@                -- Ruta a app

