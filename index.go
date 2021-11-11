package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
  	"net/url"
	"strings"
)

func request(URL string) string {
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

func parse(data string) map[string]*json.RawMessage {
	var re map[string]*json.RawMessage
	if err := json.Unmarshal([]byte(data), &re); err != nil {
		panic(err)
	}
	return re
}
func arrayparse(data string) []map[string]*json.RawMessage {
	var re []map[string]*json.RawMessage
	if err := json.Unmarshal([]byte(data), &re); err != nil {
		panic(err)
	}
	return re
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := request("https://sulphurous.cf/api/gists/all")
		objmap := parse(data)

		arr := string(*objmap["gists"])
		v := arrayparse(arr)
    fmt.Println(v)

		var resp string
		resp = ""

		for i := range v {
      de, err := url.QueryUnescape(string(*v[i]["user"]))
      check(err)
      user := strings.ReplaceAll(de, `"`, "")

      tmp := string(*parse(string(*parse(string(*parse(request("https://api.scratch.mit.edu/users/"+user))["profile"]))["images"]))["55x55"])
      fmt.Println(de, tmp)
			resp += strings.ReplaceAll(
				"<img src='"+tmp+"' height='25' width='25'>&nbsp;"+
        				"<a href='https://sulphurous.cf/gists/"+
					string(*v[i]["id"])+
					"' target='blank'>"+
					string(*v[i]["title"])+
					" by "+
					string(*v[i]["user"])+
					"</a>", `"`, "") +
				"<br>"
		}
		that, err := ioutil.ReadFile("static/index.html")
		check(err)
		fmt.Fprintf(
			w, strings.ReplaceAll(string(that), "%info", resp),
		)
	})
	log.Println("Website is up")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

