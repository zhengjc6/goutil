package godata
type Box struct {
	Itemid		int64		//物品id
	Min		int64		//最小数量
	Max		int64		//最大数量
	Boxnum		int64		//集合包id
	Random		int64		//权重
					//
}
type MapBox map[int64]Box
func CreateBoxTable() *MapBox {
	data := MapBox{
		1014000002:Box{
			Itemid:1014000002,
			Min:1,
			Max:1,
			Boxnum:4000000001,
			Random:2000,
		},
		1014000004:Box{
			Itemid:1014000004,
			Min:1,
			Max:1,
			Boxnum:4000000001,
			Random:3500,
		},
		1014000005:Box{
			Itemid:1014000005,
			Min:1,
			Max:1,
			Boxnum:4000000001,
			Random:4500,
		},
		1013010001:Box{
			Itemid:1013010001,
			Min:1,
			Max:1,
			Boxnum:4000000002,
			Random:50,
		},
		1013010002:Box{
			Itemid:1013010002,
			Min:1,
			Max:1,
			Boxnum:4000000002,
			Random:50,
		},
		1013010003:Box{
			Itemid:1013010003,
			Min:1,
			Max:1,
			Boxnum:4000000002,
			Random:50,
		},
		1019180001:Box{
			Itemid:1019180001,
			Min:1,
			Max:1,
			Boxnum:4000000002,
			Random:1000,
		},
		1019020001:Box{
			Itemid:1019020001,
			Min:1,
			Max:1,
			Boxnum:4000000002,
			Random:1000,
		},
		1019020002:Box{
			Itemid:1019020002,
			Min:1,
			Max:1,
			Boxnum:4000000002,
			Random:1000,
		},
		1019020003:Box{
			Itemid:1019020003,
			Min:1,
			Max:1,
			Boxnum:4000000002,
			Random:1000,
		},
		1010000001:Box{
			Itemid:1010000001,
			Min:50,
			Max:50,
			Boxnum:4000000002,
			Random:5850,
		},
		1010000002:Box{
			Itemid:1010000002,
			Min:10,
			Max:10,
			Boxnum:4000000003,
			Random:10000,
		},
	}
	return &data
}
