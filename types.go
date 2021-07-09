package goteleg

type Location struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

type ChatLocation struct {
	Location *Location `json:"location"`
	Addres   string    `json:"address"`
}

type ChatPermisstions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_message,omitempty"`
}

type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	Username                string `json:"username,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supprot_inline_queries,omitempty"`
}

type Chat struct {
	ID                    int64             `json:"id"`
	Type                  string            `json:"type"`
	Title                 string            `json:"title,omitempty"`
	Username              string            `json:"username,omitempty"`
	FirstName             string            `json:"first_name,omitempty"`
	LastName              string            `json:"last_name,omitempty"`
	Photo                 *ChatPhoto        `json:"photo,omitempty"`
	Bio                   string            `json:"bio,omitempty"`
	Description           string            `json:"description,omitempty"`
	InviteLink            string            `json:"invite_link,omitempty"`
	PinnedMessage         *Message          `json:"pinned_message,omitempty"`
	Premissions           *ChatPermisstions `json:"premission,omitempty"`
	SlowModeDelay         int               `json:"slow_mode_delay,omitempty"`
	MessageAutoDaleteTime int               `json:"message_auto_delete_time,omitempty"`
	StikerSetName         string            `json:"stiker_set_name,omitempty"`
	LinkedChatID          int               `json:"linked_char_id,omitempty"`
	Location              ChatLocation      `json:"location,omitempty"`
}

//Represent one special entity in a text message.
type MessageEntity struct {
	Type     string `json:"type"`
	Offset   string `json:"offset"`
	Length   int    `json:"length"`
	URL      string `json:"url,omitempty"`
	User     *User  `json:"user,omitempty"`
	Language string `json:"language,omitempty"`
}

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"`
}

type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
}
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

type Stiker struct {
	FileID       string        `json:"file_id"`
	FileUniqueID string        `json:"file_unique_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
	Emoji        string        `json:"emoji,omitempty"`
	SetName      string        `json:"set_name,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	FileSize     int           `json:"file_size,omitempty"`
}

type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size"`
}

type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int    `json:"file_size,omitempty"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []*PhotoSize     `json:"photo_size"`
	Text         string           `json:"text,omitempty"`
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation       `json:"animation,omitempty"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type Poll struct {
	ID                    string           `json:"id"`
	Question              string           `json:"question"`
	Options               []*PollOption    `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"type"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionID       int              `json:"correct_option_id,omitempty"`
	Explanation           string           `json:"explanation,omitempty"`
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int              `json:"open_preiod,omitempty"`
	CloseDate             int              `json:"close_date,omitempty"`
}

type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	FoursquareID    string    `json:"foursquare_id,omitempty"`
	FoursquareType  string    `json:"foursquare_type,omitempty"`
	GooglePalceID   string    `json:"google_place_id,omitempty"`
	GooglePalceType string    `json:"google_place_type,omitempty"`
}

type MessageAutoDaleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type SuccessfulPayment struct {
	Currency         string `json:"currency"`
	TotalAmount      int    `json:"total_amount"`
	InvoicePayload   string `json:"invoice_payload"`
	ShippingOptionID string `json:"shipping_option_id,omitempty"`
	OrderInfo
}

type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        string          `json:"data,omitempty"`
	PhoneNumber string          `json:"phone_number,omitempty"`
	Email       string          `json:"email,omitempty"`
	Files       []*PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile   `json:"front_size,omitempty"`
	ReverseSide *PassportFile   `json:"reverse_side,omitempty"`
	Selfie      *PassportFile   `json:"slfie,omitempty"`
	Translation []*PassportFile `json:"translation,omitempty"`
	Hash        string          `json:"hash"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials       `json:"credentials"`
}

type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int   `json:"distance"`
}

type VoiceChatScheduler struct {
	StartDate int `json:"start_date"`
}

type VoiceChatStarted struct{}

type VoiceChatEnded struct {
	Duration int `json:"duration"`
}

type VoiceChatParticipantsInvited struct {
	Users []*User `json:"users,omitempty"`
}

type InlineButton struct {
	Text         string `json:"text"`
	URL          string `json:"url,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}

type InlineKbdRow []InlineButton

type InlineKeyboard struct {
	InlineKeyboard []InlineKbdRow `json:"inline_keyboard"`
}

type Message struct {
	MessageID                     int                            `json:"message_id"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	Date                          int                            `json:"date"`
	Chat                          *Chat                          `json:"chat"`
	ForwardFrom                   *User                          `json:"forward_from,omitempty"`
	ForwardFromChat               *Chat                          `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID          int                            `json:"forward_from_message_id,omitempty"`
	ForwardSignature              string                         `json:"forward_signature,omitempty"`
	ForwardSenderName             string                         `json:"forward_sender_name,omitempty"`
	ForwardDate                   int                            `json:"forward_date,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      int                            `json:"edit_date,omitempty"`
	MediaGroupID                  string                         `json:"media_group_id,omitempty"`
	AuthorSignature               string                         `json:"author_signature,omitempty"`
	Text                          string                         `json:"text,omitempty"`
	Entities                      []*MessageEntity               `json:"entities,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Auido                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"doucment,omitempty"`
	Photo                         []*PhotoSize                   `json:"photo,omitempty"`
	Stiker                        *Stiker                        `json:"stiker,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       string                         `json:"caption,omitempty"`
	CaptionEntities               []*MessageEntity               `json:"caption_entities,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                []*User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	SuperGroupChatCreated         bool                           `json:"super_group_chat_created,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDaleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID               int64                          `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID             int64                          `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	VoiceChatScheduler            *VoiceChatScheduler            `json:"voice_chat_scheduler,omitempty"`
	VoiceChatStarted              *VoiceChatStarted              `json:"voice_chatstarted,omitempty"`
	VoiceChatEnded                *VoiceChatEnded                `json:"voice_chat_ended,omitempty"`
	VoiceChatParticipantsInvited  *VoiceChatParticipantsInvited  `json:"voice_chat_participants_invited,omitempty"`
	ReplyMarkup                   *InlineKeyboard                `json:"reply_markup,omitempty"`
}

type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType string    `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance,omitempty"`
	Data            string   `json:"data"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}

type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

type Update struct {
	UpdateID           int                 `json:"update_id"`
	Msg                *Message            `json:"message,omitempty"`
	EditedMsg          *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"call_back_query,omitempty"`
}

type Responce struct {
	OK          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
}

type MessageResponce struct {
	Responce
	Result *Message `json:"result,omitempty"`
}

type UpdatesResponce struct {
	Responce
	Result []*Update `json:"result,omitempty"`
}
