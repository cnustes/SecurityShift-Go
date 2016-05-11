package dao

/**
 *
 * securityShift dao user
 *
 * Clase dao user
 *
 * @author Jeison Gaviria
 * @since 10/05/2016
 *
 *
 * Historia de Modificaciones
 * ------------------------------------------------------------- Autor	Fecha
 * Modificacion ----------------- -------------- ----------------------------
 */


/**
     * estructura de los usuarios.
**/
type User struct {
	Id           string
	Password     string
	Status       string
	Email        string
	CellPhone    string
}