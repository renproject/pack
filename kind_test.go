package pack_test

import (
	"encoding/json"
	"math/rand"

	"github.com/renproject/pack"
	"github.com/renproject/surge"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kind", func() {

	numTrials := 100

	randomKind := func() pack.Kind {
		r := rand.New(rand.NewSource(GinkgoRandomSeed()))
		switch r.Int() % 14 {
		// Nil
		case 0:
			return pack.KindNil
		// Scalar
		case 1:
			return pack.KindU8
		case 2:
			return pack.KindU16
		case 3:
			return pack.KindU32
		case 4:
			return pack.KindU64
		case 5:
			return pack.KindU128
		case 6:
			return pack.KindU256
		// Bytes
		case 7:
			return pack.KindString
		case 8:
			return pack.KindBytes
		case 9:
			return pack.KindBytes32
		case 10:
			return pack.KindBytes65
		case 11:
			return pack.KindBytes64
		// Abstract
		case 12:
			return pack.KindStruct
		case 13:
			return pack.KindList
		}
		panic("unreachable")
	}

	Context("when marshaling and unmarshaling to binary", func() {
		It("should equal itself", func() {
			for trial := 0; trial < numTrials; trial++ {
				expected := randomKind()
				data, err := surge.ToBinary(expected)
				Expect(err).ToNot(HaveOccurred())
				got := pack.KindNil
				err = surge.FromBinary(&got, data)
				Expect(err).ToNot(HaveOccurred())
				Expect(expected).To(Equal(got))
			}
		})
	})

	Context("when marshaling and unmarshaling to JSON", func() {
		It("should equal itself", func() {
			for trial := 0; trial < numTrials; trial++ {
				expected := randomKind()
				data, err := json.Marshal(expected)
				Expect(err).ToNot(HaveOccurred())
				got := pack.KindNil
				err = json.Unmarshal(data, &got)
				Expect(err).ToNot(HaveOccurred())
				Expect(expected).To(Equal(got))
			}
		})
	})
})
