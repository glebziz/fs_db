package iso_level

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
)

func TestConvert(t *testing.T) {
	for _, tc := range []struct {
		name     string
		lvl      store.TxIsoLevel
		localLvl model.TxIsoLevel
	}{
		{
			name:     "read uncommitted",
			lvl:      store.TxIsoLevel_ISO_LEVEL_READ_UNCOMMITTED,
			localLvl: fs_db.IsoLevelReadUncommitted,
		},
		{
			name:     "read committed",
			lvl:      store.TxIsoLevel_ISO_LEVEL_READ_COMMITTED,
			localLvl: fs_db.IsoLevelReadCommitted,
		},
		{
			name:     "repeatable read",
			lvl:      store.TxIsoLevel_ISO_LEVEL_REPEATABLE_READ,
			localLvl: fs_db.IsoLevelRepeatableRead,
		},
		{
			name:     "serializable",
			lvl:      store.TxIsoLevel_ISO_LEVEL_SERIALIZABLE,
			localLvl: fs_db.IsoLevelSerializable,
		},
		{
			name:     "unknown level",
			lvl:      100,
			localLvl: fs_db.IsoLevelDefault,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			lvl := Convert(tc.lvl)

			require.Equal(t, tc.localLvl, lvl)
		})
	}
}

func TestConvertToGrpc(t *testing.T) {
	for _, tc := range []struct {
		name     string
		lvl      store.TxIsoLevel
		localLvl model.TxIsoLevel
	}{
		{
			name:     "read uncommitted",
			lvl:      store.TxIsoLevel_ISO_LEVEL_READ_UNCOMMITTED,
			localLvl: fs_db.IsoLevelReadUncommitted,
		},
		{
			name:     "read committed",
			lvl:      store.TxIsoLevel_ISO_LEVEL_READ_COMMITTED,
			localLvl: fs_db.IsoLevelReadCommitted,
		},
		{
			name:     "repeatable read",
			lvl:      store.TxIsoLevel_ISO_LEVEL_REPEATABLE_READ,
			localLvl: fs_db.IsoLevelRepeatableRead,
		},
		{
			name:     "serializable",
			lvl:      store.TxIsoLevel_ISO_LEVEL_SERIALIZABLE,
			localLvl: fs_db.IsoLevelSerializable,
		},
		{
			name:     "unknown level",
			lvl:      store.TxIsoLevel_ISO_LEVEL_READ_COMMITTED,
			localLvl: 100,
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			lvl := ConvertToGrpc(tc.localLvl)

			require.Equal(t, tc.lvl, lvl)
		})
	}
}
