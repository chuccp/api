package vmess

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)
var moren = "vmess://ew0KICAidiI6ICIyIiwNCiAgInBzIjogImFsaSIsDQogICJhZGQiOiAid3MuY2oyMDIwMDcwOC50ayIsDQogICJwb3J0IjogIjgwODMiLA0KICAiaWQiOiAiN2NjNzU4OWUtN2ViYy0xMWVjLWEzNTItMDAxNjNlMDBiZGVmIiwNCiAgImFpZCI6ICIwIiwNCiAgInNjeSI6ICJhdXRvIiwNCiAgIm5ldCI6ICJ3cyIsDQogICJ0eXBlIjogIm5vbmUiLA0KICAiaG9zdCI6ICIiLA0KICAicGF0aCI6ICIvNENFclBPeWovIiwNCiAgInRscyI6ICJ0bHMiLA0KICAic25pIjogIiIsDQogICJhbHBuIjogIiINCn0="
func Api(c *gin.Context) {
	buff:=bytes.NewBufferString(moren)
	resp, err := http.Get("https://call.chenxing.gq/link/zjNcC6g3mOj3oZOo?sub=3&extend=1")
	if err == nil {
		all, err := io.ReadAll(resp.Body)
		if err == nil {
			decodeString, err := base64.StdEncoding.DecodeString(string(all))
			if err == nil {
				buff.Write([]byte("\r\n"))
				buff.Write(decodeString)
			}
		}
	}
	c.Writer.WriteString(	base64.StdEncoding.EncodeToString(buff.Bytes()))
}
func Api2(c *gin.Context) {
	conn, err := sqlite.OpenConn("test.db", 0)
	defer conn.Close()
	buff := new(bytes.Buffer)
	if err == nil {
		err = sqlitex.ExecuteTransient(conn, "select id,url from t_vmess order by url desc;", &sqlitex.ExecOptions{
			ResultFunc: func(stmt *sqlite.Stmt) error {
				v:=stmt.ColumnText(1)
				if !strings.HasPrefix(v,"http"){

					if buff.Len()>0{
						buff.Write([]byte("\r\n"))
					}
					buff.Write([]byte(v))
				}else{
					resp, err := http.Get(v)
					if err == nil {
						all, err := io.ReadAll(resp.Body)
						if err == nil {
							decodeString, err := base64.StdEncoding.DecodeString(string(all))
							if err == nil {
								if buff.Len()>0{
									buff.Write([]byte("\r\n"))
								}
								buff.Write(decodeString)
							}
						}
					}
				}
				return nil
			},
		})
	}
	c.Writer.WriteString(	base64.StdEncoding.EncodeToString(buff.Bytes()))

}