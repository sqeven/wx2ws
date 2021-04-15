package msg

// Go to 718
type GoMsg struct {
	Code    int       `json:"Code"`
	Success bool      `json:"Success"`
	Message string    `json:"Message"`
	Data    GoMsgData `json:"Data"`
}

type GoMsgUserName struct {
	String string `json:"string"`
}

type GoMsgNickName struct {
	String string `json:"string"`
}

type GoMsgBindEmail struct {
}

type GoMsgBindMobile struct {
	String string `json:"string"`
}

type GoMsgNightTime struct {
	BeginTime int `json:"BeginTime"`
	EndTime   int `json:"EndTime"`
}

type GoMsgAllDayTim struct {
	BeginTime int `json:"BeginTime"`
	EndTime   int `json:"EndTime"`
}

type GoMsgDisturbSetting struct {
	NightSetting  int            `json:"NightSetting"`
	NightTime     GoMsgNightTime `json:"NightTime"`
	AllDaySetting int            `json:"AllDaySetting"`
	AllDayTim     GoMsgAllDayTim `json:"AllDayTim"`
}

type GoMsgGmailList struct {
	Count int `json:"Count"`
}

type GoMsgModUserInfos struct {
	BitFlag        int                 `json:"BitFlag"`
	UserName       GoMsgUserName       `json:"UserName"`
	NickName       GoMsgNickName       `json:"NickName"`
	BindUin        int                 `json:"BindUin"`
	BindEmail      GoMsgBindEmail      `json:"BindEmail"`
	BindMobile     GoMsgBindMobile     `json:"BindMobile"`
	Status         int                 `json:"Status"`
	ImgLen         int                 `json:"ImgLen"`
	Sex            int                 `json:"Sex"`
	Province       string              `json:"Province"`
	City           string              `json:"City"`
	PersonalCard   int                 `json:"PersonalCard"`
	DisturbSetting GoMsgDisturbSetting `json:"DisturbSetting"`
	PluginFlag     int                 `json:"PluginFlag"`
	VerifyFlag     int                 `json:"VerifyFlag"`
	Point          int                 `json:"Point"`
	Experience     int                 `json:"Experience"`
	Level          int                 `json:"Level"`
	LevelLowExp    int                 `json:"LevelLowExp"`
	LevelHighExp   int                 `json:"LevelHighExp"`
	PluginSwitch   int                 `json:"PluginSwitch"`
	GmailList      GoMsgGmailList      `json:"GmailList"`
	Alias          string              `json:"Alias"`
	WeiboFlag      int                 `json:"WeiboFlag"`
	FaceBookFlag   int                 `json:"FaceBookFlag"`
	FbuserID       int                 `json:"FbuserId"`
	AlbumStyle     int                 `json:"AlbumStyle"`
	AlbumFlag      int                 `json:"AlbumFlag"`
	TxnewsCategory int                 `json:"TxnewsCategory"`
	Country        string              `json:"Country"`
}

type GoMsgUserName0 struct {
	String string `json:"string"`
}

type GoMsgNickName0 struct {
	String string `json:"string"`
}

type GoMsgPyInitial struct {
	String string `json:"string"`
}

type GoMsgQuanPin struct {
	String string `json:"string"`
}

type GoMsgImgBuf struct {
	ILen int `json:"iLen"`
}

type GoMsgRemark struct {
}

type GoMsgRemarkPyinitial struct {
}

type GoMsgRemarkQuanPin struct {
}

type GoMsgDomainList struct {
}

type GoMsgSnsUserInfo struct {
	SnsFlag       int `json:"SnsFlag"`
	SnsBgobjectID int `json:"SnsBgobjectId"`
	SnsFlagEx     int `json:"SnsFlagEx"`
}

type GoMsgCustomizedInfo struct {
	BrandFlag int `json:"BrandFlag"`
}

type GoMsgLinkedinContactItem struct {
}

type GoMsgAdditionalContactList struct {
	LinkedinContactItem GoMsgLinkedinContactItem `json:"LinkedinContactItem"`
}

type GoMsgNewChatroomData struct {
}

