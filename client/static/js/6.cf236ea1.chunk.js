(this["webpackJsonpvoting-app-webapp"]=this["webpackJsonpvoting-app-webapp"]||[]).push([[6],{246:function(e,t,a){"use strict";var n=a(16),c=a(1),r=a(127),i=a(199),s=a(326),l=a(313),o=a(311),d=a(312),m=a(310),b=a(84),j=a(83),u=a(3);t.a=function(e){var t=e.title,a=e.message,O=e.acceptText,f=e.cancelText,h=e.renderTrigger,p=e.action,x=e.onSuccess,_=e.onError,g=Object(c.useState)(!1),v=Object(n.a)(g,2),y=v[0],N=v[1],w=Object(c.useState)(!1),C=Object(n.a)(w,2),T=C[0],A=C[1],E=function(){return N(!1)};return Object(u.jsxs)(u.Fragment,{children:[Object(u.jsx)(h,{open:function(){return N(!0)},close:E}),Object(u.jsxs)(s.a,{open:y,onClose:E,maxWidth:"sm",fullWidth:!0,disableEscapeKeyDown:T,disableBackdropClick:T,"aria-labelledby":"alert-dialog-title","aria-describedby":"alert-dialog-description",children:[Object(u.jsx)(m.a,{id:"alert-dialog-title",children:t}),a&&Object(u.jsx)(o.a,{children:Object(u.jsx)(d.a,{id:"alert-dialog-description",children:a})}),Object(u.jsxs)(l.a,{children:[Object(u.jsx)(i.a,{onClick:E,color:"secondary",variant:"contained",disabled:T,children:null!==f&&void 0!==f?f:"Cancel"}),Object(u.jsxs)(i.a,{onClick:function(){var e=p();if(!Object(j.c)(e))return E();A(!0),e.then((function(e){x&&x(e)})).catch((function(e){_&&_(e)})).finally((function(){A(!1),E()}))},color:"primary",variant:"contained",autoFocus:!0,disabled:T,children:[Object(u.jsxs)(b.a,{inCase:T,children:[Object(u.jsx)(r.a,{size:20}),Object(u.jsx)(u.Fragment,{children:null!==O&&void 0!==O?O:"Yes"})]}),T]})]})]})]})}},249:function(e,t,a){"use strict";a.d(t,"a",(function(){return r}));var n=a(16),c=a(1);function r(e){var t=Object(c.useReducer)((function(e,t){return Object.assign({},e,t)}),e),a=Object(n.a)(t,2);return[a[0],a[1]]}},253:function(e,t,a){"use strict";var n=a(16),c=a(240),r=a(314),i=a(323),s=a(199),l=a(72),o=a(254),d=a.n(o),m=a(19),b=a(22),j=a(33),u=a(1),O=a(84),f=a(99),h=a(35),p=a(83),x=a(246),_=a(13),g=a(3);t.a=function(e){var t=e.afterSubmitAction,a=e.className,o=e.team,v=void 0===o?{id:"",name:""}:o,y=e.type,N=e.isAdmin,w=void 0!==N&&N,C=Object(b.b)(),T=Object(_.g)(),A=Object(u.useState)({name:v.name,isPublic:v.isPublic}),E=Object(n.a)(A,2),S=E[0],P=E[1],I=Object(u.useState)(S.isPublic),k=Object(n.a)(I,2),M=k[0],R=k[1],B=Object(l.a)({reValidateMode:"onBlur",mode:"onBlur"}),F=B.register,U=B.handleSubmit,H=B.errors,J=F(m.a.teamName),K=Object(m.b)(H,"name",m.a.teamName),V=Object(b.c)((function(e){return e.teams})).data,D=void 0===V?{}:V,Y=Object.values(D);return Object(g.jsxs)("form",{className:d.a.form+" "+a,onSubmit:U((function(e){var a=e.name,n=e.isPublic;if("create"===y)return Y.find((function(e){return e.name===a}))?void Object(p.f)(C,{type:"error",content:"Team name already exists"}):(C(Object(j.b)({name:a,isPublic:n})),void(t&&t()));S.name===a&&S.isPublic===n||(Object(f.a)(h.a.getUrl(v.id),{method:"PATCH",body:JSON.stringify({name:a,isPublic:n})}).then((function(e){P(e),Object(p.f)(C,{content:"Saved Changes",type:"success"})})).catch((function(e){return Object(p.f)(C,{content:e.message,type:"error"})})),t&&t())})),children:[Object(g.jsx)(c.a,{margin:"normal",className:d.a.input,inputRef:J,id:"name",name:"name",label:"Name",type:"text",helperText:K,error:Boolean(K),defaultValue:v.name,disabled:!w,required:!0}),Object(g.jsx)(r.a,{className:d.a.last,label:"Public",control:Object(g.jsx)(i.a,{inputRef:F,name:"isPublic",color:"primary",checked:M,onChange:w?function(){return R(!M)}:void 0}),labelPlacement:"end"}),Object(g.jsx)(O.a,{inCase:w,children:Object(g.jsxs)("div",{className:d.a.btns,children:[Object(g.jsx)(s.a,{className:d.a.btn,variant:"contained",color:"primary",type:"submit",children:"Save"}),Object(g.jsx)(O.a,{inCase:"update"===y,children:Object(g.jsx)(x.a,{renderTrigger:function(e){var t=e.open;return Object(g.jsx)(s.a,{className:d.a.delete,color:"secondary",onClick:t,variant:"contained",children:"Delete"})},action:function(){C(Object(j.c)(v.id)),T.push("/")},title:"Are you sure you want to delete team ".concat(S.name,"?"),acceptText:"Yes",cancelText:"Cancel"})})]})})]})}},254:function(e,t,a){e.exports={"fs-lg":"team-form_fs-lg__KXAfg",form:"team-form_form__woAQs",msg:"team-form_msg__3hKwO",input:"team-form_input__1OECB",last:"team-form_last__3kpfG",btn:"team-form_btn__EPysp",btns:"team-form_btns__2lhyx",delete:"team-form_delete__1ZJSP"}},255:function(e,t,a){"use strict";a.d(t,"a",(function(){return b}));var n=a(34),c=a(104),r=a(316),i=a(315),s=a(266),l=a.n(s),o=a(256),d=a.n(o),m=a(3),b=l()({root:{fontSize:"1.2rem"}})(i.a);t.b=function(e){var t=e.isEmpty,a=e.emptyTableMessage,i=void 0===a?"No Data":a,s=Object(c.a)(e,["isEmpty","emptyTableMessage"]);return t?Object(m.jsx)("div",{className:d.a.empty,children:Object(m.jsx)("p",{children:i})}):Object(m.jsx)(r.a,Object(n.a)(Object(n.a)({className:d.a.table},s),{},{children:e.children}))}},256:function(e,t,a){e.exports={"fs-lg":"custom-table_fs-lg__3qbwE",table:"custom-table_table__26eiR",empty:"custom-table_empty__3ZLMQ"}},301:function(e,t,a){e.exports={"fs-lg":"members-table_fs-lg__2KVUd",you:"members-table_you__3bHDK",wrapper:"members-table_wrapper__1nA4o"}},302:function(e,t,a){e.exports={"fs-lg":"team-details_fs-lg__3oFHE",team:"team-details_team__3TIIc",input:"team-details_input__2fQ9V",title:"team-details_title__1i8bL","form-wrapper":"team-details_form-wrapper__1aBgq","team-form":"team-details_team-form__3FW--","name-label":"team-details_name-label__RcBwG",public:"team-details_public__XSOzF",delete:"team-details_delete__1mHrV","delete-icon":"team-details_delete-icon__1ClJ8"}},325:function(e,t,a){"use strict";a.r(t);var n=a(16),c=a(22),r=a(13),i=a(34),s=a(317),l=a(2),o=a(4),d=a(1),m=(a(7),a(5)),b=a(6),j=a(250),u={variant:"head"},O="thead",f=d.forwardRef((function(e,t){var a=e.classes,n=e.className,c=e.component,r=void 0===c?O:c,i=Object(o.a)(e,["classes","className","component"]);return d.createElement(j.a.Provider,{value:u},d.createElement(r,Object(l.a)({className:Object(m.a)(a.root,n),ref:t,role:r===O?null:"rowgroup"},i)))})),h=Object(b.a)({root:{display:"table-header-group"}},{name:"MuiTableHead"})(f),p=a(318),x=a(199),_=a(84),g=a(301),v=a.n(g),y=a(255),N=a(99),w=a(35),C=a(33),T=a(246),A=a(83),E=a(3),S=function(e){var t=e.allMembers,a=e.isAdmin,r=e.adminIds,l=e.teamId,o=e.currentUser,m=Object(d.useState)(r),b=Object(n.a)(m,2),j=b[0],u=b[1],O=t.findIndex((function(e){return e.id===o.id}));Object(A.d)(t,O);var f=Object(c.b)();return Object(E.jsx)("div",{className:v.a.wrapper,children:Object(E.jsxs)(y.b,{isEmpty:0===t.length,emptyTableMessage:"There are no members in this team.",children:[Object(E.jsx)(h,{children:Object(E.jsxs)(p.a,{children:[Object(E.jsx)(y.a,{align:"left",children:Object(E.jsx)("strong",{children:"Username"})},"username"),Object(E.jsx)(y.a,{align:"left",children:Object(E.jsx)("strong",{children:"Email"})},"email"),Object(E.jsx)(y.a,{align:"left",children:Object(E.jsx)("strong",{children:"Role"})},"role"),Object(E.jsx)(_.a,{inCase:Boolean(a),children:Object(E.jsx)(y.a,{align:"left",children:Object(E.jsx)("strong",{children:"Actions"})},"actions")})]})}),Object(E.jsx)(s.a,{children:t.map((function(e){return Object(E.jsxs)(p.a,{children:[Object(E.jsx)(y.a,{children:e.username||e.name},"username"),Object(E.jsx)(y.a,{align:"left",children:e.email},"email"),Object(E.jsx)(y.a,{align:"left",children:j[e.id]?"Admin":"Member"},"role"),Object(E.jsx)(_.a,{inCase:Boolean(a),children:Object(E.jsx)(y.a,{align:"left",children:Object(E.jsxs)(_.a,{inCase:e.id!==o.id,children:[Object(E.jsxs)(E.Fragment,{children:[Object(E.jsx)(x.a,{classes:{root:"mr-3"},variant:"contained",color:"primary",onClick:function(){return function(e){var t=j[e];Object(N.a)(w.a.toggleRole,{method:"PATCH",body:JSON.stringify({userId:e,teamId:l})}).then((function(){var a=Object(i.a)({},j);t?delete a[e]:a[e]=!0,u(a)})).catch((function(){Object(A.f)(f,{type:"error",content:"Error while changing user role"})}))}(e.id)},children:"Toggle Role"}),Object(E.jsx)(T.a,{renderTrigger:function(e){var t=e.open;return Object(E.jsx)(x.a,{classes:{root:"mr-3"},variant:"contained",color:"secondary",onClick:t,children:"Remove"})},action:function(){return t=e.id,void f(Object(C.n)({userId:t,teamId:l,wasAdmin:j[t]}));var t},title:"Are you sure you want to remove ".concat(e.username,"?"),acceptText:"Yes",cancelText:"Cancel"})]}),Object(E.jsx)("p",{className:v.a.you,children:"YOU"})]})},"actions")})]},e.username||e.name)}))})]})})},P=a(66),I=a(302),k=a.n(I),M=a(249),R=a(253);t.default=function(){var e=Object(r.g)(),t=(Object(c.b)(),Object(r.h)().teamId),a=Object(M.a)({}),i=Object(n.a)(a,2),s=i[0],l=s.team,o=s.loading,m=void 0===o||o,b=i[1];Object(d.useEffect)((function(){Object(N.a)(w.a.getUrl(t)).then((function(e){b({loading:!1,team:e})})).catch((function(){b({loading:!1,teamApiError:"Error while fetching team"})}))}),[t]);var j=Object(c.c)((function(e){return e.user})),u=j.data,O=j.loading;if(m||O||!u)return Object(E.jsx)(P.a,{});if(!l)return e.push("/"),null;var f=l.admins,h=void 0===f?[]:f,p=l.members,x=void 0===p?[]:p,_=Object(A.a)(h,"id"),g=_[(null===u||void 0===u?void 0:u.id)||""],v=h.concat(x);return Object(E.jsxs)("div",{className:k.a.team,children:[Object(E.jsx)("h3",{className:k.a.title,children:"General"}),Object(E.jsx)("div",{className:k.a["form-wrapper"],children:Object(E.jsx)(R.a,{className:k.a["team-form"],type:"update",team:l,isAdmin:g})}),Object(E.jsx)("h3",{className:k.a.title,children:"Members"}),Object(E.jsx)(S,{teamId:l.id,allMembers:v||[],adminIds:_,isAdmin:g,currentUser:u})]})}}}]);
//# sourceMappingURL=6.cf236ea1.chunk.js.map