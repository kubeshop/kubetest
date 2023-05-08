"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[8718],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>b});var r=n(67294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function s(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?s(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):s(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},s=Object.keys(e);for(r=0;r<s.length;r++)n=s[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(r=0;r<s.length;r++)n=s[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var l=r.createContext({}),u=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},c=function(e){var t=u(e.components);return r.createElement(l.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,s=e.originalType,l=e.parentName,c=i(e,["components","mdxType","originalType","parentName"]),d=u(n),b=o,g=d["".concat(l,".").concat(b)]||d[b]||p[b]||s;return n?r.createElement(g,a(a({ref:t},c),{},{components:n})):r.createElement(g,a({ref:t},c))}));function b(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var s=n.length,a=new Array(s);a[0]=d;var i={};for(var l in t)hasOwnProperty.call(t,l)&&(i[l]=t[l]);i.originalType=e,i.mdxType="string"==typeof e?e:o,a[1]=i;for(var u=2;u<s;u++)a[u]=n[u];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},62178:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>a,default:()=>p,frontMatter:()=>s,metadata:()=>i,toc:()=>u});var r=n(87462),o=(n(67294),n(3905));const s={},a="Intro",i={unversionedId:"testkube-cloud/articles/intro",id:"testkube-cloud/articles/intro",title:"Intro",description:"Testkube Cloud is the managed version of Testkube with the main purpose of orchestrating multiple clusters.",source:"@site/docs/testkube-cloud/articles/intro.md",sourceDirName:"testkube-cloud/articles",slug:"/testkube-cloud/articles/intro",permalink:"/testkube-cloud/articles/intro",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/testkube-cloud/articles/intro.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Container Executor",permalink:"/test-types/container-executor"},next:{title:"Installing the Testkube Agent",permalink:"/testkube-cloud/articles/installing-agent"}},l={},u=[{value:"Testkube Cloud Agent - Installation Manual",id:"testkube-cloud-agent---installation-manual",level:2},{value:"Installing the Agent",id:"installing-the-agent",level:2}],c={toc:u};function p(e){let{components:t,...n}=e;return(0,o.kt)("wrapper",(0,r.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"intro"},"Intro"),(0,o.kt)("p",null,"Testkube Cloud is the managed version of Testkube with the main purpose of orchestrating multiple clusters.\nAll test results and test artifacts are stored into Testkube Cloud internal data storages. Testkube cloud\nwill provide you with additional tests insights and is able to limit access for your users only to a subset\nof environments."),(0,o.kt)("p",null,"Testkube Cloud is in Alpha phase - so feel free to give us feedback! "),(0,o.kt)("h2",{id:"testkube-cloud-agent---installation-manual"},"Testkube Cloud Agent - Installation Manual"),(0,o.kt)("p",null,"Testkube Cloud is able to connect to Testkube Agents. Testkube Agent is the Testkube engine for managing test runs into your cluster.\nIt is also responsible for getting insight into Testkube resources stored in the cluster."),(0,o.kt)("p",null,"Testkube Agent opens a networking connection into Testkube Cloud API, which is always active with the main purpose of listening for Testkube Cloud commands."),(0,o.kt)("p",null,"Your existing Open Source Testkube installation can be converted into a Testkube Cloud agent; data will be passed and managed by\nTestkube servers (Coming Soon!)"),(0,o.kt)("h2",{id:"installing-the-agent"},"Installing the Agent"),(0,o.kt)("p",null,"Please follow the ",(0,o.kt)("a",{parentName:"p",href:"/testkube-cloud/articles/installing-agent"},"install steps")," to get started using the Testkube Agent."))}p.isMDXComponent=!0}}]);