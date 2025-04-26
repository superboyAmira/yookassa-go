package yookassa

type YooKassaClient struct {
	key    string
	shopId string

	redirectURL string 
}

func Load(key string, shop_id string) *YooKassaClient {
	return &YooKassaClient{key: key, shopId: shop_id}
}

func 

// func (s *YooKassaClient) parseAPI() {
// 	path := os.Getenv("SECRET_FILE")

// 	file, err := os.Open(path)
// 	if err != nil {
// 		return
// 	}
// 	reader := bufio.NewScanner(file)
// 	if reader.Scan() {
// 		s.key = strings.Split(reader.Text(), "=")[1]
// 	}
// 	if reader.Scan() {
// 		s.shop_id = strings.Split(reader.Text(), "=")[1]
// 	}

// 	if s.key == "" || s.shop_id == "" || == "" {
// 		return
// 	}
// }

// func generateIdempotenceKey() string {
// 	idempotenceKey := uuid.New().String()
// 	if len(idempotenceKey) > 64 {
// 		return idempotenceKey[:64]
// 	}
// 	return idempotenceKey
// }

// func (s *PaymentService) CreateYooKassaPayment(ctx context.Context, selector *schemas.ClientCommand, amount float32) (*PaymentResponse, error) {
// 	metadata := &UserMetadata{
// 		TgID:       selector.User.TelegramChatID,
// 		TgUsername: selector.User.TelegramUserName,
// 	}
// 	url := "https://api.yookassa.ru/v3/payments"

// 	var reqPay *PaymentRequest
// 	if selector.Payment.SelectedPaymentMethod == "bank_card" {
// 		reqPay = getBankCardPaymentRequest(fmt.Sprintf("%.1f", amount), selector.User.Email, metadata)
// 	} else {
// 		reqPay = getSBPPaymentRequest(fmt.Sprintf("%.1f", amount), selector.User.Email, metadata)
// 	}

// 	requestBody, err := json.Marshal(reqPay)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.SetBasicAuth(s.shop_id, s.key)
// 	req.Header.Set("Idempotence-Key", generateIdempotenceKey())
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var paymentResponse PaymentResponse
// 	err = json.Unmarshal(body, &paymentResponse)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if resp.StatusCode != 200 {
// 		return nil, fmt.Errorf("400 error, something wrong")
// 	}

// 	logger.GlobalLogger.Info("Created payment", zap.Any("payment", paymentResponse))
// 	return &paymentResponse, nil
// }

// func (s *PaymentService) GetStatusPayment(ctx context.Context, paymentUUID string) (*Payment, error) {
// 	logger.GlobalLogger.Debug("Start payment validation", zap.String("paymentUUID", paymentUUID))
// 	defer logger.GlobalLogger.Debug("Stop payment validation", zap.String("paymentUUID", paymentUUID))

// 	var paymentResponse Payment
// 	ticker := time.NewTicker(time.Second * 10)
// 	url := "https://api.yookassa.ru/v3/payments/" + paymentUUID

// 	for {
// 		select {
// 		case <-ticker.C:
// 			req, err := http.NewRequest("GET", url, nil)
// 			if err != nil {
// 				return nil, err
// 			}

// 			req.SetBasicAuth(s.shop_id, s.key)
// 			req.Header.Set("Content-Type", "application/json")

// 			client := &http.Client{}
// 			resp, err := client.Do(req)
// 			if err != nil {
// 				return nil, err
// 			}
// 			defer resp.Body.Close()

// 			if resp.StatusCode != http.StatusOK {
// 				return nil, errors.New("failed to fetch payment status: " + resp.Status)
// 			}

// 			err = json.NewDecoder(resp.Body).Decode(&paymentResponse)
// 			if err != nil {
// 				return nil, err
// 			}

// 			if paymentResponse.Status == "canceled" {
// 				return nil, fmt.Errorf("платеж отменен")
// 			} else if (paymentResponse.Status == "waiting_for_capture" || paymentResponse.Status == "succeeded") && paymentResponse.Paid {
// 				return &paymentResponse, nil
// 			}
// 		case <-ctx.Done():
// 			// err := s.CancelPayment(paymentUUID)
// 			// if err != nil {
// 			// 	logger.GlobalLogger.ErrorCritical("Cannot cancel payment, admins manual delete please! "+ err.Error(), zap.String("pay_uuid", paymentUUID), zap.Any("paymentRepose", paymentResponse))
// 			// }
// 			// Отмена платежа производится только при оплате уже, чеки просто будут висеть  в юкассе
// 			logger.GlobalLogger.Info("Payment Canceled", zap.String("pay_uuid", paymentUUID))
// 			return nil, context.DeadlineExceeded
// 		}
// 	}
// }

// // не используется
// func (s *PaymentService) CancelPayment(paymentUUID string) error {
// 	url := "https://api.yookassa.ru/v3/payments/" + paymentUUID + "/cancel"

