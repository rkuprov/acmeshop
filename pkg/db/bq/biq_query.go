package bq

import (
	"context"

	"AcmeShop/pkg/db"
	"google.golang.org/api/bigquery/v2"
)

var _ db.DB = (*BQDB)(nil)

var BQUtil, _ = NewBQ(context.Background())

type BQDB struct {
	Conn *bigquery.Service
}

func NewBQ(c context.Context) (BQDB, error) {
	s, err := bigquery.NewService(c)
	if err != nil {
		return BQDB{}, err
	}
	return BQDB{
		Conn: s,
	}, nil
}

func (b BQDB) Query(query string) ([]byte, error) {
	var qr *bigquery.QueryResponse
	var err error
	qr, err = b.Conn.Jobs.Query("", &bigquery.QueryRequest{Query: query}).Do()
	if err != nil {
		return nil, err
	}
	var ret []byte
	for _, r := range qr.Rows {
		b, err := r.MarshalJSON()
		if err != nil {
			return nil, err
		}
		ret = append(ret, b...)
	}
	return ret, nil
}

func (b BQDB) Service() *bigquery.Service {
	return b.Conn
}
