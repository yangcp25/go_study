package state

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMachine_Approval(t *testing.T) {
	machine := &Machine{
		state: getLeaderApproval(),
	}

	assert.Equal(t, "LeaderApproval", machine.getStateName())
	machine.Approval()
	assert.Equal(t, "FinanceApproveState", machine.getStateName())
	machine.reject()
	assert.Equal(t, "LeaderApproval", machine.getStateName())
	machine.Approval()
	assert.Equal(t, "FinanceApproveState", machine.getStateName())
	machine.Approval()
}
