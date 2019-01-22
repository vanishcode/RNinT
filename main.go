package main

import (
	"fmt"
	"time"
	"github.com/fatih/color"
	"github.com/vanishcode/RNinT/utils"
	"github.com/biezhi/moe"
	// "github.com/modood/table"
)

func main() {
	color.Green(utils.Logo)
	defer func() {
		color.Yellow("💔  用户退出")
	}()
	var name string
	color.Cyan("#请输入小说名：")
	fmt.Scanln(&name)
	moe := moe.New("正则爬取...").Spinner("dots2")
	moe.Start()
	time.Sleep(5 * time.Second)
	moe.Stop()
}