package bitfield

// a bitfield represents the pieces that a peer has
type BitField []byte

// HasPiece returns true if the peer has a particular piece
func (bf BitField) HasPiece(index int) bool {
	byteIndex := index / 8
	offset := index%8
	if byteIndex < 0 || byteIndex >= len(bf) {
		return false
	}

	return bf[byteIndex]>>uint(7-offset)&1 != 0
}

// SetPiece sets a bit in the bitfield
func (bf BitField) SetPiece(index int) {
	byteIndex := index / 8
	offset := index % 8

	if byteIndex < 0 || byteIndex >= len(bf) {
		return
	}

	bf[byteIndex] |= 1 << uint(7-offset)
}
