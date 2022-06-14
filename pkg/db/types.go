package db

import (
	"fmt"
	"strings"
)

type AcmeDb struct {
	Type    DbType
	Address string
	Port    string
	Driver  string
}

// QueryOption is a functional option to populate query builder
type QueryOption func(q AcQuery)

type DbType string

const (
	MySQL DbType = "mysql"
	BQ    DbType = "bigquery"
)

type AcQuery []string

type DB interface {
	Query(query string) ([]byte, error)
}

func NewQuery(opts ...QueryOption) AcQuery {
	q := AcQuery{}
	for _, opt := range opts {
		opt(q)
	}
	return q
}

func (q AcQuery) BuildQuery(opts ...QueryOption) AcQuery {
	for _, opt := range opts {
		opt(q)
	}

	return q
}

func (q AcQuery) String() string {
	return strings.Join(q, " ")
}

func WithID(id string) QueryOption {
	return func(q AcQuery) {
		q = append(q, fmt.Sprintf("ID = %s", id))
	}
}

func WithConjunction(conjunction string) QueryOption {
	return func(q AcQuery) {
		q = append(q, conjunction)
	}
}

func WithTable(table AcmeTable) QueryOption {
	return func(q AcQuery) {
		q = append(q, "TABLE", string(table))
	}
}

func WithAction(action string) QueryOption {
	return func(q AcQuery) {
		q = append(q, action)
	}
}
