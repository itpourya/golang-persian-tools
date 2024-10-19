package bill

import (
	"fmt"
	"strconv"
)

// BillTypes represents the type of bills.
type BillTypes string

const (
	Water             BillTypes = "آب"
	Electricity       BillTypes = "برق"
	Gas               BillTypes = "گاز"
	FixedTelephone    BillTypes = "تلفن ثابت"
	MobilePhone       BillTypes = "تلفن همراه"
	MunicipalityDues  BillTypes = "عوارض شهرداری"
	TaxOrganization   BillTypes = "سازمان مالیات"
	TrafficViolations BillTypes = "جرایم راهنمایی و رانندگی"
	Unknown           BillTypes = "unknown"
)

// BillTypesModel maps bill type indices to their corresponding types.
var billTypes = map[int]BillTypes{
	1: Water,
	2: Electricity,
	3: Gas,
	4: FixedTelephone,
	5: MobilePhone,
	6: MunicipalityDues,
	8: TaxOrganization,
	9: TrafficViolations,
}

// Currency type for currency representation.
type Currency string

const (
	Toman Currency = "toman"
	Rial  Currency = "rial"
)

// BillBarcodeModel represents the structure of a bill barcode.
type BillBarcodeModel struct {
	BillID    int
	PaymentID int
}

// BillResult contains the result of the bill verification.
type BillResult struct {
	Amount             int
	Type               BillTypes
	Barcode            string
	IsValid            bool
	IsValidBillID      bool
	IsValidBillPayment bool
}

// BillParams contains parameters for creating a Bill instance.
type BillParams struct {
	BillID    *int
	PaymentID *int
	Currency  Currency
	Barcode   *string
}

// Bill represents the bill with various properties and methods for processing.
type Bill struct {
	barcode     *string
	currency    Currency
	billID      *int
	billPayment *int
}

// NewBill creates a new Bill instance.
func NewBill(params BillParams) *Bill {
	bill := &Bill{
		barcode:  params.Barcode,
		currency: params.Currency,
	}

	if params.BillID != nil && params.PaymentID != nil {
		bill.billID = params.BillID
		bill.billPayment = params.PaymentID
	}

	return bill
}

// GetAmount calculates the bill amount.
func (b *Bill) GetAmount() int {
	currencyFactor := 100
	if b.currency == Rial {
		currencyFactor = 1000
	}

	amountStr := strconv.Itoa(*b.billPayment)
	amountStr = amountStr[:len(amountStr)-5] // Remove last 5 digits
	amount, _ := strconv.Atoi(amountStr)
	return amount * currencyFactor
}

// GetBillType retrieves the bill type based on the bill ID.
func (b *Bill) GetBillType() BillTypes {
	if b.billID == nil {
		return Unknown
	}
	index := (*b.billID % 100) / 10
	return billTypes[index]
}

// GetBarcode generates the barcode for the bill.
func (b *Bill) GetBarcode() string {
	return fmt.Sprintf("%d000%d", *b.billID, *b.billPayment)
}

// FindByBarcode extracts the bill ID and payment ID from the barcode.
func (b *Bill) FindByBarcode(barcode *string) BillBarcodeModel {
	barcodeStr := *barcode
	billID, _ := strconv.Atoi(barcodeStr[:13])
	paymentID, _ := strconv.Atoi(barcodeStr[16:26])
	return BillBarcodeModel{BillID: billID, PaymentID: paymentID}
}

// VerifyBillPayment validates the payment ID.
func (b *Bill) VerifyBillPayment() bool {
	if b.billPayment == nil {
		return false
	}
	paymentIDStr := fmt.Sprintf("%d", *b.billPayment)

	if len(paymentIDStr) < 6 {
		return false
	}

	firstControlBit := paymentIDStr[len(paymentIDStr)-2 : len(paymentIDStr)-1]
	secondControlBit := paymentIDStr[len(paymentIDStr)-1:]

	paymentIDStr = paymentIDStr[:len(paymentIDStr)-2]
	return b.calculateControlBit(paymentIDStr) == firstControlBit[0] &&
		b.calculateControlBit(fmt.Sprintf("%d%s", *b.billID, paymentIDStr+string(firstControlBit[0]))) == secondControlBit[0]
}

// VerifyBillID validates the bill ID.
func (b *Bill) VerifyBillID() bool {
	if b.billID == nil {
		return false
	}
	billIDStr := fmt.Sprintf("%d", *b.billID)

	if len(billIDStr) < 6 {
		return false
	}

	controlBit := billIDStr[len(billIDStr)-1:]
	billIDStr = billIDStr[:len(billIDStr)-1]

	return b.calculateControlBit(billIDStr) == controlBit[0] && b.GetBillType() != Unknown
}

// CalculateControlBit calculates the control bit for a given number string.
func (b *Bill) calculateControlBit(num string) byte {
	sum := 0
	base := 2

	for i := 0; i < len(num); i++ {
		if base == 8 {
			base = 2
		}
		subString := string(num[len(num)-1-i])
		n, _ := strconv.Atoi(subString)
		sum += n * base
		base++
	}
	sum %= 11
	if sum < 2 {
		sum = 0
	} else {
		sum = 11 - sum
	}
	return byte(sum)
}

// VerifyBill validates both payment and bill ID.
func (b *Bill) VerifyBill() bool {
	return b.VerifyBillPayment() && b.VerifyBillID()
}

// GetResult returns the result of the bill verification.
func (b *Bill) GetResult() BillResult {
	return BillResult{
		Amount:             b.GetAmount(),
		Type:               b.GetBillType(),
		Barcode:            b.GetBarcode(),
		IsValid:            b.VerifyBill(),
		IsValidBillID:      b.VerifyBillID(),
		IsValidBillPayment: b.VerifyBillPayment(),
	}
}
