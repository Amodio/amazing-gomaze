// File main.go
// © 2012 Jacques Boscq <jacques@boscq.fr>. Under GPL3, see COPYING.

package main

import (
    "flag"
    "fmt"
    maze "github.com/Amodio/amazing-gomaze/maze"
    "math/rand"
    "time"
)

var dimension = flag.Int("d", 20, "dimension of the Maze")

func main() {
    flag.Parse()
    rand.Seed(time.Now().UnixNano())

    b := maze.NewSquaredMaze(*dimension)
    b.Generate()
    fmt.Println(b)

    /*
       d := NewMaze(9, 50)
       d.Generate()
       fmt.Println(d)
    */
}
