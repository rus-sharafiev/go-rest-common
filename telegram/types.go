package telegram

import "encoding/json"

type BasicResponse struct {
	Ok          bool             `json:"ok"`
	ErrorCode   *int             `json:"error_code,omitempty"`
	Description *string          `json:"description,omitempty"`
	Result      *json.RawMessage `json:"result"`
}

type UserResponse struct {
	Ok          bool    `json:"ok"`
	ErrorCode   *int    `json:"error_code,omitempty"`
	Description *string `json:"description,omitempty"`
	Result      *User   `json:"result"`
}

type MessageToSend struct {
	BusinessConnectionId *string             `json:"business_connection_id,omitempty"` // Optional	Unique identifier of the business connection on behalf of which the message will be sent
	ChatId               int                 `json:"chat_id"`                          // Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageThreadId      *int                `json:"message_thread_id,omitempty"`      // Optional	Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	Text                 string              `json:"text"`                             // Yes	Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode            *string             `json:"parse_mode,omitempty"`             // Optional	Mode for parsing entities in the message text. See formatting options for more details.
	Entities             *[]MessageEntity    `json:"entities,omitempty"`               // Optional	A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	LinkPreviewOptions   *LinkPreviewOptions `json:"link_preview_options,omitempty"`   // Optional	Link preview generation options for the message
	DisableNotification  *bool               `json:"disable_notification,omitempty"`   // Optional	Sends the message silently. Users will receive a notification with no sound.
	ProtectContent       *bool               `json:"protect_content,omitempty"`        // Optional	Protects the contents of the sent message from forwarding and saving
	AllowPaidBroadcast   *bool               `json:"allow_paid_broadcast,omitempty"`   // Optional	Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
	MessageEffectId      *string             `json:"message_effect_id,omitempty"`      // Optional	Unique identifier of the message effect to be added to the message; for private chats only
	ReplyMarkup          *json.RawMessage    `json:"reply_markup,omitempty"`           // Optional	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
}

// --------------------------------------------------------------------------------

// User This object represents a Telegram user or bot.
type User struct {
	Id                      int     `json:"id"`                                    // Unique identifier for this user or bot. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit int or double-precision float type are safe for storing this identifier.
	IsBot                   bool    `json:"is_bot"`                                // True, if this user is a bot
	FirstName               string  `json:"first_name"`                            // User's or bot's first name
	LastName                *string `json:"last_name,omitempty"`                   // Optional. User's or bot's last name
	Username                *string `json:"username,omitempty"`                    // Optional. User's or bot's username
	LanguageCode            *string `json:"language_code,omitempty"`               // Optional. IETF language tag of the user's language
	IsPremium               *bool   `json:"is_premium,omitempty"`                  // Optional. True, if this user is a Telegram Premium user
	AddedToAttachmentMenu   *bool   `json:"added_to_attachment_menu,omitempty"`    // Optional. True, if this user added the bot to the attachment menu
	CanJoinGroups           *bool   `json:"can_join_groups,omitempty"`             // Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanReadAllGroupMessages *bool   `json:"can_read_all_group_messages,omitempty"` // Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	SupportsInlineQueries   *bool   `json:"supports_inline_queries,omitempty"`     // Optional. True, if the bot supports inline queries. Returned only in getMe.
	CanConnectToBusiness    *bool   `json:"can_connect_to_business,omitempty"`     // Optional. True, if the bot can be connected to a Telegram Business account to receive its messages. Returned only in getMe.
	HasMainWebApp           *bool   `json:"has_main_web_app,omitempty"`            // Optional. True, if the bot has a main Web App. Returned only in getMe.
}

// Update This object represents an incoming update.
type Update struct {
	UpdateId                int                          `json:"update_id"`                           // The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially. This identifier becomes especially handy if you're using webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	Message                 *Message                     `json:"message,omitempty"`                   // Optional. New incoming message of any kind - text, photo, sticker, etc.
	EditedMessage           *Message                     `json:"edited_message,omitempty"`            // Optional. New version of a message that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
	ChannelPost             *Message                     `json:"channel_post,omitempty"`              // Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`       // Optional. New version of a channel post that is known to the bot and was edited. This update may at times be triggered by changes to message fields that are either unavailable or not actively used by your bot.
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`       // Optional. The bot was connected to or disconnected from a business account, or a user edited an existing connection with the bot
	BusinessMessage         *Message                     `json:"business_message,omitempty"`          // Optional. New message from a connected business account
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`   // Optional. New version of a message from a connected business account
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"` // Optional. Messages were deleted from a connected business account
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`          // Optional. A reaction to a message was changed by a user. The bot must be an administrator in the chat and must explicitly specify "message_reaction" in the list of allowed_updates to receive these updates. The update isn't received for reactions set by bots.
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`    // Optional. Reactions to a message with anonymous reactions were changed. The bot must be an administrator in the chat and must explicitly specify "message_reaction_count" in the list of allowed_updates to receive these updates. The updates are grouped and can be sent with delay up to a few minutes.
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`              // Optional. New incoming inline query
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`      // Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`            // Optional. New incoming callback query
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`            // Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`        // Optional. New incoming pre-checkout query. Contains full information about checkout
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`      // Optional. A user purchased paid media with a non-empty payload sent by the bot in a non-channel chat
	Poll                    *Poll                        `json:"poll,omitempty"`                      // Optional. New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`               // Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`            // Optional. The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`               // Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`         // Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`                // Optional. A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`        // Optional. A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
}

