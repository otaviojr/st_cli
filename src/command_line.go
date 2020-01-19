/**
 * @file   command_line.go
 * @author Otavio Ribeiro
 * @date   20 jan 2020
 * @brief  command line parser
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
import "strings"
import "fmt"

type CommandLine struct {
}

func createCommandLineParser() *CommandLine {
  cmdLine := CommandLine {}
  return &cmdLine
}

func (cmdLine CommandLine) findToken(token string) (int, string) {
  position := 0
  ret := ""

  for index, value := range os.Args {
    if strings.HasPrefix(value,fmt.Sprintf("-%s=",token)) {
      position = index
      ret = strings.Split(value,"=")[1]
      break
    } else if value == fmt.Sprintf("--%s", token) {
      position = index
      if len(os.Args) > index+1 {
        argument := os.Args[index + 1]
        if !strings.HasPrefix(argument,"-") && !strings.HasPrefix(argument,"--") {
          ret = os.Args[index + 1]
        }
      }
      break
    }
  }
  return position,ret
}

func (cmdLine CommandLine) getStringParameter(name string) string {
  _,value := cmdLine.findToken(name)
  return value
}
