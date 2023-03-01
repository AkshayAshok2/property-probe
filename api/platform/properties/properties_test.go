package properties

import "testing"

func (t *SuiteTest) TestCreateUser() {
	_, err := CreateProperty(Property{
		AuctionType: "AuctionTest",
		JudgementAmount: "AmountTest",
		Address: "AddressTest",
		AssessedValue: "ValueTest"
	})
	assert.NoError(t.T(), err)

	_, err = service.UserCreate(entity.User{
		Name:  "Second",
		Email: "second@gmail.com",
	})
	assert.NoError(t.T(), err)

	_, err = service.UserCreate(entity.User{
		Name:  "Third",
		Email: "second@gmail.com",
	})
	assert.Error(t.T(), err) // Duplicate Email Error
}