package constants

import (
	"fmt"
	"os"
)

func InitializeVarEnviroment() {

	os.Setenv("INDEX", "mails")
	os.Setenv("URL", "http://localhost:4080/api/")
	os.Setenv("DB_USER", "admin")
	os.Setenv("DB_PASSWORD", "Complexpass#123")

	fmt.Println("Set environment variables")
}
