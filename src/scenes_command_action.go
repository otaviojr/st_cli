package main

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

var scenesCommands []Option = []Option {
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
  action.usageAction(scenesCommands);
}

func (action ScenesCommandAction) run() bool {
  return action.runAction(scenesCommands);
}
