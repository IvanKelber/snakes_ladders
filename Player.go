package snakes_ladders

import(
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
    name string
	position int
	board *Board
}

func newPlayer(name string, position int, board *Board) *Player{
	return &Player{name, position, board}
}


func (this Player) rollDie() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

// Public Method that the user takes
func (this *Player) TakeTurn() {
	this.takeTurn(this.rollDie())
}

// Private method that can be tested
func (this *Player) takeTurn(roll int) {
	stop := this.position + roll;
	fmt.Printf("%s rolls a %d ", this.name, roll)

	if stop > this.board.capacity{
		// Player must land exactly on the "last spot"
		fmt.Printf("but overshoots the landing by %d tiles!  They stay where they are!\n", stop - this.board.capacity)
		stop = this.position
	} else if stop == this.board.capacity {
		fmt.Printf("and lands on the last tile!  %s wins!\n", this.name)
	} else if this.board.tileHasPath(stop) {
		if this.board.tiles[stop] > stop {
			fmt.Printf("and takes a ladder from %d to %d! Lucky!\n", stop, this.board.tiles[stop])
		} else {
			fmt.Printf("and takes a snake from %d to %d! Shoot!\n", stop, this.board.tiles[stop])
		}
		stop = this.board.tiles[stop]
	} else {
		fmt.Printf("and lands on %d\n", stop)
	}
	this.updatePosition(stop)
	
}

func (this *Player) updatePosition(newPosition int) {
	this.position = newPosition
}
