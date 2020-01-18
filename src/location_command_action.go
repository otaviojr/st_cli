package main

import "os"
import "fmt"

type LocationCommandAction struct {
  BaseCommandAction
}

type LocationListCommandAction struct {
  BaseCommandAction
}

var locationCommands []Option = []Option {
  Option {
    option: "list",
    action: LocationListCommandAction{BaseCommandAction{name: "List", description: "List all locations",},},
  },
}

func (action LocationListCommandAction) usage() {
  fmt.Println("\t\tUsage: st_cli location list <options>\r\n");
  fmt.Println("\t\tOptions:\r\n");
  fmt.Println("\t\t--token|-token=\t\t\tSmartthings token\r\n");
}

func (action LocationListCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token :=  cmdLine.getStringParameter("token");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  service := createRestService(token);
  locations,err := service.listLocations();

  if err != nil {
    fmt.Println("Error searching devices: %s", err);
    return false;
  }

  fmt.Println(locations);

  return true;
}


func (action LocationCommandAction) usage() {
  fmt.Println("\tUsage: st_cli location <command> <options>\r\n");

  fmt.Println("\tListing all commands:\r\n");
  for _, command := range locationCommands {
    fmt.Printf("\tCommand %s\r\n\t%s\r\n\r\n", command.action.getName(), command.action.getDescription());
    command.action.usage();
  }}

func (action LocationCommandAction) run() bool {
  if len(os.Args) < 3 {
    fmt.Println("Device command: missing argument. Type st_cli help to usage options");
    os.Exit(2);
  }

  cmdPtr := &os.Args[2];

  var current_command *Option = nil;

  for _, command := range locationCommands {
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
