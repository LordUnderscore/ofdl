package ofdl

import (
	"fmt"
	"time"

	"github.com/ofdl/ofdl/model"
	"github.com/shurcooL/graphql"
	"github.com/spf13/viper"
)

func (o *OFDL) GetUnorganizedSubscriptions(limit int) ([]model.Subscription, error) {
	return o.Query.Subscription.Unorganized(limit)
}

func (o *OFDL) OrganizeSubscription(sub model.Subscription) error {
	lu := &PerformerLookup{}
	if err := o.Stash.Query(o.ctx, lu, map[string]interface{}{
		"name": graphql.String(sub.Username),
	}); err != nil {
		return err
	}

	// If missing, create it
	if lu.FindPerformers.Count == 0 {
		pc := &PerformerCreate{}
		if err := o.Stash.Mutate(o.ctx, pc, map[string]interface{}{
			"performer": PerformerCreateInput{
				"name":  sub.Username,
				"url":   fmt.Sprintf("https://onlyfans.com/%s", sub.Username),
				"image": sub.Avatar,
			},
		}); err != nil {
			return err
		}

		now := time.Now()
		sub.StashID = pc.PerformerCreate.ID
		sub.OrganizedAt = &now
		if err := o.DB.Save(&sub).Error; err != nil {
			return err
		}

		return nil
	}

	// If found, update it
	id := lu.FindPerformers.Performers[0].ID
	pc := &PerformerUpdate{}
	if err := o.Stash.Mutate(o.ctx, pc, map[string]interface{}{
		"performer": PerformerUpdateInput{
			"id":    id,
			"name":  sub.Username,
			"url":   fmt.Sprintf("https://onlyfans.com/%s", sub.Username),
			"image": sub.Avatar,
		},
	}); err != nil {
		return err
	}

	now := time.Now()
	sub.StashID = pc.PerformerUpdate.ID
	sub.OrganizedAt = &now
	if err := o.DB.Save(&sub).Error; err != nil {
		return err
	}

	return nil
}

func (o *OFDL) GetUnorganizedMedia(limit int) ([]model.Media, error) {
	q := o.Query.Media
	return q.Preload(
		q.Post,
		q.Post.Subscription,
	).Unorganized(limit)
}

func (o *OFDL) OrganizeMedia(m model.Media) error {
	if m.Full == "" {
		now := time.Now()
		m.OrganizedAt = &now
		if err := o.DB.Save(&m).Error; err != nil {
			return err
		}
		return nil
	}

	var sm StashLookup
	var sr interface{}

	switch m.Type {
	case "photo":
		sm = &ImageLookup{}
		sr = &ImageUpdate{}
	case "video", "gif":
		sm = &SceneLookup{}
		sr = &SceneUpdate{}
	default:
		return nil
	}

	if err := o.Stash.Query(o.ctx, sm, map[string]interface{}{
		"path": graphql.String(m.Filename()),
	}); err != nil {
		return err
	}

	if sm.Count() != 1 {
		return nil
	}

	// set the text, date, and model?
	vars := map[string]interface{}{
		"id":          sm.Medias()[0].ID,
		"title":       graphql.String(m.Post.Text),
		"performerId": graphql.ID(m.Post.Subscription.StashID),
		"studioId":    graphql.ID(viper.GetString("stash.studio_id")),
	}
	if d, err := time.Parse(time.RFC3339, m.Post.PostedAt); err == nil {
		vars["date"] = graphql.String(d.Format("2006-01-02"))
	} else {
		panic(err)
	}

	if err := o.Stash.Mutate(o.ctx, sr, vars); err != nil {
		return err
	}

	now := time.Now()
	m.StashID = sm.Medias()[0].ID
	m.OrganizedAt = &now
	if err := o.DB.Save(&m).Error; err != nil {
		return err
	}

	return nil
}
