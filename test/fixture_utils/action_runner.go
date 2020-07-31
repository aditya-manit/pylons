package fixturetest

import (
	testing "github.com/Pylons-tech/pylons/test/evtesting"
)

// ActFunc describes the type of function used for action running test
type ActFunc func(FixtureStep, *testing.T)

var actFuncs = make(map[string]ActFunc)

// RegisterActionRunner registers action runner function
func RegisterActionRunner(action string, fn ActFunc) {
	actFuncs[action] = fn
}

// GetActionRunner get registered action runner function
func GetActionRunner(action string) ActFunc {
	return actFuncs[action]
}

// RunActionRunner execute registered action runner function
func RunActionRunner(action string, step FixtureStep, t *testing.T) {
	fn := GetActionRunner(action)
	t.WithFields(testing.Fields{
		"action": step.Action,
	}).MustTrue(fn != nil, "step with unrecognizable action found")
	fn(step, t)
}

// RegisterDefaultActionRunners register default test functions
func RegisterDefaultActionRunners() {
	RegisterActionRunner("create_account", RunCreateAccount)
	RegisterActionRunner("get_pylons", RunGetPylons)
	RegisterActionRunner("mock_account", RunMockAccount) // create account + get pylons
	RegisterActionRunner("send_coins", RunSendCoins)
	RegisterActionRunner("fiat_item", RunFiatItem)
	RegisterActionRunner("update_item_string", RunUpdateItemString)
	RegisterActionRunner("send_items", RunSendItems)
	RegisterActionRunner("create_cookbook", RunCreateCookbook)
	RegisterActionRunner("mock_cookbook", RunMockCookbook) // mock_account + create_cookbook
	RegisterActionRunner("create_recipe", RunCreateRecipe)
	RegisterActionRunner("execute_recipe", RunExecuteRecipe)
	RegisterActionRunner("check_execution", RunCheckExecution)
	RegisterActionRunner("create_trade", RunCreateTrade)
	RegisterActionRunner("fulfill_trade", RunFulfillTrade)
	RegisterActionRunner("disable_trade", RunDisableTrade)
	RegisterActionRunner("multi_msg_tx", RunMultiMsgTx)
}
