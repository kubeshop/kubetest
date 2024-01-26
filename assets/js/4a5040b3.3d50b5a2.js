"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[4219],{3905:(e,t,n)=>{n.d(t,{Zo:()=>l,kt:()=>m});var o=n(67294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function s(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,o)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?s(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):s(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,o,r=function(e,t){if(null==e)return{};var n,o,r={},s=Object.keys(e);for(o=0;o<s.length;o++)n=s[o],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(o=0;o<s.length;o++)n=s[o],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var u=o.createContext({}),c=function(e){var t=o.useContext(u),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},l=function(e){var t=c(e.components);return o.createElement(u.Provider,{value:t},e.children)},k={inlineCode:"code",wrapper:function(e){var t=e.children;return o.createElement(o.Fragment,{},t)}},p=o.forwardRef((function(e,t){var n=e.components,r=e.mdxType,s=e.originalType,u=e.parentName,l=i(e,["components","mdxType","originalType","parentName"]),p=c(n),m=r,d=p["".concat(u,".").concat(m)]||p[m]||k[m]||s;return n?o.createElement(d,a(a({ref:t},l),{},{components:n})):o.createElement(d,a({ref:t},l))}));function m(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var s=n.length,a=new Array(s);a[0]=p;var i={};for(var u in t)hasOwnProperty.call(t,u)&&(i[u]=t[u]);i.originalType=e,i.mdxType="string"==typeof e?e:r,a[1]=i;for(var c=2;c<s;c++)a[c]=n[c];return o.createElement.apply(null,a)}return o.createElement.apply(null,n)}p.displayName="MDXCreateElement"},38199:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>u,contentTitle:()=>a,default:()=>k,frontMatter:()=>s,metadata:()=>i,toc:()=>c});var o=n(87462),r=(n(67294),n(3905));const s={},a="Testkube CircleCI",i={unversionedId:"articles/circleci",id:"articles/circleci",title:"Testkube CircleCI",description:"The Testkube CircleCI integration facilitates the installation of Testkube and allows the execution of any Testkube CLI command within a CircleCI pipeline. This integration can be seamlessly incorporated into your CircleCI repositories to enhance your CI/CD workflows.",source:"@site/docs/articles/circleci.md",sourceDirName:"articles",slug:"/articles/circleci",permalink:"/articles/circleci",draft:!1,editUrl:"https://github.com/kubeshop/testkube/tree/develop/docs/docs/articles/circleci.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Testkube Jenkins",permalink:"/articles/jenkins"},next:{title:"Run Tests with GitHub Actions",permalink:"/articles/run-tests-with-github-actions"}},u={},c=[{value:"Testkube Pro",id:"testkube-pro",level:2},{value:"How to configure Testkube CLI action for Testkube Pro and run a test",id:"how-to-configure-testkube-cli-action-for-testkube-pro-and-run-a-test",level:3},{value:"Testkube OSS",id:"testkube-oss",level:2},{value:"How to configure Testkube CLI action for TK OSS and run a test",id:"how-to-configure-testkube-cli-action-for-tk-oss-and-run-a-test",level:3},{value:"How to configure Testkube CLI action for TK OSS and run a test",id:"how-to-configure-testkube-cli-action-for-tk-oss-and-run-a-test-1",level:3},{value:"How to connect to GKE (Google Kubernetes Engine) cluster and run a test",id:"how-to-connect-to-gke-google-kubernetes-engine-cluster-and-run-a-test",level:3}],l={toc:c};function k(e){let{components:t,...n}=e;return(0,r.kt)("wrapper",(0,o.Z)({},l,n,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"testkube-circleci"},"Testkube CircleCI"),(0,r.kt)("p",null,"The Testkube CircleCI integration facilitates the installation of Testkube and allows the execution of any ",(0,r.kt)("a",{parentName:"p",href:"https://docs.testkube.io/cli/testkube"},"Testkube CLI")," command within a CircleCI pipeline. This integration can be seamlessly incorporated into your CircleCI repositories to enhance your CI/CD workflows.\nThe integration offers a versatile approach to align with your pipeline requirements and is compatible with Testkube Pro, Testkube Enterprise, and the open-source Testkube platform. It enables CircleCI users to leverage the powerful features of Testkube directly within their CI/CD pipelines, ensuring efficient and flexible test execution."),(0,r.kt)("h2",{id:"testkube-pro"},"Testkube Pro"),(0,r.kt)("h3",{id:"how-to-configure-testkube-cli-action-for-testkube-pro-and-run-a-test"},"How to configure Testkube CLI action for Testkube Pro and run a test"),(0,r.kt)("p",null,"To use CircleCI for ",(0,r.kt)("a",{parentName:"p",href:"https://app.testkube.io/"},"Testkube Pro"),", you need to create an ",(0,r.kt)("a",{parentName:"p",href:"https://docs.testkube.io/testkube-pro/articles/organization-management/#api-tokens"},"API token"),".\nThen, pass the ",(0,r.kt)("strong",{parentName:"p"},"organization")," and ",(0,r.kt)("strong",{parentName:"p"},"environment")," IDs, along with the ",(0,r.kt)("strong",{parentName:"p"},"token")," and other parameters specific for your use case."),(0,r.kt)("p",null,"If a test is already created, you can run it using the command ",(0,r.kt)("inlineCode",{parentName:"p"},"testkube run test test-name -f")," . However, if you need to create a test in this workflow, please add a creation command, e.g.: ",(0,r.kt)("inlineCode",{parentName:"p"},"testkube create test --name test-name --file path_to_file.json"),"."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'version: 2.1\n\njobs:\n  run-tests:\n    docker:\n      - image: kubeshop/testkube-cli\n    working_directory: /.testkube\n    environment:\n      TESTKUBE_API_KEY: tkcapi_0123456789abcdef0123456789abcd\n      TESTKUBE_ORG_ID: tkcorg_0123456789abcdef\n      TESTKUBE_ENV_ID: tkcenv_fedcba9876543210\n    steps:\n      - run:\n          name: "Set Testkube Context"\n          command: "testkube set context --api-key $TESTKUBE_API_KEY --org $TESTKUBE_ORG_ID --env $TESTKUBE_ENV_ID --cloud-root-domain testkube.dev"\n      - run:\n          name: "Trigger testkube test"\n          command: "testkube run test test-name -f"\n\nworkflows:\n  run-tests-workflow:\n    jobs:\n      - run-tests\n')),(0,r.kt)("p",null,"It is recommended that sensitive values should never be stored as plaintext in workflow files, but rather as ",(0,r.kt)("a",{parentName:"p",href:"https://circleci.com/docs/set-environment-variable/#set-an-environment-variable-in-a-project"},"project variables"),".  Secrets can be configured at the organization or project level and allow you to store sensitive information in CircleCI."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'version: 2.1\n\njobs:\n  run-tests:\n    docker:\n      - image: kubeshop/testkube-cli\n    working_directory: /.testkube\n    steps:\n      - run:\n          name: "Set Testkube Context"\n          command: "testkube set context --api-key $TESTKUBE_API_KEY --org $TESTKUBE_ORG_ID --env $TESTKUBE_ENV_ID --cloud-root-domain testkube.dev"\n      - run:\n          name: "Trigger testkube test"\n          command: "testkube run test test-name -f"\n\nworkflows:\n  run-tests-workflow:\n    jobs:\n      - run-tests\n')),(0,r.kt)("h2",{id:"testkube-oss"},"Testkube OSS"),(0,r.kt)("h3",{id:"how-to-configure-testkube-cli-action-for-tk-oss-and-run-a-test"},"How to configure Testkube CLI action for TK OSS and run a test"),(0,r.kt)("p",null,"To connect to the self-hosted instance, you need to have ",(0,r.kt)("strong",{parentName:"p"},"kubectl")," configured for accessing your Kubernetes cluster and pass an optional namespace, if Testkube is not deployed in the default ",(0,r.kt)("strong",{parentName:"p"},"testkube")," namespace. "),(0,r.kt)("p",null,"If a test is already created, you can run it using the command ",(0,r.kt)("inlineCode",{parentName:"p"},"testkube run test test-name -f")," . However, if you need to create a test in this workflow, please add a creation command, e.g.: ",(0,r.kt)("inlineCode",{parentName:"p"},"testkube create test --name test-name --file path_to_file.json"),"."),(0,r.kt)("p",null,"In order to connect to your own cluster, you can put your kubeconfig file into CircleCI variable named KUBECONFIGFILE."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'version: 2.1\n\njobs:\n  run-tests:\n    docker:\n      - image: kubeshop/testkube-cli\n    working_directory: /.testkube\n    steps:\n      - run: \n          name: "Export kubeconfig"\n          command: |\n            echo $KUBECONFIGFILE > /.testkube/tmp/kubeconfig/config\n            export KUBECONFIG=/.testkube/tmp/kubeconfig/config\n      - run:\n          name: "Set Testkube Context"\n          command: "testkube set context --api-key $TESTKUBE_API_KEY --org $TESTKUBE_ORG_ID --env $TESTKUBE_ENV_ID --cloud-root-domain testkube.dev"\n      - run:\n          name: "Trigger testkube test"\n          command: "testkube run test test-name -f"\n\nworkflows:\n  run-tests-workflow:\n    jobs:\n      - run-tests\n')),(0,r.kt)("p",null,"The steps to connect to your Kubernetes cluster differ for each provider. You should check the docs of your Cloud provider for how to connect to the Kubernetes cluster from CircleCI."),(0,r.kt)("h3",{id:"how-to-configure-testkube-cli-action-for-tk-oss-and-run-a-test-1"},"How to configure Testkube CLI action for TK OSS and run a test"),(0,r.kt)("p",null,"This workflow establishes a connection to the EKS cluster and creates and runs a test using TK CLI. In this example we also use CircleCI variables not to reveal sensitive data. Please make sure that the following points are satisfied:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"The ",(0,r.kt)("strong",{parentName:"li"},(0,r.kt)("em",{parentName:"strong"},"AwsAccessKeyId")),", ",(0,r.kt)("strong",{parentName:"li"},(0,r.kt)("em",{parentName:"strong"},"AwsSecretAccessKeyId"))," secrets should contain your AWS IAM keys with proper permissions to connect to EKS cluster."),(0,r.kt)("li",{parentName:"ul"},"The ",(0,r.kt)("strong",{parentName:"li"},(0,r.kt)("em",{parentName:"strong"},"AwsRegion"))," secret should contain the AWS region where EKS is."),(0,r.kt)("li",{parentName:"ul"},"Tke ",(0,r.kt)("strong",{parentName:"li"},"EksClusterName")," secret points to the name of the EKS cluster you want to connect.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'version: 2.1\n\njobs:\n  setup-aws:\n    docker:\n      - image: amazon/aws-cli\n    steps:\n      - run:\n          name: "Configure AWS CLI"\n          command: |\n            mkdir -p /.testkube/tmp/kubeconfig/config\n            aws configure set aws_access_key_id $AWS_ACCESS_KEY_ID\n            aws configure set aws_secret_access_key $AWS_SECRET_ACCESS_KEY\n            aws configure set region $AWS_REGION\n            aws eks update-kubeconfig --name $EKS_CLUSTER_NAME --region $AWS_REGION --kubeconfig /.testkube/tmp/kubeconfig/config\n\n  run-testkube-on-aws:\n    docker:\n      - image: kubeshop/testkube-cli\n    working_directory: /.testkube\n    environment:\n        NAMESPACE: custom-testkube\n    steps:\n      - run:\n          name: "Run Testkube Test on EKS"\n          command: |\n            export KUBECONFIG=/.testkube/tmp/kubeconfig/config\n            testkube set context --kubeconfig --namespace $NAMESPACE\n            echo "Running Testkube test..."\n            testkube run test test-name -f\n\nworkflows:\n  aws-testkube-workflow:\n    jobs:\n      - setup-aws\n      - run-testkube-on-aws:\n          requires:\n            - setup-aws\n')),(0,r.kt)("h3",{id:"how-to-connect-to-gke-google-kubernetes-engine-cluster-and-run-a-test"},"How to connect to GKE (Google Kubernetes Engine) cluster and run a test"),(0,r.kt)("p",null,"This example connects to a k8s cluster in Google Cloud then creates and runs a test using Testkube CircleCI. Please make sure that the following points are satisfied:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},"The ",(0,r.kt)("strong",{parentName:"li"},(0,r.kt)("em",{parentName:"strong"},"GKE Sevice Account"))," should already be created in Google Cloud and added to CircleCI variables along with ",(0,r.kt)("strong",{parentName:"li"},(0,r.kt)("em",{parentName:"strong"},"GKE Project"))," value."),(0,r.kt)("li",{parentName:"ul"},"The ",(0,r.kt)("strong",{parentName:"li"},(0,r.kt)("em",{parentName:"strong"},"GKE Cluster Name"))," and ",(0,r.kt)("strong",{parentName:"li"},(0,r.kt)("em",{parentName:"strong"},"GKE Zone"))," can be added as environment variables in the workflow.")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-yaml"},'version: 2.1\n\njobs:\n  setup-gcp:\n    docker:\n      - image: google/cloud-sdk:latest\n    working_directory: /.testkube\n    steps:\n      - run:\n          name: "Setup GCP"\n          command: |\n            mkdir -p /.testkube/tmp/kubeconfig/config\n            export KUBECONFIG=$CI_PROJECT_DIR/tmp/kubeconfig/config\n            echo $GKE_SA_KEY | base64 -d > gke-sa-key.json\n            gcloud auth activate-service-account --key-file=gke-sa-key.json\n            gcloud config set project $GKE_PROJECT\n            gcloud --quiet auth configure-docker\n            gcloud container clusters get-credentials $GKE_CLUSTER_NAME --zone $GKE_ZONE\n\n  run-testkube-on-gcp:\n    docker:\n      - image: kubeshop/testkube-cli\n    working_directory: /.testkube\n    steps:\n      - run:\n          name: "Run Testkube Test on GKE"\n          command: |\n            export KUBECONFIG=/.testkube/tmp/kubeconfig/config\n            testkube set context --kubeconfig --namespace $NAMESPACE\n            testkube run test test-name -f\n\nworkflows:\n  gke-testkube-workflow:\n    jobs:\n      - setup-gcp\n      - run-testkube-on-gcp:\n          requires:\n            - setup-gcp\n')))}k.isMDXComponent=!0}}]);