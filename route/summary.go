package route

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"github.com/lacazethomas/wakapi-stats/config"
	"github.com/lacazethomas/wakapi-stats/models"
)

func getToStruct(APIUrl string, APIKEY string, target interface{}) (err error) {

	req, err := http.NewRequest(
		http.MethodGet,
		APIUrl,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "error creating HTTP request")
	}

	q := req.URL.Query()
	q.Add("interval", "30_days")
	req.URL.RawQuery = q.Encode()

	token := base64.StdEncoding.EncodeToString([]byte(APIKEY))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "error sending HTTP request")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &target)
	if err != nil {
		return errors.Wrap(err, "error unmarshaling body to target struct")
	}

	return
}

//GetSummary from Wakapki
func GetSummary(cfg config.AppConfig) (s models.Summary) {
	getToStruct(cfg.APIURL+config.WakapiSummary, cfg.APIKey, &s)
	return s
}
