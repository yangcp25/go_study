package state

import "fmt"

type Machine struct {
	state IState
}

type IState interface {
	approval(m *Machine)
	reject(m *Machine)
	getName() string
}

func (m *Machine) setState(state IState) {
	m.state = state
}

func (m *Machine) getStateName() string {
	return m.state.getName()
}

func (m *Machine) Approval() {
	m.state.approval(m)
}

func (m *Machine) reject() {
	m.state.reject(m)
}

// 领导审核
type leaderApproval struct{}

func (l leaderApproval) approval(m *Machine) {
	fmt.Println("领导审核")
	m.setState(GetFinanceApproveState())
}

type financeApproveState struct {
}

func (f financeApproveState) approval(m *Machine) {
	fmt.Println("审核通过")
}

func (f financeApproveState) reject(m *Machine) {
	m.setState(getLeaderApproval())
}

func getLeaderApproval() IState {
	return &leaderApproval{}
}

func (f financeApproveState) getName() string {
	return "FinanceApproveState"
}

func GetFinanceApproveState() IState {
	return &financeApproveState{}
}

func (l leaderApproval) reject(m *Machine) {

}

func (l leaderApproval) getName() string {
	return "LeaderApproval"
}