// Message This object represents a message
type Message struct {
	MessageId                     int                            `json:"message_id"`                                 // Unique message identifier inside this chat. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageThreadId               *int                           `json:"message_thread_id,omitempty"`                // Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
	From                          *User                          `json:"from,omitempty"`                             // Optional. Sender of the message; may be empty for messages sent to channels. For backward compatibility, if the message was sent on behalf of a chat, the field contains a fake sender user in non-channel chats
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`                      // Optional. Sender of the message when sent on behalf of a chat. For example, the supergroup itself for messages sent by its anonymous administrators or a linked channel for messages automatically forwarded to the channel's discussion group. For backward compatibility, if the message was sent on behalf of a chat, the field from contains a fake sender user in non-channel chats.
	SenderBoostCount              *int                           `json:"sender_boost_count,omitempty"`               // Optional. If the sender of the message boosted the chat, the number of boosts added by the user
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`              // Optional. The bot that actually sent the message on behalf of the business account. Available only for outgoing messages sent on behalf of the connected business account.
	Date                          int                            `json:"date"`                                       // Date the message was sent in Unix time. It is always a positive number, representing a valid date.
	BusinessConnectionId          *string                        `json:"business_connection_id,omitempty"`           // Optional. Unique identifier of the business connection from which the message was received. If non-empty, the message belongs to a chat of the corresponding business account that is independent of any potential bot chat which might share the same identifier.
	Chat                          Chat                           `json:"chat"`                                       // Chat the message belongs to
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`                   // Optional. Information about the original message for forwarded messages
	IsTopicMessage                *bool                          `json:"is_topic_message,omitempty"`                 // Optional. True, if the message is sent to a forum topic
	IsAutomaticForward            *bool                          `json:"is_automatic_forward,omitempty"`             // Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`                 // Optional. For replies in the same chat and message thread, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`                   // Optional. Information about the message that is being replied to, which may come from another chat or forum topic
	Quote                         *TextQuote                     `json:"quote,omitempty"`                            // Optional. For replies that quote part of the original message, the quoted part of the message
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`                   // Optional. For replies to a story, the original story
	ViaBot                        *User                          `json:"via_bot,omitempty"`                          // Optional. Bot through which the message was sent
	EditDate                      *int                           `json:"edit_date,omitempty"`                        // Optional. Date the message was last edited in Unix time
	HasProtectedContent           *bool                          `json:"has_protected_content,omitempty"`            // Optional. True, if the message can't be forwarded
	IsFromOffline                 *bool                          `json:"is_from_offline,omitempty"`                  // Optional. True, if the message was sent by an implicit action, for example, as an away or a greeting business message, or as a scheduled message
	MediaGroupId                  *string                        `json:"media_group_id,omitempty"`                   // Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature               *string                        `json:"author_signature,omitempty"`                 // Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	Text                          *string                        `json:"text,omitempty"`                             // Optional. For text messages, the actual UTF-8 text of the message
	Entities                      *[]MessageEntity               `json:"entities,omitempty"`                         // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`             // Optional. Options used for link preview generation for the message, if it is a text message and link preview options were changed
	EffectId                      *string                        `json:"effect_id,omitempty"`                        // Optional. Unique identifier of the message effect added to the message
	Animation                     *Animation                     `json:"animation,omitempty"`                        // Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Audio                         *Audio                         `json:"audio,omitempty"`                            // Optional. Message is an audio file, information about the file
	Document                      *Document                      `json:"document,omitempty"`                         // Optional. Message is a general file, information about the file
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`                       // Optional. Message contains paid media; information about the paid media
	Photo                         *[]PhotoSize                   `json:"photo,omitempty"`                            // Optional. Message is a photo, available sizes of the photo
	Sticker                       *Sticker                       `json:"sticker,omitempty"`                          // Optional. Message is a sticker, information about the sticker
	Story                         *Story                         `json:"story,omitempty"`                            // Optional. Message is a forwarded story
	Video                         *Video                         `json:"video,omitempty"`                            // Optional. Message is a video, information about the video
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`                       // Optional. Message is a video note, information about the video message
	Voice                         *Voice                         `json:"voice,omitempty"`                            // Optional. Message is a voice message, information about the file
	Caption                       *string                        `json:"caption,omitempty"`                          // Optional. Caption for the animation, audio, document, paid media, photo, video or voice
	CaptionEntities               *[]MessageEntity               `json:"caption_entities,omitempty"`                 // Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	ShowCaptionAboveMedia         *bool                          `json:"show_caption_above_media,omitempty"`         // Optional. True, if the caption must be shown above the message media
	HasMediaSpoiler               *bool                          `json:"has_media_spoiler,omitempty"`                // Optional. True, if the message media is covered by a spoiler animation
	Contact                       *Contact                       `json:"contact,omitempty"`                          // Optional. Message is a shared contact, information about the contact
	Dice                          *Dice                          `json:"dice,omitempty"`                             // Optional. Message is a die with random value
	Game                          *Game                          `json:"game,omitempty"`                             // Optional. Message is a game, information about the game. More about games »
	Poll                          *Poll                          `json:"poll,omitempty"`                             // Optional. Message is a native poll, information about the poll
	Venue                         *Venue                         `json:"venue,omitempty"`                            // Optional. Message is a venue, information about the venue. For backward compatibility, when this field is set, the location field will also be set
	Location                      *Location                      `json:"location,omitempty"`                         // Optional. Message is a shared location, information about the location
	NewChatMembers                *[]User                        `json:"new_chat_members,omitempty"`                 // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`                 // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle                  *string                        `json:"new_chat_title,omitempty"`                   // Optional. A chat title was changed to this value
	NewChatPhoto                  *[]PhotoSize                   `json:"new_chat_photo,omitempty"`                   // Optional. A chat photo was change to this value
	DeleteChatPhoto               *bool                          `json:"delete_chat_photo,omitempty"`                // Optional. Service message: the chat photo was deleted
	GroupChatCreated              *bool                          `json:"group_chat_created,omitempty"`               // Optional. Service message: the group has been created
	SupergroupChatCreated         *bool                          `json:"supergroup_chat_created,omitempty"`          // Optional. Service message: the supergroup has been created. This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated            *bool                          `json:"channel_chat_created,omitempty"`             // Optional. Service message: the channel has been created. This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_change,omitempty"` // Optional. Service message: auto-delete timer settings changed in the chat
	MigrateToChatId               *int                           `json:"migrate_to_chat_id,omitempty"`               // Optional. The group has been migrated to a supergroup with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit int or double-precision float type are safe for storing this identifier.
	MigrateFromChatId             *int                           `json:"migrate_from_chat_id,omitempty"`             // Optional. The supergroup has been migrated from a group with the specified identifier. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit int or double-precision float type are safe for storing this identifier.
	PinnedMessage                 *Message                       `json:"pinned_message,omitempty"`                   // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	Invoice                       *Invoice                       `json:"invoice,omitempty"`                          // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`               // Optional. Message is a service message about a successful payment, information about the payment. More about payments »
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`                 // Optional. Message is a service message about a refunded payment, information about the payment. More about payments »
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`                     // Optional. Service message: users were shared with the bot
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`                      // Optional. Service message: a chat was shared with the bot
	ConnectedWebsite              *string                        `json:"connected_website,omitempty"`                // Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`             // Optional. Service message: the user allowed the bot to write messages after adding it to the attachment or side menu, launching a Web App from a link, or accepting an explicit request from a Web App sent by the method requestWriteAccess
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`                    // Optional. Telegram Passport data
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`        // Optional. Service message. A user in the chat triggered another user's proximity alert while sharing Live Location.
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`                      // Optional. Service message: user boosted the chat
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`              // Optional. Service message: chat background set
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`              // Optional. Service message: forum topic created
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`               // Optional. Service message: forum topic edited
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`               // Optional. Service message: forum topic closed
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`             // Optional. Service message: forum topic reopened
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`       // Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`     // Optional. Service message: the 'General' forum topic unhidden
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`                 // Optional. Service message: a scheduled giveaway was created
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`                         // Optional. The message is a scheduled giveaway message
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`                 // Optional. A giveaway with public winners was completed
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`               // Optional. Service message: a giveaway without public winners was completed
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`             // Optional. Service message: video chat scheduled
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`               // Optional. Service message: video chat started
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`                 // Optional. Service message: video chat ended
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`  // Optional. Service message: new participants invited to a video chat
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`                     // Optional. Service message: data sent by a Web App
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`                     // Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
}

type BusinessConnection struct {
	Id         string `json:"id"`           // Unique identifier of the business connection
	User       User   `json:"user"`         // Business account user that created the business connection
	UserChatId int    `json:"user_chat_id"` // Identifier of a private chat with the user who created the business connection. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit int or double-precision float type are safe for storing this identifier.
	Date       int    `json:"date"`         // Date the connection was established in Unix time
	CanReply   bool   `json:"can_reply"`    // True, if the bot can act on behalf of the business account in chats that were active in the last 24 hours
	IsEnabled  bool   `json:"is_enabled"`   // True, if the connection is active
}

type BusinessMessagesDeleted struct {
	BusinessConnectionId string `json:"business_connection_id"` // Unique identifier of the business connection
	Chat                 Chat   `json:"chat"`                   // Information about a chat in the business account. The bot may not have access to the chat or the corresponding user.
	MessageIds           []int  `json:"message_ids"`            // The list of identifiers of deleted messages in the chat of the business account
}

type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`                 // The chat containing the message the user reacted to
	MessageId   int            `json:"message_id"`           // Unique identifier of the message inside the chat
	User        *User          `json:"user,omitempty"`       // Optional. The user that changed the reaction, if the user isn't anonymous
	ActorChat   *Chat          `json:"actor_chat,omitempty"` // Optional. The chat on behalf of which the reaction was changed, if the user is anonymous
	Date        int            `json:"date"`                 // Date of the change in Unix time
	OldReaction []ReactionType `json:"old_reaction"`         // Previous list of reaction types that were set by the user
	NewReaction []ReactionType `json:"new_reaction"`         // New list of reaction types that have been set by the user
}

type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`       // The chat containing the message
	MessageId int             `json:"message_id"` // Unique message identifier inside the chat
	Date      int             `json:"date"`       // Date of the change in Unix time
	Reactions []ReactionCount `json:"reactions"`  // Count List of reactions that are present on the message
}

