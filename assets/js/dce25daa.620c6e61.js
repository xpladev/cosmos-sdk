"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[4012],{3905:(e,a,t)=>{t.d(a,{Zo:()=>d,kt:()=>c});var n=t(7294);function i(e,a,t){return a in e?Object.defineProperty(e,a,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[a]=t,e}function o(e,a){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);a&&(n=n.filter((function(a){return Object.getOwnPropertyDescriptor(e,a).enumerable}))),t.push.apply(t,n)}return t}function r(e){for(var a=1;a<arguments.length;a++){var t=null!=arguments[a]?arguments[a]:{};a%2?o(Object(t),!0).forEach((function(a){i(e,a,t[a])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(a){Object.defineProperty(e,a,Object.getOwnPropertyDescriptor(t,a))}))}return e}function l(e,a){if(null==e)return{};var t,n,i=function(e,a){if(null==e)return{};var t,n,i={},o=Object.keys(e);for(n=0;n<o.length;n++)t=o[n],a.indexOf(t)>=0||(i[t]=e[t]);return i}(e,a);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)t=o[n],a.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(i[t]=e[t])}return i}var s=n.createContext({}),p=function(e){var a=n.useContext(s),t=a;return e&&(t="function"==typeof e?e(a):r(r({},a),e)),t},d=function(e){var a=p(e.components);return n.createElement(s.Provider,{value:a},e.children)},m={inlineCode:"code",wrapper:function(e){var a=e.children;return n.createElement(n.Fragment,{},a)}},u=n.forwardRef((function(e,a){var t=e.components,i=e.mdxType,o=e.originalType,s=e.parentName,d=l(e,["components","mdxType","originalType","parentName"]),u=p(t),c=i,h=u["".concat(s,".").concat(c)]||u[c]||m[c]||o;return t?n.createElement(h,r(r({ref:a},d),{},{components:t})):n.createElement(h,r({ref:a},d))}));function c(e,a){var t=arguments,i=a&&a.mdxType;if("string"==typeof e||i){var o=t.length,r=new Array(o);r[0]=u;var l={};for(var s in a)hasOwnProperty.call(a,s)&&(l[s]=a[s]);l.originalType=e,l.mdxType="string"==typeof e?e:i,r[1]=l;for(var p=2;p<o;p++)r[p]=t[p];return n.createElement.apply(null,r)}return n.createElement.apply(null,t)}u.displayName="MDXCreateElement"},9142:(e,a,t)=>{t.r(a),t.d(a,{assets:()=>s,contentTitle:()=>r,default:()=>m,frontMatter:()=>o,metadata:()=>l,toc:()=>p});var n=t(7462),i=(t(7294),t(3905));const o={sidebar_position:1},r="Cosmovisor",l={unversionedId:"tooling/cosmovisor",id:"tooling/cosmovisor",title:"Cosmovisor",description:"cosmovisor is a small process manager for Cosmos SDK application binaries that monitors the governance module for incoming chain upgrade proposals. If it sees a proposal that gets approved, cosmovisor can automatically download the new binary, stop the current binary, switch from the old binary to the new one, and finally restart the node with the new binary.",source:"@site/docs/tooling/01-cosmovisor.md",sourceDirName:"tooling",slug:"/tooling/cosmovisor",permalink:"/main/tooling/cosmovisor",draft:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Protocol Buffers",permalink:"/main/tooling/protobuf"},next:{title:"Architecture Decision Records (ADR)",permalink:"/main/architecture/"}},s={},p=[{value:"Design",id:"design",level:2},{value:"Contributing",id:"contributing",level:2},{value:"Setup",id:"setup",level:2},{value:"Installation",id:"installation",level:3},{value:"Command Line Arguments And Environment Variables",id:"command-line-arguments-and-environment-variables",level:3},{value:"Folder Layout",id:"folder-layout",level:3},{value:"Usage",id:"usage",level:2},{value:"Initialization",id:"initialization",level:3},{value:"Detecting Upgrades",id:"detecting-upgrades",level:3},{value:"Auto-Download",id:"auto-download",level:3},{value:"Example: SimApp Upgrade",id:"example-simapp-upgrade",level:2},{value:"Chain Setup",id:"chain-setup",level:3},{value:"Prepare Cosmovisor and Start the Chain",id:"prepare-cosmovisor-and-start-the-chain",level:4},{value:"Update App",id:"update-app",level:4}],d={toc:p};function m(e){let{components:a,...t}=e;return(0,i.kt)("wrapper",(0,n.Z)({},d,t,{components:a,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"cosmovisor"},"Cosmovisor"),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," is a small process manager for Cosmos SDK application binaries that monitors the governance module for incoming chain upgrade proposals. If it sees a proposal that gets approved, ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," can automatically download the new binary, stop the current binary, switch from the old binary to the new one, and finally restart the node with the new binary."),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#design"},"Design")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#contributing"},"Contributing")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#setup"},"Setup"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#installation"},"Installation")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#command-line-arguments-and-environment-variables"},"Command Line Arguments And Environment Variables")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#folder-layout"},"Folder Layout")))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#usage"},"Usage"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#initialization"},"Initialization")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#detecting-upgrades"},"Detecting Upgrades")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#auto-download"},"Auto-Download")))),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#example-simapp-upgrade"},"Example: SimApp Upgrade"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#chain-setup"},"Chain Setup"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#prepare-cosmovisor-and-start-the-chain"},"Prepare Cosmovisor and Start the Chain")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"#update-app"},"Update App"))))))),(0,i.kt)("h2",{id:"design"},"Design"),(0,i.kt)("p",null,"Cosmovisor is designed to be used as a wrapper for a ",(0,i.kt)("inlineCode",{parentName:"p"},"Cosmos SDK")," app:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"it will pass arguments to the associated app (configured by ",(0,i.kt)("inlineCode",{parentName:"li"},"DAEMON_NAME")," env variable).\nRunning ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor run arg1 arg2 ....")," will run ",(0,i.kt)("inlineCode",{parentName:"li"},"app arg1 arg2 ..."),";"),(0,i.kt)("li",{parentName:"ul"},"it will manage an app by restarting and upgrading if needed;"),(0,i.kt)("li",{parentName:"ul"},"it is configured using environment variables, not positional arguments.")),(0,i.kt)("p",null,(0,i.kt)("em",{parentName:"p"},"Note: If new versions of the application are not set up to run in-place store migrations, migrations will need to be run manually before restarting ",(0,i.kt)("inlineCode",{parentName:"em"},"cosmovisor")," with the new binary. For this reason, we recommend applications adopt in-place store migrations.")),(0,i.kt)("p",null,(0,i.kt)("em",{parentName:"p"},"Note: If validators would like to enable the auto-download option (which ",(0,i.kt)("a",{parentName:"em",href:"#auto-download"},"we don't recommend"),"), and they are currently running an application using Cosmos SDK ",(0,i.kt)("inlineCode",{parentName:"em"},"v0.42"),", they will need to use Cosmovisor ",(0,i.kt)("a",{parentName:"em",href:"https://github.com/cosmos/cosmos-sdk/releases/tag/cosmovisor%2Fv0.1.0"},(0,i.kt)("inlineCode",{parentName:"a"},"v0.1")),". Later versions of Cosmovisor do not support Cosmos SDK ",(0,i.kt)("inlineCode",{parentName:"em"},"v0.44.3")," or earlier if the auto-download option is enabled.")),(0,i.kt)("h2",{id:"contributing"},"Contributing"),(0,i.kt)("p",null,"Cosmovisor is part of the Cosmos SDK monorepo, but it's a separate module with it's own release schedule."),(0,i.kt)("p",null,"Release branches have the following format ",(0,i.kt)("inlineCode",{parentName:"p"},"release/cosmovisor/vA.B.x"),", where A and B are a number (e.g. ",(0,i.kt)("inlineCode",{parentName:"p"},"release/cosmovisor/v1.3.x"),"). Releases are tagged using the following format: ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor/vA.B.C"),"."),(0,i.kt)("h2",{id:"setup"},"Setup"),(0,i.kt)("h3",{id:"installation"},"Installation"),(0,i.kt)("p",null,"You can download Cosmovisor from the ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/releases/tag/cosmovisor%2Fv1.3.0"},"GitHub releases"),"."),(0,i.kt)("p",null,"To install the latest version of ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor"),", run the following command:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@latest\n")),(0,i.kt)("p",null,"To install a previous version, you can specify the version. IMPORTANT: Chains that use Cosmos-SDK v0.44.3 or earlier (eg v0.44.2) and want to use auto-download feature MUST use ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor v0.1.0")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"go install github.com/cosmos/cosmos-sdk/cosmovisor/cmd/cosmovisor@v0.1.0\n")),(0,i.kt)("p",null,"Run ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor version")," to check the cosmovisor version."),(0,i.kt)("p",null,"You can also install from source by pulling the cosmos-sdk repository and switching to the correct version and building as follows:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"git clone git@github.com:cosmos/cosmos-sdk\ncd cosmos-sdk\ngit checkout cosmovisor/vx.x.x\nmake cosmovisor\n")),(0,i.kt)("p",null,"This will build cosmovisor in ",(0,i.kt)("inlineCode",{parentName:"p"},"/cosmovisor")," directory. Afterwards you may want to put it into your machine's PATH like as follows:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"cp cosmovisor/cosmovisor ~/go/bin/cosmovisor\n")),(0,i.kt)("p",null,(0,i.kt)("em",{parentName:"p"},"Note: If you are using go ",(0,i.kt)("inlineCode",{parentName:"em"},"v1.15")," or earlier, you will need to use ",(0,i.kt)("inlineCode",{parentName:"em"},"go get"),", and you may want to run the command outside a project directory.")),(0,i.kt)("h3",{id:"command-line-arguments-and-environment-variables"},"Command Line Arguments And Environment Variables"),(0,i.kt)("p",null,"The first argument passed to ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," is the action for ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," to take. Options are:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"help"),", ",(0,i.kt)("inlineCode",{parentName:"li"},"--help"),", or ",(0,i.kt)("inlineCode",{parentName:"li"},"-h")," - Output ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," help information and check your ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," configuration."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"run")," - Run the configured binary using the rest of the provided arguments."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"version")," - Output the ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," version and also run the binary with the ",(0,i.kt)("inlineCode",{parentName:"li"},"version")," argument.")),(0,i.kt)("p",null,"All arguments passed to ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor run")," will be passed to the application binary (as a subprocess). ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," will return ",(0,i.kt)("inlineCode",{parentName:"p"},"/dev/stdout")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"/dev/stderr")," of the subprocess as its own. For this reason, ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor run")," cannot accept any command-line arguments other than those available to the application binary."),(0,i.kt)("p",null,"*Note: Use of ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," without one of the action arguments is deprecated. For backwards compatibility, if the first argument is not an action argument, ",(0,i.kt)("inlineCode",{parentName:"p"},"run")," is assumed. However, this fallback might be removed in future versions, so it is recommended that you always provide ",(0,i.kt)("inlineCode",{parentName:"p"},"run"),"."),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," reads its configuration from environment variables:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"DAEMON_HOME")," is the location where the ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor/")," directory is kept that contains the genesis binary, the upgrade binaries, and any additional auxiliary files associated with each binary (e.g. ",(0,i.kt)("inlineCode",{parentName:"li"},"$HOME/.gaiad"),", ",(0,i.kt)("inlineCode",{parentName:"li"},"$HOME/.regend"),", ",(0,i.kt)("inlineCode",{parentName:"li"},"$HOME/.simd"),", etc.)."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"DAEMON_NAME")," is the name of the binary itself (e.g. ",(0,i.kt)("inlineCode",{parentName:"li"},"gaiad"),", ",(0,i.kt)("inlineCode",{parentName:"li"},"regend"),", ",(0,i.kt)("inlineCode",{parentName:"li"},"simd"),", etc.)."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"DAEMON_ALLOW_DOWNLOAD_BINARIES")," (",(0,i.kt)("em",{parentName:"li"},"optional"),"), if set to ",(0,i.kt)("inlineCode",{parentName:"li"},"true"),", will enable auto-downloading of new binaries (for security reasons, this is intended for full nodes rather than validators). By default, ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," will not auto-download new binaries."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"DAEMON_RESTART_AFTER_UPGRADE")," (",(0,i.kt)("em",{parentName:"li"},"optional"),", default = ",(0,i.kt)("inlineCode",{parentName:"li"},"true"),"), if ",(0,i.kt)("inlineCode",{parentName:"li"},"true"),", restarts the subprocess with the same command-line arguments and flags (but with the new binary) after a successful upgrade. Otherwise (",(0,i.kt)("inlineCode",{parentName:"li"},"false"),"), ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," stops running after an upgrade and requires the system administrator to manually restart it. Note restart is only after the upgrade and does not auto-restart the subprocess after an error occurs."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"DAEMON_RESTART_DELAY")," (",(0,i.kt)("em",{parentName:"li"},"optional"),", default none), allow a node operator to define a delay between the node halt (for upgrade) and backup by the specified time. The value must be a duration (e.g. ",(0,i.kt)("inlineCode",{parentName:"li"},"1s"),")."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"DAEMON_POLL_INTERVAL")," (",(0,i.kt)("em",{parentName:"li"},"optional"),", default 300 milliseconds), is the interval length for polling the upgrade plan file. The value must be a duration (e.g. ",(0,i.kt)("inlineCode",{parentName:"li"},"1s"),")."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"DAEMON_DATA_BACKUP_DIR")," option to set a custom backup directory. If not set, ",(0,i.kt)("inlineCode",{parentName:"li"},"DAEMON_HOME")," is used."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"UNSAFE_SKIP_BACKUP")," (defaults to ",(0,i.kt)("inlineCode",{parentName:"li"},"false"),"), if set to ",(0,i.kt)("inlineCode",{parentName:"li"},"true"),", upgrades directly without performing a backup. Otherwise (",(0,i.kt)("inlineCode",{parentName:"li"},"false"),", default) backs up the data before trying the upgrade. The default value of false is useful and recommended in case of failures and when a backup needed to rollback. We recommend using the default backup option ",(0,i.kt)("inlineCode",{parentName:"li"},"UNSAFE_SKIP_BACKUP=false"),"."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"DAEMON_PREUPGRADE_MAX_RETRIES")," (defaults to ",(0,i.kt)("inlineCode",{parentName:"li"},"0"),"). The maximum number of times to call ",(0,i.kt)("inlineCode",{parentName:"li"},"pre-upgrade")," in the application after exit status of ",(0,i.kt)("inlineCode",{parentName:"li"},"31"),". After the maximum number of retries, cosmovisor fails the upgrade.")),(0,i.kt)("h3",{id:"folder-layout"},"Folder Layout"),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"$DAEMON_HOME/cosmovisor")," is expected to belong completely to ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," and the subprocesses that are controlled by it. The folder content is organized as follows:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},".\n\u251c\u2500\u2500 current -> genesis or upgrades/<name>\n\u251c\u2500\u2500 genesis\n\u2502\xa0\xa0 \u2514\u2500\u2500 bin\n\u2502\xa0\xa0     \u2514\u2500\u2500 $DAEMON_NAME\n\u2514\u2500\u2500 upgrades\n    \u2514\u2500\u2500 <name>\n        \u251c\u2500\u2500 bin\n        \u2502\xa0\xa0 \u2514\u2500\u2500 $DAEMON_NAME\n        \u2514\u2500\u2500 upgrade-info.json\n")),(0,i.kt)("p",null,"The ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor/")," directory incudes a subdirectory for each version of the application (i.e. ",(0,i.kt)("inlineCode",{parentName:"p"},"genesis")," or ",(0,i.kt)("inlineCode",{parentName:"p"},"upgrades/<name>"),"). Within each subdirectory is the application binary (i.e. ",(0,i.kt)("inlineCode",{parentName:"p"},"bin/$DAEMON_NAME"),") and any additional auxiliary files associated with each binary. ",(0,i.kt)("inlineCode",{parentName:"p"},"current")," is a symbolic link to the currently active directory (i.e. ",(0,i.kt)("inlineCode",{parentName:"p"},"genesis")," or ",(0,i.kt)("inlineCode",{parentName:"p"},"upgrades/<name>"),"). The ",(0,i.kt)("inlineCode",{parentName:"p"},"name")," variable in ",(0,i.kt)("inlineCode",{parentName:"p"},"upgrades/<name>")," is the lowercased URI-encoded name of the upgrade as specified in the upgrade module plan. Note that the upgrade name path are normalized to be lowercased: for instance, ",(0,i.kt)("inlineCode",{parentName:"p"},"MyUpgrade")," is normalized to ",(0,i.kt)("inlineCode",{parentName:"p"},"myupgrade"),", and its path is ",(0,i.kt)("inlineCode",{parentName:"p"},"upgrades/myupgrade"),"."),(0,i.kt)("p",null,"Please note that ",(0,i.kt)("inlineCode",{parentName:"p"},"$DAEMON_HOME/cosmovisor")," only stores the ",(0,i.kt)("em",{parentName:"p"},"application binaries"),". The ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," binary itself can be stored in any typical location (e.g. ",(0,i.kt)("inlineCode",{parentName:"p"},"/usr/local/bin"),"). The application will continue to store its data in the default data directory (e.g. ",(0,i.kt)("inlineCode",{parentName:"p"},"$HOME/.gaiad"),") or the data directory specified with the ",(0,i.kt)("inlineCode",{parentName:"p"},"--home")," flag. ",(0,i.kt)("inlineCode",{parentName:"p"},"$DAEMON_HOME")," is independent of the data directory and can be set to any location. If you set ",(0,i.kt)("inlineCode",{parentName:"p"},"$DAEMON_HOME")," to the same directory as the data directory, you will end up with a configuation like the following:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-text"},".gaiad\n\u251c\u2500\u2500 config\n\u251c\u2500\u2500 data\n\u2514\u2500\u2500 cosmovisor\n")),(0,i.kt)("h2",{id:"usage"},"Usage"),(0,i.kt)("p",null,"The system administrator is responsible for:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"installing the ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," binary"),(0,i.kt)("li",{parentName:"ul"},"configuring the host's init system (e.g. ",(0,i.kt)("inlineCode",{parentName:"li"},"systemd"),", ",(0,i.kt)("inlineCode",{parentName:"li"},"launchd"),", etc.)"),(0,i.kt)("li",{parentName:"ul"},"appropriately setting the environmental variables"),(0,i.kt)("li",{parentName:"ul"},"creating the ",(0,i.kt)("inlineCode",{parentName:"li"},"<DAEMON_HOME>/cosmovisor")," directory"),(0,i.kt)("li",{parentName:"ul"},"creating the ",(0,i.kt)("inlineCode",{parentName:"li"},"<DAEMON_HOME>/cosmovisor/genesis/bin")," folder"),(0,i.kt)("li",{parentName:"ul"},"creating the ",(0,i.kt)("inlineCode",{parentName:"li"},"<DAEMON_HOME>/cosmovisor/upgrades/<name>/bin")," folders"),(0,i.kt)("li",{parentName:"ul"},"placing the different versions of the ",(0,i.kt)("inlineCode",{parentName:"li"},"<DAEMON_NAME>")," executable in the appropriate ",(0,i.kt)("inlineCode",{parentName:"li"},"bin")," folders.")),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," will set the ",(0,i.kt)("inlineCode",{parentName:"p"},"current")," link to point to ",(0,i.kt)("inlineCode",{parentName:"p"},"genesis")," at first start (i.e. when no ",(0,i.kt)("inlineCode",{parentName:"p"},"current")," link exists) and then handle switching binaries at the correct points in time so that the system administrator can prepare days in advance and relax at upgrade time."),(0,i.kt)("p",null,"In order to support downloadable binaries, a tarball for each upgrade binary will need to be packaged up and made available through a canonical URL. Additionally, a tarball that includes the genesis binary and all available upgrade binaries can be packaged up and made available so that all the necessary binaries required to sync a fullnode from start can be easily downloaded."),(0,i.kt)("p",null,"The ",(0,i.kt)("inlineCode",{parentName:"p"},"DAEMON")," specific code and operations (e.g. tendermint config, the application db, syncing blocks, etc.) all work as expected. The application binaries' directives such as command-line flags and environment variables also work as expected."),(0,i.kt)("h3",{id:"initialization"},"Initialization"),(0,i.kt)("p",null,"The ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor init <path to executable>")," command creates the folder structure required for using cosmovisor."),(0,i.kt)("p",null,"It does the following:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"creates the ",(0,i.kt)("inlineCode",{parentName:"li"},"<DAEMON_HOME>/cosmovisor")," folder if it doesn't yet exist"),(0,i.kt)("li",{parentName:"ul"},"creates the ",(0,i.kt)("inlineCode",{parentName:"li"},"<DAEMON_HOME>/cosmovisor/genesis/bin")," folder if it doesn't yet exist"),(0,i.kt)("li",{parentName:"ul"},"copies the provided executable file to ",(0,i.kt)("inlineCode",{parentName:"li"},"<DAEMON_HOME>/cosmovisor/genesis/bin/<DAEMON_NAME>")),(0,i.kt)("li",{parentName:"ul"},"creates the ",(0,i.kt)("inlineCode",{parentName:"li"},"current")," link, pointing to the ",(0,i.kt)("inlineCode",{parentName:"li"},"genesis")," folder")),(0,i.kt)("p",null,"It uses the ",(0,i.kt)("inlineCode",{parentName:"p"},"DAEMON_HOME")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"DAEMON_NAME")," environment variables for folder location and executable name."),(0,i.kt)("p",null,"The ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor init")," command is specifically for initializing cosmovisor, and should not be confused with a chain's ",(0,i.kt)("inlineCode",{parentName:"p"},"init")," command (e.g. ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor run init"),")."),(0,i.kt)("h3",{id:"detecting-upgrades"},"Detecting Upgrades"),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," is polling the ",(0,i.kt)("inlineCode",{parentName:"p"},"$DAEMON_HOME/data/upgrade-info.json")," file for new upgrade instructions. The file is created by the x/upgrade module in ",(0,i.kt)("inlineCode",{parentName:"p"},"BeginBlocker")," when an upgrade is detected and the blockchain reaches the upgrade height.\nThe following heuristic is applied to detect the upgrade:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"When starting, ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," doesn't know much about currently running upgrade, except the binary which is ",(0,i.kt)("inlineCode",{parentName:"li"},"current/bin/"),". It tries to read the ",(0,i.kt)("inlineCode",{parentName:"li"},"current/update-info.json")," file to get information about the current upgrade name."),(0,i.kt)("li",{parentName:"ul"},"If neither ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor/current/upgrade-info.json")," nor ",(0,i.kt)("inlineCode",{parentName:"li"},"data/upgrade-info.json")," exist, then ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," will wait for ",(0,i.kt)("inlineCode",{parentName:"li"},"data/upgrade-info.json")," file to trigger an upgrade."),(0,i.kt)("li",{parentName:"ul"},"If ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor/current/upgrade-info.json")," doesn't exist but ",(0,i.kt)("inlineCode",{parentName:"li"},"data/upgrade-info.json")," exists, then ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," assumes that whatever is in ",(0,i.kt)("inlineCode",{parentName:"li"},"data/upgrade-info.json")," is a valid upgrade request. In this case ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," tries immediately to make an upgrade according to the ",(0,i.kt)("inlineCode",{parentName:"li"},"name")," attribute in ",(0,i.kt)("inlineCode",{parentName:"li"},"data/upgrade-info.json"),"."),(0,i.kt)("li",{parentName:"ul"},"Otherwise, ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," waits for changes in ",(0,i.kt)("inlineCode",{parentName:"li"},"upgrade-info.json"),". As soon as a new upgrade name is recorded in the file, ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor")," will trigger an upgrade mechanism.")),(0,i.kt)("p",null,"When the upgrade mechanism is triggered, ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," will:"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},"if ",(0,i.kt)("inlineCode",{parentName:"li"},"DAEMON_ALLOW_DOWNLOAD_BINARIES")," is enabled, start by auto-downloading a new binary into ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor/<name>/bin")," (where ",(0,i.kt)("inlineCode",{parentName:"li"},"<name>")," is the ",(0,i.kt)("inlineCode",{parentName:"li"},"upgrade-info.json:name")," attribute);"),(0,i.kt)("li",{parentName:"ol"},"update the ",(0,i.kt)("inlineCode",{parentName:"li"},"current")," symbolic link to point to the new directory and save ",(0,i.kt)("inlineCode",{parentName:"li"},"data/upgrade-info.json")," to ",(0,i.kt)("inlineCode",{parentName:"li"},"cosmovisor/current/upgrade-info.json"),".")),(0,i.kt)("h3",{id:"auto-download"},"Auto-Download"),(0,i.kt)("p",null,"Generally, ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," requires that the system administrator place all relevant binaries on disk before the upgrade happens. However, for people who don't need such control and want an automated setup (maybe they are syncing a non-validating fullnode and want to do little maintenance), there is another option."),(0,i.kt)("p",null,(0,i.kt)("strong",{parentName:"p"},"NOTE: we don't recommend using auto-download")," because it doesn't verify in advance if a binary is available. If there will be any issue with downloading a binary, the cosmovisor will stop and won't restart an App (which could lead to a chain halt)."),(0,i.kt)("p",null,"If ",(0,i.kt)("inlineCode",{parentName:"p"},"DAEMON_ALLOW_DOWNLOAD_BINARIES")," is set to ",(0,i.kt)("inlineCode",{parentName:"p"},"true"),", and no local binary can be found when an upgrade is triggered, ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," will attempt to download and install the binary itself based on the instructions in the ",(0,i.kt)("inlineCode",{parentName:"p"},"info")," attribute in the ",(0,i.kt)("inlineCode",{parentName:"p"},"data/upgrade-info.json")," file. The files is constructed by the x/upgrade module and contains data from the upgrade ",(0,i.kt)("inlineCode",{parentName:"p"},"Plan")," object. The ",(0,i.kt)("inlineCode",{parentName:"p"},"Plan")," has an info field that is expected to have one of the following two valid formats to specify a download:"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"Store an os/architecture -> binary URI map in the upgrade plan info field as JSON under the ",(0,i.kt)("inlineCode",{parentName:"p"},'"binaries"')," key. For example:"),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-json"},'{\n  "binaries": {\n    "linux/amd64":"https://example.com/gaia.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f"\n  }\n}\n')),(0,i.kt)("p",{parentName:"li"},"You can include multiple binaries at once to ensure more than one environment will receive the correct binaries:"),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-json"},'{\n  "binaries": {\n    "linux/amd64":"https://example.com/gaia.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f",\n    "linux/arm64":"https://example.com/gaia.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f",\n    "darwin/amd64":"https://example.com/gaia.zip?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f"\n    }\n}\n')),(0,i.kt)("p",{parentName:"li"},"When submitting this as a proposal ensure there are no spaces. An example command using ",(0,i.kt)("inlineCode",{parentName:"p"},"gaiad")," could look like:"),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-shell"},'> gaiad tx gov submit-proposal software-upgrade Vega \\\n--title Vega \\\n--deposit 100uatom \\\n--upgrade-height 7368420 \\\n--upgrade-info \'{"binaries":{"linux/amd64":"https://github.com/cosmos/gaia/releases/download/v6.0.0-rc1/gaiad-v6.0.0-rc1-linux-amd64","linux/arm64":"https://github.com/cosmos/gaia/releases/download/v6.0.0-rc1/gaiad-v6.0.0-rc1-linux-arm64","darwin/amd64":"https://github.com/cosmos/gaia/releases/download/v6.0.0-rc1/gaiad-v6.0.0-rc1-darwin-amd64"}}\' \\\n--description "upgrade to Vega" \\\n--gas 400000 \\\n--from user \\\n--chain-id test \\\n--home test/val2 \\\n--node tcp://localhost:36657 \\\n--yes\n'))),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"Store a link to a file that contains all information in the above format (e.g. if you want to specify lots of binaries, changelog info, etc. without filling up the blockchain). For example:"),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-text"},"https://example.com/testnet-1001-info.json?checksum=sha256:deaaa99fda9407c4dbe1d04bd49bab0cc3c1dd76fa392cd55a9425be074af01e\n")))),(0,i.kt)("p",null,"When ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," is triggered to download the new binary, ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," will parse the ",(0,i.kt)("inlineCode",{parentName:"p"},'"binaries"')," field, download the new binary with ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/hashicorp/go-getter"},"go-getter"),", and unpack the new binary in the ",(0,i.kt)("inlineCode",{parentName:"p"},"upgrades/<name>")," folder so that it can be run as if it was installed manually."),(0,i.kt)("p",null,"Note that for this mechanism to provide strong security guarantees, all URLs should include a SHA 256/512 checksum. This ensures that no false binary is run, even if someone hacks the server or hijacks the DNS. ",(0,i.kt)("inlineCode",{parentName:"p"},"go-getter")," will always ensure the downloaded file matches the checksum if it is provided. ",(0,i.kt)("inlineCode",{parentName:"p"},"go-getter")," will also handle unpacking archives into directories (in this case the download link should point to a ",(0,i.kt)("inlineCode",{parentName:"p"},"zip")," file of all data in the ",(0,i.kt)("inlineCode",{parentName:"p"},"bin")," directory)."),(0,i.kt)("p",null,"To properly create a sha256 checksum on linux, you can use the ",(0,i.kt)("inlineCode",{parentName:"p"},"sha256sum")," utility. For example:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"sha256sum ./testdata/repo/zip_directory/autod.zip\n")),(0,i.kt)("p",null,"The result will look something like the following: ",(0,i.kt)("inlineCode",{parentName:"p"},"29139e1381b8177aec909fab9a75d11381cab5adf7d3af0c05ff1c9c117743a7"),"."),(0,i.kt)("p",null,"You can also use ",(0,i.kt)("inlineCode",{parentName:"p"},"sha512sum")," if you would prefer to use longer hashes, or ",(0,i.kt)("inlineCode",{parentName:"p"},"md5sum")," if you would prefer to use broken hashes. Whichever you choose, make sure to set the hash algorithm properly in the checksum argument to the URL."),(0,i.kt)("h2",{id:"example-simapp-upgrade"},"Example: SimApp Upgrade"),(0,i.kt)("p",null,"The following instructions provide a demonstration of ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmovisor")," using the simulation application (",(0,i.kt)("inlineCode",{parentName:"p"},"simapp"),") shipped with the Cosmos SDK's source code. The following commands are to be run from within the ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmos-sdk")," repository."),(0,i.kt)("h3",{id:"chain-setup"},"Chain Setup"),(0,i.kt)("p",null,"Let's create a new chain using the ",(0,i.kt)("inlineCode",{parentName:"p"},"v0.44")," version of simapp (the Cosmos SDK demo app):"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"git checkout v0.44.6\nmake build\n")),(0,i.kt)("p",null,"Clean ",(0,i.kt)("inlineCode",{parentName:"p"},"~/.simapp")," (never do this in a production environment):"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"./build/simd unsafe-reset-all\n")),(0,i.kt)("p",null,"Set up app config:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"./build/simd config chain-id test\n./build/simd config keyring-backend test\n./build/simd config broadcast-mode sync\n")),(0,i.kt)("p",null,"Initialize the node and overwrite any previous genesis file (never do this in a production environment):"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"./build/simd init test --chain-id test --overwrite\n")),(0,i.kt)("p",null,"Set the minimum gas price to ",(0,i.kt)("inlineCode",{parentName:"p"},"0stake")," in ",(0,i.kt)("inlineCode",{parentName:"p"},"~/.simapp/config/app.toml"),":"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},'minimum-gas-prices = "0stake"\n')),(0,i.kt)("p",null,"For the sake of this demonstration, amend ",(0,i.kt)("inlineCode",{parentName:"p"},"voting_period")," in ",(0,i.kt)("inlineCode",{parentName:"p"},"genesis.json")," to a reduced time of 20 seconds (",(0,i.kt)("inlineCode",{parentName:"p"},"20s"),"):"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"cat <<< $(jq '.app_state.gov.voting_params.voting_period = \"20s\"' $HOME/.simapp/config/genesis.json) > $HOME/.simapp/config/genesis.json\n")),(0,i.kt)("p",null,"Create a validator, and setup genesis transaction:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"./build/simd keys add validator\n./build/simd add-genesis-account validator 1000000000stake --keyring-backend test\n./build/simd gentx validator 1000000stake --chain-id test\n./build/simd collect-gentxs\n")),(0,i.kt)("h4",{id:"prepare-cosmovisor-and-start-the-chain"},"Prepare Cosmovisor and Start the Chain"),(0,i.kt)("p",null,"Set the required environment variables:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"export DAEMON_NAME=simd\nexport DAEMON_HOME=$HOME/.simapp\n")),(0,i.kt)("p",null,"Set the optional environment variable to trigger an automatic app restart:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"export DAEMON_RESTART_AFTER_UPGRADE=true\n")),(0,i.kt)("p",null,"Create the folder for the genesis binary and copy the ",(0,i.kt)("inlineCode",{parentName:"p"},"simd")," binary:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin\ncp ./build/simd $DAEMON_HOME/cosmovisor/genesis/bin\n")),(0,i.kt)("p",null,"Now you can run cosmovisor with simapp v0.44:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"cosmovisor run start\n")),(0,i.kt)("h4",{id:"update-app"},"Update App"),(0,i.kt)("p",null,"Update app to the latest version (e.g. v0.45)."),(0,i.kt)("p",null,"Next, we can add a migration - which is defined using ",(0,i.kt)("inlineCode",{parentName:"p"},"x/upgrade")," ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/main/docs/core/upgrade.md"},"upgrade plan")," (you may refer to a past version if you are using an older Cosmos SDK release). In a migration we can do any deterministic state change."),(0,i.kt)("p",null,"Build the new version ",(0,i.kt)("inlineCode",{parentName:"p"},"simd")," binary:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"make build\n")),(0,i.kt)("p",null,"Create the folder for the upgrade binary and copy the ",(0,i.kt)("inlineCode",{parentName:"p"},"simd")," binary:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"mkdir -p $DAEMON_HOME/cosmovisor/upgrades/test1/bin\ncp ./build/simd $DAEMON_HOME/cosmovisor/upgrades/test1/bin\n")),(0,i.kt)("p",null,"Open a new terminal window and submit an upgrade proposal along with a deposit and a vote (these commands must be run within 20 seconds of each other):\nNote, when using a ",(0,i.kt)("inlineCode",{parentName:"p"},"v0.46+")," chain, replace ",(0,i.kt)("inlineCode",{parentName:"p"},"submit-proposal")," by ",(0,i.kt)("inlineCode",{parentName:"p"},"submit-legacy-proposal"),"."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-shell"},"./build/simd tx gov submit-proposal software-upgrade test1 --title upgrade --description upgrade --upgrade-height 200 --from validator --yes\n./build/simd tx gov deposit 1 10000000stake --from validator --yes\n./build/simd tx gov vote 1 yes --from validator --yes\n")),(0,i.kt)("p",null,"The upgrade will occur automatically at height 200. Note: you may need to change the upgrade height in the snippet above if your test play takes more time."))}m.isMDXComponent=!0}}]);