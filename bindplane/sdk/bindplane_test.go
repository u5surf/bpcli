package sdk

import (
	"testing"
)

func TestInit(t *testing.T) {
	var bp BindPlane

	err := bp.Init()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// make sure default url is set
	if len(bp.BaseURL) == 0 {
		t.Errorf("Expected BaseURL to be https://public-api.bindplane.bluemedora.com/v1/, got " + bp.BaseURL)
	}

	// make sure api key is read from the env and set
	if len(bp.APIKey) == 0 {
		t.Errorf("Expected APIKey to be set")
	}

	// make sure default version is set
	if len(bp.APIVersion) == 0 {
		t.Errorf("Expected APIVersion to be set")
	}

	// make sure api paths are set
	if len(bp.paths.collectors) == 0 {
		t.Errorf("Expected bp.paths.collectors to be set")
	}
	if len(bp.paths.credentialTypes) == 0 {
		t.Errorf("Expected bp.paths.credentialTypes to be set")
	}
	if len(bp.paths.credentials) == 0 {
		t.Errorf("Expected bp.paths.credentials to be set")
	}
	if len(bp.paths.jobs) == 0 {
		t.Errorf("Expected bp.paths.jobs to be set")
	}
	if len(bp.paths.sourceTypes) == 0 {
		t.Errorf("Expected bp.paths.sourceTypes to be set")
	}
	if len(bp.paths.sources) == 0 {
		t.Errorf("Expected bp.paths.sources to be set")
	}

	// modify url and make sure Init() does not set it
	bp.BaseURL = "https://dev.bluemedora.com"
	bp.APIKey = "ffae6858-1055-482a-90d0-826b20af2081" // not real

	err = bp.Init()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if bp.BaseURL != "https://dev.bluemedora.com" {
		t.Errorf("Expected BaseURL to be 'https://dev.bluemedora.com', got: " + bp.BaseURL)
	}

	if bp.APIKey != "ffae6858-1055-482a-90d0-826b20af2081" {
		t.Errorf("Expected APIKey to be 'ffae6858-1055-482a-90d0-826b20af2081', got: " + bp.APIKey)
	}

	bp.APIVersion = "v444"
	err = bp.Init()
	if err == nil {
		t.Errorf("Expected APIVersion 'v444 to cause an error', APIVersion is set to: " + bp.APIVersion)
	}
}
