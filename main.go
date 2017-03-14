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
  dataStr := string(data)

  var value interface{}
  switch valueType {
  case "[string]": value = strings.Split(strings.TrimSpace(dataStr), "\n")
  case "float": {
    value,err = strconv.ParseFloat(strings.TrimSpace(dataStr), 64)
    if err != nil { fmt.Println("ERROR: ", err) }
  }
  case "int": {
    value,err = strconv.ParseInt(strings.TrimSpace(dataStr), 10, 0)
    if err != nil { fmt.Println("ERROR: ", err) }
  }
  case "string": fallthrough
  default: value = dataStr
  }

  var jsonData map[string]interface{}
  _ = json.Unmarshal(jsonStr, &jsonData)

  finger := jsonData
  for idx, propName := range propNames {
    if (idx == len(propNames) -1) {
      finger[propName] = value
    } else {
      finger = finger[propName].(map[string]interface{})
    }
  }

  str, err := json.Marshal(jsonData)
  fmt.Print(string(str))
}
