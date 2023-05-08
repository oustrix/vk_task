package botapi

type Repost struct {
	Count        int `json:"count"`
	WallCount    int `json:"wall_count"`
	MailCount    int `json:"mail_count"`
	UserReposted int `json:"user_reposted"`
}

type NewMessage struct {
	Message    Message    `json:"message"`
	ClientInfo ClientInfo `json:"client_info"`
}

type Message struct {
	ID                    int          `json:"id"`
	Date                  int          `json:"date"`
	PeerID                int          `json:"peer_id"`
	FromID                int          `json:"from_id"`
	Text                  string       `json:"text"`
	RandomID              int          `json:"random_id"`
	Ref                   string       `json:"ref"`
	RefSource             string       `json:"ref_source"`
	Out                   int          `json:"out"`
	Attachments           []Attachment `json:"attachments"`
	Important             bool         `json:"important"`
	Geo                   Geo          `json:"geo"`
	Payload               string       `json:"payload"`
	Keyboard              Keyboard     `json:"keyboard"`
	FwdMessages           []Message    `json:"fwd_messages"`
	ReplyMessage          *Message     `json:"reply_message"`
	Action                string       `json:"action"`
	AdminAuthorID         int          `json:"admin_author_id"`
	ConversationMessageID int          `json:"conversation_message_id"`
	IsCropped             bool         `json:"is_cropped"`
	MembersCount          int          `json:"members_count"`
	WasListened           bool         `json:"was_listened"`
	PinnedAt              int          `json:"pinned_at"`
	MessageTag            string       `json:"message_tag"`
	IsHidden              bool         `json:"is_hidden"`
}

type Attachment struct {
	Type        string      `json:"type"`
	Photo       Photo       `json:"photo"`
	Video       Video       `json:"video"`
	Audio       Audio       `json:"audio"`
	Doc         Doc         `json:"doc"`
	Link        Link        `json:"link"`
	Market      Market      `json:"market"`
	MarketAlbum MarketAlbum `json:"market_album"`
	Wall        Wall        `json:"wall"`
	WallReply   WallReply   `json:"wall_reply"`
	Sticker     Sticker     `json:"sticker"`
	Gift        Gift        `json:"gift"`
}

type Photo struct {
	ID      int         `json:"id"`
	AlbumID int         `json:"album_id"`
	OwnerID int         `json:"owner_id"`
	UserID  int         `json:"user_id"`
	Text    string      `json:"text"`
	Date    int         `json:"date"`
	Sizes   []PhotoSize `json:"sizes"`
	Widht   int         `json:"width"`
	Height  int         `json:"height"`
}

type PhotoSize struct {
	Type   string `json:"type"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Video struct {
	ID            int               `json:"id"`
	OwnerID       int               `json:"owner_id"`
	Title         string            `json:"title"`
	Description   string            `json:"description"`
	Duration      int               `json:"duration"`
	Image         []VideoCover      `json:"image"`
	FirstFrame    []VideoFirstFrame `json:"first_frame"`
	Date          int               `json:"date"`
	Views         int               `json:"views"`
	LocalViews    int               `json:"local_views"`
	Comments      int               `json:"comments"`
	Player        string            `json:"player"`
	Platform      string            `json:"platform"`
	CanAdd        int               `json:"can_add"`
	IsPrivate     int               `json:"is_private"`
	AccessKey     string            `json:"access_key"`
	Processing    int               `json:"processing"`
	IsFavorite    bool              `json:"is_favorite"`
	CanComment    int               `json:"can_comment"`
	CanEdit       int               `json:"can_edit"`
	CanLike       int               `json:"can_like"`
	CanRepost     int               `json:"can_repost"`
	CanSubscribe  int               `json:"can_subscribe"`
	CanAddToFaves int               `json:"can_add_to_faves"`
	CanAttachLink int               `json:"can_attach_link"`
	Width         int               `json:"width"`
	Height        int               `json:"height"`
	UserID        int               `json:"user_id"`
	Converting    int               `json:"converting"`
	Added         int               `json:"added"`
	IsSubscribed  int               `json:"is_subscribed"`
	Repeat        int               `json:"repeat"`
	Type          string            `json:"type"`
	Balance       int               `json:"balance"`
	LiveStatus    string            `json:"live_status"`
	Live          int               `json:"live"`
	Upcoming      int               `json:"upcoming"`
	Likes         Likes             `json:"likes"`
	Reposts       Repost            `json:"reposts"`
}

type VideoCover struct {
	Height      int    `json:"height"`
	URL         string `json:"url"`
	Width       int    `json:"width"`
	WithPadding int    `json:"with_padding"`
}

type VideoFirstFrame struct {
	Height int    `json:"height"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
}

