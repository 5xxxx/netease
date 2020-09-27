/*
 *
 * path.go
 * path
 *
 * Created by lintao on 2020/6/9 10:08 上午
 * Copyright © 2020-2020 LINTAO. All rights reserved.
 *
 */

package path

type Path string

const (
	CreateAccount  Path = "user/create.action"
	UpdateAccount  Path = "user/update.action"
	BlockAccount   Path = "user/block.action"
	UnblockAccount Path = "user/unblock.action"
	RefreshToken   Path = "user/refreshToken.action"
)

const (
	UpdateUinfo Path = "user/updateUinfo.action"
	GetUinfos   Path = "user/getUinfos.action"
)

const (
	SetDonnop Path = "user/setDonnop.action"
	Mute      Path = "user/mute.action"
	MuteAv    Path = "user/muteAv.action"
)

const (
	AddFriend            Path = "friend/add.action"
	UpdateFriend         Path = "friend/update.action"
	DeleteFriend         Path = "friend/delete.action "
	GetFriend            Path = "friend/get.action"
	SetSpecialRelation   Path = "user/setSpecialRelation.action"
	ListBlackAndMuteList Path = "user/listBlackAndMuteList.action"
)

const (
	SendMsg            Path = "msg/sendMsg.action"
	SendBatchMsg       Path = "msg/sendBatchMsg.action"
	SendAttachMsg      Path = "msg/sendAttachMsg.action"
	SendBatchAttachMsg Path = "msg/sendBatchAttachMsg.action"
	Upload             Path = "msg/upload.action"
	FileUpload         Path = "msg/fileUpload.action"
	NOSClean           Path = "job/nos/del.action"
	Recall             Path = "msg/recall.action"
)

const (
	ChatRoomCreate   Path = "chatroom/create.action"
	GetChatRoom      Path = "chatroom/get.action"
	GetBatchChatRoom Path = "chatroom/getBatch.action "
	UpdateChatRoom   Path = "chatroom/update.action"
	ToggleCloseStat  Path = "chatroom/toggleCloseStat.action"
	SetMemberRole    Path = "chatroom/setMemberRole.action"
	RequestAddr      Path = "chatroom/requestAddr.action"
	ChatRoomSendMsg  Path = "chatroom/sendMsg.action"
	AddRobot         Path = "chatroom/addRobot.action"
	RemoveRobot      Path = "chatroom/removeRobot.action"
)
