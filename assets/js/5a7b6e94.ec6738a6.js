"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[2207],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>m});var a=n(67294);function u(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function s(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?s(Object(n),!0).forEach((function(t){u(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):s(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,a,u=function(e,t){if(null==e)return{};var n,a,u={},s=Object.keys(e);for(a=0;a<s.length;a++)n=s[a],t.indexOf(n)>=0||(u[n]=e[n]);return u}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(a=0;a<s.length;a++)n=s[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(u[n]=e[n])}return u}var i=a.createContext({}),l=function(e){var t=a.useContext(i),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},c=function(e){var t=l(e.components);return a.createElement(i.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},b=a.forwardRef((function(e,t){var n=e.components,u=e.mdxType,s=e.originalType,i=e.parentName,c=o(e,["components","mdxType","originalType","parentName"]),b=l(n),m=u,d=b["".concat(i,".").concat(m)]||b[m]||p[m]||s;return n?a.createElement(d,r(r({ref:t},c),{},{components:n})):a.createElement(d,r({ref:t},c))}));function m(e,t){var n=arguments,u=t&&t.mdxType;if("string"==typeof e||u){var s=n.length,r=new Array(s);r[0]=b;var o={};for(var i in t)hasOwnProperty.call(t,i)&&(o[i]=t[i]);o.originalType=e,o.mdxType="string"==typeof e?e:u,r[1]=o;for(var l=2;l<s;l++)r[l]=n[l];return a.createElement.apply(null,r)}return a.createElement.apply(null,n)}b.displayName="MDXCreateElement"},94495:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>i,contentTitle:()=>r,default:()=>p,frontMatter:()=>s,metadata:()=>o,toc:()=>l});var a=n(87462),u=(n(67294),n(3905));const s={},r="KubePug",o={unversionedId:"test-types/executor-kubepug",id:"test-types/executor-kubepug",title:"KubePug",description:"KubePug is a kubectl plugin checking for deprecated Kubernetes clusters or deprecated versions of Kubernetes manifests. It can connect to both your cluster directly and it can run on input files.",source:"@site/docs/test-types/executor-kubepug.md",sourceDirName:"test-types",slug:"/test-types/executor-kubepug",permalink:"/test-types/executor-kubepug",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/test-types/executor-kubepug.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"K6",permalink:"/test-types/executor-k6"},next:{title:"Maven",permalink:"/test-types/executor-maven"}},i={},l=[{value:"Testing Manifests",id:"testing-manifests",level:2},{value:"Testing the Output of <code>kubectl get</code>",id:"testing-the-output-of-kubectl-get",level:2},{value:"Testing Against a Previous Kubernetes Version",id:"testing-against-a-previous-kubernetes-version",level:2}],c={toc:l};function p(e){let{components:t,...n}=e;return(0,u.kt)("wrapper",(0,a.Z)({},c,n,{components:t,mdxType:"MDXLayout"}),(0,u.kt)("h1",{id:"kubepug"},"KubePug"),(0,u.kt)("p",null,(0,u.kt)("a",{parentName:"p",href:"https://github.com/rikatz/kubepug"},"KubePug")," is a kubectl plugin checking for deprecated Kubernetes clusters or deprecated versions of Kubernetes manifests. It can connect to both your cluster directly and it can run on input files.\nFor security, Testkube only supports scanning input files via the KubePug executor."),(0,u.kt)("p",null,"Default command for this executor: kubepug\nDefault arguments for this executor command: --format=json --input-file ","<","runPath",">","\n(parameters in ","<",">"," are calculated at test execution)"),(0,u.kt)("p",null,"Running the KubePug Testkube executor does not require any special installation; Testkube comes with the ability to run Kubepug immediately after installation."),(0,u.kt)("h2",{id:"testing-manifests"},"Testing Manifests"),(0,u.kt)("p",null,"By default, ",(0,u.kt)("inlineCode",{parentName:"p"},"kubepug")," downloads the latest ",(0,u.kt)("inlineCode",{parentName:"p"},"swagger.json")," from the ",(0,u.kt)("inlineCode",{parentName:"p"},"Kubernetes")," GitHub repository. When running it using Testkube, the same behavior is applied, unless a version is specified in the arguments."),(0,u.kt)("p",null,"Example input file:"),(0,u.kt)("pre",null,(0,u.kt)("code",{parentName:"pre",className:"language-yaml"},'apiVersion: v1\nconditions:\n- message: \'{"health":"true"}\'\n  status: "True"\n  type: Healthy\nkind: ComponentStatus\nmetadata:\n  creationTimestamp: null\n  name: etcd-1\n')),(0,u.kt)("p",null,"To test this using Testkube, first create a test, then run it:"),(0,u.kt)("pre",null,(0,u.kt)("code",{parentName:"pre",className:"language-bash"},'$ kubectl testkube create test --file file_name.yaml --type kubepug/yaml --name kubepug-example-test-1\nTest created testkube / kubepug-example-test-1 \ud83e\udd47\n\n\n$ kubectl testkube run test kubepug-example-test-1\nType          : kubepug/yaml\nName          : kubepug-example-test-1\nExecution ID  : 62b59ae1657713ea1b003a25\nExecution name: completely-helped-fowl\n\n\n\nTest execution started\n\nWatch test execution until complete:\n$ kubectl testkube watch execution 62b59ae1657713ea1b003a25\n\n\nUse following command to get test execution details:\n$ kubectl testkube get execution 62b59ae1657713ea1b003a25\n\n\n$ kubectl testkube get execution 62b59ae1657713ea1b003a25\nID:        62b59ae1657713ea1b003a25\nName:      completely-helped-fowl\nType:      kubepug/yaml\nDuration:  00:00:05\n\nStatus test execution failed:\n\n\u2a2f\n{"DeprecatedAPIs":[{"Description":"ComponentStatus (and ComponentStatusList) holds the cluster validation info. Deprecated: This API is deprecated in v1.19+","Group":"","Kind":"ComponentStatus","Version":"v1","Name":"","Deprecated":true,"Items":[{"Scope":"OBJECT","ObjectName":"etcd-1","Namespace":"","location":"/tmp/test-content4075001618"}]}],"DeletedAPIs":null}\n')),(0,u.kt)("p",null,"These tests also support input strings, file URIs, Git files and Git directories."),(0,u.kt)("h2",{id:"testing-the-output-of-kubectl-get"},"Testing the Output of ",(0,u.kt)("inlineCode",{parentName:"h2"},"kubectl get")),(0,u.kt)("p",null,"Another way to test Kubernetes objects is to create the Testkube Test with the output of ",(0,u.kt)("inlineCode",{parentName:"p"},"kubectl get"),". The output has to be in the correct format in order for KubePug to be able to scan it using the ",(0,u.kt)("inlineCode",{parentName:"p"},"-o yaml")," argument."),(0,u.kt)("pre",null,(0,u.kt)("code",{parentName:"pre",className:"language-bash"},"$ kubectl get PodSecurityPolicy gce.gke-metrics-agent -o yaml | kubectl testkube create test --type kubepug/yaml --name kubepug-example-test2\nWarning: policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+\nTest created testkube / kubepug-example-test2 \ud83e\udd47\n\n$ kubectl testkube run test kubepug-example-test2\nType          : kubepug/yaml\nName          : kubepug-example-test2\nExecution ID  : 62c8110338a672dc415ce98e\nExecution name: mostly-rapid-lark\n\n\nTest execution started\n\nWatch test execution until complete:\n$ kubectl testkube watch execution 62c8110338a672dc415ce98e\n\nUse following command to get test execution details:\n$ kubectl testkube get execution 62c8110338a672dc415ce98e\n")),(0,u.kt)("h2",{id:"testing-against-a-previous-kubernetes-version"},"Testing Against a Previous Kubernetes Version"),(0,u.kt)("p",null,"It is possible to run the same Testkube KubePug test using different Kubernetes versions to compare to using the ",(0,u.kt)("inlineCode",{parentName:"p"},"--k8s-version=${VERSION}")," argument as shown below:"),(0,u.kt)("pre",null,(0,u.kt)("code",{parentName:"pre",className:"language-bash"},'$ kubectl testkube run test kubepug-example-test-1 --args \'--k8s-version=v1.18.0\'\nType          : kubepug/yaml\nName          : kubepug-example-test-1\nExecution ID  : 62b59d52657713ea1b003a2d\nExecution name: notably-healthy-cricket\n\n\n\nTest execution started\n\nWatch test execution until complete:\n$ kubectl testkube watch execution 62b59d52657713ea1b003a2d\n\nUse following command to get test execution details:\n$ kubectl testkube get execution 62b59d52657713ea1b003a2d\n\n\n$ kubectl testkube get execution 62b59d52657713ea1b003a2d\nID:        62b59d52657713ea1b003a2d\nName:      sincerely-real-marten\nType:      kubepug/yaml\nDuration:  00:00:05\nArgs:     --k8s-version=v1.18.0\n\n{"DeprecatedAPIs":null,"DeletedAPIs":null}\nStatus Test execution completed with success \ud83e\udd47\n')),(0,u.kt)("p",null,"It is also possible to pass other arguments to the executor. For the options please consult the ",(0,u.kt)("a",{parentName:"p",href:"https://github.com/rikatz/kubepug#how-to-use-it-as-a-standalone-program"},"KubePug documentation"),"."))}p.isMDXComponent=!0}}]);