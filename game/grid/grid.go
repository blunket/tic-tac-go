// Package grid provides a Tic-Tac-Toe grid with which to interact.
package grid

type Grid [9]int

// Place sets a square at x,y on a grid to p if the square is open and returns whether successful.
func (g *Grid) Place(p, x, y int) bool {
	if g.squareAt(x, y) == 0 {
		g[x+y*3] = p
		return true
	}
	return false
}

// Slice simply returns the grid as a [9]int.
func (g Grid) Slice() [9]int {
	return [9]int(g)
}

// State returns whether the grid is in a win state, and the winning player.
func (g *Grid) State() (bool, int) {
	if checkLine(g.diag()) {
		return true, g.squareAt(0, 0)
	}

	if checkLine(g.adiag()) {
		return true, g.squareAt(2, 0)
	}

	for i := 0; i < 3; i++ {
		if checkLine(g.row(i)) {
			return true, g.squareAt(0, i)
		}
		if checkLine(g.col(i)) {
			return true, g.squareAt(i, 0)
		}
	}

	return false, 0
}

// checkLine returns whether all three items of a [3]int are equal but not 0.
func checkLine(l [3]int) bool {
	return l[0] != 0 && l[0] == l[1] && l[1] == l[2]
}

// squareAt returns the value of a square at the given x,y position.
func (g *Grid) squareAt(x, y int) int {
	return g[x+y*3]
}

// row returns a [3]int containing the values of the squares in any given row.
func (g *Grid) row(r int) [3]int {
	return [3]int{
		g.squareAt(0, r),
		g.squareAt(1, r),
		g.squareAt(2, r),
	}
}

// col returns a [3]int containing the values of the squares in any given column.
func (g *Grid) col(c int) [3]int {
	return [3]int{
		g.squareAt(c, 0),
		g.squareAt(c, 1),
		g.squareAt(c, 2),
	}
}

// diag returns a [3]int containing the values of the squares in the main diagonal.
func (g *Grid) diag() [3]int {
	return [3]int{
		g.squareAt(0, 0),
		g.squareAt(1, 1),
		g.squareAt(2, 2),
	}
}

// adiag returns a [3]int containing the values of the squares in the antidiagonal.
func (g *Grid) adiag() [3]int {
	return [3]int{
		g.squareAt(2, 0),
		g.squareAt(1, 1),
		g.squareAt(0, 2),
	}
}
