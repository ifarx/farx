/**
 * @create  2019/12/28 16:25
 * @since 1.0.0.0
 * @author  ant
 */
package util

import (
	"bytes"
	suuid "github.com/satori/go.uuid"
)

type UUID interface {
	Version() byte
	Variant() byte
	String() string
	Bytes() []byte
	Equal(u UUID) bool
}

type _SUUID struct {
	_s suuid.UUID
}

func (s *_SUUID) Version() byte {
	return s._s.Version()
}

func (s *_SUUID) Variant() byte {
	return s._s.Variant()
}

func (s *_SUUID) String() string {
	return s._s.String()
}

func (s *_SUUID) Bytes() []byte {
	return s._s.Bytes()
}

func (s *_SUUID) Equal(u UUID) bool {
	if s1, ok := u.(*_SUUID); ok {
		return suuid.Equal(s._s, s1._s)
	}
	return bytes.Equal(s.Bytes(), u.Bytes())
}

func RandomUUID() (UUID, error) {
	u4, err := suuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &_SUUID{_s: u4}, nil
}

func FromBytes(bys []byte) (UUID, error) {
	u, err := suuid.FromBytes(bys)
	if err != nil {
		return nil, err
	}
	return &_SUUID{_s: u}, nil
}

func FromString(str string) (UUID, error) {
	u, err := suuid.FromString(str)
	if err != nil {
		return nil, err
	}
	return &_SUUID{_s: u}, nil
}
