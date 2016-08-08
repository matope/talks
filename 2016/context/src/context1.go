package main

import (
	"context"
	"fmt"
	"time"
)

// START TYPE
type Context interface {
	Deadline() (deadline time.Time, ok bool) // Contextのデッドライン時刻(設定されていれば)
	Done() <-chan struct{}                   // Contextがcancelまたはデッドラインを超過したらcloseするchanを返す //HL
	Err() error                              // Doneがcloseした理由を返す(Canceled or DeadlineExceeded) //HL
	Value(key interface{}) interface{}       // contextにkeyで関連付けられた値を返す
}

// END TYPE

func main() {
	// 5秒後にタイムアウトするContextを作成
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // タイムアウトのリソース解放を忘れないこと

	fmt.Println("doing something...")
	select {
	case <-ctx.Done(): // デッドラインを超過するとClose
		fmt.Println(ctx.Err())
	}
}

// START OMIT
// 第一引数にContextを受ける(契約)
func slowProcess(ctx context.Context) error {
	// ...
	select {
	case <-ctx.Done(): // ctx.Done()がcloseしたら処理を終了する // HL
		return ctx.Err() // context.Canceled または context.DeadlineExceeded
	}
	// END OMIT
	return nil
}
