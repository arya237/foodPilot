package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	pkg "github.com/arya237/foodPilot/pkg/food_reserve"
	service "github.com/arya237/foodPilot/pkg/food_reserve/samad_service"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Samad struct {
	rf pkg.RequiredFunctions
}

func NewSamad(rf pkg.RequiredFunctions) *Samad {
	return &Samad{rf: rf}
}

func (s *Samad) GetAccessToken(studentNumber string, password string) (string, error) {

	baseUrl := GetTokenUrl
	const authHeader = "Basic c2FtYWQtbW9iaWxlOnNhbWFkLW1vYmlsZS1zZWNyZXQ="

	param := url.Values{}
	param.Set("username", studentNumber)
	param.Set("password", password)
	param.Set("captchaText", "")
	param.Set("validation", "undefined")
	param.Set("nonce", "")
	param.Set("grant_type", "password")
	param.Set("scope", "read write")

	client := &http.Client{}
	req, err := http.NewRequest("POST", baseUrl, strings.NewReader(param.Encode()))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status: %s, response: %s", resp.Status, string(body))
	}

	var tokenResp TokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("error parsing JSON: %v", err)
	}

	if tokenResp.AccessToken == "" {
		return "", fmt.Errorf("access token not found in response")
	}

	return tokenResp.AccessToken, nil
}

func (s *Samad) GetFoodProgram(token string, startDate time.Time) (*pkg.WeekFood, error) {
	baseURL := GetProgramUrl
	params := url.Values{}

	params.Add("selfId", "1")
	params.Add("weekStartDate", `2025-09-20 00:00:00`)

	myurl := baseURL + "?" + params.Encode()

	req, err := http.NewRequest("GET", myurl, nil)
	if err != nil {
		log.Println("line 90", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", `Bearer `+token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("line 101: ", err.Error())
		return nil, err
	}

	datas, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("line 107: ", err.Error())
		return nil, err
	}

	var income map[string]any

	err = json.Unmarshal(datas, &income)
	if err != nil {
		log.Println("line 114: ", err.Error())
		return nil, err
	}

	tmp, _ := income["payload"].(map[string]interface{})
	ProgramWeekFoodList := tmp["selfWeekPrograms"].([]interface{})

	weekFood := service.CreateWeekFood(ProgramWeekFoodList)

	return &weekFood, nil
}

func (s *Samad) ReserveFood(token string, meal pkg.ReserveModel) (string, error) {

	url := fmt.Sprintf(ReserveUrl, meal.ProgramId)

	body := fmt.Sprintf(`{"foodTypeId":%s,"mealTypeId":%s,"selectedCount":1,"freeFoodSelected":false,"selected":true}`,
		meal.FoodTypeId, meal.MealTypeId)

	req, err := http.NewRequest("PUT", url, bytes.NewBufferString(body))

	if err != nil {
		log.Println("line 133", err)
		return "", nil
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("line 144", err.Error())
		return "", err
	}

	datas, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("line 150: ", err.Error())
		return "", err
	}

	return string(datas), nil
}
