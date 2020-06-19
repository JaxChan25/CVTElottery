package model

//执行数据迁移

func migration() {
	// 自动迁移模式

	DB.AutoMigrate(&GameUser{})
	DB.AutoMigrate(&Address{})
	DB.AutoMigrate(&UserAction{})
	DB.AutoMigrate(&Activity{})
	DB.AutoMigrate(&GamePrize{})
	DB.AutoMigrate(&GameManager{})

}
