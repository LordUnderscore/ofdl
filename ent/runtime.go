// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/ofdl/ofdl/ent/schema"
	"github.com/ofdl/ofdl/ent/subscription"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	subscriptionFields := schema.Subscription{}.Fields()
	_ = subscriptionFields
	// subscriptionDescEnabled is the schema descriptor for enabled field.
	subscriptionDescEnabled := subscriptionFields[7].Descriptor()
	// subscription.DefaultEnabled holds the default value on creation for the enabled field.
	subscription.DefaultEnabled = subscriptionDescEnabled.Default.(bool)
}
