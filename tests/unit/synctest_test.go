package unit

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_synctest_sleepInSingleGoroutine(t *testing.T) {
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
