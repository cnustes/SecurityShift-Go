package dao

/**
 *
 * securityShift dao configuration
 *
 * Clase dao configuration
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
     * estructura de la configuración.
**/

type Configuration struct {
	//HostServer         string  //host de servidor.
	PortServer         string  //puerto del servidor por donde van a escuchar las peticiones de securityShift(nombre del archivo .pem)
	Private_key        string  //certificado privado (nombre del archivo .pem)
	Public_key         string  //certificado público (nombre del archivo .pem)
	/*Numfile            int     //Numero de intentos de loguin
	TimeSessionUser    float64 //Tiempo de session, sin actividad en segundos
	TimeSessionDestroy int     //tiempo en minutos que se activa el liberador de sessiones caducadas
	EmailServer        string  //Email para envio de mensajes
	PasswordEmail      string  //contraseña del email
	HostMail           string  //Mail del correo de
	PortMail           int     //Puerto de mail
	Subject            string  //Asunto del mensaje de recuperación de password
	Message            string  //mensage del correo para recuperación de password
	KeySMS             string  //llave de cliente de Gonexmo
	SecretSMS          string  //codigo secreto de cliente de Gonexmo
	FromSMS            string  //numero de celular de envio de mensajes*/
}
