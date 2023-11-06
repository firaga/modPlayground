package redisLock

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	r := NewRedis()
	ctx := context.Background()
	key := "test_key"
	expected := "hello"
	r.Set(ctx, key, expected, 1*time.Minute)
	res := r.Get(ctx, key)
	assert.Equal(t, expected, res.Val())
}

func TestLock(t *testing.T) {
	n := "biz"
	l := NewLock(n, 60)
	l.ForceRelease()
	res := l.Lock()
	assert.Equal(t, true, res)
	if res {
		time.Sleep(30 * time.Second)
		l.Release()
	}
}

func TestLockFail(t *testing.T) {
	n := "biz"
	l := NewLock(n, 60)
	res := l.Lock()
	assert.Equal(t, false, res)
	if res {
		assert.Fail(t, "pre lock lock failed")
	}
	assert.Equal(t, false, res)
}
