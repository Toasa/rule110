package board

import (
    "fmt"
    // "strconv"
)

const (
    BOARD_SIZE = 50
)

func NextStep(b [BOARD_SIZE]int) [BOARD_SIZE]int {
    b_new := b
    for i:= 1; i < BOARD_SIZE - 1; i++ {
        b_new[i] = rule(b[i - 1], b[i], b[i + 1])
    }
    // edge case
    b_new[0] = rule(0, b[0], b[1])
    b_new[BOARD_SIZE-1] = rule(b[BOARD_SIZE-2], b[BOARD_SIZE-1], 0)

    return b_new
}

//
// 111 110 101 100 011 010 001 000
//
// 000 010 010 000 010 010 010 000
//
func rule(l, c, r int) int {
    if l == 1 {
        if c == 1 {
            if r == 1 {
                // 111
                return 0
            } else {
                // 110
                return 1
            }
        } else {
            if r == 1 {
                // 101
                return 1
            } else {
                // 100
                return 0
            }
        }
    } else {
        if c == 1 {
            if r == 1 {
                // 011
                return 1
            } else {
                // 010
                return 1
            }
        } else {
            if r == 1 {
                // 001
                return 1
            } else {
                // 000
                return 0
            }
        }
    }
}

func PrintBoard(b [BOARD_SIZE]int) {
    for _, num := range b {
        if num == 0 {
            fmt.Printf("□")
        } else {
            fmt.Printf("■")
        }
    }
    fmt.Printf("\n")
}

// func MakeBoard(seed int64) [BOARD_SIZE]int{
//     b := [BOARD_SIZE]int{}
//     // seedを2進数値に変換
//     strconv.FormatInt(seed, 2)
// }
