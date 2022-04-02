package keeper_test

import (
	"testing"

	oracletypes "github.com/Sifchain/sifnode/x/oracle/types"
	"github.com/Sifchain/sifnode/x/tokenregistry/test"
	"github.com/Sifchain/sifnode/x/tokenregistry/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestGetRegistryPaginatedEmpty(t *testing.T) {
	app, ctx, _ := test.CreateTestApp(false)

	registry, err := app.TokenRegistryKeeper.GetRegistryPaginated(ctx, 0, 100)
	assert.NoError(t, err)
	assert.Equal(t, types.Registry{}, registry)

	registry, err = app.TokenRegistryKeeper.GetRegistryPaginated(ctx, 1, 100)
	assert.NoError(t, err)
	assert.Equal(t, types.Registry{}, registry)

	registry, err = app.TokenRegistryKeeper.GetRegistryPaginated(ctx, 100, 100)
	assert.NoError(t, err)
	assert.Equal(t, types.Registry{}, registry)

	registry, err = app.TokenRegistryKeeper.GetRegistryPaginated(ctx, 1, 0)
	assert.NoError(t, err)
	assert.Equal(t, types.Registry{}, registry)

	registry, err = app.TokenRegistryKeeper.GetRegistryPaginated(ctx, 1, 1)
	assert.NoError(t, err)
	assert.Equal(t, types.Registry{}, registry)

	registry, err = app.TokenRegistryKeeper.GetRegistryPaginated(ctx, 1, 99)
	assert.NoError(t, err)
	assert.Equal(t, types.Registry{}, registry)

}

func TestGetRegistryPaginatedInvalidSize(t *testing.T) {
	app, ctx, _ := test.CreateTestApp(false)
	registry, err := app.TokenRegistryKeeper.GetRegistryPaginated(ctx, 0, 101)
	assert.ErrorIs(t, err, errors.Wrap(errors.ErrTxTooLarge, "Registry Requests limited to 100 or less"))
	assert.Equal(t, types.Registry{}, registry)
}

func TestGetRegistryPaginated(t *testing.T) {
	app, ctx, _ := test.CreateTestApp(false)
	entry1 := types.RegistryEntry{
		Denom:       "rowan",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_CLP},
	}
	entry2 := types.RegistryEntry{
		Denom:       "t1",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_UNSPECIFIED},
	}
	entry3 := types.RegistryEntry{
		Denom:       "t2",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCEXPORT},
	}
	entry4 := types.RegistryEntry{
		Denom:       "t3",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_CLP},
	}
	entry5 := types.RegistryEntry{
		Denom:       "t4",
		Decimals:    6,
		Permissions: []types.Permission{types.Permission_UNSPECIFIED},
	}
	entry6 := types.RegistryEntry{
		Denom:       "t5",
		Decimals:    23,
		Permissions: []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCEXPORT},
	}
	entry7 := types.RegistryEntry{
		Denom:       "t6",
		Decimals:    255,
		Permissions: []types.Permission{types.Permission_CLP},
	}
	entry8 := types.RegistryEntry{
		Denom:       "t7",
		Decimals:    0,
		Permissions: []types.Permission{types.Permission_UNSPECIFIED},
	}
	entry9 := types.RegistryEntry{
		Denom:       "t8",
		Decimals:    1,
		Permissions: []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCEXPORT},
	}
	entry10 := types.RegistryEntry{
		Denom:       "t9",
		Decimals:    128,
		Permissions: []types.Permission{types.Permission_CLP},
	}
	entry11 := types.RegistryEntry{
		Denom:       "t10",
		Decimals:    5,
		Permissions: []types.Permission{types.Permission_UNSPECIFIED},
	}
	entry12 := types.RegistryEntry{
		Denom:       "t11",
		Decimals:    3,
		Permissions: []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCEXPORT},
	}

	entries := []*types.RegistryEntry{
		&entry1,
		&entry2,
		&entry3,
		&entry4,
		&entry5,
		&entry6,
		&entry7,
		&entry8,
		&entry9,
		&entry10,
		&entry11,
		&entry12,
	}

	for _, entry := range entries {
		app.TokenRegistryKeeper.SetToken(ctx, entry)
	}

	// Test fetching the entire registry since it can fit in one call
	resultRegistry, err := app.TokenRegistryKeeper.GetRegistryPaginated(ctx, 1, 12)
	assert.NoError(t, err)
	assert.ElementsMatch(t, entries, resultRegistry.Entries)

	// Test fetching the registry over many calls
	// Page size of 3 total entries in 4 pages
	var appendEntries []*types.RegistryEntry
	var limit uint = 3
	for i := uint(1); i <= uint(len(entries))/limit; i++ {
		resultRegistry, err = app.TokenRegistryKeeper.GetRegistryPaginated(ctx, i, limit)
		assert.NoError(t, err)
		for _, registryEntry := range resultRegistry.Entries {
			assert.Contains(t, entries, registryEntry)
		}
		appendEntries = append(appendEntries, resultRegistry.Entries...)
	}
	assert.ElementsMatch(t, entries, appendEntries)
}

