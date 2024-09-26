package main

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"math/rand"
)

// ---------------------------------------

type randomNumbers interface {
	randomMaximumNumber(max int, useless int) int
}

type calculateRandomNumbers struct {
}

func (s calculateRandomNumbers) randomMaximumNumber(m int, useless int) int {
	_ = useless
	return rand.Intn(m)
}

// ---------------------------------------

type mockRandomNumbers struct {
	mock.Mock
}

func (s *mockRandomNumbers) randomMaximumNumber(m int, useless int) int {
	args := s.Called(m, useless)

	return args.Get(0).(int)
}

func newMockRandomNumbers() *mockRandomNumbers {
	return &mockRandomNumbers{}
}

// ---------------------------------------

func randomTimes99(s randomNumbers) int {
	return 99 * s.randomMaximumNumber(727, 1000)
}

func main() {
	s := new(calculateRandomNumbers)

	fmt.Println(randomTimes99(s))
}
