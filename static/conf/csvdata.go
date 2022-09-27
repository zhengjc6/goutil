package csvdata

type DataStruct struct {
		UserTable		mapuser
}

var  CSVData DataStruct

func init(){
	CSVData.UserTable = CreateUserTable()
}
