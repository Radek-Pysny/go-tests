//go:build integration && !unit

package integration_tests

import (
	"sort"

	"github.com/jackc/pgx/v5"
)

// TestInsertRuleOnlyAffectsFutureInserts demonstrates a key property of PostgreSQL
// rules: a rule rewrites a statement at parse time, so it only acts on inserts issued
// *after* the rule is created. Rows that already exist in `tags` (loaded as a fixture
// before the rule) are never retroactively propagated to `a_tags`.
//
// Scenario:
//  1. Create `tags` and `a_tags`, seed `tags` with a fixture (some a-prefixed, some not).
//  2. Add a rule: inserting an a-prefixed tag also inserts it into `a_tags`.
//  3. Insert new tags and assert:
//     - a-prefixed inserts land in both `tags` and `a_tags`;
//     - other prefixes land only in `tags`;
//     - the pre-rule fixture rows are absent from `a_tags`.
func (s *PostgresTestSuite) Test_PostgresRule_AlsoInsert() {
	// --- 1. schema preparation
	s.Run("schema-preparation", func() {
		_, err := s.conn.Exec(
			s.ctx,
			`CREATE TABLE tags
        (
			id   serial PRIMARY KEY,
			name text   NOT NULL UNIQUE
		)`,
		)

		s.Require().NoError(err, "create tags table")

		_, err = s.conn.Exec(
			s.ctx,
			`CREATE TABLE a_tags
		(
			id   serial PRIMARY KEY,
			name text   NOT NULL UNIQUE
		)`,
		)

		s.Require().NoError(err, "create a_tags table")
	})

	s.Run("fixture-application", func() {
		fixture := []string{"apple", "avocado", "banana", "cherry"}

		for _, name := range fixture {
			_, err := s.conn.Exec(s.ctx, `INSERT INTO tags (name) VALUES ($1)`, name)

			s.Require().NoError(err, "insert fixture %q", name)
		}

		s.ElementsMatch(fixture, s.tagNames("tags"), "tags contains whole fixture")
		s.Empty(s.tagNames("a_tags"), "a_tags is empty")
	})

	// ON INSERT to tags, when the name has prefix 'a', ALSO insert it into a_tags.
	s.Run("add-rule", func() {
		_, err := s.conn.Exec(
			s.ctx,
			` CREATE RULE tags_mirror_a_prefix AS
			ON INSERT TO tags
			WHERE NEW.name LIKE 'a%'
			DO ALSO INSERT INTO a_tags (name) VALUES (NEW.name)
		`,
		)

		s.Require().NoError(err, "create rule")
	})

	// b-prefixed tag is skipped by the rule
	s.Run("insert-b-prefixed-item", func() {
		_, err := s.conn.Exec(s.ctx, `INSERT INTO tags (name) VALUES ($1)`, "blueberry")

		s.Require().NoError(err, "insert non-a-prefixed tag after rule")
		s.Require().Empty(s.tagNames("a_tags"), "a_tags is still empty")
		s.ElementsMatch(
			[]string{"apple", "avocado", "banana", "cherry", "blueberry"},
			s.tagNames("tags"),
			"tags holds the fixture plus every post-rule insert",
		)
	})

	// a-prefixed tag is processed by the rule
	s.Run("insert-1st-a-prefixed-item", func() {
		_, err := s.conn.Exec(s.ctx, `INSERT INTO tags (name) VALUES ($1)`, "apricot")

		s.Require().NoError(err, "insert a-prefixed tag after rule")

		aTags := s.tagNames("a_tags")

		s.Require().Equal(1, len(aTags))
		s.ElementsMatch([]string{"apricot"}, aTags, "a_tags contains just apricot")
		s.ElementsMatch(
			[]string{"apple", "avocado", "banana", "cherry", "blueberry", "apricot"},
			s.tagNames("tags"),
			"tags holds the fixture plus every post-rule insert",
		)
	})

	s.Run("insert-2nd-a-prefixed-item", func() {
		_, err := s.conn.Exec(s.ctx, `INSERT INTO tags (name) VALUES ($1)`, "ambarella")

		s.Require().NoError(err, "insert a-prefixed tag after rule")

		aTags := s.tagNames("a_tags")

		s.Require().Equal(2, len(aTags))
		s.ElementsMatch([]string{"apricot", "ambarella"}, aTags, "a_tags contains just two items")
		s.ElementsMatch(
			[]string{"apple", "avocado", "banana", "cherry", "blueberry", "apricot", "ambarella"},
			s.tagNames("tags"),
			"tags holds the fixture plus every post-rule insert",
		)
	})
}

// tagNames returns the sorted list of names from the given table in the test schema.
func (s *PostgresTestSuite) tagNames(table string) []string {
	rows, err := s.conn.Query(s.ctx, "SELECT name FROM "+table)

	s.Require().NoError(err, "query %s", table)

	names, err := pgx.CollectRows(rows, pgx.RowTo[string])

	s.Require().NoError(err, "collect %s", table)

	sort.Strings(names)
	return names
}
