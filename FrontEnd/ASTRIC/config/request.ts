import axios from 'axios';
import { envRequest } from './env'
import { request, response } from '../helpers/interceptor'
import { requestValidate } from '../helpers/validateRequest'

export let configHTTP = () => {
    axios.defaults.baseURL = envRequest.baseUrl ? envRequest.baseUrl : "";
    axios.defaults.headers.post['Content-Type'] = 'application/json; charset=utf-8';
    axios.defaults.timeout = parseFloat(envRequest.timeOut ? envRequest.timeOut : "3") * 1000;
    axios.defaults.validateStatus = requestValidate
};

export let configInterceptor = () => {
    axios.interceptors.request.use(request.fn_request, request.fn_request)
    axios.interceptors.response.use(response.fn_response, response.error_response)
}
