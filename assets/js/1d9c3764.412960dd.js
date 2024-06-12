"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[8601],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>m});var a=n(67294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function s(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?s(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):s(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},s=Object.keys(e);for(a=0;a<s.length;a++)n=s[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(a=0;a<s.length;a++)n=s[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var i=a.createContext({}),u=function(e){var t=a.useContext(i),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},p=function(e){var t=u(e.components);return a.createElement(i.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},d=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,s=e.originalType,i=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),d=u(n),m=r,f=d["".concat(i,".").concat(m)]||d[m]||c[m]||s;return n?a.createElement(f,o(o({ref:t},p),{},{components:n})):a.createElement(f,o({ref:t},p))}));function m(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var s=n.length,o=new Array(s);o[0]=d;var l={};for(var i in t)hasOwnProperty.call(t,i)&&(l[i]=t[i]);l.originalType=e,l.mdxType="string"==typeof e?e:r,o[1]=l;for(var u=2;u<s;u++)o[u]=n[u];return a.createElement.apply(null,o)}return a.createElement.apply(null,n)}d.displayName="MDXCreateElement"},85162:(e,t,n)=>{n.d(t,{Z:()=>o});var a=n(67294),r=n(86010);const s="tabItem_Ymn6";function o(e){let{children:t,hidden:n,className:o}=e;return a.createElement("div",{role:"tabpanel",className:(0,r.Z)(s,o),hidden:n},t)}},74866:(e,t,n)=>{n.d(t,{Z:()=>x});var a=n(87462),r=n(67294),s=n(86010),o=n(12466),l=n(76775),i=n(91980),u=n(67392),p=n(50012);function c(e){return function(e){var t;return(null==(t=r.Children.map(e,(e=>{if(!e||(0,r.isValidElement)(e)&&function(e){const{props:t}=e;return!!t&&"object"==typeof t&&"value"in t}(e))return e;throw new Error(`Docusaurus error: Bad <Tabs> child <${"string"==typeof e.type?e.type:e.type.name}>: all children of the <Tabs> component should be <TabItem>, and every <TabItem> should have a unique "value" prop.`)})))?void 0:t.filter(Boolean))??[]}(e).map((e=>{let{props:{value:t,label:n,attributes:a,default:r}}=e;return{value:t,label:n,attributes:a,default:r}}))}function d(e){const{values:t,children:n}=e;return(0,r.useMemo)((()=>{const e=t??c(n);return function(e){const t=(0,u.l)(e,((e,t)=>e.value===t.value));if(t.length>0)throw new Error(`Docusaurus error: Duplicate values "${t.map((e=>e.value)).join(", ")}" found in <Tabs>. Every value needs to be unique.`)}(e),e}),[t,n])}function m(e){let{value:t,tabValues:n}=e;return n.some((e=>e.value===t))}function f(e){let{queryString:t=!1,groupId:n}=e;const a=(0,l.k6)(),s=function(e){let{queryString:t=!1,groupId:n}=e;if("string"==typeof t)return t;if(!1===t)return null;if(!0===t&&!n)throw new Error('Docusaurus error: The <Tabs> component groupId prop is required if queryString=true, because this value is used as the search param name. You can also provide an explicit value such as queryString="my-search-param".');return n??null}({queryString:t,groupId:n});return[(0,i._X)(s),(0,r.useCallback)((e=>{if(!s)return;const t=new URLSearchParams(a.location.search);t.set(s,e),a.replace({...a.location,search:t.toString()})}),[s,a])]}function k(e){const{defaultValue:t,queryString:n=!1,groupId:a}=e,s=d(e),[o,l]=(0,r.useState)((()=>function(e){let{defaultValue:t,tabValues:n}=e;if(0===n.length)throw new Error("Docusaurus error: the <Tabs> component requires at least one <TabItem> children component");if(t){if(!m({value:t,tabValues:n}))throw new Error(`Docusaurus error: The <Tabs> has a defaultValue "${t}" but none of its children has the corresponding value. Available values are: ${n.map((e=>e.value)).join(", ")}. If you intend to show no default tab, use defaultValue={null} instead.`);return t}const a=n.find((e=>e.default))??n[0];if(!a)throw new Error("Unexpected error: 0 tabValues");return a.value}({defaultValue:t,tabValues:s}))),[i,u]=f({queryString:n,groupId:a}),[c,k]=function(e){let{groupId:t}=e;const n=function(e){return e?`docusaurus.tab.${e}`:null}(t),[a,s]=(0,p.Nk)(n);return[a,(0,r.useCallback)((e=>{n&&s.set(e)}),[n,s])]}({groupId:a}),h=(()=>{const e=i??c;return m({value:e,tabValues:s})?e:null})();(0,r.useLayoutEffect)((()=>{h&&l(h)}),[h]);return{selectedValue:o,selectValue:(0,r.useCallback)((e=>{if(!m({value:e,tabValues:s}))throw new Error(`Can't select invalid tab value=${e}`);l(e),u(e),k(e)}),[u,k,s]),tabValues:s}}var h=n(72389);const b="tabList__CuJ",g="tabItem_LNqP";function w(e){let{className:t,block:n,selectedValue:l,selectValue:i,tabValues:u}=e;const p=[],{blockElementScrollPositionUntilNextRender:c}=(0,o.o5)(),d=e=>{const t=e.currentTarget,n=p.indexOf(t),a=u[n].value;a!==l&&(c(t),i(a))},m=e=>{var t;let n=null;switch(e.key){case"Enter":d(e);break;case"ArrowRight":{const t=p.indexOf(e.currentTarget)+1;n=p[t]??p[0];break}case"ArrowLeft":{const t=p.indexOf(e.currentTarget)-1;n=p[t]??p[p.length-1];break}}null==(t=n)||t.focus()};return r.createElement("ul",{role:"tablist","aria-orientation":"horizontal",className:(0,s.Z)("tabs",{"tabs--block":n},t)},u.map((e=>{let{value:t,label:n,attributes:o}=e;return r.createElement("li",(0,a.Z)({role:"tab",tabIndex:l===t?0:-1,"aria-selected":l===t,key:t,ref:e=>p.push(e),onKeyDown:m,onClick:d},o,{className:(0,s.Z)("tabs__item",g,null==o?void 0:o.className,{"tabs__item--active":l===t})}),n??t)})))}function v(e){let{lazy:t,children:n,selectedValue:a}=e;const s=(Array.isArray(n)?n:[n]).filter(Boolean);if(t){const e=s.find((e=>e.props.value===a));return e?(0,r.cloneElement)(e,{className:"margin-top--md"}):null}return r.createElement("div",{className:"margin-top--md"},s.map(((e,t)=>(0,r.cloneElement)(e,{key:t,hidden:e.props.value!==a}))))}function y(e){const t=k(e);return r.createElement("div",{className:(0,s.Z)("tabs-container",b)},r.createElement(w,(0,a.Z)({},e,t)),r.createElement(v,(0,a.Z)({},e,t)))}function x(e){const t=(0,h.Z)();return r.createElement(y,(0,a.Z)({key:String(t)},e))}},52508:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>p,contentTitle:()=>i,default:()=>m,frontMatter:()=>l,metadata:()=>u,toc:()=>c});var a=n(87462),r=(n(67294),n(3905)),s=n(74866),o=n(85162);const l={},i="Test Workflows - Test Suites",u={unversionedId:"articles/test-workflows-test-suites",id:"articles/test-workflows-test-suites",title:"Test Workflows - Test Suites",description:"With Test Workflows it is possible to run downstream Test Workflows and Tests with execute operation,",source:"@site/docs/articles/test-workflows-test-suites.md",sourceDirName:"articles",slug:"/articles/test-workflows-test-suites",permalink:"/articles/test-workflows-test-suites",draft:!1,editUrl:"https://github.com/kubeshop/testkube/tree/develop/docs/docs/articles/test-workflows-test-suites.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Test Workflows Examples - Templates",permalink:"/articles/test-workflows-examples-templates"},next:{title:"Test Workflows - Parallel Steps",permalink:"/articles/test-workflows-parallel"}},p={},c=[{value:"Advantages over original Test Suite",id:"advantages-over-original-test-suite",level:2},{value:"Syntax",id:"syntax",level:2},{value:"Running Test Workflows",id:"running-test-workflows",level:3},{value:"Running Tests",id:"running-tests",level:3},{value:"Controlling the concurrency level",id:"controlling-the-concurrency-level",level:3},{value:"Passing input from files",id:"passing-input-from-files",level:2},{value:"Specific files",id:"specific-files",level:3},{value:"Multiple files transfer",id:"multiple-files-transfer",level:3},{value:"Matrix and sharding",id:"matrix-and-sharding",level:3}],d={toc:c};function m(e){let{components:t,...l}=e;return(0,r.kt)("wrapper",(0,a.Z)({},d,l,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"test-workflows---test-suites"},"Test Workflows - Test Suites"),(0,r.kt)("p",null,"With Test Workflows it is possible to run downstream Test Workflows and Tests with ",(0,r.kt)("inlineCode",{parentName:"p"},"execute")," operation,\nsimilarly to what you can do in Test Suites."),(0,r.kt)("h2",{id:"advantages-over-original-test-suite"},"Advantages over original Test Suite"),(0,r.kt)("admonition",{type:"tip"},(0,r.kt)("p",{parentName:"admonition"},"We consider Test Workflows as a long-term solution, so keep in mind that the original Test Suites will ",(0,r.kt)("a",{parentName:"p",href:"https://testkube.io/blog/the-future-of-testkube-with-test-workflows"},(0,r.kt)("strong",{parentName:"a"},"become deprecated")),".")),(0,r.kt)("p",null,"As it is regular Test Workflow, where a single step is dispatching downstream Test Workflows and Tests,\nthe execution is very flexible. You can:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Fetch input data before (i.e. by using ",(0,r.kt)("inlineCode",{parentName:"li"},"curl"),"/",(0,r.kt)("inlineCode",{parentName:"li"},"wget")," to download data, or fetching Git repository)"),(0,r.kt)("li",{parentName:"ul"},"Run setup operations (i.e. start shared instance of database, or generate API key)"),(0,r.kt)("li",{parentName:"ul"},"Process the results (i.e. by notifying about the status)"),(0,r.kt)("li",{parentName:"ul"},"Run other tests based on the previous results")),(0,r.kt)("h2",{id:"syntax"},"Syntax"),(0,r.kt)("p",null,"You have to use ",(0,r.kt)("inlineCode",{parentName:"p"},"execute")," operation in the step, and provide definition of the Test Workflows and Tests to run."),(0,r.kt)(s.Z,{mdxType:"Tabs"},(0,r.kt)(o.Z,{value:"yaml",label:"YAML",default:!0,mdxType:"TabItem"},(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: testworkflows.testkube.io/v1\nkind: TestWorkflow\nmetadata:\n  name: example-test-suite\nspec:\n  steps:\n  - execute:\n      workflows:\n      - name: example-distributed-k6\n        description: Run {{ index + 1 }} of {{ count }}\n        count: 2\n        config:\n          vus: 8\n          duration: 1s\n          workers: 2\n      - name: example-sharded-cypress\n      tests:\n      - name: example-test\n        description: Example without request\n      - name: example-test\n        description: Example with env variables\n        executionRequest:\n          variables:\n            SOME_VARIABLE:\n              type: basic\n              name: SOME_VARIABLE\n              value: some-value\n"))),(0,r.kt)(o.Z,{value:"log",label:"Log Output",mdxType:"TabItem"},(0,r.kt)("p",null,(0,r.kt)("img",{alt:"example-test-suite.png",src:n(64425).Z,width:"2158",height:"1572"})))),(0,r.kt)("h3",{id:"running-test-workflows"},"Running Test Workflows"),(0,r.kt)("p",null,"To run Test Workflow as part of the ",(0,r.kt)("inlineCode",{parentName:"p"},"execute")," step, you have to add its reference in the ",(0,r.kt)("inlineCode",{parentName:"p"},"workflows")," list."),(0,r.kt)("p",null,"You need to provide ",(0,r.kt)("inlineCode",{parentName:"p"},"name"),", along with optional ",(0,r.kt)("inlineCode",{parentName:"p"},"config")," values for parametrization."),(0,r.kt)("h3",{id:"running-tests"},"Running Tests"),(0,r.kt)("admonition",{type:"tip"},(0,r.kt)("p",{parentName:"admonition"},"We consider Test Workflows as a long-term solution, so keep in mind that the Tests will ",(0,r.kt)("a",{parentName:"p",href:"https://testkube.io/blog/the-future-of-testkube-with-test-workflows"},(0,r.kt)("strong",{parentName:"a"},"become deprecated")),".")),(0,r.kt)("p",null,"To run Tests as part of the ",(0,r.kt)("inlineCode",{parentName:"p"},"execute")," step, you have to add its reference in the ",(0,r.kt)("inlineCode",{parentName:"p"},"tests")," list."),(0,r.kt)("p",null,"You need to provide ",(0,r.kt)("inlineCode",{parentName:"p"},"name"),", along with optional ",(0,r.kt)("inlineCode",{parentName:"p"},"executionRequest")," values for parametrization,\nthat are similar to the regular Test execution request."),(0,r.kt)("h3",{id:"controlling-the-concurrency-level"},"Controlling the concurrency level"),(0,r.kt)("p",null,"You can use ",(0,r.kt)("inlineCode",{parentName:"p"},"parallelism")," property to control how many Test Workflows and Tests will be running at once."),(0,r.kt)("p",null,"In example, to run all the downstream jobs sequentially, you can use ",(0,r.kt)("inlineCode",{parentName:"p"},"parallelism: 1"),".\nIt affects jobs instantiated by ",(0,r.kt)("a",{parentName:"p",href:"/articles/test-workflows-matrix-and-sharding"},(0,r.kt)("strong",{parentName:"a"},"matrix and sharding"))," properties (like ",(0,r.kt)("inlineCode",{parentName:"p"},"count"),") too."),(0,r.kt)(s.Z,{mdxType:"Tabs"},(0,r.kt)(o.Z,{value:"yaml",label:"YAML",default:!0,mdxType:"TabItem"},(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: testworkflows.testkube.io/v1\nkind: TestWorkflow\nmetadata:\n  name: example-sequential-test-suite\nspec:\n  steps:\n  - execute:\n      parallelism: 1\n      workflows:\n      - name: example-distributed-k6\n        count: 2\n        config:\n          vus: 8\n          duration: 1s\n          workers: 2\n      - name: example-sharded-cypress\n      tests:\n      - name: example-test\n        count: 5\n"))),(0,r.kt)(o.Z,{value:"log",label:"Log Output",mdxType:"TabItem"},(0,r.kt)("p",null,(0,r.kt)("img",{alt:"example-sequential-test-suite.png",src:n(18770).Z,width:"2158",height:"1580"})))),(0,r.kt)("h2",{id:"passing-input-from-files"},"Passing input from files"),(0,r.kt)("p",null,"It may happen that you will need to pass information from the file system. You can either pass the files using Test Workflow expressions (like ",(0,r.kt)("inlineCode",{parentName:"p"},'file("./file-content.txt")'),") or using a ",(0,r.kt)("inlineCode",{parentName:"p"},"tarball")," syntax."),(0,r.kt)("h3",{id:"specific-files"},"Specific files"),(0,r.kt)("p",null,"You can easily use Test Workflow expressions to fetch some files and send them as a configuration variable:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: testworkflows.testkube.io/v1\nkind: TestWorkflow\nmetadata:\n  name: example-test-suite-with-file-input\nspec:\n  content:\n    git:\n      uri: https://github.com/kubeshop/testkube\n      revision: main\n      paths:\n      - test/k6/executor-tests/k6-smoke-test-without-envs.js\n  steps:\n  - execute:\n      workflows:\n      - name: example-distributed-k6\n        config:\n          vus: 8\n          duration: 1s\n          workers: 2\n          script: '{{ file(\"/data/repo/test/k6/executor-tests/k6-smoke-test-without-envs.js\") }}'\n")),(0,r.kt)("h3",{id:"multiple-files-transfer"},"Multiple files transfer"),(0,r.kt)("p",null,"To transfer multiple files, similarly to ",(0,r.kt)("inlineCode",{parentName:"p"},"transfer")," in ",(0,r.kt)("a",{parentName:"p",href:"/articles/test-workflows-parallel#copying-content-inside"},(0,r.kt)("strong",{parentName:"a"},"Parallel Steps")),",\nyou can use a ",(0,r.kt)("inlineCode",{parentName:"p"},"tarball")," syntax that will pack selected files and return the URL to download them:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: testworkflows.testkube.io/v1\nkind: TestWorkflow\nmetadata:\n  name: example-test-suite-with-file-input-packaged\nspec:\n  content:\n    git:\n      uri: https://github.com/kubeshop/testkube\n      revision: main\n      paths:\n      - test/k6/executor-tests/k6-smoke-test-without-envs.js\n  steps:\n  - execute:\n      workflows:\n      - name: example-test-reading-files\n        tarball:\n          scripts:\n            from: /data/repo\n        config:\n          input: '{{ tarball.scripts.url }}'\n")),(0,r.kt)("p",null,"You can later use i.e. ",(0,r.kt)("inlineCode",{parentName:"p"},"content.tarball")," to unpack them in destination test:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'apiVersion: testworkflows.testkube.io/v1\nkind: TestWorkflow\nmetadata:\n  name: example-test-reading-files\nspec:\n  config:\n    input: {type: string}\n  content:\n    tarball:\n    - url: "{{ config.input }}" # extract provided tarball\n      path: "/data/repo"        # to local /data/repo directory (or any other)\n  steps:\n  - shell: tree /data/repo\n')),(0,r.kt)("h3",{id:"matrix-and-sharding"},"Matrix and sharding"),(0,r.kt)("p",null,"The ",(0,r.kt)("inlineCode",{parentName:"p"},"execute")," operation supports matrix and sharding, to run multiple replicas and/or distribute the load across multiple runs.\nIt is supported by regular matrix/sharding properties (",(0,r.kt)("inlineCode",{parentName:"p"},"matrix"),", ",(0,r.kt)("inlineCode",{parentName:"p"},"shards"),", ",(0,r.kt)("inlineCode",{parentName:"p"},"count")," and ",(0,r.kt)("inlineCode",{parentName:"p"},"maxCount"),") for each Test Workflow or Test reference."),(0,r.kt)("p",null,"You can read more about it in the general ",(0,r.kt)("a",{parentName:"p",href:"/articles/test-workflows-matrix-and-sharding"},(0,r.kt)("strong",{parentName:"a"},"Matrix and Sharding"))," documentation."))}m.isMDXComponent=!0},18770:(e,t,n)=>{n.d(t,{Z:()=>a});const a=n.p+"assets/images/example-sequential-test-suite-28022f0b5b062f77b2a8af72f4aab148.png"},64425:(e,t,n)=>{n.d(t,{Z:()=>a});const a=n.p+"assets/images/example-test-suite-b62cc83973bceaf6e8d56f05b6af3ecb.png"}}]);