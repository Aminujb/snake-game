package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)


var row, column int
var GameOver = errors.New("Oops Game over")

func init(){
	fmt.Print("Enter board size\n")

    fmt.Print("Enter height: ")
	_, err:= fmt.Scanln(&row)
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(0)
	}

    fmt.Print("Enter width: ")
	_, err= fmt.Scanln(&column)
	if err != nil{
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

type SnakeGame struct{
	Board [][]int
	SnakeHead [2]int
	FoodPosition [2]int
	NumberOfFoodEaten int
	LengthOfSnake int
	Trail [][]int
}

func main() {
	sg := SnakeGame{}

	// set up board
	sg.Board = setupBoard(row, column)
	sg.placeFood()
	sg.SnakeHead[0] = 0
	sg.SnakeHead[1] = 0
	
	sg.printBoard()

	for {
		var direction string
		fmt.Print("Move snake (Options-> r(right)--l(left)--u(up)--d(down)): ")
		fmt.Scanln(&direction)
		
		err := sg.move(direction)
		if err != nil{
			fmt.Println(err.Error())
			fmt.Println("No of food eaten: ", sg.NumberOfFoodEaten)
			fmt.Println("Snake length: ", sg.LengthOfSnake)
			os.Exit(0)
		}

		sg.printBoard()
		
		if sg.SnakeHead[0] == sg.FoodPosition[0] &&  sg.SnakeHead[1] == sg.FoodPosition[1]{
			sg.NumberOfFoodEaten++
			fmt.Println("Success! You got the meal!!! No of food eaten: ", sg.NumberOfFoodEaten)

			sg.cleanTrail()
			sg.placeFood()
			sg.printBoard()
		}

	}
}

func (sg *SnakeGame) placeFood(){
	// get possible positions first
	possibleXY := []string{}

	for x := range sg.Board{
		for y := range sg.Board[x]{
			if sg.Board[x][y] == 0{
				possibleXY = append(possibleXY, fmt.Sprintf("%d,%d", x, y))
			}
		}
	}

	// get random position
	rand.Seed(time.Now().UnixNano())
	lowerBoundry := 0
	higherBoundry:= len(possibleXY) - 1
	randomPosition := lowerBoundry + rand.Intn(higherBoundry-lowerBoundry+1)

	value := strings.Split(possibleXY[randomPosition], ",")
	pX, _ := strconv.Atoi(value[0])
	pY, _ := strconv.Atoi(value[1])

	// 5 represents food 
	sg.FoodPosition[0] = pX
	sg.FoodPosition[1] = pY
	sg.Board[pX][pY] = 5
}

func setupBoard(row, column int) [][]int{
	a := make([][]int, row)
	for i := range a {
		a[i] = make([]int, column)
	}
	a[0][0] = 2

	return a
}

func (sg SnakeGame) printBoard(){
	for x := range sg.Board{
		for y := range sg.Board[x]{
			fmt.Print(sg.Board[x][y])
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Printf("Coordinate of snake head: (%d,%d)\n", sg.SnakeHead[0], sg.SnakeHead[1])
}

func (sg *SnakeGame) move(direction string) error{
	if direction != "r" && direction != "l" && direction != "u" && direction !="d" {
		return nil
	}

	//prev head cordinate
	prevX := sg.SnakeHead[0]
	prevY := sg.SnakeHead[1]

	//old head cleared
	sg.Board[sg.SnakeHead[0]][sg.SnakeHead[1]] = 0

	if direction == "r"{
		if sg.SnakeHead[1] + 1 < column{
			oldY := sg.SnakeHead[1]
			sg.SnakeHead[1] = oldY + 1
		}
	}

	if direction == "l"{
		if sg.SnakeHead[1] - 1 > -1{
			oldY := sg.SnakeHead[1]
			sg.SnakeHead[1] = oldY - 1
		}
	}

	if direction == "d"{
		if sg.SnakeHead[0] + 1 < row{
			oldX := sg.SnakeHead[0]
			sg.SnakeHead[0] = oldX + 1
		}
	}

	if direction == "u"{
		if sg.SnakeHead[0] - 1 > -1{
			oldX := sg.SnakeHead[0]
			sg.SnakeHead[0] = oldX - 1
		}
	}

	x := sg.SnakeHead[0]
	y := sg.SnakeHead[1]
	
	// check if the value is -1
	if sg.Board[x][y] == -1 || sg.Board[x][y] == 1 {
		return GameOver
	}
	
	// new head on board
	sg.Board[x][y] = 2

	// make previous steps with 1
	sg.Board[prevX][prevY] = 1
	sg.Trail = append(sg.Trail, []int{prevX,prevY})

	return nil
}

func (sg *SnakeGame) cleanTrail(){
	sg.LengthOfSnake += len(sg.Trail)
	for cord := range sg.Trail{
		xy := sg.Trail[cord]
		sg.Board[xy[0]][xy[1]] = -1
	}

	sg.Trail = [][]int{}
}
