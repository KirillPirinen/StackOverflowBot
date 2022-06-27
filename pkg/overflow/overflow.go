package overflow

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type overQuestion struct {
	Title string
	Link string
	Tags []string
}

type overResp struct {
	Items []overQuestion
	Has_more bool
	Quota_max int
	Quota_remaining int
}

const (
	host = "https://api.stackexchange.com/2.3/search"
	defQuery = "?site=stackoverflow&order=desc&sort=relevance&filter=default&answers=1"
)


func SearchThroughOverflow(query string, tag string) (ans overResp) {
	
	parsedQuery := strings.Join(strings.Fields(query), "%")
	titleSearch := fmt.Sprintf("&intitle=%v", parsedQuery)

	if tag != "" {
		tag = fmt.Sprintf("&tagged=%v", tag)
	}

	resp, err := http.Get(host + defQuery + titleSearch + tag)
	body, err := io.ReadAll(resp.Body)

	json.Unmarshal(body, &ans)

	if err != nil {
			panic(err)
	}
	
	defer resp.Body.Close()

	return
}
