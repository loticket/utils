package stringx

import (
	"encoding/json"
	"testing"
)

func TestQrocde(t *testing.T){
   var ddddd map[string]interface{} = map[string]interface{}{
   	   "user_id":10,
   	   "user_info":30,
   }

	var Pubkey = `MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDloPwzAYD6esRgu+BVqlQepqXg3gxHNH9aRxbQOlgDVbjTlwLa7OaEWGirIasZBDlebrAjj+giaQmO38U3HCUwUEhVqEnZYh0B8dJUYw9q4XFYtsB2rzz97b914hY+owNF6QDocZMVLSAivUWfnlBcfHnnrCfZ9ZZErDkOQ7BZuwIDAQAB`

	//ssd,_:= CreateQrCodeAction(ddddd,"saoma",Pubkey)
	//t.Log(string(ssd))
	dddd,_:=CreateQrCodeActions(ddddd,"saoma",Pubkey)
	var dddddd map[string]interface{} = map[string]interface{}{
		"user_id":10,
		"user_info":30,
		"qrcode":dddd,
	}

	hst,_:=json.Marshal(dddddd)

	t.Log(string(hst))
}
