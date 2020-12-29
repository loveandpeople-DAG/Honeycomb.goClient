package curl_test

import (
	"strings"

	"github.com/loveandpeople-DAG/goClient/consts"
	. "github.com/loveandpeople-DAG/goClient/curl"
	"github.com/loveandpeople-DAG/goClient/trinary"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Curl", func() {

	DescribeTable("Hash",
		func(in trinary.Trytes, expSqueeze trinary.Trytes, rounds ...CurlRounds) {

			By("tryte", func() {
				c := NewCurl(rounds...)
				err := c.AbsorbTrytes(trinary.MustPad(in, consts.HashTrytesSize))
				Expect(err).ToNot(HaveOccurred())
				squeeze, err := c.SqueezeTrytes(len(expSqueeze) * consts.TritsPerTryte)
				Expect(err).ToNot(HaveOccurred())
				Expect(squeeze).To(Equal(expSqueeze))
			})

			By("trits", func() {
				c := NewCurl(rounds...)
				err := c.Absorb(trinary.MustPadTrits(trinary.MustTrytesToTrits(in), consts.HashTrinarySize))
				Expect(err).ToNot(HaveOccurred())
				squeeze, err := c.Squeeze(len(expSqueeze) * consts.TritsPerTryte)
				Expect(err).ToNot(HaveOccurred())
				Expect(squeeze).To(Equal(trinary.MustTrytesToTrits(expSqueeze)))
			})
		},
		Entry("Curl-P-81: empty trytes", "", consts.NullHashTrytes, CurlP81),
		Entry("Curl-P-81: normal trytes", "A", "TJVKPMTAMIZVBVHIVQUPTKEMPROEKV9SB9COEDQYRHYPTYSKQIAN9PQKMZHCPO9TS9BHCORFKW9CQXZEE", CurlP81),
		Entry("Curl-P-81: normal trytes #2", "Z", "FA9WYZSJJWSD9AEEBOGGDHFTMIZVHFURFLJLFBTNENDDCMSXGAGLXFMYZTAMKVIYDQSZEDKXSWVAOPZMK"),
		Entry("Curl-P-81: normal trytes #3", "NOPQRSTUVWXYZ9ABSDEFGHIJKLM", "GWFZSXPZPAFSVPEGEIVWOTD9MY9KVP9HYVCIWSJEITEGVOVGQGV99RONTWDXOPUBIQPIWXK9L9OHZYFUB", CurlP81),
		Entry("Curl-P-81: long absorb", strings.Repeat("ABC", consts.TransactionTrytesSize/3), "UHZVKZCGDIPNGFNPBNFZGIM9GAKYLCPTHTRFRXMNDJLZNXSGRPREFWTBKZWVTKV9BISPXEECVIXFJERAC", CurlP81),
		Entry("Curl-P-81: long squeeze", "ABC", "LRJMQXFSZSLCIMKZTWFTEIHKWJZMUOHPSOVXZOHOEVHC9D9DROUQGRPTBZWOIJFTMGMXEYKXEJROQLWNUPSFJJRVTLUUJYW9GBQVXNCAUEGEBV9IJQ9TWFDHCFPUUYPCYLACTAIK9UZAJLVXLI9NPGCJN9ICFTEIYY", CurlP81),
		Entry("Curl-P-27: empty trytes", "", consts.NullHashTrytes, CurlP27),
		Entry("Curl-P-27: normal trytes", "TWENTYSEVEN", "RQPYXJPRXEEPLYLAHWTTFRXXUZTV9SZPEVOQ9FZATCXJOZLZ9A9BFXTUBSHGXN9OOA9GWIPGAAWEDVNPN", CurlP27),
	)

	It("Clone", func() {
		a := strings.Repeat("A", consts.HashTrytesSize)
		b := strings.Repeat("B", consts.HashTrytesSize)

		c1 := NewCurlP81()
		err := c1.AbsorbTrytes(a)
		Expect(err).ToNot(HaveOccurred())

		c2 := c1.Clone()
		err = c1.AbsorbTrytes(b)
		Expect(err).ToNot(HaveOccurred())
		err = c2.AbsorbTrytes(b)
		Expect(err).ToNot(HaveOccurred())

		Expect(c2.MustSqueezeTrytes(consts.HashTrinarySize)).To(Equal(c1.MustSqueezeTrytes(consts.HashTrinarySize)))
	})
})
