package forms

type RelationActionReq struct {
	ToUserID   int64 `json:"to_user_id" binding:"gte=1,lte=32"`
	ActionType int32 `json:"action_type" binding:"required,oneof=1 2"`
}
