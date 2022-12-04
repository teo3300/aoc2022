package main

import (
    "fmt"
    "io"
    "unicode"
)

func charToIndex(char rune) (index int) {
    if unicode.IsUpper(char) {
        index = int(char)-int('A')
        index += 26
    } else {
        index = int(char)-int('a')
    }
    return
}

func count1(str[2] string) int {
    var cnt[2][52]int

    for _, letter := range str[0] {
        cnt[0][charToIndex(letter)] = 1
    }
    for _, letter := range str[1] {
        cnt[1][charToIndex(letter)] = 1
    }
    for i := 0; i<len(cnt[0]); i++ {
        if cnt[0][i] > 0 && cnt[1][i] > 0 {
            return i+1
        }
    }
    return 0
}

func count2(str[3] string) int {
    var cnt[3][52]int

    for i := 0; i<3; i++ {
        for _, letter := range str[i] {
            cnt[i][charToIndex(letter)] = 1
        }
    }

    for j := 0; j<52; j++ {
        sum := 0
        for i := 0; i < 3; i++ {
            sum += cnt[i][j]
        }
        if sum == 3 {
            return j+1
        }
        
    }

    return 0
}

func main () {
    var n int
    var err error
    w := ""
    acc1 := 0
    acc2 := 0

    read := 0
    var first [2]string
    var second [3]string

    for n, err = fmt.Scanln(&w); n > 0 && err != io.EOF; n, err = fmt.Scanln(&w) {
        first[0] = w[0:len(w)/2]
        first[1] = w[len(w)/2:len(w)]
        acc1 += count1(first)
        second[read] = w
        read = (read + 1) % 3
        if read == 0 {
            acc2 += count2(second)
        }
    }
    fmt.Println(acc1, acc2)
    return
}
