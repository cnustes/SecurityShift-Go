package email

import (
	"fmt"
	"libs/github.com/zemirco/email"
)

/**
     * estructura del envió del email.
**/

type StructSendEmail struct {
	From    string
	To      string
	Subject string
	Body    string
}

/**
     * Método para cambiar el password en el proceso de recuperar password.
     *
     * @emailTo     emails a los que se va enviar el correo([]string).
     * @host        host de la cuenta de correo de envió(string).
     * @emailserver email quien envia el correo(string).
     * @password    password de la cuenta de correo(string).
     * @subject     asunto del correo(string).
     * @body        mensaje del correo(string).
     * @port        Puerto de protocolo de envió de correo(int).
     * @return      string de respuesta.
     * @return      boolean de estado de respuesta.
**/

func SendEmail2(emailTo []string, host, emailserver, password, subject, body string, port int) (erro error) {

	mail := email.Mail{
		From:    emailserver,
		To:      emailTo[0],
		Subject: subject,
		HTML:    body,
	}

	err := mail.Send(host, port, emailserver, password)

	if err != nil {
		fmt.Println("error mail2---->", err)
	}

	return err
}
