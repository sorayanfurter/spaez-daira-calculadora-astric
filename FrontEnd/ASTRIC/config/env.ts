// @ts-nocheck
export let envRequest = {
    baseUrl: process.env.URL,
    timeOut: process.env.REQUEST_TIMEOUT
}

export let env = {
    dev: process.env.DEV,
    user: process.env.USER,
    login: process.env.LOGUIN
}

export let envMensages = {
    falloConexion: process.env.FALLO_CONEXION,
    falloRuta: process.env.FALLO_RUTA,
    falloTimeout: process.env.TIMEUOT,
    falloDefault: process.env.DEFAULT
}

export let envWebSocket = {
    time: process.env.WS_TIME,
    enable: process.env.WS_EDABLE,
    URL: process.env.WS_URL
}

