(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-74939142"],{"00b4":function(e,t,r){"use strict";r("ac1f");var a=r("23e7"),i=r("da84"),l=r("c65b"),n=r("e330"),s=r("1626"),o=r("861d"),u=function(){var e=!1,t=/[ac]/;return t.exec=function(){return e=!0,/./.exec.apply(this,arguments)},!0===t.test("abc")&&e}(),c=i.Error,h=n(/./.test);a({target:"RegExp",proto:!0,forced:!u},{test:function(e){var t=this.exec;if(!s(t))return h(this,e);var r=l(t,this,e);if(null!==r&&!o(r))throw new c("RegExp exec method returned something other than an Object or null");return!!r}})},1017:function(e,t,r){"use strict";r.r(t);var a=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{attrs:{id:"setting"}},[e._m(0),r("div",{staticClass:"setting-wrapper-box"},[r("div",{staticClass:"section-title"},[e._v(" 配置告警规则 ")]),r("div",{staticClass:"section-content"},[r("alarm-setting")],1),r("div",{staticClass:"button-line"})]),r("div",{staticClass:"item-wrapper-box"},[r("div",{staticClass:"section-title"},[e._v(" 告警规则列表 ")]),r("div",{staticClass:"section-content"},[r("alarm-list")],1)])])},i=[function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"tips"},[r("i",{staticClass:"el-icon-warning-outline"}),r("span",[e._v("在此处可以自定义配置告警阈值，联系邮箱等")])])}],l=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{attrs:{id:"setting-wrapper"}},[r("div",{staticClass:"configure-wrapper"},[r("div",{staticClass:"label-group"},[e._v("cpu使用率阈值")]),r("div",{staticClass:"input-group"},[r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"提示（0-1）",size:"small"},model:{value:e.configure.cpu_threshold.warning,callback:function(t){e.$set(e.configure.cpu_threshold,"warning",t)},expression:"configure.cpu_threshold.warning"}}),r("div",{staticClass:"split"},[e._v(" ＜ ")]),r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"严重（0-1）",size:"small"},model:{value:e.configure.cpu_threshold.serious,callback:function(t){e.$set(e.configure.cpu_threshold,"serious",t)},expression:"configure.cpu_threshold.serious"}}),r("div",{staticClass:"split"},[e._v(" ＜ ")]),r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"致命（0-1）",size:"small"},model:{value:e.configure.cpu_threshold.mortal,callback:function(t){e.$set(e.configure.cpu_threshold,"mortal",t)},expression:"configure.cpu_threshold.mortal"}})],1)]),r("div",{staticClass:"configure-wrapper"},[r("div",{staticClass:"label-group"},[e._v("内存使用率阈值")]),r("div",{staticClass:"input-group"},[r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"提示（0-1）",size:"small"},model:{value:e.configure.memory_threshold.warning,callback:function(t){e.$set(e.configure.memory_threshold,"warning",t)},expression:"configure.memory_threshold.warning"}}),r("div",{staticClass:"split"},[e._v(" ＜ ")]),r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"严重（0-1）",size:"small"},model:{value:e.configure.memory_threshold.serious,callback:function(t){e.$set(e.configure.memory_threshold,"serious",t)},expression:"configure.memory_threshold.serious"}}),r("div",{staticClass:"split"},[e._v(" ＜ ")]),r("el-input",{staticClass:"threshold-group",staticStyle:{width:"100px"},attrs:{placeholder:"致命（0-1）",size:"small"},model:{value:e.configure.memory_threshold.mortal,callback:function(t){e.$set(e.configure.memory_threshold,"mortal",t)},expression:"configure.memory_threshold.mortal"}})],1)]),r("div",{staticClass:"configure-wrapper"},[r("div",{staticClass:"label-group"},[e._v("监控粒度")]),r("el-select",{staticClass:"input-group",staticStyle:{width:"150px"},attrs:{placeholder:"请选择监控间隔",size:"small"},model:{value:e.configure.time_granularity,callback:function(t){e.$set(e.configure,"time_granularity",t)},expression:"configure.time_granularity"}},e._l(e.granularity_options,(function(e){return r("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1),r("div",{staticClass:"configure-wrapper"},[r("div",{staticClass:"label-group"},[e._v("联系邮箱")]),r("el-input",{staticStyle:{width:"550px"},attrs:{placeholder:"请输入告警联系邮箱，可填写多个，用分号分隔,如XXX@163.com;XXX@qq.com",size:"small"},model:{value:e.configure.email,callback:function(t){e.$set(e.configure,"email",t)},expression:"configure.email"}})],1),r("div",{staticClass:"configure-wrapper"},[r("div",{staticClass:"label-group"}),r("el-button",{attrs:{type:"primary"},on:{click:e.addConfigure}},[e._v("添加配置")]),r("el-button",{on:{click:e.resetConfigure}},[e._v("重置")])],1)])},n=[];function s(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}var o=r("1da1"),u=(r("ac1f"),r("00b4"),r("5319"),r("1276"),r("96cf"),r("fd03")),c={data:function(){return{type:"",configure:{cpu_threshold:{warning:"",serious:"",mortal:""},memory_threshold:{warning:"",serious:"",mortal:""},email:"",time_granularity:""},email_arr:[],granularity_options:[{value:"1m",label:"1分钟"},{value:"2m",label:"2分钟"},{value:"5m",label:"5分钟"},{value:"10m",label:"10分钟"}]}},mounted:function(){},methods:{addConfigure:function(){var e=this;return Object(o["a"])(regeneratorRuntime.mark((function t(){var r,a,i;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(e.threshold_verify(e.configure.cpu_threshold||!e.threshold_verify(e.configure.memory_threshold))){t.next=4;break}return t.abrupt("return");case 4:if(!e.email_verify(e.configure.email)){t.next=10;break}return r={CpuNoteworthyThreshold:e.configure.cpu_threshold.warning,CpuSeriousThreshold:e.configure.cpu_threshold.serious,CpuDeadlyThreshold:e.configure.cpu_threshold.mortal,MemNoteworthyThreshold:e.configure.memory_threshold.warning,MemSeriousThreshold:e.configure.memory_threshold.serious},s(r,"MemSeriousThreshold",e.configure.memory_threshold.mortal),s(r,"Granularity",e.configure.time_granularity),s(r,"ContactEmail",e.email_arr),a=r,t.next=8,Object(u["f"])(a);case 8:i=t.sent,200==i.data.CODE?e.$message.success("成功添加配置"):e.$message.error(i.data.msg);case 10:case"end":return t.stop()}}),t)})))()},email_verify:function(e){if(0==e.length)return this.$message.warning("联系邮箱不能为空"),!1;var t=/\uff1b/;t.test(e)&&(e=e.replace(/\uff1b/g,";"));for(var r=e.split(";"),a=/^([a-zA-Z0-9])[a-zA-Z0-9]+@[a-zA-Z0-9]+\.([a-zA-Z0-9]{2,5})$/,i=0;i<r.length;i++){if(!a.test(r[i]))return this.$message.warning("存在不合法邮箱格式"),!1;console.log("true")}return this.email_arr=r,!0},threshold_verify:function(e){var t=e.warning,r=e.serious,a=e.mortal;return""==t||""==r||""==a?(this.$message.warning("阈值设置不能为空"),!1):t<=0||t>=1||r<=0||r>=1||a<=0||a>=1?(this.$message.warning("阈值数值请设置在0-1之间"),!1):!(t>=r||t>=a||r>=a)||(this.$message.warning("阈值区间设置有误"),!1)},resetConfigure:function(){this.configure={cpu_threshold:{warning:"",serious:"",mortal:""},memory_threshold:{warning:"",serious:"",mortal:""},time_granularity:"",email:""}}}},h=c,d=(r("4b00"),r("2877")),p=Object(d["a"])(h,l,n,!1,null,"204d0b99",null),m=p.exports,f=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{attrs:{id:"alarm-list"}},[r("el-table",{attrs:{data:e.tableData,"header-cell-style":{background:"#eef1f6",color:"#606266",border:"1px solid #DFE1E4"},"span-method":e.arraySpan}},[r("el-table-column",{attrs:{prop:"id",label:"规则id"}}),r("el-table-column",{attrs:{prop:"metrics",label:"监控指标"}}),r("el-table-column",{attrs:{label:"阈值","header-align":"center"}},[r("el-table-column",{attrs:{prop:"threshold.warning",label:"提示"}}),r("el-table-column",{attrs:{prop:"threshold.serious",label:"严重"}}),r("el-table-column",{attrs:{prop:"threshold.mortal",label:"致命"}})],1),r("el-table-column",{attrs:{prop:"granularity",label:"粒度"}}),r("el-table-column",{attrs:{prop:"email",label:"联系邮箱"}}),r("el-table-column",{attrs:{prop:"create_time",label:"创建时间"}})],1)],1)},g=[],v={data:function(){return{tableData:[]}},mounted:function(){this.initData()},methods:{initData:function(){var e=this;return Object(o["a"])(regeneratorRuntime.mark((function t(){var r,a,i,l,n,s,o,c,h,d,p,m,f,g;return regeneratorRuntime.wrap((function(t){while(1)switch(t.prev=t.next){case 0:return t.next=2,Object(u["e"])();case 2:if(r=t.sent,a=r.data,200==a.CODE){t.next=7;break}return console.log(a.msg),t.abrupt("return",!1);case 7:for(i=0;i<a.DATA.length;i++)l=a.DATA[i],n=l.ID,s=l.CpuNoteworthyThreshold,o=l.CpuSeriousThreshold,c=l.CpuDeadlyThreshold,h=l.MemNoteworthyThreshold,d=l.MemSeriousThreshold,p=l.MemDeadlyThreshold,m=l.ContactEmail,f=l.Granularity,g=l.CreatedAt,e.tableData.push({id:n,metrics:"cpu使用率",threshold:{warning:s,serious:o,mortal:c},granularity:f,email:m,create_time:e.timefilter(g)}),e.tableData.push({id:n,metrics:"内存使用率",threshold:{warning:h,serious:d,mortal:p},granularity:f,email:m,create_time:e.timefilter(g)});case 8:case"end":return t.stop()}}),t)})))()},timefilter:function(e){var t=/T|\.|(?<=\.).*/g;return e.replace(t," ")},arraySpan:function(e){e.row,e.column;var t=e.rowIndex,r=e.columnIndex;if(0==r||5==r||6==r||7==r)return t%2===0?[2,1]:[0,0]}}},b=v,_=(r("6a1a"),Object(d["a"])(b,f,g,!1,null,"0b7e1510",null)),y=_.exports,w={components:{AlarmSetting:m,AlarmList:y},data:function(){return{basePath:"",activateName:this.$route.query.type}},mounted:function(){},methods:{}},x=w,C=(r("e975"),Object(d["a"])(x,a,i,!1,null,"09385f01",null));t["default"]=C.exports},1276:function(e,t,r){"use strict";var a=r("2ba4"),i=r("c65b"),l=r("e330"),n=r("d784"),s=r("44e7"),o=r("825a"),u=r("1d80"),c=r("4840"),h=r("8aa5"),d=r("50c4"),p=r("577e"),m=r("dc4a"),f=r("4dae"),g=r("14c3"),v=r("9263"),b=r("9f7f"),_=r("d039"),y=b.UNSUPPORTED_Y,w=4294967295,x=Math.min,C=[].push,k=l(/./.exec),$=l(C),S=l("".slice),D=!_((function(){var e=/(?:)/,t=e.exec;e.exec=function(){return t.apply(this,arguments)};var r="ab".split(e);return 2!==r.length||"a"!==r[0]||"b"!==r[1]}));n("split",(function(e,t,r){var l;return l="c"=="abbc".split(/(b)*/)[1]||4!="test".split(/(?:)/,-1).length||2!="ab".split(/(?:ab)*/).length||4!=".".split(/(.?)(.?)/).length||".".split(/()()/).length>1||"".split(/.?/).length?function(e,r){var l=p(u(this)),n=void 0===r?w:r>>>0;if(0===n)return[];if(void 0===e)return[l];if(!s(e))return i(t,l,e,n);var o,c,h,d=[],m=(e.ignoreCase?"i":"")+(e.multiline?"m":"")+(e.unicode?"u":"")+(e.sticky?"y":""),g=0,b=new RegExp(e.source,m+"g");while(o=i(v,b,l)){if(c=b.lastIndex,c>g&&($(d,S(l,g,o.index)),o.length>1&&o.index<l.length&&a(C,d,f(o,1)),h=o[0].length,g=c,d.length>=n))break;b.lastIndex===o.index&&b.lastIndex++}return g===l.length?!h&&k(b,"")||$(d,""):$(d,S(l,g)),d.length>n?f(d,0,n):d}:"0".split(void 0,0).length?function(e,r){return void 0===e&&0===r?[]:i(t,this,e,r)}:t,[function(t,r){var a=u(this),n=void 0==t?void 0:m(t,e);return n?i(n,t,a,r):i(l,p(a),t,r)},function(e,a){var i=o(this),n=p(e),s=r(l,i,n,a,l!==t);if(s.done)return s.value;var u=c(i,RegExp),m=i.unicode,f=(i.ignoreCase?"i":"")+(i.multiline?"m":"")+(i.unicode?"u":"")+(y?"g":"y"),v=new u(y?"^(?:"+i.source+")":i,f),b=void 0===a?w:a>>>0;if(0===b)return[];if(0===n.length)return null===g(v,n)?[n]:[];var _=0,C=0,k=[];while(C<n.length){v.lastIndex=y?0:C;var D,E=g(v,y?S(n,C):n);if(null===E||(D=x(d(v.lastIndex+(y?C:0)),n.length))===_)C=h(n,C,m);else{if($(k,S(n,_,C)),k.length===b)return k;for(var T=1;T<=E.length-1;T++)if($(k,E[T]),k.length===b)return k;C=_=D}}return $(k,S(n,_)),k}]}),!D,y)},"4b00":function(e,t,r){"use strict";r("53c9")},"4dae":function(e,t,r){var a=r("da84"),i=r("23cb"),l=r("07fa"),n=r("8418"),s=a.Array,o=Math.max;e.exports=function(e,t,r){for(var a=l(e),u=i(t,a),c=i(void 0===r?a:r,a),h=s(o(c-u,0)),d=0;u<c;u++,d++)n(h,d,e[u]);return h.length=d,h}},"53c9":function(e,t,r){},"6a1a":function(e,t,r){"use strict";r("b978")},8418:function(e,t,r){"use strict";var a=r("a04b"),i=r("9bf2"),l=r("5c6c");e.exports=function(e,t,r){var n=a(t);n in e?i.f(e,n,l(0,r)):e[n]=r}},"8f93":function(e,t,r){},b978:function(e,t,r){},e975:function(e,t,r){"use strict";r("8f93")}}]);
//# sourceMappingURL=chunk-74939142.65759730.js.map