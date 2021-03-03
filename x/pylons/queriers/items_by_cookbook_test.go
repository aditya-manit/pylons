package queriers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
	"testing"

	"github.com/Pylons-tech/pylons/x/pylons/handlers"
	"github.com/Pylons-tech/pylons/x/pylons/keep"
	"github.com/Pylons-tech/pylons/x/pylons/types"
	"github.com/stretchr/testify/require"
)

func TestQueriersItemsByCookbook(t *testing.T) {
	tci := keep.SetupTestCoinInput()
	tci.PlnH = handlers.NewMsgServerImpl(tci.PlnK)
	tci.PlnQ = NewQuerierServerImpl(tci.PlnK)

	sender1, _, _, _ := keep.SetupTestAccounts(t, tci, types.NewPylon(1000000), nil, nil, nil)

	// mock cookbook
	cbData := handlers.MockCookbookByName(tci, sender1, "cookbook-00001")
	cbData1 := handlers.MockCookbookByName(tci, sender1, "cookbook-00002")

	item := keep.GenItem(cbData.CookbookID, sender1, "Raichu")
	err := tci.PlnK.SetItem(tci.Ctx, *item)
	require.True(t, err == nil)

	cases := map[string]struct {
		cookbookID    string
		desiredError  string
		desiredLength int
		showError     bool
	}{
		"not existing cookbook id": {
			cookbookID:    "invalidCookbookID",
			desiredError:  "",
			desiredLength: 0,
			showError:     false,
		},
		"error check when not providing cookbookID": {
			cookbookID:    "",
			desiredError:  "no cookbook id is provided in path",
			desiredLength: 0,
			showError:     true,
		},
		"cookbook with no item": {
			cookbookID:    cbData1.CookbookID,
			desiredError:  "",
			desiredLength: 0,
			showError:     false,
		},
		"cookbook with 1 item": {
			cookbookID:    cbData.CookbookID,
			desiredError:  "",
			desiredLength: 1,
			showError:     false,
		},
	}
	for testName, tc := range cases {
		t.Run(testName, func(t *testing.T) {
			result, err := tci.PlnQ.ItemsByCookbook(
				sdk.WrapSDKContext(tci.Ctx),
				&types.ItemsByCookbookRequest{
					CookbookID: tc.cookbookID,
				},
			)

			if tc.showError {
				require.True(t, strings.Contains(err.Error(), tc.desiredError))
			} else {
				require.True(t, err == nil)

				cItems := result.Items
				require.True(t, len(cItems) == tc.desiredLength)
			}
		})
	}
}
