package dispatcher

import (
	"github.com/valyala/gorpc"
	"time"
)

var Disp *gorpc.Dispatcher

func init() {
	Disp = gorpc.NewDispatcher()

	Disp.AddFunc("PiCalc", func(req *Request0) *Response0 {
		t := time.Now()
		r := Response0{}
		r.Answer = req.M1.PiCalc()
		r.Time = time.Since(t)
		return &r
	})

	Disp.AddFunc("MatrixByMatrix", func(req *Request1) *Response1 {
		// res.Answer = req.M1.MulByMatrix(req.M2)
		t := time.Now()
		r := Response1{}
		r.Answer = req.M1.MulByMatrix(req.M2)
		r.Time = time.Since(t)
		return &r
	})

	Disp.AddFunc("MatrixByVVector", func(req *Request2) *Response2 {
		t := time.Now()
		r := Response2{}
		r.Answer = req.M1.MulByVVector(req.M2)
		r.Time = time.Since(t)
		return &r
	})

	Disp.AddFunc("VVectorByHVector", func(req *Request3) *Response3 {
		t := time.Now()
		r := Response3{}
		r.Answer = req.M1.MulByHVector(req.M2)
		r.Time = time.Since(t)
		return &r
	})

	Disp.AddFunc("Brute", func(req *Request4) *Response4 {
		t := time.Now()
		r := Response4{}
		r.Answer = req.M1.Run(1)
		r.Time = time.Since(t)
		return &r
	})

}
