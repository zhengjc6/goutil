package godata
type MakeUp struct {
	Id		int64		//道具id
	Type		int64		//分解/合成
	Makeupnum		int64		//合成/分解需要的数量
	Makeupid		int64		//合成/分解需要的id
					//
}
type MapMakeUp map[int64]MakeUp
func CreateMakeUpTable() *MapMakeUp {
	data := MapMakeUp{
		1013020001:MakeUp{
			Id:1013020001,
			Type:2,
			Makeupnum:100,
			Makeupid:1019020001,
		},
		1013020002:MakeUp{
			Id:1013020002,
			Type:2,
			Makeupnum:100,
			Makeupid:1019020002,
		},
		1013020003:MakeUp{
			Id:1013020003,
			Type:2,
			Makeupnum:100,
			Makeupid:1019020003,
		},
		1013010001:MakeUp{
			Id:1013010001,
			Type:2,
			Makeupnum:100,
			Makeupid:1019020010,
		},
		1013010002:MakeUp{
			Id:1013010002,
			Type:2,
			Makeupnum:100,
			Makeupid:1019020011,
		},
		1013010003:MakeUp{
			Id:1013010003,
			Type:2,
			Makeupnum:100,
			Makeupid:1019020012,
		},
	}
	return &data
}
