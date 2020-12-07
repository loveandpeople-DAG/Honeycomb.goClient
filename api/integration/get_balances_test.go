package integration_test

import (
	"strings"

	. "github.com/loveandpeople-DAG/goClient/api"
	. "github.com/loveandpeople-DAG/goClient/api/integration/samples"
	. "github.com/loveandpeople-DAG/goClient/consts"

	. "github.com/loveandpeople-DAG/goClient/trinary"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
)

var _ = Describe("GetBalances()", func() {

	api, err := ComposeAPI(HTTPClientSettings{}, nil)
	if err != nil {
		panic(err)
	}

	Context("call", func() {

		It("resolves to correct response", func() {
			balances, err := api.GetBalances(SampleAddressesWithChecksum)
			Expect(err).ToNot(HaveOccurred())
			Expect(*balances).To(Equal(Balances{
				Balances:       []uint64{99, 0, 1},
				Milestone:      strings.Repeat("M", 81),
				MilestoneIndex: 1,
			}))
		})
	})

	Context("invalid input", func() {
		It("returns an error for invalid addresses", func() {
			_, err := api.GetBalances(Hashes{"balalaika"})
			Expect(errors.Cause(err)).To(Equal(ErrInvalidHash))
		})
	})

})
