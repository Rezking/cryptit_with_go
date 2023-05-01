package encrypt

func Nimbus(str string) string {
	encryptedStr := ""
	for _, letter := range str {
		asciiCode := int32(letter)
		character := string(asciiCode + 3)
		encryptedStr += character
	}
	return encryptedStr
}
