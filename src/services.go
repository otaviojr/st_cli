/**
 * @file   services.go
 * @author Otavio Ribeiro
 * @date   20 jan 2020
 * @brief  SmartThings API wrapper
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

import "net/http"
import "encoding/json"
import "bytes"
import "io"
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

func (service *RestService) serviceResquest(method string, endpoint string, body io.Reader) (string,error) {
  client := &http.Client{
  };

  req, err := http.NewRequest(method, endpoint, body);

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

func (service *RestService) getDevice(deviceId string) (string,error) {
  return service.serviceResquest("GET", service.createUri("v1/devices") + "/" + deviceId, nil);
}

func (service *RestService) getDeviceStatus(deviceId string) (string,error) {
  return service.serviceResquest("GET", service.createUri("v1/devices") + "/" + deviceId + "/status", nil);
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

  return service.serviceResquest("GET", service.createUri("v1/devices") + params, nil);
}

func (service *RestService) executeCommand(deviceId string, capability string, command string, arguments []interface{}) (string,error) {
  cmd := createCommand(capability, command, arguments);
  cmds := createCommands();
  cmds.addCommand(cmd);

  jsonStr, err := json.Marshal(cmds);

  if err != nil {
    return "", err;
  }

  return service.serviceResquest("POST", service.createUri("v1/devices") + "/" + deviceId + "/commands", bytes.NewBuffer(jsonStr));
}

func (service *RestService) listScenes() (string,error) {
  return service.serviceResquest("GET", service.createUri("v1/scenes"), nil);
}

func (service *RestService) executeScene(scene string) (string,error) {
  return service.serviceResquest("POST", service.createUri("v1/scenes/") + scene + "/execute", nil);
}

func (service *RestService) listLocations() (string,error) {
  return service.serviceResquest("GET", service.createUri("v1/locations"), nil);
}


func (service *RestService) listRules(location string) (string,error) {
  return service.serviceResquest("GET", service.createUri("v1/rules") + "?locationId=" + location, nil);
}

func (service *RestService) getRule(locationId string, ruleId string) (string,error) {
  return service.serviceResquest("GET", service.createUri("v1/rules") + "/" + ruleId + "?locationId=" + locationId, nil);
}

func (service *RestService) createRule(locationId string, rule []byte) (string,error) {
  return service.serviceResquest("POST", service.createUri("v1/rules") + "?locationId=" + locationId, bytes.NewBuffer(rule));
}

func (service *RestService) editRule(locationId string, ruleId string, rule []byte) (string,error) {
  var raw map[string]interface{};
  if err := json.Unmarshal(rule, &raw); err != nil {
    return "", err;
  }

  delete(raw,"id");

  jsonStr, err := json.Marshal(raw);
  if err != nil {
    return "", err;
  }

  return service.serviceResquest("PUT",service.createUri("v1/rules/" + ruleId) + "?locationId=" + locationId, bytes.NewBuffer(jsonStr));
}

func (service *RestService) deleteRule(locationId string, ruleId string) (string,error) {
  return service.serviceResquest("DELETE",service.createUri("v1/rules") + "/" + ruleId + "?locationId=" + locationId, nil);
}

func (service *RestService) executeRule(locationId string, ruleId string) (string,error) {
  return service.serviceResquest("POST", service.createUri("v1/rules/execute") + "/" + ruleId + "?locationId=" + locationId, nil);
}
