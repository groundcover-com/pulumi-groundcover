using System.Collections.Generic;
using Pulumi;
using Groundcover = Pulumiverse.Groundcover;

return await Deployment.RunAsync(() =>
{
    var policy = new Groundcover.Policy("test-policy", new()
    {
        Name = "pulumi-test-policy-dotnet",
        Description = "Test policy created by Pulumi acceptance tests (.NET)",
        Role = 
        {
            { "read", "read" },
        },
    });

    var serviceAccount = new Groundcover.Serviceaccount("test-sa", new()
    {
        Name = "pulumi-test-service-account-dotnet",
        Email = "pulumi-test-sa-dotnet@example.com",
        Description = "Test service account created by Pulumi acceptance tests (.NET)",
        PolicyUuids = new[]
        {
            policy.Uuid,
        },
    });

    var apiKey = new Groundcover.Apikey("test-apikey", new()
    {
        Name = "pulumi-test-api-key-dotnet",
        Description = "Test API key created by Pulumi acceptance tests (.NET)",
        ServiceAccountId = serviceAccount.Id,
    });

    return new Dictionary<string, object?>
    {
        ["policyUuid"] = policy.Uuid,
        ["serviceAccountId"] = serviceAccount.Id,
        ["apiKeyId"] = apiKey.Id,
    };
});
