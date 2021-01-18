package pack_test

import (
	"reflect"
	"testing/quick"

	"github.com/renproject/pack"
	"github.com/renproject/pack/packutil"
	"github.com/renproject/surge/surgeutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("List", func() {

	numTrials := 10

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
})