type InlineQuery struct {
	Id       string    `json:"id"`                  // Unique identifier for this query
	From     User      `json:"from"`                // Sender
	Query    string    `json:"query"`               // Text of the query (up to 256 characters)
	Offset   string    `json:"offset"`              // Offset of the results to be returned, can be controlled by the bot
	ChatType string    `json:"chat_type,omitempty"` // Optional. Type of the chat from which the inline query was sent. Can be either “sender” for a private chat with the inline query sender, “private”, “group”, “supergroup”, or “channel”. The chat type should be always known for requests sent from official clients and most third-party clients, unless the request was sent from a secret chat
	Location *Location `json:"location,omitempty"`  // Optional. Sender location, only for bots that request user location
}

type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`                   // The unique identifier for the result that was chosen
	From            User      `json:"from"`                        // The user that chose the result
	Location        *Location `json:"location,omitempty"`          // Optional. Sender location, only for bots that require user location
	InlineMessageId *string   `json:"inline_message_id,omitempty"` // Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	Query           string    `json:"query"`                       // The query that was used to obtain the result
}

type CallbackQuery struct {
	Id              string   `json:"id"`                          // Unique identifier for this query
	From            User     `json:"from"`                        // Sender
	Message         *Message `json:"message,omitempty"`           // Optional. Message sent by the bot with the callback button that originated the query
	InlineMessageId *string  `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance    string   `json:"chat_instance"`               // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data            *string  `json:"data,omitempty"`              // Optional. Data associated with the callback button. Be aware that the message originated the query can contain no callback buttons with this data.
	GameShortName   *string  `json:"game_short_name,omitempty"`   // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

type ShippingQuery struct {
	Id              string          `json:"id"`               // Unique query identifier
	From            User            `json:"from"`             // User who sent the query
	InvoicePayload  string          `json:"invoice_payload"`  // Bot-specified invoice payload
	ShippingAddress ShippingAddress `json:"shipping_address"` // User specified shipping address
}

type PreCheckoutQuery struct {
	Id               string     `json:"id"`                           // Unique query identifier
	From             User       `json:"from"`                         // User who sent the query
	Currency         string     `json:"currency"`                     // Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	TotalAmount      int        `json:"total_amount"`                 // Total price in the smallest units of the currency (int, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload   string     `json:"invoice_payload"`              // Bot-specified invoice payload
	ShippingOptionId *string    `json:"shipping_option_id,omitempty"` // Optional. Identifier of the shipping option chosen by the user
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`         // Optional. Order information provided by the user
}

type PaidMediaPurchased struct {
	From             User   `json:"from"`               // User who purchased the media
	PaidMediaPayload string `json:"paid_media_payload"` // Bot-specified paid media payload
}

type Poll struct {
	Id                    string           `json:"id"`                             // Unique poll identifier
	Question              string           `json:"question"`                       // Poll question, 1-300 characters
	QuestionEntities      *[]MessageEntity `json:"question_entities"`              // Optional. Special entities that appear in the question. Currently, only custom emoji entities are allowed in poll questions
	Options               []PollOption     `json:"options"`                        // List of poll options
	TotalVoterCount       int              `json:"total_voter_count"`              // Total number of users that voted in the poll
	IsClosed              bool             `json:"is_closed"`                      // True, if the poll is closed
	IsAnonymous           bool             `json:"is_anonymous"`                   // True, if the poll is anonymous
	Type                  string           `json:"type"`                           // Poll type, currently can be “regular” or “quiz”
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`        // True, if the poll allows multiple answers
	CorrectOptionId       *int             `json:"correct_option_id,omitempty"`    // Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	Explanation           *string          `json:"explanation,omitempty"`          // Optional. Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters
	ExplanationEntities   *[]MessageEntity `json:"explanation_entities,omitempty"` // Optional. Special entities like usernames, URLs, bot commands, etc. that appear in the explanation
	OpenPeriod            *int             `json:"open_period,omitempty"`          // Optional. Amount of time in seconds the poll will be active after creation
	CloseDate             *int             `json:"close_date,omitempty"`           // Optional. Point in time (Unix timestamp) when the poll will be automatically closed
}

type PollAnswer struct {
	PollId    string `json:"poll_id"`    // Unique poll identifier
	VoterChat *Chat  `json:"voter_chat"` // Optional. The chat that changed the answer to the poll, if the voter is anonymous
	User      *User  `json:"user"`       // Optional. The user that changed the answer to the poll, if the voter isn't anonymous
	OptionIds []int  `json:"option_ids"` // 0-based identifiers of chosen answer options. May be empty if the vote was retracted.
}

type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`                                  // Chat the user belongs to
	From                    User            `json:"from"`                                  // Performer of the action, which resulted in the change
	Date                    int             `json:"date"`                                  // Date the change was done in Unix time
	OldChatMember           ChatMember      `json:"old_chat_member"`                       // Previous information about the chat member
	NewChatMember           ChatMember      `json:"new_chat_member"`                       // New information about the chat member
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`                 // Optional. Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
	ViaJoinRequest          *bool           `json:"via_join_request,omitempty"`            // Optional. True, if the user joined the chat after sending a direct join request without using an invitation link and being approved by an administrator
	ViaChatFolderInviteLink *bool           `json:"via_chat_folder_invite_link,omitempty"` // Optional. True, if the user joined the chat via a chat folder invite link
}

type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`                  // Chat to which the request was sent
	From       User            `json:"from"`                  // User that sent the join request
	UserChatId int             `json:"user_chat_id"`          // Identifier of a private chat with the user who sent the join request. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit int or double-precision float type are safe for storing this identifier. The bot can use this identifier for 5 minutes to send messages until the join request is processed, assuming no other administrator contacted the user.
	Date       int             `json:"date"`                  // Date the request was sent in Unix time
	Bio        *string         `json:"bio,omitempty"`         // Optional. Bio of the user.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"` // Optional. Chat invite link that was used by the user to send the join request
}

type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`  // Chat which was boosted
	Boost ChatBoost `json:"boost"` // Information about the chat boost
}

type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`        // Chat which was boosted
	BoostId    string          `json:"boost_id"`    // Unique identifier of the boost
	RemoveDate int             `json:"remove_date"` // Point in time (Unix timestamp) when the boost was removed
	Source     ChatBoostSource `json:"source"`      // Source of the removed boost
}

type Chat struct {
	Id        int     `json:"id"`                   // Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit int or double-precision float type are safe for storing this identifier.
	Type      string  `json:"type"`                 // Type of the chat, can be either “private”, “group”, “supergroup” or “channel”
	Title     *string `json:"title,omitempty"`      // Optional. Title, for supergroups, channels and group chats
	Username  *string `json:"username,omitempty"`   // Optional. Username, for private chats, supergroups and channels if available
	FirstName *string `json:"first_name,omitempty"` // Optional. First name of the other party in a private chat
	LastName  *string `json:"last_name,omitempty"`  // Optional. Last name of the other party in a private chat
	IsForum   *bool   `json:"is_forum,omitempty"`   // Optional. True, if the supergroup chat is a forum (has topics enabled)
}

const (
	ReactionTypeEmoji       = "emoji"
	ReactionTypeCustomEmoji = "custom_emoji"
	ReactionTypePaid        = "paid"
)

type ReactionType struct {
	Type          string  `json:"type"`            // Type of the reaction, always “emoji”
	Emoji         *string `json:"emoji"`           // Reaction emoji
	CustomEmojiId *string `json:"custom_emoji_id"` // Custom emoji identifier
}

type ReactionCount struct {
	Type       ReactionType `json:"type"`        // Type of the reaction
	TotalCount int          `json:"total_count"` // Number of times the reaction was added
}

