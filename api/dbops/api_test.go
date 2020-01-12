package dbops

import "testing"

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video")
	dbConn.Exec("truncate connents")
	dbConn.Exec("truncate sessions")
}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", TestDeleteUser)
	t.Run("Reget", testRegetUser)
}
func testAddUser(t *testing.T) {
	err := AddUserCredential("haoyong", "123456")
	if err != nil {
		t.Errorf("Error of AddUser:%v", err)
	}
}
func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("haoyong")
	if pwd != "123456" || err != nil {
		t.Errorf("Error of GetUser:%v", err)
	}
}
func TestDeleteUser(t *testing.T) {
	err := DeleteUser("haoyong", "123456")
	if err != nil {
		t.Errorf("Error DeleteUser error :%v", err)
	}
}
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("haoyong")
	if err != nil {
		t.Errorf("Error of RegetUser:%v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}
