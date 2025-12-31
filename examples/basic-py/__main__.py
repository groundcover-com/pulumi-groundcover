import pulumi
import groundcover_groundcover as groundcover

policy = groundcover.Policy(
    "test-policy",
    name="pulumi-test-policy-py",
    description="Test policy created by Pulumi acceptance tests (Python)",
    role={"read": "read"},
)

service_account = groundcover.Serviceaccount(
    "test-sa",
    name="pulumi-test-service-account-py",
    email="pulumi-test-sa-py@example.com",
    description="Test service account created by Pulumi acceptance tests (Python)",
    policy_uuids=[policy.uuid],
)

api_key = groundcover.Apikey(
    "test-apikey",
    name="pulumi-test-api-key-py",
    description="Test API key created by Pulumi acceptance tests (Python)",
    service_account_id=service_account.id,
)

pulumi.export("policy_uuid", policy.uuid)
pulumi.export("service_account_id", service_account.id)
pulumi.export("api_key_id", api_key.id)
