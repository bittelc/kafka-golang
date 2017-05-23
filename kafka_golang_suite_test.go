package kafka_golang_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestKafkaGolang(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KafkaGolang Suite")
}
