package log

import (
	"testing"

	"jw.lib/logx"
)

func Test_log(t *testing.T) {
	logx.Info("hahahah13")
	logx.Debug("hahahah41")
	logx.Warn("hahahah15")
	logx.Error("hahahah61")
	logx.Error("hahahah71")
}
