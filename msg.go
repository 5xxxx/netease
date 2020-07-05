/*
 *
 * msg.go
 * netease-im
 *
 * Created by lintao on 2020/6/9 3:15 下午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package netease

import (
	"encoding/json"
	"errors"

	"github.com/NSObjects/netease/path"
)

//from	String	是	发送者accid，用户帐号，最大32字符，
//必须保证一个APP内唯一
//ope	int	是	0：点对点个人消息，1：群消息（高级群），其他返回414
//to	String	是	ope==0是表示accid即用户id，ope==1表示tid即群id
//type	int	是	0 表示文本消息, 1 表示图片， 2 表示语音， 3 表示视频， 4 表示地理位置信息， 6 表示文件， 10 表示提示消息， 100 自定义消息类型（特别注意，对于未对接易盾反垃圾功能的应用，该类型的消息不会提交反垃圾系统检测）
//body	String	是	最大长度5000字符，JSON格式。 具体请参考： 消息格式示例
//antispam	String	否	对于对接了易盾反垃圾功能的应用，本消息是否需要指定经由易盾检测的内容（antispamCustom）。
//true或false, 默认false。 只对消息类型为：100 自定义消息类型 的消息生效。
//antispamCustom	String	否	在antispam参数为true时生效。 自定义的反垃圾检测内容, JSON格式，长度限制同body字段，不能超过5000字符，要求antispamCustom格式如下：
//{"type":1,"data":"custom content"}
//
//字段说明：
//1. type: 1：文本，2：图片。
//2. data: 文本内容or图片地址。
//option	String	否	发消息时特殊指定的行为选项,JSON格式，可用于指定消息的漫游，存云端历史，发送方多端同步，推送，消息抄送等特殊行为;option中字段不填时表示默认值 ，option示例:
//
//{"push":false,"roam":true,"history":false,"sendersync":true,"route":false,"badge":false,"needPushNick":true}
//
//字段说明：
//1. roam: 该消息是否需要漫游，默认true（需要app开通漫游消息功能）；
//2. history: 该消息是否存云端历史，默认true；
//3. sendersync: 该消息是否需要发送方多端同步，默认true；
//4. push: 该消息是否需要APNS推送或安卓系统通知栏推送，默认true；
//5. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能);
//6. badge:该消息是否需要计入到未读计数中，默认true;
//7. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认true;
//8. persistent: 是否需要存离线消息，不设置该参数时默认true。
//pushcontent	String	否	推送文案,最长500个字符。具体参见 推送配置参数详解。
//payload	String	否	必须是JSON,不能超过2k字符。该参数与APNs推送的payload含义不同。具体参见 推送配置参数详解。
//ext	String	否	开发者扩展字段，长度限制1024字符
//forcepushlist	String	否	发送群消息时的强推用户列表（云信demo中用于承载被@的成员），格式为JSONArray，如["accid1","accid2"]。若forcepushall为true，则forcepushlist为除发送者外的所有有效群成员
//forcepushcontent	String	否	发送群消息时，针对强推列表forcepushlist中的用户，强制推送的内容
//forcepushall	String	否	发送群消息时，强推列表是否为群里除发送者外的所有有效成员，true或false，默认为false
//bid	String	否	可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
//useYidun	int	否	可选，单条消息是否使用易盾反垃圾，可选值为0。
//0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。
//
//若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
//markRead	int	否	可选，群消息是否需要已读业务（仅对群消息有效），0:不需要，1:需要
//checkFriend	boolean	否	是否为好友关系才发送消息，默认否
//注：使用该参数需要先开通功能服务
type Msg struct {
	From             string `json:"from" `
	Ope              int    `json:"ope" `
	To               string `json:"to" `
	Type             int    `json:"type" `
	Body             string `json:"body"`
	Antispam         string `json:"antispam" `
	AntispamCustom   string `json:"antispamCustom"`
	Option           string `json:"option"`
	Pushcontent      string `json:"pushcontent"`
	Payload          string `json:"payload"`
	Ext              string `json:"ext"`
	Forcepushlist    string `json:"forcepushlist"`
	Forcepushcontent string `json:"forcepushcontent"`
	Forcepushall     string `json:"forcepushall"`
	Bid              string `json:"bid"`
	UseYidun         int    `json:"useYidun"`
	MarkRead         int    `json:"mark_read"`
	CheckFriend      bool   `json:"check_friend" `
}

type MsgResult struct {
	Msgid    int64 `json:"msgid"`
	Timetag  int64 `json:"timetag"`
	Antispam bool  `json:"antispam"`
}

//发送普通消息
//给用户或者高级群发送普通消息，包括文本，图片，语音，视频和地理位置
func (n NetEaseIM) SendMsg(msg Msg) (MsgResult, error) {
	b, err := n.request(path.SendMsg, structToMap(msg))
	if err != nil {
		return MsgResult{}, err
	}
	var resp struct {
		Code int       `json:"code"`
		Data MsgResult `json:"data"`
	}

	if err = json.Unmarshal(b, &resp); err != nil {
		return MsgResult{}, err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return MsgResult{}, errors.New(s)
		}
	}

	return resp.Data, nil
}

//fromAccid	String	是	发送者accid，用户帐号，最大32字符，
//必须保证一个APP内唯一
//toAccids	String	是	["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414错误），限500人
//type	int	是	0 表示文本消息, 1 表示图片， 2 表示语音， 3 表示视频， 4 表示地理位置信息， 6 表示文件， 10 表示提示消息， 100 自定义消息类型
//body	String	是	最大长度5000字符，JSON格式。 具体请参考： 消息格式示例
//option	String	否	发消息时特殊指定的行为选项,Json格式，可用于指定消息的漫游，存云端历史，发送方多端同步，推送，消息抄送等特殊行为;option中字段不填时表示默认值 option示例:
//
//{"push":false,"roam":true,"history":false,"sendersync":true,"route":false,"badge":false,"needPushNick":true}
//
//字段说明：
//1. roam: 该消息是否需要漫游，默认true（需要app开通漫游消息功能）；
//2. history: 该消息是否存云端历史，默认true；
//3. sendersync: 该消息是否需要发送方多端同步，默认true；
//4. push: 该消息是否需要APNS推送或安卓系统通知栏推送，默认true；
//5. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能);
//6. badge:该消息是否需要计入到未读计数中，默认true;
//7. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认true;
//8. persistent: 是否需要存离线消息，不设置该参数时默认true。
//pushcontent	String	否	推送文案，最长500个字符。具体参见 推送配置参数详解。
//payload	String	否	必须是JSON,不能超过2k字符。该参数与APNs推送的payload含义不同。具体参见 推送配置参数详解。
//ext	String	否	开发者扩展字段，长度限制1024字符
//bid	String	否	可选，反垃圾业务ID，实现“单条消息配置对应反垃圾”，若不填则使用原来的反垃圾配置
//useYidun	int	否	可选，单条消息是否使用易盾反垃圾，可选值为0。
//0：（在开通易盾的情况下）不使用易盾反垃圾而是使用通用反垃圾，包括自定义消息。
//
//若不填此字段，即在默认情况下，若应用开通了易盾反垃圾功能，则使用易盾反垃圾来进行垃圾消息的判断
//returnMsgid	Boolean	否	是否需要返回消息ID
//false：不返回消息ID（默认值）
//true：返回消息ID（toAccids包含的账号数量不可以超过100个）
type BatchMsg struct {
	FromAccid   string   `json:"fromAccid" `
	ToAccids    []string `json:"toAccids" `
	Type        int      `json:"type" `
	Body        string   `json:"body"`
	Option      string   `json:"option"`
	Pushcontent string   `json:"pushcontent"`
	Payload     string   `json:"payload"`
	Ext         string   `json:"ext"`
	Bid         string   `json:"bid"`
	UseYidun    int      `json:"useYidun"`
	ReturnMsgid bool     `json:"returnMsgid" `
}

//批量发送点对点普通消息
//1.给用户发送点对点普通消息，包括文本，图片，语音，视频，地理位置和自定义消息。
//2.最大限500人，只能针对个人,如果批量提供的帐号中有未注册的帐号，会提示并返回给用户。
//3.此接口受频率控制，一个应用一分钟最多调用120次，超过会返回416状态码，并且被屏蔽一段时间；
//具体消息参考下面描述。
func (n NetEaseIM) SendBatchMsg(msg BatchMsg) (string, error) {
	b, err := n.request(path.SendBatchMsg, structToMap(msg))
	if err != nil {
		return "", err
	}
	var resp struct {
		Code       int    `json:"code"`
		Unregister string `json:"unregister" `
	}

	if err = json.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return "", errors.New(s)
		}
	}

	return resp.Unregister, nil
}

//from	String	是	发送者accid，用户帐号，最大32字符，APP内唯一
//msgtype	int	是	0：点对点自定义通知，1：群消息自定义通知，其他返回414
//to	String	是	msgtype==0是表示accid即用户id，msgtype==1表示tid即群id
//attach	String	是	自定义通知内容，第三方组装的字符串，建议是JSON串，最大长度4096字符
//pushcontent	String	否	推送文案，最长500个字符。具体参见 推送配置参数详解。
//payload	String	否	必须是JSON,不能超过2k字符。该参数与APNs推送的payload含义不同。具体参见 推送配置参数详解。
//sound	String	否	如果有指定推送，此属性指定为客户端本地的声音文件名，长度不要超过30个字符，如果不指定，会使用默认声音
//save	int	否	1表示只发在线，2表示会存离线，其他会报414错误。默认会存离线
//option	String	否	发消息时特殊指定的行为选项,Json格式，可用于指定消息计数等特殊行为;option中字段不填时表示默认值。
//option示例：
//{"badge":false,"needPushNick":false,"route":false}
//字段说明：
//1. badge:该消息是否需要计入到未读计数中，默认true;
//2. needPushNick: 推送文案是否需要带上昵称，不设置该参数时默认false(ps:注意与sendMsg.action接口有别);
//3. route: 该消息是否需要抄送第三方；默认true (需要app开通消息抄送功能)
//isForcePush	String	否	发自定义通知时，是否强制推送
//forcePushContent	String	否	发自定义通知时，强制推送文案，最长500个字符
//forcePushAll	String	否	发群自定义通知时，强推列表是否为群里除发送者外的所有有效成员
//forcePushList	String	否	发群自定义通知时，强推列表，格式为JSONArray，如"accid1","accid2"
type AttachMsg struct {
	FromAccid        string   `json:"fromAccid" `
	ToAccids         []string `json:"toAccids" `
	From             string   `json:"from" `
	Msgtype          int      `json:"msgtype"`
	To               string   `json:"to" `
	Attach           string   `json:"attach" `
	Pushcontent      string   `json:"pushcontent" `
	Payload          string   `json:"payload" `
	Sound            string   `json:"sound" `
	Save             int      `json:"save" `
	Option           string   `json:"option" `
	IsForcePush      string   `json:"isForcePush" `
	ForcePushContent string   `json:"forcePushContent" `
	ForcePushAll     string   `json:"forcePushAll" `
	ForcePushList    string   `json:"forcePushList" `
}

//发送自定义系统通知
//1.自定义系统通知区别于普通消息，方便开发者进行业务逻辑的通知；
//2.目前支持两种类型：点对点类型和群类型（仅限高级群），根据msgType有所区别。
//应用场景：如某个用户给另一个用户发送好友请求信息等，具体attach为请求消息体，第三方可以自行扩展，建议是json格式
func (n NetEaseIM) SendAttachMsg(msg AttachMsg) error {
	b, err := n.request(path.SendAttachMsg, structToMap(msg))
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

//1.系统通知区别于普通消息，应用接收到直接交给上层处理，客户端可不做展示；
//2.目前支持类型：点对点类型；
//3.最大限500人，只能针对个人,如果批量提供的帐号中有未注册的帐号，会提示并返回给用户；
//4.此接口受频率控制，一个应用一分钟最多调用120次，超过会返回416状态码，并且被屏蔽一段时间；
func (n NetEaseIM) SendBatchAttachMsg(msg AttachMsg) error {
	b, err := n.request(path.SendBatchAttachMsg, structToMap(msg))
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

//content	String	是	字符流base64串(Base64.encode(bytes)) ，最大15M的字符流
//type	String	否	上传文件类型
//ishttps	String	否	返回的url是否需要为https的url，true或false，默认false
//expireSec	Integer	否	文件过期时长，单位：秒，必须大于等于86400
//tag	String	否	文件的应用场景，不超过32个字符
type File struct {
	Content   string `json:"content" `
	FileType  string `json:"type" `
	Ishttps   string `json:"ishttps"`
	ExpireSec int    `json:"expireSec"`
	Tag       string `json:"tag"`
}

//文件上传，字符流需要base64编码，最大15M。
func (n NetEaseIM) Upload(file File) (string, error) {
	b, err := n.request(path.Upload, structToMap(file))
	if err != nil {
		return "", err
	}
	var resp struct {
		Code int    `json:"code"`
		Url  string `json:"url" `
	}

	if err = json.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return "", errors.New(s)
		}
	}

	return resp.Url, nil
}

//startTime	Long	是	被清理文件的开始时间(毫秒级)
//endTime	Long	是	被清理文件的结束时间(毫秒级)
//contentType	String	否	被清理的文件类型，文件类型包含contentType则被清理 如原始文件类型为"image/png"，contentType参数为"image",则满足被清理条件
//tag	String	否	被清理文件的应用场景，完全相同才被清理 如上传文件时知道场景为"usericon",tag参数为"usericon"，则满足被清理条件
type Nos struct {
	StartTime   string `json:"startTime" `
	EndTime     string `json:"endTime" `
	ContentType string `json:"contentType" `
	Tag         string `json:"tag"`
}

//上传NOS文件清理任务，按时间范围和文件类下、场景清理符合条件的文件
//每天提交的任务数量有限制，请合理规划
//关于startTime与endTime请注意：
//startTime必须小于endTime且大于0，endTime和startTime差值在1天以上，7天以内。
//endTime必须早于今天（即只可以清理今天以前的文件
func (n NetEaseIM) CleanNOS(nos Nos) (string, error) {
	b, err := n.request(path.NOSClean, structToMap(nos))
	if err != nil {
		return "", err
	}
	var resp struct {
		Code int `json:"code"`
		Data struct {
			Taskid string `json:"taskid" `
		} `json:"data" `
	}

	if err = json.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	if resp.Code != 200 {
		s, ok := stateCode[resp.Code]
		if ok {
			return "", errors.New(s)
		}
	}

	return resp.Data.Taskid, nil
}