type Location struct {
	Latitude             float32  `json:"latitude"`                         // Latitude as defined by the sender
	Longitude            float32  `json:"longitude"`                        // Longitude as defined by the sender
	HorizontalAccuracy   *float32 `json:"horizontal_accuracy,omitempty"`    // Optional. The radius of uncertainty for the location, measured in meters; 0-1500
	LivePeriod           *int     `json:"live_period,omitempty"`            // Optional. Time relative to the message sending date, during which the location can be updated; in seconds. For active live locations only.
	Heading              *int     `json:"heading,omitempty"`                // Optional. The direction in which user is moving, in degrees; 1-360. For active live locations only.
	ProximityAlertRadius *int     `json:"proximity_alert_radius,omitempty"` // Optional. The maximum distance for proximity alerts about approaching another chat member, in meters. For sent live locations only.
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"` // Two-letter ISO 3166-1 alpha-2 country code
	State       string `json:"state"`        // State, if applicable
	City        string `json:"city"`         // City
	StreetLine1 string `json:"street_line1"` // First line for the address
	StreetLine2 string `json:"street_line2"` // Second line for the address
	PostCode    string `json:"post_code"`    // Address post code
}

type OrderInfo struct {
	Name            *string          `json:"name,omitempty"`             // Optional. Username
	PhoneNumber     *string          `json:"phone_number,omitempty"`     // Optional. User's phone number
	Email           *string          `json:"email,omitempty"`            // Optional. User email
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` // Optional. User shipping address
}

type MessageEntity struct {
	Type          string  `json:"type"`                      // Type of the entity. Currently, can be “mention” (@username), “hashtag” (#hashtag or #hashtag@chatusername), “cashtag” ($USD or $USD@chatusername), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “spoiler” (spoiler message), “blockquote” (block quotation), “expandable_blockquote” (collapsed-by-default block quotation), “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable text URLs), “text_mention” (for users without usernames), “custom_emoji” (for inline custom emoji stickers)
	Offset        int     `json:"offset"`                    // Offset in UTF-16 code units to the start of the entity
	Length        int     `json:"length"`                    // Length of the entity in UTF-16 code units
	Url           *string `json:"url,omitempty"`             // Optional. For “text_link” only, URL that will be opened after user taps on the text
	User          *User   `json:"user,omitempty"`            // Optional. For “text_mention” only, the mentioned user
	Language      *string `json:"language,omitempty"`        // Optional. For “pre” only, the programming language of the entity text
	CustomEmojiId *string `json:"custom_emoji_id,omitempty"` // Optional. For “custom_emoji” only, unique identifier of the custom emoji. Use getCustomEmojiStickers to get full information about the sticker
}

type PollOption struct {
	Text         string           `json:"text"`                    // Option text, 1-100 characters
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"` // Optional. Special entities that appear in the option text. Currently, only custom emoji entities are allowed in poll option texts
	VoterCount   int              `json:"voter_count"`             // Number of users that voted for this option
}

type ChatMember struct { // TODO
	Status      string `json:"status"`       // The member's status in the chat, always “creator”
	User        User   `json:"user"`         // Information about the user
	IsAnonymous bool   `json:"is_anonymous"` // True, if the user's presence in the chat is hidden
}

type ChatInviteLink struct {
	InviteLink              string  `json:"invite_link"`                          // The invite link. If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	Creator                 User    `json:"creator"`                              // Creator of the link
	CreatesJoinRequest      bool    `json:"creates_join_request"`                 // True, if users joining the chat via the link need to be approved by chat administrators
	IsPrimary               bool    `json:"is_primary"`                           // True, if the link is primary
	IsRevoked               bool    `json:"is_revoked"`                           // True, if the link is revoked
	Name                    *string `json:"name,omitempty"`                       // Optional. Invite link name
	ExpireDate              *int    `json:"expire_date,omitempty"`                // Optional. Point in time (Unix timestamp) when the link will expire or has been expired
	MemberLimit             *int    `json:"member_limit,omitempty"`               // Optional. The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	PendingJoinRequestCount *int    `json:"pending_join_request_count,omitempty"` // Optional. Number of pending join requests created using this link
	SubscriptionPeriod      *int    `json:"subscription_period,omitempty"`        // Optional. The number of seconds the subscription will be active for before the next payment
	SubscriptionPrice       *int    `json:"subscription_price,omitempty"`         // Optional. The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat using the link
}

type ChatBoost struct {
	BoostId        string          `json:"boost_id"`        // Unique identifier of the boost
	AddDate        int             `json:"add_date"`        // Point in time (Unix timestamp) when the chat was boosted
	ExpirationDate int             `json:"expiration_date"` // Point in time (Unix timestamp) when the boost will automatically expire, unless the booster's Telegram Premium subscription is prolonged
	Source         ChatBoostSource `json:"source"`          // Source of the added boost
}

type ChatBoostSource struct {
	Source string `json:"source"` // Source of the boost, always “premium”
	User   User   `json:"user"`   // User that boosted the chat
}

const (
	MessageOriginUser       = "user"
	MessageOriginHiddenUser = "hidden_user"
	MessageOriginChat       = "chat"
	MessageOriginChannel    = "channel"
)

type MessageOrigin struct {
	Type            string  `json:"type"`                       // Type of the message origin
	Date            int     `json:"date"`                       // Date the message was sent originally in Unix time
	SenderUser      *User   `json:"sender_user"`                // User that sent the message originally
	SenderUserName  *string `json:"sender_user_name"`           // Name of the user that sent the message originally
	SenderChat      *Chat   `json:"sender_chat"`                // Chat that sent the message originally
	Chat            *Chat   `json:"chat"`                       // Channel chat to which the message was originally sent
	MessageId       *int    `json:"message_id"`                 // Unique message identifier inside the chat
	AuthorSignature *string `json:"author_signature,omitempty"` // Optional. Signature of the original post author
}

