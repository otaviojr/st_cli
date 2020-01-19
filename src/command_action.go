/**
 * @file   command_action.go
 * @author Otavio Ribeiro
 * @date   20 jan 2020
 * @brief  commands base class
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

import "os"
import "fmt"

type BaseCommandAction struct {
  name string
  description string
}

type CommandAction interface {
  getName() string
  getDescription() string
  getToken(cmdLine *CommandLine) string
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

func (action BaseCommandAction) getToken(cmdLine *CommandLine) string {
  token := cmdLine.getStringParameter("token");
  if token == "" {
    token = os.Getenv("SMARTTHINGS_TOKEN");
  }

  return token;
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
