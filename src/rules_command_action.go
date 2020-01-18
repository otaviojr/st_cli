package main

import "os"
import "fmt"

type RulesCommandAction struct {
  BaseCommandAction
}

type RulesGetCommandAction struct {
  BaseCommandAction
}

type RulesListCommandAction struct {
  BaseCommandAction
}

type RulesCreateCommandAction struct {
  BaseCommandAction
}

type RulesEditCommandAction struct {
  BaseCommandAction
}

type RulesExecuteCommandAction struct {
  BaseCommandAction
}

type RulesDeleteCommandAction struct {
  BaseCommandAction
}

/*
 * All commands supported in this context
 */
var rulesCommands []Option = []Option {
  Option {
    option: "get",
    action: RulesGetCommandAction{BaseCommandAction{name: "Get", description: "Get a rule using the rule id",},},
  },
  Option {
    option: "list",
    action: RulesListCommandAction{BaseCommandAction{name: "List", description: "List all rules",},},
  },
  Option {
    option: "create",
    action: RulesCreateCommandAction{BaseCommandAction{name: "Create", description: "Create a new rule editing the json on vi,vim or nano editor",},},
  },
  Option {
    option: "edit",
    action: RulesEditCommandAction{BaseCommandAction{name: "Edit", description: "Edit a rule editing the json on vi,vim or nano editor",},},
  },
  Option {
    option: "execute",
    action: RulesExecuteCommandAction{BaseCommandAction{name: "Execute", description: "Execute a rule using the rule id",},},
  },
  Option {
    option: "delete",
    action: RulesDeleteCommandAction{BaseCommandAction{name: "Delete", description: "Delete a rule using the rule id",},},
  },
}

func (action RulesGetCommandAction) usage() {
  fmt.Println("\tUsage: st_cli rules get <options>\r\n");
}

func (action RulesGetCommandAction) run() bool {
  fmt.Println("Rules Get Action");
  return false;
}

func (action RulesListCommandAction) usage() {
  fmt.Println("\tUsage: st_cli rules list <options>\r\n");
  fmt.Println("\t\tOptions:\r\n");
  fmt.Println("\t\t--token|-token=\t\t\tSmartthings token\r\n");
  fmt.Println("\t\t--location|-location=\t\t\tLocation Id to listing rules\r\n");
}

func (action RulesListCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token :=  cmdLine.getStringParameter("token");
  location :=  cmdLine.getStringParameter("location");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  if location == "" {
    fmt.Println("Location missing. Type st_cli help to usage options.");
    return false;
  }

  service := createRestService(token);
  rules,err := service.listRules(location);

  if err != nil {
    fmt.Println("Error searching rules: %s", err);
    return false;
  }

  fmt.Println(rules);

  return true;
}

func (action RulesCreateCommandAction) usage() {
  fmt.Println("\tUsage: st_cli rules create <options>\r\n");
}

func (action RulesCreateCommandAction) run() bool {
  fmt.Println("Rules Create Action");
  return false;
}

func (action RulesExecuteCommandAction) usage() {
  fmt.Println("\tUsage: st_cli rules execute <options>\r\n");
}

func (action RulesExecuteCommandAction) run() bool {
  fmt.Println("Rules Execute Action");
  return false;
}

func (action RulesDeleteCommandAction) usage() {
  fmt.Println("\tUsage: st_cli rules delete <options>\r\n");
}

func (action RulesDeleteCommandAction) run() bool {
  fmt.Println("Rules Delete Action");
  return false;
}

func (action RulesCommandAction) usage() {
  fmt.Println("\tUsage: st_cli rules <command> <option>\r\n");

  fmt.Println("\tListing all commands:\r\n");
  for _, command := range rulesCommands {
    fmt.Printf("\tCommand %s\r\n\t%s\r\n\r\n", command.action.getName(), command.action.getDescription());
    command.action.usage();
  }
}

func (action RulesCommandAction) run() bool {
  if len(os.Args) < 3 {
    fmt.Println("Rules command: missing argument. Type st_cli help to usage options");
    os.Exit(2);
  }

  cmdPtr := &os.Args[2];

  var current_command *Option = nil;

  for _, command := range rulesCommands {
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