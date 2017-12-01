package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}
	ctx := r.Context()
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	r.WithContext(ctx) //Replaces context on this request

	// <-ctx.Done() -- With this we could know when this happens but this blocks

	ch := make(chan struct{})
	go func() {
		tm.Next.ServeHTTP(w, r)
		ch <- struct{}{} //If returns send message on channel we finished normally
	}()
	select {
	case <-ch: //Normal
		return
	case <-ctx.Done(): //Request timeout happened
		w.WriteHeader(http.StatusRequestTimeout)
	}
}
