package main

import (
	"Qischer/gshah/board"
)

func main() {
  
  b := &board.BitBoard{
    WRooks: 5,
  }
  board.Display(b)
}