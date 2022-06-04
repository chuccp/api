package vmess

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"strings"
	"testing"
	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

func TestDBName(t *testing.T) {

	conn, err := sqlite.OpenConn("test.db", 0)
	if err != nil {
		t.Log(err)
	}
	defer conn.Close()

	buff := new(bytes.Buffer)

	err = sqlitex.ExecuteTransient(conn, "select id,url from t_vmess order by url desc;", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			v:=stmt.ColumnText(1)
			if !strings.HasPrefix(v,"http"){
				buff.WriteString(v)
			}else{
				resp, err := http.Get(v)
				if err == nil {
					all, err := io.ReadAll(resp.Body)
					if err == nil {
						decodeString, err := base64.StdEncoding.DecodeString(string(all))
						if err == nil {
							buff.Write([]byte("\r\n"))
							buff.Write(decodeString)
						}
					}
				}else{
					t.Log(err)
				}
			}
			return nil
		},
	})
	if err != nil {
		t.Log(err)
	}else{

		println(buff.String())
	}

}
