"use strict";(self.webpackChunktestkube_documentation=self.webpackChunktestkube_documentation||[]).push([[809],{3905:(e,t,a)=>{a.d(t,{Zo:()=>u,kt:()=>d});var n=a(7294);function r(e,t,a){return t in e?Object.defineProperty(e,t,{value:a,enumerable:!0,configurable:!0,writable:!0}):e[t]=a,e}function l(e,t){var a=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),a.push.apply(a,n)}return a}function s(e){for(var t=1;t<arguments.length;t++){var a=null!=arguments[t]?arguments[t]:{};t%2?l(Object(a),!0).forEach((function(t){r(e,t,a[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(a)):l(Object(a)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(a,t))}))}return e}function o(e,t){if(null==e)return{};var a,n,r=function(e,t){if(null==e)return{};var a,n,r={},l=Object.keys(e);for(n=0;n<l.length;n++)a=l[n],t.indexOf(a)>=0||(r[a]=e[a]);return r}(e,t);if(Object.getOwnPropertySymbols){var l=Object.getOwnPropertySymbols(e);for(n=0;n<l.length;n++)a=l[n],t.indexOf(a)>=0||Object.prototype.propertyIsEnumerable.call(e,a)&&(r[a]=e[a])}return r}var p=n.createContext({}),i=function(e){var t=n.useContext(p),a=t;return e&&(a="function"==typeof e?e(t):s(s({},t),e)),a},u=function(e){var t=i(e.components);return n.createElement(p.Provider,{value:t},e.children)},m={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},k=n.forwardRef((function(e,t){var a=e.components,r=e.mdxType,l=e.originalType,p=e.parentName,u=o(e,["components","mdxType","originalType","parentName"]),k=i(a),d=r,b=k["".concat(p,".").concat(d)]||k[d]||m[d]||l;return a?n.createElement(b,s(s({ref:t},u),{},{components:a})):n.createElement(b,s({ref:t},u))}));function d(e,t){var a=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var l=a.length,s=new Array(l);s[0]=k;var o={};for(var p in t)hasOwnProperty.call(t,p)&&(o[p]=t[p]);o.originalType=e,o.mdxType="string"==typeof e?e:r,s[1]=o;for(var i=2;i<l;i++)s[i]=a[i];return n.createElement.apply(null,s)}return n.createElement.apply(null,a)}k.displayName="MDXCreateElement"},9285:(e,t,a)=>{a.r(t),a.d(t,{assets:()=>p,contentTitle:()=>s,default:()=>m,frontMatter:()=>l,metadata:()=>o,toc:()=>i});var n=a(7462),r=(a(7294),a(3905));const l={sidebar_position:1,sidebar_label:"Installation"},s="Installation Steps",o={unversionedId:"installing",id:"installing",title:"Installation Steps",description:"To get Testkube up and running you need to:",source:"@site/docs/1-installing.md",sourceDirName:".",slug:"/installing",permalink:"/testkube/installing",draft:!1,editUrl:"https://github.com/kubeshop/testkube/docs/docs/1-installing.md",tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1,sidebar_label:"Installation"},sidebar:"tutorialSidebar",previous:{title:"Welcome",permalink:"/testkube/"},next:{title:"Getting Started",permalink:"/testkube/getting-started"}},p={},i=[{value:"<strong>Installing the Testkube CLI</strong>",id:"installing-the-testkube-cli",level:2},{value:"<strong>From Scripts</strong>",id:"from-scripts",level:3},{value:"<strong>Through Package Managers</strong>",id:"through-package-managers",level:3},{value:"<strong>Homebrew (MacOS)</strong>",id:"homebrew-macos",level:4},{value:"<strong>Chocolatey (Windows)</strong>",id:"chocolatey-windows",level:4},{value:"<strong>APT (Debian/Ubuntu)</strong>",id:"apt-debianubuntu",level:4},{value:"<strong>Manual Download</strong>",id:"manual-download",level:3},{value:"<strong>Testkube Server Components</strong>",id:"testkube-server-components",level:3},{value:"<strong>Using Testkube&#39;s CLI to Deploy the Server Components</strong>",id:"using-testkubes-cli-to-deploy-the-server-components",level:3},{value:"<strong>Using HELM to Deploy the Server Components</strong>",id:"using-helm-to-deploy-the-server-components",level:3},{value:"<strong>Helm Properties</strong>",id:"helm-properties",level:4},{value:"<strong>Remove Testkube Server Components</strong>",id:"remove-testkube-server-components",level:2},{value:"<strong>Using Helm:</strong>",id:"using-helm",level:3},{value:"<strong>Using Testkube&#39;s CLI:</strong>",id:"using-testkubes-cli",level:3}],u={toc:i};function m(e){let{components:t,...a}=e;return(0,r.kt)("wrapper",(0,n.Z)({},u,a,{components:t,mdxType:"MDXLayout"}),(0,r.kt)("h1",{id:"installation-steps"},"Installation Steps"),(0,r.kt)("p",null,"To get Testkube up and running you need to:"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"Install the Testkube CLI."),(0,r.kt)("li",{parentName:"ol"},"Use HELM or the Testkube CLI to to install Testkube Server components in your cluster."),(0,r.kt)("li",{parentName:"ol"},"(optional) Configure Testkube's Dashboard UI Ingress for your ingress-controller, if needed.")),(0,r.kt)("p",null,"Watch the full installation video from our product experts: ",(0,r.kt)("a",{parentName:"p",href:"https://www.youtube.com/watch?v=bjQboi3Etys"},"Testkube Installation Video"),"."),(0,r.kt)("h2",{id:"installing-the-testkube-cli"},(0,r.kt)("strong",{parentName:"h2"},"Installing the Testkube CLI")),(0,r.kt)("p",null,"Package dependencies:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"https://kubernetes.io/docs/tasks/tools/"},"Kubectl")),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"https://helm.sh/"},"Helm"))),(0,r.kt)("p",null,"Installing the Testkube CLI with Chocolatey and Homebrew will automatically install these dependencies if they are not present. For Linux-based systems please install them manually in advance."),(0,r.kt)("h3",{id:"from-scripts"},(0,r.kt)("strong",{parentName:"h3"},"From Scripts")),(0,r.kt)("p",null,"To install on Linux or MacOs, run"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"bash < <(curl -sSLf https://kubeshop.github.io/testkube/install.sh )\n")),(0,r.kt)("h3",{id:"through-package-managers"},(0,r.kt)("strong",{parentName:"h3"},"Through Package Managers")),(0,r.kt)("h4",{id:"homebrew-macos"},(0,r.kt)("strong",{parentName:"h4"},"Homebrew (MacOS)")),(0,r.kt)("p",null,"You can install Testkube from ",(0,r.kt)("a",{parentName:"p",href:"https://brew.sh/"},"Homebrew"),":"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"brew install testkube\n")),(0,r.kt)("p",null,"Or directly from our tap. The Homebrew mantainers take a few days or a week to approve each one of our releases so you can use our tap to make sure you always have the most recent release."),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"brew tap kubeshop/homebrew-testkube\nbrew install kubeshop/testkube/testkube\n")),(0,r.kt)("h4",{id:"chocolatey-windows"},(0,r.kt)("strong",{parentName:"h4"},"Chocolatey (Windows)")),(0,r.kt)("p",null,(0,r.kt)("strong",{parentName:"p"},"Using ",(0,r.kt)("a",{parentName:"strong",href:"https://chocolatey.org/install"},"Chocolatey"),":")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"choco source add --name=testkube_repo --source=http://chocolatey.testkube.io/chocolatey\nchoco install testkube\n")),(0,r.kt)("h4",{id:"apt-debianubuntu"},(0,r.kt)("strong",{parentName:"h4"},"APT (Debian/Ubuntu)")),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"Download our public GPG key, and add it to the trusted keys:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"wget -qO - https://repo.testkube.io/key.pub | sudo apt-key add -\n")),(0,r.kt)("ol",{start:2},(0,r.kt)("li",{parentName:"ol"},"Add our repository to your apt sources:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},'echo "deb https://repo.testkube.io/linux linux main" | sudo tee -a /etc/apt/sources.list\n')),(0,r.kt)("ol",{start:3},(0,r.kt)("li",{parentName:"ol"},"Make sure to get the updates:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"sudo apt-get update\n")),(0,r.kt)("ol",{start:4},(0,r.kt)("li",{parentName:"ol"},"Install Testkube:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"sudo apt-get install -y testkube\n")),(0,r.kt)("h3",{id:"manual-download"},(0,r.kt)("strong",{parentName:"h3"},"Manual Download")),(0,r.kt)("p",null,"If you don't want to use scripts or package managers you can always do a manual install:"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"Download the binary for the version and platform of your choice ",(0,r.kt)("a",{parentName:"li",href:"https://github.com/kubeshop/testkube/releases"},"here")),(0,r.kt)("li",{parentName:"ol"},"Unpack it. For example, in Linux use (tar -zxvf testkube_1.5.1_Linux_arm64.tar.gz)"),(0,r.kt)("li",{parentName:"ol"},"Move it to a location in the PATH. For example, ",(0,r.kt)("inlineCode",{parentName:"li"},"mv  testkube_0.6.5_Linux_arm64/kubectl-testkube /usr/local/bin/kubectl-testkube"),".")),(0,r.kt)("p",null,"For Windows, you will need to unpack the binary and add it to the ",(0,r.kt)("inlineCode",{parentName:"p"},"%PATH%")," as well."),(0,r.kt)("p",null,"If you use a package manager that we don't support, please let us know here ",(0,r.kt)("a",{parentName:"p",href:"https://github.com/kubeshop/testkube/issues/161"},"#161"),"."),(0,r.kt)("h3",{id:"testkube-server-components"},(0,r.kt)("strong",{parentName:"h3"},"Testkube Server Components")),(0,r.kt)("p",null,"To deploy Testkube to your K8s cluster you will need the following packages installed:"),(0,r.kt)("ul",null,(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"https://kubernetes.io/docs/tasks/tools/"},"Kubectl docs")," "),(0,r.kt)("li",{parentName:"ul"},(0,r.kt)("a",{parentName:"li",href:"https://helm.sh/docs/intro/install/#through-package-managers"},"Helm docs"))),(0,r.kt)("h3",{id:"using-testkubes-cli-to-deploy-the-server-components"},(0,r.kt)("strong",{parentName:"h3"},"Using Testkube's CLI to Deploy the Server Components")),(0,r.kt)("p",null,"The Testkube CLI provides a command to easly deploy the Testkube server components to your cluster.\nRun:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"testkube init\n")),(0,r.kt)("p",null,"note: you must have your KUBECONFIG pointing to the desired location of the installation."),(0,r.kt)("p",null,"The above command will install the following components in your Kubernetes cluster:"),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"Testkube API"),(0,r.kt)("li",{parentName:"ol"},(0,r.kt)("inlineCode",{parentName:"li"},"testkube")," namespace"),(0,r.kt)("li",{parentName:"ol"},"CRDs for Tests, TestSuites, Executors"),(0,r.kt)("li",{parentName:"ol"},"MongoDB"),(0,r.kt)("li",{parentName:"ol"},"Minio - default (can be disabled with ",(0,r.kt)("inlineCode",{parentName:"li"},"--no-minio")," flag if you want to use S3 buckets)"),(0,r.kt)("li",{parentName:"ol"},"Dashboard - default (can be disabled with ",(0,r.kt)("inlineCode",{parentName:"li"},"--no-dasboard")," flag)")),(0,r.kt)("p",null,"Confirm that Testkube is running:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"kubectl get all -n testkube\n")),(0,r.kt)("p",null,"Output:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"NAME                                           READY   STATUS    RESTARTS   AGE\npod/cert-manager-847544bbd-fw2h8               1/1     Running   0          4m51s\npod/cert-manager-cainjector-5c747645bf-qgftx   1/1     Running   0          4m51s\npod/cert-manager-webhook-77b946cb6d-dl6gb      1/1     Running   0          4m51s\npod/testkube-dashboard-748cbcbb66-q8zzp        1/1     Running   0          4m51s\npod/testkube-api-server-546777c9f7-7g4kg       1/1     Running   0          4m51s\npod/testkube-mongodb-5d95f44fdd-cxqz6          1/1     Running   0          4m51s\npod/testkube-minio-testkube-64cd475b94-562hz   1/1     Running   0          4m51s\n\nNAME                                      TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                        AGE\nservice/cert-manager                      ClusterIP   10.106.81.214   <none>        9402/TCP                                       2d20h\nservice/cert-manager-webhook              ClusterIP   10.104.228.254  <none>        443/TCP                                        2d20h\nservice/testkube-minio-service-testkube   NodePort    10.43.121.107   <none>        9000:31222/TCP,9090:32002/TCP,9443:32586/TCP   4m51s\nservice/testkube-api-server               NodePort    10.43.66.13     <none>        8088:32203/TCP                                 4m51s\nservice/testkube-mongodb                  ClusterIP   10.43.126.230   <none>        27017/TCP                                      4m51s\nservice/testkube-dashboard                NodePort    10.43.136.34    <none>        80:31991/TCP                                   4m51s\n\nNAME                                      READY   UP-TO-DATE   AVAILABLE   AGE\ndeployment.apps/cert-manager              1/1     1            1           4m51s\ndeployment.apps/cert-manager-cainjector   1/1     1            1           4m51s\ndeployment.apps/cert-manager-webhook      1/1     1            1           4m51s\ndeployment.apps/testkube-dashboard        1/1     1            1           4m51s\ndeployment.apps/testkube-api-server       1/1     1            1           4m51s\ndeployment.apps/testkube-mongodb          1/1     1            1           4m51s\ndeployment.apps/testkube-minio-testkube   1/1     1            1           4m51s\n\nNAME                                                 DESIRED   CURRENT   READY   AGE\nreplicaset.apps/cert-manager-847544bbd               1         1         1       4m51s\nreplicaset.apps/cert-manager-cainjector-5c747645bf   1         1         1       4m51s\nreplicaset.apps/cert-manager-webhook-77b946cb6d      1         1         1       4m51s\nreplicaset.apps/testkube-dashboard-748cbcbb66        1         1         1       4m51s\nreplicaset.apps/testkube-api-server-546777c9f7       1         1         1       4m51s\nreplicaset.apps/testkube-mongodb-5d95f44fdd          1         1         1       4m51s\nreplicaset.apps/testkube-minio-testkube-64cd475b94   1         1         1       4m51s\n")),(0,r.kt)("p",null,"By default Testkube is installed in the ",(0,r.kt)("inlineCode",{parentName:"p"},"testkube")," namespace."),(0,r.kt)("h3",{id:"using-helm-to-deploy-the-server-components"},(0,r.kt)("strong",{parentName:"h3"},"Using HELM to Deploy the Server Components")),(0,r.kt)("ol",null,(0,r.kt)("li",{parentName:"ol"},"Add the Kubeshop Helm repository as follows:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"helm repo add testkube https://kubeshop.github.io/helm-charts\n")),(0,r.kt)("p",null,"If this repo already exists, run ",(0,r.kt)("inlineCode",{parentName:"p"},"helm repo update")," to retrieve\nthe ",(0,r.kt)("inlineCode",{parentName:"p"},"latest")," versions of the packages.  You can then run ",(0,r.kt)("inlineCode",{parentName:"p"},"helm search repo\ntestkube")," to see the charts."),(0,r.kt)("ol",{start:2},(0,r.kt)("li",{parentName:"ol"},"To install the ",(0,r.kt)("inlineCode",{parentName:"li"},"testkube")," chart:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"helm install --create-namespace my-testkube testkube/testkube\n")),(0,r.kt)("p",null,"Please note that, by default, the namespace for the intstallation will be ",(0,r.kt)("inlineCode",{parentName:"p"},"testkube"),". If the ",(0,r.kt)("inlineCode",{parentName:"p"},"testkube")," namespace does not exist, it will be created for you."),(0,r.kt)("p",null,"If you wish to install into a different namespace, please use following command:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"helm install --namespace namespace_name my-testkube testkube/testkube\n")),(0,r.kt)("p",null,"To uninstall the ",(0,r.kt)("inlineCode",{parentName:"p"},"testkube")," chart if it was installed into default namespace:"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"helm delete my-testkube testkube/testkube\n")),(0,r.kt)("p",null,"And from a namespace other than ",(0,r.kt)("inlineCode",{parentName:"p"},"testkube"),":"),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"helm delete --namespace namespace_name my-testkube testkube/testkube\n")),(0,r.kt)("h4",{id:"helm-properties"},(0,r.kt)("strong",{parentName:"h4"},"Helm Properties")),(0,r.kt)("p",null,"The following Helm defaults are used in the ",(0,r.kt)("inlineCode",{parentName:"p"},"testkube")," chart:"),(0,r.kt)("table",null,(0,r.kt)("thead",{parentName:"table"},(0,r.kt)("tr",{parentName:"thead"},(0,r.kt)("th",{parentName:"tr",align:null},"Parameter"),(0,r.kt)("th",{parentName:"tr",align:null},"Is optional"),(0,r.kt)("th",{parentName:"tr",align:null},"Default"))),(0,r.kt)("tbody",{parentName:"table"},(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"mongodb.auth.enabled"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},"false")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"mongodb.service.port"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'"27017"')),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"mongodb.service.portName"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'"mongodb"')),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"mongodb.service.nodePort"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},"true")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"mongodb.service.clusterIP"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'""')),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"mongodb.nameOverride"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'"mongodb"')),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"mongodb.fullnameOverride"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'"testkube-mongodb"')),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.image.repository"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'"kubeshop/testkube-api-server"')),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.image.pullPolicy"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'"Always"')),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.image.tag"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'"latest"')),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.service.type"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'"NodePort"')),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.service.port"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},"8088")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.mongoDSN"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'"mongodb://testkube-mongodb:27017"')),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.telemetryEnabled"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},"true")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.storage.endpoint"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},"testkube-minio-service-testkube:9000")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.storage.accessKeyId"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},"minio")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.storage.accessKey"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},"minio123")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.storage.scrapperEnabled"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},"true")),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.slackToken"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'""')),(0,r.kt)("tr",{parentName:"tbody"},(0,r.kt)("td",{parentName:"tr",align:null},"testkube-api.slackChannelId"),(0,r.kt)("td",{parentName:"tr",align:null},"yes"),(0,r.kt)("td",{parentName:"tr",align:null},'""')))),(0,r.kt)("blockquote",null,(0,r.kt)("p",{parentName:"blockquote"},"For more configuration parameters of ",(0,r.kt)("inlineCode",{parentName:"p"},"MongoDB")," chart please visit:\n",(0,r.kt)("a",{parentName:"p",href:"https://github.com/bitnami/charts/tree/master/bitnami/mongodb#parameters"},"https://github.com/bitnami/charts/tree/master/bitnami/mongodb#parameters"))),(0,r.kt)("h2",{id:"remove-testkube-server-components"},(0,r.kt)("strong",{parentName:"h2"},"Remove Testkube Server Components")),(0,r.kt)("h3",{id:"using-helm"},(0,r.kt)("strong",{parentName:"h3"},"Using Helm:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"helm delete testkube\n")),(0,r.kt)("h3",{id:"using-testkubes-cli"},(0,r.kt)("strong",{parentName:"h3"},"Using Testkube's CLI:")),(0,r.kt)("pre",null,(0,r.kt)("code",{parentName:"pre",className:"language-bash"},"testkube purge\n")))}m.isMDXComponent=!0}}]);