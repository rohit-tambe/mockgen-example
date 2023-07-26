package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/rohit-tambe/mockgen-example/mocks/mockPlayer"
	"github.com/stretchr/testify/assert"
)

func TestPlayer_GetName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockU := mocks.NewMockUserI(ctrl)

	gomock.InOrder(
		mockU.EXPECT().GetName("Rohit").Return("Hi...Rohit").Times(1))

	p := NewUser(mockU)
	name := p.GetName("Rohit")
	assert.Equal(t, name, "Hi...Rohit")
}
