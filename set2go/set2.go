// Placeholder
package main

import (
	"fmt"
)

func main() {
	fmt.Println(vigenere_cipher("A C", "KEY"))         // Example usage
	fmt.Println(vigenere_cipher("hello world", "key")) // Example usage

}

func shiftLetter(letter string, shift int) string {
	if letter == " " {
		return " "
	}

	if len(letter) != 1 || (letter[0] < 'A' || letter[0] > 'Z') {
		return " "
	}

	asciiValue := int(letter[0])
	newAscii := ((asciiValue-'A')+shift)%26 + 'A'
	return string(newAscii)
}

// utilizes shiftLetter function
func caesarCipher(message string, shift int) string {
	cipherResult := ""

	for _, letter := range message {
		letterStr := string(letter)

		//references shiftLetter function from above
		cipherResult += shiftLetter(letterStr, shift)
	}

	return cipherResult
}

func shiftByLetter(letter string, letterShift string) string {
	if letter == " " {
		return " "
	}

	letterValue := int(letter[0] - 'A')
	shiftValue := int(letterShift[0] - 'A')

	shiftedValue := (letterValue + shiftValue) % 26
	return string(byte(shiftedValue + 'A'))
}

func vigenere_cipher(message, key string) string {
	keyLen := len(key)
	messageLen := len(message)
	extendedKey := []rune(key)

	for i := 0; len(extendedKey) < messageLen; i++ {
		extendedKey = append(extendedKey, rune(key[i%keyLen]))
	}

	encryptedText := []rune(message)

	for i, char := range message {
		if char >= 'A' && char <= 'Z' {
			encryptedText[i] = rune((int(char-'A')+int(extendedKey[i]-'A'))%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			encryptedText[i] = rune((int(char-'a')+int(extendedKey[i]-'a'))%26 + 'a')
		}
	}

	return string(encryptedText)
}

// error concerns about taking longer than 7 seconds to run
func scytaleCipher(message string, shift int) string {
	messageLen := len(message)
	if messageLen%shift != 0 {
		for i := messageLen % shift; i < shift; i++ {
			message += "_"
		}
		messageLen = len(message)
	}

	result := make([]byte, messageLen)
	for i := 0; i < messageLen; i++ {
		result[i] = message[(i/shift)+(messageLen/shift)*(i%shift)]
	}
	return string(result)
}

func scytaleDecipher(message string, shift int) string {
	messageLen := len(message)
	result := make([]byte, messageLen)

	for i := 0; i < messageLen; i++ {
		result[(i%shift)*(messageLen/shift)+i/shift] = message[i]
	}
	return string(result)
}