type GoMsgModContacts struct {
	UserName              GoMsgUserName0             `json:"UserName"`
	NickName              GoMsgNickName0             `json:"NickName"`
	PyInitial             GoMsgPyInitial             `json:"PyInitial"`
	QuanPin               GoMsgQuanPin               `json:"QuanPin"`
	Sex                   int                        `json:"Sex"`
	ImgBuf                GoMsgImgBuf                `json:"ImgBuf"`
	BitMask               int64                      `json:"BitMask"`
	BitVal                int                        `json:"BitVal"`
	ImgFlag               int                        `json:"ImgFlag"`
	Remark                GoMsgRemark                `json:"Remark"`
	RemarkPyinitial       GoMsgRemarkPyinitial       `json:"RemarkPyinitial"`
	RemarkQuanPin         GoMsgRemarkQuanPin         `json:"RemarkQuanPin"`
	ContactType           int                        `json:"ContactType"`
	RoomInfoCount         int                        `json:"RoomInfoCount"`
	DomainList            []GoMsgDomainList          `json:"DomainList"`
	ChatRoomNotify        int                        `json:"ChatRoomNotify"`
	AddContactScene       int                        `json:"AddContactScene"`
	PersonalCard          int                        `json:"PersonalCard"`
	HasWeiXinHdHeadImg    int                        `json:"HasWeiXinHdHeadImg"`
	VerifyFlag            int                        `json:"VerifyFlag"`
	Level                 int                        `json:"Level"`
	Source                int                        `json:"Source"`
	WeiboFlag             int                        `json:"WeiboFlag"`
	AlbumStyle            int                        `json:"AlbumStyle"`
	AlbumFlag             int                        `json:"AlbumFlag"`
	SnsUserInfo           GoMsgSnsUserInfo           `json:"SnsUserInfo"`
	BigHeadImgURL         string                     `json:"BigHeadImgUrl"`
	SmallHeadImgURL       string                     `json:"SmallHeadImgUrl"`
	CustomizedInfo        GoMsgCustomizedInfo        `json:"CustomizedInfo"`
	EncryptUserName       string                     `json:"EncryptUserName"`
	AdditionalContactList GoMsgAdditionalContactList `json:"AdditionalContactList"`
	ChatroomMaxCount      int                        `json:"ChatroomMaxCount"`
	NewChatroomData       GoMsgNewChatroomData       `json:"NewChatroomData"`
	DeleteFlag            int                        `json:"DeleteFlag"`
	Description           string                     `json:"Description"`
	ChatroomStatus        int                        `json:"ChatroomStatus"`
	Extflag               int                        `json:"Extflag"`
	ChatRoomBusinessType  int                        `json:"ChatRoomBusinessType"`
	Province              string                     `json:"Province"`
	City                  string                     `json:"City"`
	Signature             string                     `json:"Signature"`
	VerifyInfo            string                     `json:"VerifyInfo"`
	VerifyContent         string                     `json:"VerifyContent"`
	Alias                 string                     `json:"Alias"`
	Country               string                     `json:"Country"`
	AlbumBGImgID          string                     `json:"AlbumBGImgID"`
	ChatRoomOwner         string                     `json:"ChatRoomOwner"`
}

type GoMsgUserName1 struct {
	String string `json:"string"`
}

type GoMsgDelContacts struct {
	UserName          GoMsgUserName1 `json:"UserName"`
	DeleteContactScen int            `json:"DeleteContactScen"`
}

type GoMsgModUserImgs struct {
	ImgType         int    `json:"ImgType"`
	ImgLen          int    `json:"ImgLen"`
	ImgBuf          string `json:"ImgBuf"`
	ImgMd5          string `json:"ImgMd5"`
	BigHeadImgURL   string `json:"BigHeadImgUrl"`
	SmallHeadImgURL string `json:"SmallHeadImgUrl"`
}

type GoMsgFunctionSwitchs struct {
	FunctionID  int `json:"FunctionId"`
	SwitchValue int `json:"SwitchValue"`
}

type GoMsgSnsUserInfo0 struct {
	SnsFlag       int    `json:"SnsFlag"`
	SnsBgimgID    string `json:"SnsBgimgId"`
	SnsBgobjectID int64  `json:"SnsBgobjectId"`
	SnsFlagEx     int    `json:"SnsFlagEx"`
}

