package pack_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"

	"github.com/renproject/pack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encoding", func() {

	ts := []reflect.Type{
		// Pack types.
		reflect.TypeOf(pack.Bool(false)),
		reflect.TypeOf(pack.U8(0)),
		reflect.TypeOf(pack.U16(0)),
		reflect.TypeOf(pack.U32(0)),
		reflect.TypeOf(pack.U64(0)),
		reflect.TypeOf(pack.NewU128([16]byte{})),
		reflect.TypeOf(pack.NewU256([32]byte{})),
		reflect.TypeOf(pack.String("")),
		reflect.TypeOf(pack.Bytes{}),
		reflect.TypeOf(pack.Bytes32{}),
		reflect.TypeOf(pack.Bytes65{}),
		reflect.TypeOf(pack.Struct{}),
		reflect.TypeOf(pack.List{}),

		// Standard types.
		reflect.TypeOf(false),
		reflect.TypeOf(uint8(0)),
		reflect.TypeOf(uint16(0)),
		reflect.TypeOf(uint32(0)),
		reflect.TypeOf(uint64(0)),
		reflect.TypeOf(""),
		reflect.TypeOf([]byte{}),
		reflect.TypeOf([32]byte{}),
		reflect.TypeOf([65]byte{}),
		reflect.TypeOf(struct{}{}),
		reflect.TypeOf([]string{}),
		reflect.TypeOf([]uint64{}),
		reflect.TypeOf(struct {
			X       uint8  `json:"x"`
			Y       uint16 `json:"y"`
			Omit    uint32 `json:"z,omitempty"`
			Dash    uint64 `json:"-"`
			Unnamed uint64

			Foo string   `json:"foo"`
			Bar []byte   `json:"bar"`
			Baz [32]byte `json:"baz"`
			Boo [65]byte `json:"boo"`

			Inner struct {
				InnerX       uint8  `json:"x"`
				InnerY       uint16 `json:"y"`
				InnerOmit    uint32 `json:"z,omitempty"`
				InnerDash    uint64 `json:"-"`
				InnerUnnamed uint64
			} `json:"inner"`

			ListOfStrings []string `json:"listOfStrings"`
			ListOfUints   []uint64 `json:"listOfUints"`
		}{}),

		// Mixed types.
		reflect.TypeOf(struct {
			X       pack.U64 `json:"x"`
			Y       pack.U64 `json:"y"`
			Omit    pack.U64 `json:"z,omitempty"`
			Dash    pack.U64 `json:"-"`
			Unnamed pack.U64

			Foo pack.String  `json:"foo"`
			Bar pack.Bytes   `json:"bar"`
			Baz pack.Bytes32 `json:"baz"`
			Boo pack.Bytes65 `json:"boo"`

			Inner struct {
				InnerX       pack.U64 `json:"x"`
				InnerY       pack.U64 `json:"y"`
				InnerOmit    pack.U64 `json:"z,omitempty"`
				InnerDash    pack.U64 `json:"-"`
				InnerUnnamed pack.U64
			} `json:"inner"`

			List pack.List `json:"list"`
		}{}),
	}
	numTrials := 100

	for _, t := range ts {
		t := t

		Context(fmt.Sprintf("when encoding and then decoding %v", t), func() {
			It("should equal itself", func() {
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				for trial := 0; trial < numTrials; trial++ {
					x, ok := quick.Value(t, r)
					Expect(ok).To(BeTrue())

					v, err := pack.Encode(x.Interface())
					Expect(err).ToNot(HaveOccurred())

					y := reflect.New(t)
					err = pack.Decode(y.Interface(), v)
					Expect(err).ToNot(HaveOccurred())

					w, err := pack.Encode(y.Elem().Interface())
					Expect(err).ToNot(HaveOccurred())

					Expect(reflect.DeepEqual(v, w)).To(BeTrue())
				}
			})
		})
	}

	Context("when encoding and then decoding a value", func() {
		It("should equal itself", func() {
			r := rand.New(rand.NewSource(GinkgoRandomSeed()))

			var x pack.Value
			x = pack.Generate(r, 1, true, true).Interface().(pack.Value)
			encoded, err := pack.Encode(x)
			Expect(err).ToNot(HaveOccurred())

			var y pack.Value
			err = pack.Decode(&y, encoded)
			Expect(err).ToNot(HaveOccurred())

			Expect(reflect.DeepEqual(x, y)).To(BeTrue())
		})
	})

	type PartialStruct struct {
		Foo pack.U64 `json:"foo"`
	}

	type A struct {
		Foo pack.U64    `json:"foo"`
		Bar pack.String `json:"bar"`
	}

	type B struct {
		Foo pack.U64   `json:"foo"`
		Baz pack.Bytes `json:"baz"`
	}

	type C struct {
		Foo pack.U64  `json:"foo"`
		Boo pack.List `json:"boo"`
	}

	Context("when decoding into a struct with new fields", func() {
		It("should not error", func() {
			v, err := pack.Encode(PartialStruct{})
			Expect(err).ToNot(HaveOccurred())
			var x A
			err = pack.Decode(&x, v)
			Expect(err).ToNot(HaveOccurred())
			var y B
			err = pack.Decode(&y, v)
			Expect(err).ToNot(HaveOccurred())
			var z C
			err = pack.Decode(&z, v)
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
