// Code generated by ent, DO NOT EDIT.

package media

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ofdl/ofdl/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldID, id))
}

// PostID applies equality check predicate on the "post_id" field. It's identical to PostIDEQ.
func PostID(v int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldPostID, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldType, v))
}

// Full applies equality check predicate on the "full" field. It's identical to FullEQ.
func Full(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldFull, v))
}

// DownloadedAt applies equality check predicate on the "downloaded_at" field. It's identical to DownloadedAtEQ.
func DownloadedAt(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldDownloadedAt, v))
}

// StashID applies equality check predicate on the "stash_id" field. It's identical to StashIDEQ.
func StashID(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldStashID, v))
}

// OrganizedAt applies equality check predicate on the "organized_at" field. It's identical to OrganizedAtEQ.
func OrganizedAt(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldOrganizedAt, v))
}

// PostIDEQ applies the EQ predicate on the "post_id" field.
func PostIDEQ(v int) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldPostID, v))
}

// PostIDNEQ applies the NEQ predicate on the "post_id" field.
func PostIDNEQ(v int) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldPostID, v))
}

// PostIDIn applies the In predicate on the "post_id" field.
func PostIDIn(vs ...int) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldPostID, vs...))
}

// PostIDNotIn applies the NotIn predicate on the "post_id" field.
func PostIDNotIn(vs ...int) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldPostID, vs...))
}

// PostIDGT applies the GT predicate on the "post_id" field.
func PostIDGT(v int) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldPostID, v))
}

// PostIDGTE applies the GTE predicate on the "post_id" field.
func PostIDGTE(v int) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldPostID, v))
}

// PostIDLT applies the LT predicate on the "post_id" field.
func PostIDLT(v int) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldPostID, v))
}

// PostIDLTE applies the LTE predicate on the "post_id" field.
func PostIDLTE(v int) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldPostID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Media {
	return predicate.Media(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Media {
	return predicate.Media(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Media {
	return predicate.Media(sql.FieldContainsFold(FieldType, v))
}

// FullEQ applies the EQ predicate on the "full" field.
func FullEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldFull, v))
}

// FullNEQ applies the NEQ predicate on the "full" field.
func FullNEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldFull, v))
}

// FullIn applies the In predicate on the "full" field.
func FullIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldFull, vs...))
}

// FullNotIn applies the NotIn predicate on the "full" field.
func FullNotIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldFull, vs...))
}

// FullGT applies the GT predicate on the "full" field.
func FullGT(v string) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldFull, v))
}

// FullGTE applies the GTE predicate on the "full" field.
func FullGTE(v string) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldFull, v))
}

// FullLT applies the LT predicate on the "full" field.
func FullLT(v string) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldFull, v))
}

// FullLTE applies the LTE predicate on the "full" field.
func FullLTE(v string) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldFull, v))
}

// FullContains applies the Contains predicate on the "full" field.
func FullContains(v string) predicate.Media {
	return predicate.Media(sql.FieldContains(FieldFull, v))
}

// FullHasPrefix applies the HasPrefix predicate on the "full" field.
func FullHasPrefix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasPrefix(FieldFull, v))
}

// FullHasSuffix applies the HasSuffix predicate on the "full" field.
func FullHasSuffix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasSuffix(FieldFull, v))
}

// FullEqualFold applies the EqualFold predicate on the "full" field.
func FullEqualFold(v string) predicate.Media {
	return predicate.Media(sql.FieldEqualFold(FieldFull, v))
}

// FullContainsFold applies the ContainsFold predicate on the "full" field.
func FullContainsFold(v string) predicate.Media {
	return predicate.Media(sql.FieldContainsFold(FieldFull, v))
}

// DownloadedAtEQ applies the EQ predicate on the "downloaded_at" field.
func DownloadedAtEQ(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldDownloadedAt, v))
}

// DownloadedAtNEQ applies the NEQ predicate on the "downloaded_at" field.
func DownloadedAtNEQ(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldDownloadedAt, v))
}

// DownloadedAtIn applies the In predicate on the "downloaded_at" field.
func DownloadedAtIn(vs ...time.Time) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldDownloadedAt, vs...))
}

// DownloadedAtNotIn applies the NotIn predicate on the "downloaded_at" field.
func DownloadedAtNotIn(vs ...time.Time) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldDownloadedAt, vs...))
}

// DownloadedAtGT applies the GT predicate on the "downloaded_at" field.
func DownloadedAtGT(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldDownloadedAt, v))
}

// DownloadedAtGTE applies the GTE predicate on the "downloaded_at" field.
func DownloadedAtGTE(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldDownloadedAt, v))
}

// DownloadedAtLT applies the LT predicate on the "downloaded_at" field.
func DownloadedAtLT(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldDownloadedAt, v))
}

