//go:build separate && !unit && !integration

package allfails

import (
	"testing"

	"github.com/Radek-Pysny/go-tests/tests/mocks/github.com/Radek-Pysny/go-tests/putmock"
	"github.com/stretchr/testify/mock"
)

func Test_mock_notEnoughCalls(t *testing.T) {
	t.Run("not-called-at-all", func(t *testing.T) {
		m := putmock.NewMockMocked(t)

		m.EXPECT().Predicate(mock.Anything, mock.Anything)
	})

	t.Run("called-just-once-1", func(t *testing.T) {
		m := putmock.NewMockMocked(t)
		m.EXPECT().Predicate(mock.Anything, mock.Anything).Return(false, nil).Twice()

		_, _ = m.Predicate(nil, nil)
	})

	t.Run("called-just-once-2", func(t *testing.T) {
		m := putmock.NewMockMocked(t)
		m.EXPECT().Predicate(mock.Anything, mock.Anything).Return(false, nil).Times(2)

		_, _ = m.Predicate(nil, nil)
	})

	t.Run("called-just-once-3", func(t *testing.T) {
		m := putmock.NewMockMocked(t)
		m.EXPECT().Predicate(mock.Anything, mock.Anything).Return(false, nil).Times(3)

		_, _ = m.Predicate(nil, nil)
	})
}

func Test_mock_unexpectedCalls(t *testing.T) {
	t.Run("not-expected", func(t *testing.T) {
		m := putmock.NewMockMocked(t)

		_, _ = m.Predicate(nil, nil)
	})

	t.Run("expected-just-once-1", func(t *testing.T) {
		m := putmock.NewMockMocked(t)

		m.EXPECT().Predicate(mock.Anything, mock.Anything).Return(false, nil).Once()

		_, _ = m.Predicate(nil, nil)
		_, _ = m.Predicate(nil, nil)
	})

	t.Run("expected-just-once-2", func(t *testing.T) {
		m := putmock.NewMockMocked(t)

		m.EXPECT().Predicate(mock.Anything, mock.Anything).Return(false, nil).Once()

		_, _ = m.Predicate(nil, nil)
		_, _ = m.Predicate(nil, nil)
		_, _ = m.Predicate(nil, nil)
		_, _ = m.Predicate(nil, nil)
		_, _ = m.Predicate(nil, nil)
	})

	t.Run("expected-just-once-3", func(t *testing.T) {
		m := putmock.NewMockMocked(t)

		m.EXPECT().Predicate(mock.Anything, mock.Anything).Return(false, nil).Times(1)

		_, _ = m.Predicate(nil, nil)
		_, _ = m.Predicate(nil, nil)
	})

	t.Run("expected-just-once-4", func(t *testing.T) {
		m := putmock.NewMockMocked(t)

		m.EXPECT().Predicate(mock.Anything, mock.Anything).Return(false, nil).Times(1)

		_, _ = m.Predicate(nil, nil)
		_, _ = m.Predicate(nil, nil)
		_, _ = m.Predicate(nil, nil)
		_, _ = m.Predicate(nil, nil)
		_, _ = m.Predicate(nil, nil)
	})
}
