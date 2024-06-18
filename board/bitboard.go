package board

type BitBoard struct {
  WPawns uint64
  WKnights uint64
  WBishops uint64
  WRooks uint64
  WQueens uint64
  WKing uint64

  BPawns uint64
  BKnights uint64
  BBishops uint64
  BRooks uint64
  BQueens uint64
  BKing uint64
}

func NewEmptyBoard() *BitBoard {
  return &BitBoard {}
}

func NewDefaultBoard() *BitBoard {

  return &BitBoard{
    WRooks: 1 << 63 + 1 << 56,
    WKnights: 1 << 62 + 1 << 57,
    WBishops: 1 << 61 + 1 << 58,
    WQueens: 1 << 59,
    WKing: 1 << 60,
    WPawns: (1 << 8-1) << 48,

    BRooks: 1 + 1 << 7,
    BKnights: 1 << 1 + 1 << 6,
    BBishops: 1 << 2 + 1 << 5,
    BQueens: 1 << 3,
    BKing: 1 << 4,
    BPawns : (1 << 8-1) << 8,
  }
}

