(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-063e3de1"],{2695:function(t,e,a){},4364:function(t,e,a){"use strict";a("2695")},d504:function(t,e,a){"use strict";a.r(e);var s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{attrs:{id:"index"}},[a("header",{staticClass:"header"},[t._m(0),a("div",{staticClass:"head-top"},[a("el-breadcrumb",{attrs:{"separator-class":"el-icon-arrow-right"}},[a("el-breadcrumb-item",[t._v(t._s(t.$route.meta[0]))]),"Monitor"==t.$route.name?a("el-breadcrumb-item",{staticClass:"select-wrapper"},[a("div",{staticClass:"select-box"},[a("div",{staticClass:"label"},[t._v("主机id")]),a("div",{staticClass:"input"},[a("el-select",{staticClass:"custom-select",attrs:{placeholder:"请选择",size:"small"},on:{change:t.selectChange},model:{value:t.value,callback:function(e){t.value=e},expression:"value"}},t._l(t.options,(function(t){return a("el-option",{key:t.ID,attrs:{label:t.label,value:t.value}})})),1)],1)])]):t._e()],1),a("div",{staticClass:"head-top-line"})],1)]),a("div",{staticClass:"container"},[a("div",{staticClass:"aside"},[a("el-menu",{staticClass:"menu-left",staticStyle:{"min-height":"100%"},attrs:{"background-color":"#324057","text-color":"#ffffff","active-text-color":"#ffd04b","default-active":t.activateIndex,router:""}},[a("el-menu-item",{attrs:{index:"/monitor"}},[a("i",{staticClass:"el-icon-odometer"}),t._v("主机监控")]),a("el-menu-item",{attrs:{index:"/alarm"}},[a("i",{staticClass:"el-icon-alarm-clock"}),t._v("告警配置")]),a("el-menu-item",{attrs:{index:"/agent"}},[a("i",{staticClass:"el-icon-setting"}),t._v("agent管理")])],1)],1),a("div",{staticClass:"main"},[a("keep-alive",[a("router-view")],1)],1)])])},i=[function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"title"},[a("div",{staticClass:"icon iconfont icon-monitorjiankong"}),a("div",{staticClass:"name"},[t._v("监控平台")])])}],l=a("fd03"),n={data:function(){return{activateIndex:this.$route.meta[1],options:[],value:""}},mounted:function(){var t=this;Object(l["a"])().then((function(e){if(200==e.data.CODE){for(var a=e.data.DATA,s=0;s<a.length;s++)t.options.push({label:a[s].HostID,value:a[s].HostID,id:a[s].ID});t.$store.commit("getHost",a),t.value=t.options[0].value,sessionStorage.setItem("HostID",t.value)}}))},methods:{selectChange:function(t){this.$store.commit("selectChange",t)}}},o=n,c=(a("4364"),a("2877")),r=Object(c["a"])(o,s,i,!1,null,"1c02df47",null);e["default"]=r.exports}}]);
//# sourceMappingURL=chunk-063e3de1.13d2d70e.js.map