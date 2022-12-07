package main

import (
        "bufio"
        "fmt"
        "io"
        "os"
        "unicode"
       )

type Stack struct {
    maxSize int
    actSize int
    data []byte
}

func initStack (maxSize int) Stack {
    return Stack {maxSize, 0, make([]byte, maxSize) }
}

func (stack *Stack) Push (value byte) {
    stack.data[stack.actSize] = value
    stack.actSize++
    return
}

func (stack *Stack) PushMulti(values []byte) {
    n := len(values)
        for i := 0; i < n; i++ {
        stack.data[stack.actSize+i] = values[i]
    }
    stack.actSize += n
        return
}

func (stack *Stack) Pop() (buf byte) {
    stack.actSize--
    buf = stack.data[stack.actSize]
    return
}

func (stack *Stack) PopMulti(n int) (buf []byte) {
        stack.actSize -= n
    buf = make([]byte, n)
    for i := 0; i<n; i++ {
        buf[i] = stack.data[stack.actSize+i]
    }
        return
}

func initCargo(stacks []Stack, str []string) {
    w := len(str[0])
    h := len(str)-1
    for j := 0; j < h; j++ {
        for i, I := 1, 0; i < w; i, I = i+4, I+1 {
            val := str[j][i]
            stacks[I].data[h-j-1] = val
            if stacks[I].actSize == 0 && val != byte(' ') {
                stacks[I].actSize = h-j
            }
        }
    }
}



func parseLine(str string) (qnt int, src int, dst int) {
    var val [3]int
    t := 0
    acc := 0
    for _, c := range str {
        if c == ' ' || c == '\n' {
            if acc != 0 {
                val[t] = acc
                t++
                acc = 0
            }
        } else if unicode.IsDigit(c) {
            acc = acc * 10 + int(c) - int('0')
        }
    }
    qnt = val[0]
    src = val[1] - 1
    dst = val[2] - 1
    return
}

func move1(stacks []Stack, qnt int, src int, dst int) {
    for i := 0; i < qnt; i++ {
        container := stacks[src].Pop()
        stacks[dst].Push(container)
    }
}

func move2(stacks []Stack, qnt int, src int, dst int) {
    container := stacks[src].PopMulti(qnt)
        stacks[dst].PushMulti(container)
}

func read(stacks []Stack) (code string) {
    code = ""
    n := len(stacks)
    for i := 0; i < n; i++ {
        container := stacks[i].Pop()
        code += string(container)
    }
    return
}

func main () {
    
    reader := bufio.NewReader(os.Stdin)

    // Parsing stacks
    stackSize := 0
    var matrix []string
    for text, err := reader.ReadString('\n'); text != "\n" && err != io.EOF; text, err = reader.ReadString('\n') {
        matrix = append(matrix, text)
        stackSize++
    }
    stackSize--
    var numOfStacks int = (len(matrix[0]) + 1) / 4
    maximumStackSize := stackSize * numOfStacks
    cargo1 := make([]Stack, numOfStacks)
    cargo2 := make([]Stack, numOfStacks)
    for i := 0; i<numOfStacks; i++ {
        cargo1[i] = initStack(maximumStackSize)
        cargo2[i] = initStack(maximumStackSize)
    }

    initCargo(cargo1, matrix)
    initCargo(cargo2, matrix)

    for text, err := reader.ReadString('\n'); err != io.EOF; text, err = reader.ReadString('\n') {
        qnt, src, dst := parseLine(text)
        move1(cargo1, qnt, src, dst)
        move2(cargo2, qnt, src, dst)
    }
    output1 := read(cargo1)
    output2 := read(cargo2)
    
    fmt.Println(output1, output2)
}
