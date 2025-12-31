import * as pulumi from "@pulumi/pulumi";
import * as groundcover from "@groundcover-com/groundcover";

const policy = new groundcover.Policy("test-policy", {
  name: "pulumi-test-policy",
  description: "Test policy created by Pulumi acceptance tests",
  role: {
    read: "read",
  },
});

const serviceAccount = new groundcover.Serviceaccount("test-sa", {
  name: "pulumi-test-service-account",
  email: "pulumi-test-sa@example.com",
  description: "Test service account created by Pulumi acceptance tests",
  policyUuids: [policy.uuid],
});

const apiKey = new groundcover.Apikey("test-apikey", {
  name: "pulumi-test-api-key",
  description: "Test API key created by Pulumi acceptance tests",
  serviceAccountId: serviceAccount.id,
});

export const policyUuid = policy.uuid;
export const serviceAccountId = serviceAccount.id;
export const apiKeyId = apiKey.id;
