package models

import (
	"errors"
	"time"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {

	// Id uint `json:"id" gorm:"primary_key"`
	Username string    `gorm:"size:255;not null;unique" json:"username"`
	ID       uuid.UUID `json:"-"`
	Birthday time.Time
}

func (user *User) SaveUser() (*User, error) {
	var err error
	err = DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave() error {

	return nil
}

// Verify password correct or not
func VerifyPassword(password, hasedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
}

// Check user exist or not and return token
//func LoginCheck(username string, password string) (string, error) {
//	var err error
//	user := User{}
//	err = DB.Model(User{}).Where("username = ?", username).Take(&user).Error
//	if err != nil {
//		return "", err
//	}
//	// typed password compare with user's database password
//	//err = VerifyPassword(password, user.Password)
//
//	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
//		return "", err
//	}
//	//token, err := token.GenerateToken(user.ID)
//	//fmt.Println("verify token", token, err)
//	if err != nil {
//		return "", err
//	}
//	//return nil, nil
//}

// Get user by ID

func GetUserByID(uid uint) (User, error) {
	var user User
	if err := DB.First(&user, uid).Error; err != nil {
		return user, errors.New("User not found")
	}
	//user.PrepareGive()
	return user, nil
}

//func (user *User) PrepareGive() {
//	user.Password = ""
//}

type UserData struct {
	Users struct {
		ID                      string `json:"id"`
		Username                string `json:"username"`
		ScreenName              string `json:"screenName"`
		IsPrivate               string `json:"isPrivate"`
		IsProtected             string `json:"isProtected"`
		CreatedAt               string `json:"createdAt"`
		UpdatedAt               string `json:"updatedAt"`
		Type                    string `json:"type"`
		Description             string `json:"description"`
		ProfilePictureLargeURL  string `json:"profilePictureLargeUrl"`
		ProfilePictureMediumURL string `json:"profilePictureMediumUrl"`
		Email                   string `json:"email"`
		FrontendPreferences     struct {
			NetFreefeed struct {
				DisplayNames struct {
					DisplayOption int  `json:"displayOption"`
					UseYou        bool `json:"useYou"`
				} `json:"displayNames"`
				RealtimeActive bool `json:"realtimeActive"`
				Comments       struct {
					OmitRepeatedBubbles bool `json:"omitRepeatedBubbles"`
					HighlightComments   bool `json:"highlightComments"`
					ShowTimestamps      bool `json:"showTimestamps"`
					HideRepliesToBanned bool `json:"hideRepliesToBanned"`
				} `json:"comments"`
				AllowLinksPreview bool   `json:"allowLinksPreview"`
				ReadMoreStyle     string `json:"readMoreStyle"`
				HomeFeedSort      string `json:"homeFeedSort"`
				HomeFeedMode      string `json:"homeFeedMode"`
				Homefeed          struct {
					HideUsers []any `json:"hideUsers"`
					HideTags  []any `json:"hideTags"`
				} `json:"homefeed"`
				HidesInNonHomeFeeds     bool     `json:"hidesInNonHomeFeeds"`
				PinnedGroups            []string `json:"pinnedGroups"`
				HideUnreadNotifications bool     `json:"hideUnreadNotifications"`
				TimeDisplay             struct {
					Absolute bool `json:"absolute"`
					AmPm     bool `json:"amPm"`
				} `json:"timeDisplay"`
				TimeDifferenceForSpacer int `json:"timeDifferenceForSpacer"`
			} `json:"net.freefeed"`
			MeFreefeed struct {
				DisplayNames struct {
					DisplayOption int  `json:"displayOption"`
					UseYou        bool `json:"useYou"`
				} `json:"displayNames"`
				RealtimeActive bool `json:"realtimeActive"`
				Comments       struct {
					OmitRepeatedBubbles bool `json:"omitRepeatedBubbles"`
					HighlightComments   bool `json:"highlightComments"`
					ShowTimestamps      bool `json:"showTimestamps"`
					HideRepliesToBanned bool `json:"hideRepliesToBanned"`
				} `json:"comments"`
				AllowLinksPreview bool   `json:"allowLinksPreview"`
				ReadMoreStyle     string `json:"readMoreStyle"`
				HomeFeedSort      string `json:"homeFeedSort"`
				HomeFeedMode      string `json:"homeFeedMode"`
				Homefeed          struct {
					HideUsers []string `json:"hideUsers"`
					HideTags  []any    `json:"hideTags"`
				} `json:"homefeed"`
				HidesInNonHomeFeeds     bool     `json:"hidesInNonHomeFeeds"`
				PinnedGroups            []string `json:"pinnedGroups"`
				HideUnreadNotifications bool     `json:"hideUnreadNotifications"`
				TimeDisplay             struct {
					Absolute bool `json:"absolute"`
					AmPm     bool `json:"amPm"`
				} `json:"timeDisplay"`
				TimeDifferenceForSpacer int `json:"timeDifferenceForSpacer"`
			} `json:"me.freefeed"`
		} `json:"frontendPreferences"`
		PrivateMeta struct {
			Vote2020 string `json:"vote2020"`
		} `json:"privateMeta"`
		Preferences struct {
			HideCommentsOfTypes     []int  `json:"hideCommentsOfTypes"`
			SendNotificationsDigest bool   `json:"sendNotificationsDigest"`
			SendDailyBestOfDigest   bool   `json:"sendDailyBestOfDigest"`
			SendWeeklyBestOfDigest  bool   `json:"sendWeeklyBestOfDigest"`
			AcceptDirectsFrom       string `json:"acceptDirectsFrom"`
			SanitizeMediaMetadata   bool   `json:"sanitizeMediaMetadata"`
		} `json:"preferences"`
		BanIds                    []string `json:"banIds"`
		UnreadDirectsNumber       string   `json:"unreadDirectsNumber"`
		UnreadNotificationsNumber int      `json:"unreadNotificationsNumber"`
		Statistics                struct {
			Posts         string `json:"posts"`
			Likes         string `json:"likes"`
			Comments      string `json:"comments"`
			Subscribers   string `json:"subscribers"`
			Subscriptions string `json:"subscriptions"`
		} `json:"statistics"`
		YouCan                      []string `json:"youCan"`
		TheyDid                     []any    `json:"theyDid"`
		PendingGroupRequests        bool     `json:"pendingGroupRequests"`
		PendingSubscriptionRequests []any    `json:"pendingSubscriptionRequests"`
		SubscriptionRequests        []any    `json:"subscriptionRequests"`
		Subscriptions               []string `json:"subscriptions"`
		Subscribers                 []struct {
			ID                      string `json:"id"`
			Username                string `json:"username"`
			ScreenName              string `json:"screenName"`
			IsPrivate               string `json:"isPrivate"`
			IsProtected             string `json:"isProtected"`
			CreatedAt               string `json:"createdAt"`
			UpdatedAt               string `json:"updatedAt"`
			Type                    string `json:"type"`
			Description             string `json:"description"`
			ProfilePictureLargeURL  string `json:"profilePictureLargeUrl"`
			ProfilePictureMediumURL string `json:"profilePictureMediumUrl"`
			Statistics              struct {
				Posts         string `json:"posts"`
				Likes         string `json:"likes"`
				Comments      string `json:"comments"`
				Subscribers   string `json:"subscribers"`
				Subscriptions string `json:"subscriptions"`
			} `json:"statistics"`
			YouCan  []string `json:"youCan"`
			TheyDid []string `json:"theyDid"`
			IsGone  bool     `json:"isGone,omitempty"`
		} `json:"subscribers"`
	} `json:"users"`
	Subscribers []struct {
		ID                      string `json:"id"`
		Username                string `json:"username"`
		ScreenName              string `json:"screenName"`
		IsPrivate               string `json:"isPrivate"`
		IsProtected             string `json:"isProtected"`
		CreatedAt               string `json:"createdAt"`
		UpdatedAt               string `json:"updatedAt"`
		Type                    string `json:"type"`
		Description             string `json:"description"`
		ProfilePictureLargeURL  string `json:"profilePictureLargeUrl"`
		ProfilePictureMediumURL string `json:"profilePictureMediumUrl"`
		IsRestricted            string `json:"isRestricted,omitempty"`
		Statistics              struct {
			Posts         string `json:"posts"`
			Likes         string `json:"likes"`
			Comments      string `json:"comments"`
			Subscribers   string `json:"subscribers"`
			Subscriptions string `json:"subscriptions"`
		} `json:"statistics"`
		Administrators []string `json:"administrators,omitempty"`
		YouCan         []string `json:"youCan"`
		TheyDid        []any    `json:"theyDid"`
		IsGone         bool     `json:"isGone,omitempty"`
	} `json:"subscribers"`
	Subscriptions []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		User string `json:"user"`
	} `json:"subscriptions"`
	Requests      []any `json:"requests"`
	ManagedGroups []struct {
		ID                      string `json:"id"`
		Username                string `json:"username"`
		ScreenName              string `json:"screenName"`
		IsPrivate               string `json:"isPrivate"`
		IsProtected             string `json:"isProtected"`
		CreatedAt               string `json:"createdAt"`
		UpdatedAt               string `json:"updatedAt"`
		Type                    string `json:"type"`
		Description             string `json:"description"`
		ProfilePictureLargeURL  string `json:"profilePictureLargeUrl"`
		ProfilePictureMediumURL string `json:"profilePictureMediumUrl"`
		IsRestricted            string `json:"isRestricted"`
		Statistics              struct {
			Posts         string `json:"posts"`
			Likes         string `json:"likes"`
			Comments      string `json:"comments"`
			Subscribers   string `json:"subscribers"`
			Subscriptions string `json:"subscriptions"`
		} `json:"statistics"`
		Administrators []string `json:"administrators"`
		YouCan         []string `json:"youCan"`
		TheyDid        []any    `json:"theyDid"`
		Requests       []any    `json:"requests"`
	} `json:"managedGroups"`
}
