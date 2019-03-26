package snakes_ladders

import(
	"math/rand"
	"time"
	"fmt"
)
const MAX_PLAYERS = 10

type Board struct {
	tiles []int
	capacity int
	players []*Player
}

func EmptyBoard(size int) *Board {
	tiles := make([]int, size)
	for i := range tiles {
		tiles[i] = -1
	}
	return &Board{tiles,size,[]*Player{}}
}

func (this *Board) createPlayers(num int) []*Player {
	players := make([]*Player, num, MAX_PLAYERS)
	for i := range players {
		players[i] = newPlayer(fmt.Sprintf("Player %d", i), 0, this)
	}
	return players
}

func NewBoard(size, players int, percentage float64) *Board {
	tiles := make([]int, size)
	for i := range tiles {
		tiles[i] = randomPathLength(percentage, i, size/10, size/2, size)
	}
	board := &Board{tiles,size,nil}
	board.createPlayers(players)
	return board
}

func (this Board) tileHasPath(index int) bool {
	return this.tiles[index] != -1
}

func (this *Board) addPlayer(player *Player) {
	player.board = this
	this.players = append(this.players, player)
}

func (this *Board) clearPlayers() {
	this.players = []*Player{}
}

func (this Board) numPlayers() int {
	return len(this.players)
}

func (this *Board) customTiles(tiles []int) {
	this.tiles = tiles
}

func randomSign(r *rand.Rand) int {
	if direction := int(r.Float64()+.5); direction == 0 {
		return -1
	} 
	return 1
}

func randomPathLength(percentage float64, start, min, max, size int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Float64() < percentage {
		// create a path
		direction := randomSign(r)
		length := min + r.Intn(max-min)
		destination := start + (length * direction)
		if destination >= 0 && destination < size {
			return destination
		}
	}
	// no path
	return -1
}

func (this Board) Print() {
	for i := range this.tiles {
		if this.tileHasPath(i) {
			fmt.Printf("**")
		}
		fmt.Printf("%d -> ", this.tiles[i])
	}
	fmt.Printf("END\n")
}
