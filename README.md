![Testkube Logo](https://raw.githubusercontent.com/kubeshop/testkube/main/assets/logo-dark-text-full.png)

![Release](https://img.shields.io/github/v/release/kubeshop/testkube) [![Releases](https://img.shields.io/github/downloads/kubeshop/testkube/total.svg)](https://github.com/kubeshop/testkube/tags?label=Downloads) ![Go version](https://img.shields.io/github/go-mod/go-version/kubeshop/testkube)

![Docker builds](https://img.shields.io/docker/automated/kubeshop/testkube-api-server) ![Code build](https://img.shields.io/github/workflow/status/kubeshop/testkube/Code%20build%20and%20checks) ![Release date](https://img.shields.io/github/release-date/kubeshop/testkube)

![Twitter](https://img.shields.io/twitter/follow/thekubeshop?style=social) ![Discord](https://img.shields.io/discord/884464549347074049)
 #### [Documentation](https://kubeshop.github.io/testkube) | [Discord](https://discord.gg/hfq44wtR6Q) 


<!-- try to enable it after snyk resolves https://github.com/snyk/snyk/issues/347

Known vulnerabilities: [![Testkube](https://snyk.io/test/github/kubeshop/testkube/badge.svg)](https://snyk.io/test/github/kubeshop/testkube)
[![testkube-operator](https://snyk.io/test/github/kubeshop/testkube-operator/badge.svg)](https://snyk.io/test/github/kubeshop/testkube-operator)
[![helm-charts](https://snyk.io/test/github/kubeshop/helm-charts/badge.svg)](https://snyk.io/test/github/kubeshop/helm-charts)
-->

# Welcome to Testkube
Your friendly cloud-native framework for reliable software testing in Kubernetes.
Testkube natively integrates tests into the cluster by making it very simple to configure the testing phase of your CI/CD while adding lots of extra capabilities.

Testkube decouples test artifacts and execution from CI/CD tooling; tests are meant to be part of your clusters state and can be executed as needed:
- Kubectl pluggin
- Externally triggered via API (CI, external tooling, etc)
- Automatically on deployment of annotated/labeled services/pods/etc (WIP)

Testkube advantages:
- Avoid vendor lock-in for test orchestration and execution in CI/CD  pipelines
- Makes it easy to orchestrate and run any kind of tests - functional, load/performance, security, compliance, etc. in your clusters, without having to wrap them in docker-images or providing network access
- Makes it possible to decouple test execution from build processes; engineers should be able to run specific tests whenever needed
- Centralize all test results in a consistent format for "actionable QA analytics"
- Provide a modular architecture for adding new types of test scripts and executors

Main TestKube components are:

- kubectl Testkube plugin - simple - installed w/o 3rd party repositories (like Krew etc), communicates with
- API Server - work orchestrator, runs executors, gather execution results
- [CRDs Operator](https://github.com/kubeshop/testkube-operator) - watches TestKube CR, handles changes, communicates with API Server
- Executors - runs tests defined for specific runner
  - [Postman Executor](https://github.com/kubeshop/testkube-executor-postman) - runs Postman Collections
  - [Cypress Executor](https://github.com/kubeshop/testkube-executor-cypress) - runs Cypress Tests
  - [Curl Executor](https://github.com/kubeshop/testkube-executor-curl) - runs simple Curl commands
  - [Executor Template](https://github.com/kubeshop/testkube-executor-template) - for creating your own executors
- Results DB - for centralized test results aggregation and analysis
- [Testkube Dashboard](https://github.com/kubeshop/testkube-dashboard) - standalone web application for viewing real-time Testkube test results


## Getting Started

Check out the [Installation](https://kubeshop.github.io/testkube/installing/) and
[Getting Started](https://kubeshop.github.io/testkube/getting-started/) guides to set up Testkube and
run your first tests!

# Discord

Don't hesitate to say hi to the team and ask questions on our [Discord server](https://discord.gg/6zupCZFQbe).

# Documentation

Is available at [https://kubeshop.github.io/testkube](https://kubeshop.github.io/testkube)

## Contributing

Go to [contribution document](CONTRIBUTING.md) to read more how can you help us 🔥

# Feedback

Whether it helps you or not - we'd LOVE to hear from you.  Please let us know what you think and of course, how we can make it better.
