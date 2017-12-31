package fields

import (
	"time"

	"github.com/sniperkit/xlogger/pkg/core"
)

//Fields is a struct to send paramaters to log messages
type Fields map[string]interface{}

func (f Fields) NormalizeTimeValues() Fields {
	for key, val := range f {
		switch v := val.(type) {
		case time.Time:
			f[key] = v.Format(core.JSONTimeFormat)
		default:
			f[key] = v
		}
	}
	return f
}

//FieldType is a type identifier for logger fields
type FieldType int8

//Field is a struct to send paramaters to log messages
type Field struct {
	key     string
	val     interface{}
	valType FieldType
}

/*
func String(key, val string) Field {
	return Field{key: key, val: val, valType: StringField}
}

func Int(key string, val int) Field {
	return Field{key: key, val: val, valType: IntField}
}

func Int64(key string, val int64) Field {
	return Field{key: key, val: val, valType: Int64Field}
}

func Float(key string, val float32) Field {
	return Field{key: key, val: val, valType: FloatField}
}

func Float64(key string, val float64) Field {
	return Field{key: key, val: val, valType: Float64Field}
}

func Bool(key string, val bool) Field {
	return Field{key: key, val: val, valType: BoolField}
}

func Duration(key string, val time.Duration) Field {
	return Field{key: key, val: val, valType: DurationField}
}

func Time(key string, val time.Time) Field {
	return Field{key: key, val: val, valType: TimeField}
}

func Struct(key string, val interface{}) Field {
	return Field{key: key, val: val, valType: StructField}
}
*/