type Likes struct {
	Count     int `json:"count"`
	UserLikes int `json:"user_likes"`
}

type Audio struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Artist   string `json:"artist"`
	Title    string `json:"title"`
	Duration int    `json:"duration"`
	URL      string `json:"url"`
	LyricsID int    `json:"lyrics_id"`
	AlbumID  int    `json:"album_id"`
	GenreID  int    `json:"genre_id"`
	Date     int    `json:"date"`
	NoSearch int    `json:"no_search"`
	IsHq     int    `json:"is_hq"`
}

type Doc struct {
	ID      int    `json:"id"`
	OwnerID int    `json:"owner_id"`
	Title   string `json:"title"`
	Size    int    `json:"size"`
	Ext     string `json:"ext"`
	URL     string `json:"url"`
	Date    int    `json:"date"`
	Type    int    `json:"type"`
}

type DocPreview struct {
	Photo DocPreviewPhoto `json:"photo"`
}

type DocPreviewPhoto struct {
	Sizes        []PhotoSize  `json:"sizes"`
	Graffiti     Graffiti     `json:"graffiti"`
	AudioMessage AudioMessage `json:"audio_message"`
}

type Graffiti struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type AudioMessage struct {
	Duration int    `json:"duration"`
	Waveform []int  `json:"waveform"`
	LinkOgg  string `json:"link_ogg"`
	LinkMp3  string `json:"link_mp3"`
}

type Link struct {
	URL         string  `json:"url"`
	Title       string  `json:"title"`
	Caption     string  `json:"caption"`
	Description string  `json:"description"`
	Photo       Photo   `json:"photo"`
	Product     Product `json:"product"`
	Button      Button  `json:"button"`
	PreviewPage string  `json:"preview_page"`
	PreviewURL  string  `json:"preview_url"`
}

type Product struct {
	Price ProductPrice `json:"price"`
}

type ProductPrice struct {
	Amount   string        `json:"amount"`
	Currency PriceCurrency `json:"currency"`
	Text     string        `json:"text"`
}

type PriceCurrency struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Button struct {
	Title  string       `json:"title"`
	Action ButtonAction `json:"action"`
}

type ButtonAction struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Market struct {
	ID           int              `json:"id"`
	OwnerID      int              `json:"owner_id"`
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Price        MarketPrice      `json:"price"`
	Dimensions   Dimensions       `json:"dimensions"`
	Weight       int              `json:"weight"`
	Category     MarketCategory   `json:"category"`
	ThumbPhoto   string           `json:"thumb_photo"`
	Date         int              `json:"date"`
	Availability int              `json:"availability"`
	IsFavorite   bool             `json:"is_favorite"`
	Sku          string           `json:"sku"`
	RejectInfo   MarketRejectInfo `json:"reject_info"`

	// extended
	Photos      []Photo `json:"photos"`
	CanComment  int     `json:"can_comment"`
	CanRepost   int     `json:"can_repost"`
	Likes       Likes   `json:"likes"`
	URL         string  `json:"url"`
	ButtonTitle string  `json:"button_title"`
}

type MarketPrice struct {
	ProductPrice
	OldAmount string `json:"old_amount"`
}

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Length int `json:"length"`
}

type MarketCategory struct {
	ID      int                   `json:"id"`
	Name    string                `json:"name"`
	Section MarketCategorySection `json:"section"`
}

type MarketCategorySection struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MarketRejectInfo struct {
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Buttons            []Button `json:"buttons"`
	ModerationStatus   int      `json:"moderation_status"`
	InfoLink           string   `json:"info_link"`
	WhiteToSupportLink string   `json:"write_to_support_link"`
}

type MarketAlbum struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Title    string `json:"title"`
	IsMain   bool   `json:"is_main"`
	IsHidden bool   `json:"is_hidden"`
	Photo    Photo  `json:"photo"`
	Count    int    `json:"count"`
}

type Wall struct {
	ID           int              `json:"id"`
	OwnerID      int              `json:"owner_id"`
	FromID       int              `json:"from_id"`
	CreatedBy    int              `json:"created_by"`
	Date         int              `json:"date"`
	Text         string           `json:"text"`
	ReplyOwnerID int              `json:"reply_owner_id"`
	ReplyPostID  int              `json:"reply_post_id"`
	FriendsOnly  int              `json:"friends_only"`
	Comments     WallComments     `json:"comments"`
	Copyright    WallCopyright    `json:"copyright"`
	Likes        WallLikes        `json:"likes"`
	Reposts      WallReposts      `json:"reposts"`
	Views        WallViews        `json:"views"`
	PostType     string           `json:"post_type"`
	PostSource   PostSource       `json:"post_source"`
	Attachments  []WallAttachment `json:"attachments"`
	Geo          Geo              `json:"geo"`
	SignerID     int              `json:"signer_id"`
	CopyHistory  []Wall           `json:"copy_history"`
	CanPin       int              `json:"can_pin"`
	CanDelete    int              `json:"can_delete"`
	CanEdit      int              `json:"can_edit"`
	IsPinned     int              `json:"is_pinned"`
	MarkedAsAds  int              `json:"marked_as_ads"`
	IsFavorite   bool             `json:"is_favorite"`
	Donut        Donut            `json:"donut"`
	PostponedID  int              `json:"postponed_id"`
}

