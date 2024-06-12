"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[6212],{3905:(e,t,n)=>{n.d(t,{Zo:()=>o,kt:()=>k});var a=n(67294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},l=Object.keys(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(a=0;a<l.length;a++)n=l[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var u=a.createContext({}),p=function(e){var t=a.useContext(u),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},o=function(e){var t=p(e.components);return a.createElement(u.Provider,{value:t},e.children)},m={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},d=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,l=e.originalType,u=e.parentName,o=i(e,["components","mdxType","originalType","parentName"]),d=p(n),k=r,c=d["".concat(u,".").concat(k)]||d[k]||m[k]||l;return n?a.createElement(c,s(s({ref:t},o),{},{components:n})):a.createElement(c,s({ref:t},o))}));function k(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var l=n.length,s=new Array(l);s[0]=d;var i={};for(var u in t)hasOwnProperty.call(t,u)&&(i[u]=t[u]);i.originalType=e,i.mdxType="string"==typeof e?e:r,s[1]=i;for(var p=2;p<l;p++)s[p]=n[p];return a.createElement.apply(null,s)}return a.createElement.apply(null,n)}d.displayName="MDXCreateElement"},34127:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>u,contentTitle:()=>s,default:()=>m,frontMatter:()=>l,metadata:()=>i,toc:()=>p});var a=n(87462),r=(n(67294),n(3905));const l={},s="Creating Test Suites",i={unversionedId:"articles/creating-test-suites",id:"articles/creating-test-suites",title:"Creating Test Suites",description:"A large IT department has a frontend team and a backend team, everything is",source:"@site/docs/articles/creating-test-suites.md",sourceDirName:"articles",slug:"/articles/creating-test-suites",permalink:"/articles/creating-test-suites",draft:!1,editUrl:"https://github.com/kubeshop/testkube/tree/develop/docs/docs/articles/creating-test-suites.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Adding Timeouts",permalink:"/articles/adding-timeout"},next:{title:"Running Test Suites",permalink:"/articles/running-test-suites"}},u={},p=[{value:"Passing Test Suite Artifacts between Steps",id:"passing-test-suite-artifacts-between-steps",level:2},{value:"Test Suite Creation",id:"test-suite-creation",level:2},{value:"Test Suite Steps",id:"test-suite-steps",level:2},{value:"Usage Example",id:"usage-example",level:3}],o={toc:p};function m(e){let{components:t,...n}=e;return(0,r.kt)("wrapper",(0,a.Z)({},o,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"creating-test-suites"},"Creating Test Suites"),(0,r.kt)("p",null,"A large IT department has a frontend team and a backend team, everything is\ndeployed on Kubernetes clusters, and each team is responsible for its part of the work. The frontend engineers test their code using the Cypress testing framework, but the backend engineers prefer simpler tools like Postman. They have many Postman collections defined and want to run them against a Kubernetes cluster but some of their services are not exposed externally."),(0,r.kt)("p",null,"A QA leader is responsible for release trains and wants to be sure that before the release all tests are completed successfully. The QA leader will need to create pipelines that orchestrate each teams' tests into a common platform."),(0,r.kt)("p",null,"This is easily done with Testkube. Each team can run their tests against clusters on their own, and the QA manager can create test resources and add tests written by all teams."),(0,r.kt)("p",null,(0,r.kt)("inlineCode",{parentName:"p"},"Test Suites")," stands for the orchestration of different test steps, which can run sequentially and/or in parallel.\nOn each batch step you can define either one or multiple steps such as test execution, delay, or other (future) steps.\nBy default the concurrency level for parallel tests is set to 10, you can redefine it using ",(0,r.kt)("inlineCode",{parentName:"p"},"--concurrency")," option for CLI command."),(0,r.kt)("h2",{id:"passing-test-suite-artifacts-between-steps"},"Passing Test Suite Artifacts between Steps"),(0,r.kt)("p",null,"In some scenarios you need to access artifacts generated on previous steps of the test suite. Testkube provides two options to define which artifacts to download in the init container: all previous step artifacts or artifacts for selected steps (step number is started from 1) or artifacts for latest executions of previously executed tests (identified by names). All downloaded artifacts are stored in /data/downloaded-artifacts/{execution id} folder. See a few examples below."),(0,r.kt)("h2",{id:"test-suite-creation"},"Test Suite Creation"),(0,r.kt)("p",null,"Creating tests is really simple - create the test definition in a JSON file and pass it to the ",(0,r.kt)("inlineCode",{parentName:"p"},"testkube")," ",(0,r.kt)("inlineCode",{parentName:"p"},"kubectl")," plugin."),(0,r.kt)("p",null,"An example test file could look like this:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-sh"},'echo \'\n{\n    "name": "testkube-suite",\n    "description": "Testkube test suite, api, dashboard and performance",\n    "steps": [\n        {"execute": [{"test": "testkube-api"}, {""test": "testkube-dashboard"}]},\n        {"execute": [{"delay": "1s"}]},\n        {"downloadArtifacts": {"previousTestNames": ["testkube-api"]}, "execute": [{"test": "testkube-dashboard"}, {"delay": "1s"}, {""test": "testkube-homepage"}]},\n        {"execute": [{"delay": "1s"}]},\n        {"downloadArtifacts": {"previousStepNumbers": [1, 3]}, "execute": [{"test": "testkube-api-performance"}]},\n        {"execute": [{"delay": "1s"}]},\n        {"downloadArtifacts": {"allPreviousSteps": true}, "execute": [{"test": "testkube-homepage-performance"}]}\n    ]\n}\' | kubectl testkube create testsuite\n')),(0,r.kt)("p",null,"To check if the test was created correctly, you can look at ",(0,r.kt)("inlineCode",{parentName:"p"},"TestSuite")," Custom Resource in your Kubernetes cluster:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-sh"},"kubectl get testsuites -ntestkube\n")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-sh",metastring:'title="Expected output:"',title:'"Expected','output:"':!0},"NAME                  AGE\ntestkube-suite           1m\ntestsuite-example-2   2d21h\n")),(0,r.kt)("p",null,"To get the details of a test:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-sh"},"kubectl get testsuites -ntestkube testkube-suite -oyaml\n")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml",metastring:'title="Expected output:"',title:'"Expected','output:"':!0},'apiVersion: tests.testkube.io/v3\nkind: TestSuite\nmetadata:\n  creationTimestamp: "2022-01-11T07:46:12Z"\n  generation: 4\n  name: testkube-suite\n  namespace: testkube\n  resourceVersion: "57695094"\n  uid: ea90a79e-bb46-49ee-a3ef-a5d99cee0a2c\nspec:\n  description: "Testkube test suite, api, dashboard and performance"\n  steps:\n  - stopOnFailure: false\n    execute:\n    - test: testkube-api\n    - test: testkube-dashboard\n  - stopOnFailure: false\n    execute:\n    - delay: 1s\n  - stopOnFailure: false\n    downloadArtifacts:\n      allPreviousSteps: false\n      previousTestNames:\n      - testkube-api\n    execute:\n    - test: testkube-dashboard\n    - delay: 1s\n    - test: testkube-homepage\n  - stopOnFailure: false\n    execute:\n    - delay: 1s\n  - stopOnFailure: false\n    downloadArtifacts:\n      allPreviousSteps: false\n      previousStepNumbers:\n      - 1\n      - 3\n    execute:\n    - test: testkube-api-performance\n  - stopOnFailure: false\n    execute:\n    - delay: 1s\n  - stopOnFailure: false\n    downloadArtifacts:\n      allPreviousSteps: true\n    execute:\n    - test: testkube-homepage-performance\n')),(0,r.kt)("p",null,"Your ",(0,r.kt)("inlineCode",{parentName:"p"},"Test Suite")," is defined and you can start running testing workflows."),(0,r.kt)("h2",{id:"test-suite-steps"},"Test Suite Steps"),(0,r.kt)("p",null,"Test Suite Steps are the individual components or actions that make up a Test Suite. They are typically a sequence of tests that are run in a specific order. There are two types of Test Suite Steps:"),(0,r.kt)("p",null,"Tests: These are the actual tests to be run. They could be unit tests, integration tests, functional tests, etc., depending on the context."),(0,r.kt)("p",null,"Delays: These are time delays inserted between tests. They are used to wait for a certain period of time before proceeding to the next test. This can be useful in situations where you need to wait for some process to complete or some condition to be met before proceeding."),(0,r.kt)("p",null,"Similar to running a Test, running a Test Suite Step based on a test allows for specific execution request parameters to be overwritten. Step level parameters overwrite Test Suite level parameters, which in turn overwrite Test level parameters. The Step level parameters are configurable only via CRDs at the moment."),(0,r.kt)("p",null,"For details on which parameters are available in the CRDs, please consult the table below:"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Parameter"),(0,r.kt)("th",{parentName:"tr",align:null},"Test"),(0,r.kt)("th",{parentName:"tr",align:null},"Test Suite"),(0,r.kt)("th",{parentName:"tr",align:null},"Test Step"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"name"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testSuiteName"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"number"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"executionLabels"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"namespace"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"variablesFile"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"isVariablesFileUploaded"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"variables"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testSecretUUID"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testSuiteSecretUUID"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"args"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"argsMode"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"command"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"image"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"imagePullSecrets"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"sync"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"httpProxy"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"httpsProxy"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"negativeTest"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"activeDeadlineSeconds"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"artifactRequest"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"jobTemplate"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"jobTemplateReference"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"cronJobTemplate"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"cronJobTemplateReference"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"preRunScript"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"postRunScript"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"executePostRunScriptBeforeScraping"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"sourceScripts"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"scraperTemplate"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"scraperTemplateReference"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"pvcTemplate"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"pvcTemplateReference"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"envConfigMaps"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"envSecrets"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"runningContext"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"slavePodRequest"),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"secretUUID"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"labels"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null})),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"timeout"),(0,r.kt)("td",{parentName:"tr",align:null}),(0,r.kt)("td",{parentName:"tr",align:null},"\u2713"),(0,r.kt)("td",{parentName:"tr",align:null})))),(0,r.kt)("p",null,"Similar to Tests and Test Suites, Test Suite Steps can also have a field of type ",(0,r.kt)("inlineCode",{parentName:"p"},"executionRequest")," like in the example below:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'apiVersion: tests.testkube.io/v3\nkind: TestSuite\nmetadata:\n  name: jmeter-special-cases\n  namespace: testkube\n  labels:\n    core-tests: special-cases\nspec:\n  description: "jmeter and jmeterd executor - special-cases"\n  steps:\n  - stopOnFailure: false\n    execute:\n    - test: jmeterd-executor-smoke-custom-envs-replication\n      executionRequest:\n        args: ["-d", "-s"]\n      ...\n  - stopOnFailure: false\n    execute:\n    - test: jmeterd-executor-smoke-env-value-in-args\n')),(0,r.kt)("p",null,"The ",(0,r.kt)("inlineCode",{parentName:"p"},"Definition")," section of each Test Suite in the Testkube UI offers the opportunity to directly edit the Test Suite CRDs. Besides that, consider also using ",(0,r.kt)("inlineCode",{parentName:"p"},"kubectl edit testsuite/jmeter-special-cases -n testkube")," on the command line."),(0,r.kt)("h3",{id:"usage-example"},"Usage Example"),(0,r.kt)("p",null,"An example of use case for test suite step parameters would be running the same K6 load test with different arguments and memory and CPU requirements."),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"Create and Configure the Test")),(0,r.kt)("p",null,"Let's say our test CRD stored in the file ",(0,r.kt)("inlineCode",{parentName:"p"},"k6-test.yaml")," looks the following:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'apiVersion: tests.testkube.io/v3\nkind: Test\nmetadata:\n  name: k6-test-parallel\n  labels:\n    core-tests: executors\n  namespace: testkube\nspec:\n  type: k6/script\n  content:\n    type: git\n    repository:\n      type: git\n      uri: https://github.com/kubeshop/testkube.git\n      branch: main\n      path: test/k6/executor-tests/\n  executionRequest:\n      args:\n        - k6-smoke-test-without-envs.js\n      jobTemplate: "apiVersion: batch/v1\\nkind: Job\\nspec:\\n  template:\\n    spec:\\n      containers:\\n        - name: \\"{{ .Name }}\\"\\n          image: {{ .Image }}\\n          resources:\\n            requests:\\n              memory: 128Mi\\n              cpu: 128m\\n"\n      activeDeadlineSeconds: 180\n')),(0,r.kt)("p",null,"We can apply this from the command line using:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl apply -f k6-test.yaml\n")),(0,r.kt)("ol",{start:2},(0,r.kt)("li",{parentName:"ol"},"Run the Test")),(0,r.kt)("p",null,"To run this test, execute:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"testkube run test k6-test-parallel\n")),(0,r.kt)("p",null,"A new Testkube execution will be created. If you investigate the new job assigned to this execution, you will see the memory and cpu limit specified in the job template was set. Checking the arguments from the ",(0,r.kt)("inlineCode",{parentName:"p"},"executionRequest")," is also possible with:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube get execution k6-test-parallel-1\n")),(0,r.kt)("ol",{start:3},(0,r.kt)("li",{parentName:"ol"},"Create and Configure the Test Suite")),(0,r.kt)("p",null,"We are content with the test created, but we need to make sure our application works with different kinds of loads. We could create a new Test with different parameters, but that would come with the overhead of having to manage and sync two instances of the same test. Creating a test suite makes test orchestration a more robust operation."),(0,r.kt)("p",null,"We have the following ",(0,r.kt)("inlineCode",{parentName:"p"},"k6-test-suite.yaml")," file:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'apiVersion: tests.testkube.io/v3\nkind: TestSuite\nmetadata:\n  name: k6-parallel\n  namespace: testkube\nspec:\n  description: "k6 parallel testsuite"\n  steps:\n  - stopOnFailure: false\n    execute:\n    - test: k6-test-parallel\n      executionRequest:\n        argsMode: override\n        args:\n          - -vu\n          - "1"\n          - k6-smoke-test-without-envs.js\n        jobTemplate: "apiVersion: batch/v1\\nkind: Job\\nspec:\\n  template:\\n    spec:\\n      containers:\\n        - name: \\"{{ .Name }}\\"\\n          image: {{ .Image }}\\n          resources:\\n            requests:\\n              memory: 64Mi\\n              cpu: 128m\\n"\n    - test: k6-test-parallel\n      executionRequest:\n        argsMode: override\n        args:\n          - -vu\n          - "2"\n          - k6-smoke-test-without-envs.js\n')),(0,r.kt)("p",null,"Note that there are two steps in there running the same test. The difference is in their ",(0,r.kt)("inlineCode",{parentName:"p"},"executionRequest"),". The first step is setting the number of virtual users to one and updating the jobTemplate to use a different memory requirement. The second test updates the VUs to 2."),(0,r.kt)("p",null,"Create the test suite with the command:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl apply -f k6-test-suite.yaml\n")),(0,r.kt)("ol",{start:4},(0,r.kt)("li",{parentName:"ol"},"Run the Test Suite")),(0,r.kt)("p",null,"Run the test suite with:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube run testsuite k6-parallel\n")),(0,r.kt)("p",null,"The output of both of the test runs can be examined with:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"testkube get execution k6-parallel-k6-test-parallel-2\n\ntestkube get execution k6-parallel-k6-test-parallel-3\n")),(0,r.kt)("p",null,"The logs show the exact commands:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"...\n\ud83d\udd2c Executing in directory /data/repo:\n $ k6 run test/k6/executor-tests/k6-smoke-test-without-envs.js -vu 1\n...\n\ud83d\udd2c Executing in directory /data/repo:\n $ k6 run test/k6/executor-tests/k6-smoke-test-without-envs.js -vu 2\n...\n")),(0,r.kt)("p",null,"The job template configuration will be visible on the job level, running ",(0,r.kt)("inlineCode",{parentName:"p"},"kubectl get jobs -n testkube")," and ",(0,r.kt)("inlineCode",{parentName:"p"},"kubectl get job ${job_id} -o yaml -n testkube")," should be enough to check the settings."),(0,r.kt)("p",null,"Now we know how to increase the flexibility, reusability and scalability of your tests using test suites. By setting parameters on test suite step levels, we are making our testing automation more robust and easier to manage."))}m.isMDXComponent=!0}}]);