package controllers

import (
	"github.com/revel/revel"
	"github.com/randomtask1155/hqserver/device"
	"os"
	"net/http"
	"crypto/tls"
	"time"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

var (
	hqAddress string 
	hqClient *http.Client
)


func init() {
	hqAddress = os.Getenv("HQSERVER_ADDRESS")
	tr := &http.Transport{
		MaxIdleConns:       -1,
		IdleConnTimeout:    30 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	hqClient = &http.Client{Transport: tr,
			Timeout: 10 * time.Second,
	}

}

type JsonErr struct {
	ErrorStr string `json:"error"`
}

func jsonError(e error) JsonErr {
	return JsonErr{e.Error()}
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}


func (c App) Status() revel.Result {
	log := c.Log.New("id","hqweb")
	res,err := hqClient.Get(hqAddress)
	if err != nil {
		log.Errorf("failed to get status form server", err.Error())
		return c.RenderJSON(jsonError(err))
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Errorf("Bad response from server: HTTP Status %d", res.StatusCode)
		return c.RenderJSON(fmt.Sprintf("Bad response from server: HTTP Status %d", res.StatusCode))
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("failed to read response body form server", err.Error())
		return c.RenderJSON(jsonError(err))
	}
	status := make([]device.DeviceStatus,0)
	err = json.Unmarshal(b, &status)
	if err != nil {
		log.Errorf("failed to unmarshal response", err.Error())
		return c.RenderJSON(jsonError(err))
	}
	return c.RenderJSON(status)
}