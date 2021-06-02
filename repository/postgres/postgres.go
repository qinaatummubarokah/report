package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"report/repository"
	"sync"
)

type postgres struct {
	db *sql.DB
	mu sync.RWMutex
}

func (p *postgres) GetReport(ctx context.Context) (result []repository.Data, err error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	sql := "select id, name from accounts where id in (62,26)"

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
		err = row.Scan(&each.ID, &each.Name)
		log.Println("err scan: ", err)
		if err != nil {
			return nil, fmt.Errorf("row scan: %v", err)
		}
		result = append(result, each)
	}
	log.Println("data", result)
	return result, nil
}

// NewPostgresConn ...
func NewPostgresConn(conf repository.DBConfiguration) (repository.DBReaderWriter, error) {
	connURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s%s",
		conf.DBUser,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
		conf.DBOptions)

	logger.Info(fmt.Sprintf(
		"Postgres connection: postgres://%s:%s@%s:%s/%s%s",
		conf.DBUser,
		conf.DBPassword,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
		conf.DBOptions))

	db, err := sql.Open(driverName, connURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	logger.Info("DB Says PONG!")

	// GORM for Auto Migrate
	// gormClient, err := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: db,
	// }), &gorm.Config{})
	// if err != nil {
	// 	logger.Info(fmt.Sprintf("create gorm client instance failed with message: %s", err.Error()))
	// 	os.Exit(1)
	// }
	// gormClient.AutoMigrate(
	// 	&datamaster.WritePoolRequest{}, &datamaster.WriteRegionRequest{},
	// )
	// logger.Info("Table Migrated")

	// // Adding Dummy Data using GORM
	// regionID, _ := uuid.NewV4()
	// if err := gormClient.Create(&datamaster.WriteRegionRequest{
	// 	RegionID:  regionID,
	// 	Kelurahan: "Klender",
	// 	Kecamatan: "Duren Sawit",
	// 	JenisKota: "KOTA",
	// 	Kota:      "Jakarta Timur",
	// 	Provinsi:  "DKI Jakarta",
	// 	KodePos:   "13470",
	// }).Error; err != nil {
	// 	logger.Error(err)
	// }
	// logger.Info("Data Inserted")

	return &postgresConn{
		db: db,
	}, nil
}
