# Frameworks ASTRIC 

## Proyecto _______

## Configuracion de VSCs

<https://engineering.truora.com/posts/2020-02-18-como-configurar-vscode-para-golang>

>La configuración obligatoria se encuentra en `doc/SetupVSC.json`

## Configuracion de SWAGGER

>Lick de descarga para el generador <https://github.com/go-swagger/go-swagger/releases/>
>Documentacion <https://goswagger.io/use/spec.html>
>Comande de consola para el auto generado "swagger generate spec -o ./swagger.json --scan-models"

### Ejemplo de cometario de documentacion

- Para endpoint o controlador

---

// swagger:operation POST /admin/usuarioLoguin (Grupo de la ruta) (Nombre de la ruta)
// ---
// summary: (Descripción de la ruta)
// parameters:
// - name: request
//   in: body
//   description: (Descripción, ayuda, etc. del request)
//   required: true
//   schema:
//     "$ref": "#/definitions/(Modelo a mostrar como ejemplo en este caso loguin)"
// responses:
//   default:
//     description: Respuesta por defecto
//     schema:
//       "$ref": "#/definitions/Response"

---

- Para los modelos de datos

---

//Login modelo de login
//swagger:model (Nombre del modelo en este caso login, el mismo se usa en el $ref del endpoints)
type Login struct {

//swagger:ignore "Ignora la propiedad"

  ID                `bson:"id,omitempty" json:"id" validate:"required"`

// required: true "Hace la propiedad requerida"

 Usuario string     `bson:"usuario" json:"usuario" validate:"required"`

 Pass string        `bson:"pass" json:"pass" validate:"required"`

}

---

## Rutas

### Existen 3 tipos de rutas

**_Rutas:_** Estas son rutas comunes y estan dentro del menu de opciones se utilizan cuando la informacion obtenida es para mostrar un apantalla, se declaran en:

`rutes/routes.go`

func rutas(ruta routes.TipoRuta string) {

> ruta("/pruebaPrueba", controller.Prueba, "POST", "Descripcion de la ruta")

}

**_Acciones:_** Son rutas que no comprenden parte del menú, se utilizan para acciones de botones, por ejemplo eleiminar o guardar un registro, se declaran en:

`routes/acciones.go`

func acciones(accion routes.TipoAccion string) {

>accion("/pruebaAccion", controller.Accion, "POST", "Descripcion de la Accion")

}

**_Servicios:_** Estas son rutas de servicios, no requieren autorizacion, el usuario solo debe tener **autorizacion** al modulo donde se encuentra la ruta y estar logeado mediante token, por ejemplo rutas donde se solicitan valores predefinidos, se declaran en:

`routes/servises.go`

func services(service routes.TipoService) {

>service("/pruebaService", controller.Service, "POST")

}

## Controller

Estructura de un controller:

Solo pueden tener una funcion publica debidamete comentada, esta sera llamda desde las rutas. El resto de las funciones seran privadas y de uso interno del controlador, habra un controlador por cada ruta declarada.

Como primera instruccion obligatoria sera un **DEFER** con el controlador de errores correspondinte y la descripcion lo mas detallada y clara del controlador.

func Ejemlp(w http.ResponseWriter, r *http.Request) {

>defer ep.ErrorControlResponse(**"ruta de llamada/nombre del controlador/ detalles"**, w, r)

}

Para trabajar en los controladores deberasn importar el pakage ep **"ASTRIC/BackEnd/shared/ep"** el mismo contiene todas las errameintas e informacion nesesaria para realizar todas las tareas.

## Detalles del pakage ep

- DatabaseContexto() debuelve un contexto de trabajo

  ctx, cancel := ep.DatabaseContexto()

  defer cancel() **_Es nesesario para cancelar el contexto y que no se acumulen_**

- ErrorControlResponse() Funcion que controla los errores de tipo Panic Ej:

  >defer ep.ErrorControlResponse(**"ruta de llamada/nombre del controlador/ detalles"**, w, r)

  ---

  >Panic(**"Informacion del error producido"**)

- Reaponses: Generan respuestas APIS Vercion 1.0 "ESTE METODO ESTA DEPRECADO EN ESTA VERCION"

  >Response("Accion que se esta realizando", errores , "datos", w) "ESTE METODO ESTA DEPRECADO EN ESTA VERCION"

  ---

  Genera una respuesta request, si NO se envia ningun error la misma es 200 si se envia algun error es 401
  - Los errores puesden ser strings o objetos generados por funciones
  - Los datos son generalmente los que se recivieron desde el request
  - W es el request propiemate dicho generalmente w

  ---

  >ResponseMSG func("Accion que se esta realizando", errores, datos, w, **"Mensage"**) "ESTE METODO ESTA DEPRECADO EN ESTA VERCION"

  ---

  Genera un response con un mensage este aparece con la leyenda "1 - !ATENCION¡" mas el mensage escrito

  ---

  >ResponseUnauthorized func("Accion que se esta realizando", errores, datos, w) "ESTE METODO ESTA DEPRECADO EN ESTA VERCION"

  ---

  Genera una reapuesta igual a Responce pero con codigo 401

