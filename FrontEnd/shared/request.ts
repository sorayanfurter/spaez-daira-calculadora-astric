import axios from 'axios'
import { instance } from '../ASTRIC/helpers/instanciasRequest'

export let http = {
    get: axios.get,
    post: axios.post,
    put: axios.put,
    delete: axios.delete,
    instance,
}



