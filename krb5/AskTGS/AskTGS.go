package AskTGS

import (
	"github.com/Serizao/GoRottenTomato/krb5/netWork"
	"github.com/Serizao/GoRottenTomato/krb5/procedure"
	"github.com/Serizao/GoRottenTomato/krb5/types"
	"fmt"
)

func AskTGS(tgsreq procedure.TGS_REQ, dcIP string, key types.EncryptionKey) (*procedure.TGS_REP, error) {
	mtgsreq, err := tgsreq.Marshal()
	if err != nil {
		return nil, fmt.Errorf("asktgs failed %v", err)
	}
	resp, err := netWork.SendToKDC(dcIP, mtgsreq)
	if err != nil {
		return nil, fmt.Errorf("asktgs requesting for %s failed %v", tgsreq.Req_Body.SName.Name_String, err)
	}


	return GetTGS(resp, key)
}
