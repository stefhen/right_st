// The MIT License (MIT)
//
// Copyright (c) 2015 Douglas Thrift
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main_test

import (
	. "github.com/rightscale/right_st"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Account", func() {
	var account = Account{
		Id:           54321,
		Host:         "localhost",
		RefreshToken: "def1234567890abcdef1234567890abcdef12345",
	}

	It("Gets an API 1.5 client singleton", func() {
		firstClient, err := account.Client15()
		Expect(err).NotTo(HaveOccurred())
		Expect(firstClient).NotTo(BeNil())

		secondClient, err := account.Client15()
		Expect(err).NotTo(HaveOccurred())
		Expect(secondClient).To(BeIdenticalTo(firstClient))
	})

	It("Gets an API 1.6 client singleton", func() {
		firstClient, err := account.Client16()
		Expect(err).NotTo(HaveOccurred())
		Expect(firstClient).NotTo(BeNil())

		secondClient, err := account.Client16()
		Expect(err).NotTo(HaveOccurred())
		Expect(secondClient).To(BeIdenticalTo(firstClient))
	})

	Context("With an invalid host", func() {
		var invalidHostAccount = Account{
			Id:           54321,
			Host:         "localhost/api/oauth2",
			RefreshToken: "def1234567890abcdef1234567890abcdef12345",
		}

		It("Returns an error for API 1.5", func() {
			client, err := invalidHostAccount.Client15()
			Expect(err).To(MatchError(MatchRegexp(`^Invalid host name for account \(host: .+, id: .+\): .+$`)))
			Expect(client).To(BeNil())
		})

		It("Returns an error for API 1.6", func() {
			client, err := invalidHostAccount.Client16()
			Expect(err).To(MatchError(MatchRegexp(`^Invalid host name for account \(host: .+, id: .+\): .+$`)))
			Expect(client).To(BeNil())
		})
	})
})
