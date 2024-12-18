package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

// Define the base URL
const baseURL = "https://localhost:8080/v1/members"

// Define the structure for a member
type Member struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Create a custom HTTP client that skips certificate verification
var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

func TestMemberAPI(t *testing.T) {
	// Sample members to add
	members := []Member{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
		{ID: "3", Name: "Charlie"},
		{ID: "4", Name: "David"},
		{ID: "5", Name: "Eve"},
		{ID: "6", Name: "Frank"},
		{ID: "7", Name: "Grace"},
		{ID: "8", Name: "Hank"},
	}

	// Add members using POST
	for _, member := range members {
		body, _ := json.Marshal(member)
		resp, err := client.Post(baseURL, "application/json", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("Failed to add member %s: %v", member.ID, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Unexpected status code %d while adding member %s", resp.StatusCode, member.ID)
		}

		var addedMember Member
		err = json.NewDecoder(resp.Body).Decode(&addedMember)
		if err != nil {
			t.Fatalf("Failed to decode response for member %s: %v", member.ID, err)
		}
		if addedMember.ID == "" || addedMember.Name == "" {
			t.Fatalf("API returned incomplete member data: %+v", addedMember)
		}

		if addedMember.ID != member.ID || addedMember.Name != member.Name {
			t.Fatalf("Mismatch in added member. Expected %+v, got %+v", member, addedMember)
		}
	}

	// List all members using GET
	resp, err := client.Get(baseURL)
	if err != nil {
		t.Fatalf("Failed to fetch members list: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected status code %d while fetching members list", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var fetchedMembers []Member
	err = json.Unmarshal(body, &fetchedMembers)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Verify all members are in the list
	membersMap := make(map[string]string)
	for _, member := range fetchedMembers {
		membersMap[member.ID] = member.Name
	}

	for _, member := range members {
		name, exists := membersMap[member.ID]
		if !exists {
			t.Fatalf("Member ID %s not found in list", member.ID)
		}
		if name != member.Name {
			t.Fatalf("Mismatch in member name for ID %s. Expected %s, got %s", member.ID, member.Name, name)
		}
	}
}
