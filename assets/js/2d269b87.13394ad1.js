"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[7487],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>m});var s=n(67294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function r(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);t&&(s=s.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,s)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?r(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):r(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,s,a=function(e,t){if(null==e)return{};var n,s,a={},r=Object.keys(e);for(s=0;s<r.length;s++)n=r[s],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(s=0;s<r.length;s++)n=r[s],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var u=s.createContext({}),l=function(e){var t=s.useContext(u),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},c=function(e){var t=l(e.components);return s.createElement(u.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return s.createElement(s.Fragment,{},t)}},d=s.forwardRef((function(e,t){var n=e.components,a=e.mdxType,r=e.originalType,u=e.parentName,c=o(e,["components","mdxType","originalType","parentName"]),d=l(n),m=a,g=d["".concat(u,".").concat(m)]||d[m]||p[m]||r;return n?s.createElement(g,i(i({ref:t},c),{},{components:n})):s.createElement(g,i({ref:t},c))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var r=n.length,i=new Array(r);i[0]=d;var o={};for(var u in t)hasOwnProperty.call(t,u)&&(o[u]=t[u]);o.originalType=e,o.mdxType="string"==typeof e?e:a,i[1]=o;for(var l=2;l<r;l++)i[l]=n[l];return s.createElement.apply(null,i)}return s.createElement.apply(null,n)}d.displayName="MDXCreateElement"},75197:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>u,contentTitle:()=>i,default:()=>p,frontMatter:()=>r,metadata:()=>o,toc:()=>l});var s=n(87462),a=(n(67294),n(3905));const r={},i="Running Tests",o={unversionedId:"guides/tests/tests-running",id:"guides/tests/tests-running",title:"Running Tests",description:"Tests are stored in Kubernetes clusters as Custom Resources. Testkube tests are reusable and can get results with the use of kubectl testkube plugin or with an API.",source:"@site/docs/guides/tests/tests-running.md",sourceDirName:"guides/tests",slug:"/guides/tests/tests-running",permalink:"/guides/tests/tests-running",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/guides/tests/tests-running.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Creating Tests",permalink:"/guides/tests/tests-creating"},next:{title:"Getting Results",permalink:"/guides/tests/tests-getting-results"}},u={},l=[{value:"<strong>Running</strong>",id:"running",level:2},{value:"<strong>Standard Run Command</strong>",id:"standard-run-command",level:3},{value:"<strong>Run with Watch for Changes</strong>",id:"run-with-watch-for-changes",level:2},{value:"<strong>Passing Parameters</strong>",id:"passing-parameters",level:3},{value:"<strong>Mapping local files</strong>",id:"mapping-local-files",level:3},{value:"<strong>Summary</strong>",id:"summary",level:2}],c={toc:l};function p(e){let{components:t,...n}=e;return(0,a.kt)("wrapper",(0,s.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"running-tests"},"Running Tests"),(0,a.kt)("p",null,"Tests are stored in Kubernetes clusters as Custom Resources. Testkube tests are reusable and can get results with the use of kubectl testkube plugin or with an API."),(0,a.kt)("h2",{id:"running"},(0,a.kt)("strong",{parentName:"h2"},"Running")),(0,a.kt)("p",null,"Running tests looks the same for any type of test.\nIn this example, we have previously created a test with the name ",(0,a.kt)("inlineCode",{parentName:"p"},"api-incluster-test"),"."),(0,a.kt)("h3",{id:"standard-run-command"},(0,a.kt)("strong",{parentName:"h3"},"Standard Run Command")),(0,a.kt)("p",null,"This is an example of the simplest run command:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube run test api-incluster-test\n")),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"Type          : postman/collection\nName          : api-incluster-test\nExecution ID  : 615d6398b046f8fbd3d955d4\nExecution name: openly-full-bream\n\nTest queued for execution\nUse the following command to get test execution details:\n$ kubectl testkube get execution 615d6398b046f8fbd3d955d4\n\nOr watch test execution until complete:\n$ kubectl testkube watch execution 615d6398b046f8fbd3d955d4\n\n")),(0,a.kt)("p",null,"Testkube will inform us about possible commands to get test results:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"kubectl testkube get execution 615d6398b046f8fbd3d955d4")," to get execution details."),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"kubectl testkube watch execution 615d6398b046f8fbd3d955d4")," to watch the current pending execution. Watch will also get details when a test is completed and is good for long running tests to lock your terminal until test execution completes.")),(0,a.kt)("h2",{id:"run-with-watch-for-changes"},(0,a.kt)("strong",{parentName:"h2"},"Run with Watch for Changes")),(0,a.kt)("p",null,"If we want to wait until a test execution completes we can pass the ",(0,a.kt)("inlineCode",{parentName:"p"},"-f")," flag (follow) to the test run command:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube run test api-incluster-test -f\n")),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"Type          : postman/collection\nName          : api-incluster-test\nExecution ID  : 615d7e1ab046f8fbd3d955d6\nExecution name: monthly-sure-finch\n\nTest queued for execution\nUse the following command to get test execution details:\n$ kubectl testkube get execution 615d7e1ab046f8fbd3d955d6\n\nOr watch test execution until complete:\n$ kubectl testkube watch execution 615d7e1ab046f8fbd3d955d6\n\n\nWatching for changes\nStatus: pending, Duration: 222.387ms\nStatus: pending, Duration: 1.210689s\nStatus: pending, Duration: 2.201346s\nStatus: pending, Duration: 3.198539s\nStatus: success, Duration: 595ms\n\nGetting results\nName: monthly-sure-finch, Status: success, Duration: 595ms\nnewman\n\nAPI-Health\n\n\u2192 Health\n  GET http://testkube-api-server:8088/health [200 OK, 124B, 282ms]\n  \u2713  Status code is 200\n\n\u250c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u252c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u252c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2510\n\u2502                         \u2502           executed \u2502            failed \u2502\n\u251c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2524\n\u2502              iterations \u2502                  1 \u2502                 0 \u2502\n\u251c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2524\n\u2502                requests \u2502                  1 \u2502                 0 \u2502\n\u251c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2524\n\u2502            test-scripts \u2502                  2 \u2502                 0 \u2502\n\u251c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2524\n\u2502      prerequest-scripts \u2502                  1 \u2502                 0 \u2502\n\u251c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u253c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2524\n\u2502              assertions \u2502                  1 \u2502                 0 \u2502\n\u251c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2534\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2534\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2524\n\u2502 total run duration: 519ms                                        \u2502\n\u251c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2524\n\u2502 total data received: 8B (approx)                                 \u2502\n\u251c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2524\n\u2502 average response time: 282ms [min: 282ms, max: 282ms, s.d.: 0\xb5s] \u2502\n\u2514\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2518\nUse the following command to get test execution details:\n$ kubectl testkube get execution 615d7e1ab046f8fbd3d955d6\n\nTest execution completed in 595ms\n")),(0,a.kt)("p",null,"This command will wait until the test execution completes."),(0,a.kt)("h3",{id:"passing-parameters"},(0,a.kt)("strong",{parentName:"h3"},"Passing Parameters")),(0,a.kt)("p",null,"For some 'real world' tests, configuration variables are passed in order to run them on different environments or with different test configurations."),(0,a.kt)("p",null,"Let's assume that our example Cypress test needs the ",(0,a.kt)("inlineCode",{parentName:"p"},"testparam")," parameter with the value ",(0,a.kt)("inlineCode",{parentName:"p"},"testvalue"),"."),(0,a.kt)("p",null,"This is done by using the ",(0,a.kt)("inlineCode",{parentName:"p"},"-p")," parameter. If you need to pass more parameters, simply pass multiple ",(0,a.kt)("inlineCode",{parentName:"p"},"-p")," flags."),(0,a.kt)("p",null,"It's possible to pass parameters securely to the executed test. It's necessary to use ",(0,a.kt)("inlineCode",{parentName:"p"},"--secret")," flag,\nwhich contains a key value pair - a name of the kubernetes secret and a secret key.\nPass it multiple times if needed."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube run test kubeshop-cypress -p testparam=testvalue -f --secret secret-name=secret-key\n")),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"Type          : cypress/project\nName          : kubeshop-cypress\nExecution ID  : 615d5372b046f8fbd3d955d2\nExecution name: nominally-able-glider\n\nTest queued for execution\nUse the following command to get test execution details:\n$ kubectl testkube get execution 615d5372b046f8fbd3d955d2\n\nor watch test execution until complete:\n$ kubectl testkube watch execution 615d5372b046f8fbd3d955d2\n\n\nWatching for changes\nStatus: queued, Duration: 0s\nStatus: pending, Duration: 383.064ms\n....\nStatus: pending, Duration: 1m45.405939s\nStatus: success, Duration: 1m45.405939s\n\nGetting results\nName: nominally-able-glider, Status: success, Duration: 2562047h47m16.854775807s\n\n====================================================================================================\n\n  (Run Starting)\n\n  \u250c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2510\n  \u2502 Cypress:    8.5.0                                                                              \u2502\n  \u2502 Browser:    Electron 91 (headless)                                                             \u2502\n  \u2502 Specs:      1 found (simple-test.js)                                                           \u2502\n  \u2514\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2518\n\n\n\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\n\n  Running:  simple-test.js                                                                  (1 of 1)\n\n  (Results)\n\n  \u250c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2510\n  \u2502 Tests:        1                                                                                \u2502\n  \u2502 Passing:      1                                                                                \u2502\n  \u2502 Failing:      0                                                                                \u2502\n  \u2502 Pending:      0                                                                                \u2502\n  \u2502 Skipped:      0                                                                                \u2502\n  \u2502 Screenshots:  0                                                                                \u2502\n  \u2502 Video:        true                                                                             \u2502\n  \u2502 Duration:     19 seconds                                                                       \u2502\n  \u2502 Spec Ran:     simple-test.js                                                                   \u2502\n  \u2514\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2518\n\n\n  (Video)\n\n  -  Started processing:  Compressing to 32 CRF\n    Compression progress:  39%\n    Compression progress:  81%\n  -  Finished processing: /tmp/testkube-scripts531364188/repo/examples/cypress/videos/   (30 seconds)\n                          simple-test.js.mp4\n\n    Compression progress:  100%\n\n====================================================================================================\n\n  (Run Finished)\n\n\n       Spec                                              Tests  Passing  Failing  Pending  Skipped\n  \u250c\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2510\n  \u2502 \u2714  simple-test.js                           00:19        1        1        -        -        - \u2502\n  \u2514\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2500\u2518\n    \u2714  All specs passed!                        00:19        1        1        -        -        -\n\nUse the following command to get test execution details:\n$ kubectl testkube get execution 615d5372b046f8fbd3d955d2\n\nTest execution completed in 1m45.405939s\n")),(0,a.kt)("h3",{id:"mapping-local-files"},(0,a.kt)("strong",{parentName:"h3"},"Mapping local files")),(0,a.kt)("p",null,"Local files can be set on the execution of a Testkube Test. Pass the file in the format ",(0,a.kt)("inlineCode",{parentName:"p"},"source_path:destination_path")," using the flag ",(0,a.kt)("inlineCode",{parentName:"p"},"--copy-files"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},'kubectl testkube run test maven-example-file-test --copy-files "/Users/local_user/local_maven_settings.xml:/tmp/settings.xml" --args "--settings" --args "/tmp/settings.xml" -v "TESTKUBE_MAVEN=true"\n')),(0,a.kt)("h2",{id:"summary"},(0,a.kt)("strong",{parentName:"h2"},"Summary")),(0,a.kt)("p",null,"As we can see, running tests in Kubernetes cluster is really easy with use of the Testkube kubectl plugin!"))}p.isMDXComponent=!0}}]);