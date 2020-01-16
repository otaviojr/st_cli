package main

import "os";
import "fmt";

type HelpCommandAction struct {
  BaseCommandAction
}

var contexts []Option = []Option {
  Option {
    option: "help",
    action: HelpCommandAction{BaseCommandAction {name: "Help", description: "Show usage options",},},
  },
  Option {
    option: "device",
    action: DeviceCommandAction{BaseCommandAction {name: "device", description: "Handle smartthings devies"},},
  },
  Option {
    option: "location",
    action: LocationCommandAction{BaseCommandAction {name: "location", description: "Handle smartthings location"},},
  },
  Option {
    option: "scenes",
    action: ScenesCommandAction{BaseCommandAction {name: "scenes", description: "Handle smartthings scenes"},},
  },
}

func (action HelpCommandAction) usage() {
  fmt.Println("\tUsage: st_cli help\r\n")
}

func (action HelpCommandAction) run() bool {
  fmt.Println("st_cli: Smartthings console client");
  fmt.Println("Otavio Ribeiro <otavio.ribeiro@gmail.com>\r\n");

  fmt.Println("Usage: st_cli <context> <command> <options>\r\n")
  fmt.Println("Listing all contexts: \r\n")
  for _, context := range contexts {
    fmt.Printf("Context %s (%s)\r\n\r\n", context.action.getName(), context.action.getDescription())
    context.action.usage()
  }

  return false;
}


func main() {

  if len(os.Args) < 2 {
    fmt.Println("Missing arguments. Type st_cli help to show usage options.");
    os.Exit(1);
  }

  contextPtr := &os.Args[1]

  if *contextPtr == "" {
    fmt.Println("Context not informed. Type st_cli help to show usage options.");
    os.Exit(1);
  }

  var current_context *Option = nil;

  for _, context := range contexts {
    if context.option == *contextPtr {
      current_context = &context;
      break;
    }
  }

  if current_context == nil {
    fmt.Printf("Context %s not fount. Type st_cli help to show usage.\r\n", *contextPtr)
    os.Exit(1)
  }

  current_context.action.run()

}
