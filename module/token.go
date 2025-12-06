package SO_Module

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

type tokenservice struct{}

func NewToken() *tokenservice {
	return &tokenservice{}
}

// TODO : tambah parameter untuk token : version, dbstring

type TokenInterface interface {
	CreateToken(userId string, database string, version string, versionId string, deviceId string, loginId string) (*TokenDetails, error)
	ExtractTokenMetadata(*gin.Context) (*AccessDetails, error)
	RefreshToken(token string) (*jwt.Token, error)
}

// Token implements the TokenInterface
var _ TokenInterface = &tokenservice{}

func (t *tokenservice) CreateToken(userId string, database string, version string, versionId string, deviceId string, loginId string) (*TokenDetails, error) {
	td := &TokenDetails{}

	access_ttl := os.Getenv("ACCESS_TTL")
	no_access_ttl, err_a := strconv.Atoi(access_ttl)
	if err_a != nil {
		fmt.Println("Error converting string to integer:", err_a)
		no_access_ttl = 30
	}

	refresh_ttl := os.Getenv("REFRESH_TTL")
	no_refresh_ttl, err_r := strconv.Atoi(refresh_ttl)
	if err_r != nil {
		fmt.Println("Error converting string to integer:", err_r)
		no_access_ttl = 7
	}

	// fmt.Println("no_access_ttl:", no_access_ttl)
	// fmt.Println("time.Duration(no_access_ttl):", time.Duration(no_access_ttl))
	// fmt.Println("no_refresh_ttl:", no_refresh_ttl)

	// td.AtExpires = time.Now().Add(time.Minute * 30).Unix() //expires after 30 min
	// td.AtExpires = time.Now().Add(time.Minute * time.Duration(no_access_ttl)).Unix() //expires after 30 min
	td.AtExpires = time.Now().Add(time.Second * time.Duration(no_access_ttl)).Unix() //expires after 30 min
	td.TokenUuid = uuid.NewV4().String()

	// td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	// td.RtExpires = time.Now().Add(time.Hour * 24 * time.Duration(no_refresh_ttl)).Unix()
	td.RtExpires = time.Now().Add(time.Second * 60 * time.Duration(no_refresh_ttl)).Unix()
	td.RefreshUuid = td.TokenUuid + "++" + userId
	td.Database = database
	td.LoginId = loginId
	td.AppId = os.Getenv("APP_ID")
	td.Version = version
	td.VersionId = versionId

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.TokenUuid
	atClaims["user_id"] = userId
	atClaims["exp"] = td.AtExpires
	atClaims["database"] = td.Database
	atClaims["login_id"] = td.LoginId
	atClaims["app_id"] = td.AppId
	atClaims["version"] = td.Version
	atClaims["version_id"] = td.VersionId
	atClaims["device_id"] = deviceId
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.RtExpires
	rtClaims["database"] = td.Database
	rtClaims["login_id"] = td.LoginId
	rtClaims["app_id"] = td.AppId
	rtClaims["version"] = td.Version
	rtClaims["version_id"] = td.VersionId
	rtClaims["device_id"] = deviceId
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func TokenValid(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// get the token from the request body
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func extract(token *jwt.Token) (*AccessDetails, error) {

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		userId, userOk := claims["user_id"].(string)
		loginId, loginOk := claims["login_id"].(string)
		appId, appOk := claims["app_id"].(string)
		deviceId, deviceIdOk := claims["device_id"].(string)
		dbschema, dbschemaOk := claims["database"].(string)
		if !ok || !userOk || !loginOk || !appOk || !dbschemaOk || !deviceIdOk {
			return nil, errors.New("unauthorized")
		} else {
			return &AccessDetails{
				TokenUuid: accessUuid,
				UserId:    userId,
				AppId:     appId,
				LoginId:   loginId,
				Database:  dbschema,
				DeviceId:  deviceId,
			}, nil
		}
	}
	return nil, errors.New("something went wrong")
}

func (t *tokenservice) ExtractTokenMetadata(c *gin.Context) (*AccessDetails, error) {
	token, err := verifyToken(c.Request)
	if err != nil {
		return nil, err
	}
	access, err := extract(token)
	if err != nil {
		return nil, err
	}

	return access, nil
}

func (t *tokenservice) RefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := verifyRefreshToken(tokenString)
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, err
	}
	return token, nil
}

func verifyRefreshToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
