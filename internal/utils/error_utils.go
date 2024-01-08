package utils

import "github.com/alichz2001/fidibo/internal/config"

func PanicErr(err error, cfgs ...*config.AppConfig) {
	b := true
	if len(cfgs) > 0 {
		b = cfgs[0].Options.PanicByErr
	}

	if b {
		if err != nil {
			panic("panic: " + err.Error())
		}
	}
}
