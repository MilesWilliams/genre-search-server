package search

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MilesWilliams/walker/src/middleware"
)

// Search func takes the search query and returns our api response
func Search(w http.ResponseWriter, r *http.Request) {
	query := new(Query)
	err := json.NewDecoder(r.Body).Decode(query)

	if err != nil {
		log.Fatal("Unable to parse body: ", err)
	}

	url := fmt.Sprintf("https://openwhyd.org/hot/%s?format=json&limit=100", query.Query)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create a new http request, trace %v", err.Error())
		return
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var results interface{}

	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		log.Printf("Failed to decode json response, trace %v", err.Error())
	}

	middleware.WriteJSON(w, results)

}

// Query struct
type Query struct {
	Query string `json:"query"`
}

// Results struct
type Results struct {
	ID   string `json:"_id"`
	UID  string `json:"uId"`
	UNm  string `json:"uNm"`
	Text string `json:"text"`
	Pl   struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"pl"`
	Name string        `json:"name"`
	EID  string        `json:"eId"`
	Ctx  string        `json:"ctx"`
	Img  string        `json:"img"`
	NbP  int           `json:"nbP"`
	Lov  []interface{} `json:"lov"`
}
