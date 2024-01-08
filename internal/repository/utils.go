package repository

import appconfig "github.com/alichz2001/fidibo/internal/config"

func PanicErr(err error, cfgs ...*appconfig.AppConfig) {
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
