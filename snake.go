package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"golang.org/x/term"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

const (
	width  = 20
	height = 10
)

type Position struct {
	x, y int
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	if _, err := exec.LookPath("clear"); err == nil {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func draw(snake []Position, food Position, score int) {
	clearScreen()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			isSnake := false
			for _, s := range snake {
				if s.x == x && s.y == y {
					fmt.Print("O")
					isSnake = true
					break
				}
			}
			if !isSnake {
				if food.x == x && food.y == y {
					fmt.Print("X")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
	fmt.Printf("Score: %d\n", score)
	fmt.Println("Controls: W=Up, S=Down, A=Left, D=Right, Q=Quit")
}

func generateFood(snake []Position) Position {
	for {
		food := Position{rand.Intn(width), rand.Intn(height)}
		collision := false
		for _, s := range snake {
			if s == food {
				collision = true
				break
			}
		}
		if !collision {
			return food
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	snake := []Position{{5, 5}}
	direction := RIGHT
	food := generateFood(snake)
	score := 0
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	input := make(chan rune)
	go func() {
		for {
			b := make([]byte, 1)
			_, err := os.Stdin.Read(b)
			if err == nil {
				input <- rune(b[0])
			}
		}
	}()
	for {
		select {
		case char := <-input:
			switch char {
			case 'w', 'W':
				if direction != DOWN {
					direction = UP
				}
			case 's', 'S':
				if direction != UP {
					direction = DOWN
				}
			case 'a', 'A':
				if direction != RIGHT {
					direction = LEFT
				}
			case 'd', 'D':
				if direction != LEFT {
					direction = RIGHT
				}
			case 'q', 'Q':
				clearScreen()
				fmt.Println("Game Over! Final Score:", score)
				return
			}
		default:
			head := snake[0]
			switch direction {
			case UP:
				head.y--
			case DOWN:
				head.y++
			case LEFT:
				head.x--
			case RIGHT:
				head.x++
			}
			if head.x < 0 {
				head.x = width - 1
			}
			if head.x >= width {
				head.x = 0
			}
			if head.y < 0 {
				head.y = height - 1
			}
			if head.y >= height {
				head.y = 0
			}
			for _, s := range snake {
				if s == head {
					clearScreen()
					fmt.Println("Game Over! You ate yourself!")
					fmt.Printf("Final Score: %d\n", score)
					return
				}
			}
			snake = append([]Position{head}, snake...)
			if head == food {
				score++
				food = generateFood(snake)
			} else {
				snake = snake[:len(snake)-1]
			}
			draw(snake, food, score)
			time.Sleep(150 * time.Millisecond)
		}
	}
}
