package database

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nleof/goyesql"
)

// Interface only for reference
//type Writer interface {
//	Query(dest interface{}, query string, args ...interface{}) error
//	Run(query string, args ...interface{}) error
//}
//
//type Reader interface {
//	Read(dest interface{}, query string, args ...interface{}) error
//}
//
//type Operations interface {
//	Writer
//	Reader
//}
//
//type Database interface {
//	Operations
//}

func New(connStr, queryPath string) (*Database, error) {
	queries, err := getQueries(queryPath)
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Connect("sqlite3", connStr)
	if err != nil {
		return nil, err
	}
	preparedQueries := make(map[string]*sqlx.Stmt)
	for tag, q := range queries {
		stmt, err := db.Preparex(q)
		if err != nil {
			return nil, fmt.Errorf("failing while creating '%s' query: %w", tag, err)
		}
		preparedQueries[string(tag)] = stmt
	}
	return &Database{db: db, queries: preparedQueries}, nil
}

// getQueries walks through all the SQL files in all provided paths to get the queries.
func getQueries(paths ...string) (goyesql.Queries, error) {
	queries := goyesql.Queries{}
	var err error
	for _, path := range paths {
		err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			q, err := goyesql.ParseFile(path)
			if err != nil {
				return err
			}
			for k, v := range q {
				queries[k] = v
			}
			return nil
		})
		if err != nil {
			return queries, err
		}
	}
	return queries, err
}


type Database struct {
	db *sqlx.DB
	queries map[string]*sqlx.Stmt
}

func (s *Database) Query(dest interface{}, query string, args ...interface{}) error {
	stmt, ok := s.queries[query]
	if !ok {
		return errors.New("query not found")
	}
	if isArrayOrSlice(dest) {
		return stmt.Select(dest, args...)
	}
	return stmt.Get(dest, args...)
}

func (s *Database) Run(query string, args ...interface{}) error {
	stmt, ok := s.queries[query]
	if !ok {
		return errors.New("query not found")
	}
	_, err := stmt.Exec(args...)
	return err
}

// This will be almost the same as query, but it will run over a read-only connection if it exists.
// For the meanwhile we will just run Query behind the scenes.
func (s *Database) Read(dest interface{}, query string, args ...interface{}) error {
	return s.Query(dest, query, args...)
}

func isArrayOrSlice(dst interface{}) bool {
	kind := reflect.Indirect(reflect.ValueOf(dst)).Kind()
	return kind == reflect.Slice || kind == reflect.Array
}