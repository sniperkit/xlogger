package logger

import (
	"testing"

	"github.com/aphistic/sweet"
	"github.com/aphistic/sweet-junit"
	. "github.com/onsi/gomega"

	"github.com/sniperkit/logger/backend/gomol"
	// 	"github.com/sniperkit/logger/backend/logrus"
	// 	"github.com/sniperkit/logger/backend/zap"
)

func TestMain(m *testing.M) {
	RegisterFailHandler(sweet.GomegaFail)

	sweet.Run(m, func(s *sweet.S) {
		s.RegisterPlugin(junit.NewPlugin())

		s.AddSuite(&LoggerSuite{})
		s.AddSuite(&CallerSuite{})
		s.AddSuite(&ConfigSuite{})
		s.AddSuite(&gomol.GomolJSONSuite{})
	})
}
