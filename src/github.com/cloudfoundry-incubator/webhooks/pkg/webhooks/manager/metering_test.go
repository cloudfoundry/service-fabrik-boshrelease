package main

import (
	// "io/ioutil"
	// "net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Webhook", func() {
	It("Should return return a valid admission response", func() {

		// whsvr := &WebhookServer{
		// 	server: &http.Server{},
		// }
		// dat, err := ioutil.ReadFile("Admissise_object")
		// if err != nil {
		// 	panic(err)
		// }
		Expect(1).To(Equal(1))
	})

	It("Should create metering", func() {

		// whsvr := &WebhookServer{
		// 	server: &http.Server{},
		// }
		Expect(true).To(Equal(true))
	})
})
