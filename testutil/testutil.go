package testutil

import (
	"math/rand"
	"strconv"
	"time"
)

// ランダム文字列を作成する
func RandomString() string {
	return strconv.FormatInt(time.Now().Unix()^rand.Int63(), 16)
}
