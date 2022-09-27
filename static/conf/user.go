package csvdata
type user struct {
	idx	int64
	name	string
	age	int64
	gender	int64
	addr	string
	nothing	string
	intarrary	[]int64
	stringarrary	[]string
	email	string
	ratio	float64
	floatarrary	[]float64
}
type mapuser map[int64]user
func CreateUserTable() mapuser {
	data := mapuser{
				11:user{
			idx:11,
			name:"武则，,天",
			age:0,
			gender:0,
			addr:"家庭地址1",
			nothing:"",
			intarrary:[]int64{1,2,3},
			stringarrary:[]string{},
			email:"safz@aaa1.com",
			ratio:1,
			floatarrary:[]float64{1.11203},
		},
		9:user{
			idx:9,
			name:"name_1",
			age:33,
			gender:1,
			addr:"address2",
			nothing:"",
			intarrary:[]int64{},
			stringarrary:[]string{"你","hello","g狗"},
			email:"",
			ratio:1.001,
			floatarrary:[]float64{1.11203,2.45},
		},
		10001:user{
			idx:10001,
			name:"a你好a",
			age:234680,
			gender:2,
			addr:"",
			nothing:"",
			intarrary:[]int64{77},
			stringarrary:[]string{},
			email:"zzz@aaa.com",
			ratio:1.11203,
			floatarrary:[]float64{},
		},
	}
	return data
}
