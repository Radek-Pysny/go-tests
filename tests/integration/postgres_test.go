//go:build integration && !unit

package integration_tests

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/suite"
)

// PostgresTestSuite exercises PostgreSQL assumptions against the shared container
// started in TestMain. Each test runs against a freshly-created `test` schema that is
// dropped completely afterwards, so no state leaks between tests.
type PostgresTestSuite struct {
	suite.Suite

	ctx  context.Context
	conn *pgx.Conn
}

func TestPostgresSuite(t *testing.T) {
	suite.Run(t, new(PostgresTestSuite))
}

func (s *PostgresTestSuite) SetupSuite() {
	s.ctx = context.Background()

	conn, err := pgx.Connect(s.ctx, _postgresDSN)
	s.Require().NoError(err, "connect to postgres")

	s.conn = conn
}

func (s *PostgresTestSuite) TearDownSuite() {
	if s.conn != nil {
		_ = s.conn.Close(s.ctx)
	}
}

// SetupTest creates a clean `test` schema and points the connection's search_path at
// it, so test code may reference objects unqualified.
func (s *PostgresTestSuite) SetupTest() {
	_, err := s.conn.Exec(s.ctx, "DROP SCHEMA IF EXISTS test CASCADE")

	s.Require().NoError(err, "drop stale test schema")

	_, err = s.conn.Exec(s.ctx, "CREATE SCHEMA test")

	s.Require().NoError(err, "create test schema")

	_, err = s.conn.Exec(s.ctx, "SET search_path TO test, public")

	s.Require().NoError(err, "set search_path")
}

// TearDownTest wipes the `test` schema completely.
func (s *PostgresTestSuite) TearDownTest() {
	_, err := s.conn.Exec(s.ctx, "DROP SCHEMA test CASCADE")

	s.Require().NoError(err, "drop test schema")

	_, err = s.conn.Exec(s.ctx, "SET search_path TO \"$user\", public")

	s.Require().NoError(err, "reset search_path")
}

// TestSchemaIsClean asserts each test starts with an empty `test` schema. Together with
// TestCreateTable it proves that objects created in one test do not leak into the next.
func (s *PostgresTestSuite) TestSchemaIsClean() {
	var count int
	err := s.conn.QueryRow(
		s.ctx,
		"SELECT count(*) FROM information_schema.tables WHERE table_schema = 'test'",
	).Scan(&count)

	s.Require().NoError(err)
	s.Equal(0, count, "test schema should start empty")
}

// TestCreateTable creates a table in the `test` schema and uses it. Because the schema
// is dropped in TearDownTest, this table is gone before TestSchemaIsClean runs.
func (s *PostgresTestSuite) TestCreateTable() {
	_, err := s.conn.Exec(s.ctx, "CREATE TABLE widgets (id int PRIMARY KEY, name text)")

	s.Require().NoError(err)

	_, err = s.conn.Exec(s.ctx, "INSERT INTO widgets (id, name) VALUES (1, 'gizmo')")

	s.Require().NoError(err)

	var name string
	err = s.conn.QueryRow(s.ctx, "SELECT name FROM widgets WHERE id = 1").Scan(&name)

	s.Require().NoError(err)
	s.Equal("gizmo", name)
}
