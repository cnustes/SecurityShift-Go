package html

/**
 *
 * html
 *
 * Clase donde esta la gui de configuración y los metodos de envio de archivos y creación de configuración. 
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
	"securityShift/web/bootstrap"
	//"bufio"
	"bytes"
	//"encoding/base64"
	"encoding/json"
	"fmt"
	"libs/hlib/response"
	"io/ioutil"
	"securityShift/web/jQuery"
	"securityShift/web/css"
	"securityShift/web/js"
	"securityShift/bo/configuration"
	"securityShift/ws"
	"net/http"
	"os"
)

/**
     * estructura que emula la clase .
**/

type Securityconfig struct {
}

/**
     * Constructor de la clase .
**/

func NewSecurityConfig() *Securityconfig {

	s := new(Securityconfig)
	return s
}


var config = configuration.Config

var bsmincss = bytes.NewBuffer(bootstrap.Bootstrapmincss)
var bsminjs = bytes.NewBuffer(bootstrap.Bootstrapminjs)
var bsthememincss = bytes.NewBuffer(bootstrap.Bootstrapthememincss)

var jq1113minjs = bytes.NewBuffer(jQuery.JQuery1113minjs)
var jqmigrate121minjs = bytes.NewBuffer(jQuery.JQuerymigrate121minjs)

/**
     * Método que representa el html de la pagina de configuración del sistema.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func (*Securityconfig) ViewConfig(w http.ResponseWriter, r *http.Request) {

	_, err := ioutil.ReadFile("config/securityShift.cfg")

	if err != nil {

		fmt.Fprintln(w, "<html> ")
		fmt.Fprintln(w, "<head> ")

		fmt.Fprintln(w, "<script type='text/javascript'> ")
		fmt.Fprintln(w, jq1113minjs.String())
		fmt.Fprintln(w, "</script> ")

		fmt.Fprintln(w, "<script type='text/javascript'> ")
		fmt.Fprintln(w, jqmigrate121minjs.String())
		fmt.Fprintln(w, "</script> ")

		fmt.Fprintln(w, "<style type='text/css'> ")
		fmt.Fprintln(w, bsmincss.String())
		fmt.Fprintln(w, "</style> ")

		fmt.Fprintln(w, "<style type='text/css'> ")
		fmt.Fprintln(w, bsthememincss.String())
		fmt.Fprintln(w, "</style> ")

		fmt.Fprintln(w, "<script type='text/javascript'> ")
		fmt.Fprintln(w, bsminjs.String())
		fmt.Fprintln(w, "</script> ")

		js.CleanFields(w,r)
		js.ValidationForm(w,r)
		js.IsEmail(w,r)
		js.IsText(w,r)
		js.IsNumeric(w,r)
		js.NormalColor(w,r)
		js.SendConfiguration(w,r)
		js.SendFile(w,r)
		js.SendFilePublic(w,r)
		js.SaveConfig(w,r)

		css.CssViewConfig(w,r)

		fmt.Fprintln(w, "</head> ")
		fmt.Fprintln(w, "<body> ")

		fmt.Fprintln(w, "<div id='container' class='row'> ")

		fmt.Fprintln(w, "<div class='row' id='intro'> ")
		fmt.Fprintln(w, "	<div class='panel-body opaque float-center'> ")
		fmt.Fprintln(w, "   	<div class='row div-labelTitle'> ")
		fmt.Fprintln(w, "       	<div class='col-xs-12 div-labelTitle'> ")
		fmt.Fprintln(w, "           	<label>Configuración principal</label> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "       </div> ")

		fmt.Fprintln(w, "       <div id='divportServer' class='row marginInput'> ")
		fmt.Fprintln(w, "       	<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "           	<label>Puerto de servicios:</label> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "           <div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           	<div class='input-group'> ")
		fmt.Fprintln(w, "               	<input type='number' id='portServer' name='portServer' class='form-control' generictype='port' obligatory='true' spa='spaportServer' required onfocus='normalColor(this)' placeholder='ejemplo= 8080' max='4' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   <span class='input-group-addon opaque' id='basic-addon2'><span id='spaportServer' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               </div> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "       </div> ")

		fmt.Fprintln(w, "		<div class='row marginInput'> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "				<label>Llave privada:</label> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "				<div class='input-group'> ")
		fmt.Fprintln(w, "					<input type='file' id='filePriv' name='filePriv' class='form-control' generictype='file' obligatory='true' spa='spafilePriv' required onfocus='normalColor(this)' placeholder='' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "					<span class='input-group-addon' id='basic-addon2'><span id='spafilePriv' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "				</div> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "		</div> ")

		fmt.Fprintln(w, "		<div class='row marginInput'> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "				<label>Llave pública:</label> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "				<div class='input-group'> ")
		fmt.Fprintln(w, "					<input type='file' id='filePublic' name='archivo[]' class='form-control' generictype='file' obligatory='true' spa='spafilePublic' required onfocus='normalColor(this)' placeholder='' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "					<span class='input-group-addon' id='basic-addon2'><span id='spafilePublic' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "				</div> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "		</div> ")

		fmt.Fprintln(w, "		<div class='row marginInput'> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 botom-size'> ")
		fmt.Fprintln(w, "				<a type='button' class='btn btn-primary bottom-nextToback' aria-haspopup='true' aria-expanded='false' onclick='sendConfiguration();'>GUARDAR</a>")
		fmt.Fprintln(w, "   		</div> ")
		fmt.Fprintln(w, "		</div> ")

		fmt.Fprintln(w, "	</div> ")
		fmt.Fprintln(w, "</div> ")

		fmt.Fprintln(w, "</div> ")

		fmt.Fprintln(w, "<div id='alert' class='row'></div> ")

		fmt.Fprintln(w, "</body> ")
		fmt.Fprintln(w, "</html> ")

	} else {

		fmt.Fprintln(w, "<script type='text/javascript'> ")
		fmt.Fprintln(w, "window.onload = function() { ")
		/****/ fmt.Fprintln(w, "window.location.href = 'https://'+window.location.hostname+'"+config.PortServer+"/services/test';")
		fmt.Fprintln(w, "}; ")
		fmt.Fprintln(w, "</script> ")

	}
}

