package gcp

import (
	"context"
	"errors"
	"testing"

	"github.com/edgelesssys/constellation/coordinator/core"
	"github.com/edgelesssys/constellation/coordinator/role"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	err := errors.New("some err")
	uid := "1234"
	instancesGenerator := func() *[]core.Instance {
		return &[]core.Instance{
			{
				Name:       "someInstance",
				ProviderID: "gce://someProject/someZone/someInstance",
				IPs:        []string{"192.0.2.0"},
			},
		}
	}

	testCases := map[string]struct {
		client             stubGCPClient
		instancesGenerator func() *[]core.Instance
		instancesMutator   func(*[]core.Instance)
		wantErr            bool
		wantInstances      []core.Instance
	}{
		"retrieve works": {
			client: stubGCPClient{
				projectID: "someProjectID",
				zone:      "someZone",
				retrieveInstanceMetadaValues: map[string]string{
					"constellation-uid": uid,
				},
			},
			instancesGenerator: instancesGenerator,
			wantInstances: []core.Instance{
				{
					Name:       "someInstance",
					ProviderID: "gce://someProject/someZone/someInstance",
					IPs:        []string{"192.0.2.0"},
				},
			},
		},
		"retrieve error is detected": {
			client: stubGCPClient{
				projectID: "someProjectID",
				zone:      "someZone",
				retrieveInstanceMetadaValues: map[string]string{
					"constellation-uid": uid,
				},
				retrieveInstancesErr: err,
			},
			instancesGenerator: instancesGenerator,
			wantErr:            true,
		},
		"project metadata retrieval error is detected": {
			client: stubGCPClient{
				retrieveProjectIDErr: err,
			},
			instancesGenerator: instancesGenerator,
			wantErr:            true,
		},
		"zone retrieval error is detected": {
			client: stubGCPClient{
				retrieveZoneErr: err,
			},
			instancesGenerator: instancesGenerator,
			wantErr:            true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			tc.client.retrieveInstancesValues = *tc.instancesGenerator()
			if tc.instancesMutator != nil {
				tc.instancesMutator(&tc.client.retrieveInstancesValues)
			}
			metadata := New(&tc.client)
			instances, err := metadata.List(context.Background())

			if tc.wantErr {
				assert.Error(err)
				return
			}
			require.NoError(err)
			assert.ElementsMatch(tc.wantInstances, instances)
		})
	}
}

func TestSelf(t *testing.T) {
	err := errors.New("some err")
	uid := "1234"

	testCases := map[string]struct {
		client       stubGCPClient
		wantErr      bool
		wantInstance core.Instance
	}{
		"retrieve works": {
			client: stubGCPClient{
				projectID: "someProjectID",
				zone:      "someZone",
				retrieveInstanceValue: core.Instance{
					Name:       "someInstance",
					ProviderID: "gce://someProject/someZone/someInstance",
					IPs:        []string{"192.0.2.0"},
				},
			},
			wantInstance: core.Instance{
				Name:       "someInstance",
				ProviderID: "gce://someProject/someZone/someInstance",
				IPs:        []string{"192.0.2.0"},
			},
		},
		"retrieve error is detected": {
			client: stubGCPClient{
				projectID: "someProjectID",
				zone:      "someZone",
				retrieveInstanceMetadaValues: map[string]string{
					"constellation-uid": uid,
				},
				retrieveInstanceErr: err,
			},
			wantErr: true,
		},
		"project id retrieval error is detected": {
			client: stubGCPClient{
				retrieveProjectIDErr: err,
			},
			wantErr: true,
		},
		"zone retrieval error is detected": {
			client: stubGCPClient{
				retrieveZoneErr: err,
			},
			wantErr: true,
		},
		"instance name retrieval error is detected": {
			client: stubGCPClient{
				retrieveInstanceNameErr: err,
			},
			wantErr: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			cloud := New(&tc.client)
			instance, err := cloud.Self(context.Background())

			if tc.wantErr {
				assert.Error(err)
				return
			}
			require.NoError(err)
			assert.Equal(tc.wantInstance, instance)
		})
	}
}

