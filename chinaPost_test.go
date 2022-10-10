package chinaPost

import (
	"log"
	"testing"
	"time"
)

func TestChinaPost_Digest(t *testing.T) {
	chinaPost := &ChinaPost{}
	dig := chinaPost.Digest("1100061201282", "687scacnlsf35r1c")
	log.Println(dig)
}

func TestChinaPost_CreateOrder(t *testing.T) {
	chinaPost := &ChinaPost{}
	chinaPostCreateOrderRequest := &ChinaPostCreateOrderRequest{}
	chinaPostLogistics := ChinaPostLogistics{}
	chinaPost.EcCompanyId = "1100061201282"
	chinaPost.Key = "687scacnlsf35r1c"
	chinaPostLogistics.Sender_no = "1100061201282"
	chinaPostLogistics.MailType = "glkj"
	chinaPostLogistics.Barcode = "gl829347892374"
	chinaPostLogistics.Battery_flag = "0"
	chinaPostLogistics.Batch_no = ""
	chinaPostLogistics.Biz_product_no = "2134214"
	timeNow := time.Now()
	chinaPostLogistics.Created_time = timeNow.Format("2006-01-02 15:04:05")
	chinaPostLogistics.Logistics_order_id = "GL3242342"
	chinaPostLogistics.Contents_total_value = 5000
	chinaPostLogistics.Declare_curr_code = "USD"
	chinaPostLogistics.Declare_source = "2"
	chinaPostLogistics.Forecastshut = "0"
	chinaPostLogistics.Length = 1.0
	chinaPostLogistics.Width = 1.0
	chinaPostLogistics.Height = 1.0
	chinaPostLogistics.Declare_type = "1"
	chinaPostLogistics.Insurance_amount = 0
	chinaPostLogistics.Contents_total_weight = 600
	chinaPostLogistics.Insurance_flag = ""
	chinaPostLogistics.Mail_sign = "2"
	chinaPostLogistics.Pickup_notes = ""
	chinaPostLogistics.Prepayment_of_vat = ""
	chinaPostLogistics.Tax_id = ""
	chinaPostLogistics.Postage_currency = "USD"
	chinaPostLogistics.Printcode = ""
	chinaPostLogistics.Undelivery_option = 2
	chinaPostLogistics.Transfer_type = "HK"
	chinaPostLogistics.Volume = 1.0
	chinaPostLogistics.Weight = 600
	chinaPostLogistics.Valuable_flag = "0"
	chinaPostLogistics.Wh_code = "234234"
	sender := ChinaPostAddress{}
	sender.City = "ShenZhen"
	sender.Address = "ewrwer"
	sender.Company = "GLKJ"
	sender.Name = "test"
	sender.Linker = "test"
	sender.Nation = "CN"
	sender.Phone = "56465465"
	sender.Mobile = "64687468"
	sender.Post_code = "518118"
	sender.Province = "GuangDong"
	chinaPostLogistics.Sender = sender
	recevier := ChinaPostAddress{}
	chinaPostLogistics.Receiver = recevier
	chinaPostItem := ChinaPostItem{}
	var items []ChinaPostItem
	items = append(items, chinaPostItem)
	chinaPostLogistics.Items = items
	chinaPostCreateOrderRequest.Logistics_interface = chinaPostLogistics
	resp := chinaPost.CreateOrder(chinaPostCreateOrderRequest)
	log.Println(resp)
}
