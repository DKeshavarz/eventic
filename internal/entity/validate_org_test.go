package entity

import "testing"

func TestValidateOrganization(t *testing.T) {
	testCases := []struct {
		name        string
		org         *Organization
		expectedErr error
	}{
		{
			name: "Valid Organization",
			org: &Organization{
				Name:        "Test Organization",
				Description: "This is a test organization",
				LogoPic:     nil,
				Email:       nil,
				Phone:       nil,
			},
			expectedErr: nil,
		},
		{
			name: "Invalid Organization - no descrption",
			org: &Organization{
				Name:        "Test Organization",
				Description: "",
				LogoPic:     nil,
				Email:       nil,
				Phone:       nil,
			},
			expectedErr: ErrInvalidDescription,
		},
		{
			name: "Invalid Organization - no name",
			org: &Organization{
				Name:        "",
				Description: "Somthing",
				LogoPic:     nil,
				Email:       nil,
				Phone:       nil,
			},
			expectedErr: ErrInvalidName,
		},
		{
			name: "Invalid Organization - bad email",
			org: &Organization{
				Name:        "name",
				Description: "Somthing",
				LogoPic:     nil,
				Email:       strPtr("dankeshavarz@som"),
				Phone:       nil,
			},
			expectedErr: ErrInvalidEmail,
		},
		{
			name: "Invalid Organization - bad phone",
			org: &Organization{
				Name:        "name",
				Description: "Somthing",
				LogoPic:     strPtr("./pic/"),
				Email:       strPtr("dankeshavarz@som.com"),
				Phone:       strPtr("0918 811 3791"),
			},
			expectedErr: ErrInvalidPhone,
		},
		{
			name: "Valid Organization - with all",
			org: &Organization{
				Name:        "name",
				Description: "Somthing",
				LogoPic:     strPtr("./pic/"),
				Email:       strPtr("dankeshavarz@som.com"),
				Phone:       strPtr("09188113791"),
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := (tc.org).Validate()
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}


// ------- helpers ----------------
func strPtr(s string) *string {
	return &s
}