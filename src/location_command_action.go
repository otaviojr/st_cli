package main

import "fmt"

type LocationCommandAction struct {
  BaseCommandAction
}

func (action LocationCommandAction) usage() {
  fmt.Println("\tUsage: st_cli location <command> <options>\r\n");
}

func (action LocationCommandAction) run() bool {
  fmt.Println("Location Command Action");
  return false;
}

