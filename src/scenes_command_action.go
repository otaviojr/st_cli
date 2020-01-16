package main

import "fmt"

type ScenesCommandAction struct {
  BaseCommandAction
}

func (action ScenesCommandAction) usage() {
  fmt.Println("\tUsage: st_cli scenes <command> <options>\r\n");
}

func (action ScenesCommandAction) run() bool {
  fmt.Println("Scenes Command Action");
  return false;
}

