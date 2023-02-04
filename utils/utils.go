package utils

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/pborman/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func MapFromBody(body io.Reader) map[string]interface{} {
	var m map[string]interface{}

	b, _ := ioutil.ReadAll(body)
	_ = json.Unmarshal(b, &m)

	return m
}

func MapFromBytes(body []byte) map[string]interface{} {
	var m map[string]interface{}

	err := json.Unmarshal(body, &m)
	if err != nil {
		fmt.Println(err.Error())
	}

	return m
}

func MapToJson(m map[string]interface{}) []byte {
	j, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	return j
}

func MapStringToJson(m map[string]string) []byte {
	j, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}

	return j
}

func MakeMapStringString(m map[string]interface{}) map[string]string {
	mss := make(map[string]string)

	for k, v := range m {
		mss[k] = v.(string)
	}

	return mss
}

func ExitIfErrExists(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ValidateSqlNullString(value sql.NullString) string {
	if value.Valid {
		return value.String
	}

	return ""
}

func MakeSlicesOfPointersAndVals(length int) ([]interface{}, [][]byte) {
	vals := make([][]byte, length)
	pointerVals := make([]interface{}, length)

	for i := range vals {
		pointerVals[i] = &vals[i]
	}

	return pointerVals, vals
}

func ToString(value interface{}) string {
	v, ok := value.(string)

	if ok {
		return string(v)
	}

	return ""
}

func ToBool(value interface{}) (bool, error) {
	v, ok := value.(bool)

	if ok {
		return bool(v), nil
	}

	return false, errors.New("Cannot convert value to boolean")
}

func ToBSON(value interface{}) (bson.M, error) {
	data, err := bson.Marshal(value)
	if err != nil {
		return nil, err
	}

	var d bson.M
	err = bson.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func IsInSlice(value string, slice []string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func DoubleQuote(s string) string {
	return "\"" + s + "\""
}

func StringInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if s == v {
			return true
		}
	}

	return false
}

func UnmarshalFrom(source interface{}, target interface{}) error {
	data, err := json.Marshal(source)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		return err
	}

	return nil
}

func GetMilliSecs() int64 {
	return (time.Now().UnixNano() / int64(time.Millisecond))
}

func GetMilliSecsForTime(t time.Time) int64 {
	return (t.UnixNano() / int64(time.Millisecond))
}

// Taken from github.com/mattermost/mattermost-server/utils
func GetIPAddress(req *http.Request) string {
	address := ""

	header := req.Header.Get("X-Forwarded-For")
	if len(header) > 0 {
		addresses := strings.Fields(header)
		if len(addresses) > 0 {
			address = strings.TrimRight(addresses[0], ",")
		}
	}

	if len(address) == 0 {
		address = req.Header.Get("X-Real-IP")
	}

	if len(address) == 0 {
		address, _, _ = net.SplitHostPort(req.RemoteAddr)
	}

	return address
}

func GetExecutableDirPath() string {
	_, file, _, _ := runtime.Caller(1)
	dirPath := path.Dir(file)

	return dirPath
}

func NewID() string {
	var encodingString = "hij58kAB3p4CRSglTUbVWXYZG67HxyIJKLMNOPQqDEFrsacdefmnotuvwz0129-_"
	var b bytes.Buffer
	var encoding = base64.NewEncoding(encodingString)
	encoder := base64.NewEncoder(encoding, &b)
	_, _ = encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(22)
	return b.String()
}

func NewLongID() string {
	return NewID() + NewID() + NewID()
}

func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
