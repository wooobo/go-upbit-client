package socket

type SubscriptionType string

const (
	TypeTicker    SubscriptionType = "ticker"
	TypeTrade     SubscriptionType = "trade"
	TypeOrderbook SubscriptionType = "orderbook"
	TypeMyOrder   SubscriptionType = "myOrder"
	TypeMyAsset   SubscriptionType = "myAsset"
)

type Request struct {
	Ticket         string           `json:"ticket,omitempty"`
	Type           SubscriptionType `json:"type,omitempty"`
	Codes          []string         `json:"codes"`
	IsOnlySnapshot bool             `json:"isOnlySnapshot,omitempty"`
	IsOnlyRealtime bool             `json:"isOnlyRealtime,omitempty"`
	Format         string           `json:"format,omitempty"`
}

type TicketConfig struct {
	Ticket string `json:"ticket"`
}
type TypeField struct {
	Ticket         string           `json:"ticket,omitempty"`
	Type           SubscriptionType `json:"type"`
	Codes          []string         `json:"codes,omitempty"`
	IsOnlySnapshot bool             `json:"isOnlySnapshot,omitempty"`
	IsOnlyRealtime bool             `json:"isOnlyRealtime,omitempty"`
	Level          *float64         `json:"level,omitempty"`
}

type FormatField struct {
	Format string `json:"format,omitempty"`
}
type TickerResponse struct {
	Type              string  `json:"type"`
	Code              string  `json:"code"`
	OpeningPrice      float64 `json:"opening_price"`
	HighPrice         float64 `json:"high_price"`
	LowPrice          float64 `json:"low_price"`
	TradePrice        float64 `json:"trade_price"`
	PrevClosingPrice  float64 `json:"prev_closing_price"`
	Change            string  `json:"change"`
	ChangePrice       float64 `json:"change_price"`
	SignedChangePrice float64 `json:"signed_change_price"`
	ChangeRate        float64 `json:"change_rate"`
	SignedChangeRate  float64 `json:"signed_change_rate"`
	TradeVolume       float64 `json:"trade_volume"`
	AccTradeVolume    float64 `json:"acc_trade_volume"`
	AccTradePrice     float64 `json:"acc_trade_price"`
	TradeDate         string  `json:"trade_date"`
	TradeTime         string  `json:"trade_time"`
	TradeTimestamp    int64   `json:"trade_timestamp"`
	StreamType        string  `json:"stream_type"`
}

type TradeResponse struct {
	Type             string  `json:"type"`
	Code             string  `json:"code"`
	TradePrice       float64 `json:"trade_price"`
	TradeVolume      float64 `json:"trade_volume"`
	AskBid           string  `json:"ask_bid"`
	PrevClosingPrice float64 `json:"prev_closing_price"`
	Change           string  `json:"change"`
	ChangePrice      float64 `json:"change_price"`
	TradeDate        string  `json:"trade_date"`
	TradeTime        string  `json:"trade_time"`
	TradeTimestamp   int64   `json:"trade_timestamp"`
	StreamType       string  `json:"stream_type"`
}

type OrderbookResponse struct {
	Type           string          `json:"type"`
	Code           string          `json:"code"`
	Timestamp      int64           `json:"timestamp"`
	TotalAskSize   float64         `json:"total_ask_size"`
	TotalBidSize   float64         `json:"total_bid_size"`
	OrderbookUnits []OrderbookUnit `json:"orderbook_units"`
	StreamType     string          `json:"stream_type"`
	Level          float64         `json:"level"`
}

type MyOrderResponse struct {
	Type            string  `json:"type"`
	Code            string  `json:"code"`
	UUID            string  `json:"uuid"`
	AskBid          string  `json:"ask_bid"`
	OrderType       string  `json:"order_type"`
	Price           float64 `json:"price"`
	AvgPrice        float64 `json:"avg_price"`
	State           string  `json:"state"`
	Volume          float64 `json:"volume"`
	RemainingVolume float64 `json:"remaining_volume"`
	ExecutedVolume  float64 `json:"executed_volume"`
	TradesCount     int     `json:"trades_count"`
	Timestamp       int64   `json:"timestamp"`
	StreamType      string  `json:"stream_type"`
}

type Asset struct {
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
	Locked   float64 `json:"locked"`
}

type MyAssetResponse struct {
	Type           string  `json:"type"`
	AssetUUID      string  `json:"asset_uuid"`
	Assets         []Asset `json:"assets"`
	AssetTimestamp int64   `json:"asset_timestamp"`
	Timestamp      int64   `json:"timestamp"`
	StreamType     string  `json:"stream_type"`
}

type OrderbookUnit struct {
	AskPrice float64 `json:"ask_price"`       // 매도호가
	BidPrice float64 `json:"bid_price"`       // 매수호가
	AskSize  float64 `json:"ask_size"`        // 매도 잔량
	BidSize  float64 `json:"bid_size"`        // 매수 잔량
	Level    float64 `json:"level,omitempty"` // 호가 모아보기 단위 (default: 0, 기본 호가단위)
}
