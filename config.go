package errors

import (
	"fmt"
	"sync"
	"time"
)

var _cfg *config
var _once = new(sync.Once)

func init() {
	_once.Do(func() {
		_cfg = &config{loc: time.Local, depth: 1, skip: 3, formatFn: DefaultFormatFn}
	})
}

type config struct {
	loc      *time.Location
	depth    int
	skip     int
	formatFn Format
}

func (c *config) SetTimeLocation(loc *time.Location) {
	if loc == nil {
		panic("errors set time location failed, loc is nil")
	}
	c.loc = loc
}

func (c *config) SetStack(depth int, skip int) {
	if depth < 1 {
		panic(fmt.Errorf("errors set stack failed, depth valued %d is invalid", depth))
	}
	if skip < 1 {
		panic(fmt.Errorf("errors set stack failed, skip valued %d is invalid", skip))
	}
	c.depth = depth
	c.skip = skip
}

func (c *config) SetFormatFunc(fn Format) {
	if fn == nil {
		panic("errors set format func failed")
	}
	c.formatFn = fn
}

func Configure() *config {
	return _cfg
}
