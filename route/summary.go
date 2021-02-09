package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"github.com/lacazethomas/wakapi-stats/models"
)

func getToStruct(APIUrl string, target interface{}) (err error) {

	req, err := http.NewRequest(
		http.MethodGet,
		APIUrl,
		nil,
	)
	if err != nil {
		return errors.Wrap(err, "error creating HTTP request")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "error executing HTTP request")
	}

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("error sending HTTP request code error %d", res.StatusCode))
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
func GetSummary(APIURL string) (s models.Summary, err error) {
	err = getToStruct(APIURL, &s)
	return
}