type GoMsgExtXML struct {
}

type GoMsgSafeDeviceList struct {
	Count int `json:"Count"`
}

type GoMsgLinkedinContactItem0 struct {
}

type GoMsgSign struct {
	ILen   int    `json:"iLen"`
	Buffer string `json:"buffer"`
}

type GoMsgPatternLockInfo struct {
	PatternVersion int       `json:"PatternVersion"`
	Sign           GoMsgSign `json:"Sign"`
	LockStatus     int       `json:"LockStatus"`
}

type GoMsgUserInfoExts struct {
	SnsUserInfo         GoMsgSnsUserInfo0         `json:"SnsUserInfo"`
	MyBrandList         string                    `json:"MyBrandList"`
	BigChatRoomSize     int                       `json:"BigChatRoomSize"`
	BigChatRoomQuota    int                       `json:"BigChatRoomQuota"`
	BigChatRoomInvite   int                       `json:"BigChatRoomInvite"`
	BigHeadImgURL       string                    `json:"BigHeadImgUrl"`
	SmallHeadImgURL     string                    `json:"SmallHeadImgUrl"`
	MainAcctType        int                       `json:"MainAcctType"`
	ExtXML              GoMsgExtXML               `json:"ExtXml"`
	SafeDeviceList      GoMsgSafeDeviceList       `json:"SafeDeviceList"`
	SafeDevice          int                       `json:"SafeDevice"`
	GrayscaleFlag       int                       `json:"GrayscaleFlag"`
	RegCountry          string                    `json:"RegCountry"`
	LinkedinContactItem GoMsgLinkedinContactItem0 `json:"LinkedinContactItem"`
	PatternLockInfo     GoMsgPatternLockInfo      `json:"PatternLockInfo"`
	PayWalletType       int                       `json:"PayWalletType"`
	WalletRegion        int                       `json:"WalletRegion"`
	ExtStatus           int64                     `json:"ExtStatus"`
	UserStatus          int                       `json:"UserStatus"`
}

type GoMsgFromUserName struct {
	String string `json:"string"`
}

type GoMsgToUserName struct {
	String string `json:"string"`
}

type GoMsgContent struct {
	String string `json:"string"`
}

type GoMsgImgBuf0 struct {
	ILen int `json:"iLen"`
}

type GoMsgAddMsgs struct {
	MsgID        int               `json:"MsgId"`
	FromUserName GoMsgFromUserName `json:"FromUserName"`
	ToUserName   GoMsgToUserName   `json:"ToUserName"`
	MsgType      int               `json:"MsgType"`
	Content      GoMsgContent      `json:"Content"`
	Status       int               `json:"Status"`
	ImgStatus    int               `json:"ImgStatus"`
	ImgBuf       GoMsgImgBuf0      `json:"ImgBuf"`
	CreateTime   int               `json:"CreateTime"`
	NewMsgID     interface{}       `json:"NewMsgId"`
	MsgSeq       int               `json:"MsgSeq"`
	MsgSource    string            `json:"MsgSource"`
}

type GoMsgKeyBuf struct {
	ILen   int    `json:"iLen"`
	Buffer string `json:"buffer"`
}

type GoMsgData struct {
	ModUserInfos    []GoMsgModUserInfos    `json:"ModUserInfos"`
	ModContacts     []GoMsgModContacts     `json:"ModContacts"`
	DelContacts     []GoMsgDelContacts     `json:"DelContacts"`
	ModUserImgs     []GoMsgModUserImgs     `json:"ModUserImgs"`
	FunctionSwitchs []GoMsgFunctionSwitchs `json:"FunctionSwitchs"`
	UserInfoExts    []GoMsgUserInfoExts    `json:"UserInfoExts"`
	AddMsgs         []GoMsgAddMsgs         `json:"AddMsgs"`
	ContinueFlag    int                    `json:"ContinueFlag"`
	KeyBuf          GoMsgKeyBuf            `json:"KeyBuf"`
	Status          int                    `json:"Status"`
	Continue        int                    `json:"Continue"`
	Time            int                    `json:"Time"`
	UnknownCmdID    string                 `json:"UnknownCmdId"`
	Remarks         string                 `json:"Remarks"`
}