// 	body := []byte("{}")

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
// 	if err != nil {
// 		return fmt.Errorf("failed to create request: %w", err)
// 	}

// 	req.SetBasicAuth(s.shop_id, s.key)
// 	req.Header.Set("Idempotence-Key", generateIdempotenceKey())
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("failed to send request: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
// 		return fmt.Errorf("failed to cancel payment, status: %s", resp.Status)
// 	}
// 	var succ Payment
// 	err = json.NewDecoder(resp.Body).Decode(&succ)
// 	if err != nil {
// 		return err
// 	}

// 	if succ.Status != "canceled" {
// 		logger.GlobalLogger.ErrorCritical("Bad resp cancel", zap.Any("payment", resp))
// 		return fmt.Errorf("something err with Kassa")
// 	}
// 	logger.GlobalLogger.Info("Canceled payment", zap.Any("payment", resp))

// 	return nil
// }

// func (s *PaymentService) CapturePayment(paymentUUID string, amount *Amount) error {
// 	if amount == nil {
// 		return fmt.Errorf("amount not valid")
// 	}
// 	url := "https://api.yookassa.ru/v3/payments/" + paymentUUID + "/capture"

// 	captureRequest := PaymentRequestCapture{Amount: *amount}
// 	body, err := json.Marshal(captureRequest)
// 	if err != nil {
// 		return fmt.Errorf("failed to marshal request body: %w", err)
// 	}

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
// 	if err != nil {
// 		return fmt.Errorf("failed to create request: %w", err)
// 	}

// 	req.SetBasicAuth(s.shop_id, s.key)
// 	req.Header.Set("Idempotence-Key", generateIdempotenceKey())
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("failed to send request: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
// 		logger.GlobalLogger.ErrorCritical("Bad resp capture", zap.Any("payment", resp))
// 		return fmt.Errorf("failed to capture payment, status: %s", resp.Status)
// 	}
// 	var succ Payment
// 	err = json.NewDecoder(resp.Body).Decode(&succ)
// 	if err != nil {
// 		return err
// 	}

// 	if !succ.Paid || succ.Status != "succeeded" {
// 		logger.GlobalLogger.ErrorCritical("Failed payment capture in body req", zap.Any("payment", resp))
// 		return fmt.Errorf("something err with Kassa")
// 	}
// 	logger.GlobalLogger.Info("Captured payment", zap.Any("payment", succ))

// 	return nil
// }

// func getBankCardPaymentRequest(amount string, email string, metadata *UserMetadata) *PaymentRequest {
// 	pay := &PaymentRequest{
// 		Amount: Amount{
// 			Value:    amount,
// 			Currency: `RUB`,
// 		},
// 		PaymentMethodData: struct {
// 			Type string "json:\"type\""
// 		}{
// 			Type: "bank_card",
// 		},
// 		Confirmation: struct {
// 			Type      string "json:\"type\""
// 			ReturnURL string "json:\"return_url\""
// 		}{
// 			Type:      "redirect",
// 			ReturnURL: "https://t.me/NorkinVPN_bot",
// 		},
// 		Description: "Payment: " + time.Now().String(),
// 		Receipt: Receipt{
// 			Customer: Customer{
// 				Email: email,
// 			},
// 			Items: []Item{{
// 				Description: "Подписка (" + amount + " рублей за 1 месяц)",
// 				Quantity:    1,
// 				Amount: Amount{
// 					Value:    amount,
// 					Currency: `RUB`,
// 				},
// 				VATCode:     1,
// 				PaymentSubj: "service",
// 				PaymentMode: "full_payment",
// 			}},
// 		},
// 		Capture:  false,
// 		Metadata: *metadata,
// 	}
// 	return pay
// }

// func getSBPPaymentRequest(amount string, email string, metadata *UserMetadata) *PaymentRequest {
// 	pay := &PaymentRequest{
// 		Amount: Amount{
// 			Value:    amount,
// 			Currency: `RUB`,
// 		},
// 		PaymentMethodData: struct {
// 			Type string "json:\"type\""
// 		}{
// 			Type: "sbp",
// 		},
// 		Confirmation: struct {
// 			Type      string "json:\"type\""
// 			ReturnURL string "json:\"return_url\""
// 		}{
// 			Type:      "redirect",
// 			ReturnURL: "https://t.me/NorkinVPN_bot",
// 		},
// 		Description: "Payment: " + time.Now().String(),
// 		Receipt: Receipt{
// 			Customer: Customer{
// 				Email: email,
// 			},
// 			Items: []Item{{
// 				Description: "Подписка (" + amount + " рублей за 1 месяц)",
// 				Quantity:    1,
// 				Amount: Amount{
// 					Value:    amount,
// 					Currency: `RUB`,
// 				},
// 				VATCode:     1,
// 				PaymentSubj: "service",
// 				PaymentMode: "full_payment",
// 			}},
// 		},
// 		Capture:  true,
// 		Metadata: *metadata,
// 	}
// 	return pay
// }
