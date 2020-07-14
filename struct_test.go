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

var _ = Describe("Struct", func() {

	numTrials := 10

	Context("when fuzzing", func() {
		It("should not panic", func() {
			for trial := 0; trial < numTrials; trial++ {
				Expect(func() { surgeutil.Fuzz(reflect.TypeOf(pack.Struct{})) }).ToNot(Panic())
				Expect(func() { packutil.JSONFuzz(reflect.TypeOf(pack.Struct{})) }).ToNot(Panic())
			}
		})
	})

	Context("when marshaling", func() {
		Context("when the buffer is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.MarshalBufTooSmall(reflect.TypeOf(pack.Struct{}))).To(Succeed())
				}
			})
		})

		Context("when the remaining memory quota is too small", func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.MarshalRemTooSmall(reflect.TypeOf(pack.Struct{}))).To(Succeed())
				}
			})
		})
	})

	Context("when getting field values by name", func() {
		Context("when the field exists", func() {
			It("should return the field value", func() {
				s := pack.NewStruct("foo", pack.NewU64(42))
				v := s.Get("foo")
				Expect(v.(pack.U64).Equal(pack.NewU64(42))).To(BeTrue())
			})
		})

		Context("when the field exists", func() {
			It("should return the field value", func() {
				s := pack.NewStruct("foo", pack.NewU64(42))
				v := s.Get("bar")
				Expect(v).To(BeNil())
			})
		})
	})

	Context("when setting field values by name", func() {
		Context("when the field exists", func() {
			It("should replace the field value", func() {
				s := pack.NewStruct("foo", pack.NewU64(42))
				v := s.Set("foo", pack.NewU64(420))
				Expect(v.(pack.U64).Equal(pack.NewU64(42))).To(BeTrue())
				v = s.Get("foo")
				Expect(v.(pack.U64).Equal(pack.NewU64(420))).To(BeTrue())
			})
		})

		Context("when the field exists", func() {
			It("should do nothing", func() {
				s := pack.NewStruct("foo", pack.NewU64(42))
				v := s.Set("bar", pack.NewU64(420))
				Expect(v).To(BeNil())
				v = s.Get("bar")
				Expect(v).To(BeNil())
			})
		})
	})

	Context("when getting type information", func() {
		It("should return the struct type", func() {
			f := func(x pack.Struct) bool {
				Expect(x.Type().Kind()).To(Equal(pack.KindStruct))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when stringifying a struct", func() {
		It("should equal the JSON representation", func() {
			f := func(x pack.Struct) bool {
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
