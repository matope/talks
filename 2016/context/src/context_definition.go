package main

import "time"

// START OMIT
package context

type Context interface {
	Deadline() (deadline time.Time, ok bool) // Contextのデッドライン時刻(設定されていれば)
	Done() <-chan struct{}                   // Contextの処理の中断が要求されたときにcloseするchan // HL
	Err() error                              // closeした理由を返す(Canceled or DeadlineExceeded)
	Value(key interface{}) interface{}       // contextにkeyで関連付けられた値を返す
}

// END OMIT
