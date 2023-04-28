---@class box_cfg
---@field public Itemid integer @物品id
---@field public Min integer @最小数量
---@field public Max integer @最大数量
---@field public Boxnum integer @集合包id
---@field public Random integer @权重

local empty = {}

local M = {
	[1014000002] = {
		Itemid = 1014000002,
		Min = 1,
		Max = 1,
		Boxnum = 4000000001,
		Random = 2000,
	},
	[1014000004] = {
		Itemid = 1014000004,
		Min = 1,
		Max = 1,
		Boxnum = 4000000001,
		Random = 3500,
	},
	[1014000005] = {
		Itemid = 1014000005,
		Min = 1,
		Max = 1,
		Boxnum = 4000000001,
		Random = 4500,
	},
	[1013010001] = {
		Itemid = 1013010001,
		Min = 1,
		Max = 1,
		Boxnum = 4000000002,
		Random = 50,
	},
	[1013010002] = {
		Itemid = 1013010002,
		Min = 1,
		Max = 1,
		Boxnum = 4000000002,
		Random = 50,
	},
	[1013010003] = {
		Itemid = 1013010003,
		Min = 1,
		Max = 1,
		Boxnum = 4000000002,
		Random = 50,
	},
	[1019180001] = {
		Itemid = 1019180001,
		Min = 1,
		Max = 1,
		Boxnum = 4000000002,
		Random = 1000,
	},
	[1019020001] = {
		Itemid = 1019020001,
		Min = 1,
		Max = 1,
		Boxnum = 4000000002,
		Random = 1000,
	},
	[1019020002] = {
		Itemid = 1019020002,
		Min = 1,
		Max = 1,
		Boxnum = 4000000002,
		Random = 1000,
	},
	[1019020003] = {
		Itemid = 1019020003,
		Min = 1,
		Max = 1,
		Boxnum = 4000000002,
		Random = 1000,
	},
	[1010000001] = {
		Itemid = 1010000001,
		Min = 50,
		Max = 50,
		Boxnum = 4000000002,
		Random = 5850,
	},
	[1010000002] = {
		Itemid = 1010000002,
		Min = 10,
		Max = 10,
		Boxnum = 4000000003,
		Random = 10000,
	},
}
return M
