package sms

/**
 *
 * sms
 *
 * Clase donde se generan metodos para manejo de sms.
 *
 * @author Jeison Gaviria
 * @since 28/09/2015
 *
 * Derechos Reservados de Autor. (C) IP Total Software S.A
 *
 * Historia de Modificaciones
 * ------------------------------------------------------------- Autor	Fecha
 * Modificacion ----------------- -------------- ----------------------------
 */

import (
	"libs/github.com/njern/gonexmo"
	"strconv"
	"time"
	"fmt"
)

func SendSMS(keySMS, secretSMS, fromSMS, messageSMS, toSMS string) (string, bool) {
	nexmoClient, errcli := nexmo.NewClientFromAPI(keySMS, secretSMS)

	if errcli != nil {
		return "revise su llave y su codigo secreto para cliente de sms", false
	}
	message := &nexmo.SMSMessage{
		From:            fromSMS,
		To:              toSMS,
		Type:            nexmo.Text,
		Text:            messageSMS,
		ClientReference: strconv.FormatInt(time.Now().Unix(), 10),
		Class:           nexmo.Standard,
	}

	meesagetext, err := nexmoClient.SMS.Send(message)

	if err != nil {
		return "ocurrió un error enviando sms", false
	}
	messageReport := meesagetext.Messages
	fmt.Println(messageReport)
	if messageReport[0].Status.String() != "Success" {
		return "ocurrió un error de datos de servidor de sms", false
	}

	return "mensaje enviado", true

}
