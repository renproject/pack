package pack_test

import (
	"bytes"
	"fmt"
	"math/big"
	"reflect"
	"testing/quick"

	"github.com/renproject/pack"
	"github.com/renproject/pack/packutil"
	"github.com/renproject/surge/surgeutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ints", func() {

	numTrials := 100

	ts := []reflect.Type{
		reflect.TypeOf(pack.NewU8(uint8(0))),
		reflect.TypeOf(pack.NewU16(uint16(0))),
		reflect.TypeOf(pack.NewU32(uint32(0))),
		reflect.TypeOf(pack.NewU64(uint64(0))),
		reflect.TypeOf(pack.NewU128([16]byte{})),
		reflect.TypeOf(pack.NewU256([32]byte{})),
	}

	for _, t := range ts {
		t := t

		Context(fmt.Sprintf("when fuzzing %v", t), func() {
			It("should not panic", func() {
				Expect(func() { surgeutil.Fuzz(t) }).ToNot(Panic())
				Expect(func() { packutil.JSONFuzz(t) }).ToNot(Panic())
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

		Context(fmt.Sprintf("when doing arithmetic on %v ints", t), func() {

			Context("when adding zero", func() {
				It("should be equal to itself", func() {
					for trial := 0; trial < numTrials; trial++ {
						Expect(packutil.AddZeroCheck(t)).To(Succeed())
					}
				})
			})

			Context("when subtracting zero", func() {
				It("should be equal to itself", func() {
					for trial := 0; trial < numTrials; trial++ {
						Expect(packutil.SubZeroCheck(t)).To(Succeed())
					}
				})
			})

			Context("when adding with overflow", func() {
				It("should panic", func() {
					for trial := 0; trial < numTrials; trial++ {
						Expect(func() { packutil.AddOverflow(t, false) }).To(Panic())
						Expect(func() { packutil.AddOverflow(t, true) }).To(Panic())
					}
				})
			})

			Context("when subtracting with underflow", func() {
				It("should panic", func() {
					for trial := 0; trial < numTrials; trial++ {
						Expect(func() { packutil.SubUnderflow(t, false) }).To(Panic())
						Expect(func() { packutil.SubUnderflow(t, true) }).To(Panic())
					}
				})
			})

			Context("when adding and then subtracting", func() {
				It("should equal itself", func() {
					for trial := 0; trial < numTrials; trial++ {
						Expect(packutil.AddSubCheck(t)).To(Succeed())
					}
				})
			})
		})
	}

	Context("when creating uint16 from smaller ints", func() {
		It("should equal the smaller ints", func() {
			f := func(x1 uint8) bool {
				Expect(pack.NewU16FromU8(pack.NewU8(x1)).String()).To(Equal(pack.NewU8(x1).String()))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when creating uint32 from smaller ints", func() {
		It("should equal the smaller ints", func() {
			f := func(x1 uint8, x2 uint16) bool {
				Expect(pack.NewU32FromU8(pack.NewU8(x1)).String()).To(Equal(pack.NewU8(x1).String()))
				Expect(pack.NewU32FromU16(pack.NewU16(x2)).String()).To(Equal(pack.NewU16(x2).String()))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when creating uint64 from smaller ints", func() {
		It("should equal the smaller ints", func() {
			f := func(x1 uint8, x2 uint16, x3 uint32) bool {
				Expect(pack.NewU64FromU8(pack.NewU8(x1)).String()).To(Equal(pack.NewU8(x1).String()))
				Expect(pack.NewU64FromU16(pack.NewU16(x2)).String()).To(Equal(pack.NewU16(x2).String()))
				Expect(pack.NewU64FromU32(pack.NewU32(x3)).String()).To(Equal(pack.NewU32(x3).String()))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when creating uint128 from smaller ints", func() {
		It("should equal the smaller ints", func() {
			f := func(x1 uint8, x2 uint16, x3 uint32, x4 uint64) bool {
				Expect(pack.NewU128FromU8(pack.NewU8(x1)).String()).To(Equal(pack.NewU8(x1).String()))
				Expect(pack.NewU128FromU16(pack.NewU16(x2)).String()).To(Equal(pack.NewU16(x2).String()))
				Expect(pack.NewU128FromU32(pack.NewU32(x3)).String()).To(Equal(pack.NewU32(x3).String()))
				Expect(pack.NewU128FromU64(pack.NewU64(x4)).String()).To(Equal(pack.NewU64(x4).String()))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when creating uint256 from smaller ints", func() {
		It("should equal the smaller ints", func() {
			f := func(x1 uint8, x2 uint16, x3 uint32, x4 uint64, x5 [16]byte) bool {
				Expect(pack.NewU256FromU8(pack.NewU8(x1)).String()).To(Equal(pack.NewU8(x1).String()))
				Expect(pack.NewU256FromU16(pack.NewU16(x2)).String()).To(Equal(pack.NewU16(x2).String()))
				Expect(pack.NewU256FromU32(pack.NewU32(x3)).String()).To(Equal(pack.NewU32(x3).String()))
				Expect(pack.NewU256FromU64(pack.NewU64(x4)).String()).To(Equal(pack.NewU64(x4).String()))
				Expect(pack.NewU256FromU128(pack.NewU128(x5)).String()).To(Equal(pack.NewU128(x5).String()))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when creating negative uint128 ints", func() {
		It("should panic", func() {
			Expect(func() { pack.NewU128FromInt(big.NewInt(-1)) }).To(Panic())
		})
	})

	Context("when creating negative uint256 ints", func() {
		It("should panic", func() {
			Expect(func() { pack.NewU256FromInt(big.NewInt(-1)) }).To(Panic())
		})
	})

	Context("when getting the underlying uint8", func() {
		It("should return the underlying uint8", func() {
			f := func(x uint8) bool {
				Expect(pack.NewU8(x).Uint8()).To(Equal(x))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when getting the underlying uint16", func() {
		It("should return the underlying uint16", func() {
			f := func(x uint16) bool {
				Expect(pack.NewU16(x).Uint16()).To(Equal(x))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when getting the underlying uint32", func() {
		It("should return the underlying uint32", func() {
			f := func(x uint32) bool {
				Expect(pack.NewU32(x).Uint32()).To(Equal(x))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when getting the underlying uint64", func() {
		It("should return the underlying uint64", func() {
			f := func(x uint64) bool {
				Expect(pack.NewU64(x).Uint64()).To(Equal(x))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when getting the underlying uint128", func() {
		It("should return the underlying uint128", func() {
			f := func(x [16]byte) bool {
				Expect(pack.NewU128(x).Bytes16()).To(Equal(x))
				Expect(pack.NewU128(x).Int().Cmp(new(big.Int).SetBytes(x[:]))).To(Equal(0))
				Expect(bytes.Equal(pack.NewU128(x).Bytes(), x[:])).To(BeTrue())
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when getting the underlying uint256", func() {
		It("should return the underlying uint256", func() {
			f := func(x [32]byte) bool {
				Expect(pack.NewU256(x).Bytes32()).To(Equal(x))
				Expect(pack.NewU256(x).Int().Cmp(new(big.Int).SetBytes(x[:]))).To(Equal(0))
				Expect(bytes.Equal(pack.NewU256(x).Bytes(), x[:])).To(BeTrue())
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})
})
