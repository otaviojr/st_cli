/**
 * @file   device_command_action.go
 * @author Otavio Ribeiro
 * @date   20 jan 2020
 * @brief  devices context/commands
 *
 * Copyright (c) 2020 Otávio Ribeiro <otavio.ribeiro@gmail.com>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */
package main

import "fmt"
import "strings"
import "strconv"
/*
 * Device main context action
 */
type DeviceCommandAction struct {
  BaseCommandAction
}

type DeviceStatusCommandAction struct {
  BaseCommandAction
}

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
var deviceCommands []Option = []Option {
  Option {
    option: "status",
    action: DeviceStatusCommandAction{BaseCommandAction{name: "Status", description: "Get a device status using the device id",},},
  },
  Option {
    option: "get",
    action: DeviceGetCommandAction{BaseCommandAction{name: "Get", description: "Get a device information using the device id",},},
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

func (action DeviceStatusCommandAction) usage() {
  fmt.Println("\t\tUsage: st_cli device status <options>\r\n");
  fmt.Println("\t\tOptions:\r\n");
  fmt.Println("\t\t--token|-token=\t\tSmartthings token\r\n");
  fmt.Println("\t\t--device|-device=\tDevice Id to retrieve status\r\n");
}

func (action DeviceStatusCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);
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
  device,err := service.getDeviceStatus(deviceId);

  if err != nil {
    fmt.Println("Error searching devices: %s", err);
    return false;
  }

  fmt.Println(device);

  return true;
}

func (action DeviceGetCommandAction) usage() {
  fmt.Println("\t\tUsage: st_cli device get <options>\r\n");
  fmt.Println("\t\tOptions:\r\n");
  fmt.Println("\t\t--token|-token=\t\tSmartthings token\r\n");
  fmt.Println("\t\t--device|-device=\tDevice Id to retrieve informations\r\n");
}

func (action DeviceGetCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);
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
    fmt.Println("Error searching devices: %s", err);
    return false;
  }

  fmt.Println(device);

  return true;
}

func (action DeviceListCommandAction) usage() {
  fmt.Println("\t\tUsage: st_cli device list <options>\r\n");
  fmt.Println("\t\tOptions:\r\n");
  fmt.Println("\t\t--token|-token=\t\t\tSmartthings token\r\n");
  fmt.Println("\t\t--capability|-capability=\tSmartthings capability to filter devices");
  fmt.Println("\t\t\t\t\t\tIf not informed all devices will be returned\r\n");
}

func (action DeviceListCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);
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
    fmt.Println("Error searching devices: %s", err);
    return false;
  }

  fmt.Println(devices);

  return true;
}

func (action DeviceCommandCommandAction) usage() {
  fmt.Println("\t\tUsage: st_cli device command <options>\r\n");
  fmt.Println("\t\tOptions:\r\n");
  fmt.Println("\t\t--token|-token=\t\t\tSmartthings token\r\n");
  fmt.Println("\t\t--device|-device=\t\tDevice Id that will receive the command/arguments\r\n");
  fmt.Println("\t\t--capability|-capability=\tSmartthings capability of this device\r\n");
  fmt.Println("\t\t--command|-command=\t\tCommand that will be send\r\n");
  fmt.Println("\t\t--arguments|-arguments=\t\tArguments to be send with the command. Comma separated.\r\n");
}

func (action DeviceCommandCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);
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
    fmt.Println("Error searching devices: %s", err);
    return false;
  }

  fmt.Println(device);

  return true;
}


func (action DeviceCommandAction) usage() {
  fmt.Println("\tUsage: st_cli device <command> <option>\r\n");

  action.usageAction(deviceCommands);
}

func (action DeviceCommandAction) run() bool {
  return action.runAction(deviceCommands);
}
