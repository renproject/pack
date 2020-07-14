package pack_test

import (
	"reflect"

	"github.com/renproject/pack"
	"github.com/renproject/pack/packutil"
	"github.com/renproject/surge/surgeutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bool", func() {

	numTrials := 10

	Context("when fuzzing", func() {
		It("should not panic", func() {
			for trial := 0; trial < numTrials; trial++ {
				Expect(func() { surgeutil.Fuzz(reflect.TypeOf(pack.Bool(false))) }).ToNot(Panic())
				Expect(func() { packutil.JSONFuzz(reflect.TypeOf(pack.Bool(false))) }).ToNot(Panic())
			}
		})
	})

	Context("when marshaling and then unmarshaling", func() {
		It("should return itself", func() {
			for trial := 0; trial < numTrials; trial++ {
				Expect(surgeutil.MarshalUnmarshalCheck(reflect.TypeOf(pack.Bool(false)))).To(Succeed())
				Expect(packutil.JSONMarshalUnmarshalCheck(reflect.TypeOf(pack.Bool(false)))).To(Succeed())
			}
		})
	})

	Context("when marshaling", func() {
		Context("when the buffer is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.MarshalBufTooSmall(reflect.TypeOf(pack.Bool(false)))).To(Succeed())
				}
			})
		})

		Context("when the remaining memory quota is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.MarshalRemTooSmall(reflect.TypeOf(pack.Bool(false)))).To(Succeed())
				}
			})
		})
	})

	Context("when unmarshaling", func() {
		Context("when the buffer is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.UnmarshalBufTooSmall(reflect.TypeOf(pack.Bool(false)))).To(Succeed())
				}
			})
		})

		Context("when the remaining memory quota is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.UnmarshalRemTooSmall(reflect.TypeOf(pack.Bool(false)))).To(Succeed())
				}
			})
		})
	})

	Context("when comparing", func() {
		Context("when the booleans are equal", func() {
			It("should return true", func() {
				Expect(pack.NewBool(true).Equal(pack.NewBool(true))).To(BeTrue())
			})
		})

		Context("when the booleans are not equal", func() {
			It("should return false", func() {
				Expect(pack.NewBool(true).Equal(pack.NewBool(false))).To(BeFalse())
			})
		})
	})

	Context("when stringifying", func() {
		Context("when the boolean is true", func() {
			It("should return \"true\"", func() {
				Expect(pack.NewBool(true).String()).To(Equal("true"))
			})
		})

		Context("when the boolean is false", func() {
			It("should return \"false\"", func() {
				Expect(pack.NewBool(false).String()).To(Equal("false"))
			})
		})
	})

	Context("when getting type information", func() {
		It("should return the bool type", func() {
			Expect(pack.NewBool(false).Type().Kind()).To(Equal(pack.KindBool))
		})
	})
})
