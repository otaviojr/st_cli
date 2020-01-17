package main

import "os"
import "fmt"

type ScenesCommandAction struct {
  BaseCommandAction
}

type ScenesListCommandAction struct {
  BaseCommandAction
}

type ScenesExecuteCommandAction struct {
  BaseCommandAction
}

var sceneCommands []Option = []Option {
  Option {
    option: "list",
    action: ScenesListCommandAction{BaseCommandAction{name: "List", description: "List all available scenes",},},
  },
  Option {
    option: "execute",
    action: ScenesExecuteCommandAction{BaseCommandAction{name: "Execute", description: "Execute a scene",},},
  },
}

func (action ScenesListCommandAction) usage() {
  fmt.Println("\tUsage: st_cli scenes list <option>\r\n");
  fmt.Println("\t\tOptions:\r\n");
  fmt.Println("\t\t--token|-token=\t\t\tSmartthings token\r\n");
}

func (action ScenesListCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token :=  cmdLine.getStringParameter("token");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  service := createRestService(token);
  scenes,err := service.listScenes();

  if err != nil {
    fmt.Println("Error searching scenes: %s", err);
    return false;
  }

  fmt.Println(scenes);

  return true;
}

func (action ScenesExecuteCommandAction) usage() {
  fmt.Println("\tUsage: st_cli scenes execute <option>\r\n");
  fmt.Println("\t\tOptions:\r\n");
  fmt.Println("\t\t--token|-token=\t\t\tSmartthings token\r\n");
  fmt.Println("\t\t--scene|-scene=\t\t\tScene ID to be executed\r\n");
}

func (action ScenesExecuteCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token :=  cmdLine.getStringParameter("token");
  scene :=  cmdLine.getStringParameter("scene");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  if scene == "" {
    fmt.Println("Smartthings scene missing. Type st_cli help to usage options.");
    return false;
  }

  service := createRestService(token);
  out,err := service.executeScene(scene);

  if err != nil {
    fmt.Println("Error searching scenes: %s", err);
    return false;
  }

  fmt.Println(out);

  return true;
}

func (action ScenesCommandAction) usage() {
  fmt.Println("\tUsage: st_cli scenes <command> <option>\r\n");

  fmt.Println("\tListing all commands:\r\n");
  for _, command := range sceneCommands {
    fmt.Printf("\tCommand %s\r\n\t%s\r\n\r\n", command.action.getName(), command.action.getDescription());
    command.action.usage();
  }
}

func (action ScenesCommandAction) run() bool {
  if len(os.Args) < 3 {
    fmt.Println("Scenes command: missing argument. Type st_cli help to usage options");
    os.Exit(2);
  }

  cmdPtr := &os.Args[2];

  var current_command *Option = nil;

  for _, command := range sceneCommands {
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
