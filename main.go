package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func contains(s string, guess string) bool {
	for _, ch := range s {
		if string(ch) == guess {
			return true
		}
	}
	return false
}

func getLetter(guess []string) string {
	fmt.Println("Enter letter:", guess)
	var first string
	fmt.Scanln(&first)
	for len(first) > 1 {
		fmt.Println("Invalid input. Please enter only one character.")
		fmt.Scanln(&first)
	}
	return first
}

func updateGuess(guess []string, letter string, word string) bool {
	correct := true
	for i := 0; i < len(word); i++ {
		if letter == string(word[i]) {
			guess[i] = letter
		}
		if guess[i] == "_" {
			correct = false
		}
	}
	return correct
}

func hangman(words []string) {
	rand.Seed(time.Now().UnixNano())
	idr := rand.Intn(len(words))
	word := words[idr]
	var seen string = ""
	guessNumber := 6
	guess := []string{}
	pics := []string{"  +---+\n  |   |\n      |\n      |\n      |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n      |\n      |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n  |   |\n      |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n /|   |\n      |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n /|\\  |\n      |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n /|\\  |\n /    |\n      |\n=========",

		"  +---+\n  |   |\n  O   |\n /|\\  |\n / \\  |\n      |\n========="}
	for i := 0; i < len(word); i++ {
		guess = append(guess, "_")
	}
	for guessNumber > 0 {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println(pics[6-guessNumber])
		fmt.Println("Remaining number of guesses:", guessNumber)
		fmt.Println("Your guesses: ", seen)
		letter := getLetter(guess)
		if !contains(word, letter) || contains(seen, letter) {
			guessNumber--
		}
		if !contains(seen, letter) {
			seen += letter
		}
		done := updateGuess(guess, letter, word)
		if done {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Println("You win! Your guess is correct. The word was:", word)
			return
		}
	}
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println(pics[6-guessNumber])
	fmt.Println("Your man died! Your guess is not correct. The word was:", word)
	return
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("There is no input file!")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var words []string = []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	hangman(words)

}
