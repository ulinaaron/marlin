// main.go
package main

import (
  "fmt"

  "github.com/nsf/termbox-go"
)

func main() {

  // Create editor, renderer, etc
  editor := NewEditor()
  renderer := NewRenderer(editor.buffer, editor.cursor)
  inputHandler := NewInputHandler(editor.buffer, editor.cursor)

  // Set up key bindings
  inputHandler.BindKey(KeyCtrlS, editor.SaveFile)

  // Open file or default buffer
  if len(os.Args) > 1 {
    editor.OpenFile(os.Args[1])
  } else {
    editor.OpenDefaultBuffer() 
  }

  // Main loop
  for {

    // Handle input
    inputHandler.HandleInput()

    // Render screen
    renderer.Render()

    // Sleep to prevent busy loop  
    time.Sleep(time.Millisecond * 16)
  }

}
