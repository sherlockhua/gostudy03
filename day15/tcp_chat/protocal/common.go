package protocal


import (
	"fmt"
	"encoding/json"
	"encoding/binary"
	"bytes"
	"github.com/gostudy03/xlog"
	"net"
)

const (
	UserEnterRoomCmd = 1001
	UserLeaveRoomCmd = 1002
	UserSendTextCmd  = 1003
	UserRecvTextCmd  = 1004
	AllRoomListCmd   = 1005
	GetRoomListCmd   = 1006
	UserEnterRoomRespCmd = 1007
	BoardcastUserEnterRoomCmd = 1008
)

type Protocal struct {
	Length uint32
	CmdNo  uint16
}

type BoardcastUserEnterRoom struct {
	EnterUserId uint64
	EnterUserName string
	RoomInfo *Room
}

type UserEnterRoom struct {
	RoomId uint64
	UserId uint64
	UserName string
}

type UserEnterRoomResp struct {
	Code int
	RoomInfo *Room
}

type UserLeaveRoom struct {
	RoomId uint64
	UserId uint64
	UserName string
}

type UserSendText struct {
	RoomId uint64
	UserId uint64
	UserName string
	Content string
}

type UserRecvText struct {
	Code int
	RoomId uint64
	AuthorUserId uint64
	AuthorUserName string
	Content string
}

type Room struct {
	RoomId uint64
	Name  string
	RoomCap uint32
	Desc string
	Online uint32
}

type AllRoomList struct {
	RoomList []*Room
}

type GetRoomList struct {
	UserId uint64
}

func Pack(cmdNo uint16, body interface{})(data []byte, err error) {

	bodyJson, err := json.Marshal(body)
	if err != nil {
		xlog.LogError("marshal json failed, body:%v, err:%v", body, err)
		return
	}

	var protocal *Protocal = &Protocal{
		Length: uint32(len(bodyJson)),
		CmdNo:cmdNo,
	}

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, protocal.Length)
	binary.Write(&buffer, binary.BigEndian, protocal.CmdNo)
	binary.Write(&buffer, binary.BigEndian, bodyJson)
	
	data = buffer.Bytes()
	return
}

func UnPack(conn net.Conn) (cmdNo uint16, result interface{}, err error){
	//1. 读取包的长度和命令号
	buf := make([]byte, 6)
	_, err = conn.Read(buf)
	if err != nil {
		xlog.LogError("read package header err:", err)
		return
	}

	var proto Protocal
	buffer := bytes.NewBuffer(buf)
	binary.Read(buffer, binary.BigEndian, &proto.Length)
	binary.Read(buffer, binary.BigEndian, &proto.CmdNo)

	cmdNo = proto.CmdNo
	//2. 读取包体
	bodyBuf := make([]byte, proto.Length)
	_, err = conn.Read(bodyBuf)
	if err != nil {
		xlog.LogError("read from network failed, err:%v", err)
		return
	}

	//3. 根据命令号，反序列化成对应命令号的结构体
	var tmpBody interface{}
	switch proto.CmdNo {
	case GetRoomListCmd:
		tmpBody = &GetRoomList{}
	case AllRoomListCmd:
		tmpBody = &AllRoomList{}
	case UserEnterRoomCmd:
		tmpBody = &UserEnterRoom{}
	case UserEnterRoomRespCmd:
		tmpBody = &UserEnterRoomResp{}
	case BoardcastUserEnterRoomCmd:
		tmpBody = &BoardcastUserEnterRoom{}
	case UserSendTextCmd:
		tmpBody = &UserSendText{}
	case UserRecvTextCmd:
		tmpBody = &UserRecvText{}
	default:
		err = fmt.Errorf("unsupport command:%d", proto.CmdNo)
		xlog.LogError("unsupport command:%d, data:%v", proto.CmdNo, string(bodyBuf))
		return
	}

	err = json.Unmarshal(bodyBuf, tmpBody)
	if err != nil {
		xlog.LogError("unmarshal failed, data:%v, err:%v", string(bodyBuf), err)
		return
	}

	result = tmpBody
	xlog.LogDebug("get message succ, cmdNo:%d, data length:%d, data:%#v", 
		cmdNo, proto.Length,result)
	return
}
	