type WallComments struct {
	Count         int  `json:"count"`
	CanPost       int  `json:"can_post"`
	GroupsCanPost bool `json:"groups_can_post"`
	CanClose      bool `json:"can_close"`
	CanOpen       bool `json:"can_open"`
}

type WallCopyright struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type WallLikes struct {
	Likes
	CanLike    int `json:"can_like"`
	CanPublish int `json:"can_publish"`
}

type WallReposts struct {
	Count        int `json:"count"`
	UserReposted int `json:"user_reposted"`
}

type WallViews struct {
	Count int `json:"count"`
}

type PostSource struct {
	Type     string `json:"type"`
	Platform string `json:"platform"`
	Data     string `json:"data"`
	URL      string `json:"url"`
}

type WallAttachment struct {
	Type        string       `json:"type"`
	Photo       Photo        `json:"photo"`
	PostedPhoto PostedPhoto  `json:"posted_photo"`
	Video       Video        `json:"video"`
	Audio       Audio        `json:"audio"`
	Doc         Doc          `json:"doc"`
	Graffiti    WallGraffiti `json:"graffiti"`
	Link        Link         `json:"link"`
	Note        Note         `json:"note"`
	App         WallApp      `json:"app"`
	Poll        Poll         `json:"poll"`
	Page        Page         `json:"page"`
	Album       WallAlbum    `json:"album"`
	PhotosList  []string     `json:"photos_list"`
	Market      Market       `json:"market"`
	MarketAlbum MarketAlbum  `json:"market_album"`
	Sticker     Sticker      `json:"sticker"`
	PrettyCards []PrettyCard `json:"pretty_cards"`
	Event       Event        `json:"event"`
}

type PostedPhoto struct {
	ID       int    `json:"id"`
	OwnerID  int    `json:"owner_id"`
	Photo130 string `json:"photo_130"`
	Photo604 string `json:"photo_604"`
}

type WallGraffiti struct {
	PostedPhoto
}

type Note struct {
	ID           int    `json:"id"`
	OwnerID      int    `json:"owner_id"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	Date         int    `json:"date"`
	Comments     int    `json:"comments"`
	ReadComments int    `json:"read_comments"`
	ViewURL      string `json:"view_url"`
	PrivacyView  string `json:"privacy_view"`
	CanComment   int    `json:"can_comment"`
	TextWiki     string `json:"text_wiki"`
}

type WallApp struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Photo130 string `json:"photo_130"`
	Photo604 string `json:"photo_604"`
}

type Poll struct {
	ID         int            `json:"id"`
	OwnerID    int            `json:"owner_id"`
	Created    int            `json:"created"`
	Question   string         `json:"question"`
	Votes      int            `json:"votes"`
	Answers    []PollAnswer   `json:"answers"`
	Anonymous  bool           `json:"anonymous"`
	Multiple   bool           `json:"multiple"`
	AnswerIDS  []int          `json:"answer_ids"`
	EndDate    int            `json:"end_date"`
	Closed     bool           `json:"closed"`
	IsBoard    bool           `json:"is_board"`
	CanEdit    bool           `json:"can_edit"`
	CanVote    bool           `json:"can_vote"`
	CanReport  bool           `json:"can_report"`
	CanShare   bool           `json:"can_share"`
	AuthorID   int            `json:"author_id"`
	Photo      Photo          `json:"photo"`
	Background PollBackground `json:"background"`
	Friends    []int          `json:"friends"`
}

type PollAnswer struct {
	ID    int     `json:"id"`
	Text  string  `json:"text"`
	Votes int     `json:"votes"`
	Rate  float64 `json:"rate"`
}

type PollBackground struct {
	ID     int                   `json:"id"`
	Type   string                `json:"type"`
	Angle  int                   `json:"angle"`
	Color  string                `json:"color"`
	Width  int                   `json:"width"`
	Height int                   `json:"height"`
	Images []Photo               `json:"images"`
	Points []PollBackgroundPoint `json:"points"`
}

type PollBackgroundPoint struct {
	Position float64 `json:"position"`
	Color    string  `json:"color"`
}

type Page struct {
	ID                       int    `json:"id"`
	GroupID                  int    `json:"group_id"`
	CreatorID                int    `json:"creator_id"`
	Title                    int    `json:"title"`
	CurrentUserCanEdit       int    `json:"current_user_can_edit"`
	CurrentUserCanEditAccess int    `json:"current_user_can_edit_access"`
	WhoCanView               int    `json:"who_can_view"`
	Edited                   int    `json:"edited"`
	Created                  int    `json:"created"`
	EditorID                 int    `json:"editor_id"`
	Views                    int    `json:"views"`
	Parent                   string `json:"parent"`
	Parent2                  string `json:"parent2"`
	Source                   string `json:"source"`
	HTML                     string `json:"html"`
	ViewURL                  string `json:"view_url"`
}

type WallAlbum struct {
	ID          int    `json:"id"`
	Thumb       Photo  `json:"thumb"`
	Owner       int    `json:"owner"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Created     int    `json:"created"`
	Updated     int    `json:"updated"`
	Size        int    `json:"size"`
}

