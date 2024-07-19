package utils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/miltsm/hubung-service/internal/types"
)

func SubstringEmailPreAlias(e string) (string, error) {
	if len(e) == 0 {
		return "", &types.RequireEmailError{}
	}
	for idx, r := range e {
		if r == '@' {
			sub := e[:idx-1]
			fmt.Println("return before alias")
			return sub, nil
		}
	}
	return "", &types.RequireEmailError{}
}

func GenerateUUID() string {
	//docker issue
	// byte, err := exec.Command("uuidgen").Output()
	// if err != nil {
	// 	log.Fatal(err)
	// 	return ""
	// }
	// uuid := string(byte)
	uuid := uuid.New().String()
	return uuid
}
