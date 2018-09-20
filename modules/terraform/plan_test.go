package terraform

import (
	"terratest/modules/aws"
	"terratest/modules/files"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlanWithNoChanges(t *testing.T) {
	t.Parallel()
	testFolder, err := files.CopyTerraformFolderToTemp("../../test/fixtures/terraform-no-error", t.Name())
	if err != nil {
		t.Fatal(err)
	}
	awsRegion := aws.GetRandomRegion(t, nil, nil)
	options := &Options{
		TerraformDir: testFolder,

		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	}
	exitCode := InitAndPlan(t, options)
	assert.Equal(t, DefaultSuccessExitCode, exitCode)
}

func TestPlanWithChanges(t *testing.T) {
	t.Parallel()
	testFolder, err := files.CopyTerraformFolderToTemp("../../examples/terraform-aws-example", t.Name())
	if err != nil {
		t.Fatal(err)
	}
	awsRegion := aws.GetRandomRegion(t, nil, nil)
	options := &Options{
		TerraformDir: testFolder,

		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
	}
	exitCode := InitAndPlan(t, options)
	assert.Equal(t, TerraformPlanChangesPresentExitCode, exitCode)
}

func TestPlanWithFailure(t *testing.T) {
	t.Parallel()

	testFolder, err := files.CopyTerraformFolderToTemp("../../test/fixtures/terraform-with-plan-error", t.Name())
	if err != nil {
		t.Fatal(err)
	}

	options := &Options{
		TerraformDir: testFolder,
	}

	_, getExitCodeErr := InitAndPlanE(t, options)
	assert.Error(t, getExitCodeErr)
}