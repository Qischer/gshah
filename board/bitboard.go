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

var pieceName = []string{
  WPawn: "White Pawn",
  WKnight: "White Knight",
  WBishop: "White Bishop",
  WRook: "White Rook",
  WQueen: "White Queen",
  WKing: "White King",

  BPawn: "Black Pawn",
  BKnight: "Black Knight",
  BBishop: "Black Bishop",
  BRook: "Black Rook",
  BQueen: "Black Queen",
  BKing: "Black King",

  EMPTY: "EMPTY",
}

func (d DisplayPiece) GetName() string {

  return pieceName[d]
}

func Match(b *BitBoard, mask uint64) DisplayPiece {
 
  var piece DisplayPiece
  for piece = 0; int(piece) < len(b.Pieces); piece++ {
    if b.Pieces[piece] & mask == mask {
      return piece
    }
  }

  return EMPTY
}

type BitBoard struct {
  Pieces [12]uint64
  Empty uint64
}

func NewEmptyBoard() *BitBoard {
  b := &BitBoard {}
  b.UpdateEmptySet()

  return b
}

func NewDefaultBoard() *BitBoard {

  b := &BitBoard{
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

  b.UpdateEmptySet()

  return b
}

func (b *BitBoard) UpdateEmptySet() {
  
  empty := uint64(0)
  for p := 0; p < len(b.Pieces); p++ {
    empty |= b.Pieces[p]   
  }

  b.Empty = empty
}

