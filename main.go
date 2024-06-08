package main

import (
	"Qischer/gshah/board"
)

func main() {
  
  b := &board.BitBoard{
    WRooks: 129,
  }
  board.Display(b)
}
