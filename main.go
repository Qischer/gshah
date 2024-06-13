package main

import (
	"Qischer/gshah/board"
)

func main() {
  
  b := &board.BitBoard{
    WRooks: 129,
    WKnights: 66,
    WBishops: 36,
  }
  board.Display(b)
}
