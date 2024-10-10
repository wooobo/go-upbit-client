package private

import "time"

type NumberString string

func (n NumberString) String() string { return string(n) }

type Account struct {
	Currency            string `json:"currency"`
	Balance             string `json:"balance"`
	Locked              string `json:"locked"`
	AvgBuyPrice         string `json:"avg_buy_price"`
	AvgBuyPriceModified bool   `json:"avg_buy_price_modified"`
	UnitCurrency        string `json:"unit_currency"`
}

type PlaceOrder struct {
	UUID            string    `json:"uuid"`             // 주문의 고유 아이디
	Side            string    `json:"side"`             // 주문 종류
	OrdType         string    `json:"ord_type"`         // 주문 방식
	Price           string    `json:"price"`            // 주문 당시 화폐 가격
	State           string    `json:"state"`            // 주문 상태
	Market          string    `json:"market"`           // 마켓의 유일키
	CreatedAt       time.Time `json:"created_at"`       // 주문 생성 시간
	Volume          string    `json:"volume"`           // 사용자가 입력한 주문 양
	RemainingVolume string    `json:"remaining_volume"` // 체결 후 남은 주문 양
	ReservedFee     string    `json:"reserved_fee"`     // 수수료로 예약된 비용
	RemainingFee    string    `json:"remaining_fee"`    // 남은 수수료
	PaidFee         string    `json:"paid_fee"`         // 사용된 수수료
	Locked          string    `json:"locked"`           // 거래에 사용중인 비용
	ExecutedVolume  string    `json:"executed_volume"`  // 체결된 양
	TradesCount     int       `json:"trades_count"`     // 해당 주문에 걸린 체결 수
	TimeInForce     string    `json:"time_in_force"`    // IOC, FOK 설정
}

type Order struct {
	UUID            string       `json:"uuid"`             // 주문의 고유 아이디
	Side            string       `json:"side"`             // 주문 종류
	OrdType         string       `json:"ord_type"`         // 주문 방식
	Price           NumberString `json:"price"`            // 주문 당시 화폐 가격
	State           string       `json:"state"`            // 주문 상태
	Market          string       `json:"market"`           // 마켓 ID
	CreatedAt       time.Time    `json:"created_at"`       // 주문 생성 시간
	Volume          NumberString `json:"volume"`           // 사용자가 입력한 주문 양
	RemainingVolume NumberString `json:"remaining_volume"` // 체결 후 남은 주문 양
	ReservedFee     NumberString `json:"reserved_fee"`     // 수수료로 예약된 비용
	RemainingFee    NumberString `json:"remaining_fee"`    // 남은 수수료
	PaidFee         NumberString `json:"paid_fee"`         // 사용된 수수료
	Locked          NumberString `json:"locked"`           // 거래에 사용 중인 비용
	ExecutedVolume  NumberString `json:"executed_volume"`  // 체결된 양
	ExecutedFunds   NumberString `json:"executed_funds"`   // 현재까지 체결된 금액
	TradesCount     int          `json:"trades_count"`     // 해당 주문에 걸린 체결 수
	TimeInForce     string       `json:"time_in_force"`    // IOC, FOK 설정
}

type FilledOrder struct {
	Order
	Trades []Trade `json:"trades"` // 체결 정보 리스트
}

type Trade struct {
	Market    string    `json:"market"`
	UUID      string    `json:"uuid"`
	Price     string    `json:"price"`
	Volume    string    `json:"volume"`
	Funds     string    `json:"funds"`
	Side      string    `json:"side"`
	CreatedAt time.Time `json:"created_at"`
}
