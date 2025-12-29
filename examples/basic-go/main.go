package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-groundcover/sdk/go/groundcover"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		policy, err := groundcover.NewPolicy(ctx, "test-policy", &groundcover.PolicyArgs{
			Name:        pulumi.String("pulumi-test-policy-go"),
			Description: pulumi.String("Test policy created by Pulumi acceptance tests (Go)"),
			Role: pulumi.StringMap{
				"read": pulumi.String("read"),
			},
		})
		if err != nil {
			return err
		}

		serviceAccount, err := groundcover.NewServiceaccount(ctx, "test-sa", &groundcover.ServiceaccountArgs{
			Name:        pulumi.String("pulumi-test-service-account-go"),
			Email:       pulumi.String("pulumi-test-sa-go@example.com"),
			Description: pulumi.String("Test service account created by Pulumi acceptance tests (Go)"),
			PolicyUuids: pulumi.StringArray{
				policy.Uuid,
			},
		})
		if err != nil {
			return err
		}

		apiKey, err := groundcover.NewApikey(ctx, "test-apikey", &groundcover.ApikeyArgs{
			Name:             pulumi.String("pulumi-test-api-key-go"),
			Description:      pulumi.String("Test API key created by Pulumi acceptance tests (Go)"),
			ServiceAccountId: serviceAccount.ID(),
		})
		if err != nil {
			return err
		}

		ctx.Export("policyUuid", policy.Uuid)
		ctx.Export("serviceAccountId", serviceAccount.ID())
		ctx.Export("apiKeyId", apiKey.ID())

		return nil
	})
}
