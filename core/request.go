package core

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/imroc/req/v3"
	"nacos/core/flag"
	"strings"
	"time"
)

var client *req.Client

var token string

func InitRequest() {
	client = req.C()
	client.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:105.0) Gecko/20100101 Firefox/105.0")
	client.SetTimeout(time.Duration(5) * time.Second)
	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: false,
	}).SetTLSHandshakeTimeout(time.Duration(5) * time.Second)
	client.SetCommonContentType("application/x-www-form-urlencoded")

	var err error
	token, err = generateToken()
	if err != nil {
		fmt.Println("generate token failed, error: ", err)
		fmt.Println("try use serverIdentity to bypass auth")
		client.SetCommonHeader("serverIdentity", "security")
	} else {
		client.SetCommonHeader("accessToken", token)
	}

	if flag.Proxy != "" {
		client.SetProxyURL(flag.Proxy)
	}
	client.SetBaseURL(flag.Target + "/nacos")
}

type ConfigItem struct {
	Id               string      `json:"id"`
	DataId           string      `json:"dataId"`
	Group            string      `json:"group"`
	Content          string      `json:"content"`
	Md5              interface{} `json:"md5,omitempty"`
	EncryptedDataKey interface{} `json:"encryptedDataKey,omitempty"`
	Tenant           string      `json:"tenant"`
	AppName          string      `json:"appName"`
	Type             string      `json:"type"`
}

type ConfigResponse struct {
	TotalCount     int          `json:"totalCount"`
	PageNumber     int          `json:"pageNumber"`
	PagesAvailable int          `json:"pagesAvailable"`
	PageItems      []ConfigItem `json:"pageItems"`
}

type NamespaceItem struct {
	Namespace         string  `json:"namespace"`
	NamespaceShowName string  `json:"namespaceShowName,omitempty"`
	NamespaceDesc     *string `json:"namespaceDesc"`
	Quota             int     `json:"quota"`
	ConfigCount       int     `json:"configCount"`
	Type              int     `json:"type"`
}

type NamespaceResp struct {
	Code    int             `json:"code"`
	Message interface{}     `json:"message,omitempty"`
	Data    []NamespaceItem `json:"data"`
}

func generateToken() (string, error) {
	// default Key SecretKey012345678901234567890123456789012345678901234567890123456789
	claims := jwt.MapClaims{
		"sub": "nacos",
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if flag.Key == "" {
		flag.Key = "SecretKey012345678901234567890123456789012345678901234567890123456789"
	}
	byteKey, err := base64.StdEncoding.DecodeString(flag.Key)
	if err != nil && !strings.Contains(err.Error(), "illegal base64 data at input") {
		return "", err
	}
	signToken, err := token.SignedString(byteKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return signToken, nil
}