// DownloadedAtLTE applies the LTE predicate on the "downloaded_at" field.
func DownloadedAtLTE(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldDownloadedAt, v))
}

// DownloadedAtIsNil applies the IsNil predicate on the "downloaded_at" field.
func DownloadedAtIsNil() predicate.Media {
	return predicate.Media(sql.FieldIsNull(FieldDownloadedAt))
}

// DownloadedAtNotNil applies the NotNil predicate on the "downloaded_at" field.
func DownloadedAtNotNil() predicate.Media {
	return predicate.Media(sql.FieldNotNull(FieldDownloadedAt))
}

// StashIDEQ applies the EQ predicate on the "stash_id" field.
func StashIDEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldStashID, v))
}

// StashIDNEQ applies the NEQ predicate on the "stash_id" field.
func StashIDNEQ(v string) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldStashID, v))
}

// StashIDIn applies the In predicate on the "stash_id" field.
func StashIDIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldStashID, vs...))
}

// StashIDNotIn applies the NotIn predicate on the "stash_id" field.
func StashIDNotIn(vs ...string) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldStashID, vs...))
}

// StashIDGT applies the GT predicate on the "stash_id" field.
func StashIDGT(v string) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldStashID, v))
}

// StashIDGTE applies the GTE predicate on the "stash_id" field.
func StashIDGTE(v string) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldStashID, v))
}

// StashIDLT applies the LT predicate on the "stash_id" field.
func StashIDLT(v string) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldStashID, v))
}

// StashIDLTE applies the LTE predicate on the "stash_id" field.
func StashIDLTE(v string) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldStashID, v))
}

// StashIDContains applies the Contains predicate on the "stash_id" field.
func StashIDContains(v string) predicate.Media {
	return predicate.Media(sql.FieldContains(FieldStashID, v))
}

// StashIDHasPrefix applies the HasPrefix predicate on the "stash_id" field.
func StashIDHasPrefix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasPrefix(FieldStashID, v))
}

// StashIDHasSuffix applies the HasSuffix predicate on the "stash_id" field.
func StashIDHasSuffix(v string) predicate.Media {
	return predicate.Media(sql.FieldHasSuffix(FieldStashID, v))
}

// StashIDEqualFold applies the EqualFold predicate on the "stash_id" field.
func StashIDEqualFold(v string) predicate.Media {
	return predicate.Media(sql.FieldEqualFold(FieldStashID, v))
}

// StashIDContainsFold applies the ContainsFold predicate on the "stash_id" field.
func StashIDContainsFold(v string) predicate.Media {
	return predicate.Media(sql.FieldContainsFold(FieldStashID, v))
}

// OrganizedAtEQ applies the EQ predicate on the "organized_at" field.
func OrganizedAtEQ(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldEQ(FieldOrganizedAt, v))
}

// OrganizedAtNEQ applies the NEQ predicate on the "organized_at" field.
func OrganizedAtNEQ(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldNEQ(FieldOrganizedAt, v))
}

// OrganizedAtIn applies the In predicate on the "organized_at" field.
func OrganizedAtIn(vs ...time.Time) predicate.Media {
	return predicate.Media(sql.FieldIn(FieldOrganizedAt, vs...))
}

// OrganizedAtNotIn applies the NotIn predicate on the "organized_at" field.
func OrganizedAtNotIn(vs ...time.Time) predicate.Media {
	return predicate.Media(sql.FieldNotIn(FieldOrganizedAt, vs...))
}

// OrganizedAtGT applies the GT predicate on the "organized_at" field.
func OrganizedAtGT(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldGT(FieldOrganizedAt, v))
}

// OrganizedAtGTE applies the GTE predicate on the "organized_at" field.
func OrganizedAtGTE(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldGTE(FieldOrganizedAt, v))
}

// OrganizedAtLT applies the LT predicate on the "organized_at" field.
func OrganizedAtLT(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldLT(FieldOrganizedAt, v))
}

// OrganizedAtLTE applies the LTE predicate on the "organized_at" field.
func OrganizedAtLTE(v time.Time) predicate.Media {
	return predicate.Media(sql.FieldLTE(FieldOrganizedAt, v))
}

// OrganizedAtIsNil applies the IsNil predicate on the "organized_at" field.
func OrganizedAtIsNil() predicate.Media {
	return predicate.Media(sql.FieldIsNull(FieldOrganizedAt))
}

// OrganizedAtNotNil applies the NotNil predicate on the "organized_at" field.
func OrganizedAtNotNil() predicate.Media {
	return predicate.Media(sql.FieldNotNull(FieldOrganizedAt))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Media) predicate.Media {
	return predicate.Media(func(s *sql.Selector) {
		p(s.Not())
	})
}
