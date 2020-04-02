package cli

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

type HttpCliMgr struct {
	client *http.Client
}

var GlobHttpCliMgr *HttpCliMgr

//Init Http Client
func InitHttpCliMgr() (err error) {
	var httpClinet *http.Client
	//var transport *http.Transport

/*	transport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 4 * time.Second,
		}).Dial,
	}
*/
	httpClinet = &http.Client{
/*		Transport: transport,
		Timeout:   5 * time.Second,*/
	}

	GlobHttpCliMgr = &HttpCliMgr{client: httpClinet}

	return
}

func (cli *HttpCliMgr) RequestWithUrl(method string, ur string, pdata []byte) (body []byte, err error) {
	var request *http.Request
	var response *http.Response

	if request, err = http.NewRequest(method, ur, bytes.NewBuffer(pdata)); err != nil {
		return
	}
	request.Header.Set(`Content-Type`, `application/json`)

	if response, err = cli.client.Do(request); err != nil {
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		if body, err = ioutil.ReadAll(response.Body); err != nil {
			return
		}
		return

	}
	if response.StatusCode != 200 {
		err = errors.New(`request error`)
		return
	}
	return
}
