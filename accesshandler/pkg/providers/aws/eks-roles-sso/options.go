package eksrolessso

import (
	"context"

	"github.com/common-fate/granted-approvals/accesshandler/pkg/types"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// List options for arg
func (p *Provider) Options(ctx context.Context, arg string) (*types.ArgOptionsResponse, error) {

	var opts types.ArgOptionsResponse
	hasMore := true
	var nextToken string

	for hasMore {

		roles, err := p.kubeClient.RbacV1().Roles(p.namespace.Get()).List(ctx, v1.ListOptions{Continue: nextToken})
		if err != nil {
			return nil, err
		}
		for _, r := range roles.Items {
			opts.Options = append(opts.Options, types.Option{Label: r.Name, Value: r.Name})
		}
		nextToken = roles.Continue
		//exit the pagination
		if nextToken == "" {
			hasMore = false
		}

	}

	return &opts, nil

}
