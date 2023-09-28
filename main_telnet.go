package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
)

type myHandler struct{}

func (handler myHandler) ServeTELNET(ctx telnet.Context, w telnet.Writer, r telnet.Reader) {
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

	var buffer [1]byte
	p := buffer[:]
	oi.LongWrite(w, []byte("\033[H\033[2J"))
	oi.LongWrite(w, []byte(pics[6-guessNumber]))
	stringGuess := strings.Join(guess, " ")
	oi.LongWrite(w, []byte("\n"))
	oi.LongWrite(w, []byte(stringGuess))
	oi.LongWrite(w, []byte("\n"))
	oi.LongWrite(w, []byte("\n"))
	oi.LongWrite(w, []byte("Remaining number of guesses: "))
	oi.LongWrite(w, []byte(strconv.Itoa(guessNumber)))
	oi.LongWrite(w, []byte("\n"))
	oi.LongWrite(w, []byte("Your guesses: "))
	oi.LongWrite(w, []byte(seen))
	oi.LongWrite(w, []byte("\n"))
	oi.LongWrite(w, []byte("Enter new letter: "))

	var buf [1]byte
	first := buf[:]

	for {
		r.Read(first)
		counter := 0
		correct := true
		for {
			r.Read(p)
			counter++
			if counter == 3 {
				oi.LongWrite(w, []byte("Invalid input. Please enter only one character. \n"))
				correct = false
			}
			if p[0] == 10 {
				break
			}
		}
		if correct {
			letter := string(string(first)[0])
			if !contains(word, letter) || contains(seen, letter) {
				guessNumber--
			}
			if !contains(seen, letter) {
				seen += letter
			}
			if guessNumber == 0 {
				oi.LongWrite(w, []byte(pics[6-guessNumber]))
				oi.LongWrite(w, []byte("\n"))
				oi.LongWrite(w, []byte("Your man died! Your guess is not correct. The word was: "))
				oi.LongWrite(w, []byte(word))
				return
			}
			done := updateGuess(guess, letter, word)
			if done {
				oi.LongWrite(w, []byte("You win! Your guess is correct. The word was: "))
				oi.LongWrite(w, []byte(word))
				return
			}
			oi.LongWrite(w, []byte(pics[6-guessNumber]))
			stringGuess := strings.Join(guess, " ")
			oi.LongWrite(w, []byte("\n"))
			oi.LongWrite(w, []byte(stringGuess))
			oi.LongWrite(w, []byte("\n"))
			oi.LongWrite(w, []byte("\n"))
			oi.LongWrite(w, []byte("Remaining number of guesses: "))
			oi.LongWrite(w, []byte(strconv.Itoa(guessNumber)))
			oi.LongWrite(w, []byte("\n"))
			oi.LongWrite(w, []byte("Your guesses: "))
			oi.LongWrite(w, []byte(seen))
			oi.LongWrite(w, []byte("\n"))
			oi.LongWrite(w, []byte("Enter new letter: "))

		}

	}
}

func contains(s string, guess string) bool {
	for _, ch := range s {
		if string(ch) == guess {
			return true
		}
	}
	return false
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
	var handler myHandler

	erro := telnet.ListenAndServe(":5555", handler)
	if nil != erro {

		panic(erro)
	}
}
