package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"report/repository"

	_ "github.com/lib/pq"
)

type ConnParam struct {
	Host        string
	Port        string
	DBName      string
	User        string
	Pass        string
	Options     string
	MaxOpenConn int
	MaxIdleConn int
	MaxLifetime time.Duration
}

type postgres struct {
	db *sql.DB
	mu sync.RWMutex
}

func NewPostgresSql(p *ConnParam) (repository.Repository, error) {
	// psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	p.Host, p.Port, p.User, p.Pass, p.DBName)

	psqlconn := "postgres://postgres:@127.0.0.1:5432/postgres?sslmode=disable&search_path=public"
	print(psqlconn)
	db, err := sql.Open("postgres", psqlconn)
	print("\n db: ", db)
	print("\n dberr: ", err)
	if err != nil {
		print("\n here: ", err)
		return nil, fmt.Errorf("db open: %v", err)
	}
	if err := db.Ping(); err != nil {
		print("\n PONG22")
		return nil, err
	}
	print("\n PONG")
	db.SetMaxOpenConns(p.MaxOpenConn)
	db.SetMaxIdleConns(p.MaxIdleConn)
	db.SetConnMaxLifetime(p.MaxLifetime)
	return &postgres{db: db}, nil
}

// Close ...
func (p *postgres) Close() error {
	if p.db != nil {
		if err := p.db.Close(); err != nil {
			return err
		}
		p.db = nil
	}
	return nil
}

const (
	getTransaction = "select pt.id,pt.identifier ,pt.fare ,pt.cv_number ,pt.created_at ,pt.sap_sent_at,pt.order_id ,pt.itop_id ,pt.paid_amount ,pt.discount_amt ,a.account_code from payment_transactions pt"
)

func (p *postgres) GetData(ctx context.Context) (result []repository.Data, err error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	sql := "select id, name, account_code from accounts where id in (62,26)"
	// sql := `
	// SELECT distinct
	// 	payment_transactions.*,
	// 	departments.name AS department_name,
	// 	accounts.account_code AS account_code,
	// 	users.full_name AS user_name,
	// 	info.invoice_number AS invoice_number,
	// 	info.posting_date AS posting_date,
	// 	info.distance AS distance,
	// 	users.other_information AS other_information
	// 	FROM "payment_transactions"
	// 	LEFT JOIN voucher_profiles vp ON vp.id = payment_transactions.cc_identifier::integer
	// 	LEFT JOIN users ON users.id = vp.user_id
	// 	LEFT JOIN departments ON departments.id = vp.department_id
	// 	LEFT JOIN accounts ON accounts.id = payment_transactions.account_id
	// 	LEFT JOIN payment_transaction_infos info ON info.payment_transaction_id = payment_transactions.id
	// WHERE
	// 	"payment_transactions"."state" IN (2, 3) AND
	// 	"payment_transactions"."payment_type" IN ('ecv', 'edc') AND
	// 	("payment_transactions"."completed_at" BETWEEN '2017-07-01 00:00:00.000000' AND '2021-05-28 00:00:00.000000')
	// ORDER BY "payment_transactions"."picked_up_at" ASC, "payment_transactions"."completed_at" ASC`

	// rows, err := p.db.Query(sql)
	row, err := p.db.Query(sql)
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
	log.Println("data", result)
	return result, nil
}
