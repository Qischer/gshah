package game

import "Qischer/gshah/board"

func generate(piece board.DisplayPiece, pos uint64) uint64 {
  
  mask := uint64(0)

  switch piece {
    case board.BKing:
      generateKing(pos)
      break
    default:
      break
  }
  
  return mask
}

func generateKing(pos uint64) uint64 {

  mask := (pos << 7) | (pos << 8) | (pos << 9) | (pos >> 1) | (pos << 1) | (pos >> 9) | (pos >> 8) | (pos >> 7)

  return mask
}
