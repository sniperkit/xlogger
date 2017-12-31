package logger

import (
	lf "github.com/sniperkit/xlogger/pkg/fields"
)

type NotFound struct{}

// func (nf *NotFound) WithFields(fields Fields) *Logger                                    { return nf }
func (nf *NotFound) Debug(format string, args ...interface{})                               {}
func (nf *NotFound) Info(format string, args ...interface{})                                {}
func (nf *NotFound) Warning(format string, args ...interface{})                             {}
func (nf *NotFound) Error(format string, args ...interface{})                               {}
func (nf *NotFound) Fatal(format string, args ...interface{})                               {}
func (nf *NotFound) DebugWithFields(fields lf.Fields, format string, args ...interface{})   {}
func (nf *NotFound) InfoWithFields(fields lf.Fields, format string, args ...interface{})    {}
func (nf *NotFound) WarningWithFields(fields lf.Fields, format string, args ...interface{}) {}
func (nf *NotFound) ErrorWithFields(fields lf.Fields, format string, args ...interface{})   {}
func (nf *NotFound) FatalWithFields(fields lf.Fields, format string, args ...interface{})   {}
func (nf *NotFound) Sync() error                                                            { return nil }
