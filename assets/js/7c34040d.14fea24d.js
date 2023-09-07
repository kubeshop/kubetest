"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[8298],{3905:(e,t,r)=>{r.d(t,{Zo:()=>p,kt:()=>m});var n=r(67294);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function s(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},i=Object.keys(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)r=i[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var c=n.createContext({}),l=function(e){var t=n.useContext(c),r=t;return e&&(r="function"==typeof e?e(t):o(o({},t),e)),r},p=function(e){var t=l(e.components);return n.createElement(c.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},f=n.forwardRef((function(e,t){var r=e.components,a=e.mdxType,i=e.originalType,c=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),f=l(r),m=a,g=f["".concat(c,".").concat(m)]||f[m]||u[m]||i;return r?n.createElement(g,o(o({ref:t},p),{},{components:r})):n.createElement(g,o({ref:t},p))}));function m(e,t){var r=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=r.length,o=new Array(i);o[0]=f;var s={};for(var c in t)hasOwnProperty.call(t,c)&&(s[c]=t[c]);s.originalType=e,s.mdxType="string"==typeof e?e:a,o[1]=s;for(var l=2;l<i;l++)o[l]=r[l];return n.createElement.apply(null,o)}return n.createElement.apply(null,r)}f.displayName="MDXCreateElement"},47610:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>c,contentTitle:()=>o,default:()=>u,frontMatter:()=>i,metadata:()=>s,toc:()=>l});var n=r(87462),a=(r(67294),r(3905));const i={},o=void 0,s={unversionedId:"cli/testkube_create_test",id:"cli/testkube_create_test",title:"testkube_create_test",description:"testkube create test",source:"@site/docs/cli/testkube_create_test.md",sourceDirName:"cli",slug:"/cli/testkube_create_test",permalink:"/cli/testkube_create_test",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/cli/testkube_create_test.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"testkube_create_template",permalink:"/cli/testkube_create_template"},next:{title:"testkube_create_testsource",permalink:"/cli/testkube_create_testsource"}},c={},l=[{value:"testkube create test",id:"testkube-create-test",level:2},{value:"Synopsis",id:"synopsis",level:3},{value:"Options",id:"options",level:3},{value:"Options inherited from parent commands",id:"options-inherited-from-parent-commands",level:3},{value:"SEE ALSO",id:"see-also",level:3}],p={toc:l};function u(e){let{components:t,...r}=e;return(0,a.kt)("wrapper",(0,n.Z)({},p,r,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h2",{id:"testkube-create-test"},"testkube create test"),(0,a.kt)("p",null,"Create new Test"),(0,a.kt)("h3",{id:"synopsis"},"Synopsis"),(0,a.kt)("p",null,"Create new Test Custom Resource"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},"testkube create test [flags]\n")),(0,a.kt)("h3",{id:"options"},"Options"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},'      --args-mode string                           usage mode for arguments. one of append|override (default "append")\n      --artifact-dir stringArray                   artifact dirs for scraping\n      --artifact-omit-folder-per-execution         don\'t store artifacts in execution folder\n      --artifact-storage-bucket string             artifact storage class name for container executor\n      --artifact-storage-class-name string         artifact storage class name for container executor\n      --artifact-volume-mount-path string          artifact volume mount path for container executor\n      --command stringArray                        command passed to image in executor\n      --copy-files stringArray                     file path mappings from host to pod of form source:destination\n      --cronjob-template string                    cron job template file path for extensions to cron job template\n      --cronjob-template-reference string          reference to cron job template to use for the test\n      --description string                         test description\n      --env stringToString                         envs in a form of name1=val1 passed to executor (default [])\n      --execution-name string                      execution name, if empty will be autogenerated\n      --executor-args stringArray                  executor binary additional arguments\n  -f, --file string                                test file - will be read from stdin if not specified\n      --git-auth-type string                       auth type for git requests one of basic|header (default "basic")\n      --git-branch string                          if uri is git repository we can set additional branch parameter\n      --git-certificate-secret string              if git repository is private we can use certificate as an auth parameter stored in a kubernetes secret name\n      --git-commit string                          if uri is git repository we can use commit id (sha) parameter\n      --git-path string                            if repository is big we need to define additional path to directory/file to checkout partially\n      --git-token string                           if git repository is private we can use token as an auth parameter\n      --git-token-secret stringToString            git token secret in a form of secret_name1=secret_key1 for private repository (default [])\n      --git-uri string                             Git repository uri\n      --git-username string                        if git repository is private we can use username as an auth parameter\n      --git-username-secret stringToString         git username secret in a form of secret_name1=secret_key1 for private repository (default [])\n      --git-working-dir string                     if repository contains multiple directories with tests (like monorepo) and one starting directory we can set working directory parameter\n  -h, --help                                       help for test\n      --http-proxy string                          http proxy for executor containers\n      --https-proxy string                         https proxy for executor containers\n      --image string                               image for container executor\n      --image-pull-secrets stringArray             secret name used to pull the image in container executor\n      --job-template string                        job template file path for extensions to job template\n      --job-template-reference string              reference to job template to use for the test\n  -l, --label stringToString                       label key value pair: --label key1=value1 (default [])\n      --mount-configmap stringToString             config map value pair for mounting it to executor pod: --mount-configmap configmap_name=configmap_mountpath (default [])\n      --mount-secret stringToString                secret value pair for mounting it to executor pod: --mount-secret secret_name=secret_mountpath (default [])\n  -n, --name string                                unique test name - mandatory\n      --negative-test                              negative test, if enabled, makes failure an expected and correct test result. If the test fails the result will be set to success, and vice versa\n      --postrun-script string                      path to script to be run after test execution\n      --prerun-script string                       path to script to be run before test execution\n      --pvc-template string                        pvc template file path for extensions to pvc template\n      --pvc-template-reference string              reference to pvc template to use for the test\n      --schedule string                            test schedule in a cron job form: * * * * *\n      --scraper-template string                    scraper template file path for extensions to scraper template\n      --scraper-template-reference string          reference to scraper template to use for the test\n      --secret-env stringToString                  secret envs in a form of secret_key1=secret_name1 passed to executor (default [])\n  -s, --secret-variable stringToString             secret variable key value pair: --secret-variable key1=value1 (default [])\n      --secret-variable-reference stringToString   secret variable references in a form name1=secret_name1=secret_key1 (default [])\n      --source string                              source name - will be used together with content parameters\n      --test-content-type string                   content type of test one of string|file-uri|git\n      --timeout int                                duration in seconds for test to timeout. 0 disables timeout.\n  -t, --type string                                test type\n      --upload-timeout string                      timeout to use when uploading files, example: 30s\n      --uri string                                 URI of resource - will be loaded by http GET\n  -v, --variable stringToString                    variable key value pair: --variable key1=value1 (default [])\n      --variable-configmap stringArray             config map name used to map all keys to basis variables\n      --variable-secret stringArray                secret name used to map all keys to secret variables\n      --variables-file string                      variables file path, e.g. postman env file - will be passed to executor if supported\n')),(0,a.kt)("h3",{id:"options-inherited-from-parent-commands"},"Options inherited from parent commands"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre"},'  -a, --api-uri string     api uri, default value read from config if set (default "https://demo.testkube.io/results/v1")\n  -c, --client string      client used for connecting to Testkube API one of proxy|direct (default "proxy")\n      --crd-only           generate only crd\n      --namespace string   Kubernetes namespace, default value read from config if set (default "testkube")\n      --oauth-enabled      enable oauth\n      --verbose            show additional debug messages\n')),(0,a.kt)("h3",{id:"see-also"},"SEE ALSO"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/cli/testkube_create"},"testkube create"),"\t - Create resource")))}u.isMDXComponent=!0}}]);