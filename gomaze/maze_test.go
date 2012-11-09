// File maze/maze_test.go
// © 2012 Jacques Boscq <jacques@boscq.fr>. Under GPL3, see COPYING.
// Creation date 09 nov. 2012 01:25 +0100

package gomaze

import "testing"

func TestNewMaze(t *testing.T) {
    invalid_dimensions := [...]int{ DimensionMax + 1, 0, -1, -1337 }
    for _, d := range invalid_dimensions {
        _, err := NewMaze(d, d)
        if err == nil {
            t.Fatalf("NewMaze() accepts invalid dimension: %v", d)
        }
    }
}

func TestNewSquaredMaze(t *testing.T) {
    invalid_dimensions := [...]int{ DimensionMax + 1, 0, -1, -1337 }
    for _, d := range invalid_dimensions {
        _, err := NewSquaredMaze(d)
        if err == nil {
            t.Fatalf("NewSquaredMaze() accepts invalid dimension: %v", d)
        }
    }
}

func TestNewTarjan(t *testing.T) {
    const n = 20
    m, err := NewSquaredMaze(n)
    if err != nil {
        t.Fatalf("NewSquaredMaze() failed: %v", err)
    }
    tarjan := m.newTarjan()
    if tarjan == nil {
        t.Fatalf("newTarjan() failed")
    }
    if len(tarjan.RemainingCells) != n * n {
        t.Fatalf("newTarjan() failed: %d != %d", len(tarjan.RemainingCells), n)
    }
}

func BenchmarkGenerateMaze(b *testing.B) {
    if b.N > DimensionMax {
        b.N = DimensionMax
    }
    m, err := NewSquaredMaze(b.N)
    if err != nil {
        b.Fatalf("NewSquaredMaze() failed: %v", err)
    }
    m.Generate()
}
/*
 * Remaining functions to test:

func (b *Maze) mergeableCell(c *Cell, t *Tarjan) *Cell {
func (b *Maze) updateTarjan(from, to int, t *Tarjan) {
func (b *Maze) removeCell(c *Cell, t *Tarjan) {
func (b *Maze) unsetWalls(c1, c2 *Cell) {
func (b *Maze) mergeEm(c1, c2 *Cell, t *Tarjan) {
func (b *Maze) String() string {
*/
