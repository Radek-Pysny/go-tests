package unit

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_synctest_bubbleTimeInThePast(t *testing.T) {
	currentTime := time.Now()

	var bubbleTime time.Time
	synctest.Test(t, func(t *testing.T) {
		bubbleTime = time.Now()
	})

	require.False(t, bubbleTime.Equal(currentTime))
	require.True(t, bubbleTime.Before(currentTime))
}

func Test_synctest_bubbleTimeInitialValue(t *testing.T) {
	expectedTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	synctest.Test(t, func(t *testing.T) {
		bubbleTime := time.Now().UTC()

		require.True(t, bubbleTime.Equal(expectedTime))
		require.Equal(t, expectedTime, bubbleTime)
	})
}

func Test_synctest_sleepInSingleGoroutine(t *testing.T) {
	// That test will finish immediately thanks to bubble introduced by synctest.Test, but the time captures are as
	// expected.
	t.Run("in-synctest", func(t *testing.T) {
		const duration = 5 * time.Second
		var (
			older time.Time
			newer time.Time
			start time.Time
			end   time.Time
		)

		start = time.Now()
		synctest.Test(t, func(t *testing.T) {
			older = time.Now()
			time.Sleep(duration)
			newer = time.Now()
		})
		end = time.Now()

		inBubbleDiff := newer.Sub(older)
		require.GreaterOrEqual(t, inBubbleDiff, duration)

		outOfBubbleDiff := end.Sub(start)
		require.Less(t, outOfBubbleDiff, duration)
		require.Less(t, outOfBubbleDiff, 1*time.Millisecond) // this is alchemy TBH

		t.Log("older:", older)
		t.Log("newer:", newer)
		t.Log("diff:", inBubbleDiff)
		t.Log("start:", start)
		t.Log("end:", end)
		t.Log("diff:", outOfBubbleDiff)
	})

	// It will take that 50 ms to finish that unit test.
	t.Run("out-of-synctest", func(t *testing.T) {
		const duration = 50 * time.Millisecond
		var (
			older time.Time
			newer time.Time
		)

		older = time.Now()
		time.Sleep(duration)
		newer = time.Now()

		require.GreaterOrEqual(t, newer.Sub(older), duration)

		t.Log("older:", older)
		t.Log("newer:", newer)
		t.Log("diff:", newer.Sub(older))
	})
}
