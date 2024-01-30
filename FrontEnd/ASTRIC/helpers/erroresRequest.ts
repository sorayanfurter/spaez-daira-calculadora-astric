import { envMensages } from '../config/env'

export let errorRequest = (error: any) => {

    let res: any

    if (error.code === "ERR_NETWORK") {
        res = {
            datos: null,
            status: error.response.status,
            statusText: error.response.statusTex,
            tag: null,
            resolve: false,
            error: envMensages.falloConexion
        }
        return res
    }

    if (error.code === "ECONNABORTED") {
        res = {
            datos: null,
            status: 0,
            statusText: error.message,
            tag: null,
            resolve: false,
            error: envMensages.falloTimeout
        }
        return res
    }

    if (error.code === "ERR_BAD_REQUEST") {

        res = {
            datos: error.response.data ? error.response.data : null,
            status: error.response.status,
            statusText: error.response.statusText,
            tag: null,
            resolve: false,
            error: envMensages.falloRuta
        }
        return res

    }

    res = {
        datos: null,
        status: 0,
        statusText: null,
        tag: null,
        resolve: false,
        error: envMensages.falloDefault
    }

    return res
}