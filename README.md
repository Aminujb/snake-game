Commandline Snake Game
========================

Commandline Snake Game built with Golang. Took about a whole day to complete. Had breaks inbetween progress made. 

## Assumptions

1. Head of snake is denoted by integer 2 on the board.
1. Trail of snake is denoted by integer 1 on the board.
1. Body of the snake (length) is denoted by -1.

1. Movement of snake head is controlled with characters r(right)--l(left)--u(up)--d(down))
1. Snake food is randomly placed in coordinates with 0 value denoted by 5
1. Snake "eats" food when movement of head intersects with coordinates of food.

1. Game continues until snake head intersects with -1 (body of snake)

## Run program
go run ./cmd/snake_game.go  
