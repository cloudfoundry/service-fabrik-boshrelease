package interfaces

import (
	"iaas-utils/ha/common/models"
)

type IAuthorizationService interface {
	Initialize(...interface{})
	Authorize(authorizationRequest models.AuthorizationRequest) (*models.AuthorizationToken, bool)
}
