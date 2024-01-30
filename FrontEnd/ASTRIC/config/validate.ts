import { setLocale } from 'yup';

export let configValidator = () => {


    setLocale({
        mixed: {
            default: 'No valido',
            required: 'El valor es obligatorio',
            defined: 'No esta definido',
            notOneOf: 'notOneOf',
            notType: 'No es el tipo de datos requerido',
            oneOf: 'oneOf'
        },
        number: {
            min: 'Debe ser mayor a ${min}',
            max: 'Debe ser menor a ${max}',
            integer: 'Debe ser un entero',
            lessThan: "Debe ser anterior a ${lessThan}",
            moreThan: "Debe estar despues de ${moreThan}",
            negative: "Debe ser un numero negativo",
            positive: "Debe ser un numero positov"
        },
        string: {
            email: 'El correo no es valido',
            length: 'La cantidad de caracteres es invlida',
            min: 'Debe ser mayor a ${min}',
            max: 'Debe ser menor a ${max}',
            lowercase: 'Todos los caracteres deben ser minusculas',
            uppercase: "Todos los caracteres deben ser mayuscula",
            matches: "Debe contener los siguientes caracteres: ${matches}",
            trim: "No debe tener espacios al inicio ni al fial",
            url: "URL no valida",
            uuid: "Numero de identificacion no valido"
        },
        date: {
            min: 'Debe ser mayor a ${min}',
            max: 'Debe ser menor a ${max}',
        },
        boolean: {
            isValue: 'No valido'
        },
        array: {
            length: 'El tamaño debe ser ${length}',
            min: 'Debe ser mayor a ${min}',
            max: 'Debe ser menor a ${max}',
        },
        object: {
            noUnknown: 'El objeto esta vacío'
        }
    });
}

