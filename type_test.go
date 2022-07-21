package pack_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"

	"github.com/renproject/pack"
	"github.com/renproject/pack/packutil"
	"github.com/renproject/surge"
	"github.com/renproject/surge/surgeutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Types", func() {

	ts := []reflect.Type{
		reflect.TypeOf(pack.NewBool(false).Type()),
		reflect.TypeOf(pack.NewU8(0).Type()),
		reflect.TypeOf(pack.NewU16(0).Type()),
		reflect.TypeOf(pack.NewU32(0).Type()),
		reflect.TypeOf(pack.NewU64(0).Type()),
		reflect.TypeOf(pack.NewU128([16]byte{}).Type()),
		reflect.TypeOf(pack.NewU256([32]byte{}).Type()),
		reflect.TypeOf(pack.NewString("").Type()),
		reflect.TypeOf(pack.NewBytes([]byte{}).Type()),
		reflect.TypeOf(pack.NewBytes32([32]byte{}).Type()),
		reflect.TypeOf(pack.NewBytes65([65]byte{}).Type()),
		reflect.TypeOf(pack.NewBytes64([64]byte{}).Type()),
		reflect.TypeOf(pack.NewStruct(
			"foo", pack.NewU32(0),
			"bar", pack.NewString(""),
			"baz", pack.NewBytes65([65]byte{}),
		).Type()),
	}
	numTrials := 10

	for _, t := range ts {
		t := t

		Context(fmt.Sprintf("when fuzzing %v", t), func() {
			It("should not panic", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(func() { surgeutil.Fuzz(t) }).ToNot(Panic())
					Expect(func() { packutil.JSONFuzz(t) }).ToNot(Panic())
				}
			})
		})

		Context(fmt.Sprintf("when marshaling and then unmarshaling %v", t), func() {
			It("should return itself", func() {
				for trial := 0; trial < numTrials; trial++ {
					Expect(surgeutil.MarshalUnmarshalCheck(t)).To(Succeed())
					Expect(packutil.JSONMarshalUnmarshalCheck(t)).To(Succeed())
				}
			})
		})

		Context(fmt.Sprintf("when marshaling %v", t), func() {
			Context("when the buffer is too small", func() {
				It("should return itself", func() {
					for trial := 0; trial < numTrials; trial++ {
						Expect(surgeutil.MarshalBufTooSmall(t)).To(Succeed())
					}
				})
			})

			Context("when the remaining memory quota is too small", func() {
				It("should return itself", func() {
					for trial := 0; trial < numTrials; trial++ {
						Expect(surgeutil.MarshalRemTooSmall(t)).To(Succeed())
					}
				})
			})
		})

		Context(fmt.Sprintf("when unmarshaling %v", t), func() {
			Context("when the buffer is too small", func() {
				It("should return itself", func() {
					for trial := 0; trial < numTrials; trial++ {
						Expect(surgeutil.UnmarshalBufTooSmall(t)).To(Succeed())
					}
				})
			})

			Context("when the remaining memory quota is too small", func() {
				It("should return itself", func() {
					for trial := 0; trial < numTrials; trial++ {
						Expect(surgeutil.UnmarshalRemTooSmall(t)).To(Succeed())
					}
				})
			})
		})

		Context(fmt.Sprintf("when checking equality of %v", t), func() {
			Context("if the types are the same", func() {
				It("should return true", func() {
					r := rand.New(rand.NewSource(time.Now().UnixNano()))
					for trial := 0; trial < numTrials; trial++ {
						x, ok := quick.Value(t, r)
						Expect(ok).To(BeTrue())
						xType := x.Interface().(pack.Type)

						Expect(xType.Equals(xType)).To(BeTrue())
					}
				})
			})

			Context("if the types are different", func() {
				It("should return false", func() {
					r := rand.New(rand.NewSource(time.Now().UnixNano()))
					for trial := 0; trial < numTrials; trial++ {
						x, ok := quick.Value(t, r)
						Expect(ok).To(BeTrue())
						xType := x.Interface().(pack.Type)
						xTypeBytes, err := surge.ToBinary(xType)
						Expect(err).ToNot(HaveOccurred())

						var yType pack.Type
						var yTypeBytes []byte
						for yTypeBytes == nil || bytes.Equal(xTypeBytes, yTypeBytes) {
							value := pack.Generate(r, 5, true, true)
							yType = value.Interface().(pack.Value).Type()
							yTypeBytes, err = surge.ToBinary(yType)
							Expect(err).ToNot(HaveOccurred())
						}

						Expect(xType.Equals(yType)).To(BeFalse())
					}
				})
			})
		})
	}
})
