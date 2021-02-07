/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package links

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Fetch(url string) (*url.URL, string, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Cancel = make(chan struct{})
	resp, err := http.DefaultClient.Do(req)
	//resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, "", err
	}

	select {
	case <-req.Cancel:
		return nil, "", nil
	default:
		break
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("get failed, statusCode: %d\n", resp.StatusCode)
		return nil, "", errors.New(string(resp.StatusCode))
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, "", err
	}
	return resp.Request.URL, string(b), nil
}
