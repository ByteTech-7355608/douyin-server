package social_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSocial(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Social Suite")
}
