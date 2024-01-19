package main

import (
	"testing"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func TestCheckValidity(t *testing.T) {
	stack := NewMyStack(cdktf.Testing_App(nil), "stack")
	assertion := cdktf.Testing_ToBeValidTerraform(cdktf.Testing_FullSynth(stack))

	if !*assertion {
		t.Error("Assertion Failed")
	}
}
