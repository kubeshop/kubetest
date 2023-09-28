"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[7735],{3905:(e,t,n)=>{n.d(t,{Zo:()=>c,kt:()=>d});var a=n(67294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var l=a.createContext({}),u=function(e){var t=a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},c=function(e){var t=u(e.components);return a.createElement(l.Provider,{value:t},e.children)},m={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},p=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,o=e.originalType,l=e.parentName,c=s(e,["components","mdxType","originalType","parentName"]),p=u(n),d=r,k=p["".concat(l,".").concat(d)]||p[d]||m[d]||o;return n?a.createElement(k,i(i({ref:t},c),{},{components:n})):a.createElement(k,i({ref:t},c))}));function d(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=n.length,i=new Array(o);i[0]=p;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:r,i[1]=s;for(var u=2;u<o;u++)i[u]=n[u];return a.createElement.apply(null,i)}return a.createElement.apply(null,n)}p.displayName="MDXCreateElement"},23612:(e,t,n)=>{n.d(t,{Z:()=>k});var a=n(67294),r=n(86010),o=n(35281),i=n(95999);const s="admonition_LlT9",l="admonitionHeading_tbUL",u="admonitionIcon_kALy",c="admonitionContent_S0QG";const m={note:{infimaClassName:"secondary",iconComponent:function(){return a.createElement("svg",{viewBox:"0 0 14 16"},a.createElement("path",{fillRule:"evenodd",d:"M6.3 5.69a.942.942 0 0 1-.28-.7c0-.28.09-.52.28-.7.19-.18.42-.28.7-.28.28 0 .52.09.7.28.18.19.28.42.28.7 0 .28-.09.52-.28.7a1 1 0 0 1-.7.3c-.28 0-.52-.11-.7-.3zM8 7.99c-.02-.25-.11-.48-.31-.69-.2-.19-.42-.3-.69-.31H6c-.27.02-.48.13-.69.31-.2.2-.3.44-.31.69h1v3c.02.27.11.5.31.69.2.2.42.31.69.31h1c.27 0 .48-.11.69-.31.2-.19.3-.42.31-.69H8V7.98v.01zM7 2.3c-3.14 0-5.7 2.54-5.7 5.68 0 3.14 2.56 5.7 5.7 5.7s5.7-2.55 5.7-5.7c0-3.15-2.56-5.69-5.7-5.69v.01zM7 .98c3.86 0 7 3.14 7 7s-3.14 7-7 7-7-3.12-7-7 3.14-7 7-7z"}))},label:a.createElement(i.Z,{id:"theme.admonition.note",description:"The default label used for the Note admonition (:::note)"},"note")},tip:{infimaClassName:"success",iconComponent:function(){return a.createElement("svg",{viewBox:"0 0 12 16"},a.createElement("path",{fillRule:"evenodd",d:"M6.5 0C3.48 0 1 2.19 1 5c0 .92.55 2.25 1 3 1.34 2.25 1.78 2.78 2 4v1h5v-1c.22-1.22.66-1.75 2-4 .45-.75 1-2.08 1-3 0-2.81-2.48-5-5.5-5zm3.64 7.48c-.25.44-.47.8-.67 1.11-.86 1.41-1.25 2.06-1.45 3.23-.02.05-.02.11-.02.17H5c0-.06 0-.13-.02-.17-.2-1.17-.59-1.83-1.45-3.23-.2-.31-.42-.67-.67-1.11C2.44 6.78 2 5.65 2 5c0-2.2 2.02-4 4.5-4 1.22 0 2.36.42 3.22 1.19C10.55 2.94 11 3.94 11 5c0 .66-.44 1.78-.86 2.48zM4 14h5c-.23 1.14-1.3 2-2.5 2s-2.27-.86-2.5-2z"}))},label:a.createElement(i.Z,{id:"theme.admonition.tip",description:"The default label used for the Tip admonition (:::tip)"},"tip")},danger:{infimaClassName:"danger",iconComponent:function(){return a.createElement("svg",{viewBox:"0 0 12 16"},a.createElement("path",{fillRule:"evenodd",d:"M5.05.31c.81 2.17.41 3.38-.52 4.31C3.55 5.67 1.98 6.45.9 7.98c-1.45 2.05-1.7 6.53 3.53 7.7-2.2-1.16-2.67-4.52-.3-6.61-.61 2.03.53 3.33 1.94 2.86 1.39-.47 2.3.53 2.27 1.67-.02.78-.31 1.44-1.13 1.81 3.42-.59 4.78-3.42 4.78-5.56 0-2.84-2.53-3.22-1.25-5.61-1.52.13-2.03 1.13-1.89 2.75.09 1.08-1.02 1.8-1.86 1.33-.67-.41-.66-1.19-.06-1.78C8.18 5.31 8.68 2.45 5.05.32L5.03.3l.02.01z"}))},label:a.createElement(i.Z,{id:"theme.admonition.danger",description:"The default label used for the Danger admonition (:::danger)"},"danger")},info:{infimaClassName:"info",iconComponent:function(){return a.createElement("svg",{viewBox:"0 0 14 16"},a.createElement("path",{fillRule:"evenodd",d:"M7 2.3c3.14 0 5.7 2.56 5.7 5.7s-2.56 5.7-5.7 5.7A5.71 5.71 0 0 1 1.3 8c0-3.14 2.56-5.7 5.7-5.7zM7 1C3.14 1 0 4.14 0 8s3.14 7 7 7 7-3.14 7-7-3.14-7-7-7zm1 3H6v5h2V4zm0 6H6v2h2v-2z"}))},label:a.createElement(i.Z,{id:"theme.admonition.info",description:"The default label used for the Info admonition (:::info)"},"info")},caution:{infimaClassName:"warning",iconComponent:function(){return a.createElement("svg",{viewBox:"0 0 16 16"},a.createElement("path",{fillRule:"evenodd",d:"M8.893 1.5c-.183-.31-.52-.5-.887-.5s-.703.19-.886.5L.138 13.499a.98.98 0 0 0 0 1.001c.193.31.53.501.886.501h13.964c.367 0 .704-.19.877-.5a1.03 1.03 0 0 0 .01-1.002L8.893 1.5zm.133 11.497H6.987v-2.003h2.039v2.003zm0-3.004H6.987V5.987h2.039v4.006z"}))},label:a.createElement(i.Z,{id:"theme.admonition.caution",description:"The default label used for the Caution admonition (:::caution)"},"caution")}},p={secondary:"note",important:"info",success:"tip",warning:"danger"};function d(e){const{mdxAdmonitionTitle:t,rest:n}=function(e){const t=a.Children.toArray(e),n=t.find((e=>{var t;return a.isValidElement(e)&&"mdxAdmonitionTitle"===(null==(t=e.props)?void 0:t.mdxType)})),r=a.createElement(a.Fragment,null,t.filter((e=>e!==n)));return{mdxAdmonitionTitle:n,rest:r}}(e.children);return{...e,title:e.title??t,children:n}}function k(e){const{children:t,type:n,title:i,icon:k}=d(e),v=function(e){const t=p[e]??e;return m[t]||(console.warn(`No admonition config found for admonition type "${t}". Using Info as fallback.`),m.info)}(n),h=i??v.label,{iconComponent:f}=v,b=k??a.createElement(f,null);return a.createElement("div",{className:(0,r.Z)(o.k.common.admonition,o.k.common.admonitionType(e.type),"alert",`alert--${v.infimaClassName}`,s)},a.createElement("div",{className:l},a.createElement("span",{className:u},b),h),a.createElement("div",{className:c},t))}},22167:(e,t,n)=>{n.r(t),n.d(t,{ExecutorInfo:()=>m,assets:()=>u,contentTitle:()=>s,default:()=>d,frontMatter:()=>i,metadata:()=>l,toc:()=>c});var a=n(87462),r=(n(67294),n(3905)),o=n(23612);const i={},s="Maven",l={unversionedId:"test-types/executor-maven",id:"test-types/executor-maven",title:"Maven",description:"Testkube allows you to run Maven-based tasks which could be also tests. For example, we can easily run JUnit tests in Testkube now.",source:"@site/docs/test-types/executor-maven.md",sourceDirName:"test-types",slug:"/test-types/executor-maven",permalink:"/test-types/executor-maven",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/test-types/executor-maven.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Kube no trouble",permalink:"/test-types/executor-kubent"},next:{title:"Playwright",permalink:"/test-types/executor-playwright"}},u={},c=[{value:"Test Environment",id:"test-environment",level:2},{value:"Create a New Maven-based Test",id:"create-a-new-maven-based-test",level:2},{value:"Running a Test",id:"running-a-test",level:2},{value:"Getting Test Results",id:"getting-test-results",level:2},{value:"Using Different JDKs",id:"using-different-jdks",level:2},{value:"Summary",id:"summary",level:2}],m=()=>(0,r.kt)("div",null,(0,r.kt)(o.Z,{type:"info",icon:"\ud83c\udf93",title:"What is Maven?",mdxType:"Admonition"},(0,r.kt)("ul",null,(0,r.kt)("li",null,"Maven is a build automation tool used primarily for Java projects."),(0,r.kt)("li",null,"Paired with JUnit, a testing framework that is built in the Maven project format, you can build and run unit tests for your projects.")))),p={toc:c,ExecutorInfo:m};function d(e){let{components:t,...n}=e;return(0,r.kt)("wrapper",(0,a.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"maven"},"Maven"),(0,r.kt)("p",null,"Testkube allows you to run Maven-based tasks which could be also tests. For example, we can easily run JUnit tests in Testkube now. "),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"Default command for this executor: ",(0,r.kt)("inlineCode",{parentName:"li"},"mvn")),(0,r.kt)("li",{parentName:"ul"},"Default arguments for this executor command: ",(0,r.kt)("inlineCode",{parentName:"li"},"--settings")," ",(0,r.kt)("inlineCode",{parentName:"li"},"<settingsFile>")," ",(0,r.kt)("inlineCode",{parentName:"li"},"<goalName>")," ",(0,r.kt)("inlineCode",{parentName:"li"},"-Duser.home")," ",(0,r.kt)("inlineCode",{parentName:"li"},"<mavenHome>"),"\n(parameters in ",(0,r.kt)("inlineCode",{parentName:"li"},"<>")," are calculated at test execution)")),(0,r.kt)(m,{mdxType:"ExecutorInfo"}),(0,r.kt)("h2",{id:"test-environment"},"Test Environment"),(0,r.kt)("p",null,"We'll try to add a simple JUnit test to our cluster and run it. Testkube Maven Executor handles ",(0,r.kt)("inlineCode",{parentName:"p"},"mvn")," and ",(0,r.kt)("inlineCode",{parentName:"p"},"mvnw")," binaries.\nBecause Maven projects are quite complicated in terms of directory structure. We'll need to load them from a Git directory."),(0,r.kt)("p",null,"You can find example projects in the repository here: ",(0,r.kt)("a",{parentName:"p",href:"https://github.com/kubeshop/testkube-executor-maven/tree/main/examples"},"https://github.com/kubeshop/testkube-executor-maven/tree/main/examples"),"."),(0,r.kt)("p",null,"Let's create a simple test which will check if an env variable is set to true: "),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-java"},'package hello.maven;\n\nimport org.junit.jupiter.api.Test;\nimport static org.junit.jupiter.api.Assertions.*;\n\nclass LibraryTest {\n    @Test void someLibraryMethodReturnsTrue() {\n        String env = System.getenv("TESTKUBE_MAVEN");\n        assertTrue(Boolean.parseBoolean(env), "TESTKUBE_MAVEN env should be true");\n    }\n}\n')),(0,r.kt)("p",null,"The default Maven executor: "),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: executor.testkube.io/v1\nkind: Executor\nmetadata:\n  name: maven-executor\n  namespace: testkube\nspec:\n  image: kubeshop/testkube-maven-executor:0.1.4\n  types:\n  - maven/project\n  - maven/test\n  - maven/integration-test \n")),(0,r.kt)("p",null,"As we can see, there are several types. The Maven executor handles the second part after ",(0,r.kt)("inlineCode",{parentName:"p"},"/")," as a task name, so ",(0,r.kt)("inlineCode",{parentName:"p"},"maven/test")," will run ",(0,r.kt)("inlineCode",{parentName:"p"},"mvn test")," and so on. "),(0,r.kt)("p",null,"One exception from this rule is ",(0,r.kt)("inlineCode",{parentName:"p"},"project")," which is a generic one and forces you to pass additional arguments during test execution. For example:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube run maven-example-project --args='runMyCustomTask' \n")),(0,r.kt)("h2",{id:"create-a-new-maven-based-test"},"Create a New Maven-based Test"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube create test --git-uri https://github.com/kubeshop/testkube-executor-maven.git --git-path examples/hello-maven --type maven/test --name maven-example-test --git-branch main\n")),(0,r.kt)("h2",{id:"running-a-test"},"Running a Test"),(0,r.kt)("p",null,"Let's pass the env variable to our test run:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube run test maven-example-test -f -v TESTKUBE_MAVEN=true\n\n# ...... after some time\n\nTest execution completed with success in 16.555s \ud83e\udd47\n\nWatch the test execution until complete:\n$ kubectl testkube watch execution 62d148db0260f256c1a1e993\n\n\nUse the following command to get test execution details:\n$ kubectl testkube get execution 62d148db0260f256c1a1e993\n")),(0,r.kt)("h2",{id:"getting-test-results"},"Getting Test Results"),(0,r.kt)("p",null,"Now we can watch/get test execution details:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube get execution 62d148db0260f256c1a1e993\n")),(0,r.kt)("p",null,"Output:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"# ....... a lot of Maven logs\n\nDownloaded from central: https://repo.maven.apache.org/maven2/org/junit/platform/junit-platform-launcher/1.7.2/junit-platform-launcher-1.7.2.pom (3.0 kB at 121 kB/s)\n[INFO] \n[INFO] -------------------------------------------------------\n[INFO]  T E S T S\n[INFO] -------------------------------------------------------\n[INFO] Running hello.maven.LibraryTest\n[INFO] Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.052 s - in hello.maven.LibraryTest\n[INFO] \n[INFO] Results:\n[INFO] \n[INFO] Tests run: 1, Failures: 0, Errors: 0, Skipped: 0\n[INFO] \n[INFO] ------------------------------------------------------------------------\n[INFO] BUILD SUCCESS\n[INFO] ------------------------------------------------------------------------\n[INFO] Total time:  9.851 s\n[INFO] Finished at: 2022-07-18T09:06:15Z\n[INFO] ------------------------------------------------------------------------\n\nStatus Test execution completed with success \ud83e\udd47\n")),(0,r.kt)("h2",{id:"using-different-jdks"},"Using Different JDKs"),(0,r.kt)("p",null,"In the Java world, usually you want to have control over your Runtime environment. Testkube can easily handle that for you!\nWe're building several Java images to handle constraints which Maven can put in its build file."),(0,r.kt)("p",null,"To use a different executor you can use one of our pre-built ones (for Java 8,11,17,18) or build your own Docker image based on a Maven executor."),(0,r.kt)("p",null,"Let's assume we need JDK18 for our test runs. To handle that issue, create a new Maven executor:"),(0,r.kt)("p",null,"content of ",(0,r.kt)("inlineCode",{parentName:"p"},"maven-jdk18-executor.yaml")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'apiVersion: executor.testkube.io/v1\nkind: Executor\nmetadata:\n  name: maven-jdk18-executor\n  namespace: testkube\nspec:\n  image: kubeshop/testkube-maven-executor:0.1.0-jdk18   # <-- we\'re building jdk\n  types:\n  - maven:jdk18/project # <-- just create different test type with naming convention "framework:version/type"\n  - maven:jdk18/test\n  - maven:jdk18/integration-test \n')),(0,r.kt)("blockquote",null,(0,r.kt)("p",{parentName:"blockquote"},"Tip: Look for recent executor versions here ",(0,r.kt)("a",{parentName:"p",href:"https://hub.docker.com/repository/registry-1.docker.io/kubeshop/testkube-maven-executor/tags?page=1&ordering=last_updated"},"https://hub.docker.com/repository/registry-1.docker.io/kubeshop/testkube-maven-executor/tags?page=1&ordering=last_updated"),".")),(0,r.kt)("p",null,"And add it to your cluster: "),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl apply -f maven-jdk18-executor.yaml \n")),(0,r.kt)("p",null,"Now, create a new test with a type which our new executor can handle e.g.: ",(0,r.kt)("inlineCode",{parentName:"p"},"maven:jdk18/test")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"}," # create test\n kubectl testkube create test --git-uri https://github.com/kubeshop/testkube-executor-maven.git --git-path examples/hello-maven-jdk18 --type maven:jdk18/test --name maven-jdk18-example-test --git-branch main\n\n# and run it\nkubectl testkube run test maven-jdk18-example-test -f -v TESTKUBE_MAVEN=true\n")),(0,r.kt)("h2",{id:"summary"},"Summary"),(0,r.kt)("p",null,"Testkube simplifies running Java tests based on Maven and simplifies the merging of Java based tests into your global testing ecosystem."))}d.isMDXComponent=!0}}]);