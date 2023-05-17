"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[7588],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>d});var r=n(67294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},i=Object.keys(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(r=0;r<i.length;r++)n=i[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var u=r.createContext({}),l=function(e){var t=r.useContext(u),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},c=function(e){var t=l(e.components);return r.createElement(u.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},m=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,i=e.originalType,u=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),m=l(n),d=o,f=m["".concat(u,".").concat(d)]||m[d]||p[d]||i;return n?r.createElement(f,a(a({ref:t},c),{},{components:n})):r.createElement(f,a({ref:t},c))}));function d(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var i=n.length,a=new Array(i);a[0]=m;var s={};for(var u in t)hasOwnProperty.call(t,u)&&(s[u]=t[u]);s.originalType=e,s.mdxType="string"==typeof e?e:o,a[1]=s;for(var l=2;l<i;l++)a[l]=n[l];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}m.displayName="MDXCreateElement"},44873:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>u,contentTitle:()=>a,default:()=>p,frontMatter:()=>i,metadata:()=>s,toc:()=>l});var r=n(87462),o=(n(67294),n(3905));const i={},a="Migrating from Testkube Open Source",s={unversionedId:"testkube-cloud/articles/transition-from-oss",id:"testkube-cloud/articles/transition-from-oss",title:"Migrating from Testkube Open Source",description:"If you have started using Testkube using the Open Source installation, you can migrate this instace to be managed using Testkube Cloud.",source:"@site/docs/testkube-cloud/articles/transition-from-oss.md",sourceDirName:"testkube-cloud/articles",slug:"/testkube-cloud/articles/transition-from-oss",permalink:"/testkube-cloud/articles/transition-from-oss",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/testkube-cloud/articles/transition-from-oss.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Installing the Testkube Agent",permalink:"/testkube-cloud/articles/installing-agent"},next:{title:"Organizations management",permalink:"/testkube-cloud/articles/organization-management"}},u={},l=[{value:"Instructions",id:"instructions",level:2}],c={toc:l};function p(e){let{components:t,...n}=e;return(0,o.kt)("wrapper",(0,r.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"migrating-from-testkube-open-source"},"Migrating from Testkube Open Source"),(0,o.kt)("p",null,"If you have started using Testkube using the Open Source installation, you can migrate this instace to be managed using Testkube Cloud. "),(0,o.kt)("p",null,"To connect your Testkube Open Source instance you will need to modify your Testkube installation to be in Cloud Agent mode. Testkube Cloud Agent is the Testkube engine for controlling your Testkube instance using the managed solution. It sends data to Testkube's Cloud Servers."),(0,o.kt)("p",null,"::: note"),(0,o.kt)("p",null,"Currently we are supporting uploading existing test logs and artifacts from your Testkube Open Source instance. This is planned for coming releases.\ns\n::: "),(0,o.kt)("h2",{id:"instructions"},"Instructions"),(0,o.kt)("ol",null,(0,o.kt)("li",{parentName:"ol"},"Run the following command which will walk you through the migration process:")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-sh"},"testkube cloud connect\n")),(0,o.kt)("ol",{start:2},(0,o.kt)("li",{parentName:"ol"},(0,o.kt)("a",{parentName:"li",href:"/testkube-cloud/articles/managing-cli-context"},"Set your CLI Context to talk to Testkube Cloud"))))}p.isMDXComponent=!0}}]);