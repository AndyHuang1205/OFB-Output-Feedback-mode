/*
3/26/2023
OFB Output Feedback mode
*/
package main

import (
	"fmt"
	"strconv"
)

/* an array with 4 rows and 2 columns*/
var codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var iv int = 0b10
var lookupValue int = 0
var xor int = 0

func textToBinary(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%.8b", binString, c)
	}
	return binString
}
func binaryToString(binaryString string) (textString string) {
	textString = ""
	var start, end int = 0, 8
	for end < len(binaryString)+1 {
		binaryInt, _ := strconv.ParseInt(binaryString[start:end], 2, 0)
		text := string(binaryInt)
		textString += text
		end += 8
		start += 8
	}
	return textString
}

func codebookLookup(xor int) (lookupValue int) {
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		if codebook[i][j] == xor {
			j++
			lookupValue = codebook[i][j]
			break
		}
	}
	return lookupValue
}
func OFB_ciphering(text string) (ciphertext string) {
	// Ciphering
	var i int = 0
	for i = 0; i < len(text)-1; i += 2 {
		lookupValue = codebookLookup(iv)
		bin, _ := strconv.ParseInt(text[i:i+2], 2, 64)
		xor = int(bin) ^ lookupValue
		iv = lookupValue
		ciphertext += fmt.Sprintf("%02b", xor)
		fmt.Printf("The ciphered value is %b\n", xor)
	}
	fmt.Println("Ciphertext: " + ciphertext)
	return ciphertext
}

func OFB_deciphering(ciphertext string) (plaintext string) {
	var plaintextBinary string = ""
	var j int = 0
	// Deciphering
	for j = 0; j < len(ciphertext)-1; j += 2 {
		lookupValue = codebookLookup(iv)
		iv = lookupValue
		bin, _ := strconv.ParseInt(ciphertext[j:j+2], 2, 64)
		xor = int(bin) ^ lookupValue
		plaintextBinary += fmt.Sprintf("%02b", xor)
		fmt.Printf("The deciphered value is %b\n", xor)
	}
	plaintext = binaryToString(plaintextBinary)
	fmt.Printf("The deciphered text is: " + plaintext)
	return plaintext
}

func main() {
	var text string
	fmt.Printf("Enter your message: ")
	fmt.Scanln(&text)
	text = textToBinary(text)
	fmt.Println("Binary plaintext: " + text)

	ciphertext := OFB_ciphering(text)
	OFB_deciphering(ciphertext)
}
