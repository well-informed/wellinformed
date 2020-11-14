package graph

import (
	"errors"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/well-informed/wellinformed/graph/model"
)

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
	//TODO: shave off .xml from end of these strings so that strings match was src library stores in database
	existingFeed, err := r.DB.GetSrcRSSFeed(model.SrcRSSFeedInput{FeedLink: &feedLink})
	if err != nil {
		return nil, false, err
	}
	if existingFeed != nil {
		log.Debug("found existing RSS feed with link: ", feedLink)
		return existingFeed, true, nil
	}
	extStartIdx := strings.LastIndex(feedLink, ".")
	noExtFeedLink := feedLink[0:extStartIdx]
	existingFeed, err = r.DB.GetSrcRSSFeed(model.SrcRSSFeedInput{FeedLink: &noExtFeedLink})
	if err != nil {
		return nil, false, err
	}
	if existingFeed != nil {
		log.Debug("found existing RSS feed with no extension link: ", noExtFeedLink)
		return existingFeed, true, nil
	}
	return nil, false, nil
}
