import type { AxiosRequestConfig, AxiosResponse } from 'axios'
import { errorRequest } from './erroresRequest'
import { env } from '../config/env'
import { dbStore } from '@astric-store'

let error_request = (error: any) => {
    return Promise.reject(error)
}


let error_response = (error: any) => {

    let res: any = errorRequest(error)

    if (env.dev) {
        if (error.response?.data?.MSG) res.MSG = error.response.data.MSG
        console.group("ERROR url:", error.config.url)
        console.warn("RESPONSE ---:", res)
        console.warn("INFO ---:", error)
        console.groupEnd()
    }
    return Promise.reject(res)
}

let fn_request = (request: AxiosRequestConfig) => {


    if (env.dev) {
        request.headers = {
            "Authorization": env.user ? env.user : "default",
            "Database": dbStore.db,
            "Accept": "application/json, text/plain, */*"
        }
        return request
    }

    let tkn = localStorage.getItem("tkn") != null ? localStorage.getItem("tkn") : ""

    request.headers = {
        "Authorization": tkn != null ? tkn : "",
        "Database": dbStore.db,
        "Accept": "application/json, text/plain, */*"
    }

    return request
}



let fn_response = (response: AxiosResponse) => {

    let res: any = {
        datos: response.data.datos,
        status: response.status,
        statusText: response.statusText,
        tag: response.data.tag,
        resolve: response.data.status,
    }
    if (env.dev) {
        if (response.data.MSG) res.MSG = response.data.MSG
        console.info("RESPONCE url:", response.config.url, res)
    }
    return res
}



export let request = {
    fn_request,
    error_request
}

export let response = {
    fn_response,
    error_response
}


