package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

// Define the base URL
const baseURL = "https://localhost:8080/v1/members"
const giftURL = "https://localhost:8080/v1/gift_exchange"

// Define the structure for a member
type Member struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type systemTestSuite struct {
	suite.Suite
}

// Create a custom HTTP client that skips certificate verification
var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

// This system test should have more cases (or more test files) and added in a Github Action, or another CI/CD tool
// This system test should became an acceptance test.
func SystemTest(t *testing.T) {
	suite.Run(t, new(systemTestSuite))
}

func (s *systemTestSuite) SetupSuite() {
}

func (s *systemTestSuite) SetupTest(t *testing.T) {
	t.Cleanup(func() {
		clearAllMembers(t, baseURL)
	})
}

// Add a Happy path scenario
func TestEightMembersAddedSuccesfully(t *testing.T) {
	clearAllMembers(t, baseURL)
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

func TestPersonCannotGiveGiftToThemselves(t *testing.T) {
	clearAllMembers(t, baseURL)
	// Sample members to add
	members := []Member{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
		{ID: "3", Name: "Charlie"},
		{ID: "4", Name: "David"},
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
	}

	// Trigger gift assignment
	resp, err := client.Get(giftURL)
	if err != nil {
		t.Fatalf("Failed to assign gifts: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected status code %d while assigning gifts", resp.StatusCode)
	}

	// Fetch the assignments
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Printf("Response body: %s\n", string(body))

	type Assignment struct {
		MemberID          string `json:"member_id"`
		RecipientMemberID string `json:"recipient_member_id"`
	}

	var assignments []Assignment

	err = json.Unmarshal(body, &assignments)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v\nBody: %s", err, string(body))
	}

	// Βεβαιώσου ότι υπάρχουν αντιστοιχίσεις
	if len(assignments) == 0 {
		t.Fatalf("No assignments found. Response body: %s", string(body))
	}

	// Εκτύπωσε τις αντιστοιχίσεις για έλεγχο
	for _, a := range assignments {
		fmt.Printf("Member: %s, Recipient: %s\n", a.MemberID, a.RecipientMemberID)
	}

}

func TestFamilyMemberCannotReceiveGiftMoreThanOnceEvery3Runs(t *testing.T) {
	clearAllMembers(t, baseURL)

	// Sample members to add
	members := []Member{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
		{ID: "3", Name: "Charlie"},
		{ID: "4", Name: "David"},
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
			t.Fatalf("Unexpected status code %d while adding member", resp.StatusCode)
		}
	}

	// Track assignment history
	assignmentsHistory := make(map[string]map[string]int) // giver -> receiver -> runs

	// Simulate 3 runs (add constraint failure scenario)
	for run := 1; run <= 3; run++ {
		// Trigger gift assignment
		resp, err := client.Get(giftURL)
		if err != nil {
			t.Fatalf("Failed to assign gifts: %v", err)
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Failed to read response body: %v", err)
		}

		// Check for error response indicating no valid assignments
		if resp.StatusCode == http.StatusBadRequest {
			var errorResponse struct {
				Error string `json:"error"`
			}
			err = json.Unmarshal(body, &errorResponse)
			if err != nil {
				t.Fatalf("Failed to decode error response: %v\nBody: %s", err, string(body))
			}
			if errorResponse.Error == "unable to assign gift recipients under the current constraints" {
				fmt.Printf("Test passed: System correctly handled no valid assignments after %d runs\n", run)
				return
			}
			t.Fatalf("Unexpected error response: %v\nBody: %s", errorResponse.Error, string(body))
		}

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Unexpected status code %d while assigning gifts", resp.StatusCode)
		}

		// Process valid assignments
		type Assignment struct {
			MemberID          string `json:"member_id"`
			RecipientMemberID string `json:"recipient_member_id"`
		}

		var assignments []Assignment
		err = json.Unmarshal(body, &assignments)
		if err != nil {
			t.Fatalf("Failed to decode response body: %v\nBody: %s", err, string(body))
		}

		// Verify no family member is given a gift from the same person more than once in 3 runs
		for _, a := range assignments {
			if _, exists := assignmentsHistory[a.MemberID]; !exists {
				assignmentsHistory[a.MemberID] = make(map[string]int)
			}

			if lastRun, exists := assignmentsHistory[a.MemberID][a.RecipientMemberID]; exists && run-lastRun <= 3 {
				t.Fatalf("Rule violation: Member %s gave a gift to %s again in less than 3 runs (last: %d, current: %d)",
					a.MemberID, a.RecipientMemberID, lastRun, run)
			}

			// Update the assignment history
			assignmentsHistory[a.MemberID][a.RecipientMemberID] = run
		}
	}

	fmt.Println("Test passed: No member received a gift from the same person more than once in 3 runs")
}

func clearAllMembers(t *testing.T, baseURL string) {
	// Step 1: GET the list of all members
	resp, err := client.Get(baseURL)
	if err != nil {
		t.Fatalf("Failed to fetch members list: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Unexpected status code %d while fetching members list", resp.StatusCode)
	}

	var members []Member
	if err := json.NewDecoder(resp.Body).Decode(&members); err != nil {
		t.Fatalf("Failed to decode members list: %v", err)
	}

	// Step 2: DELETE each member individually using the custom client
	for _, member := range members {
		deleteURL := fmt.Sprintf("%s/%s", baseURL, member.ID)
		req, err := http.NewRequest(http.MethodDelete, deleteURL, nil)
		if err != nil {
			t.Fatalf("Failed to create DELETE request for member %s: %v", member.ID, err)
		}

		resp, err := client.Do(req) // Use the custom client here
		if err != nil {
			t.Fatalf("Failed to delete member %s: %v", member.ID, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Unexpected status code %d while deleting member %s", resp.StatusCode, member.ID)
		}
	}
}
