package graph

import (
	"context"
	"testing"

	"github.com/well-informed/wellinformed/graph/model"
)

func TestRegister(t *testing.T) {
	resolver, _ := NewTestHarness("register")
	_, err := RegisterMutation(resolver, nil)
	if err != nil {
		t.Error("failed to register. err: ", err)
	}
}

func TestRegisterAndLogin(t *testing.T) {
	resolver, _ := NewTestHarness("register_and_login")
	email := "danielveenstra@protonmail.com"
	password := "ScoobyDoo69"
	registerInput := &model.RegisterInput{
		Username:        "deviator",
		Email:           email,
		Password:        password,
		ConfirmPassword: password,
		Firstname:       "Dan",
		Lastname:        "Veenstra",
	}
	loginInput := model.LoginInput{
		Email:    email,
		Password: password,
	}
	_, err := RegisterMutation(resolver, registerInput)
	if err != nil {
		t.Error("couldn't register. err: ", err)
	}
	auth, err := resolver.Mutation().Login(context.Background(), loginInput)
	if err != nil {
		t.Error("couldn't log in. err: ", err)
	}
	if auth.AuthToken == nil {
		t.Error("auth token should not be nil")
	}
	if auth.User.ID == 0 {
		t.Errorf("userID should be nonzero")
	}
}

func TestAddSrcRSSFeed(t *testing.T) {
	resolver, _ := NewTestHarness("add_src_rss_feed")
	authResponse, err := RegisterMutation(resolver, nil)
	if err != nil {
		t.Error("could not register. err: ", err)
	}
	_, ctx := NewMockAuthenticatedContext(resolver.DB, authResponse.User.ID)

	srcRSSFeeds, errs := AddStockSrcRSSFeeds(resolver, ctx)
	for i, srcRSSFeed := range srcRSSFeeds {
		if errs[i] != nil {
			t.Fatal("could not add srcRSSFeed. err: ", err)
		}
		if srcRSSFeed.ID == 0 {
			t.Errorf("missing ID")
		}
		if srcRSSFeed.Title == "" {
			t.Errorf("missing title")
		}
	}
}

func TestAddUserFeed(t *testing.T) {
	resolver, _, ctx := NewAuthedUserTestEnv("add_user_feed")
	name := "FEEDME"
	input := model.AddUserFeedInput{
		Name: name,
	}
	userFeed, err := resolver.Mutation().AddUserFeed(ctx, input)
	if err != nil {
		t.Error("error creating new userFeed")
	}
	if userFeed.ID == 0 {
		t.Error("userFeed ID has invalid value 0")
	}
	if userFeed.EngineID == 0 {
		t.Error("userFeed engineID has invalid value 0")
	}
	if userFeed.Name != name {
		t.Errorf("userFeed name should be %v but was %v", name, userFeed.Name)
	}
}

func TestAddSource(t *testing.T) {
	//Register 2 users
	//Subscribe to a source with User A
	//Subscribe to User A's new feed with User B
	resolver, _ := NewTestHarness("add_source")
	userAEmail := "dude@bro.com"
	userAPass := "something"
	userARegisterInput := model.RegisterInput{
		Username:        "dudebro",
		Email:           userAEmail,
		Password:        userAPass,
		ConfirmPassword: userAPass,
		Firstname:       "Dude",
		Lastname:        "Bro",
	}

	userBEmail := "chick@lady.com"
	userBPass := "mydogsname"
	userBRegisterInput := model.RegisterInput{
		Username:        "chicklady",
		Email:           userBEmail,
		Password:        userBPass,
		ConfirmPassword: userBPass,
		Firstname:       "Chick",
		Lastname:        "Lady",
	}
	authUserA, err := resolver.Mutation().Register(context.Background(), userARegisterInput)
	if err != nil {
		t.Error("couldn't register userA")
	}
	_, Actx := NewMockAuthenticatedContext(resolver.DB, authUserA.User.ID)
	authUserB, err := resolver.Mutation().Register(context.Background(), userBRegisterInput)
	if err != nil {
		t.Error("couldn't register userB")
	}
	userB, Bctx := NewMockAuthenticatedContext(resolver.DB, authUserB.User.ID)

	cryptoFeed, err := resolver.Mutation().AddSrcRSSFeed(Actx, "https://bankless.substack.com/feed")
	if err != nil {
		t.Error("couldn't add src to user A's feed. err: ", err)
	}

	input := model.AddSourceInput{
		SourceFeedID: cryptoFeed.ID,
		SourceType:   model.SourceTypeUserFeed,
		TargetFeedID: &userB.ActiveUserFeedID, //Problem is this attribute doesn't exist in the graphql schema
	}
	BFeedSubsription, err := resolver.Mutation().AddSource(Bctx, input)
	if err != nil {
		t.Error("could not add User A's userFeed to active feed. err: ", err)
	}
	if BFeedSubsription.ID == 0 {
		t.Error("B feed subscription had invalid value 0")
	}
	if BFeedSubsription.SourceID != cryptoFeed.ID {
		t.Errorf("subscription source ID %v does not equal intended source ID %v", BFeedSubsription.SourceID, cryptoFeed.ID)
	}
	//Get content from feed, make sure it includes original srcRSSFeed
	userFeed, err := resolver.Query().UserFeed(Bctx)
	if err != nil {
		t.Error("could not get User B's userFeed. err: ", err)
	}
	contentItemConnInput := model.ContentItemConnectionInput{
		First: 10,
	}
	contentItemConn, err := resolver.UserFeed().ContentItems(Bctx, userFeed, contentItemConnInput)
	if err != nil {
		t.Error("could not serve user's content items. err: ", err)
	}
	if len(contentItemConn.Edges) == 0 {
		t.Errorf("contentItemConnections is empty")
	} else {
		//Check that served content item came from original srcRSSFeed
		contentSourceID := contentItemConn.Edges[0].Node.SourceID
		if contentSourceID != cryptoFeed.ID {
			t.Errorf("served content source ID %v did not equal original added ID %v", contentSourceID, cryptoFeed.ID)
		}
	}

}