func TestGetInstance(t *testing.T) {
	err := errors.New("some err")

	testCases := map[string]struct {
		providerID   string
		client       stubGCPClient
		wantErr      bool
		wantInstance core.Instance
	}{
		"retrieve works": {
			providerID: "gce://someProject/someZone/someInstance",
			client: stubGCPClient{
				retrieveInstanceValue: core.Instance{
					Name:       "someInstance",
					ProviderID: "gce://someProject/someZone/someInstance",
					IPs:        []string{"192.0.2.0"},
				},
			},
			wantInstance: core.Instance{
				Name:       "someInstance",
				ProviderID: "gce://someProject/someZone/someInstance",
				IPs:        []string{"192.0.2.0"},
			},
		},
		"retrieve error is detected": {
			providerID: "gce://someProject/someZone/someInstance",
			client: stubGCPClient{
				retrieveInstanceErr: err,
			},
			wantErr: true,
		},
		"malformed providerID with too many fields is detected": {
			providerID: "gce://someProject/someZone/someInstance/tooMany/fields",
			wantErr:    true,
		},
		"malformed providerID with too few fields is detected": {
			providerID: "gce://someProject",
			wantErr:    true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			cloud := New(&tc.client)
			instance, err := cloud.GetInstance(context.Background(), tc.providerID)

			if tc.wantErr {
				assert.Error(err)
				return
			}
			require.NoError(err)
			assert.Equal(tc.wantInstance, instance)
		})
	}
}

func TestSignalRole(t *testing.T) {
	err := errors.New("some err")

	testCases := map[string]struct {
		client   stubGCPClient
		wantErr  bool
		wantRole role.Role
	}{
		"signaling role works": {
			client: stubGCPClient{
				projectID:    "someProjectID",
				zone:         "someZone",
				instanceName: "someName",
			},
			wantRole: role.Coordinator,
		},
		"project metadata retrieval error is detected": {
			client: stubGCPClient{
				retrieveProjectIDErr: err,
			},
			wantErr: true,
		},
		"instance zone retrieval error is detected": {
			client: stubGCPClient{
				retrieveZoneErr: err,
			},
			wantErr: true,
		},
		"instance name retrieval error is detected": {
			client: stubGCPClient{
				retrieveInstanceNameErr: err,
			},
			wantErr: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			cloud := New(&tc.client)
			err := cloud.SignalRole(context.Background(), tc.wantRole)

			if tc.wantErr {
				assert.Error(err)
				return
			}
			require.NoError(err)

			assert.ElementsMatch([]string{"someProjectID"}, tc.client.instanceMetadataProjects)
			assert.ElementsMatch([]string{"someZone"}, tc.client.instanceMetadataZones)
			assert.ElementsMatch([]string{"someName"}, tc.client.instanceMetadataInstanceNames)
			assert.ElementsMatch([]string{core.RoleMetadataKey}, tc.client.instanceMetadataKeys)
			assert.ElementsMatch([]string{tc.wantRole.String()}, tc.client.instanceMetadataValues)
		})
	}
}

