package samad

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/arya237/foodPilot/pkg/logger"
	"github.com/arya237/foodPilot/pkg/reservations"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Samad struct {
	*Config
	logger logger.Logger
}

func NewSamad(conf *Config) reservations.ReserveFunctions {
	return &Samad{
		Config: conf,
		logger: logger.New("samadService"),
	}
}

func (s *Samad) GetProperSelfID(token string) (map[string]int, error) {
	URL := s.GetSelfIDUrl

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		s.logger.Info(err.Error())
		return nil, reservations.ErrorInternal
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Info(err.Error())
		return nil, reservations.ErrorInternal
	}

	defer resp.Body.Close()
	datas, err := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		s.logger.Info(string(datas))
		return nil, reservations.ErrorSamadReserveation
	}
	if err != nil {
		s.logger.Info(err.Error())
		return nil, err
	}

	// log.Println(string(datas), "\nNigggerrr\n\n", resp)

	var income map[string]any

	err = json.Unmarshal(datas, &income)
	if err != nil {
		s.logger.Info(err.Error())
		return nil, err
	}

	tmp, _ := income["payload"].([]interface{})
	selfIDs := make(map[string]int)

	for _, key := range tmp {
		new := key.(map[string]interface{})
		if new["name"] == "مرکزی برادران" || new["name"] == "مرکزی خواهران" ||
			new["name"] == "رستوران مکمل 1 (آشکده)برادران" || new["name"] == "رستوران مکمل 1 (آشکده)خواهران" {

			selfIDs[new["name"].(string)] = int(new["id"].(float64))
		}
	}

	if len(selfIDs) != 0 {
		return selfIDs, nil
	}

	return nil, reservations.ErrorInternal
}

func (s *Samad) GetAccessToken(studentNumber string, password string) (string, error) {

	baseUrl := s.GetTokenUrl
	authHeader := s.AuthHeader

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

func (s *Samad) GetFoodProgram(token string, selfID int, startDate time.Time) (*reservations.WeekFood, error) {

	self := strconv.Itoa(selfID)

	baseURL := s.GetProgramUrl
	params := url.Values{}

	params.Add("selfId", self)
	params.Add("weekStartDate", startDate.Format("2006-01-02 00:00:00"))

	myurl := baseURL + "?" + params.Encode()

	req, err := http.NewRequest("GET", myurl, nil)
	if err != nil {
		s.logger.Info(err.Error())
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", `Bearer `+token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		s.logger.Info(err.Error())
		return nil, err
	}

	datas, err := io.ReadAll(resp.Body)

	if err != nil {
		s.logger.Info(err.Error())
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		s.logger.Info(string(datas))
		return nil, errors.New("failed to get food program")
	}

	var income map[string]any

	err = json.Unmarshal(datas, &income)
	if err != nil {
		s.logger.Info(err.Error())
		return nil, err
	}

	tmp, _ := income["payload"].(map[string]interface{})
	ProgramWeekFoodList := tmp["selfWeekPrograms"].([]interface{})

	weekFood := CreateWeekFood(ProgramWeekFoodList)

	return &weekFood, nil
}

func (s *Samad) ReserveFood(token string, meal reservations.ReserveModel) (string, error) {

	url := fmt.Sprintf(s.ReserveUrl, meal.ProgramId)

	body := fmt.Sprintf(`{"foodTypeId":%s,"mealTypeId":%s,"selectedCount":1,"freeFoodSelected":false,"selected":true}`,
		meal.FoodTypeId, meal.MealTypeId)

	req, err := http.NewRequest("PUT", url, bytes.NewBufferString(body))

	if err != nil {
		s.logger.Info(err.Error())
		return "", nil
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		s.logger.Info(err.Error())
		return "", err
	}

	datas, err := io.ReadAll(resp.Body)

	if err != nil {
		s.logger.Info(err.Error())
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		s.logger.Info(string(datas))
		return "", errors.New(string(datas))
	}

	return string(datas), nil
}
