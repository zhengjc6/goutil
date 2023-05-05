package gameservice

import (
	"fmt"
	"testing"
)

func TestPluralItems(t *testing.T) {
	GameService := *(NewGameService("http://192.168.1.33:8001/console"))
	var itemlist []GameItem
	item := GameItem{
		CfgID: 1010000001,
		Count: 44,
	}
	itemlist = append(itemlist, item)

	item = GameItem{
		CfgID: 1010000002,
		Count: 33,
	}
	itemlist = append(itemlist, item)

	item = GameItem{
		CfgID: 1013010002,
		Count: 1,
	}
	itemlist = append(itemlist, item)

	item = GameItem{
		CfgID: 1013010002,
		Count: 2,
	}
	itemlist = append(itemlist, item)

	item = GameItem{
		CfgID: 1015010002,
		Count: 2,
	}
	itemlist = append(itemlist, item)

	item = GameItem{
		CfgID: 1015010002,
		Count: 1,
	}
	itemlist = append(itemlist, item)

	err := GameService.AddItems(115651910697914368, &itemlist, 3333)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("AddItems Success")
	}
	// getitems, err := GameService.GetItems(115651910697914368, &itemlist)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("GetItems Success,%+v", getitems)
	// }

	err2 := GameService.CostItems(115651910697914368, &itemlist, 3333)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("AddItems Success")
	}

	getitems2, err := GameService.GetItems(115651910697914368, &itemlist)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("GetItems Success,%+v", getitems2)
	}

}

func TestItems(t *testing.T) {
	GameService := *(NewGameService("http://192.168.1.33:8001/console"))
	var itemlist []GameItem
	item := GameItem{
		CfgID: 1010000001,
		Count: 44,
	}
	itemlist = append(itemlist, item)

	// item = GameItem{
	// 	CfgID:    1013010002,
	// 	Count:    1,
	// 	UniqueID: 122342342424222,
	// }
	// itemlist = append(itemlist, item)

	// item = GameItem{
	// 	CfgID:      1013010002,
	// 	Count:      1,
	// 	UniqueID:   122342342424333,
	// 	RawGoodsID: 999999999,
	// }
	// itemlist = append(itemlist, item)

	item = GameItem{
		CfgID:      1016010001,
		Count:      1,
		UniqueID:   0,
		RawGoodsID: 111111,
	}
	itemlist = append(itemlist, item)

	err := GameService.AddItems(115651910697914368, &itemlist, 3333)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("AddItems Success")
	}
	getitems, err := GameService.GetItems(115651910697914368, &itemlist)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("GetItems Success,%+v", getitems)
	}

	err = GameService.CostItems(115651910697914368, &itemlist, 3333)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("AddItems Success")
	}

	getitems, err = GameService.GetItems(115651910697914368, &itemlist)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("GetItems Success,%+v", getitems)
	}

}

func TestChat(t *testing.T) {
	GameService := *(NewGameService("http://192.168.1.33:8001/console"))
	uidlist, err := GameService.GetAreaPlayers(115651910697914368, 0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetAreaPlayers Success")
		fmt.Println(*uidlist)
	}

	info, err := GameService.GetPlayerChatInfo(115651910697914368)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("GetPlayerChatInfo Success")
		fmt.Printf("%+v\n", *info)
	}
}

func TestLandLevel(t *testing.T) {
	GameService := *(NewGameService("http://192.168.1.33:8001/console"))
	GameService.SendLandLevel(115651910697914368, 1)
}

func TestGold(t *testing.T) {
	GameService := *(NewGameService("http://192.168.1.32:8000/console"))
	gold, err := GameService.GetGold(145995654097702912)
	fmt.Printf("%+v", err)
	fmt.Printf("%+v", gold)
}

func TestPlayerBase(t *testing.T) {
	GameService := *(NewGameService("http://192.168.1.18:8001/console"))
	playerbase, err := GameService.GetPlayerBaseInfo(115651910697914368)
	fmt.Printf("%+v", err)
	fmt.Printf("%+v", playerbase)
}
