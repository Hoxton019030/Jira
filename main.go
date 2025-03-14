package main

import (
	"jira/config"
	"jira/database"
	"jira/router"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//建立與資料庫ㄉ連線
	database.Init()
	//用viper取得config.yaml的資訊
	config.Init()
	//綁定url path 跟 gin
	router.Init()
	//啟動Gin的服務
	router.Run()
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
