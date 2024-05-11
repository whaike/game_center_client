package game_center_client

type CommonResp struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	ServerTime int64       `json:"server_time"`
}

// 约牌回调
type YuepaiInviteCallbackReqDTO struct {
	Id       int64 `json:"id" form:"id" validate:"required"`       // 约牌ID
	UcId     int64 `json:"uc_id" form:"uc_id" validate:"required"` // 用户中心ID
	MatchId  int64 `json:"match_id" form:"match_id"`               // 比赛ID
	SeatDown int64 `json:"seat_down" form:"seat_down"`             // 是否坐下（1 离开 2 坐下）
	Status   int64 `json:"status" form:"status"`                   // 约牌状态（0 创建中 1 已创建 2 已开赛 3 已结束 4 已取消）
}
