package model

//ExchangeOrder 兑换订单表
type ExchangeOrder struct {
	UID     *int64              `bson:"uid"`    //用户ID
	OrderID string             `bson:"_id"`    //订单号

	PrizeID        int32  `bson:"prize_id"`         //奖品ID
	PrizeName      string `bson:"prize_name"`       //奖品名称
	PrizeImg       string `bson:"prize_img"`        //奖品图片地址
	PrizeNum       int32  `bson:"prize_num"`        //奖品数量
	PrizePrice     int64  `bson:"prize_price"`      //奖品价格
	PrizeType      int32  `bson:"prize_type"`       //拓展类型（服务端使用）
	PrizeExtraInfo string `bson:"prize_extra_info"` //拓展信息（服务端使用）
	TotalPrize     int64  `bson:"total_prize"`      //总价

	Rule     ExchangeRule   `bson:"rules"`     //兑换券组合
	PropInfo []ExchangeProp `bson:"prop_info"` //兑换券信息

	SubOrders []SubOrder `bson:"sub_orders"` //兑换券信息

	TakeCode          string `bson:"take_code"`            //提取码
	TakeCodeExpiredAt int64  `bson:"take_code_expired_at"` //提取码失效时间
	TakeSiteName      string `bson:"take_site_name"`       //自提点名称
	TakeSiteAddr      string `bson:"take_site_addr"`       //自提点地址

	ExpressCode   string          `bson:"express_code"`   //快递单号
	ConsigneeName string          `bson:"consignee_name"` //收货人姓名
	ConsigneeTel  string          `bson:"consignee_tel"`  //收货人电话
	ConsigneeAddr string          `bson:"consignee_addr"` //收货人地址
	LogicInfo     []LogicInfoItem `bson:"logic_info"`     //用户提交信息

	CreatedAt   int64 `bson:"created_at"`   //创建时间
	ExchangedAt int64 `bson:"exchanged_at"` //兑换时间
	ProcessedAt int64 `bson:"processed_at"` //处理时间
	FinishedAt  int64 `bson:"finished_at"`  //完成时间
	Boom int
}

//SubOrder 话费子订单
type SubOrder struct {
	ChargeID    int64                        `bson:"charge_id"`
	Fee         int32                        `bson:"fee"`
	Mobile      string                       `bson:"mobile"`
}

//ExchangeRule 兑换方案
type ExchangeRule struct {
	Items       []*VoucherItem `bson:"voucher_group"` //兑换券组合
	DiamondNum  int32          `bson:"diamond_num"`   //钻石数量
	PropIDGroup string         `bson:"prop_id_group"` //rule唯一标识
}

//VoucherItem 兑换券Item
type VoucherItem struct {
	PropID  int32 `bson:"prop_id"`  //道具ID
	PropNum int32 `bson:"prop_num"` //道具数量
}

//ExchangeProp 订单中道具信息
type ExchangeProp struct {
	PropID   int32  `bson:"prop_id"`   //道具ID
	PropName string `bson:"prop_name"` //道具名称
	SmallImg string `bson:"small_img"` //列表小图片地址
}

type LogicInfoItem struct {
	Key   string `bson:"key"`
	Desc  string `bson:"desc"`
	Value string `bson:"value"`
}
