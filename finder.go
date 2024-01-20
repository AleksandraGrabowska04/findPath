package finder

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth        = 800
	windowHeight       = 800
	gridSize           = 20
	fontPath           = "fonts/Terminal.ttf"
	fontSize           = 24
	numRandomSquares   = 300
	maxSquarePositionx = 800
	maxSquarePositiony = 800
)

var (
	startPointColor = sdl.Color{R: 0, G: 255, B: 0, A: 255}
	endPointColor   = sdl.Color{R: 255, G: 0, B: 0, A: 255}
)

type Point struct {
	X     int32
	Y     int32
	Color sdl.Color
}

type Square struct {
	Rect  *sdl.Rect
	Color string
}

func GenerateRandomSquares(numSquares int) []*Square {
	squares := make([]*Square, numSquares)
	for i := 0; i < numSquares; i++ {
		x := rand.Int31n(maxSquarePositionx/gridSize) * gridSize
		y := rand.Int31n(maxSquarePositiony/gridSize) * gridSize
		square := &Square{Rect: &sdl.Rect{X: x, Y: y, W: gridSize, H: gridSize}}
		squares[i] = square
		fmt.Printf("Kwadrat %d: X=%d, Y=%d, W=%d, H=%d\n", i+1, square.Rect.X, square.Rect.Y, square.Rect.W, square.Rect.H)

	}
	return squares
}

func PopulateGrid(grid [][]string, squares []*Square, startPoint *Point, endPoint *Point) [][]string {
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = "empty"
		}
	}

	for _, square := range squares {
		grid[square.Rect.Y/gridSize][square.Rect.X/gridSize] = "white"
	}

	if startPoint != nil {
		grid[startPoint.Y/gridSize][startPoint.X/gridSize] = "startPoint"
	}

	if endPoint != nil {
		grid[endPoint.Y/gridSize][endPoint.X/gridSize] = "endPoint"
	}

	// Wypisz tablicę grid w terminalu
	fmt.Println("Grid:")
	for _, row := range grid {
		fmt.Println(row)
	}

	return grid
}

func FindShortestPath(grid *[][]string) {
	rows := len((*grid))
	if rows == 0 {
		fmt.Println("Empty grid")
		return
	}

	cols := len((*grid)[0])
	if cols == 0 {
		fmt.Println("Empty grid")
		return
	}

	startX, startY := -1, -1
	endX, endY := -1, -1

	// Znalezienie współrzędnych startPoint i endPoint w tablicy grid
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (*grid)[i][j] == "startPoint" {
				startX, startY = i, j
			} else if (*grid)[i][j] == "endPoint" {
				endX, endY = i, j
			}
		}
	}

	if startX == -1 || startY == -1 {
		fmt.Println("Start point not found")
		return
	}

	if endX == -1 || endY == -1 {
		fmt.Println("End point not found")
		return
	}

	// Utworzenie tablicy odległości i inicjalizacja wartości maksymalnych
	distances := make([][]float64, rows)
	for i := range distances {
		distances[i] = make([]float64, cols)
		for j := range distances[i] {
			distances[i][j] = math.MaxFloat64
		}
	}

	// Ustawienie odległości dla startPoint na 0
	distances[startX][startY] = 0

	// Kolejka do przetwarzania punktów
	queue := []int{startX, startY}

	// Tablica przechowująca poprzednie punkty na najkrótszej ścieżce
	previous := make([][]int, rows)
	for i := range previous {
		previous[i] = make([]int, cols)
	}

	// Przetwarzanie punktów w kolejce
	for len(queue) > 0 {
		x := queue[0]
		y := queue[1]
		queue = queue[2:]

		// Sprawdzenie sąsiednich punktów
		neighbors := [][]int{
			{x - 1, y}, // góra
			{x + 1, y}, // dół
			{x, y - 1}, // lewo
			{x, y + 1}, // prawo
		}

		for _, neighbor := range neighbors {
			nx := neighbor[0]
			ny := neighbor[1]

			if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
				if (*grid)[nx][ny] != "white" && distances[nx][ny] == math.MaxFloat64 {
					distances[nx][ny] = distances[x][y] + 1
					previous[nx][ny] = x*cols + y
					queue = append(queue, nx, ny)
				}
			}
		}
	}

	// Tworzenie najkrótszej ścieżki od endPoint do startPoint
	path := []int{}
	x := endX
	y := endY

	for x != startX || y != startY {
		path = append(path, x*cols+y)
		prevX := previous[x][y] / cols
		prevY := previous[x][y] % cols
		x = prevX
		y = prevY
	}

	// Oznaczanie ścieżki na tablicy grid
	for _, p := range path {
		px := p / cols
		py := p % cols
		if (*grid)[px][py] != "startPoint" && (*grid)[px][py] != "endPoint" && (*grid)[px][py] != "white" {
			(*grid)[px][py] = "path"
		}
	}

	// Wyświetlanie zmodyfikowanej tablicy grid w terminalu
	fmt.Println("Droga: ")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%s ", (*grid)[i][j])
		}
		fmt.Println()
	}
}

func DrawPathCells(renderer *sdl.Renderer, grid [][]string) {
	fmt.Println("start")
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "path" {
				gridX := j * gridSize
				gridY := i * gridSize
				rect := &sdl.Rect{
					X: int32(gridX),
					Y: int32(gridY),
					W: gridSize,
					H: gridSize,
				}
				renderer.FillRect(rect)
				fmt.Println("draw")
			}
		}
	}
	fmt.Println("end")
}
