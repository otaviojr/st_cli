/**
 * @file   st_cli.go
 * @author Otavio Ribeiro
 * @date   20 jan 2020
 * @brief  smartthings command line client
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
  Option {
    option: "rules",
    action: RulesCommandAction{BaseCommandAction {name: "rules", description: "Handle smartthings rules"},},
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
