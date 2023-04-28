---@class shop_cfg
---@field public Itemid integer @物品id
---@field public Npcid integer @商人id
---@field public Buymore integer @重复购买
---@field public Stocknumber integer @库存数量
---@field public Moneytype integer @购买货币类型
---@field public Moneynumber integer @购买货币数量

local empty = {}

local M = {
	[1013020001] = {
		Itemid = 1013020001,
		Npcid = 1033,
		Buymore = 2,
		Stocknumber = 1,
		Moneytype = 1019020001,
		Moneynumber = 99,
	},
	[1013020002] = {
		Itemid = 1013020002,
		Npcid = 1033,
		Buymore = 2,
		Stocknumber = 1,
		Moneytype = 1019020002,
		Moneynumber = 99,
	},
	[1013020003] = {
		Itemid = 1013020003,
		Npcid = 1033,
		Buymore = 2,
		Stocknumber = 1,
		Moneytype = 1019020003,
		Moneynumber = 99,
	},
	[1014000002] = {
		Itemid = 1014000002,
		Npcid = 1033,
		Buymore = 2,
		Stocknumber = 1000,
		Moneytype = 2,
		Moneynumber = 100,
	},
	[1014000004] = {
		Itemid = 1014000004,
		Npcid = 1033,
		Buymore = 2,
		Stocknumber = 1000,
		Moneytype = 2,
		Moneynumber = 100,
	},
	[1014000005] = {
		Itemid = 1014000005,
		Npcid = 1033,
		Buymore = 2,
		Stocknumber = 1000,
		Moneytype = 2,
		Moneynumber = 100,
	},
}
return M