func TestGetRegistryEntry(t *testing.T) {
	app, ctx, _ := test.CreateTestApp(false)
	entry1 := types.RegistryEntry{
		Denom:       "rowan",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_CLP},
	}
	entry2 := types.RegistryEntry{
		Denom:       "t1",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_UNSPECIFIED},
	}
	entry3 := types.RegistryEntry{
		Denom:       "t2",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCEXPORT},
	}

	app.TokenRegistryKeeper.SetToken(ctx, &entry1)
	app.TokenRegistryKeeper.SetToken(ctx, &entry2)
	// Duplicate permission is interpreted correctly.
	app.TokenRegistryKeeper.SetToken(ctx, &entry3)

	// Test Entry 2, Entry 1, Entry 3 in that order followed by an invalid entry
	// Entry 2
	actualEntry, err := app.TokenRegistryKeeper.GetRegistryEntry(ctx, "t1")
	assert.NoError(t, err)
	assert.Equal(t, entry2, *actualEntry)
	// Entry 1
	actualEntry, err = app.TokenRegistryKeeper.GetRegistryEntry(ctx, "rowan")
	assert.NoError(t, err)
	assert.Equal(t, entry1, *actualEntry)
	// Entry 3
	actualEntry, err = app.TokenRegistryKeeper.GetRegistryEntry(ctx, "t2")
	assert.NoError(t, err)
	assert.Equal(t, entry3, *actualEntry)
	// Invalid Entry
	actualEntry, err = app.TokenRegistryKeeper.GetRegistryEntry(ctx, "InvalidToken")
	assert.ErrorIs(t, err, errors.Wrap(errors.ErrKeyNotFound, "registry entry not found"))
	assert.Nil(t, actualEntry)
}

func TestKeeper_CheckDenomPermissions(t *testing.T) {
	app, ctx, _ := test.CreateTestApp(false)
	app.TokenRegistryKeeper.SetToken(ctx, &types.RegistryEntry{
		Denom:       "rowan",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_CLP},
	})
	app.TokenRegistryKeeper.SetToken(ctx, &types.RegistryEntry{
		Denom:       "t1",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_UNSPECIFIED},
	})
	// Duplicate permission is interpreted correctly.
	app.TokenRegistryKeeper.SetToken(ctx, &types.RegistryEntry{
		Denom:       "t2",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCEXPORT},
	})
	entry, err := app.TokenRegistryKeeper.GetRegistryEntry(ctx, "rowan")
	assert.NoError(t, err)
	entry2, err := app.TokenRegistryKeeper.GetRegistryEntry(ctx, "t2")
	assert.NoError(t, err)
	assert.True(t, app.TokenRegistryKeeper.CheckEntryPermissions(entry, []types.Permission{types.Permission_CLP}))
	assert.False(t, app.TokenRegistryKeeper.CheckEntryPermissions(entry, []types.Permission{types.Permission_IBCIMPORT}))
	assert.False(t, app.TokenRegistryKeeper.CheckEntryPermissions(entry, []types.Permission{types.Permission_CLP, types.Permission_IBCIMPORT}))
	assert.False(t, app.TokenRegistryKeeper.CheckEntryPermissions(entry2, []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCIMPORT}))
	assert.True(t, app.TokenRegistryKeeper.CheckEntryPermissions(entry, []types.Permission{}))
}

