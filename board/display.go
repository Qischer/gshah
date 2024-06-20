package board

import (
	"fmt"
)

type DisplayPiece int 
var icons = []string {
  WPawn: "♙",
  WKnight: "♘",
  WBishop: "♗",
  WRook: "♖",
  WQueen: "♕",
  WKing: "♔",
  BPawn: "♟︎",
  BKnight: "♞",
  BBishop: "♝",
  BRook: "♜",
  BQueen: "♛",
  BKing: "♚",
  EMPTY: ".",
}

func (d DisplayPiece) String() string {
  return icons[d] 
}

func Display(b *BitBoard) {
  
  var sq [64]string
  //scan

  var mask uint64 = 1
  for i:=0; i < len(sq); i++ {
    piece := Match(b, mask)
    
    sq[i] = piece.String()
    mask = mask << 1
  }

  printBoard(sq)

  fmt.Print("   a b c d e f g h\n")
}

func printBoard(pieces [64]string) {
  

  for i := 0; i < 64; i++ {
    //Start line
    if i % 8 == 0 {
      fmt.Printf("%d[", 8 - i/8)
    }

    fmt.Printf(" %v", pieces[i])

    //End line
    if i % 8 == 7 {
      fmt.Print(" ]\n")
    }
  }
}

