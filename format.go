package errors

import (
	"fmt"
	"io"
	"runtime"
	"time"
)

type Format func(s fmt.State, verb rune, e Errors)

func DefaultFormatFn(s fmt.State, verb rune, e Errors) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('+'):
			_, _ = fmt.Fprintf(s, "- %s\n", e.Error())
			for i, pc := range e.PCS() {
				fn := runtime.FuncForPC(pc)
				if fn == nil {
					_, _ = io.WriteString(s, "unknown")
				} else {
					file, line := fn.FileLine(pc)
					_, filename := fileName(file)
					if i == 0 {
						_, _ = fmt.Fprintf(s, "\t[T] %s\n\t[F] %s\n\t[F] %s:%d \n", e.OccurTime().Format(time.RFC3339), fn.Name(), filename, line)
					} else {
						_, _ = fmt.Fprintf(s, "\t[F] %s\n\t[F] %s:%d \n", fn.Name(), filename, line)
					}
				}
			}
			if e.Cause() != nil {
				hasCause, ok := e.Cause().(Errors)
				if !ok {
					_, _ = fmt.Fprintf(s, "%v\n", e.Cause())
				} else {
					hasCause.Format(s, verb)
				}
			}
		case s.Flag('-'):
			_, _ = io.WriteString(s, "{")
			_, _ = fmt.Fprintf(s, `"msg":"%s", "occurTime":"%s", "stack":[`, e.Error(), e.OccurTime().Format(time.RFC3339))
			for i, pc := range e.PCS() {
				if i > 0 {
					_, _ = io.WriteString(s, ",")
				}
				fn := runtime.FuncForPC(pc)
				if fn == nil {
					_, _ = fmt.Fprintf(s, `{"fn":"%s", "file":"%s", "line":%d}`, "unknown", "unknown", 0)
				} else {
					file, line := fn.FileLine(pc)
					_, filename := fileName(file)
					_, _ = fmt.Fprintf(s, `{"fn":"%s", "file":"%s", "line":%d}`, fn.Name(), filename, line)
				}
			}
			_, _ = io.WriteString(s, "]")
			if e.Cause() != nil {
				_, _ = io.WriteString(s, ",")
				hasCause, ok := e.Cause().(Errors)
				if !ok {
					_, _ = fmt.Fprintf(s, `"cause":{"msg":"%s"}`, e.Cause().Error())
				} else {
					_, _ = io.WriteString(s, `"cause":`)
					hasCause.Format(s, verb)
				}
			}
			_, _ = io.WriteString(s, "}")
		default:
			_, _ = fmt.Fprintf(s, "%s", e.Error())
		}
	}
}

func JsonFormatFn(s fmt.State, verb rune, e Errors) {
	switch verb {
	case 'v':
		switch {
		case s.Flag('+'):
			_, _ = io.WriteString(s, "{")
			_, _ = fmt.Fprintf(s, `"msg":"%s", "occurTime":"%s", "stack":[`, e.Error(), e.OccurTime().Format(time.RFC3339))
			for i, pc := range e.PCS() {
				if i > 0 {
					_, _ = io.WriteString(s, ",")
				}
				fn := runtime.FuncForPC(pc)
				if fn == nil {
					_, _ = fmt.Fprintf(s, `{"fn":"%s", "home":"%s", "file":"%s", "line":%d}`, "unknown", "unknown", "unknown", 0)
				} else {
					file, line := fn.FileLine(pc)
					home, filename := fileName(file)
					_, _ = fmt.Fprintf(s, `{"fn":"%s", "home":"%s", "file":"%s", "line":%d}`, fn.Name(), home, filename, line)
				}
			}
			_, _ = io.WriteString(s, "]")
			if e.Cause() != nil {
				_, _ = io.WriteString(s, ",")
				hasCause, ok := e.Cause().(Errors)
				if !ok {
					_, _ = fmt.Fprintf(s, `"cause":{"msg":"%s"}`, e.Cause().Error())
				} else {
					_, _ = io.WriteString(s, `"cause":`)
					hasCause.Format(s, verb)
				}
			}
			_, _ = io.WriteString(s, "}")
		default:
			_, _ = fmt.Fprintf(s, "%s", e.Error())
		}
	}
}
