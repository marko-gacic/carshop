package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"carshop/pkg/logger"

	"go.uber.org/zap"
)

// if you are using this file, you are smart :)
// this is why I will help you
// To run tests, use these commands:

// go run test.go --action "auth" --subAction "login"
// go run test.go --action "create"
// go run test.go --action "read"
// go run test.go --action "list"
// go run test.go --action "delete"

var (
	URL string
)

func init() {
	URL = "http://127.0.0.1:8000/"
}

func main() {
	log := logger.Log.WithOptions()

	actionFlag := flag.String("action", "", "choose action")
	fmt.Println("actionFlag: ", *actionFlag)

	subActionFlag := flag.String("subAction", "", "choose action")
	fmt.Println("subActionFlag: ", *subActionFlag)

	flag.Parse()

	switch *actionFlag {
	case "create":
		file, err := os.ReadFile(fmt.Sprintf("%v.json", *actionFlag))
		if err != nil {
			err := fmt.Errorf("file %v does not exist", *actionFlag)
			panic(err)
		}

		data := make(map[string]any)

		err = json.Unmarshal(file, &data)
		if err != nil {
			log.Info("failed to unmarshal data")
		}

		err = request(data, "POST", "car/"+*actionFlag, log, true)
		if err != nil {
			log.Info("failed to perform request",
				zap.Error(err))
		}
	case "read":
		file, err := os.ReadFile(fmt.Sprintf("%v.json", *actionFlag))
		if err != nil {
			err := fmt.Errorf("file %v does not exist", *actionFlag)
			panic(err)
		}

		data := make(map[string]any)

		err = json.Unmarshal(file, &data)
		if err != nil {
			log.Info("failed to unmarshal data")
		}

		err = request(data, "POST", "car/"+*actionFlag, log, true)
		if err != nil {
			log.Info("failed to perform request",
				zap.Error(err))
		}
	case "list":

		file, err := os.ReadFile(fmt.Sprintf("%v.json", *actionFlag))
		if err != nil {
			err := fmt.Errorf("file %v does not exist", *actionFlag)
			panic(err)
		}

		data := make(map[string]any)

		err = json.Unmarshal(file, &data)
		if err != nil {
			log.Info("failed to unmarshal data")
		}

		err = request(data, "POST", "car/"+*actionFlag, log, true)
		if err != nil {
			log.Info("failed to perform request",
				zap.Error(err))
		}

	case "delete":
		file, err := os.ReadFile(fmt.Sprintf("%v.json", *actionFlag))
		if err != nil {
			err := fmt.Errorf("file %v does not exist", *actionFlag)
			panic(err)
		}

		data := make(map[string]any)

		err = json.Unmarshal(file, &data)
		if err != nil {
			log.Info("failed to unmarshal data")
		}

		err = request(data, "POST", "car/"+*actionFlag, log, true)
		if err != nil {
			log.Info("failed to perform request",
				zap.Error(err))
		}
	case "auth":
		file, err := os.ReadFile(fmt.Sprintf("%v/%v.json", *actionFlag, *subActionFlag))
		if err != nil {
			err := fmt.Errorf("file %v does not exist", *actionFlag)
			panic(err)
		}

		data := make(map[string]any)

		err = json.Unmarshal(file, &data)
		if err != nil {
			log.Info("failed to unmarshal data")
		}

		err = request(data, "POST", *actionFlag+"/"+*subActionFlag, log, true)
		if err != nil {
			log.Info("failed to perform request",
				zap.Error(err))
		}

	}
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

var count = 0

func request(data map[string]any, method, action string, logger *zap.Logger, setTokens bool) (err error) {

	var (
		tokens Token
	)

	raw, err := json.Marshal(data)
	if err != nil {
		return
	}

	req, err := http.NewRequest(method, URL+action, bytes.NewBuffer(raw))
	if err != nil {
		return
	}

	if setTokens {
		file, err := os.ReadFile("auth/tokens.json")
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(file, &tokens)
		if err != nil {
			panic(err)
		}

		req.Header.Set("Auth-Access-Token", tokens.AccessToken)
	}

	//fmt.Printf("\nrequest body: %+v\n", req.Body)
	req.Header.Set("Content-Type", "application/json")

	now := time.Now()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	//fmt.Println("response Status:", resp.Status)
	fmt.Println(time.Since(now))
	accessToken := resp.Header.Get("Auth-Access-Token")
	if accessToken != "" {
		fmt.Println("response Access token:", accessToken)

		data := map[string]string{
			"accessToken": accessToken,
		}

		raw, _ := json.Marshal(data)
		ioutil.WriteFile("auth/tokens.json", raw, 0744)
	}

	response, _ := ioutil.ReadAll(resp.Body)
	count++
	fmt.Println("response", string(response))
	return
}
