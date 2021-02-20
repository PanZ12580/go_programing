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
)

func Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("get failed, statusCode: %d\n", resp.StatusCode)
		return "", errors.New(string(resp.StatusCode))
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(b), nil
}
