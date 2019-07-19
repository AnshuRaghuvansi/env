package env

import (
	"fmt"
	"os"
	"testing"
)

func TestConfigure(t *testing.T) {

	os.Setenv("environment", "development")

	Configure()

	port := GetValue("PORT")
	fmt.Println(port)

	prod, ok := GetBoolValue("PRODUCTION")
	if ok {
		fmt.Println(prod)
	}

	url := GetValue("DB_URL")
	fmt.Println(url)
}
