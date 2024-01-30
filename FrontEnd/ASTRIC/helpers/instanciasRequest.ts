import axios from 'axios';
import type { AxiosRequestConfig, AxiosRequestHeaders, AxiosInstance } from 'axios'

let httpInstans = (baseURL: string, timeout: number, headers: AxiosRequestHeaders): AxiosInstance => {
    let config: AxiosRequestConfig = {
        baseURL,
        timeout,
        headers
    }
    return axios.create(config);
}

let httpInstansConfig = (config: AxiosRequestConfig): AxiosInstance => {
    return axios.create(config);
}

export let instance = {
    httpInstans,
    httpInstansConfig
}