// ExternalReplyInfo This object contains information about a message that is being replied to, which may come from another chat or forum topic.
type ExternalReplyInfo struct {
	Origin             MessageOrigin       `json:"origin"`                         // Origin of the message replied to by the given message
	Chat               *Chat               `json:"chat,omitempty"`                 // Optional. Chat the original message belongs to. Available only if the chat is a supergroup or a channel.
	MessageId          *int                `json:"message_id,omitempty"`           // Optional. Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel.
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"` // Optional. Options used for link preview generation for the original message, if it is a text message
	Animation          *Animation          `json:"animation,omitempty"`            // Optional. Message is an animation, information about the animation
	Audio              *Audio              `json:"audio,omitempty"`                // Optional. Message is an audio file, information about the file
	Document           *Document           `json:"document,omitempty"`             // Optional. Message is a general file, information about the file
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`           // Optional. Message contains paid media; information about the paid media
	Photo              *[]PhotoSize        `json:"photo,omitempty"`                // Optional. Message is a photo, available sizes of the photo
	Sticker            *Sticker            `json:"sticker,omitempty"`              // Optional. Message is a sticker, information about the sticker
	Story              *Story              `json:"story,omitempty"`                // Optional. Message is a forwarded story
	Video              *Video              `json:"video,omitempty"`                // Optional. Message is a video, information about the video
	VideoNote          *VideoNote          `json:"video_note,omitempty"`           // Optional. Message is a video note, information about the video message
	Voice              *Voice              `json:"voice,omitempty"`                // Optional. Message is a voice message, information about the file
	HasMediaSpoiler    *bool               `json:"has_media_spoiler,omitempty"`    // Optional. True, if the message media is covered by a spoiler animation
	Contact            *Contact            `json:"contact,omitempty"`              // Optional. Message is a shared contact, information about the contact
	Dice               *Dice               `json:"dice,omitempty"`                 // Optional. Message is a die with random value
	Game               *Game               `json:"game,omitempty"`                 // Optional. Message is a game, information about the game. More about games »
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`             // Optional. Message is a scheduled giveaway, information about the giveaway
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`     // Optional. A giveaway with public winners was completed
	Invoice            *Invoice            `json:"invoice,omitempty"`              // Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	Location           *Location           `json:"location,omitempty"`             // Optional. Message is a shared location, information about the location
	Poll               *Poll               `json:"poll,omitempty"`                 // Optional. Message is a native poll, information about the poll
	Venue              *Venue              `json:"venue,omitempty"`                // Optional. Message is a venue, information about the venue
}

type TextQuote struct {
	Text     string           `json:"text"`                // Text of the quoted part of a message that is replied to by the given message
	Entities *[]MessageEntity `json:"entities,omitempty"`  // Optional. Special entities that appear in the quote. Currently, only bold, italic, underline, strikethrough, spoiler, and custom_emoji entities are kept in quotes.
	Position *int             `json:"position,omitempty"`  // Approximate quote position in the original message in UTF-16 code units as specified by the sender
	IsManual *bool            `json:"is_manual,omitempty"` // Optional. True, if the quote was chosen manually by the message sender. Otherwise, the quote was added automatically by the server.
}

type Story struct {
	Chat Chat `json:"chat"` // Chat that posted the story
	Id   int  `json:"id"`   // Unique identifier for the story in the chat
}

type LinkPreviewOptions struct {
	IsDisabled       *bool   `json:"is_disabled,omitempty"`        // Optional. True, if the link preview is disabled
	Url              *string `json:"url,omitempty"`                // Optional. URL to use for the link preview. If empty, then the first URL found in the message text will be used
	PreferSmallMedia *bool   `json:"prefer_small_media,omitempty"` // Optional. True, if the media in the link preview is supposed to be shrunk; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	PreferLargeMedia *bool   `json:"prefer_large_media,omitempty"` // Optional. True, if the media in the link preview is supposed to be enlarged; ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	ShowAboveText    *bool   `json:"show_above_text,omitempty"`    // Optional. True, if the link preview must be shown above the message text; otherwise, the link preview will be shown below the message text
}
type Animation struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int        `json:"width"`               // Video width as defined by the sender
	Height       int        `json:"height"`              // Video height as defined by the sender
	Duration     int        `json:"duration"`            // Duration of the video in seconds as defined by the sender
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Animation thumbnail as defined by the sender
	FileName     *string    `json:"file_name,omitempty"` // Optional. Original animation filename as defined by the sender
	MimeType     *string    `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by the sender
	FileSize     *int       `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit int or double-precision float type are safe for storing this value.
}
type Audio struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int        `json:"duration"`            // Duration of the audio in seconds as defined by the sender
	Performer    *string    `json:"performer,omitempty"` // Optional. Performer of the audio as defined by the sender or by audio tags
	Title        *string    `json:"title,omitempty"`     // Optional. Title of the audio as defined by the sender or by audio tags
	FileName     *string    `json:"file_name,omitempty"` // Optional. Original filename as defined by the sender
	MimeType     *string    `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by the sender
	FileSize     *int       `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit int or double-precision float type are safe for storing this value.
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Thumbnail of the album cover to which the music file belongs
}
type Document struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Document thumbnail as defined by the sender
	FileName     *string    `json:"file_name,omitempty"` // Optional. Original filename as defined by the sender
	MimeType     *string    `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by the sender
	FileSize     *int       `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit int or double-precision float type are safe for storing this value.
}

type PaidMediaInfo struct {
	StarCount int         `json:"star_count"` // The number of Telegram Stars that must be paid to buy access to the media
	PaidMedia []PaidMedia `json:"paid_media"` // Information about the paid media
}

const (
	PaidMediaPreview = "preview"
	PaidMediaPhoto   = "photo"
	PaidMediaVideo   = "video"
)

type PaidMedia struct {
	Type     string       `json:"type"`               // Type of the paid media
	Width    *int         `json:"width,omitempty"`    // Optional. Media width as defined by the sender
	Height   *int         `json:"height,omitempty"`   // Optional. Media height as defined by the sender
	Duration *int         `json:"duration,omitempty"` // Optional. Duration of the media in seconds as defined by the sender
	Photo    *[]PhotoSize `json:"photo"`              // The photo
	Video    *Video       `json:"video"`              // The video
}

type PhotoSize struct {
	FileId       string `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int    `json:"width"`               // Photo width
	Height       int    `json:"height"`              // Photo height
	FileSize     *int   `json:"file_size,omitempty"` // Optional. File size in bytes
}
type Sticker struct {
	FileId           string        `json:"file_id"`                     // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId     string        `json:"file_unique_id"`              // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Type             string        `json:"type"`                        // Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”. The type of the sticker is independent of its format, which is determined by the fields is_animated and is_video.
	Width            int           `json:"width"`                       // Sticker width
	Height           int           `json:"height"`                      // Sticker height
	IsAnimated       *bool         `json:"is_animated,omitempty"`       // True, if the sticker is animated
	IsVideo          *bool         `json:"is_video,omitempty"`          // True, if the sticker is a video sticker
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`         // Optional. Sticker thumbnail in the .WEBP or .JPG format
	Emoji            *string       `json:"emoji,omitempty"`             // Optional. Emoji associated with the sticker
	SetName          *string       `json:"set_name,omitempty"`          // Optional. Name of the sticker set to which the sticker belongs
	PremiumAnimation *File         `json:"premium_animation,omitempty"` // Optional. For premium regular stickers, premium animation for the sticker
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`     // Optional. For mask stickers, the position where the mask should be placed
	CustomEmojiId    *string       `json:"custom_emoji_id ,omitempty"`  // Optional. For custom emoji stickers, unique identifier of the custom emoji
	NeedsRepainting  *bool         `json:"needs_repainting,omitempty"`  // Optional. True, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places
	FileSize         *int          `json:"file_size,omitempty"`         // Optional. File size in bytes
}
type File struct {
	FileId       string  `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string  `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     *int    `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit int or double-precision float type are safe for storing this value.
	FilePath     *string `json:"file_path,omitempty"` // Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}
type MaskPosition struct {
	Point  string  `json:"point"`   // The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
	XShift float64 `json:"x_shift"` // Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	YShift float64 `json:"y_shift"` // Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	Scale  float64 `json:"scale"`   // Mask scaling coefficient. For example, 2.0 means double size.
}
type Video struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int        `json:"width"`               // Video width as defined by the sender
	Height       int        `json:"height"`              // Video height as defined by the sender
	Duration     int        `json:"duration"`            // Duration of the video in seconds as defined by the sender
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Video thumbnail
	FileName     *string    `json:"file_name,omitempty"` // Optional. Original filename as defined by the sender
	MimeType     *string    `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by the sender
	FileSize     *int       `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit int or double-precision float type are safe for storing this value.
}
type VideoNote struct {
	FileId       string     `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Length       int        `json:"length"`              // Video width and height (diameter of the video message) as defined by the sender
	Duration     int        `json:"duration"`            // Duration of the video in seconds as defined by the sender
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"` // Optional. Video thumbnail
	FileSize     *int       `json:"file_size,omitempty"` // Optional. File size in bytes
}
type Voice struct {
	FileId       string  `json:"file_id"`             // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string  `json:"file_unique_id"`      // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int     `json:"duration"`            // Duration of the audio in seconds as defined by the sender
	MimeType     *string `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by the sender
	FileSize     *int    `json:"file_size,omitempty"` // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit int or double-precision float type are safe for storing this value.
}
type Contact struct {
	PhoneNumber string  `json:"phone_number"`        // Contact's phone number
	FirstName   string  `json:"first_name"`          // Contact's first name
	LastName    *string `json:"last_name,omitempty"` // Optional. Contact's last name
	UserId      *int    `json:"user_id,omitempty"`   // Optional. Contact's user identifier in Telegram. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit int or double-precision float type are safe for storing this identifier.
	Vcard       *string `json:"vcard,omitempty"`     // Optional. Additional data about the contact in the form of a vCard
}
type Dice struct {
	Emoji string `json:"emoji"` // Emoji on which the dice throw animation is based
	Value int    `json:"value"` // Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀” and “⚽” base emoji, 1-64 for “🎰” base emoji
}
type Game struct {
	Title        string           `json:"title"`                   // Title of the game
	Description  string           `json:"description"`             // Description of the game
	Photo        []PhotoSize      `json:"photo"`                   // Photo that will be displayed in the game message in chats.
	Text         *string          `json:"text,omitempty"`          // Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	TextEntities *[]MessageEntity `json:"text_entities,omitempty"` // Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	Animation    *Animation       `json:"animation,omitempty"`     // Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
}
type Giveaway struct {
	Chats                         []Chat    `json:"chats"`                                      // The list of chats which the user must join to participate in the giveaway
	WinnersSelectionDate          int       `json:"winners_selection_date"`                     // Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnerCount                   int       `json:"winner_count"`                               // The number of users which are supposed to be selected as winners of the giveaway
	OnlyNewMembers                *bool     `json:"only_new_members"`                           // Optional. True, if only users who join the chats after the giveaway started should be eligible to win
	HasPublicWinners              *bool     `json:"has_public_winners,omitempty"`               // Optional. True, if the list of giveaway winners will be visible to everyone
	PrizeDescription              *string   `json:"prize_description,omitempty"`                // Optional. Description of additional giveaway prize
	CountryCodes                  *[]string `json:"country_codes,omitempty"`                    // Optional. A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from which eligible users for the giveaway must come. If empty, then all users can participate in the giveaway. Users with a phone number that was bought on Fragment can always participate in giveaways.
	PrizeStarCount                *int      `json:"prize_star_count,omitempty"`                 // Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PremiumSubscriptionMonthCount *int      `json:"premium_subscription_month_count,omitempty"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
}
type GiveawayWinners struct {
	Chat                          Chat    `json:"chat"`                                       // The chat that created the giveaway
	GiveawayMessageId             int     `json:"giveaway_message_id"`                        // Identifier of the message with the giveaway in the chat
	WinnersSelectionDate          int     `json:"winners_selection_date"`                     // Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnerCount                   int     `json:"winner_count"`                               // Total number of winners in the giveaway
	Winners                       []User  `json:"winners"`                                    // List of up to 100 winners of the giveaway
	AdditionalChatCount           *int    `json:"additional_chat_count,omitempty"`            // Optional. The number of other chats the user had to join in order to be eligible for the giveaway
	PrizeStarCount                *int    `json:"prize_star_count,omitempty"`                 // Optional. The number of Telegram Stars that were split between giveaway winners; for Telegram Star giveaways only
	PremiumSubscriptionMonthCount *int    `json:"premium_subscription_month_count,omitempty"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for; for Telegram Premium giveaways only
	UnclaimedPrizeCount           *int    `json:"unclaimed_prize_count,omitempty"`            // Optional. Number of undistributed prizes
	OnlyNewMembers                *bool   `json:"only_new_members,omitempty"`                 // Optional. True, if only users who had joined the chats after the giveaway started were eligible to win
	WasRefunded                   *bool   `json:"was_refunded,omitempty"`                     // Optional. True, if the giveaway was canceled because the payment for it was refunded
	PrizeDescription              *string `json:"prize_description,omitempty"`                // Optional. Description of additional giveaway prize
}
type Invoice struct {
	Title          string `json:"title"`           // Product name
	Description    string `json:"description"`     // Product description
	StartParameter string `json:"start_parameter"` // Unique bot deep-linking parameter that can be used to generate this invoice
	Currency       string `json:"currency"`        // Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	TotalAmount    int    `json:"total_amount"`    // Total price in the smallest units of the currency (int, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}
type Venue struct {
	Location        Location `json:"location"`                    // Venue location. Can't be a live location
	Title           string   `json:"title"`                       // Name of the venue
	Address         string   `json:"address"`                     // Address of the venue
	FoursquareId    *string  `json:"foursquare_id,omitempty"`     // Optional. Foursquare identifier of the venue
	FoursquareType  *string  `json:"foursquare_type,omitempty"`   // Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	GooglePlaceId   *string  `json:"google_place_id,omitempty"`   // Optional. Google Places identifier of the venue
	GooglePlaceType *string  `json:"google_place_type,omitempty"` // Optional. Google Places type of the venue. (See supported types.)
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"` // New auto-delete time for messages in the chat; in seconds
}

