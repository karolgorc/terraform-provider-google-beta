// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package firebaserules

import (
	"context"
	"log"
	"testing"

	firebaserules "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebaserules/beta"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/envvar"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/sweeper"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func init() {
	sweeper.AddTestSweepersLegacy("FirebaserulesRelease", testSweepFirebaserulesRelease)
}

func testSweepFirebaserulesRelease(region string) error {
	log.Print("[INFO][SWEEPER_LOG] Starting sweeper for FirebaserulesRelease")

	config, err := sweeper.SharedConfigForRegion(region)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
		return err
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
		return err
	}

	t := &testing.T{}
	billingId := envvar.GetTestBillingAccountFromEnv(t)

	// Setup variables to be used for Delete arguments.
	d := map[string]string{
		"project":         config.Project,
		"region":          region,
		"location":        region,
		"zone":            "-",
		"billing_account": billingId,
	}

	client := transport_tpg.NewDCLFirebaserulesClient(config, config.UserAgent, "", 0)
	err = client.DeleteAllRelease(context.Background(), d["project"], isDeletableFirebaserulesRelease)
	if err != nil {
		return err
	}
	return nil
}

func isDeletableFirebaserulesRelease(r *firebaserules.Release) bool {
	return sweeper.IsSweepableTestResource(*r.Name)
}
