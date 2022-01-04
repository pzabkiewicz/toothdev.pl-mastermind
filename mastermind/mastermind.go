package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	COLS             = 4
	ROWS             = 9
	COLOR_SYMBOLS_NO = 8
	RIGHT_COLOR_NUM  = 1
	RIGHT_SPOT_NUM   = 2
)

var (
	NUM_2_SYMBOL = map[int]string{
		0:  "_",
		1:  "x",
		2:  "o",
		3:  "G",
		4:  "O",
		5:  "Y",
		6:  "B",
		7:  "K",
		8:  "R",
		9:  "C",
		10: "V",
	}
	SYMBOLS_2_NUMS = reverseMapping(NUM_2_SYMBOL)
)

func main() {
	secret := InitSecret()
	board := InitMatrix()
	hints := InitMatrix()
	attempt := 0
	guessed := false

	printBoard(board, hints)

	for !guessed && attempt < ROWS {
		fmt.Println("******** Attempt #", attempt+1, " ********")
		printInstructions()
		reader := bufio.NewReader(os.Stdin)
		playerInput, err := reader.ReadString('\n')
		// err -> wrong input
		check(err)
		playerInput = strings.TrimSpace(playerInput)
		symbols := strings.Split(playerInput, "-")
		// validate color's symbols
		guesses := symbols2nums(symbols)
		attempt++
		UpdateMatrix(&board, guesses, attempt)
		currentHints := AnalyzeGuessesAndGetHints(secret, guesses)
		UpdateMatrix(&hints, currentHints, attempt)
		guessed = CheckWin(currentHints)
		printBoard(board, hints)
	}
	printResult(guessed, attempt, secret)
}

func reverseMapping(num2symbols map[int]string) map[string]int {
	symbols2nums := make(map[string]int)
	for k, v := range num2symbols {
		symbols2nums[v] = k
	}
	return symbols2nums
}

func InitSecret() []int {
	rand.Seed(time.Now().Unix())
	var secret [COLS]int
	scopeStart := len(NUM_2_SYMBOL) - COLOR_SYMBOLS_NO
	for i, _ := range secret {
		symbolNum := rand.Intn(COLOR_SYMBOLS_NO) + scopeStart
		secret[i] = symbolNum
	}
	return secret[:]
}

func InitMatrix() [][]int {
	matrix := make([][]int, ROWS)
	for i, _ := range matrix {
		matrix[i] = make([]int, COLS)
	}
	return matrix
}

func symbols2nums(symbols []string) []int {
	var nums [COLS]int
	for i, symbol := range symbols {
		nums[i] = SYMBOLS_2_NUMS[symbol]
	}
	return nums[:]
}

func CheckWin(hints []int) bool {
	rightSpotCount := 0
	for _, hint := range hints {
		if hint == RIGHT_SPOT_NUM {
			rightSpotCount++
		}
	}
	return COLS == rightSpotCount
}

func UpdateMatrix(matrixPointer *[][]int, tab []int, attempt int) {
	matrix := *matrixPointer
	row := matrix[ROWS-attempt]
	for i, tabElement := range tab {
		row[i] = tabElement
	}
}

func AnalyzeGuessesAndGetHints(secret []int, guesses []int) []int {
	guessedColors := 0
	guessedSpots := 0
	processedColors := make([]int, len(guesses))
	for i, secretElement := range secret {
		if secretElement == guesses[i] {
			guessedSpots++
		}
		for j, guess := range guesses {
			if processedColors[j] == guess {
				continue
			}
			if secretElement == guess {
				processedColors[j] = guess
				guessedColors++
				break
			}
		}
	}

	hints := make([]int, COLS)

	for i := 0; i < guessedSpots; i++ {
		hints[i] = RIGHT_SPOT_NUM
	}

	if COLS == guessedSpots {
		return hints
	}

	colorsCountToMark := guessedColors - guessedSpots
	for i := guessedSpots; i < guessedSpots+colorsCountToMark; i++ {
		hints[i] = RIGHT_COLOR_NUM
	}
	return hints
}

func printBoard(board [][]int, hints [][]int) {
	for i, row := range board {
		fmt.Print(ROWS-i, ": ")
		for _, col := range row {
			fmt.Print(NUM_2_SYMBOL[col], " ")
		}
		fmt.Print(" | ")
		for _, hint := range hints[i] {
			fmt.Print(NUM_2_SYMBOL[hint], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func printInstructions() {
	fmt.Println()
	fmt.Println("Legend: ")
	fmt.Println("Color symbols: G - Green, O - Orange, Y - Yellow, B - Blue, K - Black, R - Red, C - Cyan, V - violet")
	fmt.Println("Hints: \n  - correct color - 'x'\n  - correct color and spot - 'o'\n  - no guess - '_'") 
	fmt.Println("Please choose 4 of the above given color symbols delimited with '-'")
	fmt.Println()
}

func printResult(guessed bool, attempt int, secret []int) {
	if guessed {
		fmt.Println("Perfect! You guessed the secret! The attempt number is:", attempt)
	} else {
		fmt.Println("Unlucky, unlucky! Unfortunetely you didn't guess the secret")
	}
	symbolSecret := make([]string, len(secret))
	for i, secretElement := range secret {
		symbolSecret[i] = NUM_2_SYMBOL[secretElement]
	}
	fmt.Println("The secret is:", strings.Join(symbolSecret, "-"))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
