package main_test

var _ = Describe("Event", func() {
	Describe("NewCreateAccountEvent", func() {
		It("can create a create account event", func() {
			name := "John Smith"

			event := NewCreateAccountEvent(name)

			Expect(event.AccName).To(Equal(name))
			Expect(event.AccId).NotTo(BeNil())
			Expect(event.Type).To(Equal("CreateEvent"))
		})
	})
})
