package main

import (
	vault "github.com/hashicorp/vault/api"
	"testing"
)

const expectedPolicy = `
#
# Policy for Kubeflow profile: profile-test
# (policy managed by the custom Kubeflow Profiles controller)
#

# Grant full access to the KV created for this profile
path "kv_profile-test/*" {
	capabilities = ["create", "update", "delete", "read", "list"]
}

# Grant access to MinIO keys associated with this profile
path "minio1/keys/profile-test" {
	capabilities = ["read"]
}
path "minio2/keys/profile-test" {
	capabilities = ["read"]
}
`

func TestGeneratePolicy(t *testing.T) {

	policy, err := generatePolicy("profile-test", []string{"minio1", "minio2"})

	if err != nil {
		t.Fatal(err)
	}

	if policy != expectedPolicy {
		println(policy)
		t.Fail()
	}
}

// Tests updatePolicy when there is already a policy
// defined in Vault
func TestDoPolicy_updatePolicy(t *testing.T) {
	var vc = VaultConfigurerStruct{
		Logical: &VaultLogicalAPIMock{
			ReadFunc: func(path string) (*vault.Secret, error) {

				if path != "/sys/policies/acl/profile-test" {
					t.Log("Wrong path")
					t.Fail()
				}
				return &vault.Secret{
					Renewable: false,
					Data: map[string]interface{}{
						"name":   "profile-test",
						"policy": "This policy is out of data",
					},
				}, nil
			},
			WriteFunc: func(path string, data map[string]interface{}) (*vault.Secret, error) {

				if path != "/sys/policies/acl/profile-test" {
					t.Log("Wrong path")
					t.Fail()
				}

				if data["policy"] != expectedPolicy {
					t.Log("Policy was not generated as expected")
					t.Fail()
				}

				return &vault.Secret{}, nil
			},
		},
		MinioInstances: []string{"minio1", "minio2"},
	}
	policyName, _ := vc.doPolicy("profile-test", "kv_profile-test")

	if policyName != "profile-test" {
		t.Logf("Expected profile-test as policy name, got %s", policyName)
		t.Fail()
	}
}

//func TestDoKVMount_NoMount(t *testing.T) {
//	var vc = VaultConfigurerStruct{
//		Logical: nil,
//		Mounts: &VaultMountsAPIMock{
//			ListMountsFunc: func() (map[string]*vault.MountOutput, error) {
//
//			},
//			MountFunc: func(path string, mountInfo *vault.MountInput) error {
//
//			},
//		},
//		MinioInstances:     []string{"minio1", "minio2"},
//	}
//}
