package css

/**
 *
 * styleCssConfig
 *
 * Clase css para los estilos de la viewConfig
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
     * Método que representa el método css para los estilos de la hoja de configuración.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func CssViewConfig(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "<style type='text/css'> ")
		fmt.Fprintln(w, "@charset 'utf-8'  ")
		fmt.Fprintln(w, "body{	overflow: hidden; }  ")
		fmt.Fprintln(w, "</style> ")

		fmt.Fprintln(w, "<style type='text/css'> ")
		fmt.Fprintln(w, ".fontLabel{ color: 5d5d5d; font-size: 12px;font-family: sans-serif; }  ")
		fmt.Fprintln(w, "</style> ")

		fmt.Fprintln(w, "<style type='text/css'> ")
		fmt.Fprintln(w, ".opaque{ background-color: rgba(255, 250, 250, 0.7); border-radius: 4px; box-shadow: 0px 1px 2px rgba(0,0,0,0.45); }  ")
		fmt.Fprintln(w, "</style> ")

		fmt.Fprintln(w, "<style type='text/css'> ")
		fmt.Fprintln(w, ".marginInput{ margin-top: 6px; margin-bottom: 6px; } ")
		fmt.Fprintln(w, "</style> ")

		fmt.Fprintln(w, "<style type='text/css'> ")
		fmt.Fprintln(w, ".div-labelTitle{ font-size: 24px; font-family: sans-serif; color: #406085;} ")
		fmt.Fprintln(w, "</style> ")

		fmt.Fprintln(w, "<style type='text/css'> ")
		fmt.Fprintln(w, ".float-center{ margin-top: 3%; margin-left: 20%; margin-right: 20%; }  ")
		fmt.Fprintln(w, "</style> ")

		fmt.Fprintln(w, "<style type='text/css'> ")
		fmt.Fprintln(w, ".colorHeader{ background-color: #FDA910; color: aliceblue;} ")
		fmt.Fprintln(w, "</style> ")

		fmt.Fprintln(w, "<style type='text/css'> ")
		fmt.Fprintln(w, ".colorHeaderSucces{ background-color: #21981F; color: aliceblue;} ")
		fmt.Fprintln(w, "</style> ")

	}