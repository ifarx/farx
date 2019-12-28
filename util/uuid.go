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
	/**
	 * 版本
	 */
	Version() byte
	/**
	 * 变种
	 */
	Variant() byte
	String() string
	Bytes() []byte
	Equals(u UUID) bool
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

func (s *_SUUID) Equals(u UUID) bool {
	if s1, ok := u.(*_SUUID); ok {
		return suuid.Equal(s._s, s1._s)
	}
	return bytes.Equal(s.Bytes(), u.Bytes())
}

func RandomUUID() (UUID, error) {
	u4 := suuid.NewV4()
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
