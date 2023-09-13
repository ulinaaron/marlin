// renderer.go

func (r *Renderer) Render() {

  // Clear screen
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  // Draw text
  x := 0 
  for _, rune := range r.buffer.text {
    termbox.SetCell(x, 0, rune, termbox.ColorDefault, termbox.ColorDefault)
    x++
  }

  // Draw cursors
  for _, cursor := range r.cursorManager.cursors {
    termbox.SetCursor(cursor.x, cursor.y) 
  }

  // Refresh screen
  termbox.Flush()

}

