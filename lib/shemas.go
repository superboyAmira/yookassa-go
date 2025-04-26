package yookassa

import "time"

type Payment struct {
	ID                   string                 `json:"id"`
	Status               string                 `json:"status"`
	Paid                 bool                   `json:"paid"`
	Amount               Amount                 `json:"amount"`
	AuthorizationDetails AuthorizationDetails   `json:"authorization_details"`
	CreatedAt            time.Time              `json:"created_at"`
	Description          string                 `json:"description"`
	ExpiresAt            time.Time              `json:"expires_at"`
	Metadata             map[string]interface{} `json:"metadata"`
	PaymentMethod        PaymentMethod          `json:"payment_method"`
	Recipient            Recipient              `json:"recipient"`
	Refundable           bool                   `json:"refundable"`
	Test                 bool                   `json:"test"`
	IncomeAmount         Amount                 `json:"income_amount"`
}

type PaymentRequest struct {
	Amount            Amount `json:"amount"`
	PaymentMethodData struct {
		Type string `json:"type"`
	} `json:"payment_method_data"`
	Confirmation struct {
		Type      string `json:"type"`
		ReturnURL string `json:"return_url"`
	} `json:"confirmation"`
	Description string       `json:"description"`
	Receipt     Receipt      `json:"receipt"`
	Capture     bool         `json:"capture"`
	Metadata    UserMetadata `json:"metadata"`
}

type PaymentRequestCapture struct {
	Amount Amount `json:"amount"`
}

type PaymentResponse struct {
	ID            string                 `json:"id"`
	Status        string                 `json:"status"`
	Paid          bool                   `json:"paid"`
	Amount        Amount                 `json:"amount"`
	Confirmation  Confirmation           `json:"confirmation"`
	CreatedAt     string                 `json:"created_at"`
	Description   string                 `json:"description"`
	Metadata      map[string]interface{} `json:"metadata"`
	PaymentMethod PaymentMethod          `json:"payment_method"`
	Recipient     Recipient              `json:"recipient"`
	Refundable    bool                   `json:"refundable"`
	Test          bool                   `json:"test"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type AuthorizationDetails struct {
	RRN          string       `json:"rrn"`
	AuthCode     string       `json:"auth_code"`
	ThreeDSecure ThreeDSecure `json:"three_d_secure"`
}

type ThreeDSecure struct {
	Applied bool `json:"applied"`
}

type PaymentMethod struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Saved bool   `json:"saved"`
	Card  Card   `json:"card"`
	Title string `json:"title"`
}

type Card struct {
	First6        string      `json:"first6"`
	Last4         string      `json:"last4"`
	ExpiryMonth   string      `json:"expiry_month"`
	ExpiryYear    string      `json:"expiry_year"`
	CardType      string      `json:"card_type"`
	CardProduct   CardProduct `json:"card_product"`
	IssuerCountry string      `json:"issuer_country"`
	IssuerName    string      `json:"issuer_name"`
}

type CardProduct struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Recipient struct {
	AccountID string `json:"account_id"`
	GatewayID string `json:"gateway_id"`
}

type Confirmation struct {
	Type      string `json:"type"`
	ReturnURL string `json:"return_url"`

	ConfirmationURL string `json:"confirmation_url"`
}

type Receipt struct {
	Customer Customer `json:"customer"`
	Items    []Item   `json:"items"`
}

type Customer struct {
	Email string `json:"email"`
}

// https://yookassa.ru/developers/api#create_payment
type Item struct {
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	Amount      Amount  `json:"amount"`
	VATCode     int     `json:"vat_code"`
	PaymentSubj string  `json:"payment_subject"`
	PaymentMode string  `json:"payment_mode"`
}

type UserMetadata struct {
	TgID       int64  `json:"tg_id"`
	TgUsername string `json:"tg_username"`
}
