package game

type move struct {
  from uint64
  to uint64
}

func Translate(b []byte) uint64 {
  
  cShift := b[0] - 97
  rShift := 9 - (b[1] - 47)
  
  var bit uint64 = (1 << cShift) << (8 * rShift)
  
  return bit
} 
