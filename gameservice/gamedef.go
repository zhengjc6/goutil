package gameservice

type GameDropType int32

type GameItem struct {
	CfgID      int32  `json:"cfgid"`
	Count      int32  `json:"cnt"`
	UniqueID   int64  `json:"uniqueid"` //扣除不可堆叠物品时必填
	RawGoodsID int64  `json:"goodsid"`  //订单类物品关联的初始上架商品id
	CDK        string `json:"cdk"`
	Isused     bool   `json:"isused"`
}

// game资源添加原因
type GameAddResouceType int32

const (
	GameAddResouceType_GoodsOffShelf   GameAddResouceType = 10001 ////商品下架返还
	GameAddResouceType_GoodsBuyFail    GameDelResouceType = 10002 //商品购买失败回退
	GameAddResouceType_GoodsSubmitFail GameDelResouceType = 10003 //商品上架失败返还
	GameAddResouceType_GoodsBuySuc     GameDelResouceType = 10004 //商品购买获得
	GameAddResouceType_LandProduce     GameDelResouceType = 10005 //土地产出
)

// game资源删除原因
type GameDelResouceType int32

const (
	GameDelResouceType_GoodsSubmit  GameDelResouceType = 15001 //商品上架
	GameDelResouceType_GoodsBuySuc  GameDelResouceType = 15002 //商品购买消耗
	GameDelResouceType_GoodsBuyFail GameDelResouceType = 10002 //商品购买失败收回
	GameDelResouceType_LandUpgrade  GameDelResouceType = 15003 //土地升级
)

type GameService struct {
	GamePath string
	BodyType string
}

type gameRet struct {
	Ret  int32  `json:"ret"`
	Info string `json:"info"`
}

type ChatBaseInfo struct {
	Name      string `json:"name"`      // 名称
	Img       string `json:"img"`       // 头像
	Gender    int32  `json:"gender"`    //性别
	HeadFrame int32  `json:"headframe"` // 头像框
}

type AvatarPortion struct {
	PortionID int32 `json:"portionid"`
	CfgID     int32 `json:"cfgid"`
	UniqueID  int64 `json:"uniqueid"`
}

type AvatarData struct {
	PortionList []AvatarPortion `json:"portionlist"`
	FaceList    []AvatarPortion `json:"facelist"`
}

type PlayerBaseInfo struct {
	Name      string     `json:"name"`      // 名称
	Img       string     `json:"img"`       // 头像
	Gender    int32      `json:"gender"`    // 性别
	HeadFrame int32      `json:"headframe"` // 头像框
	CurAvatr  AvatarData `json:"curavatr"`  // 形象
}

const GameGoldID int32 = 1010000001
const GameDiamondID int32 = 1010000002
