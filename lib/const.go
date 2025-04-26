package yookassa

const (
	createPaymentUrl  = "https://api.yookassa.ru/v3/payments"
	getPaymentUrl     = "https://api.yookassa.ru/v3/payments/%s"         // + paymentUUID
	cancelPaymentUrl  = "https://api.yookassa.ru/v3/payments/%s/cancel"  // + paymentUUID
	capturePaymentUrl = "https://api.yookassa.ru/v3/payments/%s/capture" // + paymentUUID
)
