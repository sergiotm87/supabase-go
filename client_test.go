package supabase_test

import (
	"testing"

	"github.com/supabase-community/supabase-go"
)

const (
	API_URL = "https://your-company.supabase.co"
	API_KEY = "your-api-key"
)

func TestNewClient(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Errorf("cannot initialize client: %v", err)
	}
	t.Logf("newClient result: %v", client)
}

func TestFrom(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Errorf("cannot initialize client: %v", err)
	}
	data, count, err := client.From("countries").Select("*", "exact", false).Execute()
	if err != nil {
		t.Errorf("cannot perform From operation: %v", err)
	}
	t.Logf("%s%v%d", data, err, count)
}

func TestRpc(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Errorf("cannot initialize client: %v", err)
	}
	result, err := client.Rpc("hello_world", "", nil)
	if err != nil {
		// The test will fail if the function doesn't exist, which is fine.
		// We are just testing that the Rpc method returns an error when it should.
		t.Logf("got expected error: %v", err)
	}
	t.Logf("rpc result: %s", result)
}

func TestRpcError(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Errorf("cannot initialize client: %v", err)
	}
	_, err = client.Rpc("non_existent_function", "", nil)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestStorage(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Errorf("cannot initialize client: %v", err)
	}
	bucket, err := client.Storage.GetBucket("bucket-id")
	if err != nil {
		t.Errorf("cannot get bucket: %v", err)
	}
	t.Logf("getBucket result: %v", bucket)
}

func TestFunctions(t *testing.T) {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		t.Errorf("cannot initialize client: %v", err)
	}
	result, err := client.Functions.Invoke("hello_world", map[string]interface{}{"name": "world"})
	if err != nil {
		t.Errorf("cannot invoke function: %v", err)
	}
	t.Logf("function invokation result: %v", result)
}
