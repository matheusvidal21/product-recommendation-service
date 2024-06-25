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
