package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Consts ...
const (
	LoginURL       = "http://common.iyouqu.com.cn"
	SignURL        = "http://iyouqu.com.cn"
	Port           = "8080"
	LoginPath      = "/app/user/service.do"
	SignPath       = "/app/sign/service.do"
	YqVersion      = "v2.3.0"
	YqPlatform     = "ios"
	System         = "12.0"
	Device         = "iPhone X"
	RegistrationID = "111a010c6310a7dd5bc6aa9f19439d84"
	SystemType     = 2
	Version        = "v2.3.0"
	MsgID          = "APP129"
	City           = "武汉市"
	Country        = "中国"
	Position       = "在武汉市花城大道签到啦！"
	Longitude      = 114.51166703172403
	Imei           = "0034666FACBF44F4BA2302763BD086B7"
	Latitude       = 30.560763466475613
	MsgIDSign      = "APP_SIGN"
	GroupID        = 2004464
	Province       = "湖北省"
)

// LoginBody ...
type LoginBody struct {
	System         string `json:"system"`
	Password       string `json:"password"`
	Mobile         string `json:"mobile"`
	Device         string `json:"device"`
	RegistrationID string `json:"registrationId"`
	SystemType     int    `json:"systemType"`
	Version        string `json:"version"`
	MsgID          string `json:"msgId"`
}

// Login In ...
func Login(mobile, password string) (map[string]interface{}, error) {
	loginBody := LoginBody{
		System:         System,
		Password:       password,
		Mobile:         mobile,
		Device:         Device,
		RegistrationID: RegistrationID,
		SystemType:     SystemType,
		Version:        Version,
		MsgID:          MsgID,
	}

	data, err := json.Marshal(loginBody)
	if err != nil {
		return nil, err
	}
	res, err := http.PostForm(LoginURL+":"+Port+LoginPath, url.Values{"text": {string(data)}})
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resBodyMap := make(map[string]interface{})
	if err := json.Unmarshal(body, &resBodyMap); err != nil {
		return nil, err
	}

	return resBodyMap, nil
}

// SignInBody ...
type SignInBody struct {
	UserID    int64   `json:"userId"`
	City      string  `json:"city"`
	UserName  string  `json:"userName"`
	Country   string  `json:"country"`
	Position  string  `json:"position"`
	Longitude float64 `json:"longitude"`
	Imei      string  `json:"imei"`
	Latitude  float64 `json:"latitude"`
	MsgID     string  `json:"msgId"`
	GroupID   int64   `json:"groupId"`
	Province  string  `json:"province"`
}

// SignIn ...
func SignIn(id int64, employee *Employee) (map[string]interface{}, error) {
	signInBody := SignInBody{
		UserID:    employee.EmployeeID,
		City:      City,
		UserName:  employee.Name,
		Country:   Country,
		Position:  Position,
		Longitude: Longitude,
		Imei:      Imei,
		Latitude:  Latitude,
		MsgID:     MsgIDSign,
		GroupID:   GroupID,
		Province:  Province,
	}

	data, err := json.Marshal(signInBody)
	if err != nil {
		return nil, err
	}
	c := http.DefaultClient
	reader := url.Values{"text": {string(data)}}
	req, err := http.NewRequest("POST", SignURL+":"+Port+SignPath, strings.NewReader(reader.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("YQ-Token", employee.UserToken)
	req.Header.Set("YQ-Platform", YqPlatform)
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	resBodyMap := make(map[string]interface{})
	if err := json.Unmarshal(body, &resBodyMap); err != nil {
		return nil, err
	}
	return resBodyMap, nil
}
