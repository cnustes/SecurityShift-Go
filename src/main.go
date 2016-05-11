package main

/**
 *
 * main
 *
 * Clase main donde se levanta el servidor de configuración puerto 8085.
 *
 * @author Jeison Gaviria
 * @since 11/05/2016
 *
 *
 * Historia de Modificaciones
 * ------------------------------------------------------------- Autor	Fecha
 * Modificacion ----------------- -------------- ----------------------------
 */

import (
	"bytes"
	"encoding/json"
	"fmt"
	"libs/github.com/open-golang/open"
	"securityShift/web/html"
	"securityShift/ws"
	"securityShift/bo/configuration"
	"libs/hlib/hrest"
	"libs/hlib/util"
	"io/ioutil"
	"net/http"
	//"runtime"
)

var hostbase = "localhost"//security.Hostbase

const (
	keyEncript string = "1234567890123456ASDFGHJKLZXCVBNM"
)



/**
	Método que activa la configuración si existe y si no existe la crea.
**/

func ActivateConfig() bool {

	fileComfig, err := ioutil.ReadFile("config/securityShift.cfg")

	if err != nil {
		
		return false
	}

	fileConfigString := util.Decrypt(string(fileComfig), keyEncript)

	fileConfigBytes := bytes.NewBufferString(fileConfigString)

	err = json.Unmarshal(fileConfigBytes.Bytes(), &configuration.Config)

	if err != nil {
		return false
	}

	fmt.Println("active...")
	return true
}

/**
     * Método que levanta los servicios de configuración.
     *
     * @flag   bool para levantar los servicios como hilo o no.
     .
**/

func activateGui(state bool) {
	varHtml := html.NewSecurityConfig()
	hrest := hrest.NewRest()

	hrest.RegisterService("/viewConfig", "GET", varHtml.ViewConfig)
	hrest.RegisterService("/", "GET", varHtml.ViewConfig)
	hrest.RegisterService("/services/createConfig", "POST", varHtml.CreateConfig)
	hrest.RegisterService("/services/uploadfilePrivate", "POST", varHtml.UploadFilePrivate)
	hrest.RegisterService("/services/uploadfilePublic", "POST", varHtml.UploadFilePublic)
	

	fmt.Println("Abra el navegador e ingrese a la pagina https://" + hostbase + ":8085/viewConfig")
	err := open.Run("http://" + hostbase + ":8085/viewConfig")

	if err != nil {
		fmt.Println("error browser-->", err)
	}

	if state {
		http.ListenAndServe(":8085", nil)
	} else {
		go http.ListenAndServe(":8085", nil)
	}
}

func main() {
	
	//runtime.LockOSThread()

	//go func() {
		
		activation := ActivateConfig()

		if activation {
			activateGui(false)

			serv := ws.NewConfigurationWS()
			serv.ActivateServices(true)

		} else {
			activateGui(true)
		}

	//}()

}
