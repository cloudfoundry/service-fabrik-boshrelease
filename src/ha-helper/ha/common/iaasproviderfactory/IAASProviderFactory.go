package iaasproviderfactory

import (
	gcpproviders "ha-helper/ha/gcp/providers"	
	"ha-helper/ha/common/beans"
	"ha-helper/ha/common/constants"
	"ha-helper/ha/common/interfaces"
	"strings"
)

func GetProvider(config beans.ConfigParams) interfaces.IIAASProvider {

	var provider interfaces.IIAASProvider

	if strings.TrimSpace(strings.ToUpper(config.Landscape)) == constants.GCP_LANDSCAPE { 	
		provider = &gcpproviders.GCPIAAS{}
		return provider
	} else {
		return nil
	}

}
