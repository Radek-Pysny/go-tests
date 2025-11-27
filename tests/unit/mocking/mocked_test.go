package mocking

import (
	"testing"

	"github.com/Radek-Pysny/go-tests/tests/mocks/github.com/Radek-Pysny/go-tests/putmock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_mock_notCalled(t *testing.T) {
	m := putmock.NewMockMocked(t)

	// if we do not call maybe, it might fail
	m.EXPECT().Predicate(mock.Anything, mock.Anything).Maybe()
}

func Test_mock_calledOnce(t *testing.T) {
	t.Run("no-call-count-modifier", func(t *testing.T) {
		m := putmock.NewMockMocked(t)
		m.EXPECT().Predicate(mock.Anything, 1).Return(true, nil)

		returned, err := m.Predicate(nil, 1)

		require.NoError(t, err)
		require.True(t, returned)
	})

	t.Run("once", func(t *testing.T) {
		m := putmock.NewMockMocked(t)
		m.EXPECT().Predicate(mock.Anything, 1).Return(true, nil).Once()

		returned, err := m.Predicate(nil, 1)

		require.NoError(t, err)
		require.True(t, returned)
	})

	t.Run("times-1", func(t *testing.T) {
		m := putmock.NewMockMocked(t)
		m.EXPECT().Predicate(mock.Anything, 1).Return(true, nil).Times(1)

		returned, err := m.Predicate(nil, 1)

		require.NoError(t, err)
		require.True(t, returned)
	})
}

func Test_mock_calledTwice(t *testing.T) {
	const twice = 2

	t.Run("no-call-count-modifier", func(t *testing.T) {
		m := putmock.NewMockMocked(t)
		m.EXPECT().Predicate(mock.Anything, 1).Return(true, nil)

		for range twice {
			returned, err := m.Predicate(nil, 1)

			require.NoError(t, err)
			require.True(t, returned)
		}
	})

	t.Run("twice", func(t *testing.T) {
		m := putmock.NewMockMocked(t)

		m.EXPECT().Predicate(mock.Anything, 1).Return(true, nil).Twice()

		for range twice {
			returned, err := m.Predicate(nil, 1)

			require.NoError(t, err)
			require.True(t, returned)
		}
	})

	t.Run("times-2", func(t *testing.T) {
		m := putmock.NewMockMocked(t)

		m.EXPECT().Predicate(mock.Anything, 1).Return(true, nil).Times(twice)

		for range twice {
			returned, err := m.Predicate(nil, 1)

			require.NoError(t, err)
			require.True(t, returned)
		}
	})
}

func Test_mock_calledThreeTimes(t *testing.T) {
	const times = 3

	t.Run("no-call-count-modifier", func(t *testing.T) {
		m := putmock.NewMockMocked(t)
		m.EXPECT().Predicate(mock.Anything, 1).Return(true, nil)

		for range times {
			returned, err := m.Predicate(nil, 1)

			require.NoError(t, err)
			require.True(t, returned)
		}
	})

	t.Run("times-2", func(t *testing.T) {
		m := putmock.NewMockMocked(t)

		m.EXPECT().Predicate(mock.Anything, 1).Return(true, nil).Times(times)

		for range times {
			returned, err := m.Predicate(nil, 1)

			require.NoError(t, err)
			require.True(t, returned)
		}
	})
}

func Test_mock_smallerExpectations(t *testing.T) {
	m := putmock.NewMockMocked(t)
	for i := range 3 {
		m.EXPECT().Predicate(mock.Anything, i+1).Return(true, nil)
	}

	for i := range 10 {
		returned, err := m.Predicate(nil, i)

		require.NoError(t, err)
		require.Equal(t, i > 0, returned)
	}
}
