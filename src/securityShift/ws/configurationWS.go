package ws

/**
 *
 * configurationWS
 *
 * Clase donde se encuentra los servicios rest
 *
 * @author Jeison Gaviria
 * @since 09/05/2016
 *
 *
 * Historia de Modificaciones
 * ------------------------------------------------------------- Autor	Fecha
 * Modificacion ----------------- -------------- ----------------------------
 */

import (
		"net/http"
 		"libs/hlib/hrest"
 		"securityShift/web/htmlServer"
 		"securityShift/bo/configuration"
 		)
 /**
     * estructura que emula la clase .
 **/

type ConfigurationWS struct{}

/**
     * Constructor de la clase .
**/

func NewConfigurationWS() *ConfigurationWS {
	s := new(ConfigurationWS)
	return s
}

/**
     * Método que representa el servicio que levanta los servicios https del sistema.
     *
     * @return flag   bool para levantar los servicios como hilo o no.
     .
**/

func (*ConfigurationWS) ActivateServices(flag bool) {
	hrest := hrest.NewRest()
	config := configuration.Config

	hrest.RegisterService("/services/test", "GET", htmlServer.ViewTestHttps)

	if flag {
		//http.ListenAndServe(config.PortServices, nil)
		http.ListenAndServeTLS(config.PortServer, config.Public_key, config.Private_key, nil)
	} else {
		//go http.ListenAndServe(config.PortServices, nil)
		go http.ListenAndServeTLS(config.PortServer, config.Public_key, config.Private_key, nil)
	}

}
