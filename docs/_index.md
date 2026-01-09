---
# *** WARNING: This file was auto-generated. Do not edit by hand unless you're certain you know what you are doing! ***
title: Groundcover Provider
meta_desc: Provides an overview on how to configure the Pulumi Groundcover provider.
layout: package
---

## Installation

The Groundcover provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/groundcover`](https://www.npmjs.com/package/@pulumi/groundcover)
* Python: [`pulumi-groundcover`](https://pypi.org/project/pulumi-groundcover/)
* Go: [`github.com/groundcover-com/pulumi-groundcover/sdk/go/groundcover`](https://github.com/pulumi/pulumi-groundcover)
* .NET: [`Pulumi.Groundcover`](https://www.nuget.org/packages/Pulumi.Groundcover)
* Java: [`com.pulumi/groundcover`](https://central.sonatype.com/artifact/com.pulumi/groundcover)

## Overview

Pulumi provider for managing groundcover resources.
## Configuration Reference

- `apiKey` (String, Sensitive) groundcover API Key. Can also be set via the GROUNDCOVER_API_KEY environment variable.
- `apiUrl` (String) groundcover API URL. Defaults to the groundcover production URL. Can also be set via the GROUNDCOVER_API_URL environment variable.
- `backendId` (String) groundcover Backend ID. Can also be set via the GROUNDCOVER_BACKEND_ID environment variable.
- `orgName` (String) groundcover Organization Name. Can also be set via the GROUNDCOVER_ORG_NAME environment variable. Deprecated: Use backendId instead.