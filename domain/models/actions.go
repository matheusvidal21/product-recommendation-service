package models

import "strings"

type Action string

const (
	ActionView           Action = "View"
	ActionAddToCart      Action = "AddTocart"
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
	switch strings.ToLower(action) {
	case "view":
		return ActionView
	case "addtocart":
		return ActionAddToCart
	case "purchase":
		return ActionPurchase
	case "removefromcart":
		return ActionRemoveFromCart
	case "wishlist":
		return ActionWishlist
	case "search":
		return ActionSearch
	case "rate":
		return ActionRate
	case "review":
		return ActionReview
	case "click":
		return ActionClick
	case "share":
		return ActionShare
	}
	return ActionView
}
