package public

// Market represents the market information provided by Upbit
type Market struct {
	Market        string      `json:"market"`
	KoreanName    string      `json:"korean_name"`
	EnglishName   string      `json:"english_name"`
	MarketWarning string      `json:"market_warning"`
	MarketEvent   MarketEvent `json:"market_event"`
}

// MarketEvent represents the market event information
type MarketEvent struct {
	Warning bool         `json:"warning"`
	Caution CautionEvent `json:"caution"`
}

// CautionEvent represents the different types of caution events
type CautionEvent struct {
	PriceFluctuations          bool `json:"price_fluctuations"`
	TradingVolumeSoaring       bool `json:"trading_volume_soaring"`
	DepositAmountSoaring       bool `json:"deposit_amount_soaring"`
	GlobalPriceDifferences     bool `json:"global_price_differences"`
	ConcentrationSmallAccounts bool `json:"concentration_of_small_accounts"`
}

type Candle struct {
	Market               string  `json:"market"`                  // 종목 코드
	CandleDateTimeUTC    string  `json:"candle_date_time_utc"`    // 캔들 기준 시각 (UTC 기준)
	CandleDateTimeKST    string  `json:"candle_date_time_kst"`    // 캔들 기준 시각 (KST 기준)
	OpeningPrice         float64 `json:"opening_price"`           // 시가
	HighPrice            float64 `json:"high_price"`              // 고가
	LowPrice             float64 `json:"low_price"`               // 저가
	TradePrice           float64 `json:"trade_price"`             // 종가
	Timestamp            int64   `json:"timestamp"`               // 마지막 틱이 저장된 시각
	CandleAccTradePrice  float64 `json:"candle_acc_trade_price"`  // 누적 거래 금액
	CandleAccTradeVolume float64 `json:"candle_acc_trade_volume"` // 누적 거래량
	Unit                 int     `json:"unit,omitempty"`          // 분 단위 (유닛), optional
}

// TradeTick 최근 체결 내역
type TradeTick struct {
	Market           string  `json:"market"`
	TradeDateUTC     string  `json:"trade_date_utc"`
	TradeTimeUTC     string  `json:"trade_time_utc"`
	Timestamp        int64   `json:"timestamp"`
	TradePrice       float64 `json:"trade_price"`
	TradeVolume      float64 `json:"trade_volume"`
	PrevClosingPrice float64 `json:"prev_closing_price"`
	ChangePrice      float64 `json:"change_price"`
	AskBid           string  `json:"ask_bid"`
	SequentialID     int64   `json:"sequential_id"`
}

// TickerSnapshot 구조체는 거래소의 종목 정보를 나타냅니다.
type TickerSnapshot struct {
	Market             string  `json:"market"`                // 종목 구분 코드
	TradeDate          string  `json:"trade_date"`            // 최근 거래 일자(UTC), 형식: yyyyMMdd
	TradeTime          string  `json:"trade_time"`            // 최근 거래 시각(UTC), 형식: HHmmss
	TradeDateKst       string  `json:"trade_date_kst"`        // 최근 거래 일자(KST), 형식: yyyyMMdd
	TradeTimeKst       string  `json:"trade_time_kst"`        // 최근 거래 시각(KST), 형식: HHmmss
	TradeTimestamp     int64   `json:"trade_timestamp"`       // 최근 거래 일시(UTC), Unix Timestamp
	OpeningPrice       float64 `json:"opening_price"`         // 시가
	HighPrice          float64 `json:"high_price"`            // 고가
	LowPrice           float64 `json:"low_price"`             // 저가
	TradePrice         float64 `json:"trade_price"`           // 종가(현재가)
	PrevClosingPrice   float64 `json:"prev_closing_price"`    // 전일 종가(UTC 0시 기준)
	Change             string  `json:"change"`                // 변화 상태 (EVEN: 보합, RISE: 상승, FALL: 하락)
	ChangePrice        float64 `json:"change_price"`          // 변화액의 절대값
	ChangeRate         float64 `json:"change_rate"`           // 변화율의 절대값
	SignedChangePrice  float64 `json:"signed_change_price"`   // 부호가 있는 변화액
	SignedChangeRate   float64 `json:"signed_change_rate"`    // 부호가 있는 변화율
	TradeVolume        float64 `json:"trade_volume"`          // 가장 최근 거래량
	AccTradePrice      float64 `json:"acc_trade_price"`       // 누적 거래대금(UTC 0시 기준)
	AccTradePrice24h   float64 `json:"acc_trade_price_24h"`   // 24시간 누적 거래대금
	AccTradeVolume     float64 `json:"acc_trade_volume"`      // 누적 거래량(UTC 0시 기준)
	AccTradeVolume24h  float64 `json:"acc_trade_volume_24h"`  // 24시간 누적 거래량
	Highest52WeekPrice float64 `json:"highest_52_week_price"` // 52주 신고가
	Highest52WeekDate  string  `json:"highest_52_week_date"`  // 52주 신고가 달성일, 형식: yyyy-MM-dd
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`  // 52주 신저가
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`   // 52주 신저가 달성일, 형식: yyyy-MM-dd
	Timestamp          int64   `json:"timestamp"`             // 타임스탬프
}

// OrderBook 호가 정보
type OrderBook struct {
	Market         string          `json:"market"`          // 종목 코드
	Timestamp      int64           `json:"timestamp"`       // 호가 생성 시각
	TotalAskSize   float64         `json:"total_ask_size"`  // 호가 매도 총 잔량
	TotalBidSize   float64         `json:"total_bid_size"`  // 호가 매수 총 잔량
	OrderbookUnits []OrderbookUnit `json:"orderbook_units"` // 호가 리스트
}

type OrderbookUnit struct {
	AskPrice float64 `json:"ask_price"`       // 매도호가
	BidPrice float64 `json:"bid_price"`       // 매수호가
	AskSize  float64 `json:"ask_size"`        // 매도 잔량
	BidSize  float64 `json:"bid_size"`        // 매수 잔량
	Level    float64 `json:"level,omitempty"` // 호가 모아보기 단위 (default: 0, 기본 호가단위)
}

// SupportedLevels 호가 모아보기 단위 정보 조회
type SupportedLevels struct {
	Market          string    `json:"market"`           // 종목 코드
	SupportedLevels []float64 `json:"supported_levels"` // 해당 종목에서 지원하는 모아보기 단위, 예: 0: 기본 호가단위
}
