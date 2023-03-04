(this["webpackJsonpvoting-app-webapp"]=this["webpackJsonpvoting-app-webapp"]||[]).push([[8],{257:function(t,e,n){"use strict";var a=n(130),r=n(132);Object.defineProperty(e,"__esModule",{value:!0}),e.default=void 0;var i=r(n(1)),l=(0,a(n(133)).default)(i.createElement("path",{d:"M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"}),"CheckCircle");e.default=l},306:function(t,e,n){!function(t,e){"use strict";function n(t){return t&&"object"===typeof t&&"default"in t?t:{default:t}}var a=n(e);function r(t){return t*Math.PI/180}function i(t,e,n){return t>n?n:t<e?e:t}function l(t,e){return e/100*t}function u(t,e){return t+e/2}function o(t,e){var n=r(t);return{dx:e*Math.cos(n),dy:e*Math.sin(n)}}function s(t){return"number"===typeof t}function d(t,e){return"function"===typeof t?t(e):t}function c(t){for(var e=0,n=0;n<t.length;n++)e+=t[n].value;return e}function f(t){for(var e=t.data,n=t.lengthAngle,a=t.totalValue,r=t.paddingAngle,u=t.startAngle,o=a||c(e),s=i(n,-360,360),d=360===Math.abs(s)?e.length:e.length-1,f=Math.abs(r)*Math.sign(n),h=s-f*d,g=0,v=[],m=0;m<e.length;m++){var b=e[m],p=0===o?0:b.value/o*100,y=l(h,p),M=g+u;g=g+y+f,v.push(Object.assign({percentage:p,startAngle:M,degrees:y},b))}return v}function h(t,e){if(null==t)return{};var n,a,r={},i=Object.keys(t);for(a=0;a<i.length;a++)n=i[a],e.indexOf(n)>=0||(r[n]=t[n]);return r}function g(t){t.dataEntry,t.dataIndex;var e=h(t,["dataEntry","dataIndex"]);return a.default.createElement("text",Object.assign({dominantBaseline:"central"},e))}function v(t){var e=1e14;return Math.round((t+Number.EPSILON)*e)/e}function m(t){var e=t.labelPosition,n=t.lineWidth,a=v(t.labelHorizontalShift);return 0===a?"middle":e>100?a>0?"start":"end":e<100-n?a>0?"end":"start":"middle"}function b(t,e){var n=t(e);return"string"===typeof n||"number"===typeof n?a.default.createElement(g,Object.assign({key:"label-"+(e.dataEntry.key||e.dataIndex)},e),n):a.default.isValidElement(n)?n:null}function p(t,e){return t.map((function(t,n){var a,r=null!=(a=d(e.segmentsShift,n))?a:0,i=l(e.radius,e.labelPosition)+r,s=o(u(t.startAngle,t.degrees),i),c=s.dx,f=s.dy,h={x:e.center[0],y:e.center[1],dx:c,dy:f,textAnchor:m({labelPosition:e.labelPosition,lineWidth:e.lineWidth,labelHorizontalShift:c}),dataEntry:t,dataIndex:n,style:d(e.labelStyle,n)};return e.label&&b(e.label,h)}))}var y=function(t,e,n,a,r){var i=r-a;if(0===i)return[];var l=n*Math.cos(a)+t,u=n*Math.sin(a)+e,o=n*Math.cos(r)+t,s=n*Math.sin(r)+e;return[["M",l,u],["A",n,n,0,Math.abs(i)<=Math.PI?"0":"1",i<0?"0":"1",o,s]]};function M(t,e,n,a,l){var u=i(a,-359.999,359.999);return y(t,e,l,r(n),r(n+u)).map((function(t){return t.join(" ")})).join(" ")}function x(t){var e,n,i=t.cx,d=t.cy,c=t.lengthAngle,f=t.lineWidth,g=t.radius,v=t.shift,m=void 0===v?0:v,b=t.reveal,p=t.rounded,y=t.startAngle,x=t.title,A=h(t,["cx","cy","lengthAngle","lineWidth","radius","shift","reveal","rounded","startAngle","title"]),k=g-f/2,z=o(u(y,c),m),E=M(i+z.dx,d+z.dy,y,c,k);if(s(b)){var O=r(k)*c;n=(e=Math.abs(O))-l(e,b)}return a.default.createElement("path",Object.assign({d:E,fill:"none",strokeWidth:f,strokeDasharray:e,strokeDashoffset:n,strokeLinecap:p?"round":void 0},A),x&&a.default.createElement("title",null,x))}function A(t,e,n){var a="stroke-dashoffset "+t+"ms "+e;return n&&n.transition&&(a=a+","+n.transition),{transition:a}}function k(t){return t.animate&&!s(t.reveal)?100:t.reveal}function z(t,e){return t&&function(n){t(n,e)}}function E(t,e,n){var r=null!=n?n:k(e),i=e.radius,u=e.center,o=u[0],s=u[1],c=l(i,e.lineWidth),f=t.map((function(t,n){var l=d(e.segmentsStyle,n);return a.default.createElement(x,{cx:o,cy:s,key:t.key||n,lengthAngle:t.degrees,lineWidth:c,radius:i,rounded:e.rounded,reveal:r,shift:d(e.segmentsShift,n),startAngle:t.startAngle,title:t.title,style:Object.assign({},l,e.animate&&A(e.animationDuration,e.animationEasing,l)),stroke:t.color,tabIndex:e.segmentsTabIndex,onBlur:z(e.onBlur,n),onClick:z(e.onClick,n),onFocus:z(e.onFocus,n),onKeyDown:z(e.onKeyDown,n),onMouseOver:z(e.onMouseOver,n),onMouseOut:z(e.onMouseOut,n)})}));return e.background&&f.unshift(a.default.createElement(x,{cx:o,cy:s,key:"bg",lengthAngle:e.lengthAngle,lineWidth:c,radius:i,rounded:e.rounded,startAngle:e.startAngle,stroke:e.background})),f}var O={animationDuration:500,animationEasing:"ease-out",center:[50,50],data:[],labelPosition:50,lengthAngle:360,lineWidth:100,paddingAngle:0,radius:50,startAngle:0,viewBoxSize:[100,100]};function j(t){var n=e.useState(t.animate?0:null),r=n[0],i=n[1];e.useEffect((function(){if(t.animate)return e();function e(){var t,e;return t=setTimeout((function(){t=null,e=requestAnimationFrame((function(){e=null,i(null)}))})),function(){t&&clearTimeout(t),e&&cancelAnimationFrame(e)}}}),[]);var l=f(t);return a.default.createElement("svg",{viewBox:"0 0 "+t.viewBoxSize[0]+" "+t.viewBoxSize[1],width:"100%",height:"100%",className:t.className,style:t.style},E(l,t,r),t.label&&p(l,t),t.children)}j.defaultProps=O,t.PieChart=j,Object.defineProperty(t,"__esModule",{value:!0})}(e,n(1))},307:function(t,e,n){"use strict";n.d(e,"a",(function(){return r})),n.d(e,"b",(function(){return i}));var a=n(0);function r(t){return Object(a.a)({tag:"svg",attr:{viewBox:"0 0 24 24"},child:[{tag:"g",attr:{},child:[{tag:"path",attr:{fill:"none",d:"M0 0h24v24H0z"}},{tag:"path",attr:{d:"M12 14v8H4a8 8 0 0 1 8-8zm0-1c-3.315 0-6-2.685-6-6s2.685-6 6-6 6 2.685 6 6-2.685 6-6 6zm9 4h1v5h-8v-5h1v-1a3 3 0 0 1 6 0v1zm-2 0v-1a1 1 0 0 0-2 0v1h2z"}}]}]})(t)}function i(t){return Object(a.a)({tag:"svg",attr:{viewBox:"0 0 24 24"},child:[{tag:"g",attr:{},child:[{tag:"path",attr:{fill:"none",d:"M0 0h24v24H0z"}},{tag:"path",attr:{d:"M12 14v2a6 6 0 0 0-6 6H4a8 8 0 0 1 8-8zm0-1c-3.315 0-6-2.685-6-6s2.685-6 6-6 6 2.685 6 6-2.685 6-6 6zm0-2c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm9 6h1v5h-8v-5h1v-1a3 3 0 0 1 6 0v1zm-2 0v-1a1 1 0 0 0-2 0v1h2z"}}]}]})(t)}}}]);
//# sourceMappingURL=8.23b45e0a.chunk.js.map