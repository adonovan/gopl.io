package sha512

import (
	"crypto/sha512"
)


var byteArray [64]byte
var byte48Array [48]byte
func CommandLineHash(algo string, text []byte) string {
	
//fmt.Println(algo, text)
if algo == "SHA512" {
	byteArray = sha512.Sum512([]byte(text))
	return string(byteArray[0:])
	
}else if algo == "SHA384" {
	byte48Array = sha512.Sum384([]byte(text))
	return string(byte48Array[0:])
}
return "Not a valid algorithm"
}