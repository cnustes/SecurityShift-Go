package util

/**
 *
 * util
 *
 * Clase donde se generan metodos utiles para toda la aplicación
 *
 * @author Jeison Gaviria
 * @since 10/08/2015
 *
 * Derechos Reservados de Autor. (C) IP Total Software S.A
 *
 * Historia de Modificaciones
 * ------------------------------------------------------------- Autor	Fecha
 * Modificacion ----------------- -------------- ----------------------------
 */

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	//"strconv"
	"strings"
	"time"
)

/**
     * estructira para generar el token.
**/

type tokenStruct struct {
	Name      string
	TimeStart time.Time
}

var stringRandom string = ""

/**
     * Método que concatena un array de string.
     *
     * @data    array de string([]string).
     * @return  string concatenado
**/

func Concatenate(data []string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(data); i++ {
		buffer.WriteString(data[i])
	}

	return buffer.String()
}

/**
     * Método que genera un token para cada usuario.
     *
     * @name    nombre del usuario(string).
     * @return  string token
**/

func GeneratedToken(name string) string {

	token := tokenStruct{name, time.Now()}
	tokenData, _ := json.Marshal(token)
	tokenMD5 := md5.Sum(tokenData)
	tokenMD5String := hex.EncodeToString(tokenMD5[:])

	return tokenMD5String

}

/**
     * Método que genera un code de cuatro digitos pra mensaje de activación sms.
     *
     * @return  string code
**/

func GeneratedCode() string {

	/*rand.Seed(time.Now().UTC().UnixNano())
	var code string

	for i := 0; i < 4; i++ {
		bytes := rand.Intn(9)
		s := strconv.Itoa(bytes)
		code = code + s
	}
	return code*/

	var keyCode = []rune(GeneratedStringRand())

	code := make([]rune, 7)
	for i := range code {
		code[i] = keyCode[rand.Intn(len(keyCode))]
	}
	return string(code)
}

/**
     * Método que genera un code de cuatro digitos pra mensaje de activación sms.
     *
     * @return  string code
**/

func GeneratedStringRand() string {
	if strings.Compare(stringRandom, "") == 0 {
		stringRandom = "1234567890abcdefghijklmnopqrstuvw,xyzABCDEFGHIJKLMNOPQRSTUVWXYZ.;-:,)(/&%$#?¿¡°!lkjhkahsflhpfha{kfm}q,aefqrfevferkjhkjakjkwfkhkwkfekhqk,ejfhqkjhfqkjhfkjerhfkjhefadfwk227,36980294r7349818018370=)/(/&%&$#u"
	} else {
		arrayString := strings.SplitAfterN(stringRandom, ",", 6)
		stringRandom = ""
		for i := 0; i < len(arrayString); i++ {
			stringRandom = Concatenate([]string{stringRandom, arrayString[rand.Intn(len(arrayString))]})
		}
	}

	return stringRandom
}

/**
     * Método que lee un archivo.
     *
     * @filename    ruta con nombre del archivo(string).
     * @return  lineas del archivo leido ([]string).
     * @return  Error al efectuar el proceso(err).
     * @return  Código del error (int:  1= error al leer al archivo(no existe) , 2= error al recorrer el archivo).
**/

func ReadFile(filename string) (linesResult []string, err error, codeerror int) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err, 1
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return []string{}, scanner.Err(), 2

	}

	return lines, nil, 0
}

/**
     * Método que crea o sobreescribe un archivo.
     *
     * @lines   lineas con las que vamos a crear o a sobreescribir el archivo([]string).
     * @path    ruta con nombre del archivo(string).
     * @return  Error al efectuar el proceso(err).
     * @return  Código del error (int:  1= error al crear el archivo , 2= error al sonreescribir el archivo).
**/

func WriteFile(lines []string, path string) (errores error, codeerror int) {
	file, err := os.Create(path)
	if err != nil {
		return err, 1
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush(), 2
}

/**
     * Método que encripta una cadena.
     *
     * @msg     cadena que se desea encriptar(string).
     * @key     llave con la que se va a encriptar(string).
     * @return  cadena ya encriptada(string).
**/

func Encrypt(msg, key string) string {

	encrypted := make([]byte, len(msg))

	err := EncryptAESCFB(encrypted, []byte(msg), []byte(key))
	if err != nil {
		//panic(err)
		fmt.Println("error cryp", err)
	}

	return string(encrypted)
}

/**
     * Método que desencripta una cadena.
     *
     * @msg     cadena que se desea desencriptar(string).
     * @key     llave con la que se encripto la cadena(string).
     * @return  cadena ya desencriptada(string).
**/

func Decrypt(msg, key string) string {

	decrypted := make([]byte, len(msg))
	err := DecryptAESCFB(decrypted, []byte(msg), []byte(key))
	if err != nil {
		//panic(err)
		fmt.Println("error decryp")
	}

	return string(decrypted)

}

/**
     * Método que encripta en bytes.
     *
     * @dst     bytes del tamaño de la cadena a encriptar([]bytes).
     * @src     bytes de la cadena a encriptar([]bytes).
     * @key     bytes de llave con la que se va a encriptar la cadena([]bytes).
     * @return  cualquier error que pueda ocurrir(error).
**/

func EncryptAESCFB(dst, src, key []byte) error {
	var iv = []byte(key)[:aes.BlockSize]
	aesBlockEncrypter, err := aes.NewCipher([]byte(key))

	if err != nil {
		return err
	}

	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(dst, src)
	return nil
}

/**
     * Método que desencripta en bytes.
     *
     * @dst     bytes del tamaño de la cadena a desencriptar([]bytes).
     * @src     bytes de la cadena a desencriptar([]bytes).
     * @key     bytes de la llave con la que se encripto la cadena([]bytes).
     * @return  cualquier error que pueda ocurrir(error).
**/

func DecryptAESCFB(dst, src, key []byte) error {
	var iv = []byte(key)[:aes.BlockSize]
	aesBlockDecrypter, err := aes.NewCipher([]byte(key))

	if err != nil {
		return nil
	}

	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(dst, src)
	return nil
}
