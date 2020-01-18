package main

import "os"
import "fmt"

type BaseCommandAction struct {
  name string
  description string
}

type CommandAction interface {
  getName() string
  getDescription() string
  run() bool
  runAction(options []Option) bool
  usage()
  usageAction(options []Option)
}

func (action BaseCommandAction) getName() string {
  return action.name
}

func (action BaseCommandAction) getDescription() string {
  return action.description
}

func (action BaseCommandAction) run() bool {
  fmt.Println("\tRunning run on base class. Something is wrong.\r\n")
  return false
}

func (action BaseCommandAction) runAction(options []Option) bool {
  if len(os.Args) < 3 {
    fmt.Println("Device command: missing argument. Type st_cli help to usage options");
    os.Exit(2);
  }

  cmdPtr := &os.Args[2];

  var current_command *Option = nil;

  for _, command := range options {
    if command.option == *cmdPtr {
      current_command = &command;
      break;
    }
  }

  if current_command == nil {
    fmt.Printf("Command %s not found. Type st_cli help to usage options.\r\n\r\n", *cmdPtr);
    os.Exit(2);
  }

  return current_command.action.run();
}

func (action BaseCommandAction) usage() {
  fmt.Println("\tRunning usage on base class. Something is wrong.\r\n")
}

func (action BaseCommandAction) usageAction(options []Option) {
  fmt.Println("\tListing all commands:\r\n");
  for _, command := range options {
    fmt.Printf("\tCommand %s\r\n\t%s\r\n\r\n", command.action.getName(), command.action.getDescription());
    command.action.usage();
  }
}
