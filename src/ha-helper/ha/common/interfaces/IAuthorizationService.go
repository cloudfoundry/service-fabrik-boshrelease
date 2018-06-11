package interfaces

import (
	"ha-helper/ha/common/beans"
)

type IAuthorizationService interface {
	Initialize(...interface{})
	Authorize(authorizationRequest beans.AuthorizationRequest) (*beans.AuthorizationToken, bool)
}
