// 时间戳转时间
package main

import (
	"fmt"
	"time"
)

func main() {
	var ts int64 = 1666761965
	fmt.Println(time.Unix(ts, 0))
}
