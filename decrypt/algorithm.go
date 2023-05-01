package decrypt

func Nimbus(str string) string {
	decryptedStr := ""
	for _, letter := range str {
		asciiCode := int32(letter)
		character := string(asciiCode - 3)
		decryptedStr += character
	}
	return decryptedStr
}
