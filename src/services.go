package main

import "net/http"
import "encoding/json"
import "bytes"
//import "io/ioutil"
import "fmt"

type Command struct {
  Component string		    `json:"component"`
  Capability string		    `json:"capability"`
  Command string		      `json:"command"`
  Arguments []interface{}	`json:"arguments"`
}

type Commands struct {
  Commands []*Command		`json:"commands"`
}

func createCommand(capability string, command string, arguments []interface{}) *Command {
  c := Command {Component: "main", Capability: capability, Command: command, Arguments: arguments,};
  return &c;
}

func createCommands() *Commands {
  commands := Commands{Commands: make([]*Command, 0),};
  return &commands;
}

func (commands *Commands) addCommand(command *Command) {
  commands.Commands = append(commands.Commands,command);
}

type RestService struct {
  hostname string
  port int
  protocol string
  token string
}

func createRestService(token string) *RestService {
  service := RestService {hostname: "api.smartthings.com", port: 443, protocol: "https", token: token,};
  return &service;
}

func (service *RestService) createUri(path string) string {
  return fmt.Sprintf("%s://%s:%d/%s", service.protocol, service.hostname, service.port, path);
}

func (service *RestService) getDevice(deviceId string) (string,error) {

  client := &http.Client{
  };

  req, err := http.NewRequest("GET", service.createUri("v1/devices") + "/" + deviceId, nil);

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}

func (service *RestService) getDeviceStatus(deviceId string) (string,error) {

  client := &http.Client{
  };

  req, err := http.NewRequest("GET", service.createUri("v1/devices") + "/" + deviceId + "/status", nil);

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}

func (service *RestService) listDevices(capabilities []string) (string,error) {

  params := "";

  if len(capabilities) > 0 {
    params += "?";
    for _,capability := range capabilities {
      params += "capability=" + capability + "&";
    }
    params += "capabilitiesMode=or";
  }

  client := &http.Client{
  };

  req, err := http.NewRequest("GET", service.createUri("v1/devices") + params, nil);

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}

func (service *RestService) executeCommand(deviceId string, capability string, command string, arguments []interface{}) (string,error) {

  client := &http.Client{
  };

  cmd := createCommand(capability, command, arguments);
  cmds := createCommands();
  cmds.addCommand(cmd);

  jsonStr, err := json.Marshal(cmds);

  if err != nil {
    return "", err;
  }

  req, err := http.NewRequest("POST", service.createUri("v1/devices") + "/" + deviceId + "/commands",  bytes.NewBuffer(jsonStr));

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}

func (service *RestService) listScenes() (string,error) {

  client := &http.Client{
  };

  req, err := http.NewRequest("GET", service.createUri("v1/scenes"), nil);

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}

func (service *RestService) executeScene(scene string) (string,error) {

  client := &http.Client{
  };

  req, err := http.NewRequest("POST", service.createUri("v1/scenes/") + scene + "/execute", nil);

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}

func (service *RestService) listLocations() (string,error) {

  client := &http.Client{
  };

  req, err := http.NewRequest("GET", service.createUri("v1/locations"), nil);

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}


func (service *RestService) listRules(location string) (string,error) {

  client := &http.Client{
  };

  req, err := http.NewRequest("GET", service.createUri("v1/rules") + "?locationId=" + location, nil);

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}

func (service *RestService) getRule(locationId string, ruleId string) (string,error) {

  client := &http.Client{
  };

  req, err := http.NewRequest("GET", service.createUri("v1/rules") + "/" + ruleId + "?locationId=" + locationId, nil);

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}

func (service *RestService) createRule(locationId string, rule []byte) (string,error) {

  client := &http.Client{
  };

  req, err := http.NewRequest("POST", service.createUri("v1/rules") + "?locationId=" + locationId,  bytes.NewBuffer(rule));

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}

func (service *RestService) editRule(locationId string, ruleId string, rule []byte) (string,error) {

  client := &http.Client{
  };

  var raw map[string]interface{};
  if err := json.Unmarshal(rule, &raw); err != nil {
    return "", err;
  }

  delete(raw,"id");

  jsonStr, err := json.Marshal(raw);
  if err != nil {
    return "", err;
  }

  req, err := http.NewRequest("PUT", service.createUri("v1/rules/" + ruleId) + "?locationId=" + locationId,  bytes.NewBuffer(jsonStr));

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}

func (service *RestService) deleteRule(locationId string, ruleId string) (string,error) {

  client := &http.Client{
  };

  req, err := http.NewRequest("DELETE", service.createUri("v1/rules") + "/" + ruleId + "?locationId=" + locationId, nil);

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}

func (service *RestService) executeRule(locationId string, ruleId string) (string,error) {

  client := &http.Client{
  };

  req, err := http.NewRequest("POST", service.createUri("v1/rules/execute") + "/" + ruleId + "?locationId=" + locationId, nil);

  if err != nil {
    return "", err;
  }

  req.Header.Add("Content-Type","application/json");
  req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",service.token));
  req.Header.Add("User-Agent","PostmanRuntime/7.21.0");
  resp, err := client.Do(req);

  if err != nil {
    return "", err;
  }

  defer resp.Body.Close();

  var obj map[string] interface{};
  decoder := json.NewDecoder(resp.Body);
  err = decoder.Decode(&obj);

  if err != nil {
    return "", err;
  }

  ret, err := json.MarshalIndent(obj, "", "\t");

  return string(ret),nil;
}
