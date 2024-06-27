package models

type Action string

const (
	ActionView           Action = "View"
	ActionAddToCart      Action = "AddToCart"
	ActionPurchase       Action = "Purchase"
	ActionRemoveFromCart Action = "RemoveFromCart"
	ActionWishlist       Action = "Wishlist"
	ActionSearch         Action = "Search"
	ActionRate           Action = "Rate"
	ActionReview         Action = "Review"
	ActionClick          Action = "Click"
	ActionShare          Action = "Share"
)

func (a Action) IsValid() bool {
	switch a {
	case ActionView, ActionAddToCart, ActionPurchase, ActionRemoveFromCart, ActionWishlist, ActionSearch, ActionRate, ActionReview, ActionClick, ActionShare:
		return true
	}
	return false
}

func StringParseAction(action string) Action {
	switch action {
	case "View":
		return ActionView
	case "AddToCart":
		return ActionAddToCart
	case "Purchase":
		return ActionPurchase
	case "RemoveFromCart":
		return ActionRemoveFromCart
	case "Wishlist":
		return ActionWishlist
	case "Search":
		return ActionSearch
	case "Rate":
		return ActionRate
	case "Review":
		return ActionReview
	case "Click":
		return ActionClick
	case "Share":
		return ActionShare
	}
	return ""
}
