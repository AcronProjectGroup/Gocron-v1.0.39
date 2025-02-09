package features

import (
	"fmt"
	"strings"
	"time"
	"github.com/fatih/color"
	"os/exec"
	"runtime"
	"os"
)

func SplitIntoWords(text string) []string {
	// Split the text into words
	return splitAndTrim(text, " ")
}

func splitAndTrim(s, sep string) []string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.TrimSpace(word)
	}
	return words
}

func PrintWordByWord(words []string) {
	counter := 0
	for _, word := range words {
		counter++
		if counter == 12 {
			fmt.Println()
			counter = 0
		}
		fmt.Print(word + " ")
		time.Sleep(85 * time.Millisecond) // Adjust the delay as needed
	}
}

func GoodByePrint(){
	color.HiBlue(`
┌─────────────────┐
|    GoodBye      |
└─────────────────┘
	`)
}

func ClearTerminal() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Println("Unsupported operating system")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
