package main

import (
	"Qischer/gshah/board"
	"Qischer/gshah/game"
	"bufio"
	"log"
	"os"
)

func main() {
  
  b := board.NewEmptyBoard()
  b.Pieces[board.BKing] = 1 << 28

  board.Display(b)
  reader := bufio.NewReader(os.Stdin)
  for {
    str, err := reader.ReadString('\n')

    if err != nil {
      log.Fatal("Input error")
    }

    game.Move(str, b) 

    board.Display(b)
  }
 }
