package pages

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	m "vsys.empms.commons/models"
	u "vsys.empms.commons/utils"
)

type Logs struct {
	name string
	log  []m.Log
}

func FetchLogs() ([]m.Log, error) {
	if strings.TrimSpace(webPort) == "" {
		webPort = "3200"
	}
	if strings.TrimSpace(webHost) == "" {
		webHost = "0.0.0.0"
	}
	url := fmt.Sprint("http://", webHost, ":", webPort, "/get-logs")
	res, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(res)
	if err != nil {
		log.Fatal("NO infoooo")
		return nil, err
	}
	defer resp.Body.Close()

	var log []m.Log
	json.NewDecoder(resp.Body).Decode(&log)
	return log, nil
}

func (l *Logs) BuildLogBody() string {
	var HTMLBody string
	for _, log := range l.log {
		HTMLBody += fmt.Sprint(`<a class="list-group-item list-group-item-action">`, log.Created, `&nbsp;&nbsp;&nbsp`, log.UpdatedBy, `&nbsp;&nbsp;&nbsp`, log.Operation, `&nbsp;&nbsp;&nbsp{"`, log.EmpName, `","`, log.EmpEmail, `","`, log.EmpDeg, `"} </a>`)
	}
	return HTMLBody
}

func (l *Logs) Build() string {
	l.log, _ = FetchLogs()
	l.name = "Logs"
	var html string
	html += fmt.Sprint(`
<!-- List -->
<div class="list-group" style="user-select: none;">
    <a class="list-group-item list-group-item-action active h1 m-0 " href="#">`, l.name, `</a>
    `, l.BuildLogBody(), `
</div>
<!-- List -->
	`)

	return u.JoinStr(html)
}