type SuccessfulPayment struct {
	Currency                   string     `json:"currency"`                               // Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	TotalAmount                int        `json:"total_amount"`                           // Total price in the smallest units of the currency (int, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload             string     `json:"invoice_payload"`                        // Bot-specified invoice payload
	SubscriptionExpirationDate *int       `json:"subscription_expiration_date,omitempty"` // Optional. Expiration date of the subscription, in Unix time; for recurring payments only
	IsRecurring                *bool      `json:"is_recurring,omitempty"`                 // Optional. True, if the payment is a recurring payment for a subscription
	IsFirstRecurring           *bool      `json:"is_first_recurring,omitempty"`           // Optional. True, if the payment is the first payment for a subscription
	ShippingOptionId           *string    `json:"shipping_option_id,omitempty"`           // Optional. Identifier of the shipping option chosen by the user
	OrderInfo                  *OrderInfo `json:"order_info,omitempty"`                   // Optional. Order information provided by the user
	TelegramPaymentChargeId    string     `json:"telegram_payment_charge_id"`             // Telegram payment identifier
	ProviderPaymentChargeId    string     `json:"provider_payment_charge_id"`             // Provider payment identifier
}
type RefundedPayment struct {
	Currency                string  `json:"currency"`                             // Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars. Currently, always “XTR”
	TotalAmount             int     `json:"total_amount"`                         // Total refunded price in the smallest units of the currency (int, not float/double). For example, for a price of US$ 1.45, total_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload          string  `json:"invoice_payload"`                      // Bot-specified invoice payload
	TelegramPaymentChargeId string  `json:"telegram_payment_charge_id"`           // Telegram payment identifier
	ProviderPaymentChargeId *string `json:"provider_payment_charge_id,omitempty"` // Optional. Provider payment identifier
}

type UsersShared struct {
	RequestId int          `json:"request_id"` // Identifier of the request
	Users     []SharedUser `json:"users"`      // Information about users shared with the bot.
}

type ChatShared struct {
	RequestId int          `json:"request_id"`         // Identifier of the request
	ChatId    int          `json:"chat_id "`           // Identifier of the shared chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit int or double-precision float type are safe for storing this identifier. The bot may not have access to the chat and could be unable to use this identifier, unless the chat is already known to the bot by some other means.
	Title     *string      `json:"title,omitempty"`    // Optional. Title of the chat, if the title was requested by the bot.
	Username  *string      `json:"username,omitempty"` // Optional. Username of the chat, if the username was requested by the bot and available.
	Photo     *[]PhotoSize `json:"photo,omitempty"`    // Optional. Available sizes of the chat photo, if the photo was requested by the bot
}

type WriteAccessAllowed struct {
	FromRequest        *bool   `json:"from_request,omitempty"`         // Optional. True, if the access was granted after the user accepted an explicit request from a Web App sent by the method requestWriteAccess
	WebAppName         *string `json:"web_app_name,omitempty"`         // Optional. Name of the Web App, if the access was granted when the Web App was launched from a link
	FromAttachmentMenu *bool   `json:"from_attachment_menu,omitempty"` // Optional. True, if the access was granted when the bot was added to the attachment or side menu
}

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`        // Array with information about documents and other Telegram Passport elements that was shared with the bot
	Credentials EncryptedCredentials       `json:"credentials"` // Encrypted credentials required to decrypt the data
}

type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"` // User that triggered the alert
	Watcher  User `json:"watcher"`  // User that set the alert
	Distance int  `json:"distance"` // The distance between the users
}

