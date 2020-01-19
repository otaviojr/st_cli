package main

import "os"
import "io/ioutil"
import "os/exec"
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
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);
  rule :=  cmdLine.getStringParameter("rule");
  location :=  cmdLine.getStringParameter("location");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  if location == "" {
    fmt.Println("Location missing. Type st_cli help to usage options.");
    return false;
  }

  if rule == "" {
    fmt.Println("Rule missing. Type st_cli help to usage options.");
    return false;
  }

  service := createRestService(token);
  rule,err := service.getRule(location, rule);

  if err != nil {
    fmt.Println("Error searching rules: %s", err);
    return false;
  }

  fmt.Println(rule);

  return true;
}

func (action RulesListCommandAction) usage() {
  fmt.Println("\tUsage: st_cli rules list <options>\r\n");
  fmt.Println("\t\tOptions:\r\n");
  fmt.Println("\t\t--token|-token=\t\t\tSmartthings token\r\n");
  fmt.Println("\t\t--location|-location=\t\t\tLocation Id to listing rules\r\n");
}

func (action RulesListCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);
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
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);
  location :=  cmdLine.getStringParameter("location");
  fileName :=  cmdLine.getStringParameter("file");
  editor :=  cmdLine.getStringParameter("editor");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  if location == "" {
    fmt.Println("Location missing. Type st_cli help to usage options.");
    return false;
  }

  var ruleContent []byte;

  if fileName == "" {
    command := "vi";

    if editor != "" {
      command = editor;
    }

    file, err := ioutil.TempFile(os.TempDir(), "st-cli-rules-");
    if err != nil {
      fmt.Println("Error creating temp file");
      return false;
    }

    defer os.Remove(file.Name());

    cmd := exec.Command(command, file.Name());

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmdErr := cmd.Run();

    if err != nil {
      fmt.Println("Error opening editor: {}", cmdErr);
      return false;
  	}

    dat, err := ioutil.ReadFile(file.Name());
    if err != nil {
      fmt.Println("Error reading rule file: {}", err);
      return false;
  	}

    ruleContent = dat;
  } else {
    dat, err := ioutil.ReadFile(fileName);
    if err != nil {
      fmt.Println("Error reading rule file: {}", err);
      return false;
  	}

    ruleContent = dat;
  }

  if string(ruleContent) == "" {
    fmt.Println("Rule can not be empty");
    return false;
  }

  service := createRestService(token);
  rule,err := service.createRule(location, ruleContent);

  if err != nil {
    fmt.Println("Error searching rules: %s", err);
    return false;
  }

  fmt.Println(rule);

  return true;
}

func (action RulesEditCommandAction) usage() {
  fmt.Println("\tUsage: st_cli rules edit <options>\r\n");
}

func (action RulesEditCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);
  location :=  cmdLine.getStringParameter("location");
  rule :=  cmdLine.getStringParameter("rule");
  fileName :=  cmdLine.getStringParameter("file");
  editor :=  cmdLine.getStringParameter("editor");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  if location == "" {
    fmt.Println("Location missing. Type st_cli help to usage options.");
    return false;
  }

  if rule == "" {
    fmt.Println("Rule missing. Type st_cli help to usage options.");
    return false;
  }

  var ruleContent []byte;

  service := createRestService(token);

  if fileName == "" {
    command := "vi";

    if editor != "" {
      command = editor;
    }

    file, err := ioutil.TempFile(os.TempDir(), "st-cli-rules-");
    if err != nil {
      fmt.Println("Error creating temp file");
      return false;
    }

    defer os.Remove(file.Name());

    ruleStr,err := service.getRule(location, rule);

    ioutil.WriteFile(file.Name(), []byte(ruleStr), 0644);

    cmd := exec.Command(command, file.Name());

    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmdErr := cmd.Run();

    if err != nil {
      fmt.Println("Error opening editor: {}", cmdErr);
      return false;
  	}

    dat, err := ioutil.ReadFile(file.Name());
    if err != nil {
      fmt.Println("Error reading rule file: {}", err);
      return false;
  	}

    ruleContent = dat;
  } else {
    dat, err := ioutil.ReadFile(fileName);
    if err != nil {
      fmt.Println("Error reading rule file: {}", err);
      return false;
  	}

    ruleContent = dat;
  }

  if string(ruleContent) == "" {
    fmt.Println("Rule can not be empty");
    return false;
  }

  ruleEdited,err := service.editRule(location, rule, ruleContent);

  if err != nil {
    fmt.Println("Error searching rules: %s", err);
    return false;
  }

  fmt.Println(ruleEdited);

  return true;
}

func (action RulesExecuteCommandAction) usage() {
  fmt.Println("\tUsage: st_cli rules execute <options>\r\n");
}

func (action RulesExecuteCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);
  rule :=  cmdLine.getStringParameter("rule");
  location :=  cmdLine.getStringParameter("location");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  if location == "" {
    fmt.Println("Location missing. Type st_cli help to usage options.");
    return false;
  }

  if rule == "" {
    fmt.Println("Rule missing. Type st_cli help to usage options.");
    return false;
  }

  service := createRestService(token);
  rule,err := service.executeRule(location, rule);

  if err != nil {
    fmt.Println("Error searching rules: %s", err);
    return false;
  }

  fmt.Println(rule);

  return true;
}

func (action RulesDeleteCommandAction) usage() {
  fmt.Println("\tUsage: st_cli rules delete <options>\r\n");
}

func (action RulesDeleteCommandAction) run() bool {
  cmdLine := createCommandLineParser();

  token:= action.getToken(cmdLine);
  rule :=  cmdLine.getStringParameter("rule");
  location :=  cmdLine.getStringParameter("location");

  if token == "" {
    fmt.Println("Smartthings token missing. Type st_cli help to usage options.");
    return false;
  }

  if location == "" {
    fmt.Println("Location missing. Type st_cli help to usage options.");
    return false;
  }

  if rule == "" {
    fmt.Println("Rule missing. Type st_cli help to usage options.");
    return false;
  }

  service := createRestService(token);
  rule,err := service.deleteRule(location, rule);

  if err != nil {
    fmt.Println("Error searching rules: %s", err);
    return false;
  }

  fmt.Println(rule);

  return true;
}

func (action RulesCommandAction) usage() {
  fmt.Println("\tUsage: st_cli rules <command> <option>\r\n");
  action.usageAction(rulesCommands);
}

func (action RulesCommandAction) run() bool {
  return action.runAction(rulesCommands);
}
