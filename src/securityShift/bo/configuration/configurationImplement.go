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
	//"fmt"
	//"libs/github.com/gorilla/websocket"
	//"libs/hlib/email"
	"libs/hlib/util"
	"io/ioutil"
	//"regexp"
	//"strings"
	//"time"
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
	config, _ := json.Marshal(Config)

	configencrypt := util.Encrypt(string(config), keyEncript)

	buffer := bytes.NewBufferString(configencrypt)

	err := ioutil.WriteFile("config/securityShift.cfg", buffer.Bytes(), 0644)

	if err != nil {
		return "ocurrio un error ", false
	}

	return "", true
}