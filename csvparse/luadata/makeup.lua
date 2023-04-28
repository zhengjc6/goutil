---@class makeup_cfg
---@field public Id integer @道具id
---@field public Type integer @分解/合成
---@field public Makeupnum integer @合成/分解需要的数量
---@field public Makeupid integer @合成/分解需要的id

local empty = {}

local M = {
	[1013020001] = {
		Id = 1013020001,
		Type = 2,
		Makeupnum = 100,
		Makeupid = 1019020001,
	},
	[1013020002] = {
		Id = 1013020002,
		Type = 2,
		Makeupnum = 100,
		Makeupid = 1019020002,
	},
	[1013020003] = {
		Id = 1013020003,
		Type = 2,
		Makeupnum = 100,
		Makeupid = 1019020003,
	},
	[1013010001] = {
		Id = 1013010001,
		Type = 2,
		Makeupnum = 100,
		Makeupid = 1019020010,
	},
	[1013010002] = {
		Id = 1013010002,
		Type = 2,
		Makeupnum = 100,
		Makeupid = 1019020011,
	},
	[1013010003] = {
		Id = 1013010003,
		Type = 2,
		Makeupnum = 100,
		Makeupid = 1019020012,
	},
}
return M
