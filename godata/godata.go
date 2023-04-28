package godata

type DataStruct struct {
		BoxTable		*MapBox
		ItemTable		*MapItem
		ItemAnimationTable		*MapItemAnimation
		LandUpgradeDataTable		*MapLandUpgradeData
		LanguageTable		*MapLanguage
		MakeUpTable		*MapMakeUp
		ShopTable		*MapShop
}

var  CSVData DataStruct

func init(){
	CSVData.BoxTable = CreateBoxTable()
	CSVData.ItemTable = CreateItemTable()
	CSVData.ItemAnimationTable = CreateItemAnimationTable()
	CSVData.LandUpgradeDataTable = CreateLandUpgradeDataTable()
	CSVData.LanguageTable = CreateLanguageTable()
	CSVData.MakeUpTable = CreateMakeUpTable()
	CSVData.ShopTable = CreateShopTable()
}
