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
	path int  
	expected int
}

func TestPlayerMovement(t *testing.T) {
	board := EmptyBoard(10)
	cases := []PlayerMovementTest {
		{0,1,-1,1},
		{0,2,0,0},
		{0,3,7,7},
		{7,6,-1,7},
		{7,3,-1,10},
	}
	board.customTiles([]int{-1,-1,0,7,-1,-1,-1,-1,-1,-1})
	for _, c := range cases {
		p := newPlayer("test",c.start, nil)
		board.addPlayer(p)
		p.takeTurn(c.roll)
		fmt.Println(p.position)
		if p.position != c.expected {
			path := -1
			if c.start + c.roll < 10 {
				path = board.tiles[c.start + c.roll]
			} 
			t.Errorf("Player that started on position %d, rolled a %d, and took path %d, landed on position %d.  Expected position %d", 
				c.start, c.roll, path, p.position, c.expected)
		}

	}
}

