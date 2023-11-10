"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[578],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>k});var r=n(67294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function l(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var s=r.createContext({}),u=function(e){var t=r.useContext(s),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},c=function(e){var t=u(e.components);return r.createElement(s.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,i=e.mdxType,a=e.originalType,s=e.parentName,c=l(e,["components","mdxType","originalType","parentName"]),d=u(n),k=i,f=d["".concat(s,".").concat(k)]||d[k]||p[k]||a;return n?r.createElement(f,o(o({ref:t},c),{},{components:n})):r.createElement(f,o({ref:t},c))}));function k(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var a=n.length,o=new Array(a);o[0]=d;var l={};for(var s in t)hasOwnProperty.call(t,s)&&(l[s]=t[s]);l.originalType=e,l.mdxType="string"==typeof e?e:i,o[1]=l;for(var u=2;u<a;u++)o[u]=n[u];return r.createElement.apply(null,o)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},73578:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>s,contentTitle:()=>o,default:()=>p,frontMatter:()=>a,metadata:()=>l,toc:()=>u});var r=n(87462),i=(n(67294),n(3905));const a={},o="Configure Identity Providers",l={unversionedId:"testkube-enterprise/articles/auth",id:"testkube-enterprise/articles/auth",title:"Configure Identity Providers",description:"You can configure Testkube Enterprise to authenticate users using different identity providers such as Azure AD, Google, Okta, and OIDC. To do this, you need to update the additionalConfig field in the Helm chart values with the appropriate configuration for each identity provider.",source:"@site/docs/testkube-enterprise/articles/auth.md",sourceDirName:"testkube-enterprise/articles",slug:"/testkube-enterprise/articles/auth",permalink:"/testkube-enterprise/articles/auth",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/testkube-enterprise/articles/auth.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Helm Chart Installation and Usage Guide",permalink:"/testkube-enterprise/articles/usage-guide"},next:{title:"Testkube Open Source",permalink:"/articles/testkube-oss"}},s={},u=[{value:"Quickstart",id:"quickstart",level:3},{value:"Static Users",id:"static-users",level:3},{value:"Azure AD",id:"azure-ad",level:3},{value:"Google",id:"google",level:3},{value:"Okta",id:"okta",level:3},{value:"OIDC",id:"oidc",level:3}],c={toc:u};function p(e){let{components:t,...n}=e;return(0,i.kt)("wrapper",(0,r.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"configure-identity-providers"},"Configure Identity Providers"),(0,i.kt)("p",null,"You can configure Testkube Enterprise to authenticate users using different identity providers such as Azure AD, Google, Okta, and OIDC. To do this, you need to update the ",(0,i.kt)("inlineCode",{parentName:"p"},"additionalConfig")," field in the Helm chart values with the appropriate configuration for each identity provider."),(0,i.kt)("p",null,"For a list of all supported identity providers, see ",(0,i.kt)("a",{parentName:"p",href:"https://dexidp.io/docs/connectors/"},"Connectors"),"."),(0,i.kt)("p",null,"The examples below show how to configure Testkube Enterprise with each identity provider by editing the ",(0,i.kt)("inlineCode",{parentName:"p"},"dex.configTemplate.additionalConfig")," field in the Helm chart values."),(0,i.kt)("h3",{id:"quickstart"},"Quickstart"),(0,i.kt)("p",null,"For a quickstart, or if you do not have an identity provider, you can configure Testkube Enterprise to use static users.\nSee ",(0,i.kt)("a",{parentName:"p",href:"#static-users"},"Static Users"),"."),(0,i.kt)("h3",{id:"static-users"},"Static Users"),(0,i.kt)("p",null,"To configure Testkube Enterprise with static users, add the following configuration to the ",(0,i.kt)("inlineCode",{parentName:"p"},"additionalConfig")," field:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"additionalConfig: |\n  enablePasswordDB: true\n  staticPasswords:\n    - email: <user email>\n      hash: <bcrypt hash of user password>\n      username: <username>\n")),(0,i.kt)("p",null,"Replace ",(0,i.kt)("inlineCode",{parentName:"p"},"<user email>"),", ",(0,i.kt)("inlineCode",{parentName:"p"},"<bcrypt hash of user password>"),", and ",(0,i.kt)("inlineCode",{parentName:"p"},"<username>")," with your actual values."),(0,i.kt)("h3",{id:"azure-ad"},"Azure AD"),(0,i.kt)("p",null,"To configure Testkube Enterprise with Azure AD, add the following configuration to the ",(0,i.kt)("inlineCode",{parentName:"p"},"additionalConfig")," field:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"additionalConfig: |\n  connectors:\n    - type: azuread\n      id: azuread\n      name: Azure AD\n      config:\n        clientID: <Azure AD client ID>\n        clientSecret: <Azure AD client secret>\n        redirectURI: <Testkube Enterprise redirect URI>\n")),(0,i.kt)("p",null,"Replace ",(0,i.kt)("inlineCode",{parentName:"p"},"Azure AD client ID"),", ",(0,i.kt)("inlineCode",{parentName:"p"},"Azure AD client secret"),", and ",(0,i.kt)("inlineCode",{parentName:"p"},"Testkube Enterprise redirect URI")," with your actual Azure AD configuration values."),(0,i.kt)("h3",{id:"google"},"Google"),(0,i.kt)("p",null,"To configure Testkube Enterprise with Google, add the following configuration to the 'additionalConfig' field:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"additionalConfig: |\n  connectors:\n    - type: google\n      id: google\n      name: Google\n      config:\n        clientID: <Google client ID>\n        clientSecret: <Google client secret>\n        redirectURI: <Testkube Enterprise redirect URI>\n")),(0,i.kt)("p",null,"Replace ",(0,i.kt)("inlineCode",{parentName:"p"},"Google client ID"),", ",(0,i.kt)("inlineCode",{parentName:"p"},"Google client secret"),", and ",(0,i.kt)("inlineCode",{parentName:"p"},"Testkube Enterprise redirect URI")," with your actual Google configuration values."),(0,i.kt)("h3",{id:"okta"},"Okta"),(0,i.kt)("p",null,"To configure Testkube Enterprise with Okta, add the following configuration to the ",(0,i.kt)("inlineCode",{parentName:"p"},"additionalConfig")," field:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"additionalConfig: |\n  connectors:\n    - type: okta\n      id: okta\n      name: Okta\n      config:\n        issuerURL: <Okta issuer URL>\n        clientID: <Okta client ID>\n        clientSecret: <Okta client secret>\n        redirectURI: <Testkube Enterprise redirect URI>\n")),(0,i.kt)("p",null,"Replace ",(0,i.kt)("inlineCode",{parentName:"p"},"Okta issuer URL"),", ",(0,i.kt)("inlineCode",{parentName:"p"},"Okta client ID"),", ",(0,i.kt)("inlineCode",{parentName:"p"},"Okta client secret"),", and ",(0,i.kt)("inlineCode",{parentName:"p"},"Testkube Enterprise redirect URI")," with your actual Okta configuration values."),(0,i.kt)("h3",{id:"oidc"},"OIDC"),(0,i.kt)("p",null,"To configure Testkube Enterprise with an OIDC provider, add the following configuration to the ",(0,i.kt)("inlineCode",{parentName:"p"},"additionalConfig")," field:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-yaml"},"additionalConfig: |\n  connectors:\n    - type: oidc\n      id: oidc\n      name: OIDC\n      config:\n        issuerURL: <OIDC issuer URL>\n        clientID: <OIDC client ID>\n        clientSecret: <OIDC client secret>\n        redirectURI: <Testkube Enterprise redirect URI>\n")))}p.isMDXComponent=!0}}]);