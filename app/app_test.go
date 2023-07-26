package app

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rohit-tambe/mockgen-example/mocks"
	"github.com/rohit-tambe/mockgen-example/party"
)

func TestPartyService_GreetVisitors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedVisitorLister := mocks.NewMockVisitorLister(ctrl)

	mockedVisitorLister.EXPECT().ListVisitors(party.NiceVisitor).Return([]party.Visitor{{Name: "Peter", Surname: "TheSmart"}}, nil)
	// mockedVisitorLister.EXPECT().ListVisitors(party.NotNiceVisitor).Return(nil, fmt.Errorf("dummyError"))

	s := PartyService{visitorLister: mockedVisitorLister}
	gotErr := s.GreetVisitors(true)

	if gotErr == nil {
		t.Errorf("did not get an error")
	}
}
