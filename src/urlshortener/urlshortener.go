package urlshortener

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

var key = os.Getenv("GOOGL_API_KEY")
var apiURL = "https://www.googleapis.com/urlshortener/v1/url" + "?key=" + key
var client = &http.Client{}

type data struct {
	LongURL string `json:"longUrl"`
}

func Shorten(lURL string) []byte {
	j, _ := json.Marshal(data{lURL})
	req, _ := http.NewRequest("POST", apiURL, bytes.NewReader(j))
	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)

	var t []byte
	t, _ = ioutil.ReadAll(resp.Body)

	return t

}
