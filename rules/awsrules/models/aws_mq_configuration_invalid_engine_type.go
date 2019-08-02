// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsMqConfigurationInvalidEngineTypeRule checks the pattern is valid
type AwsMqConfigurationInvalidEngineTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsMqConfigurationInvalidEngineTypeRule returns new rule with default attributes
func NewAwsMqConfigurationInvalidEngineTypeRule() *AwsMqConfigurationInvalidEngineTypeRule {
	return &AwsMqConfigurationInvalidEngineTypeRule{
		resourceType:  "aws_mq_configuration",
		attributeName: "engine_type",
		enum: []string{
			"ACTIVEMQ",
		},
	}
}

// Name returns the rule name
func (r *AwsMqConfigurationInvalidEngineTypeRule) Name() string {
	return "aws_mq_configuration_invalid_engine_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMqConfigurationInvalidEngineTypeRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsMqConfigurationInvalidEngineTypeRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsMqConfigurationInvalidEngineTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMqConfigurationInvalidEngineTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`engine_type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}