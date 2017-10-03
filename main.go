package main

import (
	"fmt"
	"time"
)

func main() {
	formatWorkLog()
}

func formatWorkLog() {
	t := time.Now()

	format := `%s
刷卡
# 今天做了什麼
# 今天遇到的難題
# 明天的計畫
`
	fmt.Printf(format, t.Format("2006-01-02 Mon"))
}
