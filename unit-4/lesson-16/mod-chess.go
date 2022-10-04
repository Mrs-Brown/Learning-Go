//Modified chess.go to display a chess game
package main

import "fmt"

func display(board [8][8]rune) {  //Function that accepts array of type rune
    for _, row := range board {  //Inside a row
        for _, column := range row {  //Inside a column
            if column == 0 {  //First column 
                fmt.Print("  ")
            } else {  //Otherwise
                fmt.Printf("%c", column)  //%c is used to display runes
            }
        }
        fmt.Println()  //Print piece
    }
}

func main() {
    var board [8][8]rune  //Array of array of type rune

    //black pieces
    board[0][0] = 'r'  //rook
    board[0][1] = 'n'  //knight
    board[0][2] = 'b'  //bishop
    board[0][3] = 'q'  //queen
    board[0][4] = 'k'  //king
    board[0][5] = 'b'  //bishop
    board[0][6] = 'n'  //knight
    board[0][7] = 'r'  //rook

    //pawns
    for column := range board[1] {
        board[1][column] = 'p'  //black pawns
        board[6][column] = 'P'  //white pawns
    }

    //white pieces
    board[7][0] = 'R'
    board[7][1] = 'N'
    board[7][2] = 'B'
    board[7][3] = 'Q'
    board[7][4] = 'K'
    board[7][5] = 'B'
    board[7][6] = 'N'
    board[7][7] = 'R'

    display(board)
}
