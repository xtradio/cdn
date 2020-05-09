package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func randomString() string {
	rand.Seed(time.Now().Unix())

	var output strings.Builder

	charSet := []rune("abcdedfghijklmnopqrstABCDEFGHIJKLMNOP01234567890")
	length := 20
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteRune(randomChar)
	}

	return output.String()
}

func doesExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func generateFilename(imgDir string) (string, string) {
	filename := fmt.Sprintf("%s.%s", randomString(), "jpg")
	filepath := fmt.Sprintf("%s/%s", imgDir, filename)

	for doesExist(filepath) != false {
		fmt.Print("Regenerating file name.")
		filename = fmt.Sprintf("%s.%s", randomString(), "jpg")
		filepath = fmt.Sprintf("%s/%s", imgDir, filename)
	}
	return filename, filepath

}