func TestSetVPNIP(t *testing.T) {
	err := errors.New("some err")

	testCases := map[string]struct {
		client    stubGCPClient
		wantErr   bool
		wantVPNIP string
	}{
		"signaling role works": {
			client: stubGCPClient{
				projectID:    "someProjectID",
				zone:         "someZone",
				instanceName: "someName",
			},
			wantVPNIP: "192.0.2.0",
		},
		"project metadata retrieval error is detected": {
			client: stubGCPClient{
				retrieveProjectIDErr: err,
			},
			wantErr: true,
		},
		"instance zone retrieval error is detected": {
			client: stubGCPClient{
				retrieveZoneErr: err,
			},
			wantErr: true,
		},
		"instance name retrieval error is detected": {
			client: stubGCPClient{
				retrieveInstanceNameErr: err,
			},
			wantErr: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			cloud := New(&tc.client)
			err := cloud.SetVPNIP(context.Background(), tc.wantVPNIP)

			if tc.wantErr {
				assert.Error(err)
				return
			}
			require.NoError(err)

			assert.ElementsMatch([]string{"someProjectID"}, tc.client.instanceMetadataProjects)
			assert.ElementsMatch([]string{"someZone"}, tc.client.instanceMetadataZones)
			assert.ElementsMatch([]string{"someName"}, tc.client.instanceMetadataInstanceNames)
			assert.ElementsMatch([]string{core.VPNIPMetadataKey}, tc.client.instanceMetadataKeys)
			assert.ElementsMatch([]string{tc.wantVPNIP}, tc.client.instanceMetadataValues)
		})
	}
}

func TestTrivialMetadataFunctions(t *testing.T) {
	assert := assert.New(t)
	metadata := Metadata{}

	assert.True(metadata.Supported())
}

type stubGCPClient struct {
	retrieveInstanceValue        core.Instance
	retrieveInstanceErr          error
	retrieveInstancesValues      []core.Instance
	retrieveInstancesErr         error
	retrieveInstanceMetadaValues map[string]string
	retrieveInstanceMetadataErr  error
	projectID                    string
	zone                         string
	instanceName                 string
	retrieveProjectIDErr         error
	retrieveZoneErr              error
	retrieveInstanceNameErr      error
	setInstanceMetadataErr       error
	unsetInstanceMetadataErr     error

	instanceMetadataProjects      []string
	instanceMetadataZones         []string
	instanceMetadataInstanceNames []string
	instanceMetadataKeys          []string
	instanceMetadataValues        []string

	unsetMetadataProjects      []string
	unsetMetadataZones         []string
	unsetMetadataInstanceNames []string
	unsetMetadataKeys          []string
}

func (s *stubGCPClient) RetrieveInstances(ctx context.Context, project, zone string) ([]core.Instance, error) {
	return s.retrieveInstancesValues, s.retrieveInstancesErr
}

func (s *stubGCPClient) RetrieveInstance(ctx context.Context, project, zone string, instanceName string) (core.Instance, error) {
	return s.retrieveInstanceValue, s.retrieveInstanceErr
}

func (s *stubGCPClient) RetrieveInstanceMetadata(attr string) (string, error) {
	return s.retrieveInstanceMetadaValues[attr], s.retrieveInstanceMetadataErr
}

func (s *stubGCPClient) RetrieveProjectID() (string, error) {
	return s.projectID, s.retrieveProjectIDErr
}

func (s *stubGCPClient) RetrieveZone() (string, error) {
	return s.zone, s.retrieveZoneErr
}

func (s *stubGCPClient) RetrieveInstanceName() (string, error) {
	return s.instanceName, s.retrieveInstanceNameErr
}

func (s *stubGCPClient) SetInstanceMetadata(ctx context.Context, project, zone, instanceName, key, value string) error {
	s.instanceMetadataProjects = append(s.instanceMetadataProjects, project)
	s.instanceMetadataZones = append(s.instanceMetadataZones, zone)
	s.instanceMetadataInstanceNames = append(s.instanceMetadataInstanceNames, instanceName)
	s.instanceMetadataKeys = append(s.instanceMetadataKeys, key)
	s.instanceMetadataValues = append(s.instanceMetadataValues, value)

	return s.setInstanceMetadataErr
}

func (s *stubGCPClient) UnsetInstanceMetadata(ctx context.Context, project, zone, instanceName, key string) error {
	s.unsetMetadataProjects = append(s.unsetMetadataProjects, project)
	s.unsetMetadataZones = append(s.unsetMetadataZones, zone)
	s.unsetMetadataInstanceNames = append(s.unsetMetadataInstanceNames, instanceName)
	s.unsetMetadataKeys = append(s.unsetMetadataKeys, key)

	return s.unsetInstanceMetadataErr
}
