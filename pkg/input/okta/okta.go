package okta

import (
	"context"
	"fmt"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"os"
	"time"
)

type Okta struct {
	client *okta.Client
}

func New() *Okta {
	_, client, err := okta.NewClient(
		context.Background(),
		okta.WithOrgUrl(os.Getenv("OKTA_ORG_URL")),
		okta.WithToken(os.Getenv("OKTA_API_TOKEN")),
	)
	if err != nil {
		panic(err)
	}

	o := &Okta{
		client: client,
	}
	return o
}

func (o *Okta) GetSystemLogEvents() error {
	events, _, err := o.client.LogEvent.GetLogs(context.Background(), &query.Params{
		SortOrder: "ASCENDING",
		Since:     time.Now().Format(time.RFC3339),
		Limit:     1000,
	})
	if err != nil {
		return err
	}
	fmt.Println(events)
	return nil
}
