/**
 * @file   scenes_command_action.go
 * @author Otavio Ribeiro
 * @date   20 jan 2020
 * @brief  scenes context/commands
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
  fmt.Println("\t\t--token|-token=\t\tSmartthings token\r\n");
}

func (action ScenesListCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);

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
  fmt.Println("\t\t--token|-token=\t\tSmartthings token\r\n");
  fmt.Println("\t\t--scene|-scene=\t\tScene ID to be executed\r\n");
}

func (action ScenesExecuteCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);
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
