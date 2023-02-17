"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[7338],{3905:(e,t,s)=>{s.d(t,{Zo:()=>i,kt:()=>h});var n=s(67294);function a(e,t,s){return t in e?Object.defineProperty(e,t,{value:s,enumerable:!0,configurable:!0,writable:!0}):e[t]=s,e}function l(e,t){var s=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),s.push.apply(s,n)}return s}function r(e){for(var t=1;t<arguments.length;t++){var s=null!=arguments[t]?arguments[t]:{};t%2?l(Object(s),!0).forEach((function(t){a(e,t,s[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(s)):l(Object(s)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(s,t))}))}return e}function u(e,t){if(null==e)return{};var s,n,a=function(e,t){if(null==e)return{};var s,n,a={},l=Object.keys(e);for(n=0;n<l.length;n++)s=l[n],t.indexOf(s)>=0||(a[s]=e[s]);return a}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(n=0;n<l.length;n++)s=l[n],t.indexOf(s)>=0||Object.prototype.propertyIsEnumerable.call(e,s)&&(a[s]=e[s])}return a}var c=n.createContext({}),o=function(e){var t=n.useContext(c),s=t;return e&&(s="function"==typeof e?e(t):r(r({},t),e)),s},i=function(e){var t=o(e.components);return n.createElement(c.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},p=n.forwardRef((function(e,t){var s=e.components,a=e.mdxType,l=e.originalType,c=e.parentName,i=u(e,["components","mdxType","originalType","parentName"]),p=o(s),h=a,m=p["".concat(c,".").concat(h)]||p[h]||d[h]||l;return s?n.createElement(m,r(r({ref:t},i),{},{components:s})):n.createElement(m,r({ref:t},i))}));function h(e,t){var s=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var l=s.length,r=new Array(l);r[0]=p;var u={};for(var c in t)hasOwnProperty.call(t,c)&&(u[c]=t[c]);u.originalType=e,u.mdxType="string"==typeof e?e:a,r[1]=u;for(var o=2;o<l;o++)r[o]=s[o];return n.createElement.apply(null,r)}return n.createElement.apply(null,s)}p.displayName="MDXCreateElement"},21257:(e,t,s)=>{s.r(t),s.d(t,{assets:()=>c,contentTitle:()=>r,default:()=>d,frontMatter:()=>l,metadata:()=>u,toc:()=>o});var n=s(87462),a=(s(67294),s(3905));const l={},r="Test and Test Suite Scheduling",u={unversionedId:"guides/scheduling",id:"guides/scheduling",title:"Test and Test Suite Scheduling",description:"In order to run Tests and Test Suites on a regular basis, we support a scheduling mechanism for these objects.",source:"@site/docs/guides/scheduling.md",sourceDirName:"guides",slug:"/guides/scheduling",permalink:"/guides/scheduling",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/guides/scheduling.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Secret Variables",permalink:"/guides/secrets"},next:{title:"Artifacts Storage",permalink:"/guides/artifacts-storage"}},c={},o=[{value:"<strong>Scheduling Architecture</strong>",id:"scheduling-architecture",level:2},{value:"<strong>Create a Test with a Schedule</strong>",id:"create-a-test-with-a-schedule",level:2},{value:"<strong>Getting Scheduled Test Results</strong>",id:"getting-scheduled-test-results",level:2},{value:"<strong>Create a Test Suite with a Schedule</strong>",id:"create-a-test-suite-with-a-schedule",level:2},{value:"<strong>Getting Scheduled Test Suite Results</strong>",id:"getting-scheduled-test-suite-results",level:2}],i={toc:o};function d(e){let{components:t,...s}=e;return(0,a.kt)("wrapper",(0,n.Z)({},i,s,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"test-and-test-suite-scheduling"},"Test and Test Suite Scheduling"),(0,a.kt)("p",null,"In order to run Tests and Test Suites on a regular basis, we support a scheduling mechanism for these objects.\nCRDs both for test and test suite contain a ",(0,a.kt)("strong",{parentName:"p"},"schedule")," field used to define rules for launching them in time.\nTestkube's schedule data format is the same that is used to define Kubernetes Cron jobs (check Wikipedia Cron format for details ",(0,a.kt)("a",{parentName:"p",href:"https://en.wikipedia.org/wiki/Cron"},"https://en.wikipedia.org/wiki/Cron"),")."),(0,a.kt)("h2",{id:"scheduling-architecture"},(0,a.kt)("strong",{parentName:"h2"},"Scheduling Architecture")),(0,a.kt)("p",null,"Testkube uses the scheduling engine from Kubernetes Cron jobs.\nIn fact, for each scheduled Test or Test Suite, a special cron job is created from this template:\n",(0,a.kt)("a",{parentName:"p",href:"https://github.com/kubeshop/helm-charts/tree/main/charts/testkube-api/cronjob-template.yml"},"https://github.com/kubeshop/helm-charts/tree/main/charts/testkube-api/cronjob-template.yml"),".\nTechnically, it is a callback to Testkube API server method launching either Test or Test Suite execution.\nThis works similarly to scheduled Test and Test Suite executions done by external scheduling platforms. "),(0,a.kt)("h2",{id:"create-a-test-with-a-schedule"},(0,a.kt)("strong",{parentName:"h2"},"Create a Test with a Schedule")),(0,a.kt)("p",null,"Let's create a test with a required schedule using Testkube CLI command:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},'kubectl testkube create test --file test/postman/TODO.postman_collection.json --name scheduled-test --schedule="*/1 * * * *"\n')),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"Detected test type postman/collection\nTest created  / scheduled-test \ud83e\udd47\n")),(0,a.kt)("p",null,"We successfully created a scheduled test and can check a list of the available tests:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube get tests\n")),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"  NAME              | TYPE               | CREATED                       | LABELS | SCHEDULE    | STATUS | EXECUTION ID              \n+-------------------+--------------------+-------------------------------+--------+-------------+--------+--------------------------+\n  scheduled-test    | postman/collection | 2022-04-13 12:37:40 +0000 UTC |        | */1 * * * * |        |                           \n")),(0,a.kt)("p",null,"The scheduled test was created and successfully scheduled for execution.\nLet's check a Cron job connected to this test."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl get cronjobs -A\n")),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"NAMESPACE   NAME                   SCHEDULE      SUSPEND   ACTIVE   LAST SCHEDULE   AGE\ntestkube    scheduled-test-tests   */1 * * * *   False     1        42s           3m22s\n")),(0,a.kt)("p",null,"The Cron job for this test was successfully created and the test was executed.\nTo the check Cron job details:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl describe cronjob scheduled-test-tests -n testkube\n")),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},'Name:                          scheduled-test-tests\nNamespace:                     testkube\nLabels:                        testkube=tests\nAnnotations:                   <none>\nSchedule:                      */1 * * * *\nConcurrency Policy:            Forbid\nSuspend:                       False\nSuccessful Job History Limit:  3\nFailed Job History Limit:      1\nStarting Deadline Seconds:     <unset>\nSelector:                      <unset>\nParallelism:                   <unset>\nCompletions:                   <unset>\nPod Template:\n  Labels:  <none>\n  Containers:\n   curlimage:\n    Image:      curlimages/curl\n    Port:       <none>\n    Host Port:  <none>\n    Command:\n      sh\n      -c\n    Args:\n      curl -X POST -H "Content-Type: application/json" -d \'{}\' "http://testkube-api-server:8088/v1/tests/scheduled-test/executions?callback=true"\n    Environment:     <none>\n    Mounts:          <none>\n  Volumes:           <none>\nLast Schedule Time:  Wed, 13 Apr 2022 15:50:00 +0300\nActive Jobs:         scheduled-test-tests-27497570\nEvents:\n  Type    Reason            Age                  From                Message\n  ----    ------            ----                 ----                -------\n  Normal  SuccessfulCreate  5m41s                cronjob-controller  Created job scheduled-test-tests-2749757\n')),(0,a.kt)("p",null,"As mentioned above, we have a scheduled callback for launching our test."),(0,a.kt)("h2",{id:"getting-scheduled-test-results"},(0,a.kt)("strong",{parentName:"h2"},"Getting Scheduled Test Results")),(0,a.kt)("p",null,"Now we can check whether the test is executed every minute for the schedule we provided."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube get execution\n")),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"  ID                       | NAME                | TYPE               | STATUS  | LABELS  \n+--------------------------+---------------------+--------------------+---------+--------+\n  6256c98f418062706814e1fc | scheduled-test      | postman/collection | passed  |         \n  6256c953418062706814e1fa | scheduled-test      | postman/collection | passed  |         \n  6256c91e418062706814e1f8 | scheduled-test      | postman/collection | passed  |         \n  6256c8db418062706814e1f6 | scheduled-test      | postman/collection | passed  |         \n  6256c89f418062706814e1f4 | scheduled-test      | postman/collection | passed  |         \n  6256c885418062706814e1f2 | scheduled-test      | postman/collection | passed  |         \n  6256c87e418062706814e1f0 | scheduled-test      | postman/collection | passed  | \n")),(0,a.kt)("p",null,"The test is successfully regularly executed."),(0,a.kt)("h2",{id:"create-a-test-suite-with-a-schedule"},(0,a.kt)("strong",{parentName:"h2"},"Create a Test Suite with a Schedule")),(0,a.kt)("p",null,"Let's create a Test Suite with a required schedule using the Testkube CLI command:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},'cat test/suites/testsuite.json | kubectl testkube create testsuite --name scheduled-testsuite --schedule="*/1 * * * *"\n')),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"TestSuite created scheduled-testsuite \ud83e\udd47\n")),(0,a.kt)("p",null,"We successfully created a scheduled Test Suite and can view a list of the available Test Suites:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube get testsuites\n")),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"  NAME                | DESCRIPTION            | STEPS | LABELS | SCHEDULE    | STATUS | EXECUTION ID  \n+---------------------+------------------------+-------+--------+-------------+--------+--------------+\n  scheduled-testsuite | Run test several times |     2 |        | */1 * * * * |        |    \n")),(0,a.kt)("p",null,"The scheduled test suite was created and successfully scheduled for execution.\nWe will skip the Cron job details, they are fully similar to test one described above."),(0,a.kt)("h2",{id:"getting-scheduled-test-suite-results"},(0,a.kt)("strong",{parentName:"h2"},"Getting Scheduled Test Suite Results")),(0,a.kt)("p",null,"The Test Suite is executed every minute for the schedule we provided."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube get tse\n")),(0,a.kt)("p",null,"Output:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-bash"},"  ID                       | TEST SUITE NAME     | EXECUTION NAME                             | STATUS | STEPS | LABELS  \n+--------------------------+---------------------+--------------------------------------------+--------+-------+--------+\n  6256ce3f418062706814e210 | scheduled-testsuite | scheduled-testsuite.abnormally-in-lark     | passed |     2 |\n  6256ce04418062706814e20c | scheduled-testsuite | scheduled-testsuite.kindly-evolved-primate | passed |     2 |\n  6256cdcc418062706814e208 | scheduled-testsuite | scheduled-testsuite.formerly-champion-dodo | passed |     2 |\n")),(0,a.kt)("p",null,"The Test Suite is successfully executed according to the schedule set."))}d.isMDXComponent=!0}}]);