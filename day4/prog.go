package main

import (
    "fmt"
    "io"
)

type Task struct {
    first int
    last int
}

func (this Task) overlaps (other Task) bool {
    return !(((this.first > other.first) || (this.last < other.first)) &&
        ((other.first > this.first) || (other.last < this.first)))
}

func (this Task) contains (other Task) bool {
    return !((this.first > other.first) || (this.last < other.last))
}

func parse (tasks []Task, buffer string) {
    acc := 0
    ind := 0
    for _, char := range buffer {
        switch char {
            case '-':
                tasks[ind].first = acc
                acc = 0
                ind ++
            case ',':
                tasks[0].last = acc
                acc = 0
            default:
                acc = (acc*10) + int(char) - int('0')
        }
    }
    tasks[1].last = acc
    return
}

func main() {
    buffer := ""
    acc1 := 0
    acc2 := 0
    var someTask Task
    var pair = []Task {someTask, someTask}
    for n, err := fmt.Scanln(&buffer); n > 0 && err != io.EOF; n, err = fmt.Scanln(&buffer) {
        parse(pair, buffer)
        if pair[0].overlaps(pair[1]) {
            if pair[0].contains(pair[1]) || pair[1].contains(pair[0]) {
                acc1++
            }
            acc2++
        }
    }
    fmt.Println(acc1,acc2)
}
