// commands.go

package main 

// Command represents an editor command
type Command struct {
  Name string
  Handler func()
}

// CommandManager handles command registration and execution
type CommandManager struct {
  commands []Command
}

func (cm *CommandManager) Register(name string, handler func()) {
  cm.commands = append(cm.commands, Command{Name: name, Handler: handler}) 
}

func (cm *CommandManager) Execute(name string) {
  // Lookup and execute command  
}


// commands.go

func (e *Editor) ExecCommand(cmd string) {

  switch cmd {

  case "w": 
    // Write file
  
  case "q":
    // Quit application
  
  case "help":
    // Show help

  // etc
  }

}
