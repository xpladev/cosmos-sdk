"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[7833],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>u});var a=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function r(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?r(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):r(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,a,i=function(e,t){if(null==e)return{};var n,a,i={},r=Object.keys(e);for(a=0;a<r.length;a++)n=r[a],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);for(a=0;a<r.length;a++)n=r[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var l=a.createContext({}),m=function(e){var t=a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},p=function(e){var t=m(e.components);return a.createElement(l.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},c=a.forwardRef((function(e,t){var n=e.components,i=e.mdxType,r=e.originalType,l=e.parentName,p=o(e,["components","mdxType","originalType","parentName"]),c=m(n),u=i,h=c["".concat(l,".").concat(u)]||c[u]||d[u]||r;return n?a.createElement(h,s(s({ref:t},p),{},{components:n})):a.createElement(h,s({ref:t},p))}));function u(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var r=n.length,s=new Array(r);s[0]=c;var o={};for(var l in t)hasOwnProperty.call(t,l)&&(o[l]=t[l]);o.originalType=e,o.mdxType="string"==typeof e?e:i,s[1]=o;for(var m=2;m<r;m++)s[m]=n[m];return a.createElement.apply(null,s)}return a.createElement.apply(null,n)}c.displayName="MDXCreateElement"},1514:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>s,default:()=>d,frontMatter:()=>r,metadata:()=>o,toc:()=>m});var a=n(7462),i=(n(7294),n(3905));const r={sidebar_position:1},s="Gas and Fees",o={unversionedId:"basics/gas-fees",id:"basics/gas-fees",title:"Gas and Fees",description:"This document describes the default strategies to handle gas and fees within a Cosmos SDK application.",source:"@site/docs/basics/04-gas-fees.md",sourceDirName:"basics",slug:"/basics/gas-fees",permalink:"/main/basics/gas-fees",draft:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Accounts",permalink:"/main/basics/accounts"},next:{title:"BaseApp",permalink:"/main/core/baseapp"}},l={},m=[{value:"Introduction to <code>Gas</code> and <code>Fees</code>",id:"introduction-to-gas-and-fees",level:2},{value:"Gas Meter",id:"gas-meter",level:2},{value:"Main Gas Meter",id:"main-gas-meter",level:3},{value:"Block Gas Meter",id:"block-gas-meter",level:3},{value:"AnteHandler",id:"antehandler",level:2}],p={toc:m};function d(e){let{components:t,...n}=e;return(0,i.kt)("wrapper",(0,a.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"gas-and-fees"},"Gas and Fees"),(0,i.kt)("admonition",{title:"Synopsis",type:"note"},(0,i.kt)("p",{parentName:"admonition"},"This document describes the default strategies to handle gas and fees within a Cosmos SDK application.")),(0,i.kt)("admonition",{type:"note"},(0,i.kt)("h3",{parentName:"admonition",id:"pre-requisite-readings"},"Pre-requisite Readings"),(0,i.kt)("ul",{parentName:"admonition"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/main/basics/app-anatomy"},"Anatomy of a Cosmos SDK Application")))),(0,i.kt)("h2",{id:"introduction-to-gas-and-fees"},"Introduction to ",(0,i.kt)("inlineCode",{parentName:"h2"},"Gas")," and ",(0,i.kt)("inlineCode",{parentName:"h2"},"Fees")),(0,i.kt)("p",null,"In the Cosmos SDK, ",(0,i.kt)("inlineCode",{parentName:"p"},"gas")," is a special unit that is used to track the consumption of resources during execution. ",(0,i.kt)("inlineCode",{parentName:"p"},"gas")," is typically consumed whenever read and writes are made to the store, but it can also be consumed if expensive computation needs to be done. It serves two main purposes:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Make sure blocks are not consuming too many resources and will be finalized. This is implemented by default in the Cosmos SDK via the ",(0,i.kt)("a",{parentName:"li",href:"#block-gas-meter"},"block gas meter"),"."),(0,i.kt)("li",{parentName:"ul"},"Prevent spam and abuse from end-user. To this end, ",(0,i.kt)("inlineCode",{parentName:"li"},"gas")," consumed during ",(0,i.kt)("a",{parentName:"li",href:"/main/building-modules/messages-and-queries#messages"},(0,i.kt)("inlineCode",{parentName:"a"},"message"))," execution is typically priced, resulting in a ",(0,i.kt)("inlineCode",{parentName:"li"},"fee")," (",(0,i.kt)("inlineCode",{parentName:"li"},"fees = gas * gas-prices"),"). ",(0,i.kt)("inlineCode",{parentName:"li"},"fees")," generally have to be paid by the sender of the ",(0,i.kt)("inlineCode",{parentName:"li"},"message"),". Note that the Cosmos SDK does not enforce ",(0,i.kt)("inlineCode",{parentName:"li"},"gas")," pricing by default, as there may be other ways to prevent spam (e.g. bandwidth schemes). Still, most applications will implement ",(0,i.kt)("inlineCode",{parentName:"li"},"fee")," mechanisms to prevent spam. This is done via the ",(0,i.kt)("a",{parentName:"li",href:"#antehandler"},(0,i.kt)("inlineCode",{parentName:"a"},"AnteHandler")),".")),(0,i.kt)("h2",{id:"gas-meter"},"Gas Meter"),(0,i.kt)("p",null,"In the Cosmos SDK, ",(0,i.kt)("inlineCode",{parentName:"p"},"gas")," is a simple alias for ",(0,i.kt)("inlineCode",{parentName:"p"},"uint64"),", and is managed by an object called a ",(0,i.kt)("em",{parentName:"p"},"gas meter"),". Gas meters implement the ",(0,i.kt)("inlineCode",{parentName:"p"},"GasMeter")," interface"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.46.0/store/types/gas.go#L40-L51\n")),(0,i.kt)("p",null,"where:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"GasConsumed()")," returns the amount of gas that was consumed by the gas meter instance."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"GasConsumedToLimit()")," returns the amount of gas that was consumed by gas meter instance, or the limit if it is reached."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"GasRemaining()")," returns the gas left in the GasMeter."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"Limit()")," returns the limit of the gas meter instance. ",(0,i.kt)("inlineCode",{parentName:"li"},"0")," if the gas meter is infinite."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"ConsumeGas(amount Gas, descriptor string)")," consumes the amount of ",(0,i.kt)("inlineCode",{parentName:"li"},"gas")," provided. If the ",(0,i.kt)("inlineCode",{parentName:"li"},"gas")," overflows, it panics with the ",(0,i.kt)("inlineCode",{parentName:"li"},"descriptor")," message. If the gas meter is not infinite, it panics if ",(0,i.kt)("inlineCode",{parentName:"li"},"gas")," consumed goes above the limit."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"RefundGas()")," deducts the given amount from the gas consumed. This functionality enables refunding gas to the transaction or block gas pools so that EVM-compatible chains can fully support the go-ethereum StateDB interface."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"IsPastLimit()")," returns ",(0,i.kt)("inlineCode",{parentName:"li"},"true")," if the amount of gas consumed by the gas meter instance is strictly above the limit, ",(0,i.kt)("inlineCode",{parentName:"li"},"false")," otherwise."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"IsOutOfGas()")," returns ",(0,i.kt)("inlineCode",{parentName:"li"},"true")," if the amount of gas consumed by the gas meter instance is above or equal to the limit, ",(0,i.kt)("inlineCode",{parentName:"li"},"false")," otherwise.")),(0,i.kt)("p",null,"The gas meter is generally held in ",(0,i.kt)("a",{parentName:"p",href:"/main/core/context"},(0,i.kt)("inlineCode",{parentName:"a"},"ctx")),", and consuming gas is done with the following pattern:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'ctx.GasMeter().ConsumeGas(amount, "description")\n')),(0,i.kt)("p",null,"By default, the Cosmos SDK makes use of two different gas meters, the ",(0,i.kt)("a",{parentName:"p",href:"#main-gas-metter"},"main gas meter")," and the ",(0,i.kt)("a",{parentName:"p",href:"#block-gas-meter"},"block gas meter"),"."),(0,i.kt)("h3",{id:"main-gas-meter"},"Main Gas Meter"),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"ctx.GasMeter()")," is the main gas meter of the application. The main gas meter is initialized in ",(0,i.kt)("inlineCode",{parentName:"p"},"BeginBlock")," via ",(0,i.kt)("inlineCode",{parentName:"p"},"setDeliverState"),", and then tracks gas consumption during execution sequences that lead to state-transitions, i.e. those originally triggered by ",(0,i.kt)("a",{parentName:"p",href:"/main/core/baseapp#beginblock"},(0,i.kt)("inlineCode",{parentName:"a"},"BeginBlock")),", ",(0,i.kt)("a",{parentName:"p",href:"/main/core/baseapp#delivertx"},(0,i.kt)("inlineCode",{parentName:"a"},"DeliverTx"))," and ",(0,i.kt)("a",{parentName:"p",href:"/main/core/baseapp#endblock"},(0,i.kt)("inlineCode",{parentName:"a"},"EndBlock")),". At the beginning of each ",(0,i.kt)("inlineCode",{parentName:"p"},"DeliverTx"),", the main gas meter ",(0,i.kt)("strong",{parentName:"p"},"must be set to 0")," in the ",(0,i.kt)("a",{parentName:"p",href:"#antehandler"},(0,i.kt)("inlineCode",{parentName:"a"},"AnteHandler")),", so that it can track gas consumption per-transaction."),(0,i.kt)("p",null,"Gas consumption can be done manually, generally by the module developer in the ",(0,i.kt)("a",{parentName:"p",href:"/main/building-modules/beginblock-endblock"},(0,i.kt)("inlineCode",{parentName:"a"},"BeginBlocker"),", ",(0,i.kt)("inlineCode",{parentName:"a"},"EndBlocker"))," or ",(0,i.kt)("a",{parentName:"p",href:"/main/building-modules/msg-services"},(0,i.kt)("inlineCode",{parentName:"a"},"Msg")," service"),", but most of the time it is done automatically whenever there is a read or write to the store. This automatic gas consumption logic is implemented in a special store called ",(0,i.kt)("a",{parentName:"p",href:"/main/core/store#gaskv-store"},(0,i.kt)("inlineCode",{parentName:"a"},"GasKv")),"."),(0,i.kt)("h3",{id:"block-gas-meter"},"Block Gas Meter"),(0,i.kt)("p",null,(0,i.kt)("inlineCode",{parentName:"p"},"ctx.BlockGasMeter()")," is the gas meter used to track gas consumption per block and make sure it does not go above a certain limit. A new instance of the ",(0,i.kt)("inlineCode",{parentName:"p"},"BlockGasMeter")," is created each time ",(0,i.kt)("a",{parentName:"p",href:"/main/core/baseapp#beginblock"},(0,i.kt)("inlineCode",{parentName:"a"},"BeginBlock"))," is called. The ",(0,i.kt)("inlineCode",{parentName:"p"},"BlockGasMeter")," is finite, and the limit of gas per block is defined in the application's consensus parameters. By default, Cosmos SDK applications use the default consensus parameters provided by Tendermint:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/tendermint/tendermint/blob/v0.34.21/types/params.go#L24-L65\n")),(0,i.kt)("p",null,"When a new ",(0,i.kt)("a",{parentName:"p",href:"/main/core/transactions"},"transaction")," is being processed via ",(0,i.kt)("inlineCode",{parentName:"p"},"DeliverTx"),", the current value of ",(0,i.kt)("inlineCode",{parentName:"p"},"BlockGasMeter")," is checked to see if it is above the limit. If it is, ",(0,i.kt)("inlineCode",{parentName:"p"},"DeliverTx")," returns immediately. This can happen even with the first transaction in a block, as ",(0,i.kt)("inlineCode",{parentName:"p"},"BeginBlock")," itself can consume gas. If not, the transaction is processed normally. At the end of ",(0,i.kt)("inlineCode",{parentName:"p"},"DeliverTx"),", the gas tracked by ",(0,i.kt)("inlineCode",{parentName:"p"},"ctx.BlockGasMeter()")," is increased by the amount consumed to process the transaction:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go"},'ctx.BlockGasMeter().ConsumeGas(\n    ctx.GasMeter().GasConsumedToLimit(),\n    "block gas meter",\n)\n')),(0,i.kt)("h2",{id:"antehandler"},"AnteHandler"),(0,i.kt)("p",null,"The ",(0,i.kt)("inlineCode",{parentName:"p"},"AnteHandler")," is run for every transaction during ",(0,i.kt)("inlineCode",{parentName:"p"},"CheckTx")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"DeliverTx"),", before a Protobuf ",(0,i.kt)("inlineCode",{parentName:"p"},"Msg")," service method for each ",(0,i.kt)("inlineCode",{parentName:"p"},"sdk.Msg")," in the transaction. "),(0,i.kt)("p",null,"The anteHandler is not implemented in the core Cosmos SDK but in a module. That said, most applications today use the default implementation defined in the ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/tree/main/x/auth"},(0,i.kt)("inlineCode",{parentName:"a"},"auth")," module"),". Here is what the ",(0,i.kt)("inlineCode",{parentName:"p"},"anteHandler")," is intended to do in a normal Cosmos SDK application:"),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Verify that the transactions are of the correct type. Transaction types are defined in the module that implements the ",(0,i.kt)("inlineCode",{parentName:"li"},"anteHandler"),", and they follow the transaction interface:")),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.46.0/types/tx_msg.go#L38-L46\n")),(0,i.kt)("p",null,"  This enables developers to play with various types for the transaction of their application. In the default ",(0,i.kt)("inlineCode",{parentName:"p"},"auth")," module, the default transaction type is ",(0,i.kt)("inlineCode",{parentName:"p"},"Tx"),": "),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-protobuf",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.46.0/proto/cosmos/tx/v1beta1/tx.proto#L13-L26\n")),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},"Verify signatures for each ",(0,i.kt)("a",{parentName:"li",href:"/main/building-modules/messages-and-queries#messages"},(0,i.kt)("inlineCode",{parentName:"a"},"message"))," contained in the transaction. Each ",(0,i.kt)("inlineCode",{parentName:"li"},"message")," should be signed by one or multiple sender(s), and these signatures must be verified in the ",(0,i.kt)("inlineCode",{parentName:"li"},"anteHandler"),"."),(0,i.kt)("li",{parentName:"ul"},"During ",(0,i.kt)("inlineCode",{parentName:"li"},"CheckTx"),", verify that the gas prices provided with the transaction is greater than the local ",(0,i.kt)("inlineCode",{parentName:"li"},"min-gas-prices")," (as a reminder, gas-prices can be deducted from the following equation: ",(0,i.kt)("inlineCode",{parentName:"li"},"fees = gas * gas-prices"),"). ",(0,i.kt)("inlineCode",{parentName:"li"},"min-gas-prices")," is a parameter local to each full-node and used during ",(0,i.kt)("inlineCode",{parentName:"li"},"CheckTx")," to discard transactions that do not provide a minimum amount of fees. This ensures that the mempool cannot be spammed with garbage transactions."),(0,i.kt)("li",{parentName:"ul"},"Verify that the sender of the transaction has enough funds to cover for the ",(0,i.kt)("inlineCode",{parentName:"li"},"fees"),". When the end-user generates a transaction, they must indicate 2 of the 3 following parameters (the third one being implicit): ",(0,i.kt)("inlineCode",{parentName:"li"},"fees"),", ",(0,i.kt)("inlineCode",{parentName:"li"},"gas")," and ",(0,i.kt)("inlineCode",{parentName:"li"},"gas-prices"),". This signals how much they are willing to pay for nodes to execute their transaction. The provided ",(0,i.kt)("inlineCode",{parentName:"li"},"gas")," value is stored in a parameter called ",(0,i.kt)("inlineCode",{parentName:"li"},"GasWanted")," for later use."),(0,i.kt)("li",{parentName:"ul"},"Set ",(0,i.kt)("inlineCode",{parentName:"li"},"newCtx.GasMeter")," to 0, with a limit of ",(0,i.kt)("inlineCode",{parentName:"li"},"GasWanted"),". ",(0,i.kt)("strong",{parentName:"li"},"This step is crucial"),", as it not only makes sure the transaction cannot consume infinite gas, but also that ",(0,i.kt)("inlineCode",{parentName:"li"},"ctx.GasMeter")," is reset in-between each ",(0,i.kt)("inlineCode",{parentName:"li"},"DeliverTx")," (",(0,i.kt)("inlineCode",{parentName:"li"},"ctx")," is set to ",(0,i.kt)("inlineCode",{parentName:"li"},"newCtx")," after ",(0,i.kt)("inlineCode",{parentName:"li"},"anteHandler")," is run, and the ",(0,i.kt)("inlineCode",{parentName:"li"},"anteHandler")," is run each time ",(0,i.kt)("inlineCode",{parentName:"li"},"DeliverTx")," is called).")),(0,i.kt)("p",null,"As explained above, the ",(0,i.kt)("inlineCode",{parentName:"p"},"anteHandler")," returns a maximum limit of ",(0,i.kt)("inlineCode",{parentName:"p"},"gas")," the transaction can consume during execution called ",(0,i.kt)("inlineCode",{parentName:"p"},"GasWanted"),". The actual amount consumed in the end is denominated ",(0,i.kt)("inlineCode",{parentName:"p"},"GasUsed"),", and we must therefore have ",(0,i.kt)("inlineCode",{parentName:"p"},"GasUsed =< GasWanted"),". Both ",(0,i.kt)("inlineCode",{parentName:"p"},"GasWanted")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"GasUsed")," are relayed to the underlying consensus engine when ",(0,i.kt)("a",{parentName:"p",href:"/main/core/baseapp#delivertx"},(0,i.kt)("inlineCode",{parentName:"a"},"DeliverTx"))," returns."))}d.isMDXComponent=!0}}]);