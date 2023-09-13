func (e *Editor) OpenFile(filepath string) {

  e.buffer.text = ReadFile(filepath)
  e.cursor.x = 0
  e.cursor.y = 0

}

func (e *Editor) OpenDefaultBuffer() {
  
  e.buffer.text = []rune("")
  e.cursor.x = 0
  e.cursor.y = 0
  
}

func (e *Editor) SaveFile(filepath string) {

  WriteFile(filepath, e.buffer.text)

}

// Handle cursor keys
func (e *Editor) HandleInput() {

  switch e.input.KeyPress {
    
  case UpArrow:
    e.cursor.y -= 1
    
  case DownArrow:  
    e.cursor.y += 1
    
  // etc
  
  }
  
}
