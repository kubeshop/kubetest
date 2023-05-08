"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[3665],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>d});var a=n(67294);function s(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function r(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?r(Object(n),!0).forEach((function(t){s(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):r(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,a,s=function(e,t){if(null==e)return{};var n,a,s={},r=Object.keys(e);for(a=0;a<r.length;a++)n=r[a],t.indexOf(n)>=0||(s[n]=e[n]);return s}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(a=0;a<r.length;a++)n=r[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(s[n]=e[n])}return s}var l=a.createContext({}),u=function(e){var t=a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},c=function(e){var t=u(e.components);return a.createElement(l.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},m=a.forwardRef((function(e,t){var n=e.components,s=e.mdxType,r=e.originalType,l=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),m=u(n),d=s,h=m["".concat(l,".").concat(d)]||m[d]||p[d]||r;return n?a.createElement(h,o(o({ref:t},c),{},{components:n})):a.createElement(h,o({ref:t},c))}));function d(e,t){var n=arguments,s=t&&t.mdxType;if("string"==typeof e||s){var r=n.length,o=new Array(r);o[0]=m;var i={};for(var l in t)hasOwnProperty.call(t,l)&&(i[l]=t[l]);i.originalType=e,i.mdxType="string"==typeof e?e:s,o[1]=i;for(var u=2;u<r;u++)o[u]=n[u];return a.createElement.apply(null,o)}return a.createElement.apply(null,n)}m.displayName="MDXCreateElement"},78604:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>o,default:()=>p,frontMatter:()=>r,metadata:()=>i,toc:()=>u});var a=n(87462),s=(n(67294),n(3905));const r={},o="Step 3 - Creating Your First Test",i={unversionedId:"articles/step3-creating-first-test",id:"articles/step3-creating-first-test",title:"Step 3 - Creating Your First Test",description:"Using the CLI or the Dashboard",source:"@site/docs/articles/step3-creating-first-test.md",sourceDirName:"articles",slug:"/articles/step3-creating-first-test",permalink:"/articles/step3-creating-first-test",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/articles/step3-creating-first-test.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Step 2 - Install Testkube Cluster Components Using Testkube's CLI",permalink:"/articles/step2-installing-cluster-components"},next:{title:"Creating Tests",permalink:"/articles/creating-tests"}},l={},u=[{value:"Using the CLI or the Dashboard",id:"using-the-cli-or-the-dashboard",level:2},{value:"Kubernetes-native Tests",id:"kubernetes-native-tests",level:2},{value:"Creating a Postman Test",id:"creating-a-postman-test",level:2},{value:"Starting a New Test Execution",id:"starting-a-new-test-execution",level:2},{value:"Getting the Result of a Test Execution",id:"getting-the-result-of-a-test-execution",level:2},{value:"Changing the Output Format",id:"changing-the-output-format",level:2}],c={toc:u};function p(e){let{components:t,...n}=e;return(0,s.kt)("wrapper",(0,a.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,s.kt)("h1",{id:"step-3---creating-your-first-test"},"Step 3 - Creating Your First Test"),(0,s.kt)("h2",{id:"using-the-cli-or-the-dashboard"},"Using the CLI or the Dashboard"),(0,s.kt)("p",null,"You can create your first test using the CLI or the Testkube Dashboard, both are great options!"),(0,s.kt)("p",null,"To explore the Testkube dashboard, run the command:"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre",className:"language-sh"},"testkube dashboard\n")),(0,s.kt)("h2",{id:"kubernetes-native-tests"},"Kubernetes-native Tests"),(0,s.kt)("p",null,"Tests in Testkube are created as a Custom Resource in Kubernetes and live inside your cluster."),(0,s.kt)("p",null,"You can create your tests directly as a Custom Resource, or use the CLI or the Testkube Dashboard to create them."),(0,s.kt)("p",null,"This section provides an example of creating a ",(0,s.kt)("em",{parentName:"p"},"Postman")," test. Nevertheless, Testkube supports a long ",(0,s.kt)("a",{parentName:"p",href:"../category/test-types"},"list of testing tools"),"."),(0,s.kt)("h2",{id:"creating-a-postman-test"},"Creating a Postman Test"),(0,s.kt)("p",null,"First, let's create a ",(0,s.kt)("inlineCode",{parentName:"p"},"postman-collection.json")," file containing a Postman test (the file content should look similar to the one below):"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre",className:"language-json",metastring:'title="postman-collection.json"',title:'"postman-collection.json"'},'{\n  "info": {\n    "_postman_id": "8af42c21-3e31-49c1-8b27-d6e60623a180",\n    "name": "Kubeshop",\n    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"\n  },\n  "item": [\n    {\n      "name": "Home",\n      "event": [\n        {\n          "listen": "test",\n          "script": {\n            "exec": [\n              "pm.test(\\"Body matches string\\", function () {",\n              "    pm.expect(pm.response.text()).to.include(\\"Accelerator\\");",\n              "});"\n            ],\n            "type": "text/javascript"\n          }\n        }\n      ],\n      "request": {\n        "method": "GET",\n        "header": [],\n        "url": {\n          "raw": "https://kubeshop.io/",\n          "protocol": "https",\n          "host": ["kubeshop", "io"],\n          "path": [""]\n        }\n      },\n      "response": []\n    },\n    {\n      "name": "Team",\n      "event": [\n        {\n          "listen": "test",\n          "script": {\n            "exec": [\n              "pm.test(\\"Status code is 200\\", function () {",\n              "    pm.response.to.have.status(200);",\n              "});"\n            ],\n            "type": "text/javascript"\n          }\n        }\n      ],\n      "request": {\n        "method": "GET",\n        "header": [],\n        "url": {\n          "raw": "https://kubeshop.io/our-team",\n          "protocol": "https",\n          "host": ["kubeshop", "io"],\n          "path": ["our-team"]\n        }\n      },\n      "response": []\n    }\n  ]\n}\n')),(0,s.kt)("p",null,"And create the test by running the command:"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre",className:"language-sh"},"testkube create test --file postman_collection.json --type postman/collection --name my-first-test\n")),(0,s.kt)("admonition",{type:"note"},(0,s.kt)("p",{parentName:"admonition"},"This example is testing that the website ",(0,s.kt)("a",{parentName:"p",href:"https://kubeshop.io"},"https://kubeshop.io")," is returning a ",(0,s.kt)("inlineCode",{parentName:"p"},"200")," status code. It demostrates a way of creating a simple Postman test. In practice, you would test an internal Kubernetes service.")),(0,s.kt)("h2",{id:"starting-a-new-test-execution"},"Starting a New Test Execution"),(0,s.kt)("p",null,"After our test is defined as a Custom Resource, we can run it:"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre",className:"language-sh"},"testkube run test my-first-test\n")),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre",className:"language-sh",metastring:'title="Expected output:"',title:'"Expected','output:"':!0},"Type:              postman/collection\nName:              my-first-test\nExecution ID:      63f4d0910ca9ed26798741ca\nExecution name:    my-frst-test-1\nExecution number:  1\nStatus:            running\nStart time:        2023-02-21 14:09:21.163713965 +0000 UTC\nEnd time:          0001-01-01 00:00:00 +0000 UTC\nDuration:\n\nTest execution started\nWatch test execution until complete:\n$ testkube watch execution my-frst-test-1\n\n\nUse following command to get test execution details:\n$ testkube get execution my-frst-test-1\n")),(0,s.kt)("h2",{id:"getting-the-result-of-a-test-execution"},"Getting the Result of a Test Execution"),(0,s.kt)("p",null,"To see the result of a Test Execution, first, you need to get the Test Execution list:"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre",className:"language-sh"},"testkube get executions\n")),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre",className:"language-sh",metastring:'title=Expected output:"',title:"Expected",'output:"':!0},"  ID                       | NAME              | TEST NAME       | TYPE               | STATUS | LABELS\n---------------------------+-------------------+-----------------+--------------------+--------+---------\n  63f4d0910ca9ed26798741ca | my-first-test-1   | my-first-test   | postman/collection | passed |\n                           |                   |                 |                    |        |\n")),(0,s.kt)("p",null,"Copy the ID of the test and see the full details of the execution by running:"),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre",className:"language-sh"},"testkube get execution 63f4d0910ca9ed26798741ca\n")),(0,s.kt)("pre",null,(0,s.kt)("code",{parentName:"pre",className:"language-sh",metastring:'title="Expected output:"',title:'"Expected','output:"':!0},"# ... test details\nTest execution completed with success in 14.342s \ud83e\udd47\n")),(0,s.kt)("h2",{id:"changing-the-output-format"},"Changing the Output Format"),(0,s.kt)("p",null,"For lists and details, you can use different output formats via the ",(0,s.kt)("inlineCode",{parentName:"p"},"--output")," flag. The following formats are currently supported:"),(0,s.kt)("ul",null,(0,s.kt)("li",{parentName:"ul"},(0,s.kt)("inlineCode",{parentName:"li"},"RAW")," - Raw output from the given executor (e.g., for Postman collection, it's terminal text with colors and tables)."),(0,s.kt)("li",{parentName:"ul"},(0,s.kt)("inlineCode",{parentName:"li"},"JSON")," - Test run data are encoded in JSON."),(0,s.kt)("li",{parentName:"ul"},(0,s.kt)("inlineCode",{parentName:"li"},"GO")," - For go-template formatting (like in Docker and Kubernetes), you'll need to add the ",(0,s.kt)("inlineCode",{parentName:"li"},"--go-template")," flag with a custom format. The default is ",(0,s.kt)("inlineCode",{parentName:"li"},'{{ . | printf("%+v") }}'),". This will help you check available fields.")))}p.isMDXComponent=!0}}]);