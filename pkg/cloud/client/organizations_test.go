package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ClientMock struct {
	body []byte
	err  error
}

func (c ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewReader(c.body)),
	}, c.err
}

func TestGet(t *testing.T) {

	t.Run("error on invalid token", func(t *testing.T) {
		o := NewOrganizationsClient("token")
		o.RESTClient.Client = ClientMock{err: http.ErrNoLocation}
		_, err := o.List()
		assert.Error(t, err)
	})

	t.Run("get orgs on invalid token", func(t *testing.T) {
		o := NewOrganizationsClient("token")
		o.RESTClient.Client = ClientMock{body: []byte(`{"elements":[{"id":"tkcorg_2bb1486705fb6997","name":"JW IT Consulting","createdAt":"2023-01-03T08:49:48Z","updatedAt":"2023-04-16T07:53:44Z","plan":{"id":"","name":"plan_pro","status":"Active","customerId":"cus_NXG4LzlhoTQMvn","checkoutSessionId":"cs_live_b1WIb7L7cZcDDr9Eb7IL1NhkAKvMh4XBBzd2CO8PX5eVglzVWX6I0NJ1jR","subscriptionId":"sub_1MmBd2LQdtI6cbSfxXzOxfRH","currentPeriodStart":"2023-04-16T07:53:40Z","currentPeriodEnd":"2023-05-16T07:53:40Z","usageRecordTime":"0001-01-01T00:00:00Z","limits":{"maxEnvironments":-1,"maxUsers":25,"maxTestResults":5000,"maxStorageGb":128}},"members":[{"email":"jacek.wysocki@gmail.com","status":"Accepted","role":"owner"},{"email":"jacek.wysocki+1@gmail.com","status":"Revoked","role":"admin","inviteID":"588733b8-5cd3-440f-9d8d-f685fe82b6f4","invitingEmail":"jacek.wysocki@gmail.com"},{"email":"jacek@kubeshop.io","status":"Accepted","role":"member","inviteID":"b8a250a3-cdcd-4a9e-8ab5-5ad74926ec63","invitingEmail":"jacek.wysocki@gmail.com"},{"email":"povilas.versockas@kubeshop.io","status":"Accepted","role":"member","inviteID":"1d9a9ed8-f087-4521-893e-59816d486ad4","invitingEmail":"jacek.wysocki@gmail.com"},{"email":"edward@kubeshop.io","status":"Revoked","role":"admin","inviteID":"d855dd3d-97ba-4183-aaaa-2c5f4b1d5094","invitingEmail":"jacek.wysocki@gmail.com"},{"email":"edward@kubeshop.io","status":"Revoked","role":"admin","inviteID":"c326ff62-e878-49a9-aad0-089f85f55390","invitingEmail":"jacek.wysocki@gmail.com"},{"email":"jacek+33@kubeshop.io","status":"Accepted","role":"biller","inviteID":"69b37629-f66e-4d14-8622-d1673ddd0893","invitingEmail":"jacek.wysocki@gmail.com"},{"email":"jacek.wysocki+123@gmail.com","status":"Revoked","role":"member","inviteID":"ab3e457c-f60c-40e4-9e2d-a7f1c9d23bd1","invitingEmail":"jacek+33@kubeshop.io"},{"email":"jacek.wysocki+123@gmail.com","status":"Revoked","role":"member","inviteID":"b6cd8c63-a42c-4f44-ae9a-9af90add1324","invitingEmail":"jacek+33@kubeshop.io"},{"email":"jacek.wysocki+123123@gmail.com","status":"Revoked","role":"member","inviteID":"ee799d4f-e22e-4954-8c28-82bbc5b60faa","invitingEmail":"jacek+33@kubeshop.io"},{"email":"jacek+1@kubeshop.io","status":"Revoked","role":"member","inviteID":"040086ed-ab9e-41ca-ab81-40995789f5dd","invitingEmail":"jacek+33@kubeshop.io"},{"email":"jacek+12334@kubeshop.io","status":"Revoked","role":"member","inviteID":"36e747b0-bf05-421f-9e32-22f222f69399","invitingEmail":"jacek+33@kubeshop.io"},{"email":"bogdan@kubeshop.io","status":"Accepted","role":"member","inviteID":"7ed8d7ec-fa15-4c3e-8087-ce6dd95d18a6","invitingEmail":"jacek+33@kubeshop.io"},{"email":"magazyn@kasia.in","status":"Accepted","role":"member","inviteID":"e4622347-8ffb-4cc2-ac4f-723e4c57b1ab","invitingEmail":"jacek+33@kubeshop.io"}]},{"id":"tkcorg_67750603c10b5438","name":"Testkube","createdAt":"2023-01-05T11:36:29Z","updatedAt":"2023-04-24T09:10:16Z","plan":{"id":"","name":"plan_free","status":"","customerId":"","checkoutSessionId":"","subscriptionId":"","currentPeriodStart":"0001-01-01T00:00:00Z","currentPeriodEnd":"0001-01-01T00:00:00Z","usageRecordTime":"0001-01-01T00:00:00Z","limits":{"maxEnvironments":2,"maxUsers":33,"maxTestResults":600,"maxStorageGb":2}},"members":[{"email":"mh@fivenp.com","status":"Accepted","role":"owner"},{"email":"tomasz.konieczny@kubeshop.io","status":"Accepted","role":"owner","inviteID":"88a379b4-0815-426b-9d96-2bbac6ac95ba","invitingEmail":"mh@fivenp.com"},{"email":"tkonieczny91@gmail.com","status":"Accepted","role":"member","inviteID":"751f1859-f35b-4aa9-a1a1-11986e1074cf","invitingEmail":"mh@fivenp.com"},{"email":"p.versockas@gmail.com","status":"Accepted","role":"admin","inviteID":"b4297457-cc5f-4944-affe-0bed5d08f6ed","invitingEmail":"tkonieczny91@gmail.com"},{"email":"bogdan@kubeshop.io","status":"Invited","role":"admin","inviteID":"9e6a6daa-4a61-4636-a584-9043956fc020","invitingEmail":"tkonieczny91@gmail.com"},{"email":"lilla@kubeshop.io","status":"Accepted","role":"owner","inviteID":"2a06c304-25c5-44eb-9492-c974aad7d22f","invitingEmail":"tkonieczny91@gmail.com"},{"email":"alejandra@kubeshop.io","status":"Accepted","role":"member","inviteID":"8345acc1-bb90-4c28-85e8-d9990881bcaa","invitingEmail":"tkonieczny91@gmail.com"},{"email":"dejan@kubeshop.io","status":"Revoked","role":"admin","inviteID":"469dd379-2c8f-4768-a36d-07c937d84898","invitingEmail":"tkonieczny91@gmail.com"},{"email":"jacek@kubeshop.io","status":"Accepted","role":"owner","inviteID":"dad48d10-2d77-4302-bb2a-f16b00f4197c","invitingEmail":"tkonieczny91@gmail.com"},{"email":"vladislav@kubeshop.io","status":"Accepted","role":"admin","inviteID":"7d46a379-3da4-4c69-a505-c83db1c1383e","invitingEmail":"tkonieczny91@gmail.com"},{"email":"nicolae@kubeshop.io","status":"Invited","role":"admin","inviteID":"85a47004-5937-4cca-9936-05f8089e8401","invitingEmail":"tkonieczny91@gmail.com"},{"email":"yulia.poplavska@kubeshop.io","status":"Accepted","role":"admin","inviteID":"e144f98c-8d56-495d-9f34-fd303c5cc625","invitingEmail":"tkonieczny91@gmail.com"},{"email":"bogdan.hanea@yahoo.com","status":"Accepted","role":"admin","inviteID":"bb1a52ba-bcf6-4789-8aa3-675411d9e435","invitingEmail":"tkonieczny91@gmail.com"},{"email":"alejandralupio.identifix@gmail.com","status":"Invited","role":"admin","inviteID":"0f77fa39-d15e-4b50-ac37-3a7e9f6e2e61","invitingEmail":"tkonieczny91@gmail.com"},{"email":"edward@kubeshop.io","status":"Accepted","role":"admin","inviteID":"eff8b8ff-0f5a-4f47-a5d7-1171b8bd0bbd","invitingEmail":"mh@fivenp.com"},{"email":"Abdallah@kubeshop.io","status":"Accepted","role":"admin","inviteID":"215cc5c0-564a-4a6c-8093-a8f0a579d20f","invitingEmail":"tkonieczny91@gmail.com"},{"email":"bruno@kubeshop.io","status":"Accepted","role":"owner","inviteID":"d55691eb-9ff8-44c7-ab13-a30c6565b4f9","invitingEmail":"jacek+33@kubeshop.io"},{"email":"dawid@kubeshop.io","status":"Accepted","role":"admin","inviteID":"868f55c4-aad0-413b-889a-b15a00b2649d","invitingEmail":"bogdan@kubeshop.io"},{"email":"dejan@kubeshop.io","status":"Accepted","role":"admin","inviteID":"830eb9f6-544d-4c1d-9ac9-7b770e715137","invitingEmail":"dawid@kubeshop.io"}]},{"id":"tkcorg_a114ae9ec19c4a8b","name":"Jacek Wysocki-personal-org","description":"default","labels":{"default":"true"},"createdAt":"2023-01-05T12:13:42Z","plan":{"id":"","name":"plan_free","status":"","customerId":"","checkoutSessionId":"","subscriptionId":"","currentPeriodStart":"0001-01-01T00:00:00Z","currentPeriodEnd":"0001-01-01T00:00:00Z","usageRecordTime":"0001-01-01T00:00:00Z","limits":{"maxEnvironments":2,"maxUsers":3,"maxTestResults":600,"maxStorageGb":2}},"members":[{"email":"jacek@kubeshop.io","status":"Accepted","role":"owner"}]}]}`)}
		orgs, err := o.List()
		assert.NoError(t, err)
		assert.Len(t, orgs, 3)
	})
}
