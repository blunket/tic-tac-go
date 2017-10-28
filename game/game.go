package game

import (
	"github.com/blunket/tic-tac-go/game/grid"
)

type Game struct {
	Board grid.Grid
	Turn  int
}

func New(turn int) Game {
	return Game{
		Turn: turn,
	}
}

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
