"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[9137],{3905:(e,t,r)=>{r.d(t,{Zo:()=>u,kt:()=>d});var n=r(67294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function s(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?s(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):s(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function i(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},s=Object.keys(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(n=0;n<s.length;n++)r=s[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var c=n.createContext({}),l=function(e){var t=n.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},u=function(e){var t=l(e.components);return n.createElement(c.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},m=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,s=e.originalType,c=e.parentName,u=i(e,["components","mdxType","originalType","parentName"]),m=l(r),d=a,f=m["".concat(c,".").concat(d)]||m[d]||p[d]||s;return r?n.createElement(f,o(o({ref:t},u),{},{components:r})):n.createElement(f,o({ref:t},u))}));function d(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var s=r.length,o=new Array(s);o[0]=m;var i={};for(var c in t)hasOwnProperty.call(t,c)&&(i[c]=t[c]);i.originalType=e,i.mdxType="string"==typeof e?e:a,o[1]=i;for(var l=2;l<s;l++)o[l]=r[l];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}m.displayName="MDXCreateElement"},51631:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>c,contentTitle:()=>o,default:()=>p,frontMatter:()=>s,metadata:()=>i,toc:()=>l});var n=r(87462),a=(r(67294),r(3905));const s={},o=void 0,i={unversionedId:"cli/testkube_generate_tests-crds",id:"cli/testkube_generate_tests-crds",title:"testkube_generate_tests-crds",description:"testkube generate tests-crds",source:"@site/docs/cli/testkube_generate_tests-crds.md",sourceDirName:"cli",slug:"/cli/testkube_generate_tests-crds",permalink:"/cli/testkube_generate_tests-crds",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/cli/testkube_generate_tests-crds.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"testkube_generate_doc",permalink:"/cli/testkube_generate_doc"},next:{title:"testkube_get",permalink:"/cli/testkube_get"}},c={},l=[{value:"testkube generate tests-crds",id:"testkube-generate-tests-crds",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],u={toc:l};function p(e){let{components:t,...r}=e;return(0,a.kt)("wrapper",(0,n.Z)({},u,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h2",{id:"testkube-generate-tests-crds"},"testkube generate tests-crds"),(0,a.kt)("p",null,"Generate tests CRD file based on directory"),(0,a.kt)("h3",{id:"synopsis"},"Synopsis"),(0,a.kt)("p",null,"Generate tests manifest based on directory (e.g. for ArgoCD sync based on tests files)"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"testkube generate tests-crds <manifestDirectory> [flags]\n")),(0,a.kt)("h3",{id:"options"},"Options"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"      --artifact-dir stringArray                   artifact dirs for container executor\n      --artifact-storage-class-name string         artifact storage class name for container executor\n      --artifact-volume-mount-path string          artifact volume mount path for container executor\n      --command stringArray                        command passed to image in executor\n      --copy-files stringArray                     file path mappings from host to pod of form source:destination\n      --cronjob-template string                    cron job template file path for extensions to cron job template\n      --env stringToString                         envs in a form of name1=val1 passed to executor (default [])\n      --execution-name string                      execution name, if empty will be autogenerated\n      --executor-args stringArray                  executor binary additional arguments\n  -h, --help                                       help for tests-crds\n      --http-proxy string                          http proxy for executor containers\n      --https-proxy string                         https proxy for executor containers\n      --image string                               image for container executor\n      --image-pull-secrets stringArray             secret name used to pull the image in container executor\n      --job-template string                        job template file path for extensions to job template\n  -l, --label stringToString                       label key value pair: --label key1=value1 (default [])\n      --mount-configmap stringToString             config map value pair for mounting it to executor pod: --mount-configmap configmap_name=configmap_mountpath (default [])\n      --mount-secret stringToString                secret value pair for mounting it to executor pod: --mount-secret secret_name=secret_mountpath (default [])\n      --negative-test                              negative test, if enabled, makes failure an expected and correct test result. If the test fails the result will be set to success, and vice versa\n      --prerun-script string                       path to script to be run before test execution\n      --schedule string                            test schedule in a cron job form: * * * * *\n      --scraper-template string                    scraper template file path for extensions to scraper template\n      --secret-env stringToString                  secret envs in a form of secret_key1=secret_name1 passed to executor (default [])\n  -s, --secret-variable stringToString             secret variable key value pair: --secret-variable key1=value1 (default [])\n      --secret-variable-reference stringToString   secret variable references in a form name1=secret_name1=secret_key1 (default [])\n      --timeout int                                duration in seconds for test to timeout. 0 disables timeout.\n  -t, --type string                                test type\n      --upload-timeout string                      timeout to use when uploading files, example: 30s\n  -v, --variable stringToString                    variable key value pair: --variable key1=value1 (default [])\n      --variable-configmap stringArray             config map name used to map all keys to basis variables\n      --variable-secret stringArray                secret name used to map all keys to secret variables\n      --variables-file string                      variables file path, e.g. postman env file - will be passed to executor if supported\n")),(0,a.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},'  -a, --api-uri string     api uri, default value read from config if set (default "https://demo.testkube.io/results/v1")\n  -c, --client string      client used for connecting to Testkube API one of proxy|direct (default "proxy")\n      --namespace string   Kubernetes namespace, default value read from config if set (default "testkube")\n      --oauth-enabled      enable oauth\n      --verbose            show additional debug messages\n')),(0,a.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/cli/testkube_generate"},"testkube generate"),"\t - Generate resources commands")))}p.isMDXComponent=!0}}]);