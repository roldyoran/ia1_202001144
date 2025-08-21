package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Estado del puzzle
type PuzzleState [9]int

var (
	goalState    = PuzzleState{1, 2, 3, 4, 5, 6, 7, 8, 0}
	currentState = PuzzleState{1, 2, 3, 4, 5, 6, 7, 8, 0}
	solutionPath []PuzzleState
	currentStep  int
	shuffleSteps = 20
)

// Encontrar la posición vacía
func findEmpty(state PuzzleState) int {
	for i, val := range state {
		if val == 0 {
			return i
		}
	}
	return -1
}

// Verificar si el movimiento es válido
func isValidMove(pos, newPos int) bool {
	row, col := pos/3, pos%3
	newRow, newCol := newPos/3, newPos%3
	return (row == newRow && (col == newCol-1 || col == newCol+1)) ||
		(col == newCol && (row == newRow-1 || row == newRow+1))
}

// Mostrar el puzzle en consola
func printPuzzle(state PuzzleState) {
	for i := 0; i < 9; i++ {
		if state[i] == 0 {
			fmt.Print("   ")
		} else {
			fmt.Printf(" %d ", state[i])
		}
		if i%3 == 2 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// Desordenar el puzzle
func shuffleState() {
	currentState = goalState
	emptyPos := findEmpty(currentState)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < shuffleSteps; i++ {
		possibleMoves := []int{}
		for move := 0; move < 9; move++ {
			if move != emptyPos && isValidMove(emptyPos, move) {
				possibleMoves = append(possibleMoves, move)
			}
		}
		if len(possibleMoves) > 0 {
			randomMove := possibleMoves[rand.Intn(len(possibleMoves))]
			currentState[emptyPos], currentState[randomMove] = currentState[randomMove], currentState[emptyPos]
			emptyPos = randomMove
		}
	}
}

// Resolver puzzle con BFS
func solvePuzzle() {
	solutionPath = nil
	currentStep = 0

	queue := [][]PuzzleState{{currentState}}
	visited := make(map[PuzzleState]bool)
	visited[currentState] = true

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		currentNode := path[len(path)-1]

		if currentNode == goalState {
			solutionPath = path
			return
		}

		emptyPos := findEmpty(currentNode)
		for move := 0; move < 9; move++ {
			if move != emptyPos && isValidMove(emptyPos, move) {
				newState := currentNode
				newState[emptyPos], newState[move] = newState[move], newState[emptyPos]

				if !visited[newState] {
					visited[newState] = true
					newPath := make([]PuzzleState, len(path))
					copy(newPath, path)
					newPath = append(newPath, newState)
					queue = append(queue, newPath)
				}
			}
		}
	}
}

// Función para mover pieza en consola
func movePiece(pos int) {
	emptyPos := findEmpty(currentState)
	if isValidMove(emptyPos, pos) {
		currentState[emptyPos], currentState[pos] = currentState[pos], currentState[emptyPos]
	}
}

// Menú principal en consola
func main() {
	var choice int
	for {
		fmt.Println("8 Puzzle Solver - Terminal")
		printPuzzle(currentState)
		fmt.Println("1. Iniciar (Reset)")
		fmt.Println("2. Desordenar")
		fmt.Println("3. Resolver automáticamente")
		fmt.Println("4. Paso a paso")
		fmt.Println("5. Mover pieza")
		fmt.Println("0. Salir")
		fmt.Print("Opción: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			currentState = goalState
			solutionPath = nil
			currentStep = 0
		case 2:
			shuffleState()
			solutionPath = nil
			currentStep = 0
		case 3:
			solvePuzzle()
			if len(solutionPath) > 0 {
				currentState = solutionPath[len(solutionPath)-1]
			} else {
				fmt.Println("No se encontró solución")
			}
		case 4:
			if len(solutionPath) == 0 {
				solvePuzzle()
			}
			if currentStep < len(solutionPath) {
				currentState = solutionPath[currentStep]
				currentStep++
			} else {
				fmt.Println("Se completó la solución paso a paso")
			}
		case 5:
			var pos int
			fmt.Print("Número de pieza a mover (0-8): ")
			fmt.Scan(&pos)
			if pos >= 0 && pos <= 8 {
				movePiece(pos)
			} else {
				fmt.Println("Valor inválido")
			}
		case 0:
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}
