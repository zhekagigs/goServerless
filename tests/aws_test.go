package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

var movies []Movie

type Movie struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func TestLambdaMoviesGet(t *testing.T) {
	resp, err := http.Get("https://uggtms2bf5.execute-api.us-east-1.amazonaws.com/movies?limit=2")
	if err != nil {
		t.Errorf("error GET /movies, %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("error reading body %s", err)
	}
	if len(body) < 1 {
		fmt.Println(resp.Header)
		t.Errorf("error empty body")
	}

	err = json.Unmarshal(body, &movies)
	if err != nil {
		t.Errorf("error unmarshalling response body %s", err)
	}
	if string(`[{"id":"2","name":"Iron Man"},{"id":"13","name":"Deadpool"}]`) != (string(body)){
		t.Errorf("body doesn't match")
	}

}
