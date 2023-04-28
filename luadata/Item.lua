---@class item
---@field public Id integer @道具id
---@field public Type integer @道具类型
---@field public Particulartype integer @小道具类型-1：基础道具-2：个性道具-3：礼物道具
---@field public Boxid integer[] @集合包id

local empty = {}

local M = {
	[1010000001] = {
		Id = 1010000001,
		Type = 10,
		Particulartype = 0,
		Boxid = empty,
	},
	[1013020001] = {
		Id = 1013020001,
		Type = 13,
		Particulartype = 0,
		Boxid = empty,
	},
	[1013020002] = {
		Id = 1013020002,
		Type = 13,
		Particulartype = 0,
		Boxid = empty,
	},
	[1013020003] = {
		Id = 1013020003,
		Type = 13,
		Particulartype = 0,
		Boxid = empty,
	},
	[1013010001] = {
		Id = 1013010001,
		Type = 13,
		Particulartype = 0,
		Boxid = empty,
	},
	[1013010002] = {
		Id = 1013010002,
		Type = 13,
		Particulartype = 0,
		Boxid = empty,
	},
	[1013010003] = {
		Id = 1013010003,
		Type = 13,
		Particulartype = 0,
		Boxid = empty,
	},
	[1019020001] = {
		Id = 1019020001,
		Type = 19,
		Particulartype = 0,
		Boxid = empty,
	},
	[1019020002] = {
		Id = 1019020002,
		Type = 19,
		Particulartype = 0,
		Boxid = empty,
	},
	[1019020003] = {
		Id = 1019020003,
		Type = 19,
		Particulartype = 0,
		Boxid = empty,
	},
	[1015010001] = {
		Id = 1015010001,
		Type = 15,
		Particulartype = 0,
		Boxid = {4000000002,4000000003},
	},
	[1015010002] = {
		Id = 1015010002,
		Type = 15,
		Particulartype = 0,
		Boxid = {4000000001},
	},
	[1019180001] = {
		Id = 1019180001,
		Type = 19,
		Particulartype = 0,
		Boxid = empty,
	},
	[1014000001] = {
		Id = 1014000001,
		Type = 14,
		Particulartype = 1,
		Boxid = empty,
	},
	[1014000002] = {
		Id = 1014000002,
		Type = 14,
		Particulartype = 3,
		Boxid = empty,
	},
	[1014000004] = {
		Id = 1014000004,
		Type = 14,
		Particulartype = 2,
		Boxid = empty,
	},
	[1014000005] = {
		Id = 1014000005,
		Type = 14,
		Particulartype = 2,
		Boxid = empty,
	},
}
return M
