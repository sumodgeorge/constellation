package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestAccessManagerMarshalUnmarshal(t *testing.T) {
	require := require.New(t)
	assert := assert.New(t)

	// Without data
	accessManagerDeplNil := NewAccessManagerDeployment(nil)
	data, err := accessManagerDeplNil.Marshal()
	require.NoError(err)

	var recreated accessManagerDeployment
	require.NoError(UnmarshalK8SResources(data, &recreated))
	assert.Equal(accessManagerDeplNil, &recreated)

	// With data
	sshUsers := make(map[string]string)
	sshUsers["test-user"] = "ssh-rsa abcdefg"
	accessManagerDeplNil = NewAccessManagerDeployment(sshUsers)
	data, err = accessManagerDeplNil.Marshal()
	require.NoError(err)

	require.NoError(UnmarshalK8SResources(data, &recreated))
	assert.Equal(accessManagerDeplNil, &recreated)
}
