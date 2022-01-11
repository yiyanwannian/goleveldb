package iterator_test

import (
	"testing"

	"github.com/yiyanwannian/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