## Reaponses: Generan respuestas APIS Versión > 2.0

>res := ep.NewResponse( tag , ResponseWriter)
>res.Msg(mensage) Setea un mensage al responce
>res.Err(error) Setea un error al responce cambiando el header a codigo 401
>res.Dato(dato) setea el dato a enviar
Los metodos menciaonados son para configurar el responce, cualquier de los metodos anteriores precedidos por Send envia el responce EJ:
>res.DatoSend(dato) envia el dato
>res.Err(error).DatoSend(dato) envia el responce con el dato mas los errores con el codigo 401
>res.Err(error).Dato(dato).MsgSend("Mensaje") envia el responce idem al anterior egregando un mensaje

- Validacion de datos y estructuras

  - ValidateStruct(Estructura a validar) retorna los errores mas un bool true si ubo errres false si no ubo

    Las estructuras a validar se configuran con la instruccion validate: "parametros" ej:

  ---

    type Usuario struct {

    Nombre string `validate:"required"`

    Apellido string `validate:"required"`

    }

  **_Para validar estructuras anidadas es nesesario enviarlas por separado._**

  ---

## Datos de utilidad

- RequestInfo() Devuelve variables de entorno creadas por los request, ej: datos del usuario, etc.

---

## El Paquete BD

## Bases de datos

Las bases de datos se setan automaticamente de acurdo al request

- MongoDb() La bd de mango lista para trabajar

- MySQL() la bd de mysql lista para trabajar

- MysqlORM() ORM de mysql listo para trabajar

---

Si fuera nesesario trabajar con una bd distinta a la configurada desda el request podemos usar lo siguiente:

- MongoNewCN("database requerida") Crea una nueva conexion a la bd pasada como parametro
- MySQLNewCN("database requerida)" Crea una nueva conexion a la bd pasada como parametro
- MysqlORMNewCN("database requerida") Crea una nueva conexion a la bd pasada como parametro con ORM

## El paquete modesls/mg y models/my

Paquete que contiene todos los modelos de conexion a las bases de datos

>models/mg modelos de mongo db
>models/my modelos de mysql

## Mensajes

>Es posuble crear mensajes en varias sircunstancias durante el periodo de debag, estos son utiles para notificar o instruir a los desarrolladores del front-end

Los mismos se configuran desde `/Messages`

>`msgInicio.go` Estos mensages aparecen en la consola al iniciar la API

---

>`msgResponse.go` Estos mensages aparecen en response siempre

---

>`msgValidResponse.go` Estos mensages aparecen en los responce cuando se produce un error de validacion de los datos

Ejemplo de como crear uno:

>messages[**"TITULO"**] = **"Mensaje"**

## Base de datos de seguridad

>ROLES

---

{
    "nombre" : "Ejemplo",
    "descripcion" : "Rol de ejemplo",
    "sigdb" : [
        "EJ"
    ],
    "app" : [
        ¡DEBE RESPETARSE EL ORDEN PRIMERO LAS RUTAS DE APP Y LUEGO LA ACCIONES!
        {
            "ruta app" : " Descripcion para menu",
            "rutas" : {
                "sub ruta" : "Descripcion para item menu"
            },
            "acciones" : {
                "ruta de accion" : "Descripcion para boton"
            }
        }
    ],
    "createdAt" : ISODate("2021-07-09T17:40:51.897Z"),
    "updatedAt" : Date(-62135596800000),
    "usuarios" : [ARRAY NOMBRE DE USUARIOS]
}

---

>SUCURSALES

---

{
    "provincia" : "",
    "localidad" : "",
    "calle" : "",
    "numerop" : 123456,
    "telefonos" : [
        {
            "carac" : 2964,
            "num" : 123456
        }
    ],
    "celulares" : [
        {
            "carac" : 2964,
            "num" : 123456
        }
    ],
    "nombre" : "",
    "rs" : "",
    "cuit" : 123456,
    "iva" : 0.0,
    "siglas" : "SIGLAS QUE IDENTIFICAN A LA SUCURSA, SE USA COMO PARAMETRO EN EL REQUEST PARA SABER A QUE BD REFERECNIAR",
    "bdmongo" : "BASE DE DATOS DE MONGO A USAR POR LA SUCURSAL",
    "bdmysql" : "BASE DE DATOS DE MYSQL A USAR POR LA SUCURSAL",
    "folder" : "CARPETA DEL SERVIDOR DONDE SE GUARDARAN LOS ARCHIVOS DE LA SUCURSAL",
    "color" : "COLOR DISTINTIVO DE LA SUCURSAL",
    "horarios" : [
        {
            "descripcion" : "",
            "dias" : [
                "lunes",
                "martes",
                "miercoles",
                "jueves",
                "viernes"
            ],
            "apertura" : "10:00:00",
            "cierre" : "18:00:00"
        },
        {
            "descripcion" : "",
            "dias" : [
                "sabados"
            ],
            "apertura" : "10:00:00",
            "cierre" : "13:00:00"
        }
    ],
    "mail" : "",
    "pventa" : NUMERO DE PUNTO DE VENTAS PARA EL WEB SERVICE DE AFIP,
    "createdAt" : ISODate("2021-06-29T13:48:32.919Z"),
    "updatedAt" : Date(-62135596800000)
}

---

>USUARIOS

---

{
    "nombre" : "Super",
    "apellido" : "Admin",
    "usuario" : "admin",
    "pass" : "123456",
    "activo" : true, BAJA LOFGICA
    "reset" : false, ACTIVA RESETEO DE CONTRASEÑA
    "id" : 0, ID PARA LA BASE DE DATOS DE MYSQL
    "hash" : "HASH CREADO POR TOCKEN PARA VALIDAR LA SECION (ES AUTOMATICO)",
    "admin" : true HACE EL USUARIO ADMINISTRADOR TOTAL --- SE RECOIENDA SU USO SOLO PARA DESARROLLO
}

---

## Archivo .ENV

>Generales

- PRODUCCION=false
- DISABLENOTIFICACIONES=false
- PANIC=false
- DISABLE_TOKEN=false
- EXPIRE_TOKEN=259200
- CONTEXT_TIME=10
- CLAVE_SECRET=prueba
- PORT_API=3000
- IP_SERVER=localhost
- DISABLE_TOKEN=true
- DISABLE_CONSOLA=false
- TIMEOUT_CLIENT=3

>Documentacion

- DISABLEDOC=false
- URL_DOC=ASTRIC
- PORT_DOC=3100
- TOKEN=StringToken

>MongoDB

- MONGO_DB=mongodb://root:66865@25.41.64.244:27017
- MONGO_DB_SEGURIDAD=ASTRIC

>MySQL

- MYSQL=true
- MYSQL_USER=root
- MYSQL_PASS=66865
- MYSQL_HOST=25.41.64.244
- MYSQL_DB=daira
- MYSQL_PORT=3306
- MYSQL_LOGER=false
- MYSQL_SET_MAX_LIFE_TIMA = 3
- MYSQL_SET_MAX_OPEN_CONN=10
- MYSQL_SET_MAX_IDLE_CONNS=10

>DISCORD

- DISCORD=true
- APP_DISCORD=ASTRIC
- URL_DISCORD_HOKS_ERROR=URL hocks de discord

>Correos

- MAIL=true
- MAIL_FROM=no-reply@ASTRIC.net
- MAIL_SMTP=sd-1819389-l.dattaweb.com
- MAIL_PORT=587
- MAIL_ADDRES= web@ASTRIC.com.ar
- MAIL_PASS=******

>Mensages

- CODE_UNAUT=1819
- MSG_UNAUT_APP=El usuario no esta autorizado a ustilizar este modulo
- MSG_UNAUT_SECCION=El usuario no esta autorizado a ustilizar esta seccion
- MSG_UNAUT_ACCION=Su usuario no esta autorizado a realizar esta accion
- MSG_UNAUT_SUCURSAL=El usuario no esta autorizado a ustilizar esta sucursal


/*	DDB PARA CREACION DE TABLA

	CREATE TABLE usuarios
	(
	id         INT       NOT NULL AUTO_INCREMENT,
	created_at timestamp NULL    ,
	updated_at timestamp NULL    ,
	usuario    TEXT      NOT NULL,
	clave      TEXT      NOT NULL,
	activo     BOOL      NOT NULL DEFAULT false,
	reset      BOOL      NOT NULL DEFAULT false,
	hash       TEXT      NULL    ,
	admin      BOOL      NOT NULL DEFAULT false,
	id_roles   INT       NOT NULL,
	PRIMARY KEY (id)
	);

*/