func TestKeeper_AddRemoveRegisterAll(t *testing.T) {
	app, ctx, _ := test.CreateTestApp(false)
	app.TokenRegistryKeeper.SetToken(ctx, &types.RegistryEntry{
		Denom:       "rowan",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_CLP},
	})
	app.TokenRegistryKeeper.SetToken(ctx, &types.RegistryEntry{
		Denom:       "t1",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_UNSPECIFIED},
	})

	app.TokenRegistryKeeper.SetToken(ctx, &types.RegistryEntry{
		Denom:       "t2",
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCEXPORT},
	})
	registry := app.TokenRegistryKeeper.GetRegistry(ctx)
	assert.Equal(t, len(registry.Entries), 3)

	entries := []*types.RegistryEntry{}
	entries = append(entries, &types.RegistryEntry{
		Denom:       "t2",
		Decimals:    19,
		Permissions: []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCEXPORT},
	})

	entries = append(entries, &types.RegistryEntry{
		Denom:       "t3",
		Decimals:    19,
		Permissions: []types.Permission{types.Permission_IBCEXPORT, types.Permission_IBCEXPORT},
	})

	// add entries
	app.TokenRegistryKeeper.AddMultipleTokens(ctx, entries)

	registry = app.TokenRegistryKeeper.GetRegistry(ctx)

	assert.Equal(t, len(registry.Entries), 4)

	// remove entries
	denom := []string{"t2", "t3", "t4"}
	app.TokenRegistryKeeper.RemoveMultipleTokens(ctx, denom)

	registry = app.TokenRegistryKeeper.GetRegistry(ctx)

	assert.Equal(t, len(registry.Entries), 2)
}

func TestKeeper_SetFirstLockDoublePeg(t *testing.T) {
	app, ctx, _ := test.CreateTestApp(false)

	denom := "denom"
	app.TokenRegistryKeeper.SetToken(ctx, &types.RegistryEntry{
		Denom:       denom,
		Decimals:    18,
		Permissions: []types.Permission{types.Permission_CLP},
	})
	networkDescriptor := oracletypes.NetworkDescriptor_NETWORK_DESCRIPTOR_GANACHE

	assert.True(t, app.TokenRegistryKeeper.GetFirstLockDoublePeg(ctx, denom, networkDescriptor))
	app.TokenRegistryKeeper.SetFirstDoublePeg(ctx, denom, networkDescriptor)
	assert.False(t, app.TokenRegistryKeeper.GetFirstLockDoublePeg(ctx, denom, networkDescriptor))

}

func TestKeeper_SetAdminAccount(t *testing.T) {
	app, ctx, admin := test.CreateTestApp(false)
	address, _ := sdk.AccAddressFromBech32(admin)
	newAddress, _ := sdk.AccAddressFromBech32("sif1azpar20ck9lpys89r8x7zc8yu0qzgvtp48ng5v")

	assert.True(t, app.TokenRegistryKeeper.IsAdminAccount(ctx, address))
	assert.False(t, app.TokenRegistryKeeper.IsAdminAccount(ctx, newAddress))
	app.TokenRegistryKeeper.SetAdminAccount(ctx, newAddress)
	assert.True(t, app.TokenRegistryKeeper.IsAdminAccount(ctx, newAddress))
	assert.False(t, app.TokenRegistryKeeper.IsAdminAccount(ctx, address))
}
