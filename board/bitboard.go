package board

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

type BitBoard struct {
  Pieces [12]uint64

}

func NewEmptyBoard() *BitBoard {
  return &BitBoard {}
}

func NewDefaultBoard() *BitBoard {

  return &BitBoard{
    Pieces : [12]uint64 {
      WRook: 1 << 63 + 1 << 56,
      WKnight: 1 << 62 + 1 << 57,
      WBishop: 1 << 61 + 1 << 58,
      WQueen: 1 << 59,
      WKing: 1 << 60,
      WPawn: (1 << 8-1) << 48,

      BRook: 1 + 1 << 7,
      BKnight: 1 << 1 + 1 << 6,
      BBishop: 1 << 2 + 1 << 5,
      BQueen: 1 << 3,
      BKing: 1 << 4,
      BPawn: (1 << 8-1) << 8,
    },
  }
}

