package request

type MenuRequest struct {
	Name      string `json:"name" binding:"required"`
	Url       string `json:"url" binding:"required"`
	Icon      string `json:"icon"`
	Ord       int    `json:"index"`
	ParentID  int    `json:"parent_id"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}
