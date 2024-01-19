package iso_level

import (
	"github.com/glebziz/fs_db"
	"github.com/glebziz/fs_db/internal/model"
	store "github.com/glebziz/fs_db/internal/proto"
)

func Convert(level store.TxIsoLevel) model.TxIsoLevel {
	switch level {
	case store.TxIsoLevel_ISO_LEVEL_READ_UNCOMMITTED:
		return fs_db.IsoLevelReadUncommitted
	case store.TxIsoLevel_ISO_LEVEL_READ_COMMITTED:
		return fs_db.IsoLevelReadCommitted
	case store.TxIsoLevel_ISO_LEVEL_REPEATABLE_READ:
		return fs_db.IsoLevelRepeatableRead
	case store.TxIsoLevel_ISO_LEVEL_SERIALIZABLE:
		return fs_db.IsoLevelSerializable
	default:
		return fs_db.IsoLevelDefault
	}
}

func ConvertToGrpc(level model.TxIsoLevel) store.TxIsoLevel {
	switch level {
	case fs_db.IsoLevelReadUncommitted:
		return store.TxIsoLevel_ISO_LEVEL_READ_UNCOMMITTED
	case fs_db.IsoLevelReadCommitted:
		return store.TxIsoLevel_ISO_LEVEL_READ_COMMITTED
	case fs_db.IsoLevelRepeatableRead:
		return store.TxIsoLevel_ISO_LEVEL_REPEATABLE_READ
	case fs_db.IsoLevelSerializable:
		return store.TxIsoLevel_ISO_LEVEL_SERIALIZABLE
	default:
		return store.TxIsoLevel_ISO_LEVEL_READ_COMMITTED
	}
}
