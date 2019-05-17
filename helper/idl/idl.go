package  idl

type Order struct {
	UID               int64                  `protobuf:"varint,1,req,name=UID" json:"UID"`
	OrderID           string                 `protobuf:"bytes,2,req,name=orderID" json:"orderID"`

	PrizeID           int32                  `protobuf:"varint,4,req,name=prizeID" json:"prizeID"`
	PrizeName         string                 `protobuf:"bytes,5,req,name=prizeName" json:"prizeName"`
	PrizeNum          int32                  `protobuf:"varint,6,req,name=prizeNum" json:"prizeNum"`
	PrizePrice        int64                  `protobuf:"varint,7,req,name=prizePrice" json:"prizePrice"`
	PrizeType         int32                  `protobuf:"varint,8,req,name=prizeType" json:"prizeType"`
	PrizeExtraInfo    string                 `protobuf:"bytes,9,opt,name=prizeExtraInfo" json:"prizeExtraInfo"`
	PrizeImg          string                 `protobuf:"bytes,10,opt,name=prizeImg" json:"prizeImg"`
	Rule              ExchangeRule    `protobuf:"bytes,11,req,name=rule" json:"rule"`
	PropInfo          []PropInfo             `protobuf:"bytes,12,rep,name=propInfo" json:"propInfo"`

	TakeCode          string                 `protobuf:"bytes,14,opt,name=takeCode" json:"takeCode"`
	TakeSiteName      string                 `protobuf:"bytes,15,opt,name=takeSiteName" json:"takeSiteName"`
	TakeSiteAddr      string                 `protobuf:"bytes,16,opt,name=takeSiteAddr" json:"takeSiteAddr"`
	ExpressCode       string                 `protobuf:"bytes,18,opt,name=expressCode" json:"expressCode"`
	ConsigneeName     string                 `protobuf:"bytes,19,opt,name=consigneeName" json:"consigneeName"`
	ConsigneeTel      string                 `protobuf:"bytes,20,opt,name=consigneeTel" json:"consigneeTel"`
	ConsigneeAddr     string                 `protobuf:"bytes,21,opt,name=consigneeAddr" json:"consigneeAddr"`
	SubOrders         []SubOrder             `protobuf:"bytes,22,rep,name=subOrders" json:"subOrders"`
	LogicInfo         []LogicInfoItem `protobuf:"bytes,23,rep,name=logicInfo" json:"logicInfo"`
	TakeCodeExpiredAt int64                  `protobuf:"varint,24,opt,name=takeCodeExpiredAt" json:"takeCodeExpiredAt"`
	CreatedAt         int64                  `protobuf:"varint,25,opt,name=createdAt" json:"createdAt"`
	ExchangedAt       int64                  `protobuf:"varint,26,opt,name=exchangedAt" json:"exchangedAt"`
	ProcessedAt       int64                  `protobuf:"varint,27,opt,name=processedAt" json:"processedAt"`
	FinishedAt        int64                  `protobuf:"varint,28,opt,name=finishedAt" json:"finishedAt"`
	TakeSiteID        int64                  `protobuf:"varint,29,opt,name=takeSiteID" json:"takeSiteID"`
	TotalPrize        int64                  `protobuf:"varint,30,opt,name=totalPrize" json:"totalPrize"`
	XXX_unrecognized  []byte                 `json:"-"`
}

type ExchangeRule struct {
	Items            []*VoucherItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	DiamondNum       int32          `protobuf:"varint,2,req,name=diamondNum" json:"diamondNum"`
	PropIDGroup      string         `protobuf:"bytes,3,req,name=propIDGroup" json:"propIDGroup"`
	XXX_unrecognized []byte         `json:"-"`
}

type PropInfo struct {
	PropID           int32  `protobuf:"varint,1,req,name=propID" json:"propID"`
	PropName         string `protobuf:"bytes,2,req,name=propName" json:"propName"`
	PropImg          string `protobuf:"bytes,3,req,name=propImg" json:"propImg"`
	XXX_unrecognized []byte `json:"-"`
}

type SubOrder struct {
	ChargeID         int64          `protobuf:"varint,1,req,name=chargeID" json:"chargeID"`
	Fee              int32          `protobuf:"varint,3,req,name=fee" json:"fee"`
	Mobile           string         `protobuf:"bytes,4,req,name=mobile" json:"mobile"`
	XXX_unrecognized []byte         `json:"-"`
}
type LogicInfoItem struct {
	Key              string `protobuf:"bytes,1,req,name=key" json:"key"`
	Desc             string `protobuf:"bytes,2,opt,name=desc" json:"desc"`
	Value            string `protobuf:"bytes,3,opt,name=value" json:"value"`
	XXX_unrecognized []byte `json:"-"`
}

type VoucherItem struct {
	PropID           int32  `protobuf:"varint,1,req,name=propID" json:"propID"`
	PropNum          int32  `protobuf:"varint,2,req,name=propNum" json:"propNum"`
	XXX_unrecognized []byte `json:"-"`
}