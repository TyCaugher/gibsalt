package main

type E6Response struct {
	Posts []struct {
		ID        int    `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		File      struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			Ext    string `json:"ext"`
			Size   int    `json:"size"`
			Md5    string `json:"md5"`
			URL    string `json:"url"`
		} `json:"file"`
		Preview struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			URL    string `json:"url"`
		} `json:"preview"`
		Sample struct {
			Has        bool   `json:"has"`
			Height     int    `json:"height"`
			Width      int    `json:"width"`
			URL        string `json:"url"`
			Alternates struct {
			} `json:"alternates"`
		} `json:"sample"`
		Score struct {
			Up    int `json:"up"`
			Down  int `json:"down"`
			Total int `json:"total"`
		} `json:"score"`
		Tags struct {
			General   []string      `json:"general"`
			Species   []string      `json:"species"`
			Character []string      `json:"character"`
			Copyright []string      `json:"copyright"`
			Artist    []string      `json:"artist"`
			Invalid   []interface{} `json:"invalid"`
			Lore      []interface{} `json:"lore"`
			Meta      []string      `json:"meta"`
		} `json:"tags"`
		LockedTags []interface{} `json:"locked_tags"`
		ChangeSeq  int           `json:"change_seq"`
		Flags      struct {
			Pending         bool `json:"pending"`
			Flagged         bool `json:"flagged"`
			NoteLocked      bool `json:"note_locked"`
			StatusLocked    bool `json:"status_locked"`
			RatingLocked    bool `json:"rating_locked"`
			CommentDisabled bool `json:"comment_disabled"`
			Deleted         bool `json:"deleted"`
		} `json:"flags"`
		Rating        string        `json:"rating"`
		FavCount      int           `json:"fav_count"`
		Sources       []string      `json:"sources"`
		Pools         []interface{} `json:"pools"`
		Relationships struct {
			ParentID          interface{}   `json:"parent_id"`
			HasChildren       bool          `json:"has_children"`
			HasActiveChildren bool          `json:"has_active_children"`
			Children          []interface{} `json:"children"`
		} `json:"relationships"`
		ApproverID   interface{} `json:"approver_id"`
		UploaderID   int         `json:"uploader_id"`
		Description  string      `json:"description"`
		CommentCount int         `json:"comment_count"`
		IsFavorited  bool        `json:"is_favorited"`
		HasNotes     bool        `json:"has_notes"`
		Duration     interface{} `json:"duration"`
	} `json:"posts"`
}
