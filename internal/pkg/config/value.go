package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Value описывает значение конфига.
type Value interface {
	IsNil() bool
	IsEqual(Value) bool

	Bool() bool
	MaybeBool() (bool, error)

	Int64() int64
	MaybeInt64() (int64, error)

	Float64() float64
	Duration() time.Duration
	String() string
}

// concreteValue реализация Value
type concreteValue struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

func (v *concreteValue) Int64() int64 {
	val, _ := v.MaybeInt64()
	return val
}

func (v *concreteValue) MaybeInt64() (int64, error) {
	raw := v.String()

	val, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("value cannot be parsed as int64: %v", err)
	}
	return val, nil
}

func (v *concreteValue) Float64() float64 {
	//TODO implement me
	panic("implement me")
}

func (v *concreteValue) Duration() time.Duration {
	//TODO implement me
	panic("implement me")
}

func (v *concreteValue) String() string {
	//TODO implement me
	panic("implement me")
}

func (v *concreteValue) IsNil() bool {
	return v == nil || v.Value == nil
}

func (v *concreteValue) IsEqual(other Value) bool {
	if v.IsNil() && other.IsNil() {
		return true
	}
	return v.String() == other.String()
}

func (v *concreteValue) Bool() bool {
	val, _ := v.MaybeBool()
	return val
}

func (v *concreteValue) MaybeBool() (bool, error) {
	if v.IsNil() {
		return false, errors.New("value is nil")
	}
	switch val := v.Value.(type) {
	case bool:
		return val, nil
	case string:
		return val == "true", nil
	default:
		return false, errors.Errorf("cannot convert %T to bool", v.Value)
	}
}
