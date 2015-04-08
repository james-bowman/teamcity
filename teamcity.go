package teamcity

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func get(url string, resource interface{}) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	
	req.Header.Set("Accept", "application/json")
	
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	
	err = json.Unmarshal(body, &resource)
	
	if err != nil {
		myErr := fmt.Errorf("%T\n%s\n%#v\n", err, err, err)
		switch v := err.(type) {
			case *json.SyntaxError:
				myErr = fmt.Errorf("Error processing message: %s\n%s", string(body[v.Offset-40:v.Offset]), myErr)
		}
		return myErr
	}
	
	return nil
}

func ChangesBetweenBuilds(baseURL string, buildType string, firstBuildNumber string, lastBuildNumber string) ([]string, error) {
	var changes []string
	var builds map[string]interface{}
// http://segvml033.segmint.local:8080/guestAuth/app/rest/builds/?locator=buildType:Retail_RemusDocumentsService_Build,sinceBuild:2.0.189

	buildsURL := fmt.Sprintf("%s/guestAuth/app/rest/builds/?locator=buildType:%s,sinceBuild:%s", baseURL, buildType, firstBuildNumber)
	err := get(buildsURL, interface{}(&builds))
	
	fmt.Printf("%s", builds)
	
	return changes, nil
}