// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package artifactory

import (
	"os"

	artiClient "github.com/jfrog/jfrog-client-go/artifactory"
	artiAuth "github.com/jfrog/jfrog-client-go/artifactory/auth"
	artiLog "github.com/jfrog/jfrog-client-go/utils/log"
)

type (
	// Artifactory client
	Artifactory struct {
		client *artiClient.ArtifactoryServicesManager
	}
)

// New creates a new client
func New(c Config) (*Artifactory, error) {
	rtDetails := artiAuth.NewArtifactoryDetails()

	url := c.URL
	if string(url[len(url)-1:]) != "/" {
		url = c.URL + "/"
	}

	rtDetails.SetUrl(url)
	rtDetails.SetUser(c.Username)

	if c.APIKey != "" {
		rtDetails.SetApiKey(c.APIKey)
	}
	if c.Password != "" {
		rtDetails.SetPassword(c.Password)
	}

	l := artiLog.NewLogger(artiLog.INFO, os.Stdout)
	if c.Debug {
		l.SetLogLevel(artiLog.DEBUG)
	}

	serviceConfig, err := artiClient.NewConfigBuilder().
		SetArtDetails(rtDetails).
		Build()

	if err != nil {
		return nil, err
	}

	rtClient, err := artiClient.New(&rtDetails, serviceConfig)

	if err != nil {
		return nil, err
	}

	return &Artifactory{
		client: rtClient,
	}, nil
}

// Ping Artifactory
func (a Artifactory) Ping() error {
	_, err := a.client.Ping()
	return err
}