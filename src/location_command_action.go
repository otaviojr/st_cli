package main

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

  token:= action.getToken(cmdLine);

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
  action.usageAction(locationCommands);
}
func (action LocationCommandAction) run() bool {
  return action.runAction(locationCommands);
}
