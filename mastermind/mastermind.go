package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

const COLS = 4
const ROWS = 9
const COLOR_SYMBOLS_NO = 8
const RIGHT_COLOR_NUM = 1
const RIGHT_SPOT_NUM = 2

var NUM_2_SYMBOL = map[int]string{
	0:"_",
	1:"x",
	2:"o",
	3:"G",
	4:"O",
	5:"Y",
	6:"B",
	7:"K",
	8:"R",
	9:"C",
	10:"V",
}
var SYMBOLS_2_NUMS = reverseMapping(NUM_2_SYMBOL)

func main() {
	secret := initSecret()
	board := initMatrix()
	hints := initMatrix()
	attempt := 0
    guessed := false
	
	printBoard(board, hints)
	
	for !guessed && attempt < ROWS {
		attempt++
		fmt.Println("******** Attempt #", attempt, " ********")
		printInstructions()
		reader := bufio.NewReader(os.Stdin)
		playerInput, err := reader.ReadString('\n')
		// err -> wrong input
		check(err)
		playerInput = strings.TrimSpace(playerInput)
		symbols := strings.Split(playerInput, "-")
		// validate color's symbols
		guesses := symbols2nums(symbols)
		updateMatrix(&board, guesses, attempt)
		currentHints := analyzeGuessesAndGetHints(secret, guesses)
		updateMatrix(&hints, currentHints, attempt)
		guessed = checkWin(currentHints)
		printBoard(board, hints)
	}
	if guessed {
		fmt.Println("Perfect! You guessed the code! The attempt number is:", attempt) 
	} else {
		fmt.Println("Unlucky, unlucky! Unfortunetely you didn't guess the right code")
		fmt.Println("The right code is:", code)
	}
}

func reverseMapping(num2symbols map[int]string) map[string]int {
	symbols2nums := make(map[string]int)
	for k, v := range num2symbols {
		symbols2nums[v] = k
	}
	return symbols2nums
}

func initSecret() []int {
	var secret [COLS]int
	scopeStart := len(NUM_2_SYMBOL) - COLOR_SYMBOLS_NO
	for i, _ := range secret {
		symbolNum := rand.Intn(COLOR_SYMBOLS_NO) + scopeStart
		secret[i] = symbolNum
	}
	return code[:]
}

func initMatrix() [][]int {
	matrix := make([][]int, ROWS)
	for i, _ := range matrix {
		matrix[i] = make([]int, COLS)
	}
	for _, row := range matrix {
		for i, _ := range row {
			row[i] = 0
		}
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

func checkWin(hints []int) bool {
	rightSpotCount := 0
	for _, hint := range hints {
		if hint == RIGHT_SPOT_NUM {
			rightSpotCount++
		}
	}
	return COLS == rightSpotCount
}

func updateMatrix(matrixPointer *[][]int, tab []int, attempt int) {
	matrix := *matrixPointer
	row := matrix[ROWS - attempt]
	for i, tabElement := range tab {
		row[i] = tabElement
	}
}

func analyzeGuessesAndGetHints(secret []int, guesses []int) []int {
	guessedColors := 0
	guessedSpots := 0
	for i, secretElement := range secret {
		if guesses[i] == secretElement {
			guessedSpots++
		}
		for _, guess := range guesses {
			if secretElement == guess {
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
	for i := guessedSpots; i < guessedSpots + colorsCountToMark; i++ {
		hints[i] = RIGHT_COLOR_NUM
	}
	return hints
}

func printBoard(board [][]int, hints [][]int) {
	for i, row := range board {
		fmt.Print(ROWS - i, ": ")
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
	fmt.Println("Symbols: G - Green, O - Orange, Y - Yellow, B - Blue, K - Black, R - Red, C - Cyan, V - violet")
	fmt.Println("Please choose 4 of the above given color symbols delimited with '-'")
	fmt.Println()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}