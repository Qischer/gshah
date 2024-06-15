package board

import (
	"fmt"
)

const (
  WPawn = iota
  WKnight
  WBishop
  WRook
  WQueen
  WKing
  BPawn
  BKnight
  BBishop
  BRook
  BQueen
  BKing
  EMPTY
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
    piece := match(b, mask)
    
    sq[i] = piece.String()
    mask = mask << 1
  }

  printBoard(sq)
  //fmt.Println(sq)

  /* fmt.Printf(`
     a b c d e f g h
   +----------------+
  8| %v %v %v %v %v %v n r |8
  7| p p p p p p p p |7
  6| . . . . . . . . |6
  5| . . . . . . . . |5
  4| . . . . . . . . |4
  3| . . . . . . . . |3
  2| P P P P P P P P |2
  1| R N B Q K B N ♟︎ |1
   +----------------+
     a b c d e f g h

  `,) */
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

func match(b *BitBoard, mask uint64) DisplayPiece {
  
  if b.WPawns & mask == mask {
    return WPawn
  }

  if b.WKnights & mask == mask {
    return WKnight
  }

  if b.WBishops & mask == mask {
     return WBishop
  }
  
  if b.WRooks & mask == mask {
    return WRook
  }

  if b.WQueens & mask  == mask {
    return WQueen
  }

  if b.WKing & mask == mask {
    return WKing
  }

  if b.WPawns & mask == mask {
    return WPawn
  }

  if b.WKnights & mask == mask {
    return WKnight
  }

  if b.WBishops & mask == mask {
     return WBishop
  }
  
  if b.WRooks & mask == mask {
    return WRook
  }

  if b.WQueens & mask  == mask {
    return WQueen
  }

  if b.WKing & mask == mask {
    return WKing
  }

  if b.BPawns & mask == mask {
    return BPawn
  }

  if b.BKnights & mask == mask {
    return BKnight
  }

  if b.BBishops & mask == mask {
     return BBishop
  }
  
  if b.BRooks & mask == mask {
    return BRook
  }

  if b.BQueens & mask  == mask {
    return BQueen
  }

  if b.BKing & mask == mask {
    return BKing
  }

  if b.BPawns & mask == mask {
    return BPawn
  }

  if b.BKnights & mask == mask {
    return BKnight
  }

  if b.BBishops & mask == mask {
     return BBishop
  }
  
  if b.BRooks & mask == mask {
    return BRook
  }

  if b.BQueens & mask  == mask {
    return BQueen
  }

  if b.BKing & mask == mask {
    return BKing
  }

  return EMPTY
}
