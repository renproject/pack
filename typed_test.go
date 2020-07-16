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

var _ = Describe("Typed", func() {

	numTrials := 100

	Context("when fuzzing", func() {
		It("should not panic", func() {
			for trial := 0; trial < numTrials; trial++ {
				surgeutil.Fuzz(reflect.TypeOf(pack.Typed{}))
				Expect(func() { surgeutil.Fuzz(reflect.TypeOf(pack.Typed{})) }).ToNot(Panic())
				Expect(func() { packutil.JSONFuzz(reflect.TypeOf(pack.Typed{})) }).ToNot(Panic())
			}
		})
	})

	Context("when marshaling and then unmarshaling", func() {
		It("should return itself", func() {
			for trial := 0; trial < numTrials; trial++ {
				Expect(surgeutil.MarshalUnmarshalCheck(reflect.TypeOf(pack.Typed{}))).To(Succeed())
				Expect(packutil.JSONMarshalUnmarshalCheck(reflect.TypeOf(pack.Typed{}))).To(Succeed())
			}
		})
	})

	Context("when marshaling", func() {
		Context("when the buffer is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.MarshalBufTooSmall(reflect.TypeOf(pack.Typed{}))).To(Succeed())
				}
			})
		})

		Context("when the remaining memory quota is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.MarshalRemTooSmall(reflect.TypeOf(pack.Typed{}))).To(Succeed())
				}
			})
		})
	})

	Context("when unmarshaling", func() {
		Context("when the buffer is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.UnmarshalBufTooSmall(reflect.TypeOf(pack.Typed{}))).To(Succeed())
				}
			})
		})

		Context("when the remaining memory quota is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.UnmarshalRemTooSmall(reflect.TypeOf(pack.Typed{}))).To(Succeed())
				}
			})
		})
	})

	Context("when getting type information", func() {
		It("should return the struct type", func() {
			f := func(x pack.Typed) bool {
				Expect(x.Type().Kind()).To(Equal(pack.KindStruct))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when stringifying a typed struct", func() {
		It("should equal the JSON representation", func() {
			f := func(x pack.Typed) bool {
				stringified := x.String()
				data, err := x.MarshalJSON()
				Expect(err).ToNot(HaveOccurred())
				Expect(stringified).To(Equal(string(data)))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when unmarshaling a non-struct", func() {
		It("should return an error", func() {
			typed := pack.NewTyped("foo", pack.NewString("bar"))
			data := []byte(`{"t":"string","v":"bar"}`)
			Expect(typed.UnmarshalJSON(data)).ToNot(Succeed())
		})
	})
})
