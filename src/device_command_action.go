package main

import "os"
import "fmt"
import "strings"
import "strconv"
/*
 * Device main context action
 */
type DeviceCommandAction struct {
  BaseCommandAction
}

/*
 * NFC commands actions
 */
type DeviceGetCommandAction struct {
  BaseCommandAction
}

type DeviceListCommandAction struct {
  BaseCommandAction
}

type DeviceCommandCommandAction struct {
  BaseCommandAction
}

/*
 * All commands supported in this context
 */
var commands []Option = []Option {
  Option {
    option: "get",
    action: DeviceGetCommandAction{BaseCommandAction{name: "Get", description: "Get a device using device id",},},
  },
  Option {
    option: "list",
    action: DeviceListCommandAction{BaseCommandAction{name: "List", description: "List all devices with a capability",},},
  },
  Option {
    option: "command",
    action: DeviceCommandCommandAction{BaseCommandAction{name: "Command", description: "Send a command to a scpecific device using the device id.",},},
  },
}

func (action DeviceGetCommandAction) usage() {
  fmt.Println("\t\tUsage: st_cli device get <options>\r\n");
}

func (action DeviceGetCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token :=  cmdLine.getStringParameter("token");
  deviceId := cmdLine.getStringParameter("device");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  if deviceId == "" {
    fmt.Println("Device ID missing. Type st_cli help to usage options.");
    return false;
  }

  service := createRestService(token);
  device,err := service.getDevice(deviceId);

  if err != nil {
    fmt.Println("Error searching devices: $s", err);
    return false;
  }

  fmt.Println(device);

  return true;
}

func (action DeviceListCommandAction) usage() {
  fmt.Println("\t\tUsage: st_cli device list <options>\r\n");
}

func (action DeviceListCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token :=  cmdLine.getStringParameter("token");
  capability := cmdLine.getStringParameter("capability");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  capabilities := make([]string,0);
  if capability != "" {
    capabilities = strings.Split(capability,",");
  }

  service := createRestService(token);
  devices,err := service.listDevices(capabilities);

  if err != nil {
    fmt.Println("Error searching devices: $s", err);
    return false;
  }

  fmt.Println(devices);

  return true;
}

func (action DeviceCommandCommandAction) usage() {
  fmt.Println("\t\tUsage: st_cli device command <options>\r\n");
}

func (action DeviceCommandCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token :=  cmdLine.getStringParameter("token");
  deviceId := cmdLine.getStringParameter("device");
  capability := cmdLine.getStringParameter("capability");
  command := cmdLine.getStringParameter("command");
  arguments := cmdLine.getStringParameter("arguments");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  if deviceId == "" {
    fmt.Println("Device ID missing. Type st_cli help to usage options.");
    return false;
  }

  if command == "" {
    fmt.Println("Command missing. Type st_cli help to usage options.");
    return false;
  }

  if capability == "" {
    fmt.Println("Capability missing. Type st_cli help to usage options.");
    return false;
  }

  args := make([]interface{},0);
  if arguments != "" {
    arguments_arr := strings.Split(arguments,",");
    for _,arg := range arguments_arr {
      if val, err := strconv.ParseFloat(string(arg),32); err == nil {
        args = append(args, val);
      } else {
        args = append(args, arg);
      }
    }
  }

  service := createRestService(token);
  device,err := service.executeCommand(deviceId, capability, command, args);

  if err != nil {
    fmt.Println("Error searching devices: $s", err);
    return false;
  }

  fmt.Println(device);

  return true;
}


func (action DeviceCommandAction) usage() {
  fmt.Println("\tUsage: st_cli device <command> <option>\r\n");

  fmt.Println("\tListing all commands:\r\n");
  for _, command := range commands {
    fmt.Printf("\tCommand %s\r\n\t%s\r\n\r\n", command.action.getName(), command.action.getDescription());
    command.action.usage();
  }
}

func (action DeviceCommandAction) run() bool {
  if len(os.Args) < 3 {
    fmt.Println("Device command: missing argument. Type st_cli help to usage options");
    os.Exit(2);
  }

  cmdPtr := &os.Args[2];

  var current_command *Option = nil;

  for _, command := range commands {
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

