package pack_test

import (
	"reflect"

	"github.com/renproject/pack"
	"github.com/renproject/pack/packutil"
	"github.com/renproject/surge/surgeutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Nil", func() {

	numTrials := 10

	Context("when fuzzing", func() {
		It("should not panic", func() {
			for trial := 0; trial < numTrials; trial++ {
				Expect(func() { surgeutil.Fuzz(reflect.TypeOf(pack.Nil{})) }).ToNot(Panic())
				Expect(func() { packutil.JSONFuzz(reflect.TypeOf(pack.Nil{})) }).ToNot(Panic())
			}
		})
	})

	Context("when marshaling and then unmarshaling", func() {
		It("should return itself", func() {
			for trial := 0; trial < numTrials; trial++ {
				Expect(surgeutil.MarshalUnmarshalCheck(reflect.TypeOf(pack.Nil{}))).To(Succeed())
				Expect(packutil.JSONMarshalUnmarshalCheck(reflect.TypeOf(pack.Nil{}))).To(Succeed())
			}
		})
	})

	Context("when marshaling", func() {
		Context("when the buffer is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.MarshalBufTooSmall(reflect.TypeOf(pack.Nil{}))).To(Succeed())
				}
			})
		})

		Context("when the remaining memory quota is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.MarshalRemTooSmall(reflect.TypeOf(pack.Nil{}))).To(Succeed())
				}
			})
		})
	})

	Context("when unmarshaling", func() {
		Context("when the buffer is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.UnmarshalBufTooSmall(reflect.TypeOf(pack.Nil{}))).To(Succeed())
				}
			})
		})

		Context("when the remaining memory quota is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.UnmarshalRemTooSmall(reflect.TypeOf(pack.Nil{}))).To(Succeed())
				}
			})
		})
	})

	Context("when comparing", func() {
		It("should return true", func() {
			Expect(pack.NewNil().Equal(pack.NewNil())).To(BeTrue())
		})
	})

	Context("when stringifying", func() {
		It("should return \"nil\"", func() {
			Expect(pack.NewNil().String()).To(Equal("nil"))
		})
	})

	Context("when getting type information", func() {
		It("should return the nil type", func() {
			Expect(pack.NewNil().Type().Kind()).To(Equal(pack.KindNil))
		})
	})
})
