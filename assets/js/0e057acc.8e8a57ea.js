"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[3858],{3905:(e,t,n)=>{n.d(t,{Zo:()=>u,kt:()=>m});var r=n(67294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var l=r.createContext({}),p=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},u=function(e){var t=p(e.components);return r.createElement(l.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,l=e.parentName,u=s(e,["components","mdxType","originalType","parentName"]),d=p(n),m=a,g=d["".concat(l,".").concat(m)]||d[m]||c[m]||o;return n?r.createElement(g,i(i({ref:t},u),{},{components:n})):r.createElement(g,i({ref:t},u))}));function m(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,i=new Array(o);i[0]=d;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:a,i[1]=s;for(var p=2;p<o;p++)i[p]=n[p];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},30099:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>i,default:()=>c,frontMatter:()=>o,metadata:()=>s,toc:()=>p});var r=n(87462),a=(n(67294),n(3905));const o={},i="Upgrade Testkube",s={unversionedId:"articles/upgrade",id:"articles/upgrade",title:"Upgrade Testkube",description:"Upgrading Testkube will upgrade the cluster components to the latest version. The following",source:"@site/docs/articles/upgrade.md",sourceDirName:"articles",slug:"/articles/upgrade",permalink:"/articles/upgrade",draft:!1,editUrl:"https://github.com/kubeshop/testkube/tree/develop/docs/docs/articles/upgrade.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Uninstall Testkube",permalink:"/articles/uninstall"},next:{title:"Test Types",permalink:"/category/test-types"}},l={},p=[{value:"Using Helm",id:"using-helm",level:2},{value:"Using Testkube&#39;s CLI",id:"using-testkubes-cli",level:2}],u={toc:p};function c(e){let{components:t,...n}=e;return(0,a.kt)("wrapper",(0,r.Z)({},u,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"upgrade-testkube"},"Upgrade Testkube"),(0,a.kt)("p",null,"Upgrading Testkube will upgrade the cluster components to the latest version. The following\napplies both to Open Source and Commercial installations."),(0,a.kt)("p",null,"There are two ways to upgrade Testkube: "),(0,a.kt)("h2",{id:"using-helm"},"Using Helm"),(0,a.kt)("admonition",{type:"note"},(0,a.kt)("p",{parentName:"admonition"},"By default, the namespace for the installation will be ",(0,a.kt)("inlineCode",{parentName:"p"},"testkube"),".\nTo upgrade the ",(0,a.kt)("inlineCode",{parentName:"p"},"testkube")," chart if it was installed into the default namespace:"),(0,a.kt)("pre",{parentName:"admonition"},(0,a.kt)("code",{parentName:"pre",className:"language-sh"},"helm upgrade my-testkube kubeshop/testkube\n")),(0,a.kt)("p",{parentName:"admonition"},"And for a namespace other than ",(0,a.kt)("inlineCode",{parentName:"p"},"testkube"),":"),(0,a.kt)("pre",{parentName:"admonition"},(0,a.kt)("code",{parentName:"pre",className:"language-sh"},"helm upgrade --namespace namespace_name my-testkube kubeshop/testkube\n"))),(0,a.kt)("h2",{id:"using-testkubes-cli"},"Using Testkube's CLI"),(0,a.kt)("p",null,"You can use the ",(0,a.kt)("inlineCode",{parentName:"p"},"upgrade")," command to upgrade your Testkube installation, see the\ncorresponding ",(0,a.kt)("a",{parentName:"p",href:"/cli/testkube_upgrade"},"CLI Documentation")," for all options."),(0,a.kt)("p",null,"Simple usage: "),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"testkube upgrade\n")))}c.isMDXComponent=!0}}]);