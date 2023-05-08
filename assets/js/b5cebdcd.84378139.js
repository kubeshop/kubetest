"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[8121],{3905:(e,t,n)=>{n.d(t,{Zo:()=>l,kt:()=>d});var s=n(67294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);t&&(s=s.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,s)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,s,r=function(e,t){if(null==e)return{};var n,s,r={},a=Object.keys(e);for(s=0;s<a.length;s++)n=a[s],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(s=0;s<a.length;s++)n=a[s],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var u=s.createContext({}),c=function(e){var t=s.useContext(u),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},l=function(e){var t=c(e.components);return s.createElement(u.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return s.createElement(s.Fragment,{},t)}},m=s.forwardRef((function(e,t){var n=e.components,r=e.mdxType,a=e.originalType,u=e.parentName,l=i(e,["components","mdxType","originalType","parentName"]),m=c(n),d=r,k=m["".concat(u,".").concat(d)]||m[d]||p[d]||a;return n?s.createElement(k,o(o({ref:t},l),{},{components:n})):s.createElement(k,o({ref:t},l))}));function d(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var a=n.length,o=new Array(a);o[0]=m;var i={};for(var u in t)hasOwnProperty.call(t,u)&&(i[u]=t[u]);i.originalType=e,i.mdxType="string"==typeof e?e:r,o[1]=i;for(var c=2;c<a;c++)o[c]=n[c];return s.createElement.apply(null,o)}return s.createElement.apply(null,n)}m.displayName="MDXCreateElement"},72781:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>u,contentTitle:()=>o,default:()=>p,frontMatter:()=>a,metadata:()=>i,toc:()=>c});var s=n(87462),r=(n(67294),n(3905));const a={},o="Testkube Custom Resources",i={unversionedId:"articles/crds",id:"articles/crds",title:"Testkube Custom Resources",description:"In Testkube, Tests, Test Suites, Executors and Webhooks are defined using Custom Resources. The current definitions can be found in the kubeshop/testkube-operator repository.",source:"@site/docs/articles/crds.md",sourceDirName:"articles",slug:"/articles/crds",permalink:"/articles/crds",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/articles/crds.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Development",permalink:"/articles/development"}},u={},c=[{value:"Tests",id:"tests",level:2},{value:"Test Suites",id:"test-suites",level:2},{value:"Executors",id:"executors",level:2},{value:"Webhooks",id:"webhooks",level:2}],l={toc:c};function p(e){let{components:t,...n}=e;return(0,r.kt)("wrapper",(0,s.Z)({},l,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"testkube-custom-resources"},"Testkube Custom Resources"),(0,r.kt)("p",null,"In Testkube, Tests, Test Suites, Executors and Webhooks are defined using ",(0,r.kt)("a",{parentName:"p",href:"https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/"},"Custom Resources"),". The current definitions can be found in the ",(0,r.kt)("a",{parentName:"p",href:"https://github.com/kubeshop/testkube-operator/tree/main/config/crd"},"kubeshop/testkube-operator")," repository."),(0,r.kt)("p",null,"You can always check the list of all CRDs using ",(0,r.kt)("inlineCode",{parentName:"p"},"kubectl")," configured to point to your Kubernetes cluster with Testkube installed:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-sh"},"$ kubectl get crds\n")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-sh",metastring:'title="Expected output:"',title:'"Expected','output:"':!0},"NAME                                  CREATED AT\ncertificaterequests.cert-manager.io   2022-04-01T10:53:54Z\ncertificates.cert-manager.io          2022-04-01T10:53:54Z\nchallenges.acme.cert-manager.io       2022-04-01T10:53:54Z\nclusterissuers.cert-manager.io        2022-04-01T10:53:54Z\nexecutors.executor.testkube.io        2022-04-13T11:44:22Z\nissuers.cert-manager.io               2022-04-01T10:53:54Z\norders.acme.cert-manager.io           2022-04-01T10:53:54Z\nscripts.tests.testkube.io             2022-04-13T11:44:22Z\ntests.tests.testkube.io               2022-04-13T11:44:22Z\ntestsuites.tests.testkube.io          2022-04-13T11:44:22Z\nwebhooks.executor.testkube.io         2022-04-13T11:44:22Z\n")),(0,r.kt)("p",null,"To check details on one of the CRDs, use ",(0,r.kt)("inlineCode",{parentName:"p"},"describe"),":"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-sh"},"$ kubectl describe crd tests.tests.testkube.io\n")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-sh",metastring:'title="Expected output:"',title:'"Expected','output:"':!0},"Name:         tests.tests.testkube.io\nNamespace:\nLabels:       app.kubernetes.io/managed-by=Helm\nAnnotations:  controller-gen.kubebuilder.io/version: v0.4.1\n              meta.helm.sh/release-name: testkube\n              meta.helm.sh/release-namespace: testkube\nAPI Version:  apiextensions.k8s.io/v1\nKind:         CustomResourceDefinition\n...\n")),(0,r.kt)("p",null,"Below, you will find short descriptions and example declarations of the custom resources defined by Testkube."),(0,r.kt)("h2",{id:"tests"},"Tests"),(0,r.kt)("p",null,"Testkube Tests can be defined as a single executable unit of tests. Depending on the test type, this can mean one or multiple test files."),(0,r.kt)("p",null,"To get all the test types available in your cluster, check the executors:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-sh"},"testkube get executors -o yaml | grep -A1 types\n")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-sh",metastring:'title="Expected output:"',title:'"Expected','output:"':!0},"    types:\n    - postman/collection\n--\n    types:\n    - curl/test\n--\n    types:\n    - cypress/project\n--\n    types:\n    - k6/script\n--\n    types:\n    - postman/collection\n--\n    types:\n    - soapui/xml\n")),(0,r.kt)("p",null,"When creating a Testkube Test, there are multiple supported input types:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"String"),(0,r.kt)("li",{parentName:"ul"},"Git directory"),(0,r.kt)("li",{parentName:"ul"},"Git file"),(0,r.kt)("li",{parentName:"ul"},"File URI")),(0,r.kt)("p",null,"Variables can be configured using the ",(0,r.kt)("inlineCode",{parentName:"p"},"variables")," field as shown below."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'apiVersion: tests.testkube.io/v3\nkind: Test\nmetadata:\n  name: example-test\n  namespace: testkube\nspec:\n  content:\n    data: "{...}"\n    type: string\n  type: postman/collection\n  executionRequest:\n    variables:\n      var1:\n        name: var1\n        type: basic\n        value: val1\n      sec1:\n        name: sec1\n        type: secret\n        valueFrom:\n          secretKeyRef:\n            key: sec1\n            name: vartest4-testvars\n')),(0,r.kt)("h2",{id:"test-suites"},"Test Suites"),(0,r.kt)("p",null,"Testkube Test Suites are collections of Testkube Tests of the same or different types."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: tests.testkube.io/v2\nkind: TestSuite\nmetadata:\n  name: example-testsuite\n  namespace: testkube\nspec:\n  description: Example Test Suite\n  steps:\n    - execute:\n        name: example-test1\n        namespace: testkube\n    - delay:\n        duration: 1000\n    - execute:\n        name: example-test2\n        namespace: testkube\n")),(0,r.kt)("h2",{id:"executors"},"Executors"),(0,r.kt)("p",null,"Executors are Testkube-specific test runners. There are predefined Executors avialable in Testkube. You can also write your own custom Testkube Executor using ",(0,r.kt)("a",{parentName:"p",href:"https://kubeshop.github.io/testkube/executor-custom/"},"this guide"),"."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml",metastring:'title="Example:"',title:'"Example:"'},"apiVersion: executor.testkube.io/v1\nkind: Executor\nmetadata:\n  name: example-executor\n  namespace: testkube\nspec:\n  executor_type: job\n  image: YOUR_USER/testkube-executor-example:1.0.0\n  types:\n    - example/test\n  content_types:\n    - string\n    - file-uri\n    - git\n  features:\n    - artifacts\n    - junit-report\n  meta:\n    iconURI: http://mydomain.com/icon.jpg\n    docsURI: http://mydomain.com/docs\n    tooltips:\n      name: please enter executor name\n")),(0,r.kt)("h2",{id:"webhooks"},"Webhooks"),(0,r.kt)("p",null,"Testkube Webhooks are HTTP POST calls having the Testkube Execution object and its current state as payload. They are sent when a test is either started or finished. This can be defined under ",(0,r.kt)("inlineCode",{parentName:"p"},"events"),"."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: executor.testkube.io/v1\nkind: Webhook\nmetadata:\n  name: example-webhook\n  namespace: testkube\nspec:\n  uri: http://localhost:8080/events\n  events:\n    - start-test\n    - end-test\n")))}p.isMDXComponent=!0}}]);