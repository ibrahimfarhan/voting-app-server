(this["webpackJsonpvoting-app-webapp"]=this["webpackJsonpvoting-app-webapp"]||[]).push([[2],{102:function(e,t,n){"use strict";n.d(t,"a",(function(){return o}));var a=n(18),r=n(1),c=n(22);function o(e){var t=e.action,n=e.selector,o=e.effectDeps,i=void 0===o?[]:o,s=e.actionParams,u=void 0===s?[]:s,l=e.useCache,d=void 0!==l&&l,b=e.cachedValueSelector,m=void 0===b?function(e){return e}:b,j=e.precedingAction,f=e.precedingActionParams,p=Object(c.b)(),h=Object(c.c)(n);return Object(r.useEffect)((function(){d&&m(h)||(j&&p(j(f)),p(t.apply(void 0,Object(a.a)(u))))}),i),h}},113:function(e,t,n){"use strict";var a=n(21),r=n(34),c=n(16),o=n(199),i=n(127),s=n(1),u=n(60),l=n.n(u),d=n(3);t.a=function(e){var t=e.buttonTitles,n=e.onTabSwitch,u=e.loading,b=e.initialSelectedIndex,m=void 0===b?0:b,j=e.stretchTabs,f=void 0===j||j,p=Object(s.useState)(m),h=Object(c.a)(p,2),g=h[0],O=h[1],v=Object(s.useMemo)((function(){return e.tabContents.reduce((function(e,t,n){return Object(r.a)(Object(r.a)({},e),{},Object(a.a)({},n,t))}),{})}),[e.tabContents]),x=function(e){var t=+e.currentTarget.id;O(t),n&&n(t)};return Object(d.jsxs)("div",{className:l.a.container,children:[Object(d.jsx)("div",{className:l.a["tab-buttons"],children:t.map((function(e,t){return Object(d.jsx)(o.a,{classes:{root:l.a[f?"btn--stretched":"btn--squashed"]},color:"primary",id:t.toString(),onClick:x,variant:g===t?"contained":"outlined",children:e},t)}))}),Object(d.jsx)("div",{className:l.a["tab-content"],children:u?Object(d.jsx)(i.a,{className:l.a.loading}):v[g]})]})}},119:function(e,t,n){e.exports={"fs-lg":"auth_fs-lg__2N9m-",container:"auth_container__3IqlF",error:"auth_error__1eoZJ"}},121:function(e,t,n){e.exports={"fs-lg":"loading_fs-lg__1lp4R",spinner:"loading_spinner__2G47J"}},123:function(e,t,n){e.exports={"fs-lg":"global-alert_fs-lg__2UNRh",message:"global-alert_message__2FBFX"}},124:function(e,t,n){e.exports={content:"root_content__1lKYC"}},19:function(e,t,n){"use strict";n.d(t,"c",(function(){return a})),n.d(t,"b",(function(){return r}));var a="Password must contain uppercase and lowercase letters, numbers and symbols",r=function(e,t,n,a){if(!e||0===Object.keys(e).length||!e[t])return"";switch(e[t].type){case"required":return"Field is required";case"minLength":return"Field length cannot be less than ".concat(n.minLength);case"maxLength":return"Field length cannot be greater than ".concat(n.maxLength);case"pattern":return a||"Field is invalid";default:return"Field is invalid"}};t.a={usernameOrEmail:{required:!0,maxLength:200,minLength:3},password:{required:!0,maxLength:15,minLength:8,pattern:/^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9])(?=.*?[#?!@$%^&*-]).{8,15}$/},username:{required:!0,maxLength:30,minLength:3,pattern:/^[a-zA-Z0-9_-]{3,30}$/},email:{required:!0,minLength:5,maxLength:200,pattern:/^[a-zA-z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/},votingSubject:{maxLength:500},teamName:{required:!0,minLength:3,maxLength:1e3,pattern:/^[a-zA-Z0-9_-\s]+$/}}},194:function(e,t,n){},195:function(e,t,n){"use strict";n.r(t);var a=n(1),r=n.n(a),c=n(10),o=n.n(c),i=n(16),s=n(239),u=n(61),l=n(13),d=n(102),b=n(12),m=n(41),j=n(25),f=Object(b.b)("receiveUser"),p=Object(b.b)("clearUserState"),h=function(){return Object(j.a)({path:m.a.current,successActionCreator:f,failureActionCreator:p})},g=n(34),O=n(240),v=n(199),x=n(72),_=n(32),y=n.n(_),C=n(19),T=n(75),w=n.n(T),k=n(22);function N(e){return[Object(k.b)(),Object(k.c)(e)]}var A=n(3),P=function(e){var t=e.location,n=Object(x.a)(),a=n.register,r=n.handleSubmit,c=n.errors,o=N((function(e){return e.user.data})),s=Object(i.a)(o,2),u=s[0],d=s[1],b=a(C.a.usernameOrEmail),p=a(C.a.password),h=Object(C.b)(c,"usernameOrEmail",C.a.usernameOrEmail),g=Object(C.b)(c,"password",C.a.password,C.c);return d?Object(A.jsx)(l.a,{to:w()(t,"state.from.pathName","/")}):Object(A.jsxs)("form",{className:y.a.form,onSubmit:r((function(e){var t={password:e.password};e.usernameOrEmail.includes("@")?t.email=e.usernameOrEmail:t.username=e.usernameOrEmail,u(function(e){return Object(j.a)({path:m.a.login,successActionCreator:f,additionalReqOptions:{method:"post",body:JSON.stringify(e)}})}(t))})),children:[Object(A.jsx)(O.a,{margin:"normal",className:y.a.item,inputRef:b,id:"usernameOrEmail",name:"usernameOrEmail",label:"Username or Email",type:"text",helperText:h,error:Boolean(h),required:!0}),Object(A.jsx)(O.a,{margin:"normal",helperText:g,error:Boolean(g),className:y.a.item,inputRef:p,id:"password",name:"password",label:"Password",type:"password",required:!0}),Object(A.jsx)(v.a,{classes:{root:y.a.btn},color:"primary",type:"submit",variant:"contained",children:"Submit"})]})},S=n(119),I=n.n(S),L=n(113),q=function(e){var t=e.location,n=Object(x.a)({reValidateMode:"onChange"}),a=n.register,r=n.handleSubmit,c=n.errors,o=N((function(e){return e.user.data})),s=Object(i.a)(o,2),u=s[0],d=s[1],b=a(C.a.username),p=a(C.a.email),h=a(C.a.password),g=Object(C.b)(c,"username",C.a.username),_=Object(C.b)(c,"email",C.a.email),T=Object(C.b)(c,"password",C.a.password,C.c);return d?Object(A.jsx)(l.a,{to:w()(t,"state.from.pathName","/")}):Object(A.jsxs)("form",{className:y.a.form,onSubmit:r((function(e){u(function(e){return Object(j.a)({path:m.a.register,successActionCreator:f,additionalReqOptions:{method:"post",body:JSON.stringify(e)}})}(e))})),children:[Object(A.jsx)(O.a,{helperText:g,error:Boolean(g),margin:"normal",className:y.a.item,inputRef:b,id:"username",name:"username",label:"Username",type:"text",required:!0}),Object(A.jsx)(O.a,{helperText:_,error:Boolean(_),margin:"normal",className:y.a.item,inputRef:p,id:"email",name:"email",label:"Email",type:"email",required:!0}),Object(A.jsx)(O.a,{helperText:T,error:Boolean(T),margin:"normal",className:y.a.item,inputRef:h,id:"password",name:"password",label:"Password",type:"password",required:!0}),Object(A.jsx)(v.a,{classes:{root:y.a.btn},color:"primary",type:"submit",variant:"contained",children:"Submit"})]})},E=n(120),R=function(e){return Object(A.jsxs)("div",{className:I.a.container,children:[Object(A.jsx)(v.a,{color:"primary",variant:"contained",startIcon:Object(A.jsx)(E.a,{}),onClick:function(){return window.location.href=m.a.googleLogin},children:"Log in with Google"}),Object(A.jsx)("h3",{children:"OR"}),Object(A.jsx)(L.a,{buttonTitles:["Log in","Register"],tabContents:[Object(A.jsx)(P,Object(g.a)({},e)),Object(A.jsx)(q,Object(g.a)({},e))]})]})},F=n(104),J=n(66),z=function(e){var t=e.component,n=Object(F.a)(e,["component"]),r=Object(k.c)((function(e){return e.user})),c=r.data,o=r.isLoggedIn,i=t;return Object(A.jsx)(l.b,Object(g.a)(Object(g.a)({},n),{},{render:function(e){return c&&o?Object(A.jsx)(a.Suspense,{fallback:Object(A.jsx)(J.a,{}),children:Object(A.jsx)(i,Object(g.a)({},e))}):Object(A.jsx)(l.a,{to:{pathname:"/login",state:{from:e.location}}})}}))},M=n(238),U=n(237),B=n(243),D=n(126),V=n(122),Z=n.n(V),H=n(84),$=n(244),G=n.p+"static/media/favicon.1cc6896d.png",X=n(96),K=n(28),Q=n.n(K),W=function(e){var t=e.onThemeChange,n=e.darkModeIsEnabled,a=Object(k.b)(),c=Object(l.g)(),o=Object(k.c)((function(e){return e.user})),s=o.data,u=o.isLoggedIn,d=void 0!==u&&u,b=r.a.useState(null),f=Object(i.a)(b,2),h=f[0],g=f[1],O=r.a.useState(null),v=Object(i.a)(O,2),x=v[0],_=v[1],y=Boolean(h),C=(Boolean(x),function(){g(null),_(null)}),T="primary-search-account-menu",w=Object(A.jsx)(D.a,{classes:{paper:Q.a.menu},anchorEl:h,anchorOrigin:{vertical:"bottom",horizontal:"right"},id:T,keepMounted:!0,getContentAnchorEl:null,transformOrigin:{vertical:"top",horizontal:"right"},open:y,onClose:C,children:Object(A.jsx)(B.a,{onClick:function(){a(Object(j.a)({path:m.a.logout,successActionCreator:p,additionalReqOptions:{method:"POST"}})),C()},children:"Log out"})});return Object(A.jsxs)("div",{className:Q.a.header,children:[Object(A.jsxs)(U.a,{onClick:function(){return c.push("/")},className:Q.a.logo,variant:"h6",noWrap:!0,children:[Object(A.jsx)("img",{className:Q.a.img,src:G}),Object(A.jsx)("strong",{className:Q.a.title,children:"Scrum Poker"})]}),Object(A.jsx)("div",{className:Q.a.grow}),Object(A.jsx)("div",{className:Q.a.sectionDesktop,children:Object(A.jsx)(H.a,{inCase:d,children:Object(A.jsxs)("div",{className:Q.a.loggedIn,children:[Object(A.jsx)("div",{onClick:t,children:n?Object(A.jsx)(X.b,{className:Q.a.lightIcon,size:"1.5rem",title:"Enable Light Mode"}):Object(A.jsx)(X.a,{className:Q.a.darkIcon,size:"1.5rem",title:"Enable Dark Mode"})}),Object(A.jsx)($.a,{className:Q.a.avatar,"aria-label":"account of current user","aria-controls":T,"aria-haspopup":"true",onClick:function(e){g(e.currentTarget)},src:null===s||void 0===s?void 0:s.pictureUrl,alt:"".concat((null===s||void 0===s?void 0:s.username)||(null===s||void 0===s?void 0:s.name)," (").concat(null===s||void 0===s?void 0:s.email,")"),title:"".concat((null===s||void 0===s?void 0:s.username)||(null===s||void 0===s?void 0:s.name)," (").concat(null===s||void 0===s?void 0:s.email,")"),children:Object(A.jsx)("strong",{children:((null===s||void 0===s?void 0:s.username)||(null===s||void 0===s?void 0:s.name)||"")[0]})})]})})}),Object(A.jsx)("div",{className:Q.a.sectionMobile,children:Object(A.jsx)(M.a,{"aria-label":"show more","aria-controls":"primary-search-account-menu-mobile","aria-haspopup":"true",onClick:function(e){_(e.currentTarget)},color:"inherit",children:Object(A.jsx)(H.a,{inCase:d,children:Object(A.jsx)(Z.a,{})})})}),w]})},Y=n(241),ee=n(40),te=n(123),ne=n.n(te),ae=function(){var e=Object(k.b)(),t=Object(k.c)((function(e){return e.global.message})),n=t.content,a=t.type,r=void 0===a?"info":a;return n?Object(A.jsx)(Y.a,{severity:r,className:ne.a.message,onClose:function(){return e(Object(ee.a)())},children:n}):null},re=n(125),ce={"--bg-color":"white","--font-color":"black","--border-color":"#8f8f8f","--lighter-border-color":"#b6b6b6","--scrollbar-color":"gray","--scrollbar-bg-color":"white","--blue-font-color":"var(--blue)"},oe={"--bg-color":"rgb(26, 26, 26)","--font-color":"white","--border-color":"#424242","--lighter-border-color":"#2c2c2c","--scrollbar-color":"#2d2d2d","--scrollbar-bg-color":"rgb(26, 26, 26)","--blue-font-color":"var(--light-blue)"},ie=document.querySelector(":root"),se=function(e){var t=getComputedStyle(ie);return Object(re.a)({typography:{button:{textTransform:"none"},fontFamily:t.getPropertyValue("--main-font-family"),fontSize:16,allVariants:{color:t.getPropertyValue("--font-color").trim()}},palette:{type:e?"dark":"light",primary:{main:t.getPropertyValue("--blue").trim()},secondary:{main:t.getPropertyValue("--red").trim()}}})},ue=function(e){Object.entries(e?oe:ce).forEach((function(e){var t=Object(i.a)(e,2),n=t[0],a=t[1];ie.style.setProperty(n,a)}))},le=n(97),de=n(65),be=n.n(de),me=function(){return Object(A.jsxs)("footer",{children:[Object(A.jsx)("p",{className:be.a.sentence,children:"Developed and maintained by Ibrahim Farhan"}),Object(A.jsxs)("div",{className:be.a.icons,children:[Object(A.jsx)("a",{target:"blank",className:be.a.icon,href:"https://linkedin.com/in/ibrahimahmadfarhan",children:Object(A.jsx)(le.b,{color:"#0072B1",size:"2rem"})}),Object(A.jsx)("a",{className:be.a.icon,target:"blank",href:"https://github.com/ibrahimfarhan",children:Object(A.jsx)(le.a,{size:"2rem",color:"black"})})]})]})},je=n(124),fe=n.n(je),pe=r.a.lazy((function(){return Promise.all([n.e(0),n.e(1),n.e(5)]).then(n.bind(null,324))})),he=r.a.lazy((function(){return n.e(9).then(n.bind(null,319))})),ge=r.a.lazy((function(){return Promise.all([n.e(0),n.e(1),n.e(6)]).then(n.bind(null,325))})),Oe=r.a.lazy((function(){return Promise.all([n.e(0),n.e(8),n.e(7)]).then(n.bind(null,322))})),ve=r.a.lazy((function(){return n.e(10).then(n.bind(null,320))})),xe=function(){var e="true"===localStorage.getItem("darkModeIsEnabled"),t=Object(a.useState)(e),n=Object(i.a)(t,2),r=n[0],c=n[1];Object(a.useEffect)((function(){ue(e),j(se(e))}),[]);var o=Object(a.useState)(null),b=Object(i.a)(o,2),m=b[0],j=b[1];return Object(d.a)({action:h,selector:function(e){return e.user}}).loading||!m?Object(A.jsx)(J.a,{}):Object(A.jsxs)(s.a,{theme:m,children:[Object(A.jsxs)(u.a,{children:[Object(A.jsx)(W,{onThemeChange:function(){ue(!r),localStorage.setItem("darkModeIsEnabled",(!r).toString()),j(se(!r)),c(!r)},darkModeIsEnabled:r}),Object(A.jsx)("div",{className:fe.a.content,children:Object(A.jsxs)(l.d,{children:[Object(A.jsx)(z,{exact:!0,path:"/",component:pe}),Object(A.jsx)(l.b,{path:"/login",component:R}),Object(A.jsx)(z,{exact:!0,path:"/team/:teamId",component:ge}),Object(A.jsx)(z,{exact:!0,path:"/team/join/:token",component:he}),Object(A.jsx)(z,{exact:!0,path:"/voting/:teamId",component:Oe}),Object(A.jsx)(l.b,{component:ve})]})})]}),Object(A.jsx)(ae,{}),Object(A.jsx)(me,{})]})},_e=function(e){e&&e instanceof Function&&n.e(11).then(n.bind(null,321)).then((function(t){var n=t.getCLS,a=t.getFID,r=t.getFCP,c=t.getLCP,o=t.getTTFB;n(e),a(e),r(e),c(e),o(e)}))},ye=n(31),Ce={loading:!0,isLoggedIn:!1},Te=Object(b.c)(Ce,(function(e){e.addCase(f,(function(e,t){e.data=t.payload,e.data.username||(e.data.username=e.data.name||""),e.loading=!1,e.isLoggedIn=!0})),e.addCase(p,(function(e){e.loading=!1,delete e.error,delete e.data,e.isLoggedIn=!1})),e.addDefaultCase((function(e){Ce}))})),we=n(21),ke=n(33),Ne={loading:!0},Ae=Object(b.c)(Ne,(function(e){e.addCase(ke.l,(function(e,t){var n=t.payload;n.currentUserIsJoined=!0,e.data||(e.data={}),e.data[n.id]=Object.assign(e.data[n.id]||{},n),e.loading=!1})),e.addCase(ke.m,(function(e,t){e.data=t.payload.reduce((function(e,t){return Object.assign(e,Object(we.a)({},t.id,Object.assign(t,{currentUserIsJoined:!0})))}),{}),e.loading=!1,e.loadingSelectedTeams=!1,e.selectedTeamsType="joined"})),e.addCase(ke.k,(function(e,t){e.data=t.payload.reduce((function(e,t){return Object.assign(e,Object(we.a)({},t.id,Object.assign(t,{currentUserIsJoined:!1})))}),{}),e.loading=!1,e.loadingSelectedTeams=!1,e.selectedTeamsType="nonJoined"})),e.addCase(ke.o,(function(e,t){e.data&&delete e.data[t.payload]})),e.addCase(ke.g,(function(e,t){var n=t.payload,a=n.teamId,r=n.userId,c=n.wasAdmin;if(e.data){var o=e.data[a],i=o.admins,s=void 0===i?[]:i,u=o.members,l=c?s:void 0===u?[]:u,d=l.findIndex((function(e){return e.id===r}));l.splice(d,1)}})),e.addCase(ke.f,(function(e,t){e.data&&(e.data[t.payload].currentUserIsJoined=!0);e.loading=!1})),e.addCase(ke.p,(function(e,t){e.loading=!0})),e.addCase(ke.a,(function(e,t){e.loadingSelectedTeams=!0,e.selectedTeamsType="joined"===e.selectedTeamsType?"nonJoined":"joined"})),e.addDefaultCase((function(e){Ne}))})),Pe=Object(b.c)({message:{}},(function(e){e.addCase(ee.b,(function(e,t){e.message=t.payload})),e.addCase(ee.a,(function(e){e.message={}}))})),Se=Object(ye.c)({user:Te,teams:Ae,global:Pe}),Ie=Object(b.a)({reducer:Se});n(194);o.a.render(Object(A.jsx)(r.a.StrictMode,{children:Object(A.jsx)(k.a,{store:Ie,children:Object(A.jsx)(xe,{})})}),document.getElementById("root")),_e()},25:function(e,t,n){"use strict";n.d(t,"a",(function(){return s}));var a=n(14),r=n.n(a),c=n(26),o=n(83),i=n(99),s=function(e){var t=e.path,n=e.successActionCreator,a=e.failureActionCreator,s=e.additionalReqOptions,l=e.successActionPayload,d=e.onSuccess,b=e.onFailure;return function(){var e=Object(c.a)(r.a.mark((function e(c){var m,j,f;return r.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.prev=0,e.next=3,fetch(t,u(s||{}));case 3:return m=e.sent,e.next=6,Object(i.b)(m);case 6:if(j=e.sent,m.ok){e.next=10;break}return b?b(c,j.message):a?c(a(j.message)):Object(o.f)(c,{content:j.message,type:"error"}),e.abrupt("return");case 10:d?d(c,j):n&&c(n(l||j)),e.next=17;break;case 13:e.prev=13,e.t0=e.catch(0),f="Connection Error",b?b(c,f):a?c(a(f)):Object(o.f)(c,{content:f,type:"error"});case 17:case"end":return e.stop()}}),e,null,[[0,13]])})));return function(t){return e.apply(this,arguments)}}()},u=function(e){return Object.assign({headers:{"content-type":"application/json"},credentials:"include"},e)}},28:function(e,t,n){e.exports={header:"mui-header_header__1H-Pg",logo:"mui-header_logo__3fx7D",menuButton:"mui-header_menuButton__3-t_j",sectionDesktop:"mui-header_sectionDesktop__3GX2h",title:"mui-header_title__xR8yN",sectionMobile:"mui-header_sectionMobile__1nO2Q",avatar:"mui-header_avatar__2Kj4W",loggedIn:"mui-header_loggedIn__3F9hD",darkIcon:"mui-header_darkIcon__3Q7fA",lightIcon:"mui-header_lightIcon__3NRmV"}},32:function(e,t,n){e.exports={"fs-lg":"auth-form_fs-lg__2L0ms",form:"auth-form_form__22C4F",item:"auth-form_item__q-1vt",btn:"auth-form_btn__3qbqp"}},33:function(e,t,n){"use strict";n.d(t,"l",(function(){return i})),n.d(t,"m",(function(){return s})),n.d(t,"k",(function(){return u})),n.d(t,"o",(function(){return l})),n.d(t,"g",(function(){return d})),n.d(t,"f",(function(){return m})),n.d(t,"p",(function(){return j})),n.d(t,"a",(function(){return f})),n.d(t,"d",(function(){return p})),n.d(t,"e",(function(){return h})),n.d(t,"n",(function(){return g})),n.d(t,"j",(function(){return O})),n.d(t,"b",(function(){return v})),n.d(t,"i",(function(){return x})),n.d(t,"h",(function(){return _})),n.d(t,"c",(function(){return y}));var a=n(12),r=n(35),c=n(41),o=n(25),i=Object(a.b)("receiveTeam"),s=Object(a.b)("receiveTeams"),u=Object(a.b)("receivePublicNonJoinedTeams"),l=Object(a.b)("removeTeamFromTable"),d=Object(a.b)("handleRemovedMember"),b=Object(a.b)("handleJoiningTeam"),m=Object(a.b)("handleJoiningPublicTeam"),j=Object(a.b)("startLoading"),f=Object(a.b)("changeSelectedTeamsType"),p=function(){return Object(o.a)({path:c.a.teams,successActionCreator:s})},h=function(){return Object(o.a)({path:r.a.getPublic,successActionCreator:u})},g=function(e){return Object(o.a)({path:r.a.removeMember,successActionCreator:d,successActionPayload:e,additionalReqOptions:{method:"PATCH",body:JSON.stringify(e)}})},O=function(e){return Object(o.a)({path:r.a.leave,successActionCreator:l,successActionPayload:e,additionalReqOptions:{method:"PATCH",body:JSON.stringify({id:e})}})},v=function(e){return Object(o.a)({path:r.a.create,successActionCreator:i,additionalReqOptions:{method:"POST",body:JSON.stringify(e)}})},x=function(e){return Object(o.a)({path:r.a.getJoinTeamUrl(e),successActionCreator:b,additionalReqOptions:{method:"PATCH"}})},_=function(e){return Object(o.a)({path:r.a.getJoinTeamUrl(e),successActionCreator:m,successActionPayload:e,additionalReqOptions:{method:"PATCH"}})},y=function(e){return Object(o.a)({path:r.a.getUrl(e),successActionCreator:l,successActionPayload:e,additionalReqOptions:{method:"DELETE"}})}},35:function(e,t,n){"use strict";var a=n(74),r="".concat(a.a||"","/api/v1/team"),c={getPublic:"".concat(r,"/all-public"),getUrl:function(e){return"".concat(r,"/").concat(e)},create:"".concat(r,"/create"),getInvitationLinkUrl:function(e){return"".concat(r,"/").concat(e,"/invite")},getJoinTeamUrl:function(e){return"".concat(r,"/join/").concat(e)},leave:"".concat(r,"/leave"),toggleRole:"".concat(r,"/toggle-role"),removeMember:"".concat(r,"/remove-member")};t.a=c},40:function(e,t,n){"use strict";n.d(t,"b",(function(){return r})),n.d(t,"a",(function(){return c}));var a=n(12),r=Object(a.b)("showMessage"),c=Object(a.b)("hideMessage")},41:function(e,t,n){"use strict";var a=n(74),r="".concat(a.a,"/api/v1/user"),c={login:"".concat(r,"/login"),register:"".concat(r,"/register"),logout:"".concat(r,"/logout"),googleLogin:"".concat(r,"/oauth/google/login"),teams:"".concat(r,"/teams"),current:"".concat(r,"/"),sendVerificationEmail:"".concat(r,"/send-verification-email"),verifyEmail:"".concat(r,"/verify-email"),changePassword:"".concat(r,"/change-password"),sendResetPasswordLink:"".concat(r,"/send-reset-password-link"),resetPassword:"".concat(r,"/reset-password")};t.a=c},60:function(e,t,n){e.exports={"fs-lg":"tab-panel_fs-lg__2dEiO",container:"tab-panel_container__3AgO6","tab-buttons":"tab-panel_tab-buttons__1ZFSs","btn--stretched":"tab-panel_btn--stretched__3XhAp","btn--squashed":"tab-panel_btn--squashed__3YN_4",loading:"tab-panel_loading__k3p4y"}},65:function(e,t,n){e.exports={sentence:"footer_sentence__2sbFv",icon:"footer_icon__2Cx79"}},66:function(e,t,n){"use strict";var a=n(127),r=n(121),c=n.n(r),o=n(3);t.a=function(){return Object(o.jsx)("div",{className:c.a.spinner,children:Object(o.jsx)(a.a,{})})}},74:function(e,t,n){"use strict";t.a="https://voting-app-bqaf.onrender.com"},83:function(e,t,n){"use strict";n.d(t,"e",(function(){return c})),n.d(t,"b",(function(){return o})),n.d(t,"c",(function(){return i})),n.d(t,"a",(function(){return s})),n.d(t,"d",(function(){return u})),n.d(t,"f",(function(){return l}));var a=n(21),r=n(40),c=function(e){if(e)return["http://","https://"].some((function(t){return e.includes(t)}))?e:void 0},o=function(e){for(var t=arguments.length,n=new Array(t>1?t-1:0),a=1;a<t;a++)n[a-1]=arguments[a];return null===n||void 0===n?void 0:n.reduce((function(t,n){return""===t?e[n]:"".concat(t," ").concat(e[n])}),"")},i=function(e){return e&&"function"===typeof e.then},s=function(e,t){return e.reduce((function(e,n){return Object.assign(e,Object(a.a)({},t?n[t]:n,!0))}),{})},u=function(e,t){if(t<0||t>=e.length)return e;var n=e.splice(t,1)[0];return e.unshift(n),e},l=function(e,t){e(Object(r.b)(t)),setTimeout((function(){return e(Object(r.a)())}),t.displayDuration||5e3)}},84:function(e,t,n){"use strict";var a=n(1),r=n(3);t.a=function(e){var t=e.children,n=e.inCase,c=Array.isArray(t)?t:[t,null];return Object(r.jsx)(a.Fragment,{children:n?c[0]:c[1]})}},99:function(e,t,n){"use strict";n.d(t,"b",(function(){return s}));var a=n(14),r=n.n(a),c=n(26),o=function(){var e=Object(c.a)(r.a.mark((function e(t,n){var a,c;return r.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.prev=0,e.next=3,fetch(t,i(n||{}));case 3:return a=e.sent,e.next=6,s(a);case 6:if(c=e.sent,a.ok){e.next=9;break}return e.abrupt("return",Promise.reject(c));case 9:return e.abrupt("return",Promise.resolve(c));case 12:return e.prev=12,e.t0=e.catch(0),e.abrupt("return",Promise.reject(e.t0));case 15:case"end":return e.stop()}}),e,null,[[0,12]])})));return function(t,n){return e.apply(this,arguments)}}(),i=function(e){return Object.assign({headers:{"content-type":"application/json"},credentials:"include"},e)};function s(e){if(204===e.status)return Promise.resolve();var t=e.headers.get("content-type");return"application/json"===t?e.json():"plain/text"===t?e.text():Promise.reject("Error while reading server response")}t.a=o}},[[195,3,4]]]);
//# sourceMappingURL=main.322e744c.chunk.js.map