package nets

import (
	"Satellite/pack"
	"errors"
)

type GoApi struct {

}

func (api *GoApi) MD5Encode(request TNetsRpcPackMD5EncodeReq, response *TNetsRpcPackMD5EncodeResp) (err error) {
	if request.Src == "" {
		err = errors.New("source string can not be empty")
		return err
	}
	response.Dest = pack.MD5Encode(request.Src)
	return err
}

func (api *GoApi) MD5Equal(request TNetsRpcPackMD5EqualReq, response *TNetsRpcPackMD5EqualResp) (err error) {
	if request.Src == "" || request.Dest == "" {
		err = errors.New("source or destination string can not be empty")
		return err
	}
	response.Equal = pack.MD5Check(request.Src, request.Dest)
	return err
}
