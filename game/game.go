// Package game provides a Tic-Tac-Toe game controller.
package game

import (
	"github.com/blunket/tic-tac-go/game/grid"
)

type Game struct {
	Board grid.Grid
	Turn  int
}

// New creates a new Game, and the provided player goes first.
func New(p int) Game {
	return Game{
		Turn: p,
	}
}

// Place sets a square at x,y on a grid to p if the square is open.
// It must be the player's turn.
// Returns whether successful.
func (g *Game) Place(p, x, y int) bool {
	if g.Turn != p {
		return false
	}

	sq := g.Board.Place(p, x, y)
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
