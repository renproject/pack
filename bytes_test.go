package pack_test

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"testing/quick"

	"github.com/renproject/pack"
	"github.com/renproject/pack/packutil"
	"github.com/renproject/surge/surgeutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bytes", func() {

	numTrials := 10

	ts := []reflect.Type{
		reflect.TypeOf(pack.NewString("")),
		reflect.TypeOf(pack.NewBytes([]byte{})),
		reflect.TypeOf(pack.NewBytes32([32]byte{})),
		reflect.TypeOf(pack.NewBytes65([65]byte{})),
		reflect.TypeOf(pack.NewBytes64([64]byte{})),
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
	}

	Context("when unmarshaling a byte32 array", func() {
		Context("when the string represents an array with a different length", func() {
			It("should return an error", func() {
				f := func(x [31]byte) bool {
					data, err := json.Marshal(pack.NewBytes(x[:]))
					Expect(err).ToNot(HaveOccurred())
					b32 := pack.Bytes32{}
					err = json.Unmarshal(data, &b32)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("expected len=32"))
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})
	})

	Context("when unmarshaling a byte65 array", func() {
		Context("when the string represents an array with a different length", func() {
			It("should return an error", func() {
				f := func(x [64]byte) bool {
					data, err := json.Marshal(pack.NewBytes(x[:]))
					Expect(err).ToNot(HaveOccurred())
					b65 := pack.Bytes65{}
					err = json.Unmarshal(data, &b65)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("expected len=65"))
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})
	})

	Context("when unmarshaling a byte64 array", func() {
		Context("when the string represents an array with a different length", func() {
			It("should return an error", func() {
				f := func(x [63]byte) bool {
					data, err := json.Marshal(pack.NewBytes(x[:]))
					Expect(err).ToNot(HaveOccurred())
					b64 := pack.Bytes64{}
					err = json.Unmarshal(data, &b64)
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("expected len=64"))
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})
	})

	Context("when creating a nil byte slice", func() {
		It("should return empty bytes", func() {
			Expect([]byte(pack.NewBytes(nil))).ToNot(BeNil())
			Expect([]byte(pack.NewBytes(nil))).To(HaveLen(0))
		})
	})

	Context("when comparing strings", func() {
		Context("when the strings are equal", func() {
			It("should return true", func() {
				f := func(x string) bool {
					Expect(pack.NewString(x).Equal(pack.NewString(x))).To(BeTrue())
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})

		Context("when the strings are not equal", func() {
			It("should return true", func() {
				f := func(x, y string) bool {
					if x == y {
						return true
					}
					Expect(pack.NewString(x).Equal(pack.NewString(y))).To(BeFalse())
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})
	})

	Context("when comparing byte slices", func() {
		Context("when the byte slices are equal", func() {
			It("should return true", func() {
				f := func(x []byte) bool {
					Expect(pack.NewBytes(x).Equal(pack.NewBytes(x))).To(BeTrue())
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})

		Context("when the byte slices are not equal", func() {
			It("should return true", func() {
				f := func(x, y []byte) bool {
					if bytes.Equal(x, y) {
						return true
					}
					Expect(pack.NewBytes(x).Equal(pack.NewBytes(y))).To(BeFalse())
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})
	})

	Context("when comparing byte32 arrays", func() {
		Context("when the byte32 arrays are equal", func() {
			It("should return true", func() {
				f := func(x [32]byte) bool {
					other := pack.NewBytes32(x)
					Expect(pack.NewBytes32(x).Equal(&other)).To(BeTrue())
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})

		Context("when the byte32 arrays are not equal", func() {
			It("should return true", func() {
				f := func(x, y [32]byte) bool {
					if bytes.Equal(x[:], y[:]) {
						return true
					}
					other := pack.NewBytes32(y)
					Expect(pack.NewBytes32(x).Equal(&other)).To(BeFalse())
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})
	})

	Context("when comparing byte65 arrays", func() {
		Context("when the byte65 arrays are equal", func() {
			It("should return true", func() {
				f := func(x [65]byte) bool {
					other := pack.NewBytes65(x)
					Expect(pack.NewBytes65(x).Equal(&other)).To(BeTrue())
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})

		Context("when the byte65 arrays are not equal", func() {
			It("should return true", func() {
				f := func(x, y [65]byte) bool {
					if bytes.Equal(x[:], y[:]) {
						return true
					}
					other := pack.NewBytes65(y)
					Expect(pack.NewBytes65(x).Equal(&other)).To(BeFalse())
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})
	})

	Context("when comparing byte64 arrays", func() {
		Context("when the byte64 arrays are equal", func() {
			It("should return true", func() {
				f := func(x [64]byte) bool {
					other := pack.NewBytes64(x)
					Expect(pack.NewBytes64(x).Equal(&other)).To(BeTrue())
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})

		Context("when the byte64 arrays are not equal", func() {
			It("should return true", func() {
				f := func(x, y [64]byte) bool {
					if bytes.Equal(x[:], y[:]) {
						return true
					}
					other := pack.NewBytes64(y)
					Expect(pack.NewBytes64(x).Equal(&other)).To(BeFalse())
					return true
				}
				Expect(quick.Check(f, nil)).To(Succeed())
			})
		})
	})

	Context("when marshaling and unmarshaling byte slices to and from text", func() {
		It("should return itself", func() {
			f := func(x pack.Bytes) bool {
				data, err := x.MarshalText()
				Expect(err).ToNot(HaveOccurred())
				y := pack.Bytes{}
				err = y.UnmarshalText(data)
				Expect(err).ToNot(HaveOccurred())
				return x.Equal(y)
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when marshaling and unmarshaling byte32 arrays to and from text", func() {
		It("should return itself", func() {
			f := func(x pack.Bytes32) bool {
				data, err := x.MarshalText()
				Expect(err).ToNot(HaveOccurred())
				y := pack.Bytes32{}
				err = y.UnmarshalText(data)
				Expect(err).ToNot(HaveOccurred())
				return x.Equal(&y)
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when marshaling and unmarshaling byte65 arrays to and from text", func() {
		It("should return itself", func() {
			f := func(x pack.Bytes65) bool {
				data, err := x.MarshalText()
				Expect(err).ToNot(HaveOccurred())
				y := pack.Bytes65{}
				err = y.UnmarshalText(data)
				Expect(err).ToNot(HaveOccurred())
				return x.Equal(&y)
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when marshaling and unmarshaling byte64 arrays to and from text", func() {
		It("should return itself", func() {
			f := func(x pack.Bytes64) bool {
				data, err := x.MarshalText()
				Expect(err).ToNot(HaveOccurred())
				y := pack.Bytes64{}
				err = y.UnmarshalText(data)
				Expect(err).ToNot(HaveOccurred())
				return x.Equal(&y)
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when stringifying strings", func() {
		It("should return itself", func() {
			f := func(x string) bool {
				Expect(pack.NewString(x).String()).To(Equal(x))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when stringifying byte slices", func() {
		It("should return raw URL base64 encodings", func() {
			f := func(x []byte) bool {
				Expect(pack.NewBytes(x).String()).To(Equal(base64.RawURLEncoding.EncodeToString(x)))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when stringifying byte32 arrays", func() {
		It("should return raw URL base64 encodings", func() {
			f := func(x [32]byte) bool {
				Expect(pack.NewBytes32(x).String()).To(Equal(base64.RawURLEncoding.EncodeToString(x[:])))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when getting bytes from a byte32 array", func() {
		It("should a copy of the underlying bytes", func() {
			f := func(x [32]byte) bool {
				Expect(bytes.Equal(pack.NewBytes32(x).Bytes(), x[:])).To(BeTrue())
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when stringifying byte65 slices", func() {
		It("should return raw URL base64 encodings", func() {
			f := func(x [65]byte) bool {
				Expect(pack.NewBytes65(x).String()).To(Equal(base64.RawURLEncoding.EncodeToString(x[:])))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when getting bytes from a byte65 array", func() {
		It("should a copy of the underlying bytes", func() {
			f := func(x [65]byte) bool {
				Expect(bytes.Equal(pack.NewBytes65(x).Bytes(), x[:])).To(BeTrue())
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when stringifying byte64 slices", func() {
		It("should return raw URL base64 encodings", func() {
			f := func(x [64]byte) bool {
				Expect(pack.NewBytes64(x).String()).To(Equal(base64.RawURLEncoding.EncodeToString(x[:])))
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when getting bytes from a byte64 array", func() {
		It("should a copy of the underlying bytes", func() {
			f := func(x [64]byte) bool {
				Expect(bytes.Equal(pack.NewBytes64(x).Bytes(), x[:])).To(BeTrue())
				return true
			}
			Expect(quick.Check(f, nil)).To(Succeed())
		})
	})

	Context("when getting type information for strings", func() {
		It("should return the string type", func() {
			Expect(pack.NewString("").Type().Kind()).To(Equal(pack.KindString))
		})
	})

	Context("when getting type information for byte slices", func() {
		It("should return the bytes type", func() {
			Expect(pack.NewBytes(nil).Type().Kind()).To(Equal(pack.KindBytes))
		})
	})

	Context("when getting type information for byte32 arrays", func() {
		It("should return the 32-byte array type", func() {
			Expect(pack.NewBytes32([32]byte{}).Type().Kind()).To(Equal(pack.KindBytes32))
		})
	})

	Context("when getting type information for byte65 arrays", func() {
		It("should return the 65-byte array type", func() {
			Expect(pack.NewBytes65([65]byte{}).Type().Kind()).To(Equal(pack.KindBytes65))
		})
	})

	Context("when getting type information for byte64 arrays", func() {
		It("should return the 64-byte array type", func() {
			Expect(pack.NewBytes64([64]byte{}).Type().Kind()).To(Equal(pack.KindBytes64))
		})
	})
})
