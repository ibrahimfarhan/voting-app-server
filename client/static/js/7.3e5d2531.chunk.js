(this["webpackJsonpvoting-app-webapp"]=this["webpackJsonpvoting-app-webapp"]||[]).push([[7],{246:function(e,t,n){"use strict";var a=n(16),s=n(1),r=n(127),i=n(199),c=n(326),o=n(313),l=n(311),d=n(312),u=n(310),b=n(84),m=n(83),j=n(3);t.a=function(e){var t=e.title,n=e.message,v=e.acceptText,g=e.cancelText,O=e.renderTrigger,f=e.action,_=e.onSuccess,h=e.onError,p=Object(s.useState)(!1),x=Object(a.a)(p,2),S=x[0],y=x[1],k=Object(s.useState)(!1),N=Object(a.a)(k,2),V=N[0],w=N[1],C=function(){return y(!1)};return Object(j.jsxs)(j.Fragment,{children:[Object(j.jsx)(O,{open:function(){return y(!0)},close:C}),Object(j.jsxs)(c.a,{open:S,onClose:C,maxWidth:"sm",fullWidth:!0,disableEscapeKeyDown:V,disableBackdropClick:V,"aria-labelledby":"alert-dialog-title","aria-describedby":"alert-dialog-description",children:[Object(j.jsx)(u.a,{id:"alert-dialog-title",children:t}),n&&Object(j.jsx)(l.a,{children:Object(j.jsx)(d.a,{id:"alert-dialog-description",children:n})}),Object(j.jsxs)(o.a,{children:[Object(j.jsx)(i.a,{onClick:C,color:"secondary",variant:"contained",disabled:V,children:null!==g&&void 0!==g?g:"Cancel"}),Object(j.jsxs)(i.a,{onClick:function(){var e=f();if(!Object(m.c)(e))return C();w(!0),e.then((function(e){_&&_(e)})).catch((function(e){h&&h(e)})).finally((function(){w(!1),C()}))},color:"primary",variant:"contained",autoFocus:!0,disabled:V,children:[Object(j.jsxs)(b.a,{inCase:V,children:[Object(j.jsx)(r.a,{size:20}),Object(j.jsx)(j.Fragment,{children:null!==v&&void 0!==v?v:"Yes"})]}),V]})]})]})]})}},249:function(e,t,n){"use strict";n.d(t,"a",(function(){return r}));var a=n(16),s=n(1);function r(e){var t=Object(s.useReducer)((function(e,t){return Object.assign({},e,t)}),e),n=Object(a.a)(t,2);return[n[0],n[1]]}},303:function(e,t,n){e.exports={"fs-lg":"voting-session_fs-lg__NcCB0",container:"voting-session_container__343Kg",header:"voting-session_header__22OUM",subject:"voting-session_subject__2uBqO","team-name-wrapper":"voting-session_team-name-wrapper__1ZqsG","team-name":"voting-session_team-name__1aU_V","admin-btns":"voting-session_admin-btns__1abX_",empty:"voting-session_empty__1AvRl",body:"voting-session_body__OY0Xw","members-container":"voting-session_members-container__1yzRm",cards:"voting-session_cards__3_dZo",results:"voting-session_results__llr5L",members:"voting-session_members__1dfaC",member:"voting-session_member__1Hf5V","admin-icon":"voting-session_admin-icon__gTCXg","name-container":"voting-session_name-container__IGLu_",name:"voting-session_name__2ic3f",moderator:"voting-session_moderator__2cPuX",admin:"voting-session_admin__1Mtmr",card:"voting-session_card__2s2KA","special-card":"voting-session_special-card__3H0Cl",icon:"voting-session_icon__3sZXa","number-wrapper":"voting-session_number-wrapper__3ORLH",number:"voting-session_number__2u3wa","number-sm":"voting-session_number-sm__XbEOK","selected-card":"voting-session_selected-card__2RELS",stats:"voting-session_stats__3UUiv"}},304:function(e,t,n){e.exports={"fs-lg":"editable-on-click_fs-lg__3J3qF",text:"editable-on-click_text__3uOma",empty:"editable-on-click_empty__2EER1",editable:"editable-on-click_editable__271Zu",error:"editable-on-click_error__1VxBt",disabled:"editable-on-click_disabled__2HuPQ"}},305:function(e,t,n){e.exports={"fs-lg":"voting-results_fs-lg__149bv",container:"voting-results_container__34O4q",avg:"voting-results_avg__3CBKI",percentages:"voting-results_percentages__1qKj9",percentage:"voting-results_percentage__t8ivr",chart:"voting-results_chart__1clPu"}},322:function(e,t,n){"use strict";n.r(t);var a,s=n(18),r=n(16),i=n(199),c=n(1),o=n(22),l=n(13),d=n(249),u=n(84),b=n(303),m=n.n(b),j=n(83),v=n(304),g=n.n(v),O=n(3),f=function(e){var t,n=e.onSubmit,a=e.disableEdit,s=e.initialText,i=void 0===s?"":s,l=e.validationRules,d=void 0===l?{}:l,u=e.customErrorMsg,b=e.classes,m=e.placeholder,v=void 0===m?"":m,f=Object(o.b)(),_=d.maxLength,h=void 0===_?500:_,p=d.pattern,x=Object(c.useState)(!1),S=Object(r.a)(x,2),y=S[0],k=S[1],N=Object(c.useState)(""),V=Object(r.a)(N,2),w=V[0],C=V[1],M=Object(c.useState)(!Boolean(i)),E=Object(r.a)(M,2),R=E[0],T=E[1],U=Object(c.useRef)(null);Object(c.useEffect)((function(){T(!Boolean(i)),(null===U||void 0===U?void 0:U.current)&&(U.current.textContent=i||v,C(U.current.textContent))}),[U,i,v]);var A=a?void 0:function(){return k(!0)};return Object(O.jsx)("a",{contentEditable:!a,className:"".concat(g.a.text," ").concat(null!==(t=null===b||void 0===b?void 0:b.input)&&void 0!==t?t:""," ").concat(y?g.a.editable:""," ").concat(R?g.a.empty:""),ref:U,href:Object(j.e)(w),id:"field",spellCheck:"false",onBlur:function(){var e,t=(null===(e=U.current)||void 0===e?void 0:e.textContent)||"";t||(T(!0),U.current&&(U.current.textContent=v)),!p||p.test(t)?t.length>h?Object(j.f)(f,{type:"error",content:u||"Content cannot be larger than ".concat(h)}):((null===U||void 0===U?void 0:U.current)&&t&&(U.current.textContent=t.trim()),n&&n(t),k(!1)):Object(j.f)(f,{type:"error",content:u||"Invalid value"})},onFocus:function(){T(!1),U.current&&R&&(U.current.textContent="")},target:"blank",onMouseDown:A,defaultValue:R?v:i})},_=n(246),h=n(257),p=n.n(h),x=n(305),S=n.n(x),y=n(306),k=["tomato","green","orange","gray","rgba(153, 102, 255, 0.2)","rgba(255, 159, 64, 0.2)"],N=function(e){var t=e.memberVotes,n=Object.keys(t).reduce((function(e,n){var a=t[n]||0;return a<=0||(e.countPerVote[a]=e.countPerVote[a]?e.countPerVote[a]+1:1,e.avgNominator+=a,e.avgDenominator++),e}),{countPerVote:{},avgNominator:0,avgDenominator:0}),a=n.countPerVote,s=void 0===a?{}:a,r=n.avgDenominator,i=void 0===r?0:r,c=n.avgNominator,o=void 0===c?0:c,l=Math.round(o/i*10)/10,d=Object.keys(s).map((function(e,t){return{title:e,value:s[e],color:k[t%k.length]}}));return Object(O.jsxs)("div",{className:S.a.container,children:[Object(O.jsx)("div",{className:S.a.avg,children:Object(O.jsxs)("strong",{children:["Average: ",l||"?"]})}),Object(O.jsx)("div",{className:S.a.percentages,children:Object(O.jsx)(y.PieChart,{className:S.a.chart,data:d,label:function(e){var t=e.dataEntry,n=t.value,a=t.title;return"".concat(n," vote").concat(n>1?"s":""," for ").concat(a)},labelStyle:{fontSize:"5px",fill:"white",fontWeight:"bold"}})})]})},V=n(66),w=n(99),C=n(35),M=n(21),E=n(34);!function(e){e.submitVote="submitVote",e.toggleResultsDisplay="toggleResultsDisplay",e.closeSession="closeSession",e.displaySubject="displaySubject",e.reset="reset",e.changeModerator="changeModerator"}(a||(a={}));var R="".concat("ws://localhost:8080","/api/v1/voting")||!1,T=[1,2,3,5,8,13,21,34,55],U={voted:-1,unknown:-2,initial:0},A={closeSession:function(e){e.conn.close()},showNewMember:function(e){var t=e.msg,n=e.setState,a=e.conn,s=t.data,r=Object(E.a)({},a.state.memberVotes),i=Object(E.a)(Object(E.a)({},a.state.members),{},Object(M.a)({},s.id,s));r[s.id]=U.initial,n({memberVotes:r,members:i})},removeLeavingMember:function(e){var t=e.msg,n=e.conn,a=e.setState,s=t.data,r=Object(E.a)({},n.state.memberVotes),i=Object(E.a)({},n.state.members);delete r[s.id],delete i[s.id],a({memberVotes:r,members:i})},showCurrentMembers:function(e){var t,n=e.msg,a=e.setState,s=e.currentUser,r=n.data,i=r.sessionModerator,c=r.currentMembers,o=s.id===i.id?Object(M.a)({},i.id,U.initial):(t={},Object(M.a)(t,s.id,U.initial),Object(M.a)(t,i.id,U.initial),t),l=c.reduce((function(e,t){return Object.assign(e,Object(M.a)({},t.id,t))}),Object(M.a)({},i.id,i));a({sessionModerator:i.id,members:l,memberVotes:c.reduce((function(e,t){return t.id===s.id?e:Object(E.a)(Object(E.a)({},e),{},Object(M.a)({},t.id,U.initial))}),o),loading:!1})},sessionCreated:function(e){var t=e.setState,n=e.currentUser;t({sessionModerator:n.id,currentUserIsSessionModerator:!0,loading:!1,memberVotes:Object(M.a)({},n.id,U.initial),members:Object(M.a)({},n.id,n)})},displaySubject:function(e){var t=e.msg,n=e.setState,a=e.conn,s=t.data,r={};s!==a.state.currentSubject&&(r.currentSubject=s),n(r)},toggleResultsDisplay:function(e){var t=e.msg,n=e.setState,a=e.conn,s=e.currentUser,r=t.data||{},i=r.votes,c=void 0===i?{}:i,o=r.show,l=void 0!==o&&o,d=a.state.memberVotes||{},u={};Object.keys(d).forEach((function(e){u[e]=l?c[e]:e!==s.id&&d[e]?U.voted:d[e]||U.initial})),n({memberVotes:u,resultsAreShown:l})},showVote:function(e){var t=e.msg,n=e.conn,a=e.setState,s=e.currentUser,r=t.data,i={memberVotes:Object.assign({},n.state.memberVotes,r)};r[s.id]&&(i.lockedOption=r[s.id]),a(i)},reset:function(e){var t=e.setState,n=e.conn;t({memberVotes:Object.keys(n.state.memberVotes||{}).reduce((function(e,t){return Object.assign({},e,Object(M.a)({},t,U.initial))}),{}),lockedOption:U.initial,resultsAreShown:!1})},changeModerator:function(e){var t=e.msg,n=e.conn,a=e.setState,s=e.currentUser,r=t.data;a({memberVotes:Object.keys(n.state.memberVotes||{}).reduce((function(e,t){return Object.assign({},e,Object(M.a)({},t,U.initial))}),{}),lockedOption:U.initial,resultsAreShown:!1,sessionModerator:r,currentUserIsSessionModerator:s.id===r})}},J=function(e){var t;return t={},Object(M.a)(t,a.submitVote,function(e){return function(t){e.send(JSON.stringify({action:a.submitVote,data:t}))}}(e)),Object(M.a)(t,a.toggleResultsDisplay,function(e){return function(){return e.send(JSON.stringify({action:a.toggleResultsDisplay}))}}(e)),Object(M.a)(t,a.closeSession,function(e){return function(){return e.send(JSON.stringify({action:a.closeSession}))}}(e)),Object(M.a)(t,a.displaySubject,function(e){return function(t){e.send(JSON.stringify({action:a.displaySubject,data:t}))}}(e)),Object(M.a)(t,a.reset,function(e){return function(){return e.send(JSON.stringify({action:a.reset}))}}(e)),Object(M.a)(t,a.changeModerator,function(e){return function(t){e.send(JSON.stringify({action:a.changeModerator,data:t}))}}(e)),t},D=n(307),B={sessionModerator:"",currentUserIsSessionModerator:!1,currentSubject:"",members:{},memberVotes:{},isRunning:!0,resultsAreShown:!1},P=new Map([[U.voted,Object(O.jsx)(p.a,{classes:{root:m.a.icon}})],[U.unknown,"?"],[U.initial,""]]),I=function(e){var t=P.get(e);return t&&"string"!==typeof t?t:Object(O.jsx)("p",{className:m.a["number-wrapper"],children:Object(O.jsx)("span",{className:Number(e)>=10?m.a.number:m.a.number+" "+m.a["number-sm"],children:void 0===t?e:t})})};t.default=function(){var e,t,n=Object(o.b)(),a=Object(l.h)().teamId,b=Object(l.g)(),v=Object(o.c)((function(e){return e.teams})).data,g=Object(d.a)({}),h=Object(r.a)(g,2),p=h[0],x=p.team,S=p.loadingTeam,y=h[1];Object(c.useEffect)((function(){v&&v[a]?y({team:v[a],loadingTeam:!1}):Object(w.a)(C.a.getUrl(a)).then((function(e){y({loadingTeam:!1,team:e})})).catch((function(){y({loadingTeam:!1,teamApiError:"Error while fetching team"})}))}),[a]);var k=Object(o.c)((function(e){return e.user.data})),M=Object(d.a)(B),P=Object(r.a)(M,2),L=P[0],F=P[1],K=L.memberVotes,X=void 0===K?{}:K,q=L.members,H=void 0===q?{}:q,W=L.sessionModerator,Y=void 0===W?"":W,Z=L.actionSenders,z=L.isRunning,G=L.currentUserIsSessionModerator,Q=void 0!==G&&G,$=L.loading,ee=void 0===$||$,te=L.currentSubject,ne=void 0===te?"":te,ae=L.resultsAreShown,se=L.lockedOption;if(Object(c.useEffect)((function(){if(!z)return Object(j.f)(n,{type:"error",content:"Voting ended by moderator"}),void b.push("/");var e=function(e,t,n,a){var s=new WebSocket("".concat(R,"/").concat(e));s.state=t;var r=function(e){s.state=Object(E.a)(Object(E.a)({},s.state),e),n(e)};return s.onopen=function(){r({actionSenders:J(s),loading:!0})},s.onmessage=function(e){var t=JSON.parse(e.data);if(Array.isArray(t))t.forEach((function(e){A[null===e||void 0===e?void 0:e.action]&&A[e.action]({msg:e,setState:r,conn:s,currentUser:a})}));else{var n=t;A[null===t||void 0===t?void 0:t.action]&&A[t.action]({msg:n,setState:r,conn:s,currentUser:a})}},s.onclose=function(){r(Object(E.a)(Object(E.a)({},t),{},{isRunning:!1,loading:!1}))},s.onerror=function(e){s.close()},s}(a,B,F,k);return function(){e&&e.close()}}),[z]),!ee&&!z&&!Z)return Object(j.f)(n,{type:"error",content:"Already Joined"}),b.push("/"),null;if(ee||S||!Z||!x)return Object(O.jsx)(V.a,{});var re=function(e){ae||(F({lockedOption:U.initial}),null===Z||void 0===Z||Z.submitVote(+e.currentTarget.id))},ie=function(){return b.push("/")};return Object(O.jsxs)("div",{className:m.a.container,children:[Object(O.jsxs)("div",{className:m.a.header,children:[Object(O.jsx)("div",{className:m.a["team-name-wrapper"],title:x.name,children:Object(O.jsx)("strong",{className:m.a["team-name"],children:x.name})}),Object(O.jsx)("div",{className:m.a.subject,children:Object(O.jsx)(f,{onSubmit:function(e){return Z.displaySubject(e)},disableEdit:!Q,initialText:ne,validationRules:{maxLength:200},placeholder:Q?"Enter Voting Subject":"Voting Subject"})}),Object(O.jsx)("div",{className:m.a["admin-btns"],children:Object(O.jsxs)(u.a,{inCase:Q,children:[Object(O.jsxs)(c.Fragment,{children:[Object(O.jsx)(i.a,{variant:"outlined",color:"primary",onClick:function(){return Z.reset()},children:"Reset"}),Object(O.jsxs)(i.a,{variant:"contained",color:"primary",onClick:Z.toggleResultsDisplay,children:[ae?"Hide":"Show"," Results"]}),Object(O.jsx)(_.a,{renderTrigger:function(e){var t=e.open;return Object(O.jsx)(i.a,{variant:"contained",color:"secondary",onClick:t,children:"End"})},action:ie,title:"Are you sure you want to end voting for ".concat(x.name," ?"),acceptText:"Yes",cancelText:"Cancel"})]}),Object(O.jsx)(i.a,{variant:"contained",color:"secondary",onClick:ie,children:"Leave"})]})})]}),Object(O.jsxs)("div",{className:m.a.body,children:[Object(O.jsxs)("div",{className:m.a["members-container"],children:[Object(O.jsx)("div",{className:m.a.admin,children:Object(O.jsx)("div",{className:m.a.moderator,title:null===(e=H[Y])||void 0===e?void 0:e.username,children:Object(O.jsxs)("strong",{children:["Moderator: ",null===(t=H[Y])||void 0===t?void 0:t.username]})})}),Object(O.jsx)("div",{className:m.a.members,children:Object(s.a)(Object.keys(X).map((function(e){return Object(O.jsxs)("div",{className:m.a.member,children:[Object(O.jsxs)("p",{className:m.a["name-container"],title:H[e].username,children:[Q&&Object(O.jsx)("span",{className:m.a["admin-icon"],children:e===Y?Object(O.jsx)(D.a,{title:"Current Moderator"}):Object(O.jsx)(_.a,{renderTrigger:function(e){var t=e.open;return Object(O.jsx)(D.b,{title:"Make Moderator",onClick:t})},action:function(){return null===Z||void 0===Z?void 0:Z.changeModerator(e)},title:"Are you sure you want to make ".concat(H[e].username," the moderator?"),acceptText:"Yes",cancelText:"Cancel"})}),Object(O.jsx)("span",{className:m.a.name,children:H[e].username})]}),I(X[e])]},e)})))})]}),Object(O.jsxs)("div",{className:m.a.cards,children:[T.map((function(e){return Object(O.jsx)("div",{className:se===e?Object(j.b)(m.a,"card","selected-card"):m.a.card,id:e.toString(),onClick:re,children:e},e)})),Object(O.jsx)("div",{className:se===U.unknown?Object(j.b)(m.a,"card","special-card","selected-card"):Object(j.b)(m.a,"card","special-card"),id:U.unknown.toString(),onClick:re,children:"?"},"?")]}),Object(O.jsxs)("div",{className:m.a.results,children:[Object(O.jsxs)("div",{className:m.a.stats,children:[Object(O.jsxs)("strong",{children:["Voted: ",Object.values(X).filter((function(e){return e&&e!==U.initial})).length]}),Object(O.jsxs)("strong",{children:["Joined: ",Object.keys(H).length]})]}),Object(O.jsx)(u.a,{inCase:Boolean(ae),children:Object(O.jsx)(N,{memberVotes:X})})]})]})]})}}}]);
//# sourceMappingURL=7.3e5d2531.chunk.js.map