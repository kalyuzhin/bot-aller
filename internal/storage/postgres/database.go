package postgres

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Database is a structure to wrap pgx
type Database struct {
	cluster *pgxpool.Pool
}

func newDatabase(cluster *pgxpool.Pool) *Database {
	return &Database{cluster: cluster}
}

// GetPool gets pool
func (db Database) GetPool() *pgxpool.Pool {
	return db.cluster
}

// Get is a get call to db
func (db Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db.cluster, dest, query, args...)
}

// Select is a select call to db
func (db Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db.cluster, dest, query, args...)
}

// Exec is an exec call to db
func (db Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.cluster.Exec(ctx, query, args...)
}

// ExecQueryRow is an exec query row call to db
func (db Database) ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.cluster.QueryRow(ctx, query, args...)
}

// SendBatch is a send batch call
func (db Database) SendBatch(ctx context.Context, batch *pgx.Batch) pgx.BatchResults {
	return db.cluster.SendBatch(ctx, batch)
}

// BeginTx is a beginTx call to db
func (db Database) BeginTx(ctx context.Context, opts pgx.TxOptions) (pgx.Tx, error) {
	return db.cluster.BeginTx(ctx, opts)
}

// Close closes connection
func (db Database) Close() {
	db.cluster.Close()
}
