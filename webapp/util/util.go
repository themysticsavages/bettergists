package util

import (
  "io/ioutil"
  "net/http"
  "fmt"
  "encoding/base64"
  "encoding/json"
)

func Request(URL string) string {
    resp, err := http.Get(URL)
    if err != nil {
        fmt.Println(err)
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    sb := string(body)
    return sb
}
func Base64FromFile(f string) string {
  bytes, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
	}

	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += base64.StdEncoding.EncodeToString(bytes)
  return base64Encoding
}

func Parse(data string) map[string] * json.RawMessage {
    var re map[string] * json.RawMessage
    if err := json.Unmarshal([] byte(data), & re);
    err != nil {
        panic(err)
    }
    return re
}
func ArrayParse(data string)[] map[string] * json.RawMessage {
    var re[] map[string] * json.RawMessage
    if err := json.Unmarshal([] byte(data), & re);
    err != nil {
        panic(err)
    }
    return re
}

func Check(e error) {
    if e != nil {
        panic(e)
    }
}
