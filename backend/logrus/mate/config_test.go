package logrus_mate

import (
	"github.com/sirupsen/logrus"
	"github.com/sniperkit/logrus_mate"
)

func testNewLogger() {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString(
			`{ mike {formatter.name = "json"} }`,
		),
	)
	mikeLoger := mate.Logger("mike")
	mikeLoger.Errorln("Hello Error Level from Mike and my formatter is json")
}

func testStandardLogger() {
	logrus_mate.Hijack(
		logrus.StandardLogger(),
		logrus_mate.ConfigString(
			`{formatter.name = "json"}`,
		),
	)

	logrus.WithField("Field", "A").Debugln("Hello JSON")
}

func testHijackLogger() {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString(
			`{ mike {formatter.name = "json"} }`,
		),
	)
	mate.Hijack(
		logrus.StandardLogger(),
		"mike",
	)
	logrus.Println("hello std logger is hijack by mike")
}

func testFallbackConfigString() {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigString(
			`{ mike {formatter.name = "json"} }`,
		),
		logrus_mate.ConfigFile(
			"mate.conf", // { mike {formatter.name = "text"} }
		),
	)
	mate.Hijack(
		logrus.StandardLogger(),
		"mike",
	)
	logrus.Println("hello std logger is hijack by mike")
}

func testFallbackConfigStringWhileHijack() {
	mate, _ := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigFile(
			"mate.conf", // { mike {formatter.name = "text"} }
		),
	)

	mate.Hijack(logrus.StandardLogger(),
		"mike",
		logrus_mate.ConfigString(
			`{formatter.name = "json"}`,
		),
	)

	logrus.Errorln("hello std logger is hijack by mike")
}
