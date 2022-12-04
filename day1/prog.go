package main // Uh? dunno what this is for

import (    // Wtf? i'm importing it as a string?
    "fmt"
    "strconv"
    "io"
)

func swap (vec []int, i int, j int) {
    tmp := vec[i]
    vec[i] = vec[j]
    vec[j] = tmp
    return
}

func insert (vec []int, val int) {
    if val > vec[0] {
        vec[0] = val
    }
    for i := 0; i < 3; i++ {
        if vec[i] > vec[i+1] {
            swap(vec, i, i+1)
        }
    }
    return
}

func main() {
    
    var max = []int{0,0,0,0}  // zero-initialized vector

    var w string
    n := 1
    var err error
    for err != io.EOF {
        acc := 0
        for n, err = fmt.Scanln(&w); n>0; n, err = fmt.Scanln(&w) {
            num, _ := strconv.Atoi(w)
            acc += num
            w = ""
        }
        insert(max, acc)
    }
    fmt.Println(max[3], max[1] + max[2] + max[3])
    return
}
