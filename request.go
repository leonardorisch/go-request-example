package main
import {
  "net/http"
  "fmt"
  "io/ioutil"
}

func main() {
  requester()
}

func requester() {
  url := "https://jsonplaceholder.typicode.com/"
  resp, err := http.Get(url)
  if err != nil {
  	// handle error
  }
  defer resp.Body.Close()
  body, err := ioutil.Rea1dAll(resp.Body)
  fmt.Printf(string(body))
}
