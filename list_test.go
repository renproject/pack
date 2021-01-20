package pack_test

import (
	"math/rand"
	"reflect"
	"testing/quick"

	"github.com/renproject/pack"
	"github.com/renproject/pack/packutil"
	"github.com/renproject/surge"
	"github.com/renproject/surge/surgeutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("List", func() {

	numTrials := 10

	randomType := func() pack.Type {
		r := rand.New(rand.NewSource(GinkgoRandomSeed()))
		switch r.Int() % 13 {
		// Nil
		case 0:
			return pack.Bool(false).Type()
		// Scalar
		case 1:
			return pack.U8(0).Type()
		case 2:
			return pack.U16(0).Type()
		case 3:
			return pack.U32(0).Type()
		case 4:
			return pack.U64(0).Type()
		case 5:
			return pack.U128{}.Type()
		case 6:
			return pack.U256{}.Type()
		// Bytes
		case 7:
			return pack.String("").Type()
		case 8:
			return pack.Bytes{}.Type()
		case 9:
			return pack.Bytes32{}.Type()
		case 10:
			return pack.Bytes65{}.Type()
		// Abstract
		case 11:
			return pack.Struct{}.Type()
		case 12:
			return pack.List{}.Type()
		}
		panic("unreachable")
	}

	Context("when fuzzing", func() {
		It("should not panic", func() {
			for trial := 0; trial < numTrials; trial++ {
				Expect(func() { surgeutil.Fuzz(reflect.TypeOf(pack.List{})) }).ToNot(Panic())
				Expect(func() { packutil.JSONFuzz(reflect.TypeOf(pack.List{})) }).ToNot(Panic())
			}
		})
	})

	Context("when marshaling", func() {
		Context("when the buffer is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.MarshalBufTooSmall(reflect.TypeOf(pack.List{}))).To(Succeed())
				}
			})
		})

		Context("when the remaining memory quota is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.MarshalRemTooSmall(reflect.TypeOf(pack.List{}))).To(Succeed())
				}
			})
		})
	})

	Context("when marshaling and unmarshaling to binary", func() {
		It("should equal itself", func() {
			f := func(x pack.List) bool {
				data, err := surge.ToBinary(x)
				Expect(err).ToNot(HaveOccurred())
				y := pack.List{
					T: x.T,
				}
				err = surge.FromBinary(&y, data)
				Expect(err).ToNot(HaveOccurred())
				Expect(y).To(Equal(x))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when getting type information", func() {
		It("should return the list type", func() {
			f := func(x pack.List) bool {
				Expect(x.Type().Kind()).To(Equal(pack.KindList))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when stringifying a list", func() {
		It("should equal the JSON representation", func() {
			f := func(x pack.List) bool {
				stringified := x.String()
				data, err := x.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())
				Expect(stringified).To(Equal(string(data)))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when constructing an empty list", func() {
		It("should return a list with no elements", func() {
			for trial := 0; trial < numTrials; trial++ {
				t := randomType()
				list := pack.EmptyList(t)
				Expect(list.T).To(Equal(t))
				Expect(list.Elems).To(HaveLen(0))
			}
		})
	})

	Context("when unmarshaling an empty list", func() {
		Context("if the list has a type specified", func() {
			It("should not return an error", func() {
				t := randomType()
				list := pack.List{
					T: t,
				}
				bytes, err := surge.ToBinary(list)
				Expect(err).ToNot(HaveOccurred())
				err = surge.FromBinary(&list, bytes)
				Expect(err).ToNot(HaveOccurred())
				Expect(list.T.Kind()).To(Equal(t.Kind()))
				Expect(list.Elems).To(HaveLen(0))
			})
		})

		Context("if the list does not have a type specified", func() {
			It("should return an error", func() {
				list := pack.List{}
				bytes, err := surge.ToBinary(list)
				Expect(err).ToNot(HaveOccurred())
				err = surge.FromBinary(&list, bytes)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