/**
     * Método que representa el servicio que permite crear la configuración del sistema.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func (*Securityconfig) CreateConfig(w http.ResponseWriter, r *http.Request) {
	response := response.NewGenericResponse()
	bodyresp, _ := ioutil.ReadAll(r.Body)
	flag := false
	resp := ""
	shift := configuration.NewConfigImpl()

	resp, flag = shift.CreateConfigGraphic(bodyresp)

	if flag {
		config = configuration.Config
		resp = config.PortServer
		
		objectResponse, _ := json.Marshal(resp)

		responseMap := response.SetResponseOK("OK", string(objectResponse))

		body, _ := json.Marshal(responseMap)

		serv := ws.NewConfigurationWS()
		serv.ActivateServices(false)

		fmt.Fprintln(w, string(body))

	} else {
			
		err := os.Remove("config/securityShift.cfg")
		if err != nil {
			fmt.Println("ocurrió un error borrando la carpeta-->", err)
		}

		responseMap := response.SetResponseFailure("FAILURE", resp, "220 ")

		body, _ := json.Marshal(responseMap)

		fmt.Fprintln(w, string(body))
	}
}

/**
     * Método que representa el servicio que carga archivos al sistema.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func (*Securityconfig) UploadFilePrivate(w http.ResponseWriter, r *http.Request) {
	response := response.NewGenericResponse()
	bodyresp, _ := ioutil.ReadAll(r.Body)
	flag := false
	resp := ""
	var errpackage error

	_, errOpen := os.Open("config")

	if errOpen != nil {
		errpackage = os.Mkdir("config", 0770)
	} else {
		errpackage = nil
	}

	if errpackage != nil {
		flag = false
		resp = "error creando carpeta de configuración"
	} else {

		err := ioutil.WriteFile("config/private_key", bodyresp, 0644)

		if err != nil {
			flag = false
			resp = "error guardando el archivo"
		} else {
			flag = true
			resp = "archivo creado exitosamente"
		}
	}

	if flag {

		objectResponse, _ := json.Marshal(resp)

		responseMap := response.SetResponseOK("OK", string(objectResponse))

		body, _ := json.Marshal(responseMap)

		fmt.Fprintln(w, string(body))

	} else {
		responseMap := response.SetResponseFailure("FAILURE", resp, "220 ")

		body, _ := json.Marshal(responseMap)

		fmt.Fprintln(w, string(body))
	}
}

/**
     * Método que representa el servicio que carga archivos al sistema.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func (*Securityconfig) UploadFilePublic(w http.ResponseWriter, r *http.Request) {
	response := response.NewGenericResponse()
	bodyresp, _ := ioutil.ReadAll(r.Body)
	flag := false
	resp := ""

	err := ioutil.WriteFile("config/public_key", bodyresp, 0644)

	if err != nil {
		flag = false
		resp = "error guardando el archivo"
	} else {
		flag = true
		resp = "archivo creado exitosamente"
	}

	if flag {

		objectResponse, _ := json.Marshal(resp)

		responseMap := response.SetResponseOK("OK", string(objectResponse))

		body, _ := json.Marshal(responseMap)

		fmt.Fprintln(w, string(body))

	} else {
		responseMap := response.SetResponseFailure("FAILURE", resp, "220 ")

		body, _ := json.Marshal(responseMap)

		fmt.Fprintln(w, string(body))
	}
}