package misc

import "sync/atomic"

var loadingAssets int32 = 0
func StartLoading() {
	atomic.AddInt32(&loadingAssets, 1)
}

func LoadingDone() {
	atomic.AddInt32(&loadingAssets, -1)
}

func IsLoadingDone() bool {
	return atomic.LoadInt32(&loadingAssets) == 0
}