type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"` // Number of boosts added by the user
}

type ChatBackground struct {
	Type BackgroundType `json:"type"` // Type of the background
}

type ForumTopicCreated struct {
	Name              string  `json:"name"`                           // Name of the topic
	IconColor         int     `json:"icon_color"`                     // Color of the topic icon in RGB format
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"` // Optional. Unique identifier of the custom emoji shown as the topic icon
}

type ForumTopicEdited struct {
	Name              *string `json:"name,omitempty"`                 // Optional. New name of the topic, if it was edited
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"` // Optional. New identifier of the custom emoji shown as the topic icon, if it was edited; an empty string if the icon was removed
}

type ForumTopicClosed struct {
}

type ForumTopicReopened struct {
}

type GeneralForumTopicHidden struct {
}

type GeneralForumTopicUnhidden struct {
}

type GiveawayCreated struct {
	PrizeStarCount *int `json:"prize_star_count"` // Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
}

type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`                    // Number of winners in the giveaway
	UnclaimedPrizeCount *int     `json:"unclaimed_prize_count,omitempty"` // Optional. Number of undistributed prizes
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`      // Optional. Message with the giveaway that was completed, if it wasn't deleted
	IsStarGiveaway      *bool    `json:"is_star_giveaway,omitempty"`      // Optional. True, if the giveaway is a Telegram Star giveaway. Otherwise, currently, the giveaway is a Telegram Premium giveaway.
}

type VideoChatScheduled struct {
	StartDate int `json:"start_date"` // Point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator
}

type VideoChatStarted struct {
}

type VideoChatEnded struct {
	Duration int `json:"duration"` // Video chat duration in seconds
}

type VideoChatParticipantsInvited struct {
	Users []User `json:"users"` // New members that were invited to the video chat
}

type WebAppData struct {
	Data       string `json:"data"`        // The data. Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"` // Text of the web_app keyboard button from which the Web App was opened. Be aware that a bad client can send arbitrary data in this field.
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`                          // Array of button rows, each represented by an Array of KeyboardButton objects
	IsPersistent          *bool              `json:"is_persistent,omitempty"`           // Optional. Requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	ResizeKeyboard        *bool              `json:"resize_keyboard,omitempty"`         // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard       *bool              `json:"one_time_keyboard,omitempty"`       // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	InputFieldPlaceholder *string            `json:"input_field_placeholder,omitempty"` // Optional. The placeholder to be shown in the input field when the keyboard is active; 1-64 characters
	Selective             *bool              `json:"selective,omitempty"`               // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply to a message in the same chat and forum topic, sender of the original message. Example: A user requests to change the bot's language, bot replies to the request with a keyboard to select the new language. Other users in the group don't see the keyboard.
}

type KeyboardButton struct {
	Text            string                      `json:"text"`                       // Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`    // Optional. If specified, pressing the button will open a list of suitable users. Identifiers of selected users will be sent to the bot in a “users_shared” service message. Available in private chats only.
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`     // Optional. If specified, pressing the button will open a list of suitable chats. Tapping on a chat will send its identifier to the bot in a “chat_shared” service message. Available in private chats only.
	RequestContact  *bool                       `json:"request_contact,omitempty"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only.
	RequestLocation *bool                       `json:"request_location,omitempty"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only.
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`     // Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only.
	WebApp          *WebAppInfo                 `json:"web_app,omitempty"`          // Optional. If specified, the described Web App will be launched when the button is pressed. The Web App will be able to send a “web_app_data” service message. Available in private chats only.
}

type KeyboardButtonRequestUsers struct {
	RequestId       int   `json:"request_id"`                 // Signed 32-bit identifier of the request that will be received back in the UsersShared object. Must be unique within the message
	UserIsBot       *bool `json:"user_is_bot,omitempty"`      // Optional. Pass True to request bots, pass False to request regular users. If not specified, no additional restrictions are applied.
	UserIsPremium   *bool `json:"user_is_premium,omitempty"`  // Optional. Pass True to request premium users, pass False to request non-premium users. If not specified, no additional restrictions are applied.
	MaxQuantity     *int  `json:"max_quantity,omitempty"`     // Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
	RequestName     *bool `json:"request_name,omitempty"`     // Optional. Pass True to request the users' first and last names
	RequestUsername *bool `json:"request_username,omitempty"` // Optional. Pass True to request the users' usernames
	RequestPhoto    *bool `json:"request_photo,omitempty"`    // Optional. Pass True to request the users' photos
}

type KeyboardButtonRequestChat struct {
	RequestId               int                      `json:"request_id"`                          // Signed 32-bit identifier of the request, which will be received back in the ChatShared object. Must be unique within the message
	ChatIsChannel           *bool                    `json:"chat_is_channel,omitempty"`           // Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsForum             *bool                    `json:"chat_is_forum,omitempty"`             // Optional. Pass True to request a forum supergroup, pass False to request a non-forum chat. If not specified, no additional restrictions are applied.
	ChatHasUsername         *bool                    `json:"chat_has_username,omitempty"`         // Optional. Pass True to request a supergroup or a channel with a username, pass False to request a chat without a username. If not specified, no additional restrictions are applied.
	ChatIsCreated           *bool                    `json:"chat_is_created,omitempty"`           // Optional. Pass True to request a chat owned by the user. Otherwise, no additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"` // Optional. A JSON-serialized object listing the required administrator rights of the user in the chat. The rights must be a superset of bot_administrator_rights. If not specified, no additional restrictions are applied.
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`  // Optional. A JSON-serialized object listing the required administrator rights of the bot in the chat. The rights must be a subset of user_administrator_rights. If not specified, no additional restrictions are applied.
	BotIsMember             *bool                    `json:"bot_is_member,omitempty"`             // Optional. Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
	RequestTitle            *bool                    `json:"request_title,omitempty"`             // Optional. Pass True to request the chat's title
	RequestUsername         *bool                    `json:"request_username,omitempty"`          // Optional. Pass True to request the chat's username
	RequestPhoto            *bool                    `json:"request_photo,omitempty"`             // Optional. Pass True to request the chat's photo
}

type ChatAdministratorRights struct {
	IsAnonymous         *bool `json:"is_anonymous,omitempty"`           // True, if the user's presence in the chat is hidden
	CanManageChat       *bool `json:"can_manage_chat,omitempty"`        // True, if the administrator can access the chat event log, get boost list, see hidden supergroup and channel members, report spam messages and ignore slow mode. Implied by any other administrator privilege.
	CanDeleteMessages   *bool `json:"can_delete_messages,omitempty"`    // True, if the administrator can delete messages of other users
	CanManageVideoChats *bool `json:"can_manage_video_chats,omitempty"` // True, if the administrator can manage video chats
	CanRestrictMembers  *bool `json:"can_restrict_members,omitempty"`   // True, if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanPromoteMembers   *bool `json:"can_promote_members,omitempty"`    // True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanChangeInfo       *bool `json:"can_change_info,omitempty"`        // True, if the user is allowed to change the chat title, photo and other settings
	CanInviteUsers      *bool `json:"can_invite_users,omitempty"`       // True, if the user is allowed to invite new users to the chat
	CanPostStories      *bool `json:"can_post_stories,omitempty"`       // True, if the administrator can post stories to the chat
	CanEditStories      *bool `json:"can_edit_stories,omitempty"`       // True, if the administrator can edit stories posted by other users, post stories to the chat page, pin chat stories, and access the chat's story archive
	CanDeleteStories    *bool `json:"can_delete_stories,omitempty"`     // True, if the administrator can delete stories posted by other users
	CanPostMessages     *bool `json:"can_post_messages,omitempty"`      // Optional. True, if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanEditMessages     *bool `json:"can_edit_messages,omitempty"`      // Optional. True, if the administrator can edit messages of other users and can pin messages; for channels only
	CanPinMessages      *bool `json:"can_pin_messages,omitempty"`       // Optional. True, if the user is allowed to pin messages; for groups and supergroups only
	CanManageTopics     *bool `json:"can_manage_topics,omitempty"`      // Optional. True, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
}

type KeyboardButtonPollType struct {
	Type *string `json:"type,omitempty"` // Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`                                       // Label text on the button
	Url                          *string                      `json:"url,omitempty"`                              // Optional. HTTP or tg:// URL to be opened when the button is pressed. Links tg://user?id=<user_id> can be used to mention a user by their identifier without using a username, if this is allowed by their privacy settings.
	CallbackData                 *string                      `json:"callback_data,omitempty"`                    // Optional. Data to be sent in a callback query to the bot when the button is pressed, 1-64 bytes
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`                          // Optional. Description of the Web App that will be launched when the user presses the button. The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery. Available only in private chats between a user and the bot. Not supported for messages sent on behalf of a Telegram Business account.
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`                        // Optional. An HTTPS URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`              // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot's username and the specified inline query in the input field. May be empty, in which case just the bot's username will be inserted. Not supported for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"` // Optional. If set, pressing the button will insert the bot's username and the specified inline query in the current chat's input field. May be empty, in which case only the bot's username will be inserted. This offers a quick way for the user to open your bot in inline mode in the same chat - good for selecting something from multiple options. Not supported in channels and for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`  // Optional. If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and insert the bot's username and the specified inline query in the input field. Not supported for messages sent on behalf of a Telegram Business account.
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`                        // Optional. Description of the button that copies the specified text to the clipboard.
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`                    // Optional. Description of the game that will be launched when the user presses the button. NOTE: This type of button must always be the first button in the first row.
	Pay                          *bool                        `json:"pay,omitempty"`                              // Optional. Specify True, to send a Pay button. Substrings “⭐” and “XTR” in the buttons text will be replaced with a Telegram Star icon. NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
}

