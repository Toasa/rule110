package main

import (
    "rule110/board"
)

const (
    // BOARD_SIZE = 30
    PRINT_LOOP_NUM = 50
)

func main() {
    // MakeBoard(seed_value)関数を作ってはどうか？
    // seed = 25899906842524
    // b := MakeBoard(seed)

    b := [board.BOARD_SIZE]int{}

    // 右端だけbitを立てる
    b[board.BOARD_SIZE - 1] = 1

    board.PrintBoard(b)
    for i:= 0; i < PRINT_LOOP_NUM; i++ {
        b = board.NextStep(b)
        board.PrintBoard(b)
    }
}
