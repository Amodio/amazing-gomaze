// File amazing-gomaze.go
// © 2012 Jacques Boscq <jacques@boscq.fr>. Under GPL3, see COPYING.

package main

import (
    "flag"
    "fmt"
    "github.com/Amodio/amazing-gomaze/gomaze"
    "math/rand"
    "time"
)

var dimension = flag.Int("d", 20, "dimension of the Maze")

func main() {
    flag.Parse()
    rand.Seed(time.Now().UnixNano())

    if b, err := gomaze.NewSquaredMaze(*dimension); err == nil {
        b.Generate()
        fmt.Println(b)
    }

    /*
       d := NewMaze(9, 50)
       d.Generate()
       fmt.Println(d)
    */
}
