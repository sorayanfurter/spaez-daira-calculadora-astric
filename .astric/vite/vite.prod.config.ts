import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import tsconfigPaths from 'vite-tsconfig-paths'
import config from '../../FrontEnd/config.json';
import prod from '../config/fronEnd.json'
import path from 'path'

const envProd = {
    'process.env.FALLO_CONEXION': JSON.stringify(config.mensajes.fallo_conexion),
    'process.env.FALLO_RUTA': JSON.stringify(config.mensajes.fallo_ruta),
    'process.env.TIMEUOT': JSON.stringify(config.mensajes.fallo_timeuot),
    'process.env.DEFAULT': JSON.stringify(config.mensajes.fallo_default),
    'process.env.URL': JSON.stringify(prod.api_url),
    'process.env.WS_URL': JSON.stringify(prod.ws_url),
    'process.env.WS_EDABLE': JSON.stringify(prod.websocket),
    'process.env.WS_TIME': JSON.stringify(prod.ws_time),
    'process.env.DEV': false,
    'process.env.LOGUIN': JSON.stringify(prod.loguin),
    'process.env.REQUEST_TIMEOUT': JSON.stringify(prod.request_timeuot_sec)
}

const alias = {
    "@astric": `${path.resolve(__dirname, '../../FrontEnd/shared/imports/controlers.ts')}`,
    "@astric-validate": `${path.resolve(__dirname, '../../FrontEnd/shared/imports/validate.ts')}`,
    "@astric-login": `${path.resolve(__dirname, '../../FrontEnd/shared/imports/login.ts')}`,
    "@astric-env": `${path.resolve(__dirname, '../../FrontEnd/config.json')}`,
    "@astric-store": `${path.resolve(__dirname, '../../FrontEnd/ASTRIC/base/store.ts')}`,
    "@": `${path.resolve(__dirname, '../../FrontEnd/src/app')}/`
}

export default defineConfig(() => {
    return {
        server: {
            host: prod.host,
            port: prod.port,
            cors: prod.cords
        },
        plugins: [
            svelte(),
            tsconfigPaths()
        ],
        define: envProd,
        resolve: {
            alias: alias
        }
    }
})