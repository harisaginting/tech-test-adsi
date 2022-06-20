package helper

import "os"
import "errors"
import "strconv"
import "encoding/json"
import "math/big"
import "encoding/base64"

// MustGetEnv get environment value
func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return ""
	}
	return value
}

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func AdjustStructToStruct(a interface{},b interface{}) interface{} {
	JsonStruct, _ := json.Marshal(a)
	json.Unmarshal([]byte(JsonStruct), &b)
	return b
}

func ForceInt(v interface{}) int {
	var result int
	switch v.(type) {
	case int:
		result = v.(int)
	case float64:
		result = int(v.(float64))
	case string:
		result, _ = strconv.Atoi(v.(string))
	}
	return result
}

func ForceString(v interface{}) string {
	var result string
	switch v.(type) {
	case int:
		result = strconv.Itoa(v.(int))
	case float64:
		result = strconv.Itoa(int(v.(float64)))
	case string:
		result, _ = v.(string)
	}
	return result
}

func ForceError(v interface{}) error {
	result := errors.New(ForceString(v))
	return result
}


func DecodeBase64BigInt(s string) *big.Int {
	buffer, _ := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(s)
	return big.NewInt(0).SetBytes(buffer)
}