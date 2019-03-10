// Code generated by immutableGen. DO NOT EDIT.

// My favourite license

package coretest_test

//go:generate echo "hello world"
//immutableVet:skipFile

import (
	"myitcv.io/immutable"
)

//
// xtestA is an immutable type and has the following template:
//
// 	struct {
// 		*XTestB
//
// 		age	int
// 	}
//
type xtestA struct {
	anonfield_XTestB *XTestB
	field_age        int

	mutable bool
	__tmpl  *_Imm_xtestA
}

var _ immutable.Immutable = new(xtestA)
var _ = new(xtestA).__tmpl

func (s *xtestA) AsMutable() *xtestA {
	if s.Mutable() {
		return s
	}

	res := *s
	res.mutable = true
	return &res
}

func (s *xtestA) AsImmutable(v *xtestA) *xtestA {
	if s == nil {
		return nil
	}

	if s == v {
		return s
	}

	s.mutable = false
	return s
}

func (s *xtestA) Mutable() bool {
	return s.mutable
}

func (s *xtestA) WithMutable(f func(si *xtestA)) *xtestA {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable(s)

	return res
}

func (s *xtestA) WithImmutable(f func(si *xtestA)) *xtestA {
	prev := s.mutable
	s.mutable = false
	f(s)
	s.mutable = prev

	return s
}

func (s *xtestA) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true
	{
		v := s.anonfield_XTestB

		if v != nil && !v.IsDeeplyNonMutable(seen) {
			return false
		}
	}
	return true
}
func (s *xtestA) Name() string {
	return s.XTestB().Name()
}
func (s *xtestA) SetName(n string) *xtestA {
	v1 := s.XTestB().SetName(n)
	v0 := s.SetXTestB(v1)
	return v0
}
func (s *xtestA) XTestB() *XTestB {
	return s.anonfield_XTestB
}

// SetXTestB is the setter for XTestB()
func (s *xtestA) SetXTestB(n *XTestB) *xtestA {
	if s.mutable {
		s.anonfield_XTestB = n
		return s
	}

	res := *s
	res.anonfield_XTestB = n
	return &res
}
func (s *xtestA) age() int {
	return s.field_age
}

// setAge is the setter for Age()
func (s *xtestA) setAge(n int) *xtestA {
	if s.mutable {
		s.field_age = n
		return s
	}

	res := *s
	res.field_age = n
	return &res
}

//
// XTestB is an immutable type and has the following template:
//
// 	struct {
// 		Name string
// 	}
//
type XTestB struct {
	field_Name string

	mutable bool
	__tmpl  *_Imm_XTestB
}

var _ immutable.Immutable = new(XTestB)
var _ = new(XTestB).__tmpl

func (s *XTestB) AsMutable() *XTestB {
	if s.Mutable() {
		return s
	}

	res := *s
	res.mutable = true
	return &res
}

func (s *XTestB) AsImmutable(v *XTestB) *XTestB {
	if s == nil {
		return nil
	}

	if s == v {
		return s
	}

	s.mutable = false
	return s
}

func (s *XTestB) Mutable() bool {
	return s.mutable
}

func (s *XTestB) WithMutable(f func(si *XTestB)) *XTestB {
	res := s.AsMutable()
	f(res)
	res = res.AsImmutable(s)

	return res
}

func (s *XTestB) WithImmutable(f func(si *XTestB)) *XTestB {
	prev := s.mutable
	s.mutable = false
	f(s)
	s.mutable = prev

	return s
}

func (s *XTestB) IsDeeplyNonMutable(seen map[interface{}]bool) bool {
	if s == nil {
		return true
	}

	if s.Mutable() {
		return false
	}

	if seen == nil {
		return s.IsDeeplyNonMutable(make(map[interface{}]bool))
	}

	if seen[s] {
		return true
	}

	seen[s] = true
	return true
}
func (s *XTestB) Name() string {
	return s.field_Name
}

// SetName is the setter for Name()
func (s *XTestB) SetName(n string) *XTestB {
	if s.mutable {
		s.field_Name = n
		return s
	}

	res := *s
	res.field_Name = n
	return &res
}