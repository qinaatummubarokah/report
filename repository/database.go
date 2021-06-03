package repository

import (
	"context"
)

// type Data struct {
// 	ID          string `json:"id"`
// 	Name        string `json:"name"`
// 	AccountCode string `json:"account_code"`
// }

type Data struct {
	ID                 *string `json:"id"`                                         //
	Fare               *string `json:"fare"`                                       //
	Extra              *string `json:"extra"`                                      //
	DiscountAmt        *string `json:"discount_amt" db:"discount_amt"`             //
	PaidAmount         *string `json:"paid_amount" db:"paid_amount"`               //
	PromoCode          *string `json:"promo_code" db:"promo_code"`                 //
	PaymentToken       *string `json:"payment_token" db:"payment_token"`           //
	TransactionTime    *string `json:"transaction_time" db:"transaction_time"`     //
	Identifier         *string `json:"identifier" db:"identifier"`                 //
	PaymentType        *string `json:"payment_type" db:"payment_type"`             //
	VehicleId          *string `json:"vehicle_id" db:"vehicle_id"`                 //
	VehicleName        *string `json:"vehicle_name" db:"vehicle_name"`             //
	ServiceType        *string `json:"service_type" db:"service_type"`             //
	DriverID           *string `json:"driver_id" db:"driver_id"`                   //
	PickUpSuburb       *string `json:"pick_up_suburb" db:"pick_up_suburb"`         //
	PickUpArea         *string `json:"pick_up_area" db:"pick_up_area"`             //
	DestinationArea    *string `json:"destination_area" db:"destination_area"`     //
	DSestinationSuburb *string `json:"destination_suburb" db:"destination_suburb"` //
	// PickUpLatitude        string `json:"pick_up_latitude" db:"pick_up_latitude"`
	PickUpLng          *string `json:"pick_up_lng" db:"pick_up_lng"`               //
	PaymentProfileID   *string `json:"payment_profile_id" db:"payment_profile_id"` //
	State              *string `json:"state"`                                      //
	ReleasedAt         *string `json:"released_at" db:"released_at"`               //
	CompletedAt        *string `json:"completed_at" db:"completed_at"`             //
	CreatedAt          *string `json:"created_at" db:"created_at"`                 //
	Updated_at         *string `json:"updated_at" db:"updated_at"`                 //
	CcIdentifier       *string `json:"cc_identifier" db:"cc_identifier"`           //
	AccountID          *string `json:"account_id" db:"account_id"`                 //
	SapSentAt          *string `json:"sap_sent_at" db:"sap_sent_at"`               //
	SapState           *string `json:"sap_state" db:"sap_state"`                   //
	MsakuState         *string `json:"msaku_state" db:"msaku_state"`               //
	CvNumber           *string `json:"cv_number" db:"cv_number"`                   //
	ValidityPeriod     *string `json:"validity_period" db:"validity_period"`       //
	ItopID             *string `json:"itop_id" db:"itop_id"`                       //
	OrderID            *string `json:"order_id" db:"order_id"`                     //
	PickupAddress      *string `json:"pickup_address" db:"pickup_address"`
	PickedUp           *string `json:"picked_up_at" db:"picked_up_at"` //
	TripPurpose        *string `json:"trip_purpose" db:"trip_purpose"` //
	MsakuTransactionID *string `json:"msaku_transaction_id,,omitempty" db:"msaku_transaction_id"`
	// TripPurposedriverName string `json:"trip_purposedriver_name" db:"trip_purposedriver_name"`
	ExternalOrderID  *string `json:"external_order_id" db:"external_order_id"`
	RouteImage       *string `json:"route_image" db:"route_image"`
	DepartmentName   *string `json:"department_name" db:"department_name"`
	AccountCode      *string `json:"account_code" db:"account_code"`
	UserName         *string `json:"user_name" db:"user_name"`
	InvoiceNumber    *string `json:"invoice_number,omitempty" db:"invoice_number"`
	PostingDate      *string `json:"posting_date" db:"posting_date"`
	Distance         *string `json:"distance"`
	OtherInformation *string `json:"other_information" db:"other_information"`
	PickUpLat        *string `json:"pick_up_lat" db:"pick_up_lat"`         //
	DestinationLat   *string `json:"destination_lat" db:"destination_lat"` //
	DestinationLng   *string `json:"destination_lng" db:"destination_lng"` //
	MsakuResponse    *string `json:"msaku_response" db:"msaku_response"`   //
	DropoffAddress   *string `json:"dropoff_address" db:"dropoff_address"` //
	Tips             *string `json:"tips" db:"tips"`
	DriverName       *string `json:"driver_name" db:"driver_name"`
}

type Repository interface {
	GetReport(ctx context.Context) ([]Data, error)
}
