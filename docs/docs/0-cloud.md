---
sidebar_position: 1
sidebar_label: Testkube Cloud
---
# Testkube Cloud 

Testkube Cloud is managed version of Testkube it's main purpose is to orchestrate multiple clusters. 
All test results and test artifacts are stored into Testkube Cloud internal data storages. Testkube Cloud 
will provide you with additional tests insights and is able to limit access for your users only to subset 
of environments.

Testkube Cloud is in Alpha phase - so feel free to give us feedback! 


# Testkube Cloud Agent - Installation manual

Testkube Cloud is able to connect to Testkube Agents. Testkube Agent is Testkube engine for managing test runs into your cluster
it's also responsible for getting insights into Testkube resources stored into cluster.

Testkube Agent opens networking connection into Testkube Cloud API, which stays active and it's main purpose is to start 
listening for Testkube Cloud commands.

Your existing Open Source Testkube installation can be also converted into Testkube Cloud agent, data will be passed and managed by 
Testkube servers (Coming Soon!)

## Installing Agent

Please follow [install steps](0-testkube-cloud/installing-agent.md) for getting ready to use Testkube Agent