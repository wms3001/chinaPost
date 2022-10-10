package chinaPost

type ChinaPostAddress struct {
	Name      string `json:"name"`
	Company   string `json:"company"`
	Post_code string `json:"post_code"`
	Phone     string `json:"phone"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Id_type   string `json:"id_type"`
	Id_no     string `json:"id_no"`
	Nation    string `json:"nation"`
	Province  string `json:"province"`
	City      string `json:"city"`
	County    string `json:"county"`
	Address   string `json:"address"`
	Gis       string `json:"gis"`
	Linker    string `json:"linker"`
}

type ChinaPostItem struct {
	Cargo_no           string  `json:"cargo_no"`
	Cargo_name         string  `json:"cargo_name"`
	Cargo_name_en      string  `json:"cargo_name_en"`
	Cargo_type_name    string  `json:"cargo_type_name"`
	Cargo_type_name_en string  `json:"cargo_type_name_en"`
	Cargo_origin_name  string  `json:"cargo_origin_name"`
	Cargo_link         string  `json:"cargo_link"`
	Cargo_quantity     int     `json:"cargo_quantity"`
	Cargo_value        float64 `json:"cargo_value"`
	Cost               float64 `json:"cost"`
	Cargo_currency     string  `json:"cargo_currency"`
	Cargo_weight       float64 `json:"cargo_weight"`
	Cargo_description  string  `json:"cargo_description"`
	Cargo_serial       string  `json:"cargo_serial"`
	Unit               string  `json:"unit"`
	Intemsize          string  `json:"intemsize"`
}

type ChinaPostLogistics struct {
	Created_time          string           `json:"created_time"`
	Sender_no             string           `json:"sender_no"`
	Wh_code               string           `json:"wh_code"`
	Logistics_order_id    string           `json:"logistics_order_id"`
	TradeId               string           `json:"tradeId"`
	LogisticsCompany      string           `json:"logisticsCompany"`
	LogisticsBiz          string           `json:"logisticsBiz"`
	MailType              string           `json:"mailType"`
	FaceType              int              `json:"faceType"`
	Rcountry              string           `json:"rcountry"`
	Batch_no              string           `json:"batch_no"`
	Biz_product_no        string           `json:"biz_product_no"`
	Weight                float64          `json:"weight"`
	Volume                float64          `json:"volume"`
	Length                float64          `json:"length"`
	Width                 float64          `json:"width"`
	Height                float64          `json:"height"`
	Postage_total         float64          `json:"postage_total"`
	Postage_currency      string           `json:"postage_currency"`
	Contents_total_weight float64          `json:"contents_total_weight"`
	Contents_total_value  float64          `json:"contents_total_value"`
	Transfer_type         string           `json:"transfer_type"`
	Battery_flag          string           `json:"battery_flag"`
	Pickup_notes          string           `json:"pickup_notes"`
	Insurance_flag        string           `json:"insurance_flag"`
	Insurance_amount      float64          `json:"insurance_amount"`
	Undelivery_option     int              `json:"undelivery_option"`
	Back_addr             string           `json:"back_addr"`
	Back_way              string           `json:"back_way"`
	Valuable_flag         string           `json:"valuable_flag"`
	Declare_source        string           `json:"declare_source"`
	Declare_type          string           `json:"declare_type"`
	Declare_curr_code     string           `json:"declare_curr_code"`
	Printcode             string           `json:"printcode"`
	Barcode               string           `json:"barcode"`
	Forecastshut          string           `json:"forecastshut"`
	Mail_sign             string           `json:"mail_sign"`
	Mail_flag             string           `json:"mail_flag"`
	Tax_id                string           `json:"tax_id"`
	S_tax_id              string           `json:"s_tax_id"`
	Prepayment_of_vat     string           `json:"prepayment_of_vat"`
	Pickup_flag           int              `json:"pickup_flag"`
	Sender                ChinaPostAddress `json:"sender"`
	Receiver              ChinaPostAddress `json:"receiver"`
	Pickup                ChinaPostAddress `json:"pickup"`
	Items                 []ChinaPostItem  `json:"items"`
}

type ChinaPostCreateOrderRequest struct {
	EcCompanyId         string             `json:"ecCompanyId"`
	Data_digest         string             `json:"data_digest"`
	Msg_type            string             `json:"msg_type"`
	Data_type           string             `json:"data_type"`
	Logistics_interface ChinaPostLogistics `json:"logistics_interface"`
	Biz_product_no      string             `json:"biz_product_no"`
}

type ChinaPostOrderRequest struct {
	EcCompanyId         string `json:"ecCompanyId"`
	Data_digest         string `json:"data_digest"`
	Msg_type            string `json:"msg_type"`
	Data_type           string `json:"data_type"`
	Logistics_interface string `json:"logistics_interface"`
	Biz_product_no      string `json:"biz_product_no"`
}
