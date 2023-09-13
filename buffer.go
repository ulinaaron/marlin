// buffer.go

// Buffer stores text contents
type Buffer struct {
  text []rune
  cursor int // current cursor position
}

// InsertRune inserts a rune at the cursor position
func (b *Buffer) InsertRune(r rune) {
  // Insert rune and update cursor
}

// DeleteRune deletes a rune at the cursor position 
func (b *Buffer) DeleteRune() {
  // Delete rune and update cursor  
}

// GetText returns the full text 
func (b *Buffer) GetText() string {
  return string(b.text)
}

// SetCursor sets the cursor position
func (b *Buffer) SetCursor(pos int) {
  b.cursor = pos
}
