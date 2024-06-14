package ascii

import (
	"errors"
	"os"
	"strings"
)

const chunkSize int = 9

// Convert .txt file of the theme into lines.
func ThemeToLines(s string) []string {
	data, _ := os.ReadFile("pkg/ascii/theme/" + s + ".txt")
	return strings.Split(string(data), "\n")
}

// Get text input from arguments and split it into differente lines.
func GetTextInput(s string) ([]string, error) {
	inputCorrected := strings.Replace(s, "\\n", "\n", -1)
	inputCorrected = strings.Replace(inputCorrected, string(rune(10))+string(rune(13)), "\n", -1)
	inputCorrected = strings.Replace(inputCorrected, string(rune(13))+string(rune(10)), "\n", -1)
	inputCorrected = strings.Replace(inputCorrected, "à", "a", -1)
	inputCorrected = strings.Replace(inputCorrected, "é", "e", -1)
	inputCorrected = strings.Replace(inputCorrected, "è", "e", -1)
	inputCorrected = strings.Replace(inputCorrected, "ç", "c", -1)
	inputCorrected = strings.Replace(inputCorrected, "ù", "u", -1)
	if verifyChar(inputCorrected) {
		return nil, errors.New("invalid char")
	}
	return strings.Split(inputCorrected, "\n"), nil
}

func verifyChar(s string) bool {
	for _, char := range s {
		if !(char >= 32 && char <= 126) && char != 10 {
			return true
		}
	}
	return false
}

// Print the ascii-art.
func PrintAsciiArt(input, lines []string) string {
	s := ""
	for _, inputLine := range input {
		if inputLine == "" {
			s += "\n"
		} else {
			for line := 1; line < chunkSize; line++ {
				for char := 0; char < len(inputLine); char++ {
					characterStart := (int(inputLine[char]) - 32) * chunkSize
					s += lines[characterStart+line]
				}
				s += "\n"
			}
		}
	}
	return s
}
