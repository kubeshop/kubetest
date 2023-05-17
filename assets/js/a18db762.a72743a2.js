"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[312],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>h});var s=n(67294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);t&&(s=s.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,s)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,s,a=function(e,t){if(null==e)return{};var n,s,a={},i=Object.keys(e);for(s=0;s<i.length;s++)n=i[s],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(s=0;s<i.length;s++)n=i[s],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var r=s.createContext({}),c=function(e){var t=s.useContext(r),n=t;return e&&(n="function"==typeof e?e(t):l(l({},t),e)),n},u=function(e){var t=c(e.components);return s.createElement(r.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return s.createElement(s.Fragment,{},t)}},p=s.forwardRef((function(e,t){var n=e.components,a=e.mdxType,i=e.originalType,r=e.parentName,u=o(e,["components","mdxType","originalType","parentName"]),p=c(n),h=a,k=p["".concat(r,".").concat(h)]||p[h]||d[h]||i;return n?s.createElement(k,l(l({ref:t},u),{},{components:n})):s.createElement(k,l({ref:t},u))}));function h(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=n.length,l=new Array(i);l[0]=p;var o={};for(var r in t)hasOwnProperty.call(t,r)&&(o[r]=t[r]);o.originalType=e,o.mdxType="string"==typeof e?e:a,l[1]=o;for(var c=2;c<i;c++)l[c]=n[c];return s.createElement.apply(null,l)}return s.createElement.apply(null,n)}p.displayName="MDXCreateElement"},14524:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>r,contentTitle:()=>l,default:()=>d,frontMatter:()=>i,metadata:()=>o,toc:()=>c});var s=n(87462),a=(n(67294),n(3905));const i={},l="Integrating with Slack",o={unversionedId:"articles/slack-integration",id:"articles/slack-integration",title:"Integrating with Slack",description:"In order to receive notifications in Slack about the status of the testing process, Testkube provides integration with Slack. Below are the steps to configure it.",source:"@site/docs/articles/slack-integration.md",sourceDirName:"articles",slug:"/articles/slack-integration",permalink:"/articles/slack-integration",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/articles/slack-integration.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Webhooks",permalink:"/articles/webhooks"},next:{title:"Test CRD Generation",permalink:"/articles/generate-test-crds"}},r={},c=[{value:"Install the Testkube bot to Your Slack Workspace",id:"install-the-testkube-bot-to-your-slack-workspace",level:2},{value:"Configure Testkube to Use the Slack bot Token",id:"configure-testkube-to-use-the-slack-bot-token",level:2},{value:"Adjust Slack Config File",id:"adjust-slack-config-file",level:2},{value:"Adjust Slack Config Using Helm Values",id:"adjust-slack-config-using-helm-values",level:2},{value:"Configure Message Template",id:"configure-message-template",level:2},{value:"Add the Testkube bot to a Channel",id:"add-the-testkube-bot-to-a-channel",level:2}],u={toc:c};function d(e){let{components:t,...i}=e;return(0,a.kt)("wrapper",(0,s.Z)({},u,i,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"integrating-with-slack"},"Integrating with Slack"),(0,a.kt)("p",null,"In order to receive notifications in Slack about the status of the testing process, Testkube provides integration with Slack. Below are the steps to configure it."),(0,a.kt)("iframe",{width:"100%",height:"315",src:"https://www.youtube.com/embed/iaiiDilAyMY",title:"YouTube video player",frameborder:"0",allow:"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share",allowfullscreen:!0}),(0,a.kt)("h2",{id:"install-the-testkube-bot-to-your-slack-workspace"},"Install the Testkube bot to Your Slack Workspace"),(0,a.kt)("p",null,"Testkube bot:"),(0,a.kt)("a",{href:"https://slack.com/oauth/v2/authorize?client_id=1943550956369.3416932538629&scope=chat:write,chat:write.public,groups:read,channels:read&user_scope="},(0,a.kt)("img",{alt:"Add Testkube bot to your Slack workspace",height:"40",width:"139",src:"https://platform.slack-edge.com/img/add_to_slack.png",srcSet:"https://platform.slack-edge.com/img/add_to_slack.png 1x, https://platform.slack-edge.com/img/add_to_slack@2x.png 2x"})),(0,a.kt)("p",null,(0,a.kt)("img",{alt:"img.gif",src:n(65998).Z,width:"600",height:"339"})),(0,a.kt)("h2",{id:"configure-testkube-to-use-the-slack-bot-token"},"Configure Testkube to Use the Slack bot Token"),(0,a.kt)("p",null,"Populate slackToken in the Helm values file, then install Testkube using Helm install, see ",(0,a.kt)("a",{parentName:"p",href:"../articles/getting-started-overview"},"Manual Testkube Helm Charts Installation")," for more information."),(0,a.kt)("h2",{id:"adjust-slack-config-file"},"Adjust Slack Config File"),(0,a.kt)("p",null,"By default the configuration ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/kubeshop/helm-charts/blob/704c71fa3b8f0138f983ea9a2fa598ecbe3868ae/charts/testkube-api/slack-config.json"},"/charts/testkube-api/slack-config.json")," looks like below; it will send notification for all events and all test or test suite names with any labels.\nIf the channel is left empty, it will send to the first channel that the Testkube bot is member of."),(0,a.kt)("p",null,"It is an array of config objects and can use any config combinations:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-json"},'[\n    {\n      "ChannelID": "",\n      "selector": {},\n      "testName": [],\n      "testSuiteName": [],\n      "events": [\n        "start-test",\n        "end-test-success",\n        "end-test-failed",\n        "end-test-aborted",\n        "end-test-timeout",\n        "start-testsuite",\n        "end-testsuite-success",\n        "end-testsuite-failed",\n        "end-testsuite-aborted",\n        "end-testsuite-timeout"\n      ]\n    }\n  ]\n')),(0,a.kt)("p",null,"For example:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-json"},'[\n    {\n      "ChannelID": "C01234567",\n      "selector": {"label1":"value1"},\n      "testName": ["sanity", "testName2"],\n      "testSuiteName": ["test-suite1", "test-suite2"],\n      "events": [\n        "end-test-failed",\n        "end-test-timeout",\n        "end-testsuite-failed",\n        "end-testsuite-timeout"\n      ]\n    },\n    {\n      "ChannelID": "C07654342",\n      "selector": {"label3":"value4"},\n      "testName": ["integration-test1", "integration-test2"],\n      "testSuiteName": ["integration-test-suite1", "integration-test-suite2"],\n      "events": [\n        "start-test",\n        "end-test-success",\n        "end-test-failed",\n        "end-test-aborted",\n        "end-test-timeout",\n        "start-testsuite",\n        "end-testsuite-success",\n        "end-testsuite-failed",\n        "end-testsuite-aborted",\n        "end-testsuite-timeout"\n      ]\n    },\n]\n')),(0,a.kt)("p",null,"This will send notifications to the channel with the id ",(0,a.kt)("inlineCode",{parentName:"p"},"C01234567")," for the test and test suites with labels ",(0,a.kt)("inlineCode",{parentName:"p"},"label1:value1"),'; tests with the labels "sanity" and "testName2" and test suites with the labels "test-suite1" and "test-suite2"; on events with the labels "end-test-failed", "end-test-timeout", "end-testsuite-failed" and "end-testsuite-timeout"; and to the channel with the id ',(0,a.kt)("inlineCode",{parentName:"p"},"C07654342")," for tests with labels ",(0,a.kt)("inlineCode",{parentName:"p"},"label3:value4"),', tests with the labels "integration-test1" and "integration-test2" and test suites with the labels "integration-test-suite1" and "integration-test-suite2" on all events.'),(0,a.kt)("h2",{id:"adjust-slack-config-using-helm-values"},"Adjust Slack Config Using Helm Values"),(0,a.kt)("p",null,"For convenience, you can also adjust the Slack config using Helm values but you have to use YAML format. For example:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},'  # -- Slack config for the events, tests, testsuites and channels\n  slackConfig:\n    - ChannelID: ""\n      selector: \n        label1: "value1"\n      testName: \n        - "sanity"\n        - "testName2"\n      testSuiteName:\n        - "test-suite1"\n        - "test-suite2"\n      events:\n        - "start-test"\n        - "end-test-success"\n        - "end-test-failed"\n        - "end-test-aborted"\n        - "end-test-timeout"\n        - "start-testsuite"\n        - "end-testsuite-success"\n        - "end-testsuite-failed"\n        - "end-testsuite-aborted"\n        - "end-testsuite-timeout"\n')),(0,a.kt)("h2",{id:"configure-message-template"},"Configure Message Template"),(0,a.kt)("p",null,"The default message is ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/kubeshop/helm-charts/blob/311ff9f6fc38dfb5196b91a6f63ee7d3f59f7f4b/charts/testkube-api/slack-template.json"},"/charts/testkube-api/slack-template.json")," and is written using ",(0,a.kt)("a",{parentName:"p",href:"https://app.slack.com/block-kit-builder"},"Slack block kit builder")," and Golang templates. You can customize it depending on your needs. The following structure is referenced in the template where it is getting the data to show:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"type MessageArgs struct {\n    ExecutionID   string\n    ExecutionName string\n    EventType     string\n    Namespace     string\n    Labels        string\n    TestName      string\n    TestType      string\n    Status        string\n    FailedSteps   int\n    TotalSteps    int\n    StartTime     string\n    EndTime       string\n    Duration      string\n}\n")),(0,a.kt)("h2",{id:"add-the-testkube-bot-to-a-channel"},"Add the Testkube bot to a Channel"),(0,a.kt)("p",null,"If the goal is to receive all the notifications in one channel, add the Testkube bot to the channel and leave the ",(0,a.kt)("inlineCode",{parentName:"p"},"ChannelID")," field empty in the ",(0,a.kt)("inlineCode",{parentName:"p"},"slack-config.json")," file.\n",(0,a.kt)("img",{alt:"img.gif",src:n(86927).Z,width:"600",height:"350"})))}d.isMDXComponent=!0},86927:(e,t,n)=>{n.d(t,{Z:()=>s});const s=n.p+"assets/images/add-testkube-bot-to-channel-c8b68124472c69707767cb5b24c23e24.gif"},65998:(e,t,n)=>{n.d(t,{Z:()=>s});const s=n.p+"assets/images/add-testkube-bot-to-workspace-798438c0fea9e7999484c2eb57cf71fa.gif"}}]);