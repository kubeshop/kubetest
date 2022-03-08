# Development

## **Running with CRDs on Kubernetes Cluster**

The minimial component which must be deployed on your local Kubernetes cluster is testkube-operator with project CRDs (<https://github.com/kubeshop/testkube-operator>)

To install CRDs into your local cluster, checkout the testkube-operator project and run:

```sh
make install 
```

## **Running on a Local Machine**

The next critical components are the Testkube API (<https://github.com/kubeshop/testkube>) and an executor. You can use your own tests executor or an existing one from Testkube.

Checkout the Testkube project and run a local API server:

```sh
make run-mongo-dev run-api
```

Next, checkout and run the Testkube Postman executor (<https://github.com/kubeshop/testkube-executor-postman>). The postman executor is MongoDB based so it will launch MongoDB with the API server step:

```sh
make run-executor
```

### **Installing Local Executors**

Install development executors by running them from the Testkube project (<https://github.com/kubeshop/testkube>):

```sh
make dev-install-local-executors
```

This will register Custom Resources for the following test types:

- local-postman/collection
- local-cypress/project
- local-curl/test

Create a `Test` Custom Resource with one of the types above to be executed on given the executor:

```sh
kubectl testkube tests create --file my_collection_file.json --name my-test-name --type local-postman/collection
```

To summarize: `type` is the single relation between `Test` and `Executor`.

## **Intercepting an API Server on a Cluster**

For debugging on Kubernetes, intercept the whole API Server (or Postman executor) service
by using [Telepresence](https://telepresence.io).

Simply intercept the API server with the local instance.

To create/run tests pointed to in-cluster executors, start the API Server with telepresence mode:

```sh
make run-api-telepresence
```

