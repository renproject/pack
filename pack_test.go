package pack_test

import (
	"github.com/renproject/pack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Value", func() {
	Context("when compiling", func() {
		It("should check that all types implement the value interface", func() {
			Expect(func() {
				func(v pack.Value) {}(new(pack.Bool))
				func(v pack.Value) {}(new(pack.U8))
				func(v pack.Value) {}(new(pack.U16))
				func(v pack.Value) {}(new(pack.U32))
				func(v pack.Value) {}(new(pack.U64))
				func(v pack.Value) {}(new(pack.U128))
				func(v pack.Value) {}(new(pack.U256))
				func(v pack.Value) {}(new(pack.String))
				func(v pack.Value) {}(new(pack.Bytes))
				func(v pack.Value) {}(new(pack.Bytes32))
				func(v pack.Value) {}(new(pack.Bytes65))
				func(v pack.Value) {}(new(pack.Struct))
			}).ToNot(Panic())
		})
	})
})
