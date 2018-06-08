package iaasproviderfactory

import (
	"ha-utils/ha/common/models"
	"ha-utils/ha/common/constants"
	"ha-utils/ha/common/interfaces"
	gcpproviders "ha-utils/ha/gcp/providers"
	"strings"
)

func GetProvider(config models.ConfigParams) interfaces.IIAASProvider {

	var provider interfaces.IIAASProvider

	if strings.TrimSpace(strings.ToUpper(config.Landscape)) == constants.GCP_LANDSCAPE {
		provider = &gcpproviders.GCPIAAS{}
		return provider
	} else {
		return nil
	}

}
