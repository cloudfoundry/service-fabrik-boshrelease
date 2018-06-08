package interfaces

import (
	"ha-helper/ha/common/models"
)

type IAuthorizationService interface {
	Initialize(...interface{})
	Authorize(authorizationRequest models.AuthorizationRequest) (*models.AuthorizationToken, bool)
}
