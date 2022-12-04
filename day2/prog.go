package main

import (
    "fmt"
    "io"
)

/*
   YOU PLAY     X   Y   Z

   OPPONENT A   1+3 2+6 3+0

            B   1+0 2+3 3+6

            C   1+6 2+0 3+3
            
 */
/*
   OUTCOME      X   Y   Z

   OPPONENT A   0+3 3+1 6+2

            B   0+1 3+2 6+3

            C   0+2 3+3 6+1
            
 */

func opponent (move string) int {
    switch move {
        case "A": return 0
        case "B": return 1
        case "C": return 2
    }
    return -1
}

func you (move string) int {
    switch move {
        case "X": return 0
        case "Y": return 1
        case "Z": return 2
    }
    return -1
}

var lookup1 = [3][3]int {
    {1+3, 2+6, 3+0},
    {1+0, 2+3, 3+6},
    {1+6, 2+0, 3+3}}

var lookup2 = [3][3]int {
    {0+3, 3+1, 6+2},
    {0+1, 3+2, 6+3},
    {0+2, 3+3, 6+1}}



func main () {

    var n int
    var err error
    w1, w2 := "", ""
    points1, points2 := 0, 0
    for n, err = fmt.Scanln(&w1, &w2); n>0 && err != io.EOF; n, err = fmt.Scanln(&w1, &w2) {
        points1 += lookup1[opponent(w1)][you(w2)]
        points2 += lookup2[opponent(w1)][you(w2)]
    }
    fmt.Println(points1, points2)
    return
}
