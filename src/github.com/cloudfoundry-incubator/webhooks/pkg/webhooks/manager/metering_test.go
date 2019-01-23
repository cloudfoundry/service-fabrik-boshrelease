package main

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Metering", func() {
	Describe("newMetering", func() {
		It("it should create the metering object", func() {
			//Create params
			co := ContextOptions{
				Platform:         "test-platform",
				OrganizationGUID: "test-org-guid",
				SpaceGUID:        "test-space",
			}
			opt := GenericOptions{
				ServiceID: "test-service-id",
				PlanID:    "test-plan-id",
				Context:   co,
			}
			crd := GenericResource{}
			signal := MeterStop
			// Test creating metering object
			m := newMetering(opt, crd, signal)
			var unmarsheledMeteringOptions MeteringOptions
			json.Unmarshal([]byte(m.Spec.Options), &unmarsheledMeteringOptions)
			Expect(unmarsheledMeteringOptions.ID).Should(MatchRegexp("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"))
			// Expect(unmarsheledMeteringOptions.Timestamp).To(Equal(opt.ServiceID))
			Expect(unmarsheledMeteringOptions.ServiceInfo.ID).To(Equal(opt.ServiceID))
			Expect(unmarsheledMeteringOptions.ServiceInfo.Plan).To(Equal(opt.PlanID))
			Expect(unmarsheledMeteringOptions.ConsumerInfo.Environment).To(Equal(""))
			Expect(unmarsheledMeteringOptions.ConsumerInfo.Region).To(Equal(""))
			Expect(unmarsheledMeteringOptions.ConsumerInfo.Org).To(Equal(opt.Context.OrganizationGUID))
			Expect(unmarsheledMeteringOptions.ConsumerInfo.Space).To(Equal(opt.Context.SpaceGUID))
			Expect(unmarsheledMeteringOptions.ConsumerInfo.Instance).To(Equal(crd.Name))
			Expect(unmarsheledMeteringOptions.InstancesMeasures[0].ID).To(Equal("instances"))
			Expect(unmarsheledMeteringOptions.InstancesMeasures[0].Value).To(Equal(MeterStop))
		})
	})
})
