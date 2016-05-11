package htmlServer

/**
 *
 * html
 *
 * Clase donde estan la gui de configuración y los métodos de 
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
	"fmt"
	"net/http"
)

/**
     * Método que representa el servicio que permite crear la gui de test de servicio https.
     *
     * @w   http response(http.ResponseWriter).
     * @r   http request(http.Request).
**/

func ViewTestHttps(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "<html> ")
		/****/ fmt.Fprintln(w, "<head> ")
		/****/ fmt.Fprintln(w, "</head> ")
		/****/ fmt.Fprintln(w, "<body> ")
		/********/ fmt.Fprintln(w, "<div id='dialogal'> ")
		/************/ fmt.Fprintln(w, "<div> ")
		/****************/ fmt.Fprintln(w, "<div> ")
		/********************/ fmt.Fprintln(w, "<h4>test servidor https</h4> ")
		/****************/ fmt.Fprintln(w, "</div> ")
		/****************/ fmt.Fprintln(w, "<div> ")
		/********************/ fmt.Fprintln(w, "<div>  ")
		/************************/ fmt.Fprintln(w, "<strong>!si funciono </strong>")
		/********************/ fmt.Fprintln(w, "</div> ")
		/****************/ fmt.Fprintln(w, "</div> ")
		/************/ fmt.Fprintln(w, "</div> ")
		/********/ fmt.Fprintln(w, "</div> ")
		/****/ fmt.Fprintln(w, "</body> ")
		fmt.Fprintln(w, "</html> ")

}