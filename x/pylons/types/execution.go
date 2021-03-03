package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TypeExecution is a store key for execution
const TypeExecution = "execution"

// Execution is a recipe execution used for tracking the execution - specifically a
// scheduled execution
type Execution struct {
	NodeVersion SemVer
	ID          string
	RecipeID    string // the recipe guid
	CookbookID  string
	CoinInputs  sdk.Coins
	ItemInputs  []Item
	BlockHeight int64
	Sender      string
	Completed   bool
}

// ExecutionList describes executions list
type ExecutionList struct {
	Executions []Execution
}

// NewExecution return a new Execution
func NewExecution(recipeID string, cookbookID string, ci sdk.Coins,
	itemInputs []Item,
	blockHeight int64, sender sdk.AccAddress,
	completed bool) Execution {

	exec := Execution{
		NodeVersion: SemVer{"0.0.1"},
		RecipeID:    recipeID,
		CookbookID:  cookbookID,
		CoinInputs:  ci,
		ItemInputs:  itemInputs,
		BlockHeight: blockHeight,
		Sender:      sender.String(),
		Completed:   completed,
	}

	exec.ID = KeyGen(sender)
	return exec
}

func (e Execution) String() string {
	return fmt.Sprintf(`
		Execution{ 
			NodeVersion: %s,
			ID: %s,
			RecipeID: %s,
			CookbookID: %s,
			CoinInputs: %+v,
			ItemInputs: %+v,
			BlockHeight: %d,
			Sender: %s,
			Completed: %t,
		}`, e.NodeVersion, e.ID, e.RecipeID, e.CookbookID, e.CoinInputs, e.ItemInputs,
		e.BlockHeight, e.Sender, e.Completed)
}

func ExecutionsToListProto(execs []Execution) []*GetExecutionResponse {
	var res []*GetExecutionResponse
	for _, exec := range execs {
		res = append(res, &GetExecutionResponse{
			NodeVersion: &exec.NodeVersion,
			ID:          exec.ID,
			RecipeID:    exec.RecipeID,
			CookbookID:  exec.CookbookID,
			CoinsInput:  exec.CoinInputs,
			ItemInputs:  ItemInputsToProto(exec.ItemInputs),
			BlockHeight: exec.BlockHeight,
			Sender:      exec.Sender,
			Completed:   exec.Completed,
		})
	}
	return res
}
