package godata
type ItemAnimation struct {
	Id		int64		//道具id
	Effectid		string		//使用时的特效id
	Modelname		string		//模型名称
	Animatorid		int64		//使用时的动画id
					//
}
type MapItemAnimation map[int64]ItemAnimation
func CreateItemAnimationTable() *MapItemAnimation {
	data := MapItemAnimation{
		1014000001:ItemAnimation{
			Id:1014000001,
			Effectid:"PropOne1014000001",
			Modelname:"Goods_1014000001",
			Animatorid:0,
		},
		1014000002:ItemAnimation{
			Id:1014000002,
			Effectid:"100002",
			Modelname:"Goods_1014000002",
			Animatorid:300009,
		},
		1014000004:ItemAnimation{
			Id:1014000004,
			Effectid:"100004",
			Modelname:"Goods_1014000004",
			Animatorid:300009,
		},
		1014000005:ItemAnimation{
			Id:1014000005,
			Effectid:"100003",
			Modelname:"Goods_1014000005",
			Animatorid:300009,
		},
	}
	return &data
}
