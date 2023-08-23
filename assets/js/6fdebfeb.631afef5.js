"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[1966],{3905:(e,t,r)=>{r.d(t,{Zo:()=>p,kt:()=>g});var n=r(67294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function s(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?s(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):s(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function i(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},s=Object.keys(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var l=n.createContext({}),c=function(e){var t=n.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},p=function(e){var t=c(e.components);return n.createElement(l.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,s=e.originalType,l=e.parentName,p=i(e,["components","mdxType","originalType","parentName"]),d=c(r),g=a,m=d["".concat(l,".").concat(g)]||d[g]||u[g]||s;return r?n.createElement(m,o(o({ref:t},p),{},{components:r})):n.createElement(m,o({ref:t},p))}));function g(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var s=r.length,o=new Array(s);o[0]=d;var i={};for(var l in t)hasOwnProperty.call(t,l)&&(i[l]=t[l]);i.originalType=e,i.mdxType="string"==typeof e?e:a,o[1]=i;for(var c=2;c<s;c++)o[c]=r[c];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},61068:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>l,contentTitle:()=>o,default:()=>u,frontMatter:()=>s,metadata:()=>i,toc:()=>c});var n=r(87462),a=(r(67294),r(3905));const s={},o="Triggers",i={unversionedId:"articles/test-triggers",id:"articles/test-triggers",title:"Triggers",description:"Testkube allows you to automate running tests and test suites by defining triggers on certain events for various Kubernetes resources.",source:"@site/docs/articles/test-triggers.md",sourceDirName:"articles",slug:"/articles/test-triggers",permalink:"/articles/test-triggers",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/articles/test-triggers.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Scheduling Tests",permalink:"/articles/scheduling-tests"},next:{title:"Webhooks",permalink:"/articles/webhooks"}},l={},c=[{value:"What is a Testkube Test Trigger?",id:"what-is-a-testkube-test-trigger",level:2},{value:"Custom Resource Definition Model",id:"custom-resource-definition-model",level:2},{value:"Selectors",id:"selectors",level:3},{value:"Name Selector",id:"name-selector",level:4},{value:"Label Selector",id:"label-selector",level:4},{value:"Resource Conditions",id:"resource-conditions",level:3},{value:"Resource Probes",id:"resource-probes",level:3},{value:"Supported Values",id:"supported-values",level:3},{value:"Example",id:"example",level:2},{value:"Architecture",id:"architecture",level:2},{value:"API",id:"api",level:2},{value:"Creating Test Triggers in the Testkube Dashboard",id:"creating-test-triggers-in-the-testkube-dashboard",level:2},{value:"Video Tutorial",id:"video-tutorial",level:2}],p={toc:c};function u(e){let{components:t,...s}=e;return(0,a.kt)("wrapper",(0,n.Z)({},p,s,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"triggers"},"Triggers"),(0,a.kt)("p",null,"Testkube allows you to automate running tests and test suites by defining triggers on certain events for various Kubernetes resources."),(0,a.kt)("h2",{id:"what-is-a-testkube-test-trigger"},"What is a Testkube Test Trigger?"),(0,a.kt)("p",null,"In generic terms, a ",(0,a.kt)("em",{parentName:"p"},"Trigger")," defines an ",(0,a.kt)("em",{parentName:"p"},"action")," which will be executed for a given ",(0,a.kt)("em",{parentName:"p"},"execution")," when a certain ",(0,a.kt)("em",{parentName:"p"},"event")," on a specific ",(0,a.kt)("em",{parentName:"p"},"resource")," occurs. For example, we could define a ",(0,a.kt)("em",{parentName:"p"},"TestTrigger")," which ",(0,a.kt)("em",{parentName:"p"},"runs")," a ",(0,a.kt)("em",{parentName:"p"},"Test")," when a ",(0,a.kt)("em",{parentName:"p"},"ConfigMap")," gets ",(0,a.kt)("em",{parentName:"p"},"modified"),"."),(0,a.kt)("p",null,"Watch our ",(0,a.kt)("a",{parentName:"p",href:"#video-tutorial"},"video guide")," on using Testkube Test Triggers to perform ",(0,a.kt)("strong",{parentName:"p"},"Asynchronous Testing in Kubernetes"),"."),(0,a.kt)("h2",{id:"custom-resource-definition-model"},"Custom Resource Definition Model"),(0,a.kt)("h3",{id:"selectors"},"Selectors"),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"resourceSelector")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"testSelector")," fields support selecting resources either by name or using\nthe Kubernetes ",(0,a.kt)("a",{parentName:"p",href:"https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#resources-that-support-set-based-requirements"},"Label Selector"),"."),(0,a.kt)("p",null,"Each selector should specify the ",(0,a.kt)("inlineCode",{parentName:"p"},"namespace")," of the object, otherwise the namespace defaults to ",(0,a.kt)("inlineCode",{parentName:"p"},"testkube"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"selector := resourceSelector | testSelector\n")),(0,a.kt)("h4",{id:"name-selector"},"Name Selector"),(0,a.kt)("p",null,"Name selectors are used when we want to select a specific resource in a specific namespace."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},"selector:\n  name: Kubernetes object name\n  namespace: Kubernetes object namespace (default is **testkube**)\n")),(0,a.kt)("h4",{id:"label-selector"},"Label Selector"),(0,a.kt)("p",null,"Label selectors are used when we want to select a group of resources in a specific namespace."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},"selector:\n  namespace: Kubernetes object namespace (default is **testkube**)\n  labelSelector:\n    matchLabels: map of key-value pairs\n    matchExpressions:\n      - key: label name\n        operator: [In | NotIn | Exists | DoesNotExist\n        values: list of values\n")),(0,a.kt)("h3",{id:"resource-conditions"},"Resource Conditions"),(0,a.kt)("p",null,"Resource Conditions allows triggers to be defined based on the status conditions for a specific resource."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},"conditionSpec:\n    timeout: duration in seconds the test trigger waits for conditions, until its stopped\n    delay: duration in seconds the test trigger waits between condition checks      \n    conditions:\n    - type: test trigger condition type\n      status: test trigger condition status, supported values - True, False, Unknown\n      reason: test trigger condition reason\n      ttl: test trigger condition ttl\n")),(0,a.kt)("h3",{id:"resource-probes"},"Resource Probes"),(0,a.kt)("p",null,"Resource Probes allows triggers to be defined based on the probe status."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},"probeSpec:\n    timeout: duration in seconds the test trigger waits for probes, until its stopped\n    delay: duration in seconds the test trigger waits between probes\n    probes:\n    - scheme: test trigger condition probe scheme to connect to host, default is http\n      host: test trigger condition probe host, default is pod ip or service name\n      path: test trigger condition probe path to check, default is /\n      port: test trigger condition probe port to connect\n      headers: test trigger condition probe headers to submit\n")),(0,a.kt)("h3",{id:"supported-values"},"Supported Values"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("strong",{parentName:"li"},"Resource"),"  - pod, deployment, statefulset, daemonset, service, ingress, event, configmap"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("strong",{parentName:"li"},"Action"),"    - run"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("strong",{parentName:"li"},"Event"),"     - created, modified, deleted"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("strong",{parentName:"li"},"Execution")," - test, testsuite")),(0,a.kt)("p",null,(0,a.kt)("strong",{parentName:"p"},"NOTE"),": All resources support the above-mentioned events, a list of finer-grained events is in the works, stay tuned..."),(0,a.kt)("h2",{id:"example"},"Example"),(0,a.kt)("p",null,"Here is an example for a ",(0,a.kt)("strong",{parentName:"p"},"Test Trigger")," ",(0,a.kt)("em",{parentName:"p"},"default/testtrigger-example")," which runs the ",(0,a.kt)("strong",{parentName:"p"},"TestSuite")," ",(0,a.kt)("em",{parentName:"p"},"frontend/sanity-test"),"\nwhen a ",(0,a.kt)("strong",{parentName:"p"},"deployment")," containing the label ",(0,a.kt)("strong",{parentName:"p"},"testkube.io/tier: backend")," gets ",(0,a.kt)("strong",{parentName:"p"},"modified")," and also has the conditions ",(0,a.kt)("strong",{parentName:"p"},"Progressing: True: NewReplicaSetAvailable")," and ",(0,a.kt)("strong",{parentName:"p"},"Available: True"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},'apiVersion: tests.testkube.io/v1\nkind: TestTrigger\nmetadata:\n  name: testtrigger-example\n  namespace: default\nspec:\n  resource: deployment\n  resourceSelector:\n    labelSelector:\n      matchLabels:\n        testkube.io/tier: backend\n  event: modified\n  conditionSpec:\n    timeout: 100\n    delay: 2\n    conditions:\n    - type: Progressing\n      status: "True"\n      reason: "NewReplicaSetAvailable"\n      ttl: 60\n    - type: Available\n      status: "True"\n  probeSpec:\n    timeout: 50\n    delay: 1\n    probes:\n    - scheme: http\n      host: testkube-api-server\n      path: /health\n      port: 8088\n      headers:\n        X-Token: "12345"\n    - host: testkube-dashboard\n      port: 8080     \n  action: run\n  execution: testsuite\n  testSelector:\n    name: sanity-test\n    namespace: frontend\n')),(0,a.kt)("h2",{id:"architecture"},"Architecture"),(0,a.kt)("p",null,"Testkube uses ",(0,a.kt)("a",{parentName:"p",href:"https://pkg.go.dev/k8s.io/client-go/informers"},"Informers")," to watch Kubernetes resources and register handlers\non certain actions on the watched Kubernetes resources."),(0,a.kt)("p",null,"Informers are a reliable, scalable and fault-tolerant Kubernetes concept where each informer registers handlers with the\nKubernetes API and gets notified by Kubernetes on each event on the watched resources."),(0,a.kt)("h2",{id:"api"},"API"),(0,a.kt)("p",null,"Testkube exposes CRUD operations on test triggers in the REST API. Check out the ",(0,a.kt)("a",{parentName:"p",href:"/openapi"},"Open API")," docs for more info."),(0,a.kt)("h2",{id:"creating-test-triggers-in-the-testkube-dashboard"},"Creating Test Triggers in the Testkube Dashboard"),(0,a.kt)("p",null,"Click on the lightening bolt icon on the left of the Testkube IDE to open the dialog to create test triggers. Any current test triggers will be listed and the ",(0,a.kt)("inlineCode",{parentName:"p"},"Create a new trigger")," button is at the top right of the screen."),(0,a.kt)("p",null,(0,a.kt)("img",{alt:"Trigger Screen",src:r(66612).Z,width:"2874",height:"1540"})),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"Create new trigger")," dialog opens:"),(0,a.kt)("p",null,(0,a.kt)("img",{alt:"Create Trigger",src:r(35746).Z,width:"2808",height:"1266"})),(0,a.kt)("p",null,"Input the condition that will cause the trigger and click ",(0,a.kt)("inlineCode",{parentName:"p"},"Next"),"."),(0,a.kt)("p",null,"Input the action that will be the result of the trigger condition happening and click ",(0,a.kt)("inlineCode",{parentName:"p"},"Create"),"."),(0,a.kt)("p",null,(0,a.kt)("img",{alt:"Create Trigger Action",src:r(2343).Z,width:"2874",height:"1564"})),(0,a.kt)("p",null,"Get all the details in the video below!"),(0,a.kt)("h2",{id:"video-tutorial"},"Video Tutorial"),(0,a.kt)("iframe",{width:"100%",height:"350px",src:"https://www.youtube.com/embed/t4V6E9rQ5W4",title:"YouTube video player",frameborder:"0",allow:"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share",allowfullscreen:!0}))}u.isMDXComponent=!0},2343:(e,t,r)=>{r.d(t,{Z:()=>n});const n=r.p+"assets/images/create-trigger-action-66f0a0cff7093f187dbe3db492dcc723.png"},35746:(e,t,r)=>{r.d(t,{Z:()=>n});const n=r.p+"assets/images/create-trigger-fee6f31a95ace01607673804370bcf86.png"},66612:(e,t,r)=>{r.d(t,{Z:()=>n});const n=r.p+"assets/images/trigger-screen-2bba2c4224d296c8175ef60f0980dbc3.png"}}]);