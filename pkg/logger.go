package logger

import (
	lf "github.com/sniperkit/xlogger/pkg/fields"
)

type Logger interface {

	////////////////////////////////////
	//////// Custom wrappers
	Sync() error

	////////////////////////////////////
	//////// Default wrappers

	// Arguments are handled in the manner of fmt.Print.
	Print(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Panic(format string, args ...interface{})

	// Arguments are handled in the manner of fmt.Printf
	Printf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	// Arguments are handled in the manner of fmt.Println.
	Println(format string, args ...interface{})
	Infoln(format string, args ...interface{})
	Debugln(format string, args ...interface{})
	Warnln(format string, args ...interface{})
	Errorln(format string, args ...interface{})
	Fatalln(format string, args ...interface{})
	Panicln(format string, args ...interface{})

	////////////////////////////////////
	//////// Prefix logger embedding initial fields set by the user

	// Arguments are handled in the manner of fmt.Print.
	PanicWithFields(fields lf.Fields, format string, args ...interface{})
	PrintWithFields(fields lf.Fields, format string, args ...interface{})
	DebugWithFields(fields lf.Fields, format string, args ...interface{})
	InfoWithFields(fields lf.Fields, format string, args ...interface{})
	WarnWithFields(fields lf.Fields, format string, args ...interface{})
	ErrorWithFields(fields lf.Fields, format string, args ...interface{})
	FatalWithFields(fields lf.Fields, format string, args ...interface{})

	// Arguments are handled in the manner of fmt.Printf
	PanicfWithFields(fields lf.Fields, format string, args ...interface{})
	PrintfWithFields(fields lf.Fields, format string, args ...interface{})
	DebugfWithFields(fields lf.Fields, format string, args ...interface{})
	InfofWithFields(fields lf.Fields, format string, args ...interface{})
	WarnfWithFields(fields lf.Fields, format string, args ...interface{})
	ErrorfWithFields(fields lf.Fields, format string, args ...interface{})
	FatalfWithFields(fields lf.Fields, format string, args ...interface{})

	// Arguments are handled in the manner of fmt.Println.
	PaniclnWithFields(fields lf.Fields, format string, args ...interface{})
	PrintlnWithFields(fields lf.Fields, format string, args ...interface{})
	DebuglnWithFields(fields lf.Fields, format string, args ...interface{})
	InfolnWithFields(fields lf.Fields, format string, args ...interface{})
	WarnlnWithFields(fields lf.Fields, format string, args ...interface{})
	ErrorlnWithFields(fields lf.Fields, format string, args ...interface{})
	FatallnWithFields(fields lf.Fields, format string, args ...interface{})
}
