package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

type Quotes struct {
  Contents struct {
    Quotes []Quote
  }
}

type Quote struct {
  Quote string
}

func main() {
  contents, err := fetchQuotes()
  if err != nil {
    fmt.Println(err)
    return
  }

  quote, err := getQuoteFromContents(contents)
  if err != nil {
    fmt.Println(err)
    return
  }
  
  fmt.Println(quote)
}

func fetchQuotes() ([]byte, error) {
  response, err := http.Get("http://api.theysaidso.com/qod")
  if err != nil {
    return nil, err
  }

  defer response.Body.Close()
  return ioutil.ReadAll(response.Body)
}

func getQuoteFromContents(contents []byte) (string, error) {
  data := &Quotes{}
  err := json.Unmarshal(contents, &data)
  if err != nil {
    return "", err
  }

  return data.Contents.Quotes[0].Quote, nil
}
