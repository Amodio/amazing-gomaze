// File maze.go
// © 2012 Jacques Boscq <jacques@boscq.fr>. Under GPL3, see COPYING.

// Package gomaze implements methods for manipulating a maze
package gomaze

import (
    "fmt"
    "math/rand"
    "strings"
)

type Maze struct {
    // dimension
    width, height int
    // cells
    cells [][]Cell
}

type Cell struct {
    // coordinates
    x, y int
    // neighbours
    up, down, left, right *Cell
}

type Tarjan struct {
    // remaining cells to compute
    RemainingCells []*Cell
    // id of the tarjan's set the cells belong to
    cells [][]int
}

const DimensionMax = 50

// Return a new Maze (partially) initialized
func NewMaze(height, width int) *Maze {
    if width <= 0 || height <= 0 || width > DimensionMax || height > DimensionMax {
        panic("Invalid dimensions for the Maze!")
    }

    // init. the cells
    cells := make([][]Cell, height)
    for i := 0; i < height; i++ {
        cells[i] = make([]Cell, width)
        for j := 0; j < width; j++ {
            // set the (x, y) coordinates for each cell
            cells[i][j].x, cells[i][j].y = i, j
        }
    }

    return &Maze{
        width,
        height,
        cells,
    }
}

// Return a new SquaredMaze (partially) initialized
func NewSquaredMaze(n int) *Maze {
    return NewMaze(n, n)
}

// Return a neighbour cell not connected to c
//        nil if every neighbour of c is reachable
func (b *Maze) mergeableCell(c *Cell, t *Tarjan) *Cell {
    dir := [4][2]int{{-1, 0}, // up
        {1, 0},  // down
        {0, -1}, // left
        {0, 1},  // right
    }
    // the directions are mixed randomly
    shuffledDir := rand.Perm(4)
    // for each randomly mixed direction
    for _, v := range shuffledDir {
        // if the targeted cell is reachable, ignore it
        if v == 0 && c.up != nil {
            continue
        } else if v == 1 && c.down != nil {
            continue
        } else if v == 2 && c.left != nil {
            continue
        } else if v == 3 && c.right != nil {
            continue
        }
        // coordinates of the targeted neighbour
        nX, nY := c.x+dir[v][0], c.y+dir[v][1]
        // boundaries check
        if nX >= 0 && nY >= 0 && nX < b.height && nY < b.width {
            // are they from the same tarjan's set ?
            if t.cells[c.x][c.y] != t.cells[nX][nY] {
                // no ? return the neighbour cell to connect to
                return &b.cells[nX][nY]
            }
        }
    }
    return nil
}

// Convert cells from a tarjan's set to another
func (b *Maze) updateTarjan(from, to int, t *Tarjan) {
    // go through each cell
    for i := 0; i < b.height; i++ {
        for j := 0; j < b.width; j++ {
            if t.cells[i][j] == from {
                // convert the selected cell
                t.cells[i][j] = to
                // remove the cell from the remaining cells to compute
                b.removeCell(&b.cells[i][j], t)
            }
        }
    }
}

// Remove a cell from the remaining cells to compute
func (b *Maze) removeCell(c *Cell, t *Tarjan) {
    var tmp []*Cell
    for _, v := range t.RemainingCells {
        if v != c {
            tmp = append(tmp, v)
        }
    }
    t.RemainingCells = tmp
}

// Unset the wall between the two cells
func (b *Maze) unsetWalls(c1, c2 *Cell) {
    dir := [4][2]int{{-1, 0}, // up
        {1, 0},  // down
        {0, -1}, // left
        {0, 1},  // right
    }
    var direction int
    // determine the direction to go from c1 to c2
    for k, v := range dir {
        if c1.x+v[0] == c2.x && c1.y+v[1] == c2.y {
            direction = k
            break
        }
    }
    // link the two cells
    switch direction {
    case 0: // up
        c1.up, c2.down = c2, c1
    case 1: // down
        c1.down, c2.up = c2, c1
    case 2: // left
        c1.left, c2.right = c2, c1
    case 3: // right
        c1.right, c1.left = c2, c1
    }
}

// Merges two sets of tarjan, the id the lower will survive
func (b *Maze) mergeEm(c1, c2 *Cell, t *Tarjan) {
    x, y := t.cells[c1.x][c1.y], t.cells[c2.x][c2.y]
    var master, slave int
    var firstCell, secCell *Cell
    // select the cell given its tarjan's set ID
    if x < y {
        master, slave = x, y
        firstCell, secCell = c1, c2
    } else {
        master, slave = y, x
        firstCell, secCell = c2, c1
    }
    // link the cells
    b.unsetWalls(firstCell, secCell)
    b.updateTarjan(master, slave, t)
}

// Create a new Tarjan structure
func (b *Maze) newTarjan() *Tarjan {
    t := new(Tarjan)
    t.cells = make([][]int, b.height)
    for i := 0; i < b.height; i++ {
        t.cells[i] = make([]int, b.width)
        for j := 0; j < b.width; j++ {
            // for each cell this is the first tarjan's set id
            t.cells[i][j] = i*b.width + j
            t.RemainingCells = append(t.RemainingCells, &b.cells[i][j])
        }
    }
    return t
}

// Generate randomly a perfect maze
func (b *Maze) Generate() {
    t := b.newTarjan()
    // while the sets of tarjan are > 1
    for len(t.RemainingCells) > 1 {
        // c is a randomly chosen cell, remaining to compute
        c := t.RemainingCells[rand.Intn(len(t.RemainingCells))]
        if neighbour := b.mergeableCell(c, t); neighbour != nil {
            // merge the mergeable cells! (a currently disconnected neighbour)
            b.mergeEm(c, neighbour, t)
        }
    }
}

// Returns a string representation of a Maze
func (b *Maze) String() string {
    s := fmt.Sprintf("Printing a Maze of %dx%d cells.\n", b.height, b.width)
    s += " " + strings.Repeat("_", b.width*2-1) + "\n"
    for i := 0; i < b.height; i++ {
        s += "|"
        for j := 0; j < b.width; j++ {
            if b.cells[i][j].down != nil || i == b.height-1 {
                s += " "
            } else {
                s += "_"
            }
            if b.cells[i][j].right == nil || j == b.width-1 {
                s += "|"
            } else {
                s += " "
            }
        }
        s += "\n"
    }
    s += " " + strings.Repeat("‾", b.width*2-1)
    return s
}
