/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SearchIssues(s []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(s, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

func Classify(res *IssuesSearchResult) map[string][]*Issue {
	resMap := make(map[string][]*Issue)
	now := time.Now()
	for _, item := range res.Items {
		subH := now.Sub(item.CreatedAt).Hours()
		fmt.Println(item.CreatedAt)
		if subH / 24 < 0 {
			resMap["less one day"] = append(resMap["less one day"], item)
		} else if subH / (24 * 7) < 0 {
			resMap["less one week"] = append(resMap["less one week"], item)
		} else if subH / (24 * 7 * 30) < 0 {
			resMap["less one month"] = append(resMap["less one month"], item)
		} else {
			resMap["more than one month"] = append(resMap["more than one month"], item)
		}
	}
	return resMap
}
