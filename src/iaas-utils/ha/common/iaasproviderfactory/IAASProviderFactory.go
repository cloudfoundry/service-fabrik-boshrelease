package iaasproviderfactory

import (
	"iaas-utils/ha/common/models"
	"iaas-utils/ha/common/constants"
	"iaas-utils/ha/common/interfaces"
	gcpproviders "iaas-utils/ha/gcp/providers"
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
