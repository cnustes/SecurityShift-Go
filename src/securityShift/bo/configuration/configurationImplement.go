package configuration

/**
 *
 * configurationImplement
 *
 * Clase donde se encuentra la implementación de los métodos de la interfaz configurationImpl
 *
 * @author Jeison Gaviria
 * @since 10/05/2016
 *
 *
 * Historia de Modificaciones
 * ------------------------------------------------------------- Autor	Fecha
 * Modificacion ----------------- -------------- ----------------------------
 */

import (
	"bytes"
	//"crypto/md5"
	//"encoding/hex"
	"encoding/json"
	"fmt"
	//"libs/github.com/gorilla/websocket"
	"libs/hlib/email"
	"libs/hlib/sms"
	"libs/hlib/util"
	"io/ioutil"
	"securityShift/dao"
	//"regexp"
	"strings"
	"time"
)

/**
     * estructura que emula la clase .
**/

type ConfigImplement struct {
}

/**
     * Constructor de la clase .
**/

func NewConfigImplement() *ConfigImplement {
	s := new(ConfigImplement)

	return s
}

/**
     * Método que crea la configuración como un archivo encriptado.
     *
     * @params  []byte parametros de la configuración en byte.
     * @retur   string de respuesta y bool de estado de respuesta.
**/

func  (s *ConfigImplement) CreateConfigGraphic(params []byte) (string, bool) {
	
	json.Unmarshal(params, &Config)
	fmt.Println("configuración >>> ",Config)
	vali, flag := ValidateConfiguration(&Config)

	if flag {

		Config.Subject = "Notificacion turno"
		Config.Message = util.Concatenate([]string{"<table width='100%' border='0' cellspacing='0' cellpadding='0'><tr><td align='center' valign='top' bgcolor='#f1f1f1;' style='background-color:#f1f1f1;'><br><br><table width='600' border='0' cellspacing='0' cellpadding='0'><tr><td align='left' valign='top'><img src='https://www.transoftsolutions.com/content/main/banner-password.jpg' width='600' height='191' style='display:block;'></td></tr><tr><td align='center' valign='top' bgcolor='#ffffff' style='background-color:#ffffff; font-family:Arial, Helvetica, sans-serif; font-size:13px; color:#000000; padding:0px 5px 0px 15px;'><table width='100%' border='0' cellspacing='0' cellpadding='0'><tr><td align='left' valign='top' style='font-size:13px; font-family:Arial, Helvetica, sans-serif; color:#000000;'><div style='font-family:Georgia, 'Times New Roman', Times, serif; font-size:72px; color:#7bab00;'><i>Notificacion de turno</i></div><div style='font-size:24px; color:#7bab00;'><br>{{.Date}}</div><div style='font-size:16px;'><br>Bienvenido a la plataforma</div><div><br>este es un correo personal, {{.Mess}}: <br>{{.Turn}}<br><br><br><br><br><br>Security shift<br></div></td><td width='210' align='center' valign='top'><!--<img src='images/mainpic.png' width='259' height='416'>--></td></tr></table></td></tr><tr><td align='left' valign='top'><img src='http://www.e-astorion.com/img/slides/slide-bg1.jpg' width='600' height='127' style='display:block;'></td></tr></table><br><br></td></tr></table>"})

		config, _ := json.Marshal(Config)
		configencrypt := util.Encrypt(string(config), keyEncript)

		buffer := bytes.NewBufferString(configencrypt)

		err := ioutil.WriteFile("config/securityShift.cfg", buffer.Bytes(), 0644)

		if err != nil {
			return "ocurrio un error ", false
		}
	}else{
		return vali, flag
	}

	return "", true
}

func ValidateConfiguration(configura *dao.Configuration) (string, bool) {
	var vali bool = true

	now := time.Now()

	messageSMS := util.Concatenate([]string{now.String(), " test servicio de SMS funciono"})
	fmt.Println("message >>> ",messageSMS)

	messageSMS, errSMS := sms.SendSMS(configura.KeySMS, configura.SecretSMS, configura.FromSMS, messageSMS, configura.FromSMS)
	
	if errSMS == false {
		return messageSMS, false
	}

	messageDate := strings.Replace("<table width='100%' border='0' cellspacing='0' cellpadding='0'><tr><td align='center' valign='top' bgcolor='#f1f1f1;' style='background-color:#f1f1f1;'><br><br><table width='600' border='0' cellspacing='0' cellpadding='0'><tr><td align='left' valign='top'><img src='https://www.transoftsolutions.com/content/main/banner-password.jpg' width='600' height='191' style='display:block;'></td></tr><tr><td align='center' valign='top' bgcolor='#ffffff' style='background-color:#ffffff; font-family:Arial, Helvetica, sans-serif; font-size:13px; color:#000000; padding:0px 5px 0px 15px;'><table width='100%' border='0' cellspacing='0' cellpadding='0'><tr><td align='left' valign='top' style='font-size:13px; font-family:Arial, Helvetica, sans-serif; color:#000000;'><div style='font-family:Georgia, 'Times New Roman', Times, serif; font-size:72px; color:#7bab00;'><i>prueba datos de servior de correo</i></div><div style='font-size:24px; color:#7bab00;'><br>{{.Date}}</div><div style='font-size:16px;'><br>Bienvenido a la plataforma</div><div><br>este correo nos permite validar el servidor de correo.<br><br><br>test servior correo<br><br><br><br>Security shift<br></div></td><td width='210' align='center' valign='top'><!--<img src='images/mainpic.png' width='259' height='416'>--></td></tr></table></td></tr><tr><td align='left' valign='top'><img src='http://www.e-astorion.com/img/slides/slide-bg1.jpg' width='600' height='127' style='display:block;'></td></tr></table><br><br></td></tr></table>", "{{.Date}}", now.String(), 1)
	fmt.Println("configuration >>> ",configura)
	errSend := email.SendEmail2([]string{configura.EmailServer}, configura.HostMail, configura.EmailServer, configura.PasswordEmail, "test servidor correo", messageDate, configura.PortMail)

	if errSend != nil {
		return "verifique los datos de servidor de correo", false
	}

	return "OK", vali
}