package storerunner

import (
	"github.com/c4pt0r/storeadapter"
)

type StoreRunner interface {
	Start()
	Stop()
	GoAway()
	ComeBack()
	NodeURLS() []string
	DiskUsage() (bytes int64, err error)
	FastForwardTime(seconds int)
	Reset()
	Adapter() storeadapter.StoreAdapter
}
