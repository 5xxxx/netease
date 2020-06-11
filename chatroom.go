/*
 *
 * chatroom.go
 * netease-im
 *
 * Created by lintao on 2020/6/9 4:42 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package netease

import (
	"encoding/json"
	"errors"
	"net/url"

	"github.com/NSObjects/netease/path"
)

//creator	String	是	聊天室属主的账号accid
//name	String	是	聊天室名称，长度限制128个字符
//announcement	String	否	公告，长度限制4096个字符
//broadcasturl	String	否	直播地址，长度限制1024个字符
//ext	String	否	扩展字段，最长4096字符
//queuelevel	int	否	队列管理权限：0:所有人都有权限变更队列，1:只有主播管理员才能操作变更。默认0
type Chatroom struct {
	Creator         string `json:"creator"`
	Name            string `json:"name" `
	Announcement    string `json:"announcement" `
	Broadcasturl    string `json:"broadcasturl" `
	Ext             string `json:"ext" `
	Queuelevel      string `json:"queuelevel" `
	Roomid          int    `json:"roomid"`
	Valid           bool   `json:"valid"`
	Muted           bool   `json:"muted"`
	Onlineusercount int    `json:"onlineusercount"`
	Ionotify        bool   `json:"ionotify"`
}

//创建聊天室
func (n NetEaseIM) CreateChatroom(room Chatroom) (int, error) {

	b, err := n.request(path.ChatRoomCreate, structToMap(room))
	if err != nil {
		return 0, err
	}
	var resp struct {
		Chatroom struct {
			Roomid       int         `json:"roomid"`
			Valid        bool        `json:"valid"`
			Announcement interface{} `json:"announcement"`
			Name         string      `json:"name"`
			Broadcasturl string      `json:"broadcasturl"`
			Ext          string      `json:"ext"`
			Creator      string      `json:"creator"`
		} `json:"chatroom"`
		Code int `json:"code"`
	}
	if err = json.Unmarshal(b, &resp); err != nil {
		return 0, err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return 0, errors.New(s)
		}
	}

	return resp.Chatroom.Roomid, nil
}

//查询聊天室信息
//roomid				long	是	聊天室id
//needOnlineUserCount	String	否	是否需要返回在线人数，true或false，默认false
func (n NetEaseIM) GetChatroom(roomid string, needOnlineUserCount string) (Chatroom, error) {
	params := url.Values{}
	params.Set("roomid", roomid)
	params.Set("needOnlineUserCount", needOnlineUserCount)

	b, err := n.request(path.GetChatRoom, params)
	if err != nil {
		return Chatroom{}, err
	}
	var resp struct {
		Chatroom Chatroom `json:"chatroom"`
		Code     int      `json:"code"`
	}
	if err = json.Unmarshal(b, &resp); err != nil {
		return Chatroom{}, err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return Chatroom{}, errors.New(s)
		}
	}

	return resp.Chatroom, nil
}

//批量查询聊天室信息
//roomids	String	是	多个roomid，格式为：["6001","6002","6003"]（JSONArray对应的roomid，如果解析出错，会报414错误），限20个roomid
//needOnlineUserCount	String	否	是否需要返回在线人数，true或false，默认false
func (n NetEaseIM) GetBatchChatroom(roomids string, needOnlineUserCount string) ([]Chatroom, error) {
	params := url.Values{}
	params.Set("roomids", roomids)
	params.Set("needOnlineUserCount", needOnlineUserCount)
	b, err := n.request(path.GetBatchChatRoom, params)
	if err != nil {
		return nil, err
	}
	var resp struct {
		SuccRooms []Chatroom `json:"succRooms"`
		Code      int        `json:"code"`
	}
	if err = json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return nil, errors.New(s)
		}
	}

	return resp.SuccRooms, nil
}

func (n NetEaseIM) UpdateChatroom(room Chatroom) error {

	b, err := n.request(path.UpdateChatRoom, structToMap(room))
	if err != nil {
		return err
	}
	var resp struct {
		Code int `json:"code"`
	}
	if err = json.Unmarshal(b, &resp); err != nil {
		return err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return errors.New(s)
		}
	}

	return nil
}

//roomid	long	是	聊天室id
//operator	String	是	操作者账号，必须是创建者才可以操作
//valid	String	是	true或false，false:关闭聊天室；true:打开聊天室
func (n NetEaseIM) ToggleCloseStat(roomid string, operator string, valid string) error {
	params := url.Values{}
	params.Set("roomid", roomid)
	params.Set("operator", operator)
	params.Set("valid", valid)
	b, err := n.request(path.ToggleCloseStat, params)
	if err != nil {
		return err
	}
	var resp struct {
		Code int `json:"code"`
	}
	if err = json.Unmarshal(b, &resp); err != nil {
		return err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return errors.New(s)
		}
	}

	return nil
}

//roomid	long	是	聊天室id
//operator	String	是	操作者账号accid
//target	String	是	被操作者账号accid
//opt	int	是	操作：
//1: 设置为管理员，operator必须是创建者
//2:设置普通等级用户，operator必须是创建者或管理员
//-1:设为黑名单用户，operator必须是创建者或管理员
//-2:设为禁言用户，operator必须是创建者或管理员
//optvalue	String	是	true或false，true:设置；false:取消设置；
//执行“取消”设置后，若成员非禁言且非黑名单，则变成游客
//notifyExt	String	否	通知扩展字段，长度限制2048，请使用json格式
type MemberRole struct {
	Roomid    string `json:"roomid" `
	Operator  string `json:"operator"`
	Target    string `json:"target"`
	Opt       int    `json:"opt"`
	Optvalue  string `json:"optvalue"`
	NotifyExt string `json:"notify_ext"`
}

//设置聊天室内用户角色
func (n NetEaseIM) SetMemberRole(r MemberRole) error {

	b, err := n.request(path.ToggleCloseStat, structToMap(r))
	if err != nil {
		return err
	}
	var resp struct {
		Code int `json:"code"`
	}
	if err = json.Unmarshal(b, &resp); err != nil {
		return err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return errors.New(s)
		}
	}

	return nil
}

//roomid	long	是	聊天室id
//accid	String	是	进入聊天室的账号
//clienttype	int	否	1:weblink（客户端为web端时使用）; 2:commonlink（客户端为非web端时使用）;3:wechatlink(微信小程序使用), 默认1
//clientip	String	否	客户端ip，传此参数时，会根据用户ip所在地区，返回合适的地址
type ChatroomRequest struct {
	Roomid     string `json:"roomid" `
	Accid      string `json:"accid" `
	Clienttype string `json:"clienttype" `
	Clientip   string `json:"clientip" `
}

//请求聊天室地址与令牌
func (n NetEaseIM) RequestAddr(r ChatroomRequest) ([]string, error) {

	b, err := n.request(path.ToggleCloseStat, structToMap(r))
	if err != nil {
		return nil, err
	}
	var resp struct {
		Code int      `json:"code"`
		Addr []string `json:"addr" `
	}
	if err = json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return nil, errors.New(s)
		}
	}

	return resp.Addr, nil
}

//roomid	long	是	聊天室id
//msgId	String	是	客户端消息id，使用uuid等随机串，msgId相同的消息会被客户端去重
//fromAccid	String	是	消息发出者的账号accid
//msgType	int	是	消息类型：
//0: 表示文本消息，
//1: 表示图片，
//2: 表示语音，
//3: 表示视频，
//4: 表示地理位置信息，
//6: 表示文件，
//10: 表示Tips消息，
//100: 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
//resendFlag	int	否	重发消息标记，0：非重发消息，1：重发消息，如重发消息会按照msgid检查去重逻辑
//attach	String	否	文本消息：填写消息文案;
//其它类型消息，请参考 消息格式示例；
//长度限制4096字符
//ext	String	否	消息扩展字段，内容可自定义，请使用JSON格式，长度限制4096字符
//skipHistory	int	否	是否跳过存储云端历史，0：不跳过，即存历史消息；1：跳过，即不存云端历史；默认0
//abandonRatio	int	否	可选，消息丢弃的概率。取值范围[0-9999]；
//其中0代表不丢弃消息，9999代表99.99%的概率丢弃消息，默认不丢弃；
//注意如果填写了此参数，下面的highPriority参数则会无效；
//此参数可用于流控特定业务类型的消息。
//highPriority	Boolean	否	可选，true表示是高优先级消息，云信会优先保障投递这部分消息；false表示低优先级消息。默认false。
//强烈建议应用恰当选择参数，以便在必要时，优先保障应用内的高优先级消息的投递。若全部设置为高优先级，则等于没有设置。 高优先级消息可以设置进入后重发，见needHighPriorityMsgResend参数
//needHighPriorityMsgResend	Boolean	否	可选，true表示会重发消息，false表示不会重发消息。默认true。注:若设置为true， 用户离开聊天室之后重新加入聊天室，在有效期内还是会收到发送的这条消息，目前有效期默认30s。在没有配置highPriority时needHighPriorityMsgResend不生效。
//useYidun	int	否	可选，单条消息是否使用易盾反垃圾，可选值为0。
//0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。
//
//若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
//bid	String	否	可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
//antispam	String	否	对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。
//true或false, 默认false。
//只对消息类型为：100 自定义消息类型 的消息生效。
//antispamCustom	String	否	在antispam参数为true时生效。
//自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：
//
//{"type":1,"data":"custom content"}
//
//字段说明：
//1. type: 1：文本，2：图片。
//2. data: 文本内容or图片地址。
type ChatroomMsg struct {
	Roomid                    string `json:"roomid" `
	MsgId                     string `json:"msgId" `
	FromAccid                 string `json:"fromAccid" `
	MsgType                   string `json:"msgType" `
	ResendFlag                int    `json:"resendFlag" `
	Attach                    string `json:"attach" `
	Ext                       string `json:"ext" `
	SkipHistory               int    `json:"skipHistory" `
	AbandonRatio              int    `json:"abandonRatio" `
	HighPriority              bool   `json:"highPriority" `
	NeedHighPriorityMsgResend bool   `json:"needHighPriorityMsgResend" `
	UseYidun                  int    `json:"useYidun" `
	Bid                       string `json:"bid" `
	Antispam                  string `json:"antispam" `
	AntispamCustom            string `json:"antispamCustom" `
}

//往聊天室内发消息
func (n NetEaseIM) SendChatRoomMsg(r ChatroomMsg) error {

	b, err := n.request(path.ChatRoomSendMsg, structToMap(r))
	if err != nil {
		return err
	}
	var resp struct {
		Code int `json:"code"`
	}
	if err = json.Unmarshal(b, &resp); err != nil {
		return err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return errors.New(s)
		}
	}

	return nil
}

//roomid	long	是	聊天室id
//accids	JSONArray	是	机器人账号accid列表，必须是有效账号，账号数量上限100个
//roleExt	String	否	机器人信息扩展字段，请使用json格式，长度4096字符
//notifyExt	String	否	机器人进入聊天室通知的扩展字段，请使用json格式，长度2048字符
type ChatroomRobot struct {
	Roomid    string   `json:"roomid" `
	Accids    []string `json:"accids" `
	RoleExt   string   `json:"roleExt" `
	NotifyExt string   `json:"notifyExt" `
}

func (n NetEaseIM) AddRobot(r ChatroomRobot) error {

	b, err := n.request(path.AddRobot, structToMap(r))
	if err != nil {
		return err
	}
	var resp struct {
		Code int `json:"code"`
	}
	if err = json.Unmarshal(b, &resp); err != nil {
		return err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return errors.New(s)
		}
	}

	return nil
}