type WebAppInfo struct {
	Url string `json:"url"` // An HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps
}

type LoginUrl struct {
	Url                string  `json:"url"`                            // An HTTPS URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data. NOTE: You must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization.
	ForwardText        *string `json:"forward_text,omitempty"`         // Optional. New text of the button in forwarded messages.
	BotUsername        *string `json:"bot_username,omitempty"`         // Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	RequestWriteAccess *bool   `json:"request_write_access,omitempty"` // Optional. Pass True to request the permission for your bot to send messages to the user.
}

type SwitchInlineQueryChosenChat struct {
	Query             *string `json:"query,omitempty"`               // Optional. The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	AllowUserChats    *bool   `json:"allow_user_chats,omitempty"`    // Optional. True, if private chats with users can be chosen
	AllowBotChats     *bool   `json:"allow_bot_chats,omitempty"`     // Optional. True, if private chats with bots can be chosen
	AllowGroupChats   *bool   `json:"allow_group_chats,omitempty"`   // Optional. True, if group and supergroup chats can be chosen
	AllowChannelChats *bool   `json:"allow_channel_chats,omitempty"` // Optional. True, if channel chats can be chosen
}

type CopyTextButton struct {
	Text string `json:"text"` // The text to be copied to the clipboard; 1-256 characters
}

type CallbackGame struct {
}

type SharedUser struct {
	UserId    int          `json:"user_id"`              // Identifier of the shared user. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so 64-bit ints or double-precision float types are safe for storing these identifiers. The bot may not have access to the user and could be unable to use this identifier, unless the user is already known to the bot by some other means.
	FirstName *string      `json:"first_name,omitempty"` // Optional. First name of the user, if the name was requested by the bot
	LastName  *string      `json:"last_name,omitempty"`  // Optional. Last name of the user, if the name was requested by the bot
	Username  *string      `json:"username,omitempty"`   // Optional. Username of the user, if the username was requested by the bot
	Photo     *[]PhotoSize `json:"photo,omitempty"`      // Optional. Available sizes of the chat photo, if the photo was requested by the bot
}

type EncryptedPassportElement struct {
	Type        string          `json:"type"`                   // Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”.
	Data        *string         `json:"data,omitempty"`         // Optional. Base64-encoded encrypted Telegram Passport element data provided by the user; available only for “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types. Can be decrypted and verified using the accompanying EncryptedCredentials.
	PhoneNumber *string         `json:"phone_number,omitempty"` // Optional. User's verified phone number; available only for “phone_number” type
	Email       *string         `json:"email,omitempty"`        // Optional. User's verified email address; available only for “email” type
	Files       *[]PassportFile `json:"files,omitempty"`        // Optional. Array of encrypted files with documents provided by the user; available only for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide   *PassportFile   `json:"front_side,omitempty"`   // Optional. Encrypted file with the front side of the document, provided by the user; available only for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile   `json:"reverse_side,omitempty"` // Optional. Encrypted file with the reverse side of the document, provided by the user; available only for “driver_license” and “identity_card”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie      *PassportFile   `json:"selfie,omitempty"`       // Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available if requested for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Translation *[]PassportFile `json:"translation,omitempty"`  // Optional. Array of encrypted files with translated versions of documents provided by the user; available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Hash        string          `json:"hash"`                   // Base64-encoded element hash for using in PassportElementErrorUnspecified
}

type PassportFile struct {
	FileId       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int    `json:"file_size"`      // File size in bytes
	FileDate     int    `json:"file_date"`      // Unix time when the file was uploaded
}

type EncryptedCredentials struct {
	Data   string `json:"data"`   // Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Hash   string `json:"hash"`   // Base64-encoded data hash for data authentication
	Secret string `json:"secret"` // Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
}

const (
	BackgroundTypeFill      = "fill"
	BackgroundTypeWallpaper = "wallpaper"
	BackgroundTypePattern   = "pattern"
	BackgroundTypeChatTheme = "chat_theme"
)

type BackgroundType struct {
	Type             string          `json:"type"`                         // Type of the background
	Fill             *BackgroundFill `json:"fill,omitempty"`               // The background fill
	DarkThemeDimming *int            `json:"dark_theme_dimming,omitempty"` // Dimming of the background in dark themes, as a percentage; 0-100
	Document         *Document       `json:"document,omitempty"`           // Document with the wallpaper
	IsBlurred        *bool           `json:"is_blurred,omitempty"`         // Optional. True, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12
	IsMoving         *bool           `json:"is_moving,omitempty"`          // Optional. True, if the background moves slightly when the device is tilted
	Intensity        *int            `json:"intensity,omitempty"`          // Intensity of the pattern when it is shown above the filled background; 0-100
	IsInverted       *bool           `json:"is_inverted,omitempty"`        // Optional. True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
	ThemeName        *string         `json:"theme_name,omitempty"`         // Name of the chat theme, which is usually an emoji
}

const (
	BackgroundFillSolid            = "solid"
	BackgroundFillGradient         = "gradient"
	BackgroundFillFreeformGradient = "freeform_gradient"
)

type BackgroundFill struct {
	Type          string `json:"type"`           // Type of the background fill. “solid” | “gradient” | “freeform_gradient”
	Color         *int   `json:"color"`          // he color of the background fill in the RGB24 format
	TopColor      *int   `json:"top_color"`      // Top color of the gradient in the RGB24 format
	BottomColor   *int   `json:"bottom_color"`   // Bottom color of the gradient in the RGB24 format
	RotationAngle *int   `json:"rotation_angle"` // Clockwise rotation angle of the background fill in degrees; 0-359
	Colors        *[]int `json:"colors"`         // A list of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24 format
}
