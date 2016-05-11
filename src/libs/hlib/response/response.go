package response

/**
 *
 * response
 *
 * Clase que representa las respuestas al front de la forma GenericResponse
 *
 * @author Jeison Gaviria
 * @since 12/08/2015
 *
 * Derechos Reservados de Autor. (C) IP Total Software S.A
 *
 * Historia de Modificaciones
 * ------------------------------------------------------------- Autor	Fecha
 * Modificacion ----------------- -------------- ----------------------------
 */

/**
     * estructira que emula la clase.
**/
type GenericResponse struct {
	genericResponseMap map[string]map[string]string
}

var responseMap map[string]string

/**
     * Constructor de la clase .
**/
func NewGenericResponse() *GenericResponse {
	response := new(GenericResponse)

	return response
}

/**
     * Retorna un GenericResponse Fallido.
     *
     * @resultStatus estado de la respuesta en este caso "FAILURE".
     * @errorMessage mensaje del error(string).
     * @errorCode    codigo del error(string).
     * @return       GenericResponse con respuesta fallida
**/

func (response *GenericResponse) SetResponseFailure(resultStatus string, errorMessage string, errorCode string) map[string]map[string]string {
	responseMap = make(map[string]string)
	responseMap["resultStatus"] = resultStatus
	responseMap["errorMessage"] = errorMessage
	responseMap["errorCode"] = errorCode

	response.genericResponseMap = make(map[string]map[string]string)
	response.genericResponseMap["genericResponse"] = responseMap

	return response.genericResponseMap
}

/**
     * Retorna un GenericResponse satisfactorio.
     *
     * @resultStatus estado de la respuesta en este caso "OK".
     * @object       objeto que se va a retornar.
     * @return       GenericResponse con respuesta satisfactoria.
**/

func (response *GenericResponse) SetResponseOK(resultStatus string, object string) map[string]map[string]string {

	responseMap = make(map[string]string)
	responseMap["resultStatus"] = resultStatus
	responseMap["object"] = object

	response.genericResponseMap = make(map[string]map[string]string)
	response.genericResponseMap["genericResponse"] = responseMap

	return response.genericResponseMap
}
