package chinaPost

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/wms3001/goCommon"
	"github.com/wms3001/goTool"
	"github.com/wms3001/http"
	"log"
)

type ChinaPost struct {
	EcCompanyId string `json:"ecCompanyId"`
	Key         string `json:"key"`
}

var url = "https://my.ems.com.cn/pcpErp-web/a/pcp"

func (chinaPost *ChinaPost) CreateOrder(request *ChinaPostCreateOrderRequest) *goCommon.Resp {
	m := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded;charset=utf-8",
	}
	a := map[string]string{
		//"Accept": "application/json",
		//"Accept": "text/plain",
	}
	var chinaPostOrderRequest ChinaPostOrderRequest
	logistics_interface, _ := json.Marshal(request.Logistics_interface)
	chinaPostOrderRequest.EcCompanyId = chinaPost.EcCompanyId
	chinaPostOrderRequest.Msg_type = "B2C_TRADE"
	chinaPostOrderRequest.Data_type = "JSON"
	chinaPostOrderRequest.Data_digest = chinaPost.Digest(string(logistics_interface), chinaPost.Key)
	chinaPostOrderRequest.Logistics_interface = string(logistics_interface)
	param, _ := json.Marshal(chinaPostOrderRequest)
	req := string(param)
	//req = "{\"data_digest\":\"tsGFXuAJ8o0TQFJjaFhvcw==\",\"msg_type\":\"B2C_TRADE\",\"ecCompanyId\":\"11000612012821\",\"data_type\":\"JSON\",\"logistics_interface\":{\"created_time\":\"2022-09-28 13:22:46\",\"sender_no\":\"11000612012821\",\"mailType\":\"GLKJ\",\"wh_code\":\"51810302\",\"logistics_order_no\":\"GL928012246960228\",\"batch_no\":\"\",\"biz_product_no\":\"003\",\"weight\":124.0,\"volume\":1.0,\"length\":1.0,\"width\":1.0,\"height\":1.0,\"postage_total\":null,\"postage_currency\":\"USD\",\"contents_total_weight\":124.0,\"contents_total_value\":919.9999999999999,\"transfer_type\":\"HK\",\"battery_flag\":\"0\",\"pickup_notes\":\"\",\"insurance_flag\":\"\",\"insurance_amount\":0.0,\"undelivery_option\":2,\"valuable_flag\":0,\"declare_source\":\"2\",\"declare_type\":1,\"declare_curr_code\":\"USD\",\"printcode\":\"\",\"barcode\":\"GL928012246960228\",\"forecastshut\":\"0\",\"mail_sign\":\"2\",\"tax_id\":null,\"s_tax_id\":\"IM4420001201\",\"prepayment_of_vat\":\"0\",\"sender\":{\"name\":\"ZhouLi\",\"company\":\"\",\"post_code\":\"510000\",\"phone\":\"8613128939352\",\"mobile\":\"8613128939352\",\"email\":null,\"id_type\":\"\",\"id_no\":null,\"nation\":\"CN\",\"province\":\"Guangdong\",\"city\":\"Shenzhen\",\"county\":null,\"address\":\"Building D, Huafeng High-tech Industrial Park, East Gate, No. 3, Keji Road, Da Industrial Zone, Pingshan New District\",\"gis\":null,\"linker\":\"ZhouLi\"},\"pickup\":null,\"receiver\":{\"name\":\"Carine Poisson\",\"company\":\"Carine Poisson\",\"post_code\":\"89510\",\"phone\":\"0616031290\",\"mobile\":\"0616031290\",\"email\":\"Carine Poisson@gmail.com\",\"id_type\":\"\",\"id_no\":\"\",\"nation\":\"FR\",\"province\":\"Veron\",\"city\":\"Veron\",\"county\":\"\",\"address\":\"1 rue du puits Valperonne \",\"gis\":\"\",\"linker\":\"Carine Poisson\"},\"items\":[{\"cargo_no\":\"TW13207A1\",\"cargo_name\":\"布鲁狗公仔\",\"cargo_name_en\":\"Brugo\",\"cargo_type_name\":\"布鲁狗公仔\",\"cargo_type_name_en\":\"Brugo\",\"cargo_origin_name\":\"CN\",\"cargo_link\":\"\",\"cargo_quantity\":1,\"cargo_value\":9.2,\"cost\":9.2,\"cargo_currency\":\"USD\",\"carogo_weight\":114.0,\"cargo_description\":\"Brugo\",\"cargo_serial\":\"9503002100\",\"unit\":\"个\",\"intemsize\":\"\"}]}}"
	log.Println(req)
	var wHttp = http.Http{
		url + "/orderService/OrderReceiveBack",
		"",
		"",
		m,
		a,
		req,
		10,
		"form",
		[]byte{},
	}
	resp := wHttp.Post()
	return resp
}

func (chinaPost *ChinaPost) Digest(str string, key string) string {
	tool := &goTool.GoTool{}
	//m := md5.New()
	//m.Write([]byte(str))
	//res := hex.EncodeToString(m.Sum(nil))
	sign := md5.Sum([]byte(str))
	signStr := fmt.Sprintf("%x", sign)
	log.Println(signStr)
	return tool.StrToBase64(signStr)
}
