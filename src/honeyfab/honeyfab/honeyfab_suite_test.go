package honeyfab_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHoneyfab(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Honeyfab Suite")
}
