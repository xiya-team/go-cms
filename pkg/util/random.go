package util

import (
	"math/rand"
	"time"
)

//算法1
func Random01() (newMember bool) {
	rand.Seed(time.Now().UnixNano())
	random01 := rand.Intn(100)
	if random01 >= 50 {
		newMember = true
	} else {
		newMember = false
	}
	return
}

//算法2
//判断上一个数，如果上一个数为真，返回真，否则返回假
func Random02(oldMember bool) (newMember bool) {
	if oldMember {
		newMember = true
	} else {
		newMember = false
	}
	return
}
