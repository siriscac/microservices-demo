// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package money

import (
	"errors"

	models "github.com/siriscac/microservices-demo/fontend/hipster/models"
)

const (
	nanosMin = -999999999
	nanosMax = +999999999
	nanosMod = 1000000000
)

var (
	ErrInvalidValue        = errors.New("one of the specified money values is invalid")
	ErrMismatchingCurrency = errors.New("mismatching currency codes")
)

// IsValid checks if specified value has a valid units/nanos signs and ranges.
func IsValid(m models.HipstershopMoney) bool {
	return signMatches(m) && validNanos(m.Nanos)
}

func signMatches(m models.HipstershopMoney) bool {
	return m.Nanos == 0 || m.Units == 0 || (m.Nanos < 0) == (m.Units < 0)
}

func validNanos(nanos int32) bool { return nanosMin <= nanos && nanos <= nanosMax }

// IsZero returns true if the specified money value is equal to zero.
func IsZero(m models.HipstershopMoney) bool { return m.Units == 0 && m.Nanos == 0 }

// IsPositive returns true if the specified money value is valid and is
// positive.
func IsPositive(m models.HipstershopMoney) bool {
	return IsValid(m) && m.Units > 0 || (m.Units == 0 && m.Nanos > 0)
}

// IsNegative returns true if the specified money value is valid and is
// negative.
func IsNegative(m models.HipstershopMoney) bool {
	return IsValid(m) && m.Units < 0 || (m.Units == 0 && m.Nanos < 0)
}

// AreSameCurrency returns true if values l and r have a currency code and
// they are the same values.
func AreSameCurrency(l, r models.HipstershopMoney) bool {
	return l.GetCurrencyCode() == r.GetCurrencyCode() && l.GetCurrencyCode() != ""
}

// AreEquals returns true if values l and r are the equal, including the
// currency. This does not check validity of the provided values.
func AreEquals(l, r models.HipstershopMoney) bool {
	return l.GetCurrencyCode() == r.GetCurrencyCode() &&
		l.Units == r.Units && l.Nanos == r.Nanos
}

// Negate returns the same amount with the sign negated.
func Negate(m models.HipstershopMoney) models.HipstershopMoney {
	return models.HipstershopMoney{
		Units:        -m.Units,
		Nanos:        -m.Nanos,
		CurrencyCode: m.GetCurrencyCode()}
}

// Must panics if the given error is not nil. This can be used with other
// functions like: "m := Must(Sum(a,b))".
func Must(v models.HipstershopMoney, err error) models.HipstershopMoney {
	if err != nil {
		panic(err)
	}
	return v
}

// Sum adds two values. Returns an error if one of the values are invalid or
// currency codes are not matching (unless currency code is unspecified for
// both).
func Sum(l, r models.HipstershopMoney) (models.HipstershopMoney, error) {
	if !IsValid(l) || !IsValid(r) {
		return models.HipstershopMoney{}, ErrInvalidValue
	} else if l.GetCurrencyCode() != r.GetCurrencyCode() {
		return models.HipstershopMoney{}, ErrMismatchingCurrency
	}
	units := l.Units + r.Units
	nanos := l.Nanos + r.Nanos

	if (units == 0 && nanos == 0) || (units > 0 && nanos >= 0) || (units < 0 && nanos <= 0) {
		// same sign <units, nanos>
		units += int64(nanos / nanosMod)
		nanos = nanos % nanosMod
	} else {
		// different sign. nanos guaranteed to not to go over the limit
		if units > 0 {
			units--
			nanos += nanosMod
		} else {
			units++
			nanos -= nanosMod
		}
	}

	return models.HipstershopMoney{
		Units:        units,
		Nanos:        nanos,
		CurrencyCode: l.GetCurrencyCode()}, nil
}

// MultiplySlow is a slow multiplication operation done through adding the value
// to itself n-1 times.
func MultiplySlow(m models.HipstershopMoney, n uint32) models.HipstershopMoney {
	out := m
	for n > 1 {
		out = Must(Sum(out, m))
		n--
	}
	return out
}
