import { writable } from 'svelte/store';
import { env } from '../../ASTRIC/config/env'
import { db } from '@astric-env'


let items: any[] = [];
let databases: any[] = [];

let menu = {
    open: false,
    items: items,
    databases: databases
}

let socket = {
    //App sirve como identificador de app
    app: "",
    //EndPoint sirve como identificador del endpoint
    endPoint: "",
    //Tag accion o transaccion realizada
    tag: "",
    //BD base de datos en la cual se produjo el cambio
    BD: "",
    //Tabla de la base de datos en la cual se produjo el cambio
    tabla: "",
    //Usuario usuario que produjo el evento
    usuario: "",
    //Data datos enviados si es nesesario
    data: "",
    //Msj mensaje si es nesesario
    msj: "",
    //Error si es nesessario
    error: ""
}

export let menuAccion = {
    open: () => {
        menuStore.update(value => {
            value.open = !value.open
            return value
        })
    }
}

export let setMenu = {
    item: (name: string, route: string) => {
        menuStore.update(value => {
            value.items.push({ name, route })
            return value
        })
    },
    items: (items: any[]) => {
        menuStore.update(value => {
            value.items = items
            return value
        })
    },
    database: (name: string, db: string) => {
        menuStore.update(value => {
            value.databases.push({ name, db })
            return value
        })
    },
    databases: (databases: any[]) => {
        menuStore.update(value => {
            value.databases = databases
            return value
        })
    },
}


export let dbStore = {
    db: db.default,
    name: db.name,
}

let login = {
    loguin: !env.login,
    userName: "Usuario Administrado",
    roll: "Administrador"
}

export let wsAccion = {
    set: (dato: any) => {
        ws.update(value => {
            const datos = JSON.parse(dato.data)
            value.app = datos.App
            value.BD = datos.BD
            value.data = datos.Data
            value.endPoint = datos.EndPoint
            value.error = datos.Error
            value.msj = datos.Msj
            value.tabla = datos.Tabla
            value.tag = datos.Tag
            value.usuario = datos.Usuario
            return value
        })
    }
}

export let dbAccion = {
    set: (name: string, db: string) => {
        dbStore.db = db
        database.update(value => {
            value.db = db
            value.name = name
            return value
        })
    }
}

export let loguinAccion = {
    login: (userName: string, roll: string) => {
        loguin.update(value => {
            value.loguin = true
            value.userName = userName
            value.roll = roll
            return value
        })
    },
    unLoguin: () => {
        loguin.update(value => {
            value.loguin = false
            value.userName = ""
            value.roll = ""
            return value
        })
    }
}

export let ws = writable(socket)
export let loguin = writable(login)
export let database = writable(dbStore)
export let menuStore = writable(menu)
