package private

type State string

const (
	StateWait  State = "wait"
	StateWatch State = "watch"
)

func (s State) String() string {
	return string(s)
}

type OrderBy string

func (o OrderBy) String() string {
	return string(o)
}

const (
	OrderByAsc  OrderBy = "asc"
	OrderByDesc OrderBy = "desc"
)

type OrderSide string

func (o OrderSide) String() string { return string(o) }

const (
	OrderSideBid OrderSide = "bid" // 매수
	OrderSideAsk OrderSide = "ask" // 매도
)

type PlaceOrderRequest struct {
	Market      string    `json:"market"`                  // 마켓 ID
	Side        OrderSide `json:"side"`                    // 주문 종류
	Volume      string    `json:"volume"`                  // 주문량
	Price       string    `json:"price"`                   // 주문 가격
	OrdType     string    `json:"ord_type"`                // 주문 타입
	Identifier  string    `json:"identifier,omitempty"`    // 조회용 사용자 지정값
	TimeInForce string    `json:"time_in_force,omitempty"` // IOC, FOK 주문 설정
}

type CancelOrderRequest struct {
	UUID       string `json:"uuid"`       // 주문 UUID
	Identifier string `json:"identifier"` // 조회용 사용자 지정 값
}

type AccountStatus struct {
	Currency            string       `json:"currency"`
	Balance             string       `json:"balance"`
	Locked              string       `json:"locked"`
	AvgBuyPrice         NumberString `json:"avg_buy_price"`
	AvgBuyPriceModified bool         `json:"avg_buy_price_modified"`
	UnitCurrency        string       `json:"unit_currency"`
}

type Constraint struct {
	Currency  string `json:"currency"`
	PriceUnit string `json:"price_unit"`
	MinTotal  string `json:"min_total"`
}

type MarketTicker struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	OrderTypes []string   `json:"order_types,omitempty"` // 만료된 필드이므로 사용하지 않음
	OrderSides []string   `json:"order_sides"`
	Bid        Constraint `json:"bid"`
	Ask        Constraint `json:"ask"`
	MaxTotal   string     `json:"max_total"`
	State      string     `json:"state"`
}

type OrderChance struct {
	BidFee     string        `json:"bid_fee"`
	AskFee     string        `json:"ask_fee"`
	Market     MarketTicker  `json:"market"`
	AskTypes   []string      `json:"ask_types"`
	BidTypes   []string      `json:"bid_types"`
	BidAccount AccountStatus `json:"bid_account"`
	AskAccount AccountStatus `json:"ask_account"`
}

type OrderSearchRequest struct {
	Market      string   `json:"market"`
	UUIDs       []string `json:"uuids,omitempty"`
	Identifiers []string `json:"identifiers,omitempty"`
	OrderBy     string   `json:"order_by,omitempty"`
}

type OrderQueryParams struct {
	Market  string  `json:"market"`
	State   State   `json:"state,omitempty"`
	States  []State `json:"states,omitempty"`
	Page    int     `json:"page,omitempty"`
	Limit   int     `json:"limit,omitempty"`
	OrderBy OrderBy `json:"order_by,omitempty"`
}

type CompletedOrderState string

const (
	StateCompletedOrderDone   CompletedOrderState = "done"
	StateCompletedOrderCancel CompletedOrderState = "cancel"
)

func (s CompletedOrderState) String() string {
	return string(s)
}

type CompletedOrderRequest struct {
	Market    string                `json:"market"`               // 마켓 ID
	State     CompletedOrderState   `json:"state,omitempty"`      // 주문 상태
	States    []CompletedOrderState `json:"states,omitempty"`     // 주문 상태 목록, 기본값: ['done', 'cancel'], state 와 states 는 동시에 사용할 수 없습니다.
	StartTime string                `json:"start_time,omitempty"` // 조회 시작 시간 (ISO-8601 포맷)
	EndTime   string                `json:"end_time,omitempty"`   // 조회 종료 시간 (ISO-8601 포맷)
	Limit     int                   `json:"limit,omitempty"`      // 요청 개수, 기본값: 100, 최대값: 1000
	OrderBy   OrderBy               `json:"order_by,omitempty"`   // 정렬 방식, 기본값: 내림차순,  asc : 오름차순, desc : 내림차순 (default)
}
