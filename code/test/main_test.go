package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"

//	"github.com/stretchr/testify/assert"
)

func TestTerraformHelloWorldExample(t *testing.T) {
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../example",
	})

	//defer terraform.Destroy(t, terraformOptions)

	terraform.Init(t, terraformOptions)
  defer terraform.Destroy(t, terraformOptions)
  // When the test is completed, teardown the infrastructure by calling terraform destroy
 

	//output := terraform.Output(t, terraformOptions, "hello_world")
	//assert.Equal(t, "Hello, World!", output)
}
