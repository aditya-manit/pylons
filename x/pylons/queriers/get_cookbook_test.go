package queriers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Pylons-tech/pylons/x/pylons/handlers"
	"github.com/Pylons-tech/pylons/x/pylons/keep"
	"github.com/Pylons-tech/pylons/x/pylons/types"
)

func TestGetCookbook(t *testing.T) {
	tci := keep.SetupTestCoinInput()
	tci.PlnH = handlers.NewMsgServerImpl(tci.PlnK)
	tci.PlnQ = NewQuerierServerImpl(tci.PlnK)

	sender1, _, _, _ := keep.SetupTestAccounts(t, tci, types.NewPylon(1000000), nil, nil, nil)

	// mock cookbook
	cbData := handlers.MockCookbook(tci, sender1)

	cases := map[string]struct {
		cookbookID    string
		desiredError  string
		showError     bool
		desiredRcpCnt int
		cbName        string
	}{
		"error check when providing invalid cookbook ID": {
			cookbookID:    "invalid cookbookID",
			showError:     true,
			desiredError:  "The cookbook doesn't exist",
			desiredRcpCnt: 0,
		},
		"error check when not providing cookbookID": {
			cookbookID:    "",
			showError:     true,
			desiredError:  "no cookbook id is provided in path",
			desiredRcpCnt: 0,
		},
		"get cookbook successful check": {
			cookbookID:    cbData.CookbookID,
			showError:     false,
			desiredError:  "",
			desiredRcpCnt: 1,
			cbName:        "cookbook-00001",
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			result, err := tci.PlnQ.GetCookbook(
				sdk.WrapSDKContext(tci.Ctx),
				&types.GetCookbookRequest{CookbookID: tc.cookbookID},
			)
			if tc.showError {
				require.True(t, strings.Contains(err.Error(), tc.desiredError))
			} else {
				require.True(t, err == nil)
				require.True(t, result.Name == tc.cbName)
			}
		})
	}
}
