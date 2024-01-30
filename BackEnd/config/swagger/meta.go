// Package swagger ASTRIC - API Rest FULL.
//
//	    Token de prueba:
//	    {{.Token}}
//			 ------------------------------------------------------------------------------
//			Usuario disponible para dev = admin
//			 ------------------------------------------------------------------------------
//			En el lugar de la sucursal debe colocarse el alias
//			Alias de sucursales de prueba:
//				Para Rio Grande = GRK
//				Para Ushuaia = USH
//			BASE DE DATOS CON ESTRUCTURA COMPLETA USH
//
//	    Schemes: http
//	    Host: {{.Host}}
//	    BasePath: /
//	    Version: 3.0.0
//	    License: Free GNU --- Propiedad de DARA It Groups
//	    Contact: Soporte DAIRA<soporte@dairaitgroups.com.ar>  https://www.dairaitgroup.com.ar/
//
//	    Consumes:
//	    - application/json
//
//	    Produces:
//	    - application/json
//
//	    Security:
//	     - ApiKeyAuth: []
//	     - Sucursal: []
//	    SecurityDefinitions:
//	       ApiKeyAuth:
//	         type: apiKey
//	         name: Authorization
//	         in: header
//	       Sucursal:
//	         type: apiKey
//	         name: Sucursal
//	         in: header
//
// swagger:meta
//
//nolint:ireturn
package swagger
