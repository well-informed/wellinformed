package graph

import (
	"context"
	"errors"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

func (r *mutationResolver) addSrcRSSFeed(ctx context.Context, feedLink string, targetFeedID int64) (*model.SrcRSSFeed, error) {
	var srcRSSFeed *model.SrcRSSFeed

	feedLink, err := cleanUserFeedLinkInput(feedLink)
	if err != nil {
		return nil, err
	}

	//Check if the provided feed URL already exists in the database, and if not, fetch and add it
	srcRSSFeed, exists, err := r.checkForExistingRSSFeed(feedLink)
	if err != nil {
		log.Error("could not check existing srcRSSFeeds for AddSrcRSSFeed. err: ", err)
	}
	if !exists {
		log.Debug("did not find existing SrcRSSFeed, subscribing to new link: ", feedLink)
		srcRSSFeed, err = r.Sub.SubscribeToRSSFeed(ctx, feedLink)
		if err != nil {
			return nil, err
		}
	}

	//Create a subscription for the user who added the new feed
	_, err = r.DB.CreateFeedSubscription(targetFeedID, srcRSSFeed.ID, model.SourceTypeSrcRSSFeed)
	if err != nil {
		log.Errorf("unable to create feed subscription between target feed: %v and sourceRSSFeed: %+v. err: %v", targetFeedID, srcRSSFeed, err)
		return srcRSSFeed, err
	}
	return srcRSSFeed, nil
}

func cleanUserFeedLinkInput(feedLink string) (string, error) {
	//parse and massage string to standardize

	link, err := url.Parse(feedLink)
	if err != nil {
		log.Error("couldn't parse feedLink: ", feedLink)
		return "", errors.New("couldn't parse feedLink")
	}
	if link.Scheme == "" {
		link.Scheme = "https"
	}
	feedLink = link.String()
	return feedLink, nil
}

//Takes a user supplied feedLink and checks if it already exists in the database
//Checks the feedlink with the extension (.xml, .html, etc) stripped and with it kept intact,
//to ensure that a matching feed is found when it should be.
func (r *mutationResolver) checkForExistingRSSFeed(feedLink string) (*model.SrcRSSFeed, bool, error) {
	existingFeed, err := r.DB.GetSrcRSSFeedByFeedLink(feedLink)
	if err != nil {
		return nil, false, err
	}
	if existingFeed != nil {
		log.Debug("found existing RSS feed with link: ", feedLink)
		return existingFeed, true, nil
	}

	//Additionally check for feedLink strings with ending extension (e.g. '.xml') stripped off
	//Required to handle differences between valid links for fetching feed and stored string for feedLink in returned data
	extStartIdx := strings.LastIndex(feedLink, ".")
	noExtFeedLink := feedLink[0:extStartIdx]
	existingFeed, err = r.DB.GetSrcRSSFeedByFeedLink(noExtFeedLink)
	if err != nil {
		return nil, false, err
	}
	if existingFeed != nil {
		log.Debug("found existing RSS feed with no extension link: ", noExtFeedLink)
		return existingFeed, true, nil
	}
	return nil, false, nil
}
