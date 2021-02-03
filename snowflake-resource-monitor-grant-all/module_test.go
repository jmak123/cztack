package test

/*
 NOTE: automatically generated by scripts/snowflake_generate_grant_all.
			 Changes made directly to this file will be overwritten.
			 Make improvements there so everyone can benefit.
*/

import (
	"encoding/json"
	"testing"

	"github.com/chanzuckerberg/go-misc/tftest"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

const vars string = "{\"monitor_name\":\"monitor\",\"roles\":[\"role_a\",\"role_b\",\"role_c\"]}"
const onFutureVars string = ""

func TestModule(t *testing.T) {
	test := tftest.Test{
		Mode: tftest.Plan,

		Setup: func(t *testing.T) *terraform.Options {
			opts := &terraform.Options{
				TerraformDir: ".",
				EnvVars:      map[string]string{},
				Vars:         mustJson(vars),
			}
			return opts
		},
		Validate: func(t *testing.T, options *terraform.Options) {},
	}

	test.Run(t)
}

func TestModule_onFuture(t *testing.T) {
	if onFutureVars == "" {
		return // nothing to test, no future vars
	}
	test := tftest.Test{
		Mode: tftest.Plan,
		Setup: func(t *testing.T) *terraform.Options {
			opts := &terraform.Options{
				TerraformDir: ".",
				EnvVars:      map[string]string{},
				Vars:         mustJson(vars),
			}
			return opts
		},
		Validate: func(t *testing.T, options *terraform.Options) {},
	}

	test.Run(t)
}

func mustJson(jsonData string) map[string]interface{} {
	unmarshalled := map[string]interface{}{}

	err := json.Unmarshal([]byte(jsonData), &unmarshalled)
	if err != nil {
		panic(err) // should never happen since generated code
	}
	return unmarshalled
}
