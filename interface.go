package game_center_client

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Client struct {
	Url string
	lg  func(logContext string)
}

func Init(url string, log func(logContext string)) *Client {
	return &Client{
		Url: url,
		lg:  log,
	}
}

// 约牌回调，补充约牌的房间信息
func (c *Client) YuepaiInviteCallback(ucId, id, matchId, seatDown, status int64) (err error) {
	request := make(map[string]interface{})

	request["id"] = id
	request["uc_id"] = ucId
	request["match_id"] = matchId
	request["seat_down"] = seatDown
	request["status"] = status

	uri, _ := url.JoinPath(c.Url, "/api/yuepai/invite_callback")
	d, _ := json.Marshal(request)
	c.lg(fmt.Sprintf("Request Uri:%s, body:%s", uri, string(d)))
	res, err := postRequest(uri, request)
	if err != nil {
		return err
	}
	c.lg(fmt.Sprintf("Response:%+v", res))
	if res.Code != 0 {
		return fmt.Errorf("code:%d, message:%s", res.Code, res.Message)
	}

	return nil
}
