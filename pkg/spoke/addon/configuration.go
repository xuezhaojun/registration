package addon

import (
	"fmt"

	addonv1alpha1 "github.com/open-cluster-management/api/addon/v1alpha1"
)

const (
	installationNamespaceAnnotation = "addon.open-cluster-management.io/installNamespace"
)

// addonConfig contains addon configuration information.
type addOnConfig struct {
	AddOnName             string
	InstallationNamespace string
}

// getAddOnConfig returns addon configuration information from addon annotations.
func getAddOnConfig(addOn *addonv1alpha1.ManagedClusterAddOn) (*addOnConfig, error) {
	installationNamespace, ok := addOn.Annotations[installationNamespaceAnnotation]
	if !ok {
		return nil, fmt.Errorf("failed to get addon installation namesapce from addon %q", addOn.Name)
	}

	addOnConfig := &addOnConfig{
		AddOnName:             addOn.Name,
		InstallationNamespace: installationNamespace,
	}

	//TODO fill out the bootstrap secret name and registrions information

	return addOnConfig, nil
}