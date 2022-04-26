package role

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshal(t *testing.T) {
	testCases := map[string]struct {
		role     Role
		wantJson string
		wantErr  bool
	}{
		"coordinator role": {
			role:     Coordinator,
			wantJson: `"Coordinator"`,
		},
		"node role": {
			role:     Node,
			wantJson: `"Node"`,
		},
		"admin role": {
			role:     Admin,
			wantJson: `"Admin"`,
		},
		"unknown role": {
			role:     Unknown,
			wantJson: `"Unknown"`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			jsonRole, err := tc.role.MarshalJSON()
			if tc.wantErr {
				assert.Error(err)
				return
			}

			require.NoError(err)
			assert.Equal(tc.wantJson, string(jsonRole))
		})
	}
}

func TestUnmarshal(t *testing.T) {
	testCases := map[string]struct {
		json     string
		wantRole Role
		wantErr  bool
	}{
		"Coordinator can be unmarshaled": {
			json:     `"Coordinator"`,
			wantRole: Coordinator,
		},
		"lowercase coordinator can be unmarshaled": {
			json:     `"coordinator"`,
			wantRole: Coordinator,
		},
		"Node can be unmarshaled": {
			json:     `"Node"`,
			wantRole: Node,
		},
		"lowercase node can be unmarshaled": {
			json:     `"node"`,
			wantRole: Node,
		},
		"Admin can be unmarshaled": {
			json:     `"Admin"`,
			wantRole: Admin,
		},
		"lowercase admin can be unmarshaled": {
			json:     `"admin"`,
			wantRole: Admin,
		},
		"other strings unmarshal to the unknown role": {
			json:     `"anything"`,
			wantRole: Unknown,
		},
		"invalid json fails": {
			json:    `"unterminated string literal`,
			wantErr: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			var role Role
			err := role.UnmarshalJSON([]byte(tc.json))

			if tc.wantErr {
				assert.Error(err)
				return
			}

			require.NoError(err)
			assert.Equal(tc.wantRole, role)
		})
	}
}
