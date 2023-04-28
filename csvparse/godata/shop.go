package godata
type Shop struct {
	Itemid		int64		//物品id
	Npcid		int64		//商人id
	Buymore		int64		//重复购买
	Stocknumber		int64		//库存数量
	Moneytype		int64		//购买货币类型
	Moneynumber		int64		//购买货币数量
					//
}
type MapShop map[int64]Shop
func CreateShopTable() *MapShop {
	data := MapShop{
		1013020001:Shop{
			Itemid:1013020001,
			Npcid:1033,
			Buymore:2,
			Stocknumber:1,
			Moneytype:1019020001,
			Moneynumber:99,
		},
		1013020002:Shop{
			Itemid:1013020002,
			Npcid:1033,
			Buymore:2,
			Stocknumber:1,
			Moneytype:1019020002,
			Moneynumber:99,
		},
		1013020003:Shop{
			Itemid:1013020003,
			Npcid:1033,
			Buymore:2,
			Stocknumber:1,
			Moneytype:1019020003,
			Moneynumber:99,
		},
		1014000002:Shop{
			Itemid:1014000002,
			Npcid:1033,
			Buymore:2,
			Stocknumber:1000,
			Moneytype:2,
			Moneynumber:100,
		},
		1014000004:Shop{
			Itemid:1014000004,
			Npcid:1033,
			Buymore:2,
			Stocknumber:1000,
			Moneytype:2,
			Moneynumber:100,
		},
		1014000005:Shop{
			Itemid:1014000005,
			Npcid:1033,
			Buymore:2,
			Stocknumber:1000,
			Moneytype:2,
			Moneynumber:100,
		},
	}
	return &data
}
