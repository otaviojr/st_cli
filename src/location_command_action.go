/**
 * @file   location_command_action.go
 * @author Otavio Ribeiro
 * @date   20 jan 2020
 * @brief  location context/commands
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
  fmt.Println("\t\t--token|-token=\t\tSmartthings token\r\n");
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
