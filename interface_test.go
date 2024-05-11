package game_center_client

import (
	"fmt"
	"testing"
)

func Lo(s string) {
	fmt.Println(s)
}

func TestClient_YuepaiInviteCallback(t *testing.T) {
	cli := Init("http://192.168.20.125:13543", Lo)
	err := cli.YuepaiInviteCallback(10275, 1, 0, 1, 0)
	if err != nil {
		Lo(err.Error())
	}
}
