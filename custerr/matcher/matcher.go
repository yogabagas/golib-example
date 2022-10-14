package matcher

import (
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"gitlab.sicepat.tech/platform/golib/custerr"
)

func HaveErrorMessage(m types.GomegaMatcher) types.GomegaMatcher {
	return WithTransform(func(c custerr.ErrChain) string { return c.Message }, m)
}

func HaveErrorType(m types.GomegaMatcher) types.GomegaMatcher {
	return WithTransform(func(c custerr.ErrChain) error { return c.Type }, m)
}
