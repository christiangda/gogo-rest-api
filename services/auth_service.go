package services

import "github.com/christiangda/gogo-rest-api/models"

type authenticationService struct {
	Token []byte
}

func Authenticate(c *models.Credential) (*authenticationService, error) {

}
