!function(l){function t(t){for(var e,r,n=t[0],i=t[1],a=t[2],s=0,o=[];s<n.length;s++)r=n[s],Object.prototype.hasOwnProperty.call(u,r)&&u[r]&&o.push(u[r][0]),u[r]=0;for(e in i)Object.prototype.hasOwnProperty.call(i,e)&&(l[e]=i[e]);for(p&&p(t);o.length;)o.shift()();return d.push.apply(d,a||[]),c()}function c(){for(var t,e=0;e<d.length;e++){for(var r=d[e],n=!0,i=1;i<r.length;i++){var a=r[i];0!==u[a]&&(n=!1)}n&&(d.splice(e--,1),t=s(s.s=r[0]))}return t}var r={},u={0:0},d=[];function s(t){if(r[t])return r[t].exports;var e=r[t]={i:t,l:!1,exports:{}};return l[t].call(e.exports,e,e.exports,s),e.l=!0,e.exports}s.m=l,s.c=r,s.d=function(t,e,r){s.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:r})},s.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},s.t=function(e,t){if(1&t&&(e=s(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var r=Object.create(null);if(s.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)s.d(r,n,function(t){return e[t]}.bind(null,n));return r},s.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return s.d(e,"a",e),e},s.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},s.p="/assets/js/";var e=window.webpackJsonp=window.webpackJsonp||[],n=e.push.bind(e);e.push=t,e=e.slice();for(var i=0;i<e.length;i++)t(e[i]);var p=n;d.push([30,1]),c()}({30:function(t,e,r){"use strict";r.r(e);var n=r(7),i=r(29),a=(r(31),r(34),n.a.start()),s=r(36);a.load(Object(i.a)(s))},31:function(t,e,r){var n=r(32);"string"==typeof n&&(n=[[t.i,n,""]]);var i={hmr:!0,transform:void 0,insertInto:void 0};r(20)(n,i);n.locals&&(t.exports=n.locals)},32:function(t,e,r){(t.exports=r(19)(!1)).push([t.i,".body-block{font-size:5px}\n",""])},34:function(t,e,r){var n=r(35);"string"==typeof n&&(n=[[t.i,n,""]]);var i={hmr:!0,transform:void 0,insertInto:void 0};r(20)(n,i);n.locals&&(t.exports=n.locals)},35:function(t,e,r){(t.exports=r(19)(!1)).push([t.i,'.tile{width:100%;display:inline-block;box-sizing:border-box;background:#fff;padding:20px;margin-bottom:10px}.tile .title{margin-top:0}.tile p,.tile span,.tile a,.tile ul,.tile li,.tile button{font-family:inherit;font-size:inherit;font-weight:inherit;line-height:inherit}.tile h1,.tile h2,.tile h3,.tile h4,.tile h5,.tile h6{font-family:"Open Sans","Segoe UI",Frutiger,"Frutiger Linotype","Dejavu Sans","Helvetica Neue",Arial,sans-serif;line-height:1.5em;font-weight:300}.tile strong{font-weight:400}.tile.purple,.tile.blue,.tile.red,.tile.orange,.tile.green{color:#fff}.tile.gray{background:#eee;color:black}.tile.gray:hover{background:#e1e1e1}.tile.purple{background:#5133ab}.tile.purple:hover{background:#3e2784}.tile.red{background:#ac193d}.tile.red:hover{background:#7f132d}.tile.green{background:#00a600}.tile.green:hover{background:#007300}.tile.blue{background:#2672ec}.tile.blue:hover{background:#125acd}.tile.orange{background:#dc572e}.tile.orange:hover{background:#b8431f}\n',""])},36:function(t,e,r){var n={"./layout_controller.js":37,"./payment_controller.js":59};function i(t){var e=a(t);return r(e)}function a(t){if(r.o(n,t))return n[t];var e=new Error("Cannot find module '"+t+"'");throw e.code="MODULE_NOT_FOUND",e}i.keys=function(){return Object.keys(n)},i.resolve=a,(t.exports=i).id=36},37:function(t,e,r){"use strict";r.r(e),r.d(e,"default",function(){return g});var n=r(9),s=r.n(n),i=r(10),o=r.n(i),a=r(11),l=r.n(a),c=r(12),u=r.n(c),d=r(4),p=r.n(d),f=r(13),h=r.n(f),b=r(5),v=r.n(b),g=function(t){function a(){var t,e;s()(this,a);for(var r=arguments.length,n=new Array(r),i=0;i<r;i++)n[i]=arguments[i];return e=l()(this,(t=u()(a)).call.apply(t,[this].concat(n))),v()(p()(e),"navbarIsOpened",void 0),e}return h()(a,t),o()(a,[{key:"toggleNavbar",value:function(){this.navbarIsOpened?$(this.navbarTarget).slideUp():$(this.navbarTarget).slideDown(),this.navbarIsOpened=!this.navbarIsOpened}}],[{key:"targets",get:function(){return["navbar"]}}]),a}(r(7).b)},59:function(t,e,r){"use strict";r.r(e);var n=r(8),o=r.n(n),i=r(17),l=r.n(i),a=r(9),s=r.n(a),c=r(10),u=r.n(c),d=r(11),p=r.n(d),f=r(12),h=r.n(f),b=r(4),v=r.n(b),g=r(13),y=r.n(g),m=r(5),T=r.n(m),w=r(7),k=r(18),x=r.n(k),j=function(t){t.classList.add("d-none"),t.classList.add("d-hide")},_=function(t){t.classList.remove("d-none"),t.classList.remove("d-hide")};var O=r(81);r.d(e,"default",function(){return L});var L=function(t){function a(){var t,e;s()(this,a);for(var r=arguments.length,n=new Array(r),i=0;i<r;i++)n[i]=arguments[i];return e=p()(this,(t=h()(a)).call.apply(t,[this].concat(n))),T()(v()(e),"list",void 0),T()(v()(e),"periods",void 0),T()(v()(e),"subjects",void 0),T()(v()(e),"classes",void 0),e}var e;return y()(a,t),u()(a,[{key:"connect",value:function(){this.list=[],this.periods=[],this.classes=[],this.subjects=[];var e=this;Array.prototype.forEach.call(this.periodTarget.options,function(t){""!==t.value&&e.periods.push({id:t.value,label:t.innerText})}),Array.prototype.forEach.call(this.subjectTarget.options,function(t){""!==t.value&&e.subjects.push({id:t.value,label:t.innerText})}),Array.prototype.forEach.call(this.classTarget.options,function(t){""!==t.value&&e.classes.push({id:t.value,label:t.innerText})})}},{key:"addToList",value:function(){var t=this.subjectTarget.value,e=this.periodTarget.value,r=this.classTarget.value;if(""!==t&&""!==e&&""!==r){var n=t+e+r;O.a.find(this.list,function(t){return t.id===n})?window.alert("Select subject exists for the specified period and class"):(this.list.push({id:n,subject_id:t,period_id:e,class_id:r,subject:this.subjectTarget.options[this.subjectTarget.selectedIndex].innerText,period:this.periodTarget.options[this.periodTarget.selectedIndex].innerText,className:this.classTarget.options[this.classTarget.selectedIndex].innerText}),this.displayList())}else window.alert("Subject, period amd class are required")}},{key:"removeFromList",value:function(t){var e=t.currentTarget.getAttribute("data-id");O.a.remove(this.list,function(t){return t.id===e}),this.displayList()}},{key:"displayList",value:function(){var i=this;this.listTblTarget.innerHTML="",this.cartTotal=0,this.list.forEach(function(t,e){var r=document.importNode(i.itemTemplateTarget.content,!0),n=r.querySelectorAll("td");n[0].innerText=e+1,n[1].innerText=t.className,n[2].innerText=t.subject,n[3].innerHTML=t.period,n[4].innerHTML='<button data-action="click->payment#removeFromList" data-id="'.concat(t.id,'">Remove</button>'),i.listTblTarget.appendChild(r)});var t=this.list.length;if(5<=t){var e=(t-t%5)/5;this.cartTotal=3e4*e,t-=5*e}if(3<=t){var r=(t-t%3)/3;this.cartTotal+=2e4*r,t-=3*r}this.cartTotal+=8e3*t,this.cartTotalTarget.textContent=this.cartTotal;var n=8e3*this.list.length-this.cartTotal,a=100*n/this.cartTotal;this.savingsTarget.textContent="".concat(n," (").concat(a.toFixed(0),"%)"),(0<this.list.length?_:j)(this.cartItemDivTarget)}},{key:"subscribe",value:(e=l()(o.a.mark(function t(){var e,s,r,n;return o.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:if(0===this.list.length)return window.alert("Please add a subject, period amd class to continue"),t.abrupt("return");t.next=3;break;case 3:if(this.loading)return t.abrupt("return");t.next=5;break;case 5:return this.loading=!0,e={count:this.list.length,subject_id:this.list[0].subject_id,class_id:this.list[0].class_id,period_id:this.list[0].period_id},s=this,t.next=10,x.a.post("/payments/initiate",e);case 10:if(r=t.sent,(n=r.data).error)return window.alert(n.error),t.abrupt("return");t.next=15;break;case 15:PaystackPop.setup({key:"pk_test_6e0424ff8e08138c2fcce3f4f2b05052c4b3f77c",email:n.student.parent_email,amount:n.amount,currency:"NGN",ref:n.id,metadata:{custom_fields:[{display_name:"Mobile Number",variable_name:"mobile_number",value:n.student.parent_phone}]},callback:function(){var e=l()(o.a.mark(function t(e){var r,n,i,a;return o.a.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:for(r=[],n=0;n<s.list.length;n++)r.push({subject_id:s.list[n].subject_id,class_id:s.list[n].class_id,period_id:s.list[n].period_id});return t.next=4,x.a.post("/payments/".concat(e.reference,"/update-status"),{items:r});case 4:i=t.sent,(a=i.data).error&&window.alert(a.error),window.alert("Success. Subscription successful, check the subscription table for your lesson start date"),window.location.reload();case 9:case"end":return t.stop()}},t)}));return function(t){return e.apply(this,arguments)}}(),onClose:function(){window.alert("window closed")}}).openIframe();case 17:case"end":return t.stop()}},t,this)})),function(){return e.apply(this,arguments)})}],[{key:"targets",get:function(){return["class","subject","period","cartItemDiv","listTbl","itemTemplate","cartTotal","savings"]}}]),a}(w.b)}});
//# sourceMappingURL=app.bundle.js.map