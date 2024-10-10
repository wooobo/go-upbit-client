package public

type CandleInterval string

const (
	Minute CandleInterval = "minute"
	Day    CandleInterval = "day"
	Week   CandleInterval = "week"
	Month  CandleInterval = "month"
)

type CandleRequest struct {
	Market string `json:"market"`       // 마켓 코드 (ex. KRW-BTC)
	To     string `json:"to,omitempty"` // 마지막 캔들 시각 (ISO8061 포맷)
	Count  int    `json:"count"`        // 요청할 캔들 개수 (최대 200개)
	CandleInterval
	UnitCount int `json:"unit,omitempty"` // 분 단위 (유닛), optional
}

// TradeTicksRequest 최근 체결 내역
type TradeTicksRequest struct {
	Market  string `json:"market"`
	To      string `json:"to,omitempty"`
	Count   int    `json:"count,omitempty"`
	Cursor  string `json:"cursor,omitempty"`
	DaysAgo int    `json:"daysAgo,omitempty"`
}
