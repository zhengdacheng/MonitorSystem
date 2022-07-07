(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-554a9eaf"],{"00b4":function(e,t,r){"use strict";r("ac1f");var a=r("23e7"),i=r("da84"),n=r("c65b"),s=r("e330"),l=r("1626"),o=r("861d"),c=function(){var e=!1,t=/[ac]/;return t.exec=function(){return e=!0,/./.exec.apply(this,arguments)},!0===t.test("abc")&&e}(),u=i.Error,h=s(/./.test);a({target:"RegExp",proto:!0,forced:!c},{test:function(e){var t=this.exec;if(!l(t))return h(this,e);var r=n(t,this,e);if(null!==r&&!o(r))throw new u("RegExp exec method returned something other than an Object or null");return!!r}})},1017:function(e,t,r){"use strict";r.r(t);var a=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{attrs:{id:"setting"}},[e._m(0),r("div",{staticClass:"setting-wrapper-box"},[r("div",{staticClass:"section-title"},[e._v(" 配置告警规则 ")]),r("div",{staticClass:"section-content"},[r("alarm-setting")],1),r("div",{staticClass:"button-line"})]),r("div",{staticClass:"item-wrapper-box"},[r("div",{staticClass:"section-title"},[e._v(" 告警规则列表 ")]),r("div",{staticClass:"section-content"},[r("alarm-list")],1)])])},i=[function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"tips"},[r("i",{staticClass:"el-icon-warning-outline"}),r("span",[e._v("在此处可以自定义配置告警阈值，联系邮箱等")])])}],n=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{attrs:{id:"setting-wrapper"}},[r("div",{staticClass:"configure-wrapper"},[r("div",{staticClass:"label-group"},[e._v("cpu使用率阈值")]),r("div",{staticClass:"input-group"},[r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"提示（0-1）",size:"small"},model:{value:e.configure.cpu_threshold.warning,callback:function(t){e.$set(e.configure.cpu_threshold,"warning",t)},expression:"configure.cpu_threshold.warning"}}),r("div",{staticClass:"split"},[e._v(" ＜ ")]),r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"严重（0-1）",size:"small"},model:{value:e.configure.cpu_threshold.serious,callback:function(t){e.$set(e.configure.cpu_threshold,"serious",t)},expression:"configure.cpu_threshold.serious"}}),r("div",{staticClass:"split"},[e._v(" ＜ ")]),r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"致命（0-1）",size:"small"},model:{value:e.configure.cpu_threshold.mortal,callback:function(t){e.$set(e.configure.cpu_threshold,"mortal",t)},expression:"configure.cpu_threshold.mortal"}})],1)]),r("div",{staticClass:"configure-wrapper"},[r("div",{staticClass:"label-group"},[e._v("内存使用率阈值")]),r("div",{staticClass:"input-group"},[r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"提示（0-1）",size:"small"},model:{value:e.configure.memory_threshold.warning,callback:function(t){e.$set(e.configure.memory_threshold,"warning",t)},expression:"configure.memory_threshold.warning"}}),r("div",{staticClass:"split"},[e._v(" ＜ ")]),r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"严重（0-1）",size:"small"},model:{value:e.configure.memory_threshold.serious,callback:function(t){e.$set(e.configure.memory_threshold,"serious",t)},expression:"configure.memory_threshold.serious"}}),r("div",{staticClass:"split"},[e._v(" ＜ ")]),r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"致命（0-1）",size:"small"},model:{value:e.configure.memory_threshold.mortal,callback:function(t){e.$set(e.configure.memory_threshold,"mortal",t)},expression:"configure.memory_threshold.mortal"}})],1)]),r("div",{staticClass:"configure-wrapper"},[r("div",{staticClass:"label-group"},[e._v("监控粒度")]),r("el-select",{staticClass:"input-group",staticStyle:{width:"150px"},attrs:{placeholder:"请选择监控间隔",size:"small"},model:{value:e.configure.time_granularity,callback:function(t){e.$set(e.configure,"time_granularity",t)},expression:"configure.time_granularity"}},e._l(e.granularity_options,(function(e){return r("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1),r("div",{staticClass:"configure-wrapper"},[r("div",{staticClass:"label-group"},[e._v("联系邮箱")]),r("el-input",{staticStyle:{width:"550px"},attrs:{placeholder:"请输入告警联系邮箱，可填写多个，用分号分隔,如XXX@163.com;XXX@qq.com",size:"small"},model:{value:e.configure.email,callback:function(t){e.$set(e.configure,"email",t)},expression:"configure.email"}})],1),r("div",{staticClass:"configure-wrapper"},[r("div",{staticClass:"label-group"}),r("el-button",{attrs:{type:"primary"},on:{click:e.addConfigure}},[e._v("添加配置")]),r("el-button",{on:{click:e.resetConfigure}},[e._v("重置")])],1)])},s=[],l=r("1da1"),o=(r("ac1f"),r("00b4"),r("5319"),r("1276"),r("96cf"),r("fd03")),c={data:function(){return{type:"",configure:{cpu_threshold:{warning:"",serious:"",mortal:""},memory_threshold:{warning:"",serious:"",mortal:""},time_granularity:"",email:""},email_arr:[],granularity_options:[{value:1,label:"1分钟"},{value:2,label:"2分钟"},{value:5,label:"5分钟"},{value:10,label:"10分钟"}]}},mounted:function(){},methods:{addConfigure:function(){var e=this;return Object(l["a"])(regeneratorRuntime.mark((function t(){var r,a;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(e.threshold_verify(e.configure.cpu_threshold||!e.threshold_verify(e.configure.memory_threshold))){t.next=4;break}return t.abrupt("return");case 4:if(""!=e.configure.time_granularity){t.next=8;break}e.$message.warning("请选择监控粒度"),t.next=14;break;case 8:if(!e.email_verify(e.configure.email)){t.next=14;break}return r={cpu_threshold:e.configure.cpu_threshold,memory_threshold:e.configure.memory_threshold,time_granularity:e.configure.time_granularity,email:e.email_arr},t.next=12,Object(o["b"])(r);case 12:a=t.sent,200==a.data.code?e.$message.success("成功添加配置"):e.$message.error(a.data.msg);case 14:case"end":return t.stop()}}),t)})))()},email_verify:function(e){if(0==e.length)return this.$message.warning("联系邮箱不能为空"),!1;var t=/\uff1b/;t.test(e)&&(e=e.replace(/\uff1b/g,";"));for(var r=e.split(";"),a=/^([a-zA-Z0-9])[a-zA-Z0-9]+@[a-zA-Z0-9]+\.([a-zA-Z0-9]{2,5})$/,i=0;i<r.length;i++){if(!a.test(r[i]))return this.$message.warning("存在不合法邮箱格式"),!1;console.log("true")}return this.email_arr=r,!0},threshold_verify:function(e){var t=e.warning,r=e.serious,a=e.mortal;return""==t||""==r||""==a?(this.$message.warning("阈值设置不能为空"),!1):t<0||t>1||r<0||r>1||a<0||a>1?(this.$message.warning("阈值数值请设置在0-1之间"),!1):!(t>r||t>a||r>a)||(this.$message.warning("阈值区间设置有误"),!1)},resetConfigure:function(){this.configure={cpu_threshold:{warning:"",serious:"",mortal:""},memory_threshold:{warning:"",serious:"",mortal:""},time_granularity:"",email:""}}}},u=c,h=(r("c5d9"),r("2877")),d=Object(h["a"])(u,n,s,!1,null,"76a492bf",null),p=d.exports,f=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{attrs:{id:"alarm-list"}},[r("el-table",{attrs:{data:e.tableData,"header-cell-style":{background:"#eef1f6",color:"#606266",border:"1px solid #DFE1E4"}}},[r("el-table-column",{attrs:{prop:"metric",label:"监控指标"}}),r("el-table-column",{attrs:{label:"阈值","header-align":"center"}},[r("el-table-column",{attrs:{prop:"threshold.warning",label:"提示"}}),r("el-table-column",{attrs:{prop:"threshold.serious",label:"严重"}}),r("el-table-column",{attrs:{prop:"threshold.mortal",label:"致命"}})],1),r("el-table-column",{attrs:{prop:"time_granularity",label:"监控粒度"}}),r("el-table-column",{attrs:{prop:"email",label:"联系邮箱"}})],1)],1)},g=[],m={data:function(){return{tableData:[]}},mounted:function(){this.initData()},methods:{initData:function(){var e=this;return Object(l["a"])(regeneratorRuntime.mark((function t(){var r,a,i,n,s,l,c,u;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,Object(o["a"])();case 2:if(r=t.sent,a=r.data,200==a.code){t.next=7;break}return console.log(a.msg),t.abrupt("return",!1);case 7:for(i=0;i<a.data.length;i++)n=a.data[i],s=n.cpu_threshold,l=n.memory_threshold,c=n.time_granularity,u=n.email,e.tableData.push({metric:"cpu使用率",threshold:s,time_granularity:c,email:u}),e.tableData.push({metric:"内存使用率",threshold:l,time_granularity:c,email:u});case 8:case"end":return t.stop()}}),t)})))()}}},v=m,b=(r("c932"),Object(h["a"])(v,f,g,!1,null,"51e82f8c",null)),_=b.exports,x={components:{AlarmSetting:p,AlarmList:_},data:function(){return{basePath:"",activateName:this.$route.query.type}},mounted:function(){},methods:{}},y=x,w=(r("e975"),Object(h["a"])(y,a,i,!1,null,"09385f01",null));t["default"]=w.exports},1276:function(e,t,r){"use strict";var a=r("2ba4"),i=r("c65b"),n=r("e330"),s=r("d784"),l=r("44e7"),o=r("825a"),c=r("1d80"),u=r("4840"),h=r("8aa5"),d=r("50c4"),p=r("577e"),f=r("dc4a"),g=r("4dae"),m=r("14c3"),v=r("9263"),b=r("9f7f"),_=r("d039"),x=b.UNSUPPORTED_Y,y=4294967295,w=Math.min,C=[].push,k=n(/./.exec),$=n(C),E=n("".slice),z=!_((function(){var e=/(?:)/,t=e.exec;e.exec=function(){return t.apply(this,arguments)};var r="ab".split(e);return 2!==r.length||"a"!==r[0]||"b"!==r[1]}));s("split",(function(e,t,r){var n;return n="c"=="abbc".split(/(b)*/)[1]||4!="test".split(/(?:)/,-1).length||2!="ab".split(/(?:ab)*/).length||4!=".".split(/(.?)(.?)/).length||".".split(/()()/).length>1||"".split(/.?/).length?function(e,r){var n=p(c(this)),s=void 0===r?y:r>>>0;if(0===s)return[];if(void 0===e)return[n];if(!l(e))return i(t,n,e,s);var o,u,h,d=[],f=(e.ignoreCase?"i":"")+(e.multiline?"m":"")+(e.unicode?"u":"")+(e.sticky?"y":""),m=0,b=new RegExp(e.source,f+"g");while(o=i(v,b,n)){if(u=b.lastIndex,u>m&&($(d,E(n,m,o.index)),o.length>1&&o.index<n.length&&a(C,d,g(o,1)),h=o[0].length,m=u,d.length>=s))break;b.lastIndex===o.index&&b.lastIndex++}return m===n.length?!h&&k(b,"")||$(d,""):$(d,E(n,m)),d.length>s?g(d,0,s):d}:"0".split(void 0,0).length?function(e,r){return void 0===e&&0===r?[]:i(t,this,e,r)}:t,[function(t,r){var a=c(this),s=void 0==t?void 0:f(t,e);return s?i(s,t,a,r):i(n,p(a),t,r)},function(e,a){var i=o(this),s=p(e),l=r(n,i,s,a,n!==t);if(l.done)return l.value;var c=u(i,RegExp),f=i.unicode,g=(i.ignoreCase?"i":"")+(i.multiline?"m":"")+(i.unicode?"u":"")+(x?"g":"y"),v=new c(x?"^(?:"+i.source+")":i,g),b=void 0===a?y:a>>>0;if(0===b)return[];if(0===s.length)return null===m(v,s)?[s]:[];var _=0,C=0,k=[];while(C<s.length){v.lastIndex=x?0:C;var z,R=m(v,x?E(s,C):s);if(null===R||(z=w(d(v.lastIndex+(x?C:0)),s.length))===_)C=h(s,C,f);else{if($(k,E(s,_,C)),k.length===b)return k;for(var S=1;S<=R.length-1;S++)if($(k,R[S]),k.length===b)return k;C=_=z}}return $(k,E(s,_)),k}]}),!z,x)},"2cbc":function(e,t,r){},"44e7":function(e,t,r){var a=r("861d"),i=r("c6b6"),n=r("b622"),s=n("match");e.exports=function(e){var t;return a(e)&&(void 0!==(t=e[s])?!!t:"RegExp"==i(e))}},"4dae":function(e,t,r){var a=r("da84"),i=r("23cb"),n=r("07fa"),s=r("8418"),l=a.Array,o=Math.max;e.exports=function(e,t,r){for(var a=n(e),c=i(t,a),u=i(void 0===r?a:r,a),h=l(o(u-c,0)),d=0;c<u;c++,d++)s(h,d,e[c]);return h.length=d,h}},"7b02":function(e,t,r){},8418:function(e,t,r){"use strict";var a=r("a04b"),i=r("9bf2"),n=r("5c6c");e.exports=function(e,t,r){var s=a(t);s in e?i.f(e,s,n(0,r)):e[s]=r}},"8f93":function(e,t,r){},c5d9:function(e,t,r){"use strict";r("2cbc")},c932:function(e,t,r){"use strict";r("7b02")},e975:function(e,t,r){"use strict";r("8f93")}}]);
//# sourceMappingURL=chunk-554a9eaf.c0d7447a.js.map