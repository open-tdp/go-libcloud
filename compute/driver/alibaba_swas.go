package drivers

import (
	"github.com/open-tdp/go-libcloud/compute"
)

type AlibabaSwasDriver struct {
	// add any necessary fields here
}

func (p *AlibabaSwasDriver) ListLocations() ([]*compute.Location, error) {
	// TDDO
	return nil, nil
}