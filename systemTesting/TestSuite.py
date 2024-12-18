import unittest
import requests

# Define the base URL for the API
BASE_URL = "https://localhost:8080/v1/members"

# Disable SSL warnings for self-signed certificates (not recommended for production)
import urllib3
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)


class TestMemberAPI(unittest.TestCase):
    # Sample members to add
    members = [
        {"id": "1", "name": "Alice"},
        {"id": "2", "name": "Bob"},
        {"id": "3", "name": "Charlie"},
        {"id": "4", "name": "David"},
        {"id": "5", "name": "Eve"},
        {"id": "6", "name": "Frank"},
        {"id": "7", "name": "Grace"},
        {"id": "8", "name": "Hank"},
    ]

    def test_add_members(self):
        """Test adding members via POST requests."""
        for member in self.members:
            response = requests.post(
                BASE_URL, json=member, verify=False
            )  # Set verify=False for self-signed certs
            self.assertEqual(response.status_code, 200, f"Failed to add member {member['id']}")
            self.assertEqual(response.json()["id"], member["id"])
            self.assertEqual(response.json()["name"], member["name"])

    def test_list_members(self):
        """Test listing all members and verifying their presence."""
        response = requests.get(BASE_URL, verify=False)
        self.assertEqual(response.status_code, 200, "Failed to fetch members list")
        members_list = response.json()

        # Convert the list of members to a dictionary for easier verification
        members_dict = {member["id"]: member["name"] for member in members_list}

        for member in self.members:
            self.assertIn(member["id"], members_dict, f"Member ID {member['id']} missing in list")
            self.assertEqual(
                members_dict[member["id"]],
                member["name"],
                f"Member name mismatch for ID {member['id']}",
            )


if __name__ == "__main__":
    unittest.main()
