package js

/**
 *
 * configJS
 *
 * Clase js para el javascript de la viewConfig
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
	"fmt"
	"net/http"
)

/**
     * Método que representa el método js para limpiar un formulario.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func CleanFields(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<script type='text/javascript'> ")
	fmt.Fprintln(w, "cleanFields = function(div){ ")
	fmt.Fprintln(w, "$('#'+div).find(':input').each(function() { ")
	fmt.Fprintln(w, "var elemento= this; ")
	fmt.Fprintln(w, "$('#'+elemento.id).val(''); ")
	fmt.Fprintln(w, "}); ")
	fmt.Fprintln(w, "$('#'+div).find(':select').each(function() { ")
	fmt.Fprintln(w, "var elemento= this; ")
	fmt.Fprintln(w, "$('#'+elemento.id).val('0'); ")
	fmt.Fprintln(w, "}); ")
	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "</script> ")
}

/**
     * Método que representa el método js para validar un formulario.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func ValidationForm(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "<script type='text/javascript'> ")
		fmt.Fprintln(w, "validationForm=function(frm){ ")

		fmt.Fprintln(w, "var messageFile =''; ")
		fmt.Fprintln(w, "var band =true; ")
		fmt.Fprintln(w, "try{ ")
		fmt.Fprintln(w, "$('#'+frm).find(':input').each(function() { ")
		fmt.Fprintln(w, "var Myvali=false; ")
		fmt.Fprintln(w, "var elemento= this; ")
		fmt.Fprintln(w, "var myType = $('#'+elemento.id).attr('generictype'); ")
		fmt.Fprintln(w, "var myRequired = $('#'+elemento.id).attr('obligatory'); ")
		fmt.Fprintln(w, "var myName = $('#'+elemento.id).attr('name'); ")
		fmt.Fprintln(w, "var myId = $('#'+elemento.id).attr('id'); ")
		fmt.Fprintln(w, "var myValue = $('#'+elemento.id).val(); ")
		fmt.Fprintln(w, "var myspa = $('#'+elemento.id).attr('spa'); ")
		fmt.Fprintln(w, "if(myType!=null && myType!=undefined && myType!='' && myType!='undefined'){ ")
		fmt.Fprintln(w, "	if(myType=='text'){ ")
		fmt.Fprintln(w, "		if(myRequired=='true'){ ")
		fmt.Fprintln(w, "			if(myValue==''){ ")
		fmt.Fprintln(w, "				$(elemento).css({ 'border': '1px solid red' }); ")
		fmt.Fprintln(w, "				$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "				band=false; ")
		fmt.Fprintln(w, "			}else{ $(elemento).css({ 'border': '1px solid white' }); } ")
		fmt.Fprintln(w, "		} ")
		fmt.Fprintln(w, "	} ")
		
		fmt.Fprintln(w, "	if(myType=='number'){ ")
		fmt.Fprintln(w, "		if(myRequired=='true'){ ")
		fmt.Fprintln(w, "			if(myValue==''){ ")
		fmt.Fprintln(w, "				$(elemento).css({ 'border': '1px solid red' }); ")
		fmt.Fprintln(w, "				$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "				band=false; ")
		fmt.Fprintln(w, "			}else{ $(elemento).css({ 'border': '1px solid white' }); } ")
		fmt.Fprintln(w, "		} ")
		fmt.Fprintln(w, "		if(myValue!=''){ ")
		fmt.Fprintln(w, "			if(isNumeric(myValue)==false){ ")
		fmt.Fprintln(w, "				$(elemento).css({'border': '1px solid red' }); ")
		fmt.Fprintln(w, "				$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "				band=false; ")
		fmt.Fprintln(w, "			}else{ $(elemento).css({ 'border': '1px solid white' }); } ")
		fmt.Fprintln(w, "		} ")
		fmt.Fprintln(w, "	} ")

		fmt.Fprintln(w, "	if(myType=='port'){ ")
		fmt.Fprintln(w, "		if(myRequired=='true'){ ")
		fmt.Fprintln(w, "			if(myValue==''){ ")
		fmt.Fprintln(w, "				$(elemento).css({ 'border': '1px solid red' }); ")
		fmt.Fprintln(w, "				$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "				band=false; ")
		fmt.Fprintln(w, "			}else{ ")
		fmt.Fprintln(w, "				$(elemento).css({ 'border': '1px solid white' }); ")
		fmt.Fprintln(w, "			} ")
		fmt.Fprintln(w, "		} ")
		fmt.Fprintln(w, "		if(myValue!=''){ ")
		fmt.Fprintln(w, "			if(isNumeric(myValue)==false){ ")
		fmt.Fprintln(w, "				$(elemento).css({'border': '1px solid red' }); ")
		fmt.Fprintln(w, "				$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "				band=false; ")
		fmt.Fprintln(w, "			} else { ")
		fmt.Fprintln(w, "				if(parseInt(myValue) < 0){ ")
		fmt.Fprintln(w, "					$(elemento).css({'border': '1px solid red' }); ")
		fmt.Fprintln(w, "					$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "					band=false; ")
		fmt.Fprintln(w, "				} else { ")
		fmt.Fprintln(w, "					$(elemento).css({ 'border': '1px solid white' }); ") 
		fmt.Fprintln(w, " 				} ")
		fmt.Fprintln(w, " 			} ")
		fmt.Fprintln(w, "		} ")
		fmt.Fprintln(w, "	} ")

		fmt.Fprintln(w, "if(myType=='phone'){ ")
		fmt.Fprintln(w, "if(myRequired=='true'){ ")
		fmt.Fprintln(w, "if(myValue==''){ ")
		fmt.Fprintln(w, "$(elemento).css({'border': '1px solid red'}); ")
		fmt.Fprintln(w, "$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "band=false; ")
		fmt.Fprintln(w, "}else{ $(elemento).css({ 'border': '1px solid white' }); } ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "if(myValue!=''){ ")
		fmt.Fprintln(w, "if(isNumeric(myValue)==false){ ")
		fmt.Fprintln(w, "$(elemento).css({'border': '1px solid red'}); ")
		fmt.Fprintln(w, "$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "band=false; ")
		fmt.Fprintln(w, "}else{ $(elemento).css({ 'border': '1px solid white' }); } ")
		fmt.Fprintln(w, "if(myValue.length<10){ ")
		fmt.Fprintln(w, " $(elemento).css({'border': '1px solid red' }); ")
		fmt.Fprintln(w, "$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "band=false; ")
		fmt.Fprintln(w, "}else{ $(elemento).css({ 'border': '1px solid white' }); } ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "if(myType=='email'){ ")
		fmt.Fprintln(w, "if(myRequired=='true'){ ")
		fmt.Fprintln(w, "if(myValue==''){ ")
		fmt.Fprintln(w, "$(elemento).css({'border': '1px solid red' }); ")
		fmt.Fprintln(w, "$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "band=false; ")
		fmt.Fprintln(w, "}else{ $(elemento).css({ 'border': '1px solid white' }); } ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "if(myValue!=''){ ")
		fmt.Fprintln(w, " if(isEmail(myValue)==false){ ")
		fmt.Fprintln(w, "$(elemento).css({'border': '1px solid red' }); ")
		fmt.Fprintln(w, "$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "band=false; ")
		fmt.Fprintln(w, "}else{ $(elemento).css({ 'border': '1px solid white' }); } ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "if(myType=='password'){ ")
		fmt.Fprintln(w, "if(myRequired=='true'){ ")
		fmt.Fprintln(w, "if(myValue==''){ ")
		fmt.Fprintln(w, "$(elemento).css({'border': '1px solid red' }); ")
		fmt.Fprintln(w, "$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, "band=false; ")
		fmt.Fprintln(w, "}else{ $(elemento).css({ 'border': '1px solid white' }); } ")
		fmt.Fprintln(w, "}")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "if(myType=='select'){ ")
		fmt.Fprintln(w, "if(myRequired=='true'){ ")
		fmt.Fprintln(w, "if(myValue=='0'){ ")
		fmt.Fprintln(w, "$(elemento).css({'border': '1px solid red' }); ")
		fmt.Fprintln(w, "$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, " band=false; ")
		fmt.Fprintln(w, "}else{ $(elemento).css({ 'border': '1px solid white' }); } ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "} ")

		fmt.Fprintln(w, "if(myType=='file'){ ")
		fmt.Fprintln(w, "if(myRequired=='true'){ ")
		fmt.Fprintln(w, "var elementofile= $(this); ")
		fmt.Fprintln(w, "var elemt= elementofile.prop('files') ")
		fmt.Fprintln(w, "if (elemt.length==0 ){ ")
		fmt.Fprintln(w, "$(elemento).css({'border': '1px solid red' }); ")
		fmt.Fprintln(w, "$('#'+myspa).css({ 'display':'block' });")
		fmt.Fprintln(w, " band=false; ")
		fmt.Fprintln(w, "} else { ")
		fmt.Fprintln(w, "var name= elemt[0].name; ")
		fmt.Fprintln(w, " var ext = (name.substring(name.lastIndexOf('.'))).toLowerCase();")
		fmt.Fprintln(w, "if (ext=='.pem' || ext=='.PEM'){ ")
		fmt.Fprintln(w, "} else { ")
		fmt.Fprintln(w, "$(elemento).css({'border': '1px solid red' }); ")
		fmt.Fprintln(w, "$('#'+myspa).css({ 'display':'block' }); ")
		fmt.Fprintln(w, "band=false; ")
		fmt.Fprintln(w, "messageFile=' y los archivos deben ser extensión .pem'; ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "} ")

		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "} ")

		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "}); ")
		fmt.Fprintln(w, "}catch(err){ ")
		fmt.Fprintln(w, "console.log('error de validate form '+err); ")
		fmt.Fprintln(w, "alert('El formulario de id '+frm+' no existe') ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "if(band==false){ ")
		fmt.Fprintln(w, "$('#alert').html(''); ")
		fmt.Fprintln(w, "var html=''; ")

		fmt.Fprintln(w, "html+=\"<div id='dialogal' class='modal fade bs-example-modal-sm' tabindex='-1' role='dialog' aria-labelledby='mySmallModalLabel'>\" ")
		fmt.Fprintln(w, "html+=\"<div class='modal-dialog modal-sm'>\" ")
		fmt.Fprintln(w, "html+=\"<div class='modal-header colorHeader'>\" ")
		fmt.Fprintln(w, "html+=\"<button type='button' class='close' data-dismiss='modal' aria-label='Close'><span aria-hidden='true'>&times;</span></button>\" ")
		fmt.Fprintln(w, "html+=\"<h4 class='modal-title' id='gridSystemModalLabel'>Atención</h4>\" ")
		fmt.Fprintln(w, "html+=\"</div>\" ")
		fmt.Fprintln(w, "html+=\"<div class='modal-content'>\" ")
		fmt.Fprintln(w, "html+=\"<div class='alert alert-warning' role='alert'> \" ")
		fmt.Fprintln(w, "html+=\"<strong>Atención!</strong> Diligencie el formulario Correctamente, valide los campos marcados en rojo \"+messageFile ")
		fmt.Fprintln(w, "html+=\"</div>\" ")
		fmt.Fprintln(w, "html+=\"</div>\" ")
		fmt.Fprintln(w, "html+=\"</div>\" ")
		fmt.Fprintln(w, "html+=\"</div>\" ")

		//fmt.Fprintln(w, "alert('Diligencie el formulario Correctamente, valide los campos marcados en rojo'); ")
		fmt.Fprintln(w, "$('#alert').append(html); ")
		fmt.Fprintln(w, "$('#dialogal').modal('show'); ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "return band; ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "</script> ")

}

/**
     * Método que representa el método js para capturar los ficheros y enviarlos a los métodos de envió.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func SendConfiguration(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "<script type='text/javascript'> ")
		fmt.Fprintln(w, "sendConfiguration = function(){ ")
		fmt.Fprintln(w, "if (validationForm('container')) { ")
		fmt.Fprintln(w, "var ficheros = []; ")
		fmt.Fprintln(w, "var cont = 0; ")
		fmt.Fprintln(w, "var band = true; ")
		fmt.Fprintln(w, "$('#container').find('input:file').each(function() { ")
		fmt.Fprintln(w, "var elemento= $(this); ")

		fmt.Fprintln(w, "ficheros[cont]= elemento.prop('files')")
		fmt.Fprintln(w, "cont=cont+1; ")

		fmt.Fprintln(w, "}); ")

		fmt.Fprintln(w, "var filePriv = ficheros[0]; ")
		fmt.Fprintln(w, "var filePub = ficheros[1]; ")

		fmt.Fprintln(w, "sendFile(filePriv[0]); ")
		fmt.Fprintln(w, "sendFilePublic(filePub[0]); ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "</script> ")

}

/**
     * Método que representa el método js para enviar el archivo privado de https.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func SendFile(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "<script type='text/javascript'> ")
		fmt.Fprintln(w, "sendFile = function(formData) { ")
		fmt.Fprintln(w, "var contenido = this; ")
		fmt.Fprintln(w, "$.ajax({ ")
		fmt.Fprintln(w, "url: 'http://'+window.location.hostname+':8085/services/uploadfilePrivate', ")
		fmt.Fprintln(w, "data: formData, ")
		fmt.Fprintln(w, "processData: false, ")
		fmt.Fprintln(w, "type: 'POST', ")
		//fmt.Fprintln(w, "contentType: false, ")
		fmt.Fprintln(w, "mimeType:  'multipart/form-data', ")
		fmt.Fprintln(w, "success: function (data) { ")
		fmt.Fprintln(w, "var response = JSON.parse(data); ")
		fmt.Fprintln(w, "if(response.genericResponse.resultStatus == 'OK') { ")
		//fmt.Fprintln(w, "contenido.sendFilePublic(); ")
		//fmt.Fprintln(w, "contenido.saveConfig(); ")
		fmt.Fprintln(w, "} else {")
		fmt.Fprintln(w, "alert('no se pudo enviar el archivo'); ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "},")
		fmt.Fprintln(w, "error: function (err) ")
		fmt.Fprintln(w, "{")
		fmt.Fprintln(w, " console.log(err); ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "}); ")
		fmt.Fprintln(w, "} ")
		fmt.Fprintln(w, "</script> ")
}

/**
     * Método que representa el método js para enviar el archivo público de https.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func SendFilePublic(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<script type='text/javascript'> ")
	fmt.Fprintln(w, "sendFilePublic = function(formData) { ")
	fmt.Fprintln(w, "var contenido = this; ")
	fmt.Fprintln(w, "$.ajax({ ")
	fmt.Fprintln(w, "url: 'http://'+window.location.hostname+':8085/services/uploadfilePublic', ")
	fmt.Fprintln(w, "data: formData, ")
	fmt.Fprintln(w, "processData: false, ")
	fmt.Fprintln(w, "type: 'POST', ")
	fmt.Fprintln(w, "contentType: false, ")
	fmt.Fprintln(w, "mimeType:  'multipart/form-data', ")
	fmt.Fprintln(w, "success: function (data) { ")
	fmt.Fprintln(w, "var response = JSON.parse(data); ")
	fmt.Fprintln(w, "if(response.genericResponse.resultStatus == 'OK') { ")
	fmt.Fprintln(w, "contenido.saveConfig(); ")
	fmt.Fprintln(w, "} else {")
	fmt.Fprintln(w, "alert('no se pudo enviar el archivo'); ")
	fmt.Fprintln(w, "} ")
	fmt.Fprintln(w, "},")
	fmt.Fprintln(w, "error: function (err) ")
	fmt.Fprintln(w, "{")
	fmt.Fprintln(w, " console.log(err); ")
	fmt.Fprintln(w, "} ")
	fmt.Fprintln(w, "}); ")
	fmt.Fprintln(w, "} ")
	fmt.Fprintln(w, "</script> ")
}

/**
     * Método que representa el método js para enviar la configuración a security Shift.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func SaveConfig(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<script type='text/javascript'> ")
	fmt.Fprintln(w, "saveConfig = function() { ")
	fmt.Fprintln(w, "var portServer = ':'+$.trim($('#portServer').val()); ")
	fmt.Fprintln(w, "var data ={'PortServer':portServer,'Private_key':'config/private_key','Public_key':'config/public_key' }; ")

	fmt.Fprintln(w, "$.ajax({ ")
	fmt.Fprintln(w, "type: 'POST', ")
	fmt.Fprintln(w, "url: 'http://'+window.location.hostname+':8085/services/createConfig', ")
	fmt.Fprintln(w, "data: JSON.stringify(data), ")
	fmt.Fprintln(w, "}).then(function(data) { ")
	fmt.Fprintln(w, "var html = ''; ")
	fmt.Fprintln(w, "var dataS=JSON.parse(data); ")
	fmt.Fprintln(w, "$('#alert').html(''); ")

	fmt.Fprintln(w, "if (dataS.genericResponse.resultStatus=='OK') { ")

	fmt.Fprintln(w, "html+=\"<div id='dialogal' class='modal fade bs-example-modal-sm' tabindex='-1' role='dialog' aria-labelledby='mySmallModalLabel' data-keyboard='false' data-backdrop='static'>\" ")
	fmt.Fprintln(w, "html+=\"<div class='modal-dialog modal-sm'>\" ")
	fmt.Fprintln(w, "html+=\"<div class='modal-header colorHeaderSucces'>\" ")
	//fmt.Fprintln(w, "html+=\"<button id='btn_closeSucces' type='button' class='close' data-dismiss='modal' aria-label='Close'><span aria-hidden='true'>&times;</span></button>\" ")
	fmt.Fprintln(w, "html+=\"<h4 class='modal-title' id='gridSystemModalLabel'>Registro exitoso</h4>\" ")
	fmt.Fprintln(w, "html+=\"</div>\" ")
	fmt.Fprintln(w, "html+=\"<div class='modal-content'>\" ")
	fmt.Fprintln(w, "html+=\"<div class='alert alert-warning' role='alert'> \" ")
	fmt.Fprintln(w, "html+=\"<strong>Exito!</strong>  Configuración guardada \" ")
	fmt.Fprintln(w, "html+=\"</div>\" ")
	fmt.Fprintln(w, "html+=\"</div>\" ")
	fmt.Fprintln(w, "html+=\"</div>\" ")
	fmt.Fprintln(w, "html+=\"</div>\" ")

	
	fmt.Fprintln(w, "setTimeout(function(){ ")
	fmt.Fprintln(w, "port = dataS.genericResponse.object.replace('\"', ''); ")
	fmt.Fprintln(w, "port = port.replace('\"', ''); ")
	fmt.Fprintln(w, "window.location.href = 'https://'+window.location.hostname+port+'/services/test';")
	fmt.Fprintln(w, " }, 3000); ")

	fmt.Fprintln(w, "}else{ ")

	fmt.Fprintln(w, "html+=\"<div id='dialogal' class='modal fade bs-example-modal-sm' tabindex='-1' role='dialog' aria-labelledby='mySmallModalLabel'>\" ")
	fmt.Fprintln(w, "html+=\"<div class='modal-dialog modal-sm'>\" ")
	fmt.Fprintln(w, "html+=\"<div class='modal-header colorHeader'>\" ")
	fmt.Fprintln(w, "html+=\"<button type='button' class='close' data-dismiss='modal' aria-label='Close'><span aria-hidden='true'>&times;</span></button>\" ")
	fmt.Fprintln(w, "html+=\"<h4 class='modal-title' id='gridSystemModalLabel'>Atención</h4>\" ")
	fmt.Fprintln(w, "html+=\"</div>\" ")
	fmt.Fprintln(w, "html+=\"<div class='modal-content'>\" ")
	fmt.Fprintln(w, "html+=\"<div class='alert alert-warning' role='alert'> \" ")
	fmt.Fprintln(w, "html+=\"<strong>Error!</strong> \"+dataS.genericResponse.errorMessage ")
	fmt.Fprintln(w, "html+=\"</div>\" ")
	fmt.Fprintln(w, "html+=\"</div>\" ")
	fmt.Fprintln(w, "html+=\"</div>\" ")
	fmt.Fprintln(w, "html+=\"</div>\" ")

	fmt.Fprintln(w, "} ")

	fmt.Fprintln(w, "var htmlP=''; ")
	fmt.Fprintln(w, "htmlP+=\"	<div class='progress-bar progress-bar-success' role='progressbar' aria-valuenow='100' aria-valuemin='0' aria-valuemax='100' style='width:100%'>100% Completado \" ")
	fmt.Fprintln(w, "htmlP+=\"	</div>\" ")
	fmt.Fprintln(w, "$('.progress').html(htmlP); ")

	fmt.Fprintln(w, "$('.modal-backdrop.in').remove(); ")
	fmt.Fprintln(w, "$('#dialogal').modal('hide'); ")
	fmt.Fprintln(w, "$('#alert').html(html); ")
	fmt.Fprintln(w, "$('#dialogal').modal('show'); ")

	fmt.Fprintln(w, "}); ")
	fmt.Fprintln(w, "} ")
	//fmt.Fprintln(w, "} ")
	fmt.Fprintln(w, "</script> ")
}

/**
     * Método que representa el método js para validar si es email.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func IsEmail(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<script type='text/javascript'> ")
	fmt.Fprintln(w, "function isEmail(email) ")
	fmt.Fprintln(w, "{ ")
	fmt.Fprintln(w, "	var tester =/^([a-zA-Z0-9_.+-])+\\@(([a-zA-Z0-9-]+)\\.)+([a-zA-Z0-9]{2,4})+$/; ")
	fmt.Fprintln(w, "	return tester.test(email); ")
	fmt.Fprintln(w, "} ")
	fmt.Fprintln(w, "</script> ")
}

/**
     * Método que representa el método js para validar si es texto.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func IsText(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<script type='text/javascript'> ")
	fmt.Fprintln(w, "function isText(text) ")
	fmt.Fprintln(w, "{ ")
	fmt.Fprintln(w, "	var tester =/^[a-zA-Z0-9_.+-]$/; ")
	fmt.Fprintln(w, "	return tester.test(email); ")
	fmt.Fprintln(w, "} ")
	fmt.Fprintln(w, "</script> ")
}

/**
     * Método que representa el método js para validar si es numerico.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func IsNumeric(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<script type='text/javascript'> ")
	fmt.Fprintln(w, "function isNumeric(number) ")
	fmt.Fprintln(w, "{ ")
	fmt.Fprintln(w, "   var tester =/^(?:\\+|-)?\\d+$/; ")
	fmt.Fprintln(w, "	return tester.test(number); ")
	fmt.Fprintln(w, "} ")
	fmt.Fprintln(w, "</script> ")
}

/**
     * Método que representa el método js para volver a poner el color normal a los campos.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func NormalColor(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<script type='text/javascript'> ")
	fmt.Fprintln(w, "normalColor=function(element){ ")
	fmt.Fprintln(w, "	var myspa = $('#'+element.id).attr('spa'); ")
	fmt.Fprintln(w, "	$('#'+myspa).css({ 'display':'none' });")
	fmt.Fprintln(w, "	$(element).css({'border': '1px solid white'}); ")
	fmt.Fprintln(w, "} ")
	fmt.Fprintln(w, "</script> ")
}