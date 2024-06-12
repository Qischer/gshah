package main

import (
	"Qischer/gshah/board"
)

func main() {
  
  b := &board.BitBoard{
    WRooks: 128,
  }
  board.Display(b)
}
