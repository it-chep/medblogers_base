package domain_error

import (
	"errors"
	"fmt"

	"github.com/samber/lo"
)

const (
	grpc     string = "gRPC"
	internal string = "internal"
)

// InternalCode возможные код ошибок
type InternalCode int

const (
	// External не будет учитываться
	External InternalCode = 0
	// NotFound описывает ошибку, когда не найдена запрошенная сущность
	NotFound InternalCode = 1
	// Aborted описывает ошибку, отражающую конфликтную ситуацию
	Aborted InternalCode = 2
	// FailedPrecondition ошибка бизнес валидации
	FailedPrecondition InternalCode = 3
	// InvalidArgument описывает ошибку некорректных данных
	InvalidArgument InternalCode = 4
)

// InternalError описывает ошибку, которая произошла внутри системы
type InternalError struct {
	// Код ошибки
	code InternalCode
	// Классификация ошибок
	classification string
	// Сообщение
	msg string
	// Базовая ошибка
	base error
}

// NewExternalErr создает новую ошибку с кодом internal NotFound
func NewExternalErr(msg string, base error) error {
	return &InternalError{
		classification: grpc,
		code:           External,
		msg:            msg,
		base:           base,
	}
}

// NewNotFoundErr создает новую ошибку с кодом internal NotFound
func NewNotFoundErr(msg string, base error) error {
	return &InternalError{
		classification: internal,
		code:           NotFound,
		msg:            msg,
		base:           base,
	}
}

// NewAbortedErr создает новую ошибку с кодом internal Aborted
func NewAbortedErr(msg string) error {
	return &InternalError{
		classification: internal,
		code:           Aborted,
		msg:            msg,
	}
}

// NewBusinessErr создает новую ошибку с кодом internal FailedPrecondition
func NewBusinessErr(msg string) error {
	return &InternalError{
		classification: internal,
		code:           FailedPrecondition,
		msg:            msg,
	}
}

func (err *InternalError) Error() string {
	return err.msg
}

func (err *InternalError) Unwrap() error {
	if err.base == nil {
		return errors.New(err.Error())
	}
	return fmt.Errorf("%s : %s", err.msg, err.base.Error())
}

// LogUnwrap возвращает внутреннюю ошибку для лога если это возможно
func LogUnwrap(err error) error {
	// Если мы понимаем, что это наша внутренняя ошибка, то необходимо взять описание
	// ошибки из вложенной ошибки, так как она содержит детали ошибки
	if monitoringErr, ok := lo.ErrorsAs[*InternalError](err); ok {
		return monitoringErr.Unwrap()
	}
	return err
}

// As представляет ошибку как InternalError
func As(err error) (*InternalError, bool) {
	return lo.ErrorsAs[*InternalError](err)
}

// IsCode проверяет что ошибка является *InternalError и имеет код code
func IsCode(err error, code InternalCode) bool {
	if rErr, ok := As(err); ok && rErr.IsCode(code) {
		return true
	}
	return false
}

// IsCode проверяет как классифицируется ошибка
func (err *InternalError) IsCode(code InternalCode) bool {
	return err.code == code
}

// SkipCode возвращает nil, если ее код совпадает с переданным,
// иначе возвращает ошибку
func SkipCode(err error, code InternalCode) error {
	internalError, ok := lo.ErrorsAs[*InternalError](err)
	if !ok {
		return err
	}

	if internalError.IsCode(code) {
		return nil
	}

	return internalError
}
