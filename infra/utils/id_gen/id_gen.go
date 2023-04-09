package id_gen

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func DoGen() (res int64) {
	str := fmt.Sprintf("%v%v", time.Now().Unix(), rand.Uint64()%10000)
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return int64(rand.Uint64())
	}
	return res
}
