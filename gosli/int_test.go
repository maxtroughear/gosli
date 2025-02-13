package gosli

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	suite "github.com/stretchr/testify/suite"
)

type IntTestSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestIntFactory(t *testing.T) {
	suite.Run(t, new(IntTestSuite))
}
