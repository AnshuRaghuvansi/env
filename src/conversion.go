package env

import (
	"os"
	"strconv"
)

//GetValue :
func GetValue(key string) string {
	return os.Getenv(key)
}

//GetIntValue :
func GetIntValue(key string) (int, bool) {
	value := os.Getenv(key)
	val, err := strconv.Atoi(value)
	if err != nil {
		return val, false
	}
	return val, true
}

//GetInt8Value :
func GetInt8Value(key string) (int8, bool) {
	value := os.Getenv(key)
	val, err := strconv.ParseInt(value, 10, 8)
	if err != nil {
		return int8(val), false
	}
	return int8(val), true
}

//GetInt16Value :
func GetInt16Value(key string) (int16, bool) {
	value := os.Getenv(key)
	val, err := strconv.ParseInt(value, 10, 16)
	if err != nil {
		return int16(val), false
	}
	return int16(val), true
}

//GetInt32Value :
func GetInt32Value(key string) (int32, bool) {
	value := os.Getenv(key)
	val, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		return int32(val), false
	}
	return int32(val), true
}

//GetInt64Value :
func GetInt64Value(key string) (int64, bool) {
	value := os.Getenv(key)
	val, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return int64(val), false
	}
	return int64(val), true
}

//GetFloat32Value :
func GetFloat32Value(key string) (float32, bool) {
	value := os.Getenv(key)
	val, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return float32(val), false
	}
	return float32(val), true
}

//GetFloat64Value :
func GetFloat64Value(key string) (float64, bool) {
	value := os.Getenv(key)
	val, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return float64(val), false
	}
	return float64(val), true
}

//GetBoolValue :
func GetBoolValue(key string) (bool, bool) {
	value := os.Getenv(key)
	val, err := strconv.ParseBool(value)
	if err != nil {
		return val, false
	}
	return val, true
}
