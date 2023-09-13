// input.go

// InputHandler handles keyboard input
type InputHandler struct {
  buffer *Buffer

  keybinds map[termbox.Key]func()
}

// BindKey binds a key combination to a handler func
func (ih *InputHandler) BindKey(key termbox.Key, handler func()) {
  ih.keybinds[key] = handler
}

// HandleKey handles a key press
func (ih *InputHandler) HandleKey(key termbox.Key) {
   if handler := ih.keybinds[key]; handler != nil {
     handler()
   }
}

