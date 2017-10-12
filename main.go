package main

import (
	"fmt"
	"time"

	"github.com/nyogjtrc/note-format/template"
)

func main() {
	formatWorkLog()
	template.NextWeekBulletJournal()
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
