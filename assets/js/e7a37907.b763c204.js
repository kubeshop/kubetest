"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[4880],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>m});var r=n(67294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function l(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?l(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):l(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},l=Object.keys(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(r=0;r<l.length;r++)n=l[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var o=r.createContext({}),u=function(e){var t=r.useContext(o),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},c=function(e){var t=u(e.components);return r.createElement(o.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,l=e.originalType,o=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),d=u(n),m=a,h=d["".concat(o,".").concat(m)]||d[m]||p[m]||l;return n?r.createElement(h,s(s({ref:t},c),{},{components:n})):r.createElement(h,s({ref:t},c))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=n.length,s=new Array(l);s[0]=d;var i={};for(var o in t)hasOwnProperty.call(t,o)&&(i[o]=t[o]);i.originalType=e,i.mdxType="string"==typeof e?e:a,s[1]=i;for(var u=2;u<l;u++)s[u]=n[u];return r.createElement.apply(null,s)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},85162:(e,t,n)=>{n.d(t,{Z:()=>s});var r=n(67294),a=n(86010);const l="tabItem_Ymn6";function s(e){let{children:t,hidden:n,className:s}=e;return r.createElement("div",{role:"tabpanel",className:(0,a.Z)(l,s),hidden:n},t)}},74866:(e,t,n)=>{n.d(t,{Z:()=>T});var r=n(87462),a=n(67294),l=n(86010),s=n(12466),i=n(76775),o=n(91980),u=n(67392),c=n(50012);function p(e){return function(e){return a.Children.map(e,(e=>{if((0,a.isValidElement)(e)&&"value"in e.props)return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)}))}(e).map((e=>{let{props:{value:t,label:n,attributes:r,default:a}}=e;return{value:t,label:n,attributes:r,default:a}}))}function d(e){const{values:t,children:n}=e;return(0,a.useMemo)((()=>{const e=t??p(n);return function(e){const t=(0,u.l)(e,((e,t)=>e.value===t.value));if(t.length>0)throw new Error(`Docusaurus error: Duplicate values "${t.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[t,n])}function m(e){let{value:t,tabValues:n}=e;return n.some((e=>e.value===t))}function h(e){let{queryString:t=!1,groupId:n}=e;const r=(0,i.k6)(),l=function(e){let{queryString:t=!1,groupId:n}=e;if("string"==typeof t)return t;if(!1===t)return null;if(!0===t&&!n)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return n??null}({queryString:t,groupId:n});return[(0,o._X)(l),(0,a.useCallback)((e=>{if(!l)return;const t=new URLSearchParams(r.location.search);t.set(l,e),r.replace({...r.location,search:t.toString()})}),[l,r])]}function b(e){const{defaultValue:t,queryString:n=!1,groupId:r}=e,l=d(e),[s,i]=(0,a.useState)((()=>function(e){let{defaultValue:t,tabValues:n}=e;if(0===n.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(t){if(!m({value:t,tabValues:n}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${t}" but none of its children has the corresponding value. Available values are: ${n.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return t}const r=n.find((e=>e.default))??n[0];if(!r)throw new Error("Unexpected error: 0 tabValues");return r.value}({defaultValue:t,tabValues:l}))),[o,u]=h({queryString:n,groupId:r}),[p,b]=function(e){let{groupId:t}=e;const n=function(e){return e?`docusaurus.tab.${e}`:null}(t),[r,l]=(0,c.Nk)(n);return[r,(0,a.useCallback)((e=>{n&&l.set(e)}),[n,l])]}({groupId:r}),f=(()=>{const e=o??p;return m({value:e,tabValues:l})?e:null})();(0,a.useLayoutEffect)((()=>{f&&i(f)}),[f]);return{selectedValue:s,selectValue:(0,a.useCallback)((e=>{if(!m({value:e,tabValues:l}))throw new Error(`Can't select invalid tab value=${e}`);i(e),u(e),b(e)}),[u,b,l]),tabValues:l}}var f=n(72389);const y="tabList__CuJ",k="tabItem_LNqP";function g(e){let{className:t,block:n,selectedValue:i,selectValue:o,tabValues:u}=e;const c=[],{blockElementScrollPositionUntilNextRender:p}=(0,s.o5)(),d=e=>{const t=e.currentTarget,n=c.indexOf(t),r=u[n].value;r!==i&&(p(t),o(r))},m=e=>{var t;let n=null;switch(e.key){case"Enter":d(e);break;case"ArrowRight":{const t=c.indexOf(e.currentTarget)+1;n=c[t]??c[0];break}case"ArrowLeft":{const t=c.indexOf(e.currentTarget)-1;n=c[t]??c[c.length-1];break}}null==(t=n)||t.focus()};return a.createElement("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,l.Z)("tabs",{"tabs--block":n},t)},u.map((e=>{let{value:t,label:n,attributes:s}=e;return a.createElement("li",(0,r.Z)({role:"tab",tabIndex:i===t?0:-1,"aria-selected":i===t,key:t,ref:e=>c.push(e),onKeyDown:m,onClick:d},s,{className:(0,l.Z)("tabs__item",k,null==s?void 0:s.className,{"tabs__item--active":i===t})}),n??t)})))}function v(e){let{lazy:t,children:n,selectedValue:r}=e;if(n=Array.isArray(n)?n:[n],t){const e=n.find((e=>e.props.value===r));return e?(0,a.cloneElement)(e,{className:"margin-top--md"}):null}return a.createElement("div",{className:"margin-top--md"},n.map(((e,t)=>(0,a.cloneElement)(e,{key:t,hidden:e.props.value!==r}))))}function w(e){const t=b(e);return a.createElement("div",{className:(0,l.Z)("tabs-container",y)},a.createElement(g,(0,r.Z)({},e,t)),a.createElement(v,(0,r.Z)({},e,t)))}function T(e){const t=(0,f.Z)();return a.createElement(w,(0,r.Z)({key:String(t)},e))}},13538:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>o,default:()=>m,frontMatter:()=>i,metadata:()=>u,toc:()=>p});var r=n(87462),a=(n(67294),n(3905)),l=n(74866),s=n(85162);const i={},o="Artillery.io",u={unversionedId:"test-types/executor-artillery",id:"test-types/executor-artillery",title:"Artillery.io",description:"Artillery.io is an open-source load testing tool. It's designed to be both straightforward in configuration (YAML files), and powerful. The Artillery executor allow you to run Artillery tests with Testkube.",source:"@site/docs/test-types/executor-artillery.mdx",sourceDirName:"test-types",slug:"/test-types/executor-artillery",permalink:"/testkube/test-types/executor-artillery",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/test-types/executor-artillery.mdx",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Test Types",permalink:"/testkube/category/test-types"},next:{title:"cURL",permalink:"/testkube/test-types/executor-curl"}},c={},p=[{value:"<strong>Test Environment</strong>",id:"test-environment",level:2},{value:"<strong>Create a Test Manifest</strong>",id:"create-a-test-manifest",level:2},{value:"Create a New Testkube Test",id:"create-a-new-testkube-test",level:2},{value:"Running a Test",id:"running-a-test",level:2},{value:"Getting Test Results",id:"getting-test-results",level:2},{value:"Additional examples",id:"additional-examples",level:2}],d={toc:p};function m(e){let{components:t,...i}=e;return(0,a.kt)("wrapper",(0,r.Z)({},d,i,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"artilleryio"},"Artillery.io"),(0,a.kt)("p",null,"Artillery.io is an open-source load testing tool. It's designed to be both straightforward in configuration (YAML files), and powerful. The Artillery executor allow you to run Artillery tests with Testkube."),(0,a.kt)("h2",{id:"test-environment"},(0,a.kt)("strong",{parentName:"h2"},"Test Environment")),(0,a.kt)("p",null,"Let's assume that our SUT (Service Under Test) is an internal Kubernetes service which has\nClusterIP ",(0,a.kt)("inlineCode",{parentName:"p"},"Service")," created and is exposed on port ",(0,a.kt)("inlineCode",{parentName:"p"},"8088"),". The service name is ",(0,a.kt)("inlineCode",{parentName:"p"},"testkube-api-server"),"\nand is exposing the ",(0,a.kt)("inlineCode",{parentName:"p"},"/health")," endpoint that we want to test."),(0,a.kt)("p",null,"To call the SUT inside a cluster:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"curl http://testkube-api-server:8088/health\n")),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"200 OK\n")),(0,a.kt)("h2",{id:"create-a-test-manifest"},(0,a.kt)("strong",{parentName:"h2"},"Create a Test Manifest")),(0,a.kt)("p",null,"The Artillery tests are defined in declarative manner, as YAML files.",(0,a.kt)("br",{parentName:"p"}),"\n","The test should warm up our service a little bit first, then we can hit a little harder."),(0,a.kt)("p",null,"Let's save our test into ",(0,a.kt)("inlineCode",{parentName:"p"},"test.yaml")," file with the content below:   "),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-yaml"},'config:\n  target: "http://testkube-api-server:8088"\n  phases:\n    - duration: 6\n      arrivalRate: 5\n      name: Warm up\n    - duration: 120\n      arrivalRate: 5\n      rampTo: 50\n      name: Ramp up load\n    - duration: 60\n      arrivalRate: 50\n      name: Sustained load\nscenarios:\n  - name: "Check health endpoint"\n    flow:\n      - get:\n          url: "/health"\n')),(0,a.kt)("p",null,"Our test is ready but how do we run it in a Kubernetes cluster? Testkube will help you with that! "),(0,a.kt)("p",null,"Let's create a new Testkube test based on the saved Artillery test definition."),(0,a.kt)("h2",{id:"create-a-new-testkube-test"},"Create a New Testkube Test"),(0,a.kt)("p",null,"If you want to upload a test file directly (like in this example) you can use Dashboard, or CLI - depending on your preferences."),(0,a.kt)(l.Z,{groupId:"dashboard-cli",mdxType:"Tabs"},(0,a.kt)(s.Z,{value:"dash",label:"Dashboard",mdxType:"TabItem"},(0,a.kt)("p",null,"If you prefer to use the Dashboard, just go to Tests, and click ",(0,a.kt)("inlineCode",{parentName:"p"},"Add a new test")," button. Then you need to fill in the test Name, choose the test Type (",(0,a.kt)("inlineCode",{parentName:"p"},"artillery/test"),"), Test Source (",(0,a.kt)("inlineCode",{parentName:"p"},"File"),", which allow you to upload specific file), and choose the File.\n",(0,a.kt)("img",{alt:"Container executor creation dialog",src:n(18285).Z,width:"879",height:"684"}))),(0,a.kt)(s.Z,{value:"cli",label:"CLI",mdxType:"TabItem"},(0,a.kt)("p",null,"If you prefer using the CLI instead, you can create the test with ",(0,a.kt)("inlineCode",{parentName:"p"},"testkube create test"),".\nYou need to set the test:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"--name")," (for example, ",(0,a.kt)("inlineCode",{parentName:"li"},"artillery-api-test"),")"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"--type")," (in this case ",(0,a.kt)("inlineCode",{parentName:"li"},"artillery/test"),")"),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("inlineCode",{parentName:"li"},"--file")," which is a path to your test file (in this case ",(0,a.kt)("inlineCode",{parentName:"li"},"test.yaml"),")")),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"testkube create test --name artillery-api-test --type artillery/test --file test.yaml\n")),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"Test created  \ud83e\udd47\n")))),(0,a.kt)("h2",{id:"running-a-test"},"Running a Test"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"$ testkube run test artillery-api-test                                                                                                                       \nType:              artillery/test\nName:              artillery-api-test\nExecution ID:      63ee9ca6872e05f0ea790d73\nExecution name:    artillery-api-test-1\nExecution number:  1\nStatus:            running\nStart time:        2023-02-16 21:14:14.451905194 +0000 UTC\nEnd time:          0001-01-01 00:00:00 +0000 UTC\nDuration:          \n\n\n\nTest execution started\nWatch test execution until complete:\n$ kubectl testkube watch execution artillery-api-test-1\n\n\nUse following command to get test execution details:\n$ kubectl testkube get execution artillery-api-test-1\n")),(0,a.kt)("p",null,"You can also watch your test results in real-time with ",(0,a.kt)("inlineCode",{parentName:"p"},"-f"),' flag (like "follow"). '),(0,a.kt)("p",null,"Test runs can be named. If no name is passed, Testkube will autogenerate a name."),(0,a.kt)("h2",{id:"getting-test-results"},"Getting Test Results"),(0,a.kt)("p",null,"Let's get back our finished test results. The test report and output will be stored in Testkube storage to revisit when necessary."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"testkube get execution artillery-api-test-1                                               \nID:         63ee9cd8872e05f0ea790d76\nName:       artillery-api-test-1\nNumber:            1\nTest name:         artillery-api-test\nType:              artillery/test\nStatus:            passed\nStart time:        2023-02-16 21:15:04.979 +0000 UTC\nEnd time:          2023-02-16 21:18:19.463 +0000 UTC\nDuration:          00:03:14\n\n...\n... (long output)\n...\n\nAll VUs finished. Total time: 3 minutes, 7 seconds\n\n--------------------------------\nSummary report @ 21:18:16(+0000)\n--------------------------------\n\nhttp.codes.200: ................................................................ 6330\nhttp.request_rate: ............................................................. 33/sec\nhttp.requests: ................................................................. 6330\nhttp.response_time:\n  min: ......................................................................... 0\n  max: ......................................................................... 11\n  median: ...................................................................... 0\n  p95: ......................................................................... 1\n  p99: ......................................................................... 2\nhttp.responses: ................................................................ 6330\nvusers.completed: .............................................................. 6330\nvusers.created: ................................................................ 6330\nvusers.created_by_name.Check health endpoint: .................................. 6330\nvusers.failed: ................................................................. 0\nvusers.session_length:\n  min: ......................................................................... 0.9\n  max: ......................................................................... 25.6\n  median: ...................................................................... 1.3\n  p95: ......................................................................... 3.3\n  p99: ......................................................................... 9.5\nLog file: /tmp/test-report.json\n\n\nTest execution completed with success in 3m14.484s \ud83e\udd47\n\n")),(0,a.kt)("h2",{id:"additional-examples"},"Additional examples"),(0,a.kt)("p",null,"Additional Artillery examples can be found in the Testkube repository ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/kubeshop/testkube/blob/main/test/artillery/executor-smoke/"},"here"),"."))}m.isMDXComponent=!0},18285:(e,t,n)=>{n.d(t,{Z:()=>r});const r=n.p+"assets/images/dashboard-create-artillery-api-test-727a136c9175a7fb4c6efc18126652b2.png"}}]);