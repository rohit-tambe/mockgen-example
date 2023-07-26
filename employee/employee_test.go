package employee

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rohit-tambe/mockgen-example/mocks"
	"github.com/stretchr/testify/assert"
)

func TestEmployee_SalaryCalaculation(t *testing.T) {
	t.Run("Salary Addition", func(t *testing.T) {
		mockCotroller := gomock.NewController(t)
		defer mockCotroller.Finish()

		mockCalculator := mocks.NewMockCalculator(mockCotroller)
		mockCalculator.EXPECT().Add(1, 2).Return(3, nil).Times(1)
		testEmployee := New(mockCalculator)
		result := testEmployee.Add(1, 2)
		assert.Equal(t, 3, result)
	})
}

// func TestEmployee_SalaryMultiplication(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		e       *Employee
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := tt.e.SalaryMultiplication(); (err != nil) != tt.wantErr {
// 				t.Errorf("Employee.SalaryMultiplication() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
