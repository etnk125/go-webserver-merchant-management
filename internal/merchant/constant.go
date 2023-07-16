package merchant

import (
	"fmt"

	"github.com/etnk125/go-webserver-merchant-management/model"
)

func ErrMerchantNotFound() error {
	return fmt.Errorf("merchant not found")
}

func DefaultCredential() *model.Credential {
	return &model.Credential{
		Username: "default_username",
		Password: "default_password",
	}
}
