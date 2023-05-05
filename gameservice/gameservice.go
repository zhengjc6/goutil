package gameservice

import (
	"encoding/json"
	"fmt"
	"io"
	"metaserver/pkg/errorx"
	"net/http"
	"strconv"
	"strings"
)

func (l *GameService) send2Game(cmd *string) (*string, error) {
	fmt.Printf("go send to game:%s\n", *cmd)
	client := &http.Client{}
	req, err := http.NewRequest("POST", l.GamePath, strings.NewReader(*cmd))
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	defer req.Body.Close()
	req.Header.Set("Content-Type", l.BodyType)
	req.Header.Set("Connection", "keep-alive")

	rep, err := client.Do(req)
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	if rep.StatusCode != 200 {
		return nil, errorx.NewDefaultError("httpe statuscode:" + strconv.FormatInt(int64(rep.StatusCode), 10))
	}
	data, err := io.ReadAll(rep.Body)
	rep.Body.Close()
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}
	rst := gameRet{}
	err = json.Unmarshal(data, &rst)
	if err != nil {
		return nil, errorx.NewDefaultError("game ret Unmarshal fail body:" + string(data))
	}
	if rst.Ret != 0 {
		return nil, errorx.NewDefaultError("game ret fail code:" + strconv.FormatInt(int64(rst.Ret), 10))
	}
	return &(rst.Info), nil
}

func (l *GameService) CostOneItem(uid int64, item *GameItem, reason int32) error {
	var itemlist []GameItem
	itemlist = append(itemlist, *item)
	return l.CostItems(uid, &itemlist, reason)
}

func (l *GameService) CostItems(uid int64, itemlist *[]GameItem, reason int32) error {
	jsonItems, err := json.Marshal(itemlist)
	if err != nil {
		return errorx.NewDefaultError("json marshal fail")
	}
	sendStr := fmt.Sprintf("U%d game_del_items %d %s", uid, reason, jsonItems)
	_, err = l.send2Game(&sendStr)
	return err
}

func (l *GameService) AddOneItem(uid int64, item *GameItem, reason int32) error {
	var itemlist []GameItem
	itemlist = append(itemlist, *item)
	return l.AddItems(uid, &itemlist, reason)
}
func (l *GameService) AddItems(uid int64, itemlist *[]GameItem, reason int32) error {
	jsonItems, err := json.Marshal(itemlist)
	if err != nil {
		return errorx.NewDefaultError("json marshal fail")
	}
	sendStr := fmt.Sprintf("U%d game_add_items %d %s", uid, reason, jsonItems)
	_, err = l.send2Game(&sendStr)
	return err
}

// w h 填 0 代表使用默认值
func (l *GameService) GetAreaPlayers(uid int64, w, h int32) (*([]int64), error) {
	sendStr := fmt.Sprintf("S1 citymap maincity_area_players %d %d %d", uid, w, h)
	strList, err := l.send2Game(&sendStr)
	uidList := []int64{}
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(*strList), &uidList)
	if err != nil {
		return nil, errorx.NewDefaultError("json marshal fail")
	}
	return &uidList, nil
}

func (l *GameService) GetPlayerChatInfo(uid int64) (*ChatBaseInfo, error) {
	sendStr := fmt.Sprintf("S1 localdata gmt_chat_baseinfo %d", uid)
	rspStr, err := l.send2Game(&sendStr)
	if err != nil {
		return nil, err
	}
	info := &ChatBaseInfo{}
	err = json.Unmarshal([]byte(*rspStr), info)
	if err != nil {
		fmt.Printf("%v\n", *rspStr)
		fmt.Println(err)
		return nil, errorx.NewDefaultError("json marshal fail")
	}
	return info, nil
}

func (l *GameService) GetItems(uid int64, itemlist *[]GameItem) (*[]GameItem, error) {
	jsonItems, err := json.Marshal(itemlist)
	if err != nil {
		return nil, errorx.NewDefaultError("json marshal fail")
	}
	sendStr := fmt.Sprintf("U%d game_get_items %s", uid, jsonItems)
	rspStr, err := l.send2Game(&sendStr)
	if err != nil {
		return nil, err
	}
	info := &[]GameItem{}
	err = json.Unmarshal([]byte(*rspStr), info)
	if err != nil {
		fmt.Printf("%v\n", *rspStr)
		fmt.Println(err)
		return nil, errorx.NewDefaultError("json unmarshal fail")
	}
	return info, nil
}

func (l *GameService) GetOneItem(uid, itemid, uniqueid int64) (*GameItem, error) {
	item := GameItem{
		CfgID:    int32(itemid),
		UniqueID: uniqueid,
	}
	itemlist := []GameItem{}
	itemlist = append(itemlist, item)
	rstlist, err := l.GetItems(uid, &itemlist)
	if err != nil {
		return nil, err
	}
	rstItem := (*rstlist)[0]
	return &rstItem, nil
}

func (l *GameService) SendLandLevel(uid int64, level int32) error {
	sendStr := fmt.Sprintf("U%d send_land_level %d", uid, level)
	_, err := l.send2Game(&sendStr)
	return err
}

func (l *GameService) AddGold(uid int64, num, reason int32) error {
	goldItem := &GameItem{
		CfgID: GameGoldID,
		Count: num,
	}
	return l.AddOneItem(uid, goldItem, reason)
}

func (l *GameService) AddDiamond(uid int64, num, reason int32) error {
	goldItem := &GameItem{
		CfgID: GameDiamondID,
		Count: num,
	}
	return l.AddOneItem(uid, goldItem, reason)
}

func (l *GameService) GetGold(uid int64) (*GameItem, error) {
	return l.GetOneItem(uid, int64(GameGoldID), 0)
}
func (l *GameService) GetDiamond(uid int64) (*GameItem, error) {
	return l.GetOneItem(uid, int64(GameDiamondID), 0)
}

func (l *GameService) CostGold(uid int64, num, reason int32) error {
	costItem := &GameItem{
		CfgID: GameGoldID,
		Count: num,
	}
	return l.CostOneItem(uid, costItem, reason)
}
func (l *GameService) CostDiamond(uid int64, num, reason int32) error {
	costItem := &GameItem{
		CfgID: GameDiamondID,
		Count: num,
	}
	return l.CostOneItem(uid, costItem, reason)
}

func (l *GameService) GetPlayerBaseInfo(uid int64) (*PlayerBaseInfo, error) {
	sendStr := fmt.Sprintf("S1 localdata gmt_player_baseinfo %d", uid)
	rspStr, err := l.send2Game(&sendStr)
	if err != nil {
		return nil, err
	}
	info := &PlayerBaseInfo{}
	err = json.Unmarshal([]byte(*rspStr), info)
	if err != nil {
		fmt.Printf("%v\n", *rspStr)
		fmt.Println(err)
		return nil, errorx.NewDefaultError("json marshal fail")
	}
	return info, nil
}

func NewGameService(gamePath string) *GameService {
	return &GameService{
		GamePath: gamePath,
		BodyType: "text/plain"}
}
