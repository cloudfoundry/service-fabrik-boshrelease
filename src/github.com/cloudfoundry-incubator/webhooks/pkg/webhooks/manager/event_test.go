package main

import (
	"encoding/json"
	"io/ioutil"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/api/admission/v1beta1"
)

var _ = Describe("Event", func() {
	var (
		ar v1beta1.AdmissionReview
	)
	dat, err := ioutil.ReadFile("test_resources/admission_request.json")
	if err != nil {
		panic(err)
	}

	BeforeEach(func() {
		err = json.Unmarshal(dat, &ar)
		if err != nil {
			panic(err)
		}
	})

	Describe("NewEvent", func() {
		It("Should create a new Event object", func() {
			evt, err := NewEvent(&ar)
			Expect(evt).ToNot(Equal(nil))
			Expect(err).To(BeNil())
		})
	})
	Describe("isMeteringEvent", func() {
		Context("Create", func() {
			It("Should should return true if create succeeds", func() {
				evt, _ := NewEvent(&ar)
				Expect(evt.isMeteringEvent()).To(Equal(true))
			})
		})
	})
	Describe("getClient", func() {
		Context("with the passed config", func() {
			It("Should should return a valid client", func() {
				Expect(getClient(tcfg)).ToNot(Equal(nil))
			})
		})
	})
	Describe("createMertering", func() {
		Context("with the passed config", func() {
			It("Should return no kind match error", func() {
				evt, _ := NewEvent(&ar)
				err := evt.createMertering(tcfg)
				Expect(err).To(Equal(&meta.NoKindMatchError{
					GroupKind: schema.GroupKind{
						Group: "metering.servicefabrik.io",
						Kind:  "Event",
					},
					SearchedVersions: []string{"v1alpha1"},
				}))
			})
		})
	})

})
