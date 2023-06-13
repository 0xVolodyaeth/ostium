package models

type Bet struct {
	ID           int64  `bun:"id,pk" json:"id"`
	LongAddress  string `bun:"long_address,notnull" json:"monthlyTurnover"`
	ShortAddress string `bun:"short_address,notnull" json:"shortAddress"`
	Amount       int64  `bun:"amount,notnull" json:"longAddress"`
	Expiration   int64  `bun:"expiration,notnull" json:"expiration"`
	CreatedAt    int64  `bun:"created_at,notnull" json:"createdAt"`
	OpeningPrice int64  `bun:"opening_price,notnull" json:"openingPrice"`
	IsActive     bool   `bun:"is_active,notnull" json:"isActive"`
	Withdrawn    bool   `bun:"withdrawn,notnull" json:"withdrawn"`
	Winner       string `bun:"winner,notnull" json:"winner"`
	Canceled     bool   `bun:"canceled,notnull" json:"canceled"`
}
