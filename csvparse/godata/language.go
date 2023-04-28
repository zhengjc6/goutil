package godata
type Language struct {
	Id		int64		//道具id
	Name		string		//道具名称
}
type MapLanguage map[int64]Language
func CreateLanguageTable() *MapLanguage {
	data := MapLanguage{
		9010000001:Language{
			Id:9010000001,
			Name:"元石",
		},
		9013020001:Language{
			Id:9013020001,
			Name:"阿里郎的套装-权赫",
		},
		9013020002:Language{
			Id:9013020002,
			Name:"阿里郎的套装-崔金水",
		},
		9013020003:Language{
			Id:9013020003,
			Name:"阿里郎的套装-邓军",
		},
		9013010001:Language{
			Id:9013010001,
			Name:"隔壁泰山-猩猩头饰",
		},
		9013010002:Language{
			Id:9013010002,
			Name:"隔壁泰山-大象头饰",
		},
		9013010003:Language{
			Id:9013010003,
			Name:"隔壁泰山-熊猫头饰",
		},
		9019020001:Language{
			Id:9019020001,
			Name:"阿里郎的套装实物碎片-权赫",
		},
		9019020002:Language{
			Id:9019020002,
			Name:"阿里郎的套装实物碎片-权赫",
		},
		9019020003:Language{
			Id:9019020003,
			Name:"阿里郎的套装碎实物片-权赫",
		},
		9015010001:Language{
			Id:9015010001,
			Name:"阿里郎礼盒",
		},
		9015010002:Language{
			Id:9015010002,
			Name:"元宇宙演唱会礼盒",
		},
		9019180001:Language{
			Id:9019180001,
			Name:"阿里郎签名专辑碎片",
		},
		9014000001:Language{
			Id:9014000001,
			Name:"星星",
		},
		9014000002:Language{
			Id:9014000002,
			Name:"应援灯-超级明星",
		},
		9014000004:Language{
			Id:9014000004,
			Name:"烟花-地球上的星",
		},
		9014000005:Language{
			Id:9014000005,
			Name:"烟花-兔年烟花",
		},
		9110000001:Language{
			Id:9110000001,
			Name:"鲸宇宙世界的通用货币。",
		},
		9113020001:Language{
			Id:9113020001,
			Name:"阿里郎乐队同款套装，穿上它，跟随魔性的旋律，狂热的舞步，一起唱跳《隔壁泰山》吧。/n2768K2NM99D-通过该SDK可去阿里郎官方主页兑换阿里郎隔壁泰山同款服装一套。",
		},
		9113020002:Language{
			Id:9113020002,
			Name:"阿里郎乐队同款套装，穿上它，跟随魔性的旋律，狂热的舞步，一起唱跳《隔壁泰山》吧。/n2768K2NM99D-通过该SDK可去阿里郎官方主页兑换阿里郎隔壁泰山同款服装一套。",
		},
		9113020003:Language{
			Id:9113020003,
			Name:"阿里郎乐队同款套装，穿上它，跟随魔性的旋律，狂热的舞步，一起唱跳《隔壁泰山》吧。/n2768K2NM99D-通过该SDK可去阿里郎官方主页兑换阿里郎隔壁泰山同款服装一套。",
		},
		9113010001:Language{
			Id:9113010001,
			Name:"隔壁泰山同款头饰，穿上它，跟随魔性的旋律，狂热的舞步，一起唱跳《隔壁泰山》吧。",
		},
		9113010002:Language{
			Id:9113010002,
			Name:"隔壁泰山同款头饰，穿上它，跟随魔性的旋律，狂热的舞步，一起唱跳《隔壁泰山》吧。",
		},
		9113010003:Language{
			Id:9113010003,
			Name:"隔壁泰山同款头饰，穿上它，跟随魔性的旋律，狂热的舞步，一起唱跳《隔壁泰山》吧。",
		},
		9119020001:Language{
			Id:9119020001,
			Name:"阿里郎乐队同款套装兑换卷碎片，凑齐30张可合成兑换卷兑换同款实物上衣，填写地址时可注明型号。",
		},
		9119020002:Language{
			Id:9119020002,
			Name:"阿里郎乐队同款套装兑换卷碎片，凑齐30张可合成兑换卷兑换同款实物上衣，填写地址时可注明型号。",
		},
		9119020003:Language{
			Id:9119020003,
			Name:"阿里郎乐队同款套装兑换卷碎片，凑齐30张可合成兑换卷兑换同款实物上衣，填写地址时可注明型号。",
		},
		9115010001:Language{
			Id:9115010001,
			Name:"任务-隔壁泰山的奖励道具，开启后可随机获得礼盒内的奖励。",
		},
		9115010002:Language{
			Id:9115010002,
			Name:"演唱会专属礼盒，开启后可获得神秘礼物。",
		},
		9119180001:Language{
			Id:9119180001,
			Name:"集齐100个阿里郎签名专辑碎片可合成1个阿里郎签名专辑。在阿里郎-权赫处兑换实体阿里郎签名专辑1张。",
		},
		9114000001:Language{
			Id:9114000001,
			Name:"开启阿里郎演唱会的必要道具，累计赠送100颗星星给阿里郎-权赫，可直接开启阿里郎演唱会。",
		},
		9114000002:Language{
			Id:9114000002,
			Name:"对自己喜欢的超级明星使用后激活炫酷的动作与特效。",
		},
		9114000004:Language{
			Id:9114000004,
			Name:"使用后可释放烟花“地球上的星”。",
		},
		9114000005:Language{
			Id:9114000005,
			Name:"使用后可释放烟花“兔年烟花”。",
		},
	}
	return &data
}