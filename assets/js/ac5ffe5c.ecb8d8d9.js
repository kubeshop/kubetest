"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[2251],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>d});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function s(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function o(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?s(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):s(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function u(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},s=Object.keys(e);for(a=0;a<s.length;a++)n=s[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(a=0;a<s.length;a++)n=s[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var i=a.createContext({}),l=function(e){var t=a.useContext(i),n=t;return e&&(n="function"==typeof e?e(t):o(o({},t),e)),n},p=function(e){var t=l(e.components);return a.createElement(i.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},m=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,s=e.originalType,i=e.parentName,p=u(e,["components","mdxType","originalType","parentName"]),m=l(n),d=r,k=m["".concat(i,".").concat(d)]||m[d]||c[d]||s;return n?a.createElement(k,o(o({ref:t},p),{},{components:n})):a.createElement(k,o({ref:t},p))}));function d(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var s=n.length,o=new Array(s);o[0]=m;var u={};for(var i in t)hasOwnProperty.call(t,i)&&(u[i]=t[i]);u.originalType=e,u.mdxType="string"==typeof e?e:r,o[1]=u;for(var l=2;l<s;l++)o[l]=n[l];return a.createElement.apply(null,o)}return a.createElement.apply(null,n)}m.displayName="MDXCreateElement"},4653:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>i,contentTitle:()=>o,default:()=>c,frontMatter:()=>s,metadata:()=>u,toc:()=>l});var a=n(7462),r=(n(7294),n(3905));const s={sidebar_position:7,sidebar_label:"Maven"},o="Maven-based Tests",u={unversionedId:"test-types/executor-maven",id:"test-types/executor-maven",title:"Maven-based Tests",description:"Testkube allows us to run Maven-based tasks which could be also tests. For example we can easily run JUnit tests in Testkube now.",source:"@site/docs/4-test-types/executor-maven.md",sourceDirName:"4-test-types",slug:"/test-types/executor-maven",permalink:"/testkube/test-types/executor-maven",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/4-test-types/executor-maven.md",tags:[],version:"current",sidebarPosition:7,frontMatter:{sidebar_position:7,sidebar_label:"Maven"},sidebar:"tutorialSidebar",previous:{title:"Artillery.io",permalink:"/testkube/test-types/executor-artillery"},next:{title:"Gradle",permalink:"/testkube/test-types/executor-gradle"}},i={},l=[{value:"<strong>Test Environment</strong>",id:"test-environment",level:2},{value:"<strong>Create a New Maven-based Test</strong>",id:"create-a-new-maven-based-test",level:2},{value:"<strong>Running a Test</strong>",id:"running-a-test",level:2},{value:"<strong>Getting Test Results</strong>",id:"getting-test-results",level:2},{value:"Using different JDKs",id:"using-different-jdks",level:2},{value:"<strong>Summary</strong>",id:"summary",level:2}],p={toc:l};function c(e){let{components:t,...n}=e;return(0,r.kt)("wrapper",(0,a.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"maven-based-tests"},"Maven-based Tests"),(0,r.kt)("p",null,"Testkube allows us to run Maven-based tasks which could be also tests. For example we can easily run JUnit tests in Testkube now. "),(0,r.kt)("h2",{id:"test-environment"},(0,r.kt)("strong",{parentName:"h2"},"Test Environment")),(0,r.kt)("p",null,"We'll try to put simple JUnit test to our cluster and run it. Testkube Maven Executor handles ",(0,r.kt)("inlineCode",{parentName:"p"},"mvn")," and ",(0,r.kt)("inlineCode",{parentName:"p"},"mvnw")," binaries.\nBecause Maven projects are quite complicated in terms of directory structure. We'll need to load them from a Git directory."),(0,r.kt)("p",null,"You can find example projects in the repository here: ",(0,r.kt)("a",{parentName:"p",href:"https://github.com/kubeshop/testkube-executor-maven/tree/main/examples"},"https://github.com/kubeshop/testkube-executor-maven/tree/main/examples"),"."),(0,r.kt)("p",null,"Let's create a simple test which will check if an env variable is set to true: "),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-java"},'package hello.maven;\n\nimport org.junit.jupiter.api.Test;\nimport static org.junit.jupiter.api.Assertions.*;\n\nclass LibraryTest {\n    @Test void someLibraryMethodReturnsTrue() {\n        String env = System.getenv("TESTKUBE_MAVEN");\n        assertTrue(Boolean.parseBoolean(env), "TESTKUBE_MAVEN env should be true");\n    }\n}\n')),(0,r.kt)("p",null,"The default Maven executor: "),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},"apiVersion: executor.testkube.io/v1\nkind: Executor\nmetadata:\n  name: maven-executor\n  namespace: testkube\nspec:\n  image: kubeshop/testkube-maven-executor:0.1.4\n  types:\n  - maven/project\n  - maven/test\n  - maven/integration-test \n")),(0,r.kt)("p",null,"As we can see, there are several types. The Maven executor handles the second part after ",(0,r.kt)("inlineCode",{parentName:"p"},"/")," as a task name, so ",(0,r.kt)("inlineCode",{parentName:"p"},"maven/test")," will run ",(0,r.kt)("inlineCode",{parentName:"p"},"mvn test")," and so on. "),(0,r.kt)("p",null,"One exception from this rule is ",(0,r.kt)("inlineCode",{parentName:"p"},"project")," which is a generic one and forces you to pass additional arguments during test execution. For example:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube run maven-example-project --args='runMyCustomTask' \n")),(0,r.kt)("h2",{id:"create-a-new-maven-based-test"},(0,r.kt)("strong",{parentName:"h2"},"Create a New Maven-based Test")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube create test --git-uri https://github.com/kubeshop/testkube-executor-maven.git --git-path examples/hello-maven --type maven/test --name maven-example-test --git-branch main\n")),(0,r.kt)("h2",{id:"running-a-test"},(0,r.kt)("strong",{parentName:"h2"},"Running a Test")),(0,r.kt)("p",null,"Let's pass the env variable to our test run:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube run test maven-example-test -f -v TESTKUBE_MAVEN=true\n\n# ...... after some time\n\nTest execution completed with success in 16.555s \ud83e\udd47\n\nWatch the test execution until complete:\n$ kubectl testkube watch execution 62d148db0260f256c1a1e993\n\n\nUse the following command to get test execution details:\n$ kubectl testkube get execution 62d148db0260f256c1a1e993\n")),(0,r.kt)("h2",{id:"getting-test-results"},(0,r.kt)("strong",{parentName:"h2"},"Getting Test Results")),(0,r.kt)("p",null,"Now we can watch/get test execution details:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl testkube get execution 62d148db0260f256c1a1e993\n")),(0,r.kt)("p",null,"Output:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"# ....... a lot of Maven logs\n\nDownloaded from central: https://repo.maven.apache.org/maven2/org/junit/platform/junit-platform-launcher/1.7.2/junit-platform-launcher-1.7.2.pom (3.0 kB at 121 kB/s)\n[INFO] \n[INFO] -------------------------------------------------------\n[INFO]  T E S T S\n[INFO] -------------------------------------------------------\n[INFO] Running hello.maven.LibraryTest\n[INFO] Tests run: 1, Failures: 0, Errors: 0, Skipped: 0, Time elapsed: 0.052 s - in hello.maven.LibraryTest\n[INFO] \n[INFO] Results:\n[INFO] \n[INFO] Tests run: 1, Failures: 0, Errors: 0, Skipped: 0\n[INFO] \n[INFO] ------------------------------------------------------------------------\n[INFO] BUILD SUCCESS\n[INFO] ------------------------------------------------------------------------\n[INFO] Total time:  9.851 s\n[INFO] Finished at: 2022-07-18T09:06:15Z\n[INFO] ------------------------------------------------------------------------\n\nStatus Test execution completed with success \ud83e\udd47\n")),(0,r.kt)("h2",{id:"using-different-jdks"},"Using different JDKs"),(0,r.kt)("p",null,"In the Java world, usually you want to have control over your Runtime environment. Testkube can easily handle that for you!\nWe're building several Java images to handle constraints which Maven can put in it's build file."),(0,r.kt)("p",null,"To use a different executor you can use one of our pre-built ones (for Java 8,11,17,18) or build your own Docker image based on a Maven executor."),(0,r.kt)("p",null,"Let's assume we need JDK18 for our test runs. To handle that issue, create a new Maven executor:"),(0,r.kt)("p",null,"content of ",(0,r.kt)("inlineCode",{parentName:"p"},"maven-jdk18-executor.yaml")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'apiVersion: executor.testkube.io/v1\nkind: Executor\nmetadata:\n  name: maven-jdk18-executor\n  namespace: testkube\nspec:\n  image: kubeshop/testkube-maven-executor:0.1.0-jdk18   # <-- we\'re building jdk\n  types:\n  - maven:jdk18/project # <-- just create different test type with naming convention "framework:version/type"\n  - maven:jdk18/test\n  - maven:jdk18/integration-test \n')),(0,r.kt)("blockquote",null,(0,r.kt)("p",{parentName:"blockquote"},"Tip: Look for recent executor versions here ",(0,r.kt)("a",{parentName:"p",href:"https://hub.docker.com/repository/registry-1.docker.io/kubeshop/testkube-maven-executor/tags?page=1&ordering=last_updated"},"https://hub.docker.com/repository/registry-1.docker.io/kubeshop/testkube-maven-executor/tags?page=1&ordering=last_updated"),".")),(0,r.kt)("p",null,"And add it to your cluster: "),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl apply -f maven-jdk18-executor.yaml \n")),(0,r.kt)("p",null,"Now, create a new test with a type which our new executor can handle e.g.: ",(0,r.kt)("inlineCode",{parentName:"p"},"maven:jdk18/test")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"}," # create test\n kubectl testkube create test --git-uri https://github.com/kubeshop/testkube-executor-maven.git --git-path examples/hello-maven-jdk18 --type maven:jdk18/test --name maven-jdk18-example-test --git-branch main\n\n# and run it\nkubectl testkube run test maven-jdk18-example-test -f -v TESTKUBE_MAVEN=true\n")),(0,r.kt)("h2",{id:"summary"},(0,r.kt)("strong",{parentName:"h2"},"Summary")),(0,r.kt)("p",null,"Testkube simplifies running Java tests based on Maven and simplifies the merging of Java based tests into your global testing ecosystem."))}c.isMDXComponent=!0}}]);