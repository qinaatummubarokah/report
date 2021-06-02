package repository

import "context"

// type DBReaderWriter interface {
// 	// TODO: Create your repo here
// 	Testing(string) error
// 	GetData(request interface{}) (interface{}, error)
// }

// Service interface
type DBReaderWriter interface {
	GetData(ctx context.Context) (interface{}, error)
}
