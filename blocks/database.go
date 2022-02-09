package blocks

import (
	"time"

	"github.com/archway-network/valuter/database"
	"github.com/archway-network/valuter/types"
)

func DBRowToBlockRecord(row database.RowType) types.BlockRecord {

	if row == nil {
		return types.BlockRecord{}
	}

	row[database.FIELD_BLOCKS_BLOCK_HASH] = ""

	return types.BlockRecord{

		BlockHash: row[database.FIELD_BLOCKS_BLOCK_HASH].(string),
		Height:    uint64(row[database.FIELD_BLOCKS_HEIGHT].(int64)),
		NumOfTxs:  uint64(row[database.FIELD_BLOCKS_NUM_OF_TXS].(int64)),
		Time:      row[database.FIELD_BLOCKS_TIME].(time.Time),
		// Signers:   []BlockSignersRecord, //TODO: Check if we really need to have signers here as well
	}
}

func DBRowsToBlockRecords(row []database.RowType) []types.BlockRecord {

	var res []types.BlockRecord
	for i := range row {
		res = append(res, DBRowToBlockRecord(row[i]))
	}

	return res

}