package quotes

import (
	"encoding/json"
	"fmt"
	"gitee.com/quant1x/gotdx/proto"
	"github.com/mymmsc/gox/api"
	"testing"
)

func TestSecurityBarsPackage_UnSerialize(t *testing.T) {
	stdApi, err := NewStdApiWithServers([]Server{Server{Host: "123.125.108.14", Port: 7709, Name: "test"}})
	if err != nil {
		panic(err)
	}
	defer stdApi.Close()
	//sq1, err := stdApi.GetSecurityQuotes([]uint8{proto.MarketIdShangHai, proto.MarketIdShangHai, proto.MarketIdShangHai, proto.MarketIdShenZhen}, []string{"600275", "600455", "600086", "300742"})
	sq1, err := stdApi.GetKLine(0, "000966", proto.KLINE_TYPE_RI_K, 0, 1)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", sq1)
	data, _ := json.Marshal(sq1)
	text := api.Bytes2String(data)
	fmt.Println(text)
}
