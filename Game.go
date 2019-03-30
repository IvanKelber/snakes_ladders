package snakes_ladders

// Game.go should be the public interface for users.
// A Game object should have a list of players, a board and a few other attributes

//TODO
import(
	"fmt"
)

type Game struct {
	board *Board
	players []*Player
}

func EmptyGame() *Game {
	return &Game{nil, nil}
}

func NewGame(size, players int, percentage float64) *Game{
	game := EmptyGame();
	game.board = NewBoard(size, percentage)
	game.createPlayers(players)
	return game
}

func (this *Game) createPlayers(num int) []*Player {
	this.players = make([]*Player, num, MAX_PLAYERS)
	for i := range this.players {
		this.players[i] = newPlayer(fmt.Sprintf("Player %d", i), 0, this.board)
	}
	return this.players
}

func (this *Game) addPlayer(name string, pos int) *Player {
	player := newPlayer(name,pos, this.board);
	this.players = append(this.players, player)
	return player
}

func (this *Game) clearPlayers() {
	this.players = []*Player{}
}

func (this Game) numPlayers() int {
	return len(this.players)
}
