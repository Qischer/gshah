package main

import (
	"Qischer/gshah/board"
	"Qischer/gshah/game"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
  
  b := board.NewEmptyBoard()
  
  reader := bufio.NewReader(os.Stdin)
  for {
    str, _, err := reader.ReadLine()

    if err != nil {
      log.Fatal("Input error")
    }

    b.BPawns = game.Translate(str[0:2]) 

    board.Display(b)
    fmt.Println(str)
  }
 }
