package testhelpers

//
//import (
//	_ "bytes"
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"testing"
//)
//
//// Member represents the structure of a member in the system
//type Member struct {
//	ID   string `json:"id"`
//	Name string `json:"name"`
//}
//
//// ClearAllMembers performs GET and DELETE requests to remove all members
//func ClearAllMembers(t *testing.T, httpClient *http.Client, baseURL string) {
//	// Step 1: GET the list of all members
//	resp, err := http.Get(baseURL)
//	if err != nil {
//		t.Fatalf("Failed to fetch members list: %v", err)
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		t.Fatalf("Unexpected status code %d while fetching members list", resp.StatusCode)
//	}
//
//	var members []Member
//	if err := json.NewDecoder(resp.Body).Decode(&members); err != nil {
//		t.Fatalf("Failed to decode members list: %v", err)
//	}
//
//	// Step 2: DELETE each member individually
//	for _, member := range members {
//		deleteURL := fmt.Sprintf("%s/members/%s", baseURL, member.ID)
//		req, err := http.NewRequest(http.MethodDelete, deleteURL, nil)
//		if err != nil {
//			t.Fatalf("Failed to create DELETE request for member %s: %v", member.ID, err)
//		}
//
//		resp, err := http.DefaultClient.Do(req)
//		if err != nil {
//			t.Fatalf("Failed to delete member %s: %v", member.ID, err)
//		}
//		defer resp.Body.Close()
//
//		if resp.StatusCode != http.StatusOK {
//			t.Fatalf("Unexpected status code %d while deleting member %s", resp.StatusCode, member.ID)
//		}
//	}
//}
