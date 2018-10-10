package main

import (
    "rule110/board"
)

const (
    BOARD_SIZE = 50
)

func main() {
    // MakeBoard(seed_value)関数を作ってはどうか？
    // seed = 25899906842524
    // b := MakeBoard(seed)

    b := [BOARD_SIZE]int{}

    // 右端だけbitを立てる
    b[BOARD_SIZE - 1] = 1
    
    board.PrintBoard(b)
    for i:= 0; i < 500; i++ {
        b = board.NextStep(b)
        board.PrintBoard(b)
    }
}
