package dispatcher

import (
	"github.com/valyala/gorpc"
	"time"
)

var Disp *gorpc.Dispatcher

func init() {
	Disp = gorpc.NewDispatcher()

	Disp.AddFunc("matrix", func(req *Request1) *Response1 {
		// res.Answer = req.M1.MulByMatrix(req.M2)
		t := time.Now()
		r := Response1{}
		r.Answer = req.M1.MulByMatrix(req.M2)
		s := time.Since(t)
		r.Time = s
		return &r
	})
}
