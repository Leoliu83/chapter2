package os

import "testing"

func TestRunSqlplus(t *testing.T) {
	// RunSqlWithSqlplus("SELECT rownum,table_name,t.* FROM dba_tables t where rownum<50;", SqlplusEnvPureResult())
	Ping("127.0.0.1")
}

func TestInitInstancesUsers(t *testing.T) {
	IsValidCommand("sqlplus1")
}
