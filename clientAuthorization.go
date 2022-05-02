package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "https://cbdevice.okta.com/oauth2/aus1l4r7i7vJE9zRF3l7/v1/device/authorize"
  method := "POST"

  payload := strings.NewReader("client_id=0oa1l4rl4c3IQluOA3l7&scope=openid%20profile%20offline_access")

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Cookie", "JSESSIONID=A8F3018CBF307D2A40647339F66EC3F6")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}

