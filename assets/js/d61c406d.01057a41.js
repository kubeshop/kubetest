"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[1987,4701,9960],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>m});var a=n(67294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function r(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?r(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):r(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,a,i=function(e,t){if(null==e)return{};var n,a,i={},r=Object.keys(e);for(a=0;a<r.length;a++)n=r[a],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(a=0;a<r.length;a++)n=r[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var l=a.createContext({}),u=function(e){var t=a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},c=function(e){var t=u(e.components);return a.createElement(l.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},d=a.forwardRef((function(e,t){var n=e.components,i=e.mdxType,r=e.originalType,l=e.parentName,c=o(e,["components","mdxType","originalType","parentName"]),d=u(n),m=i,k=d["".concat(l,".").concat(m)]||d[m]||p[m]||r;return n?a.createElement(k,s(s({ref:t},c),{},{components:n})):a.createElement(k,s({ref:t},c))}));function m(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var r=n.length,s=new Array(r);s[0]=d;var o={};for(var l in t)hasOwnProperty.call(t,l)&&(o[l]=t[l]);o.originalType=e,o.mdxType="string"==typeof e?e:i,s[1]=o;for(var u=2;u<r;u++)s[u]=n[u];return a.createElement.apply(null,s)}return a.createElement.apply(null,n)}d.displayName="MDXCreateElement"},23612:(e,t,n)=>{n.d(t,{Z:()=>k});var a=n(67294),i=n(86010),r=n(35281),s=n(95999);const o="admonition_LlT9",l="admonitionHeading_tbUL",u="admonitionIcon_kALy",c="admonitionContent_S0QG";const p={note:{infimaClassName:"secondary",iconComponent:function(){return a.createElement("svg",{viewBox:"0 0 14 16"},a.createElement("path",{fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))},label:a.createElement(s.Z,{id:"theme.admonition.note",description:"The default label used for the Note admonition (:::note)"},"note")},tip:{infimaClassName:"success",iconComponent:function(){return a.createElement("svg",{viewBox:"0 0 12 16"},a.createElement("path",{fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"}))},label:a.createElement(s.Z,{id:"theme.admonition.tip",description:"The default label used for the Tip admonition (:::tip)"},"tip")},danger:{infimaClassName:"danger",iconComponent:function(){return a.createElement("svg",{viewBox:"0 0 12 16"},a.createElement("path",{fillRule:"evenodd",d:"M5.05.31c.81 2.17.41 3.38-.52 4.31C3.55 5.67 1.98 6.45.9 7.98c-1.45 2.05-1.7 6.53 3.53 7.7-2.2-1.16-2.67-4.52-.3-6.61-.61 2.03.53 3.33 1.94 2.86 1.39-.47 2.3.53 2.27 1.67-.02.78-.31 1.44-1.13 1.81 3.42-.59 4.78-3.42 4.78-5.56 0-2.84-2.53-3.22-1.25-5.61-1.52.13-2.03 1.13-1.89 2.75.09 1.08-1.02 1.8-1.86 1.33-.67-.41-.66-1.19-.06-1.78C8.18 5.31 8.68 2.45 5.05.32L5.03.3l.02.01z"}))},label:a.createElement(s.Z,{id:"theme.admonition.danger",description:"The default label used for the Danger admonition (:::danger)"},"danger")},info:{infimaClassName:"info",iconComponent:function(){return a.createElement("svg",{viewBox:"0 0 14 16"},a.createElement("path",{fillRule:"evenodd",d:"M7 2.3c3.14 0 5.7 2.56 5.7 5.7s-2.56 5.7-5.7 5.7A5.71 5.71 0 0 1 1.3 8c0-3.14 2.56-5.7 5.7-5.7zM7 1C3.14 1 0 4.14 0 8s3.14 7 7 7 7-3.14 7-7-3.14-7-7-7zm1 3H6v5h2V4zm0 6H6v2h2v-2z"}))},label:a.createElement(s.Z,{id:"theme.admonition.info",description:"The default label used for the Info admonition (:::info)"},"info")},caution:{infimaClassName:"warning",iconComponent:function(){return a.createElement("svg",{viewBox:"0 0 16 16"},a.createElement("path",{fillRule:"evenodd",d:"M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"}))},label:a.createElement(s.Z,{id:"theme.admonition.caution",description:"The default label used for the Caution admonition (:::caution)"},"caution")}},d={secondary:"note",important:"info",success:"tip",warning:"danger"};function m(e){const{mdxAdmonitionTitle:t,rest:n}=function(e){const t=a.Children.toArray(e),n=t.find((e=>{var t;return a.isValidElement(e)&&"mdxAdmonitionTitle"===(null==(t=e.props)?void 0:t.mdxType)})),i=a.createElement(a.Fragment,null,t.filter((e=>e!==n)));return{mdxAdmonitionTitle:n,rest:i}}(e.children);return{...e,title:e.title??t,children:n}}function k(e){const{children:t,type:n,title:s,icon:k}=m(e),h=function(e){const t=d[e]??e;return p[t]||(console.warn(`No admonition config found for admonition type "${t}". Using Info as fallback.`),p.info)}(n),f=s??h.label,{iconComponent:g}=h,b=k??a.createElement(g,null);return a.createElement("div",{className:(0,i.Z)(r.k.common.admonition,r.k.common.admonitionType(e.type),"alert",`alert--${h.infimaClassName}`,o)},a.createElement("div",{className:l},a.createElement("span",{className:u},b),f),a.createElement("div",{className:c},t))}},85162:(e,t,n)=>{n.d(t,{Z:()=>s});var a=n(67294),i=n(86010);const r="tabItem_Ymn6";function s(e){let{children:t,hidden:n,className:s}=e;return a.createElement("div",{role:"tabpanel",className:(0,i.Z)(r,s),hidden:n},t)}},74866:(e,t,n)=>{n.d(t,{Z:()=>w});var a=n(87462),i=n(67294),r=n(86010),s=n(12466),o=n(76775),l=n(91980),u=n(67392),c=n(50012);function p(e){return function(e){var t;return(null==(t=i.Children.map(e,(e=>{if(!e||(0,i.isValidElement)(e)&&function(e){const{props:t}=e;return!!t&&"object"==typeof t&&"value"in t}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)})))?void 0:t.filter(Boolean))??[]}(e).map((e=>{let{props:{value:t,label:n,attributes:a,default:i}}=e;return{value:t,label:n,attributes:a,default:i}}))}function d(e){const{values:t,children:n}=e;return(0,i.useMemo)((()=>{const e=t??p(n);return function(e){const t=(0,u.l)(e,((e,t)=>e.value===t.value));if(t.length>0)throw new Error(`Docusaurus error: Duplicate values "${t.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[t,n])}function m(e){let{value:t,tabValues:n}=e;return n.some((e=>e.value===t))}function k(e){let{queryString:t=!1,groupId:n}=e;const a=(0,o.k6)(),r=function(e){let{queryString:t=!1,groupId:n}=e;if("string"==typeof t)return t;if(!1===t)return null;if(!0===t&&!n)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return n??null}({queryString:t,groupId:n});return[(0,l._X)(r),(0,i.useCallback)((e=>{if(!r)return;const t=new URLSearchParams(a.location.search);t.set(r,e),a.replace({...a.location,search:t.toString()})}),[r,a])]}function h(e){const{defaultValue:t,queryString:n=!1,groupId:a}=e,r=d(e),[s,o]=(0,i.useState)((()=>function(e){let{defaultValue:t,tabValues:n}=e;if(0===n.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(t){if(!m({value:t,tabValues:n}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${t}" but none of its children has the corresponding value. Available values are: ${n.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return t}const a=n.find((e=>e.default))??n[0];if(!a)throw new Error("Unexpected error: 0 tabValues");return a.value}({defaultValue:t,tabValues:r}))),[l,u]=k({queryString:n,groupId:a}),[p,h]=function(e){let{groupId:t}=e;const n=function(e){return e?`docusaurus.tab.${e}`:null}(t),[a,r]=(0,c.Nk)(n);return[a,(0,i.useCallback)((e=>{n&&r.set(e)}),[n,r])]}({groupId:a}),f=(()=>{const e=l??p;return m({value:e,tabValues:r})?e:null})();(0,i.useLayoutEffect)((()=>{f&&o(f)}),[f]);return{selectedValue:s,selectValue:(0,i.useCallback)((e=>{if(!m({value:e,tabValues:r}))throw new Error(`Can't select invalid tab value=${e}`);o(e),u(e),h(e)}),[u,h,r]),tabValues:r}}var f=n(72389);const g="tabList__CuJ",b="tabItem_LNqP";function y(e){let{className:t,block:n,selectedValue:o,selectValue:l,tabValues:u}=e;const c=[],{blockElementScrollPositionUntilNextRender:p}=(0,s.o5)(),d=e=>{const t=e.currentTarget,n=c.indexOf(t),a=u[n].value;a!==o&&(p(t),l(a))},m=e=>{var t;let n=null;switch(e.key){case"Enter":d(e);break;case"ArrowRight":{const t=c.indexOf(e.currentTarget)+1;n=c[t]??c[0];break}case"ArrowLeft":{const t=c.indexOf(e.currentTarget)-1;n=c[t]??c[c.length-1];break}}null==(t=n)||t.focus()};return i.createElement("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,r.Z)("tabs",{"tabs--block":n},t)},u.map((e=>{let{value:t,label:n,attributes:s}=e;return i.createElement("li",(0,a.Z)({role:"tab",tabIndex:o===t?0:-1,"aria-selected":o===t,key:t,ref:e=>c.push(e),onKeyDown:m,onClick:d},s,{className:(0,r.Z)("tabs__item",b,null==s?void 0:s.className,{"tabs__item--active":o===t})}),n??t)})))}function v(e){let{lazy:t,children:n,selectedValue:a}=e;const r=(Array.isArray(n)?n:[n]).filter(Boolean);if(t){const e=r.find((e=>e.props.value===a));return e?(0,i.cloneElement)(e,{className:"margin-top--md"}):null}return i.createElement("div",{className:"margin-top--md"},r.map(((e,t)=>(0,i.cloneElement)(e,{key:t,hidden:e.props.value!==a}))))}function N(e){const t=h(e);return i.createElement("div",{className:(0,r.Z)("tabs-container",g)},i.createElement(y,(0,a.Z)({},e,t)),i.createElement(v,(0,a.Z)({},e,t)))}function w(e){const t=(0,f.Z)();return i.createElement(N,(0,a.Z)({key:String(t)},e))}},23264:(e,t,n)=>{n.r(t),n.d(t,{ExecutorInfo:()=>m,assets:()=>p,contentTitle:()=>u,default:()=>h,frontMatter:()=>l,metadata:()=>c,toc:()=>d});var a=n(87462),i=(n(67294),n(3905)),r=n(74866),s=n(85162),o=n(23612);const l={},u="K6",c={unversionedId:"test-types/executor-k6",id:"test-types/executor-k6",title:"K6",description:"Testkube's k6 executor provides a convenient way of running k6 tests.",source:"@site/docs/test-types/executor-k6.mdx",sourceDirName:"test-types",slug:"/test-types/executor-k6",permalink:"/test-types/executor-k6",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/test-types/executor-k6.mdx",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"JMeter",permalink:"/test-types/executor-jmeter"},next:{title:"KubePug",permalink:"/test-types/executor-kubepug"}},p={},d=[{value:"Example k6 test",id:"example-k6-test",level:2},{value:"Test Source",id:"test-source",level:3},{value:"Creating and Running a Test",id:"creating-and-running-a-test",level:2},{value:"File",id:"file",level:3},{value:"Git File",id:"git-file",level:3},{value:"String",id:"string",level:3},{value:"File",id:"file-1",level:3},{value:"Git file",id:"git-file-1",level:3},{value:"Git Directory",id:"git-directory",level:3},{value:"Using Additional K6 Arguments in Your Tests",id:"using-additional-k6-arguments-in-your-tests",level:2},{value:"Git File",id:"git-file-2",level:4},{value:"Git Directory",id:"git-directory-1",level:4},{value:"String",id:"string-1",level:5},{value:"K6 Test Results",id:"k6-test-results",level:2}],m=()=>(0,i.kt)("div",null,(0,i.kt)(o.Z,{type:"info",icon:"\ud83c\udf93",title:"What is k6?",mdxType:"Admonition"},(0,i.kt)("ul",null,(0,i.kt)("li",null,(0,i.kt)("a",{href:"https://k6.io/docs/"},"k6")," is a free, developer-centric, and extensible open-source load testing tool that makes performance testing easy and productive for engineering teams.")),(0,i.kt)("b",null,"What can I do with k6?"),(0,i.kt)("ul",null,(0,i.kt)("li",null,"With k6, you can test the reliability and performance of your systems and catch performance regressions and problems earlier. K6 will help you to build resilient and performant applications that scale.")),"K6 is developed by Grafana Labs and the Open-Source community.")),k={toc:d,ExecutorInfo:m};function h(e){let{components:t,...o}=e;return(0,i.kt)("wrapper",(0,a.Z)({},k,o,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"k6"},"K6"),(0,i.kt)("p",null,"Testkube's k6 executor provides a convenient way of running k6 tests."),(0,i.kt)("p",null,"Default command for this executor: k6\nDefault arguments for this executor command: ","<","k6Command",">"," ","<","envVars",">"," ","<","runPath",">","\n(parameters in ","<",">"," are calculated at test execution)"),(0,i.kt)("iframe",{width:"100%",height:"315",src:"https://www.youtube.com/embed/e0NjGvGv_0c",title:"YouTube video \nplayer",frameborder:"0",allow:"accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; \npicture-in-picture; web-share",allowfullscreen:!0}),(0,i.kt)(m,{mdxType:"ExecutorInfo"}),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"Check out our ",(0,i.kt)("a",{parentName:"strong",href:"https://kubeshop.io/blog/load-testing-in-kubernetes-with-k6-and-testkube"},"blog post")," to follow tutorial steps to harness the power of k6 load testing in Kubernetes with Testkube's CLI and API.")),(0,i.kt)("h2",{id:"example-k6-test"},"Example k6 test"),(0,i.kt)("p",null,"In this example we will use the following k6 test:\n",(0,i.kt)("a",{parentName:"p",href:"https://github.com/kubeshop/testkube/blob/main/test/k6/executor-tests/k6-smoke-test-without-envs.js"},"https://github.com/kubeshop/testkube/blob/main/test/k6/executor-tests/k6-smoke-test-without-envs.js")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-js"},'import http from "k6/http";\n\nexport default function () {\n  http.get("https://testkube.kubeshop.io/");\n}\n')),(0,i.kt)("h3",{id:"test-source"},"Test Source"),(0,i.kt)("p",null,"K6 tests may vary significantly. The test may be just a single file, but may also consist of multiple files (modules, dependencies, or test data files). That's why all of the available Test Sources may be used with K6:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Git file"),(0,i.kt)("li",{parentName:"ul"},"Git directory"),(0,i.kt)("li",{parentName:"ul"},"File"),(0,i.kt)("li",{parentName:"ul"},"String")),(0,i.kt)("h2",{id:"creating-and-running-a-test"},"Creating and Running a Test"),(0,i.kt)(r.Z,{groupId:"dashboard-cli",mdxType:"Tabs"},(0,i.kt)(s.Z,{value:"dash",label:"Dashboard",mdxType:"TabItem"},(0,i.kt)("p",null,"If you prefer to use the Dashboard, just go to Tests, and click ",(0,i.kt)("inlineCode",{parentName:"p"},"Add a new test")," button. Then you need to fill in the test Name, choose the test Type (",(0,i.kt)("inlineCode",{parentName:"p"},"k6 script"),"), and then choose Test Source."),(0,i.kt)("h3",{id:"file"},"File"),(0,i.kt)("p",null,"If the source is ",(0,i.kt)("strong",{parentName:"p"},"File"),", the test file is uploaded directly."),(0,i.kt)("p",null,(0,i.kt)("img",{alt:"K6 test - creation dialog - file",src:n(65785).Z,width:"878",height:"680"})),(0,i.kt)("h3",{id:"git-file"},"Git File"),(0,i.kt)("p",null,"If the source is a ",(0,i.kt)("strong",{parentName:"p"},"Git file"),", you need to fill in repository details - Git repository URI (in this case ",(0,i.kt)("inlineCode",{parentName:"p"},"https://github.com/kubeshop/testkube.git"),"), branch (",(0,i.kt)("inlineCode",{parentName:"p"},"main"),"), and the path to k6 script in your repository (",(0,i.kt)("inlineCode",{parentName:"p"},"test/k6/executor-tests/k6-smoke-test-without-envs.js"),"). In this example, the repository is public, but in the case of private ones you would need to additionally fill in Git credentials."),(0,i.kt)("p",null,(0,i.kt)("img",{alt:"K6 test - creation dialog - git file",src:n(83581).Z,width:"876",height:"1072"})),(0,i.kt)("h3",{id:"string"},"String"),(0,i.kt)("p",null,"If the source is a ",(0,i.kt)("strong",{parentName:"p"},"String"),", the test script is added directly."),(0,i.kt)("p",null,(0,i.kt)("img",{alt:"K6 test - creation dialog - string",src:n(84532).Z,width:"876",height:"882"}))),(0,i.kt)(s.Z,{value:"cli",label:"CLI",mdxType:"TabItem"},"If you prefer using the CLI, you can create the test with `testkube create test`.",(0,i.kt)("p",null,"You need to set:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--name")," (for example, ",(0,i.kt)("inlineCode",{parentName:"li"},"k6-test"),")"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--type")," (in this case ",(0,i.kt)("inlineCode",{parentName:"li"},"k6/script"),")")),(0,i.kt)("p",null,"Then choose the Test Content type based on Test Source you want to use:"),(0,i.kt)("h3",{id:"file-1"},"File"),(0,i.kt)("p",null,"In the case of File test source:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--test-content-type")," (",(0,i.kt)("inlineCode",{parentName:"li"},"file-uri"),")"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--file")," (path to your k6 script - in this case ",(0,i.kt)("inlineCode",{parentName:"li"},"test/k6/executor-tests/k6-smoke-test-without-envs.js"),")")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-sh"},"testkube create test --name k6-test --type k6/script --test-content-type file-uri --file test/k6/executor-tests/k6-smoke-test-without-envs.js\n")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-sh",metastring:'title="Expected output:"',title:'"Expected','output:"':!0},"Test created testkube / k6-test \ud83e\udd47\n")),(0,i.kt)("h3",{id:"git-file-1"},"Git file"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--test-content-type")," (",(0,i.kt)("inlineCode",{parentName:"li"},"git-file"),", so specific file will be checked out from the Git repository)"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--git-uri")," - repository URI (in case of this example, ",(0,i.kt)("inlineCode",{parentName:"li"},"https://github.com/kubeshop/testkube.git"),")"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--git-branch")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--git-path")," - path to the k6 script in the repository (in this case ",(0,i.kt)("inlineCode",{parentName:"li"},"test/k6/executor-tests/k6-smoke-test-without-envs.js"),")")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-sh"},"testkube create test --name k6-test --type k6/script --test-content-type git-file --git-uri https://github.com/kubeshop/testkube.git --git-branch main --git-path test/k6/executor-tests/k6-smoke-test-without-envs.js\n")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-sh",metastring:'title="Expected output:"',title:'"Expected','output:"':!0},"Test created testkube / k6-test \ud83e\udd47\n")),(0,i.kt)("h3",{id:"git-directory"},"Git Directory"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--test-content-type")," (",(0,i.kt)("inlineCode",{parentName:"li"},"git-directory"),", so the whole directory will be checked out from the Git repository)"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--git-uri")," - repository URI (in case of this example, ",(0,i.kt)("inlineCode",{parentName:"li"},"https://github.com/kubeshop/testkube.git"),")"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--git-branch")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--git-path")," (path to the directory that should be checked out)"),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"--executor-args")," (whole directory will be checked out - specific test file must be set as k6 argument - in this example ",(0,i.kt)("inlineCode",{parentName:"li"},"test/k6/executor-tests/k6-smoke-test-without-envs.js"),")")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-sh"},"testkube create test --name k6-test --type k6/script --test-content-type git-dir --git-uri https://github.com/kubeshop/testkube.git --git-branch main --git-path test/k6/executor-tests --executor-args test/k6/executor-tests/k6-smoke-test-without-envs.js\n")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-sh",metastring:'title="Expected output:"',title:'"Expected','output:"':!0},"Test created testkube / k6-test \ud83e\udd47\n")),(0,i.kt)("h2",{id:"using-additional-k6-arguments-in-your-tests"},"Using Additional K6 Arguments in Your Tests"),(0,i.kt)("p",null,"Additional arguments can be passed to the ",(0,i.kt)("inlineCode",{parentName:"p"},"k6")," binary both on test creation (",(0,i.kt)("inlineCode",{parentName:"p"},"--executor-args"),"), and during test execution (",(0,i.kt)("inlineCode",{parentName:"p"},"--args"),")."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-sh"},"testkube run test -k6-test --args '--vus 100 --no-connection-reuse'\n"))),(0,i.kt)(s.Z,{value:"crd",label:"Custom Resource",mdxType:"TabItem"},(0,i.kt)("h4",{id:"git-file-2"},"Git File"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: tests.testkube.io/v3\nkind: Test\nmetadata:\n  name: k6-test\n  namespace: testkube\nspec:\n  type: k6/script\n  content:\n    type: git-file\n    repository:\n      type: git\n      uri: https://github.com/kubeshop/testkube.git\n      branch: main\n      path: test/k6/executor-tests/k6-smoke-test-without-envs.js\n")),(0,i.kt)("h4",{id:"git-directory-1"},"Git Directory"),(0,i.kt)("p",null,"Check out the entire Git directory (in the following example ",(0,i.kt)("inlineCode",{parentName:"p"},"test/k6/executor-tests"),"), and run a specific test file (",(0,i.kt)("inlineCode",{parentName:"p"},"test/k6/executor-tests/k6-smoke-test-without-envs.js"),"):"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: tests.testkube.io/v3\nkind: Test\nmetadata:\n  name: k6-test\n  namespace: testkube\nspec:\n  type: k6/script\n  content:\n    type: git-dir\n    repository:\n      type: git\n      uri: https://github.com/kubeshop/testkube.git\n      branch: main\n      path: test/k6/executor-tests\n  executionRequest:\n    args:\n      - test/k6/executor-tests/k6-smoke-test-without-envs.js\n")),(0,i.kt)("h5",{id:"string-1"},"String"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: tests.testkube.io/v3\nkind: Test\nmetadata:\n  name: k6-test\n  namespace: testkube\nspec:\n  type: k6/script\n  content:\n    type: string\n    data: \"import http from 'k6/http';\\n\\nexport default function () {\\n  http.get('https://testkube.kubeshop.io/');\\n}\"\n")))),(0,i.kt)("h2",{id:"k6-test-results"},"K6 Test Results"),(0,i.kt)("p",null,"A k6 test will be successful in Testkube when all checks and thresholds are successful. In the case of an error, the test will have ",(0,i.kt)("inlineCode",{parentName:"p"},"failed")," status, even if there is no failure in the summary report in the test logs. For details check ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/grafana/k6/issues/1680"},"this k6 issue"),"."))}h.isMDXComponent=!0},65785:(e,t,n)=>{n.d(t,{Z:()=>a});const a=n.p+"assets/images/dashboard-k6-create-test-file-82965b1d544c08a51c7f1994e5bfec27.png"},83581:(e,t,n)=>{n.d(t,{Z:()=>a});const a=n.p+"assets/images/dashboard-k6-create-test-git-file-cfbe7ee143c89452f12c077cd8f111f4.png"},84532:(e,t,n)=>{n.d(t,{Z:()=>a});const a=n.p+"assets/images/dashboard-k6-create-test-string-fccdf7090c09b924c5a738a53251cc0a.png"}}]);