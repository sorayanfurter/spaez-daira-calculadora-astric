package routes

import "ASTRIC/BackEnd/ASTRIC/helper/routes"

// TipoRuta tipo ruta
type TipoRuta routes.TipoRuta

// TipoAccion tipo accion
type TipoAccion routes.TipoAccion

// TipoService tipo service
type TipoService routes.TipoService

// RoutersStruct tipo ruta struct
type RoutersStruct routes.RoutersStruct

/*
DisableSecurity Desactiva la seguridad de la ruta
@ruta: la ruta propiamente dicha
*/
var DisableSecurity = routes.DisableSecurity

/*
DisableSucrsal desactiva el requerimiento de la sucursal en la ruta
@ruta: La direccion de la ruta propiamente dicha
*/
var DisableSucrsal = routes.DisableSucrsal
