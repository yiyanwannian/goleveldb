package table

import (
	"testing"

	"github.com/yiyanwannian/goleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
