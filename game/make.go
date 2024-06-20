package game

import (
	"Qischer/gshah/board"
	"fmt"
	"strings"
)

type move struct {
  from uint64
  to uint64
}

func translate(b []byte) uint64 {
  
  cShift := b[0] - 97
  rShift := 9 - (b[1] - 47)
  
  var bit uint64 = (1 << cShift) << (8 * rShift)
  
  return bit
} 

func checkPiece(posByte []byte, b *board.BitBoard) {
  //TODO:
  mask := translate(posByte)  
  piece := board.Match(b, mask)

  fmt.Println(piece.GetName())
}

func movePiece(fromByte []byte, toByte []byte, b *board.BitBoard) {
  
  fromMask := translate(fromByte)
  toMask := translate(toByte)

  fromTo := fromMask ^ toMask
 
  piece := board.Match(b, fromMask)
  if piece == board.EMPTY {
    fmt.Println("No Piece to Move")
    return
  }

  b.Pieces[piece] ^= fromTo
  //b.BPawns ^= fromTo

} 


func Move(str string, b *board.BitBoard) {

  move_bytes := parseStr(str)
  switch len(move_bytes) {
    case 1: 
      fmt.Println("Check Piece")
      checkPiece(move_bytes[0], b)
      break
    case 2: 
      fmt.Println("Make move")
      movePiece(move_bytes[0], move_bytes[1], b)
      break
  }
}

func parseStr(str string) [][]byte {
  
  str = strings.TrimSuffix(str, "\n")
  move_str_arr := strings.Split(str, " ")
  var move_byte_arr [][]byte

  for i := 0; i < len(move_str_arr); i++ {
    move_byte_arr = append(move_byte_arr, []byte(move_str_arr[i]))
  }
  return move_byte_arr 
}
