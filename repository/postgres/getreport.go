package postgres

import (
	"context"
	"fmt"
	"log"
	"report/repository"
)

func (p *postgres) GetReport(ctx context.Context) (result []repository.Data, err error) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	// mutex.RLock()
	// sql := "select id, name from accounts where id in (62,26)"

	sql := `
	SELECT distinct
		payment_transactions.*,
		departments.name AS department_name,
		accounts.account_code AS account_code,
		users.full_name AS user_name,
		info.invoice_number AS invoice_number,
		info.posting_date AS posting_date,
		info.distance AS distance,
		users.other_information AS other_information
		FROM "payment_transactions"
		LEFT JOIN voucher_profiles vp ON vp.id = payment_transactions.cc_identifier::integer
		LEFT JOIN users ON users.id = vp.user_id
		LEFT JOIN departments ON departments.id = vp.department_id
		LEFT JOIN accounts ON accounts.id = payment_transactions.account_id
		LEFT JOIN payment_transaction_infos info ON info.payment_transaction_id = payment_transactions.id
	WHERE
		"payment_transactions"."state" IN (2, 3) AND
		"payment_transactions"."payment_type" IN ('ecv', 'edc') AND
		("payment_transactions"."completed_at" BETWEEN '2017-07-01 00:00:00.000000' AND '2021-05-28 00:00:00.000000')
	ORDER BY "payment_transactions"."picked_up_at" ASC, "payment_transactions"."completed_at" ASC`

	// rows, err := p.db.Query(sql)
	row, err := p.db.Query(sql)
	// mutex.RUnlock()
	log.Println("row: ", row)
	log.Println("connect db err: ", err)
	if err != nil {
		return nil, err
	}
	log.Println("AAAAA")
	defer func() {
		if e := row.Close(); e != nil {
			err = e
		}
	}()
	log.Println("BBBBB")
	// var data []repository.Transaction
	for row.Next() {
		each := repository.Data{}
		// err = row.Scan(
		// 	&each.ID,
		// 	&each.Name,
		// 	&each.AccountCode)
		err = row.Scan(
			&each.ID,
			&each.Fare,
			&each.Extra,
			&each.DiscountAmt,
			&each.PaidAmount,
			&each.PromoCode,
			&each.PaymentToken,
			&each.TransactionTime,
			&each.Identifier,
			&each.PaymentType,
			&each.VehicleId,
			&each.VehicleName,
			&each.ServiceType,
			&each.DriverID,
			&each.PickUpSuburb,
			&each.PickUpArea,
			&each.DestinationArea,
			&each.DSestinationSuburb,
			// &each.PickUpLatitude,
			&each.PickUpLat,
			&each.PickUpLng,
			&each.PaymentProfileID,
			&each.State,
			&each.ReleasedAt,
			&each.CompletedAt,
			&each.CreatedAt,
			&each.Updated_at,
			&each.CcIdentifier,
			&each.AccountID,
			&each.SapSentAt,
			&each.SapState,
			&each.MsakuState,
			&each.CvNumber,
			&each.ValidityPeriod,
			&each.ItopID,
			&each.OrderID,
			&each.PickedUp,
			&each.TripPurpose,
			&each.MsakuTransactionID,
			&each.ExternalOrderID,
			&each.RouteImage,
			&each.DepartmentName,
			&each.AccountCode,
			&each.UserName,
			&each.InvoiceNumber,
			&each.PostingDate,
			&each.Distance,
			&each.OtherInformation,
			&each.DestinationLat,
			&each.DestinationLng,
			&each.MsakuResponse,
			&each.PickupAddress,
			&each.DropoffAddress,
			&each.Tips,
			&each.DriverName)
		log.Println("err scan: ", err)
		if err != nil {
			return nil, fmt.Errorf("row scan: %v", err)
		}
		result = append(result, each)
	}
	// log.Println("data", result)
	return result, nil
}
