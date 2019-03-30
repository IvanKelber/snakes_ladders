package snakes_ladders

import (
	"testing"
	"fmt"
)

type TileHasPathTest struct {
    path int
    expected bool
}

func TestTileHasPath(t *testing.T) {
    board := EmptyBoard(10)
	cases := []TileHasPathTest {
        {-1, false},
        {1, true}, 
    }
	for _, c := range cases {
        board.tiles[0] = c.path
		output := board.tileHasPath(0)
		if output != c.expected {
			t.Errorf("{tiles[0] = %d } tileHasPath(0) == %t, expected %t",  c.path, output, c.expected)
		}
	}
}

type PlayerMovementTest struct {
	start int
	roll int
	expected int
}

func TestPlayerMovement(t *testing.T) {
	//Setup
	game := EmptyGame();
	game.board = EmptyBoard(10)
	game.board.customTiles([]int{-1,-1,0,7,-1,-1,-1,-1,-1,-1})

	cases := []PlayerMovementTest {
		{0,1,1}, // no path
		{0,2,0}, // snake back to 0
		{0,3,7}, // ladder up to 7
		{7,6,7}, // overshoot
		{7,3,10}, // win the game
	}
	for _, c := range cases {
		p := game.addPlayer("test", c.start)
		p.takeTurn(c.roll)
		fmt.Println(p.position)
		if p.position != c.expected {
			path := game.board.getPath(c.start + c.roll)
			t.Errorf("Player that started on position %d, rolled a %d, and took path %d, landed on position %d.  Expected position %d", 
				c.start, c.roll, path, p.position, c.expected)
		}

	}
}

