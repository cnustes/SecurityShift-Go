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
		js.NewTypeTurn(w,r)
		js.RemoveDiv(w,r)

		css.CssViewConfig(w,r)

		fmt.Fprintln(w, "</head> ")
		fmt.Fprintln(w, "<body> ")

		fmt.Fprintln(w, "<div id='container' class='row'> ")

		//////* Configuración principal */////

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

		fmt.Fprintln(w, "	</div> ")
		fmt.Fprintln(w, "</div> ")

		//////* Configuración servidor email */////

		fmt.Fprintln(w, "<div class='row' id='intro'> ")
		fmt.Fprintln(w, "	<div class='panel-body opaque float-center'> ")
		fmt.Fprintln(w, "   	<div class='row div-labelTitle'> ")
		fmt.Fprintln(w, "       	<div class='col-xs-12 div-labelTitle'> ")
		fmt.Fprintln(w, "           	<label>Configuración servidor Email</label> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "       </div> ")

		fmt.Fprintln(w, "       <div id='divportServer' class='row marginInput'> ")
		fmt.Fprintln(w, "       	<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "           	<label>Correo servidor:</label> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "           <div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           	<div class='input-group'> ")
		fmt.Fprintln(w, "               	<input type='email' id='emailServer' name='emailServer' class='form-control' generictype='email' obligatory='true' spa='spaemailServer' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   <span class='input-group-addon opaque' id='basic-addon2'><span id='spaemailServer' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               </div> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "       </div> ")

		fmt.Fprintln(w, "		<div class='row marginInput'> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "				<label>Password correo servidor:</label> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           	<div class='input-group'> ")
		fmt.Fprintln(w, "               	<input type='password' id='emailServerPass' name='emailServerPass' class='form-control' generictype='text' obligatory='true' spa='spaemailServerPass' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   <span class='input-group-addon opaque' id='basic-addon2'><span id='spaemailServerPass' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               </div> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "		</div> ")

		fmt.Fprintln(w, "		<div class='row marginInput'> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "				<label>Host mail servidor:</label> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           	<div class='input-group'> ")
		fmt.Fprintln(w, "               	<input type='text' id='emailServerHost' name='emailServerHost' class='form-control' generictype='text' obligatory='true' spa='spaemailServerHost' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   <span class='input-group-addon opaque' id='basic-addon2'><span id='spaemailServerHost' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               </div> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "		</div> ")

		fmt.Fprintln(w, "		<div class='row marginInput'> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "				<label>Puerto correo servidor:</label> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           	<div class='input-group'> ")
		fmt.Fprintln(w, "               	<input type='number' id='emailServerPort' name='emailServerPort' class='form-control' generictype='number' obligatory='true' spa='spaemailServerPort' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   <span class='input-group-addon opaque' id='basic-addon2'><span id='spaemailServerPort' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               </div> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "		</div> ")

		fmt.Fprintln(w, "	</div> ")
		fmt.Fprintln(w, "</div> ")

		//////* Configuración servidor sms */////

		fmt.Fprintln(w, "<div class='row' id='intro'> ")
		fmt.Fprintln(w, "	<div class='panel-body opaque float-center'> ")
		fmt.Fprintln(w, "   	<div class='row div-labelTitle'> ")
		fmt.Fprintln(w, "       	<div class='col-xs-12 div-labelTitle'> ")
		fmt.Fprintln(w, "           	<label>Configuración servidor SMS</label> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "       </div> ")

		fmt.Fprintln(w, "       <div id='divportServer' class='row marginInput'> ")
		fmt.Fprintln(w, "       	<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "           	<label>Llave del servidor sms:</label> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "           <div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           	<div class='input-group'> ")
		fmt.Fprintln(w, "               	<input type='text' id='keyServerSms' name='keyServerSms' class='form-control' generictype='text' obligatory='true' spa='spakeyServerSms' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   <span class='input-group-addon opaque' id='basic-addon2'><span id='spakeyServerSms' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               </div> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "       </div> ")

		fmt.Fprintln(w, "		<div class='row marginInput'> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "				<label>Código del servidor sms:</label> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           	<div class='input-group'> ")
		fmt.Fprintln(w, "               	<input type='text' id='codeServerSms' name='codeServerSms' class='form-control' generictype='text' obligatory='true' spa='spacodeServerSms' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   <span class='input-group-addon opaque' id='basic-addon2'><span id='spacodeServerSms' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               </div> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "		</div> ")

		fmt.Fprintln(w, "		<div class='row marginInput'> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "				<label>Número de envios servidor de sms:</label> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           	<div class='input-group'> ")
		fmt.Fprintln(w, "               	<input type='number' id='numServerSms' name='numServerSms' class='form-control' generictype='number' obligatory='true' spa='spanumServerSms' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   <span class='input-group-addon opaque' id='basic-addon2'><span id='spanumServerSms' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               </div> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "		</div> ")

		fmt.Fprintln(w, "	</div> ")
		fmt.Fprintln(w, "</div> ")

		//////* Configuración sesiones */////

		fmt.Fprintln(w, "<div class='row' id='intro'> ")
		fmt.Fprintln(w, "	<div class='panel-body opaque float-center'> ")
		fmt.Fprintln(w, "   	<div class='row div-labelTitle'> ")
		fmt.Fprintln(w, "       	<div class='col-xs-12 div-labelTitle'> ")
		fmt.Fprintln(w, "           	<label>Configuración de sesión</label> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "       </div> ")

		fmt.Fprintln(w, "       <div id='divportServer' class='row marginInput'> ")
		fmt.Fprintln(w, "       	<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "           	<label>Intentos de logueo:</label> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "           <div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           	<div class='input-group'> ")
		fmt.Fprintln(w, "               	<input type='number' id='numFile' name='numFile' class='form-control' generictype='number' obligatory='true' spa='spanumFile' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   <span class='input-group-addon opaque' id='basic-addon2'><span id='spanumFile' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               </div> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "       </div> ")

		fmt.Fprintln(w, "		<div class='row marginInput'> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "				<label>Tiempo de caducidad sesión:</label> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           	<div class='input-group'> ")
		fmt.Fprintln(w, "               	<input type='number' id='timeSesion' name='timeSesion' class='form-control' generictype='number' obligatory='true' spa='spatimeSesion' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   <span class='input-group-addon opaque' id='basic-addon2'><span id='spatimeSesion' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               </div> ")
		fmt.Fprintln(w, "			</div> ")
		fmt.Fprintln(w, "		</div> ")

		fmt.Fprintln(w, "	</div> ")
		fmt.Fprintln(w, "</div> ")

		///////////////////////////////////////////////

		//////* Configuración turnos */////

		fmt.Fprintln(w, "<div class='row' id='intro'> ")
		fmt.Fprintln(w, "	<div class='panel-body opaque float-center'> ")
		fmt.Fprintln(w, "   	<div class='row div-labelTitle'> ")
		fmt.Fprintln(w, "       	<div class='col-xs-12 div-labelTitle'> ")
		fmt.Fprintln(w, "           	<label>Configuración de turnos</label> ")
		fmt.Fprintln(w, "           </div> ")
		fmt.Fprintln(w, "       </div> ")

		fmt.Fprintln(w, "       <div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "          <a type='button' class='btn btn-default' aria-haspopup='true' aria-expanded='false' onclick='newTypeTurn();'>+</a>")
		fmt.Fprintln(w, "       </div> ")

		fmt.Fprintln(w, "       <div id='divShift' class='row marginInput'> ")

		fmt.Fprintln(w, "       <div id='div-turn-1' class='col-xs-12' > ")

		fmt.Fprintln(w, "       	<div class='col-xs-3 fontLabel'> ")
		fmt.Fprintln(w, "       		<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "           		<label>Typo de turno:</label> ")
		fmt.Fprintln(w, "           	</div> ")
		fmt.Fprintln(w, "           	<div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           		<div class='input-group'> ")
		fmt.Fprintln(w, "               		<input type='text' id='typeTurn-1' name='typeTurn-1' class='form-control' generictype='text' obligatory='true' spa='spatypeTurn' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   	<span class='input-group-addon opaque' id='basic-addon2'><span id='spatypeTurn' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               	</div> ")
		fmt.Fprintln(w, "           	</div> ")
		fmt.Fprintln(w, "       	</div> ")

		fmt.Fprintln(w, "       	<div class='col-xs-3 fontLabel'> ")
		fmt.Fprintln(w, "       		<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "           		<label>Rango inicial:</label> ")
		fmt.Fprintln(w, "           	</div> ")
		fmt.Fprintln(w, "           	<div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           		<div class='input-group'> ")
		fmt.Fprintln(w, "               		<input type='number' id='initial-1' name='initial-1' class='form-control' generictype='number' obligatory='true' spa='spainitial' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   	<span class='input-group-addon opaque' id='basic-addon2'><span id='spainitial' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               	</div> ")
		fmt.Fprintln(w, "           	</div> ")
		fmt.Fprintln(w, "       	</div> ")

		fmt.Fprintln(w, "       	<div class='col-xs-3 fontLabel'> ")
		fmt.Fprintln(w, "       		<div class='col-xs-12 fontLabel'> ")
		fmt.Fprintln(w, "           		<label>Rango final:</label> ")
		fmt.Fprintln(w, "           	</div> ")
		fmt.Fprintln(w, "           	<div class='col-xs-12 fieldRight'> ")
		fmt.Fprintln(w, "           		<div class='input-group'> ")
		fmt.Fprintln(w, "               		<input type='number' id='final-1' name='final-1' class='form-control' generictype='number' obligatory='true' spa='spatypeTurn' required onfocus='normalColor(this)' aria-describedby='basic-addon2'> ")
		fmt.Fprintln(w, "                   	<span class='input-group-addon opaque' id='basic-addon2'><span id='spatypeTurn' class='glyphicon glyphicon-asterisk' style='display: none;'></span> ")
		fmt.Fprintln(w, "               	</div> ")
		fmt.Fprintln(w, "           	</div> ")
		fmt.Fprintln(w, "       	</div> ")

		fmt.Fprintln(w, "       	<div class='col-xs-3 fontLabel'> ")
		fmt.Fprintln(w, "       		<div class='col-xs-12 fontLabel'> ")
		//fmt.Fprintln(w, "           		<a type='button' class='btn btn-danger size-close' aria-haspopup='true' aria-expanded='false' onclick='removeDiv(1);'>x</a>")
		fmt.Fprintln(w, "           	</div> ")
		fmt.Fprintln(w, "       	</div> ")

		fmt.Fprintln(w, "       </div> ")
		
		fmt.Fprintln(w, "       </div> ")

		fmt.Fprintln(w, "		<div class='row marginInput'> ")
		fmt.Fprintln(w, "			<div class='col-xs-12 botom-size'> ")
		fmt.Fprintln(w, "				<a type='button' class='btn btn-primary bottom-nextToback' aria-haspopup='true' aria-expanded='false' onclick='sendConfiguration();'>GUARDAR</a>")
		fmt.Fprintln(w, "   		</div> ")
		fmt.Fprintln(w, "		</div> ")

		fmt.Fprintln(w, "	</div> ")
		fmt.Fprintln(w, "</div> ")

		///////////////////////////////////////////////

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