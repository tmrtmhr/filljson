/*
Usage: cat template.json | filljson ${ValueType} "path.to.target" <(some command)
*/
package main

import (
  "flag"
  "encoding/json"
  "io/ioutil"
  "fmt"
  "os"
  "strings"
  "strconv"
)

type Json map[string]interface{}

func main() {
  flag.Parse()
  jsonStr, err := ioutil.ReadAll(os.Stdin)

  valueType := flag.Arg(0)
  keyPath := flag.Arg(1)
  propNames := strings.Split(keyPath, ".")
  inputFile := flag.Arg(2)

  data, err := ioutil.ReadFile(inputFile)
  if err != nil {
    fmt.Println("ERROR: ", err)
  }

  value := Parse(valueType, string(data))

  var jsonData map[string]interface{}
  _ = json.Unmarshal(jsonStr, &jsonData)

  jsonData = Fill(jsonData,propNames,value)

  str, err := json.Marshal(jsonData)
  fmt.Print(string(str))
}

func Parse(valueType string, valueStr string) interface{} {
  var value interface{}
  var err error
  switch valueType {
  case "[string]": value = strings.Split(strings.TrimSpace(valueStr), "\n")
  case "float": {
    value,err = strconv.ParseFloat(strings.TrimSpace(valueStr), 64)
    if err != nil { fmt.Println("ERROR: ", err) }
  }
  case "int": {
    value,err = strconv.ParseInt(strings.TrimSpace(valueStr), 10, 0)
    if err != nil { fmt.Println("ERROR: ", err) }
  }
  case "string": fallthrough
  default: value = valueStr
  }
  return value
}


func Fill(json Json, propNames []string, value interface{}) Json {
  finger := json
  for idx, propName := range propNames {
    if (idx == len(propNames) -1) {
      finger[propName] = value
    } else {
      finger = finger[propName].(map[string]interface{})
    }
  }
  return json
}
