package logger

import (
	lf "github.com/sniperkit/xlogger/pkg/fields"
)

type NotFound struct{}

func (nf *NotFound) Sync() error { return nil }

// func (nf *NotFound) WithFields(fields Fields) *Logger                                    { return nf }

// Arguments are handled in the manner of fmt.Print.
func (nf *NotFound) Panic(format string, args ...interface{}) {}
func (nf *NotFound) Print(format string, args ...interface{}) {}
func (nf *NotFound) Debug(format string, args ...interface{}) {}
func (nf *NotFound) Info(format string, args ...interface{})  {}
func (nf *NotFound) Warn(format string, args ...interface{})  {}
func (nf *NotFound) Error(format string, args ...interface{}) {}
func (nf *NotFound) Fatal(format string, args ...interface{}) {}

// Arguments are handled in the manner of fmt.Printf
func (nf *NotFound) Panicf(format string, args ...interface{}) {}
func (nf *NotFound) Printf(format string, args ...interface{}) {}
func (nf *NotFound) Debugf(format string, args ...interface{}) {}
func (nf *NotFound) Infof(format string, args ...interface{})  {}
func (nf *NotFound) Warnf(format string, args ...interface{})  {}
func (nf *NotFound) Errorf(format string, args ...interface{}) {}
func (nf *NotFound) Fatalf(format string, args ...interface{}) {}

// Arguments are handled in the manner of fmt.Println.
func (nf *NotFound) Panicln(format string, args ...interface{}) {}
func (nf *NotFound) Println(format string, args ...interface{}) {}
func (nf *NotFound) Debugln(format string, args ...interface{}) {}
func (nf *NotFound) Infoln(format string, args ...interface{})  {}
func (nf *NotFound) Warnln(format string, args ...interface{})  {}
func (nf *NotFound) Errorln(format string, args ...interface{}) {}
func (nf *NotFound) Fatalln(format string, args ...interface{}) {}

// Prefix logger with initial fields set by the user and handle arguments in the manner of fmt.Println.
func (nf *NotFound) PanicWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) PrintWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) DebugWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) InfoWithFields(fields lf.Fields, format string, args ...interface{})  {}
func (nf *NotFound) WarnWithFields(fields lf.Fields, format string, args ...interface{})  {}
func (nf *NotFound) ErrorWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) FatalWithFields(fields lf.Fields, format string, args ...interface{}) {}

// Prefix logger with initial fields set by the user and handle arguments in the manner of fmt.Println.
func (nf *NotFound) PrintlnWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) PaniclnWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) DebuglnWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) InfolnWithFields(fields lf.Fields, format string, args ...interface{})  {}
func (nf *NotFound) WarnlnWithFields(fields lf.Fields, format string, args ...interface{})  {}
func (nf *NotFound) ErrorlnWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) FatallnWithFields(fields lf.Fields, format string, args ...interface{}) {}

// Prefix logger with initial fields set by the user and handle arguments in the manner of fmt.Println.
func (nf *NotFound) PrintfWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) PanicfWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) DebugfWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) InfofWithFields(fields lf.Fields, format string, args ...interface{})  {}
func (nf *NotFound) WarnfWithFields(fields lf.Fields, format string, args ...interface{})  {}
func (nf *NotFound) ErrorfWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) FatalfWithFields(fields lf.Fields, format string, args ...interface{}) {}
