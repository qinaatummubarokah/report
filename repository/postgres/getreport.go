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
	sql := "select id, name from accounts where id in (62,26)"

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
		err = row.Scan(&each.ID, &each.Name)
		log.Println("err scan: ", err)
		if err != nil {
			return nil, fmt.Errorf("row scan: %v", err)
		}
		result = append(result, each)
	}
	// log.Println("data", result)
	return result, nil
}
