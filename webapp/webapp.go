package webapp

import (
  "./util"
  "net/http"
  "io/ioutil"
  "strings"
  "net/url"
  "fmt"
)

// Handle the index page, the only page in the website
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    data := util.Request("https://sulphurous.cf/api/gists/all")
    objmap := util.Parse(data)

        arr := string( * objmap["gists"])
    v := util.ArrayParse(arr)

    var resp string
    resp = ""

    for i := range v {
        de, err := url.QueryUnescape(string( * v[i]["user"]))
        util.Check(err)
        user := strings.ReplaceAll(de, `"`, "")

        tmp := string(*util.Parse(string(*util.Parse(string(*util.Parse(util.Request("https://api.scratch.mit.edu/users/" + user))["profile"]))["images"]))["55x55"])
        resp += strings.ReplaceAll(
        "<img src='" + tmp + "' height='25' width='25'>&nbsp;" +
          "<a href='https://sulphurous.cf/gists/" +
            string( * v[i]["id"]) +
                "' target='blank'>" +
                string( * v[i]["title"]) +
            " by " +
          string( * v[i]["user"]) +
        "</a>", `"`, "") +
        "<br>"
    }
    that,
    err := ioutil.ReadFile("webapp/static/index.html")
    this,
    errno := ioutil.ReadFile("webapp/static/styles.css")

    util.Check(err)
    util.Check(errno)

    fresp := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(string(that), "%info", resp), "%ldimg", util.Base64FromFile("webapp/static/loading.png")), "%css", string(this))
		fmt.Fprintf(
			w, fresp,
		)
}

func JavaScriptHandler(w http.ResponseWriter, r *http.Request) {
  that,
  err := ioutil.ReadFile("webapp/static/script.js")
  util.Check(err)
  fmt.Fprintf(
    w, string(that),
  )
}
