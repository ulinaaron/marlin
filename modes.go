// modes.go

const (
  ModeNormal = iota
  ModeInsert
  ModeVisual
  ModeCommand
)

type Editor struct {
  // ...
  mode int
}

func (e *Editor) HandleKeyPress(key rune) {

  switch e.mode {

  case ModeNormal:
    switch key {
    case 'i':
      e.mode = ModeInsert
    
    case 'v': 
      e.mode = ModeVisual

    case ':':
      e.mode = ModeCommand
    
    // Other keys for navigation etc
    }

  case ModeInsert: 
    if key == escKey {
      e.mode = ModeNormal 
    }

  // Handle input 

  case ModeVisual:
    // Handle selection

  case ModeCommand:
    if key == escKey {
      e.mode = ModeNormal
    }
    // Handle commands
  }

}
