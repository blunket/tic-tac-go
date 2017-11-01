// Package game provides a Tic-Tac-Toe game controller.
package game

import (
	"github.com/blunket/tic-tac-go/game/grid"
	"sync"
)

// Game is a Tic-Tac-Toe game controller.
type Game struct {
	board grid.Grid
	turn  int // Turn identifies whose turn it is. 1 or 2.
	sync.Mutex
}

// New creates a new Game, and the provided player goes first.
func New(p int) Game {
	return Game{
		Turn: p,
	}
}

// Board returns the board state as a slice.
func (g *Game) Board() [9]int {
	return g.board.Slice()
}

// Place sets a square at x,y on a grid to p if the square is open.
// It must be the player's turn.
// Returns whether successful.
func (g *Game) Place(p, x, y int) bool {
	g.Lock()
	defer g.Unlock()

	if g.Turn != p {
		return false
	}

	sq := g.board.Place(p, x, y)
	if sq == false {
		return false
	}

	if g.Turn == 1 {
		g.Turn = 2
	} else {
		g.Turn = 1
	}
	return sq
}
