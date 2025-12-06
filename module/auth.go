package SO_Module

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v7"
)

type AuthInterface interface {
	CreateAuth(string, *TokenDetails) error
	FetchAuth(string) (string, error)
	DeleteRefresh(string) error
	DeleteTokens(*AccessDetails) error
}

type service struct {
	client *redis.Client
}

var _ AuthInterface = &service{}

func NewAuth(client *redis.Client) *service {
	return &service{client: client}
}

type AccessDetails struct {
	TokenUuid string
	AppId     string
	UserId    string
	LoginId   string
	Database  string
	DeviceId  string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUuid    string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
	Database     string
	Version      string
	VersionId    string
	LoginId      string
	AppId        string
}

// Save token metadata to Redis
func (tk *service) CreateAuth(userId string, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()
	appname := os.Getenv("APP_NAME")
	atCreated, err := tk.client.Set(appname+td.TokenUuid, userId, at.Sub(now)).Result()
	if err != nil {
		return err
	}
	rtCreated, err := tk.client.Set(appname+td.RefreshUuid, userId, rt.Sub(now)).Result()
	if err != nil {
		return err
	}
	if atCreated == "0" || rtCreated == "0" {
		return errors.New("no record inserted")
	}
	return nil
}

// Check the metadata saved
func (tk *service) FetchAuth(tokenUuid string) (string, error) {
	appname := os.Getenv("APP_NAME")
	userid, err := tk.client.Get(appname + tokenUuid).Result()
	fmt.Println(appname + tokenUuid)
	if err != nil {
		return "", err
	}
	return userid, nil
}

// Once a user row in the token table
func (tk *service) DeleteTokens(authD *AccessDetails) error {
	appname := os.Getenv("APP_NAME")
	//get the refresh uuid
	refreshUuid := fmt.Sprintf("%s++%s", authD.TokenUuid, authD.UserId)
	//delete access token
	deletedAt, err := tk.client.Del(appname + authD.TokenUuid).Result()
	if err != nil {
		return err
	}
	//delete refresh token
	deletedRt, err := tk.client.Del(appname + refreshUuid).Result()
	if err != nil {
		return err
	}
	//When the record is deleted, the return value is 1
	if deletedAt != 1 || deletedRt != 1 {
		return errors.New("something went wrong")
	}
	return nil
}

func (tk *service) DeleteRefresh(refreshUuid string) error {
	appname := os.Getenv("APP_NAME")
	//delete refresh token
	deleted, err := tk.client.Del(appname + refreshUuid).Result()
	if err != nil || deleted == 0 {
		return err
	}
	return nil
}
