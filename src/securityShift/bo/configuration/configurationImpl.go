package configuration

/**
 *
 * security
 *
 * Clase Interface configuration security shift
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
	//"bytes"
	//"encoding/json"
	//"fmt"
	//"libs/hlib/connection"
	//"libs/hlib/util"
	//"io/ioutil"
	//"time"
	"securityShift/dao"
)

/**
	 * llave de encryptación
**/

const (
	keyEncript    string = "1234567890123456ASDFGHJKLZXCVBNM"
	valiEmail     string = `^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,3})$`
	valiCellPhone string = `^[0-9]{10}.+$`
)

/**
     * Interface conlos metodos abstractos de configuration security Shift.
**/

type ConfigImpl interface {
	CreateConfigGraphic([]byte) (string, bool)
}

var Config dao.Configuration

/**
     * Constructor de la clase .
**/

func NewConfigImpl() ConfigImpl {
	
	var s ConfigImpl

		s = NewConfigImplement()

	return s
}