type PrettyCard struct {
	CardID   string            `json:"card_id"`
	LinkURL  string            `json:"link_url"`
	Title    string            `json:"title"`
	Images   []PrettyCardImage `json:"images"`
	Button   Button            `json:"button"`
	Price    string            `json:"price"`
	PriceOld string            `json:"price_old"`
}

type PrettyCardImage struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Event struct {
	ID           int    `json:"id"`
	Time         int    `json:"time"`
	MemberStatus int    `json:"member_status"`
	IsFavorite   int    `json:"is_favorite"`
	Address      string `json:"address"`
	Text         string `json:"text"`
	ButtonText   string `json:"button_text"`
	Friends      []int  `json:"friends"`
}

type Donut struct {
	IsDonut            bool                   `json:"is_donut"`
	PaidDuration       int                    `json:"paid_duration"`
	Placeholder        map[string]interface{} `json:"placeholder"` // there is no documentation for this field
	CanPublishFreeCopy bool                   `json:"can_publish_free_copy"`
	EditMode           string                 `json:"edit_mode"`
}

type WallReply struct {
	Comment
	PostID  int `json:"post_id"`
	OwnerID int `json:"owner_id"`
}

type Comment struct {
	ID             int           `json:"id"`
	FromID         int           `json:"from_id"`
	Date           int           `json:"date"`
	Text           int           `json:"text"`
	Donut          CommentDonut  `json:"donut"`
	ReplyToUser    int           `json:"reply_to_user"`
	ReplyToComment int           `json:"reply_to_comment"`
	Attachments    []Attachment  `json:"attachments"`
	ParentsStack   []int         `json:"parents_stack"`
	Thread         CommentThread `json:"thread"`
}

type CommentDonut struct {
	IsDonut     bool   `json:"is_donut"`
	Placeholder string `json:"placeholder"`
}

type CommentThread struct {
	Count           int       `json:"count"`
	Items           []Comment `json:"items"`
	CanPost         bool      `json:"can_post"`
	ShowReplyButton bool      `json:"show_reply_button"`
	GroupsCanPost   bool      `json:"groups_can_post"`
}

type Sticker struct {
	ProductID            int            `json:"product_id"`
	StickerID            int            `json:"sticker_id"`
	Images               []StickerImage `json:"images"`
	ImagesWithBackground []StickerImage `json:"images_with_background"`
	AnimationURL         string         `json:"animation_url"`
	IsAllowed            bool           `json:"is_allowed"`
}

type StickerImage struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Gift struct {
	ID       int    `json:"id"`
	Thumb256 string `json:"thumb_256"`
	Thumb96  string `json:"thumb_96"`
	Thumb48  string `json:"thumb_48"`
}

type Geo struct {
	Type        string           `json:"type"`
	Coordinates []GeoCoordinates `json:"coordinates"`
	Place       GeoPlace         `json:"place"`
	Showmap     int              `json:"showmap"`
}

type GeoCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type GeoPlace struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Created   int     `json:"created"`
	Icon      string  `json:"icon"`
	Country   string  `json:"country"`
	City      string  `json:"city"`
}

type Keyboard struct {
	Type    string `json:"type"`
	Label   string `json:"label"`
	Payload string `json:"payload"`
	Link    string `json:"link"`
	Hash    string `json:"hash"`
	AppID   int    `json:"app_id"`
	OwnerID int    `json:"owner_id"`
}

type ClientInfo struct {
	ButtonActions  []string `json:"button_actions"`
	Keyboard       bool     `json:"keyboard"`
	InlineKeyboard bool     `json:"inline_keyboard"`
	Carousel       bool     `json:"carousel"`
	LangID         int      `json:"lang_id"`
}
