package main_test

import (
	"io"
	"strings"

	. "github.com/rightscale/right_st"

	"github.com/go-yaml/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServerTemplate", func() {
	var script io.Reader

	Describe("Parse", func() {
		Context("With valid YAML", func() {
			script = strings.NewReader(`---
Name: Test ST
Description: Test ST Description
RightScripts:
  Boot:
    - Dummy.sh
  Operational:
    - Dummy2.sh
  Decommission:
    - Dummy3.sh
Inputs:
  SERVER_HOSTNAME: text:test.local
MultiCloudImages:
  - Href: /api/multi_cloud_images/403042003
  - Name: FooImage
    Revision: 100
`)
			It("should parse correctly", func() {
				st, err := ParseServerTemplate(script)
				Expect(err).To(Succeed())
				Expect(st).NotTo(BeNil())
				Expect(st.Name).To(Equal("Test ST"))
				Expect(st.Description).To(Equal("Test ST Description"))
				Expect(st.Inputs).To(Equal(map[string]*InputValue{"SERVER_HOSTNAME": &InputValue{Type: "text", Value: "test.local"}}))
				Expect(st.RightScripts["Boot"]).To(Equal([]string{"Dummy.sh"}))
				Expect(st.RightScripts["Operational"]).To(Equal([]string{"Dummy2.sh"}))
				Expect(st.RightScripts["Decommission"]).To(Equal([]string{"Dummy3.sh"}))
				Expect(st.MultiCloudImages).To(Equal([]*MultiCloudImage{
					&MultiCloudImage{Href: "/api/multi_cloud_images/403042003"},
					&MultiCloudImage{Name: "FooImage", Revision: 100},
				}))
			})
		})

		Context("With invalid structure in YAML", func() {
			It("should return an error", func() {
				script := strings.NewReader(`---
Name: Test ST
Description: Test ST Description
RightScripts:
  Boot:
    - Dummy.sh
Inputs:
  - TEXT_INPUT
# The Inputs field should have a map not an array
`)
				_, err := ParseServerTemplate(script)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(&yaml.TypeError{
					Errors: []string{
						"line 8: cannot unmarshal !!seq into map[string]*main.InputValue",
					},
				}))
			})
		})

		Context("With invalid RightScript type in YAML", func() {
			It("should return an error", func() {
				script := strings.NewReader(`---
Name: Test ST
Description: Test ST Description
RightScripts:
  Unknown:
    - Dummy.sh
Inputs: {}
`)
				_, err := ParseServerTemplate(script)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("not a valid sequence name"))
			})
		})

		Context("With an unknown field in YAML", func() {
			It("should return an error", func() {
				script := strings.NewReader(`---
Name: Test ST
Description: Test ST Description
Some Bogus Field: Some bogus value
`)
				_, err := ParseServerTemplate(script)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(&yaml.TypeError{
					Errors: []string{
						"line 2: no such field 'Some Bogus Field' in struct 'main.ServerTemplate'",
					},
				}))
			})
		})
	})
})
