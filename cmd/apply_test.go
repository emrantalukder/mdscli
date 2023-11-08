package cmd

import "testing"

func TestProcess(t *testing.T) {
	applyRoleBindings("./../data/role-bindings.yml")
}
