package bct_test

import (
	"strings"

	"github.com/loveandpeople-DAG/goClient/consts"
	"github.com/loveandpeople-DAG/goClient/curl"
	"github.com/loveandpeople-DAG/goClient/curl/bct"
	"github.com/loveandpeople-DAG/goClient/trinary"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("BCT Curl", func() {

	DescribeTable("Hash",
		func(src []trinary.Trits, hashLen int, rounds curl.CurlRounds) {
			c := bct.NewCurl(rounds)
			err := c.Absorb(src, len(src[0]))
			Expect(err).ToNot(HaveOccurred())

			dst := make([]trinary.Trits, len(src))
			err = c.Squeeze(dst, hashLen)
			Expect(err).ToNot(HaveOccurred())

			for i := range dst {
				// compare against the non batched Curl
				Expect(dst[i]).To(Equal(CurlSum(src[i], hashLen, rounds)))
			}
		},
		Entry("Curl-P-81: trits and hash", Trits(bct.MaxBatchSize, consts.HashTrinarySize), consts.HashTrinarySize, curl.CurlP81),
		Entry("Curl-P-81: multi trits and hash", Trits(bct.MaxBatchSize, consts.TransactionTrinarySize), consts.HashTrinarySize, curl.CurlP81),
		Entry("Curl-P-81: trits and multi squeeze", Trits(bct.MaxBatchSize, consts.HashTrinarySize), 3*consts.HashTrinarySize, curl.CurlP81),
		Entry("Curl-P-27: trits and hash", Trits(bct.MaxBatchSize, consts.HashTrinarySize), consts.HashTrinarySize, curl.CurlP27),
		Entry("Curl-P-27: multi trits and hash", Trits(bct.MaxBatchSize, consts.TransactionTrinarySize), consts.HashTrinarySize, curl.CurlP27),
		Entry("Curl-P-27: trits and multi squeeze", Trits(bct.MaxBatchSize, consts.HashTrinarySize), 3*consts.HashTrinarySize, curl.CurlP27),
	)

	It("Clone", func() {
		a := []trinary.Trits{trinary.MustTrytesToTrits(strings.Repeat("A", consts.HashTrytesSize))}
		b := []trinary.Trits{trinary.MustTrytesToTrits(strings.Repeat("B", consts.HashTrytesSize))}

		c1 := bct.NewCurlP81()
		err := c1.Absorb(a, len(a[0]))
		Expect(err).ToNot(HaveOccurred())

		c2 := c1.Clone()
		err = c1.Absorb(b, len(b[0]))
		Expect(err).ToNot(HaveOccurred())
		err = c2.Absorb(b, len(b[0]))
		Expect(err).ToNot(HaveOccurred())

		hash1 := make([]trinary.Trits, 1)
		err = c1.Squeeze(hash1, consts.HashTrinarySize)
		Expect(err).ToNot(HaveOccurred())
		hash2 := make([]trinary.Trits, 1)
		err = c2.Squeeze(hash2, consts.HashTrinarySize)
		Expect(err).ToNot(HaveOccurred())

		Expect(hash2[0]).To(Equal(hash1[0]))
	})
})

func Trits(size int, tritsCount int) []trinary.Trits {
	trytesCount := tritsCount / consts.TritsPerTryte
	src := make([]trinary.Trits, size)
	for i := range src {
		trytes := strings.Repeat("ABC", trytesCount/3+1)[:trytesCount-2] + trinary.IntToTrytes(int64(i), 2)
		src[i] = trinary.MustTrytesToTrits(trytes)
	}
	return src
}

func CurlSum(data trinary.Trits, tritsCount int, rounds curl.CurlRounds) trinary.Trits {
	c := curl.NewCurl(rounds)
	if err := c.Absorb(data); err != nil {
		panic(err)
	}
	out, err := c.Squeeze(tritsCount)
	if err != nil {
		panic(err)
	}
	return out
}
