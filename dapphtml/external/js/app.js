!function(e){function t(t){for(var c,n,d=t[0],r=t[1],o=t[2],m=0,b=[];m<d.length;m++)n=d[m],Object.prototype.hasOwnProperty.call(i,n)&&i[n]&&b.push(i[n][0]),i[n]=0;for(c in r)Object.prototype.hasOwnProperty.call(r,c)&&(e[c]=r[c]);for(l&&l(t);b.length;)b.shift()();return s.push.apply(s,o||[]),a()}function a(){for(var e,t=0;t<s.length;t++){for(var a=s[t],c=!0,d=1;d<a.length;d++){var r=a[d];0!==i[r]&&(c=!1)}c&&(s.splice(t--,1),e=n(n.s=a[0]))}return e}var c={},i={0:0},s=[];function n(t){if(c[t])return c[t].exports;var a=c[t]={i:t,l:!1,exports:{}};return e[t].call(a.exports,a,a.exports,n),a.l=!0,a.exports}n.m=e,n.c=c,n.d=function(e,t,a){n.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:a})},n.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},n.t=function(e,t){if(1&t&&(e=n(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var a=Object.create(null);if(n.r(a),Object.defineProperty(a,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var c in e)n.d(a,c,function(t){return e[t]}.bind(null,c));return a},n.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return n.d(t,"a",t),t},n.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},n.p="/";var d=window.webpackJsonp=window.webpackJsonp||[],r=d.push.bind(d);d.push=t,d=d.slice();for(var o=0;o<d.length;o++)t(d[o]);var l=r;s.push([222,1]),a()}({222:function(e,t,a){e.exports=a(344)},330:function(e,t,a){(t=a(40)(!1)).push([e.i,'body {\n  background: white;\n  font-family: "Source Sans Pro", sans-serif;\n  line-height: 1.5;\n  margin: 0px;\n}\n\n.root {\n  font: "Source Sans Pro", sans-serif;\n  --font-monospace: "Source Sans Pro", sans-serif;\n  --font-monospace-size: 0.8rem !important;\n  -webkit-font-smoothing: antialiased;\n}\n\n.privacy-banner {\n  margin-bottom: 20px;\n}\n',""]),e.exports=t},332:function(e,t,a){var c={"./af":82,"./af.js":82,"./ar":83,"./ar-dz":84,"./ar-dz.js":84,"./ar-kw":85,"./ar-kw.js":85,"./ar-ly":86,"./ar-ly.js":86,"./ar-ma":87,"./ar-ma.js":87,"./ar-sa":88,"./ar-sa.js":88,"./ar-tn":89,"./ar-tn.js":89,"./ar.js":83,"./az":90,"./az.js":90,"./be":91,"./be.js":91,"./bg":92,"./bg.js":92,"./bm":93,"./bm.js":93,"./bn":94,"./bn.js":94,"./bo":95,"./bo.js":95,"./br":96,"./br.js":96,"./bs":97,"./bs.js":97,"./ca":98,"./ca.js":98,"./cs":99,"./cs.js":99,"./cv":100,"./cv.js":100,"./cy":101,"./cy.js":101,"./da":102,"./da.js":102,"./de":103,"./de-at":104,"./de-at.js":104,"./de-ch":105,"./de-ch.js":105,"./de.js":103,"./dv":106,"./dv.js":106,"./el":107,"./el.js":107,"./en-SG":108,"./en-SG.js":108,"./en-au":109,"./en-au.js":109,"./en-ca":110,"./en-ca.js":110,"./en-gb":111,"./en-gb.js":111,"./en-ie":112,"./en-ie.js":112,"./en-il":113,"./en-il.js":113,"./en-nz":114,"./en-nz.js":114,"./eo":115,"./eo.js":115,"./es":116,"./es-do":117,"./es-do.js":117,"./es-us":118,"./es-us.js":118,"./es.js":116,"./et":119,"./et.js":119,"./eu":120,"./eu.js":120,"./fa":121,"./fa.js":121,"./fi":122,"./fi.js":122,"./fo":123,"./fo.js":123,"./fr":124,"./fr-ca":125,"./fr-ca.js":125,"./fr-ch":126,"./fr-ch.js":126,"./fr.js":124,"./fy":127,"./fy.js":127,"./ga":128,"./ga.js":128,"./gd":129,"./gd.js":129,"./gl":130,"./gl.js":130,"./gom-latn":131,"./gom-latn.js":131,"./gu":132,"./gu.js":132,"./he":133,"./he.js":133,"./hi":134,"./hi.js":134,"./hr":135,"./hr.js":135,"./hu":136,"./hu.js":136,"./hy-am":137,"./hy-am.js":137,"./id":138,"./id.js":138,"./is":139,"./is.js":139,"./it":140,"./it-ch":141,"./it-ch.js":141,"./it.js":140,"./ja":142,"./ja.js":142,"./jv":143,"./jv.js":143,"./ka":144,"./ka.js":144,"./kk":145,"./kk.js":145,"./km":146,"./km.js":146,"./kn":147,"./kn.js":147,"./ko":148,"./ko.js":148,"./ku":149,"./ku.js":149,"./ky":150,"./ky.js":150,"./lb":151,"./lb.js":151,"./lo":152,"./lo.js":152,"./lt":153,"./lt.js":153,"./lv":154,"./lv.js":154,"./me":155,"./me.js":155,"./mi":156,"./mi.js":156,"./mk":157,"./mk.js":157,"./ml":158,"./ml.js":158,"./mn":159,"./mn.js":159,"./mr":160,"./mr.js":160,"./ms":161,"./ms-my":162,"./ms-my.js":162,"./ms.js":161,"./mt":163,"./mt.js":163,"./my":164,"./my.js":164,"./nb":165,"./nb.js":165,"./ne":166,"./ne.js":166,"./nl":167,"./nl-be":168,"./nl-be.js":168,"./nl.js":167,"./nn":169,"./nn.js":169,"./pa-in":170,"./pa-in.js":170,"./pl":171,"./pl.js":171,"./pt":172,"./pt-br":173,"./pt-br.js":173,"./pt.js":172,"./ro":174,"./ro.js":174,"./ru":175,"./ru.js":175,"./sd":176,"./sd.js":176,"./se":177,"./se.js":177,"./si":178,"./si.js":178,"./sk":179,"./sk.js":179,"./sl":180,"./sl.js":180,"./sq":181,"./sq.js":181,"./sr":182,"./sr-cyrl":183,"./sr-cyrl.js":183,"./sr.js":182,"./ss":184,"./ss.js":184,"./sv":185,"./sv.js":185,"./sw":186,"./sw.js":186,"./ta":187,"./ta.js":187,"./te":188,"./te.js":188,"./tet":189,"./tet.js":189,"./tg":190,"./tg.js":190,"./th":191,"./th.js":191,"./tl-ph":192,"./tl-ph.js":192,"./tlh":193,"./tlh.js":193,"./tr":194,"./tr.js":194,"./tzl":195,"./tzl.js":195,"./tzm":196,"./tzm-latn":197,"./tzm-latn.js":197,"./tzm.js":196,"./ug-cn":198,"./ug-cn.js":198,"./uk":199,"./uk.js":199,"./ur":200,"./ur.js":200,"./uz":201,"./uz-latn":202,"./uz-latn.js":202,"./uz.js":201,"./vi":203,"./vi.js":203,"./x-pseudo":204,"./x-pseudo.js":204,"./yo":205,"./yo.js":205,"./zh-cn":206,"./zh-cn.js":206,"./zh-hk":207,"./zh-hk.js":207,"./zh-tw":208,"./zh-tw.js":208};function i(e){var t=s(e);return a(t)}function s(e){if(!a.o(c,e)){var t=new Error("Cannot find module '"+e+"'");throw t.code="MODULE_NOT_FOUND",t}return c[e]}i.keys=function(){return Object.keys(c)},i.resolve=s,e.exports=i,i.id=332},338:function(e,t,a){var c=a(39),i=a(339);"string"==typeof(i=i.__esModule?i.default:i)&&(i=[[e.i,i,""]]);var s={insert:"head",singleton:!1},n=(c(i,s),i.locals?i.locals:{});e.exports=n},339:function(e,t,a){(t=a(40)(!1)).push([e.i,'body {\n  margin: 0;\n  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen",\n    "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",\n    sans-serif;\n  -webkit-font-smoothing: antialiased;\n  -moz-osx-font-smoothing: grayscale;\n}\n\ncode {\n  font-family: source-code-pro, Menlo, Monaco, Consolas, "Courier New",\n    monospace;\n}',""]),e.exports=t},344:function(e,t,a){"use strict";a.r(t);var c=a(3),i=a.n(c),s=a(25),n=a.n(s),d=a(346),r=a(10),o=a(217),l=a(5),m=(a(14),a(81),a(0));var b={name:"11tluwr",styles:'position:absolute;top:0;left:0;width:100%;height:100%;opacity:0.6;display:none;background-image:url("/static/images/watermark.svg");background-repeat:repeat;@media print{display:block;}'};const j=()=>Object(m.d)("div",{css:b});var O=a.p+"aaab203210feb25570cdd124a79370ff.svg",u=a.p+"c32209aa68332e2f221b5516c56990bf.png",g=a.p+"43d70e329eccff324df7aac26017d4e3.png";const p=Object(r.a)("div",{target:"ev00jpv0"})("max-width:297mm;margin:0 auto;position:relative;background-image:",e=>e.certificateBg,";background-position:center;background-size:cover;border:10px solid #324353;.logo-oc{width:320px;@media (min-width:1024px){width:600px;}}.logo-gt{width:200px;@media (min-width:1024px){width:300px;}}.signature{width:160px;@media (min-width:1024px){width:230px;}}.text-sm{font-size:12px;@media (min-width:1024px){font-size:21px;}}.text-md{font-size:16px;@media (min-width:1024px){font-size:26px;}}.text-lg{font-size:24px;@media (min-width:1024px){font-size:34px;}}.spacer{margin:24px;@media (min-width:1024px){margin:48px;}}footer{max-width:960px;margin:0 auto;padding:64px 24px;@media (min-width:1024px){padding:96px;}}@media print{/* @page{size:A4 landscape;}*/ .logo-oc{width:320px;}.logo-gt{width:200px;}.signature{width:160px;}.text-sm{font-size:12px;}.text-md{font-size:16px;}.text-lg{font-size:24px;}.spacer{margin:24px;}footer{padding:64px 24px;}}");var f=a(42);const N=e=>{if(!e)return null;const t=new Date(e);return Object(f.tz)(t,"Asia/Singapore").format("MMM YYYY").toUpperCase()};var v=a(345),h=a.p+"cbfc062552a3ea7c7909ff47dc07cd80.png",I=a(27),T=a(28);const y={name:"bz64z4",styles:"@media print{display:none;}background-color:whitesmoke;display:flex;justify-content:space-between;padding:20px;.icon{cursor:pointer;}"},w=({onToggleEditable:e,className:t=""})=>Object(m.d)("div",{css:y,className:t},Object(m.d)("div",{className:"text-container"},Object(m.d)("div",null,Object(m.d)("h4",null,"OpenCerts Privacy Filter Enabled")),Object(m.d)("div",null,"Edit this certificate by removing sensitive information by clicking on the edit button. Downloaded certificate remains valid!")),Object(m.d)("h5",{className:"icon",onClick:e},Object(m.d)(I.a,{icon:T.a,title:"toggle certificate obfuscation"}))),x=[{id:"certificate",label:"Certificate",template:({document:e})=>Object(m.d)(p,{certificateBg:`url('${g}')`,className:"p-4"},Object(m.d)(j,null),Object(m.d)("main",{className:"text-center"},Object(m.d)("div",{className:"spacer"},Object(m.d)("img",{src:O,className:"img-fluid logo-oc",alt:"OpenCerts Logo"})),Object(m.d)("div",{className:"spacer text-md"},Object(m.d)("i",null,"This is to certify that")),Object(m.d)("div",{className:"spacer text-lg"},Object(m.d)("b",null,e.recipient.name)),Object(m.d)("div",{className:"spacer text-md"},Object(m.d)("i",null,"has successfully completed the")),Object(m.d)("div",{className:"spacer text-lg"},"OpenCerts Demo"),Object(m.d)("div",{className:"spacer text-md"},Object(m.d)("i",null,"certification through training administered by")),Object(m.d)("img",{className:"img-fluid logo-gt",src:u,alt:"Govtech Logo"})),Object(m.d)("footer",null,Object(m.d)("div",{className:"row align-items-center"},Object(m.d)("div",{className:"col"},Object(m.d)("div",{className:"text-center text-sm"},Object(m.d)("img",{className:"img-fluid signature",src:Object(l.get)(e,"additionalData.certSignatories[0].signature"),alt:"Signature"}),Object(m.d)("hr",{style:{backgroundColor:"#333"}}),Object(m.d)("div",null,Object(m.d)("b",null,Object(l.get)(e,"additionalData.certSignatories[0].name")),Object(m.d)("br",null),Object(l.get)(e,"additionalData.certSignatories[0].position"),","," ",Object(l.get)(e,"additionalData.certSignatories[0].organisation")))),Object(m.d)("div",{className:"col"}),Object(m.d)("div",{className:"col"},Object(m.d)("div",{className:"text-sm text-right"},"Dated ",Object(o.format)(e.issuedOn,"DD/MM/YYYY"))))))},{id:"transcript",label:"Transcript",template:({document:e,handleObfuscation:t})=>{const[a,s]=Object(c.useState)(!1),n=Object(l.get)(e,"name"),d=Object(l.get)(e,"id"),r=Object(l.get)(e,"issuedOn"),o=Object(l.get)(e,"admissionDate"),b=Object(l.get)(e,"graduationDate"),O=Object(l.get)(e,"recipient.name"),g=Object(l.get)(e,"recipient.nric"),p=Object(l.get)(e,"recipient.course"),f=Object(l.get)(e,"additionalData.studentId"),I=e.transcript||[],T=I.map((e,c)=>Object(m.d)("tr",{key:c},Object(m.d)("td",null,Object(m.d)(v.a,{editable:a,value:e.courseCode,onObfuscationRequested:()=>t(`transcript[${c}].courseCode`)})),Object(m.d)("td",null,Object(m.d)(v.a,{editable:a,value:e.name,onObfuscationRequested:()=>t(`transcript[${c}].name`)})),Object(m.d)("td",null,Object(m.d)(v.a,{editable:a,value:e.grade,onObfuscationRequested:()=>t(`transcript[${c}].grade`)})),Object(m.d)("td",null,Object(m.d)(v.a,{editable:a,value:e.courseCredit,onObfuscationRequested:()=>t(`transcript[${c}].courseCredit`)})),Object(m.d)("td",null,Object(m.d)(v.a,{editable:a,value:e.semester,onObfuscationRequested:()=>t(`transcript[${c}].semester`)}))));return Object(m.d)(i.a.Fragment,null,Object(m.d)(j,null),Object(m.d)("div",{className:"container"},Object(m.d)(w,{onToggleEditable:()=>s(!a),className:"privacy-banner"}),Object(m.d)("div",{className:"p-2 container",style:{backgroundImage:`url('${h}')`,backgroundRepeat:"repeat"}},Object(m.d)("div",{className:"row root cert-title",style:{paddingLeft:"3%"}},Object(m.d)("b",null,n)),Object(m.d)("div",{className:"row transcript",style:{paddingTop:"3%",paddingLeft:"2%"}},Object(m.d)("div",{className:"col"},Object(m.d)("div",{className:"row"},Object(m.d)("div",{className:"col"},"NAME"),Object(m.d)("div",{className:"col"},":  ",O)),Object(m.d)("div",{className:"row"},Object(m.d)("div",{className:"col"},"COURSE"),Object(m.d)("div",{className:"col"},":  ",p)),Object(m.d)("div",{className:"row"},Object(m.d)("div",{className:"col"},"NRIC/FIN"),Object(m.d)("div",{className:"col"},":  ",g)),Object(m.d)("div",{className:"row"},Object(m.d)("div",{className:"col"},"STUDENT ID"),Object(m.d)("div",{className:"col"},":  ",f))),Object(m.d)("div",{className:"col"},Object(m.d)("div",{className:"row"},Object(m.d)("div",{className:"col"},"DOCUMENT ID"),Object(m.d)("div",{className:"col"},":  ",d)),Object(m.d)("div",{className:"row"},Object(m.d)("div",{className:"col"},"DATE OF ISSUANCE"),Object(m.d)("div",{className:"col"},":  ",N(r))),Object(m.d)("div",{className:"row"},Object(m.d)("div",{className:"col"},"DATE OF ADMISSION"),Object(m.d)("div",{className:"col"},":  ",N(o))),Object(m.d)("div",{className:"row"},Object(m.d)("div",{className:"col"},"DATE OF GRADUATION"),Object(m.d)("div",{className:"col"},":  ",N(b))))),I!==[]&&Object(m.d)("div",{className:"row mb-4",style:{paddingLeft:"3%",paddingTop:"5%"}},Object(m.d)("div",{className:"root cert-title"},Object(m.d)("b",null,"Transcript")),Object(m.d)("table",{className:"w-100 transcript"},Object(m.d)("tbody",null,Object(m.d)("tr",null,Object(m.d)("th",null,"Course Code"),Object(m.d)("th",null,"Name"),Object(m.d)("th",null,"Grade"),Object(m.d)("th",null,"Units"),Object(m.d)("th",null,"Semester")),T))),Object(m.d)("div",{className:"row"},Object(m.d)("div",{className:"col"},Object(m.d)("img",{className:"w-100",style:{paddingTop:"40%",paddingLeft:"3%",width:"100%",height:"auto"},src:u,alt:"Govtech Logo"})),Object(m.d)("div",{className:"col"}),Object(m.d)("div",{className:"col text-center",style:{paddingTop:"5%",paddingRight:"5%",width:"100%",height:"auto"}},Object(m.d)("img",{className:"w-100",src:Object(l.get)(e,"additionalData.certSignatories[0].signature")}),Object(m.d)("hr",{className:"m-1"}),Object(m.d)("div",{className:"transcript"},Object(m.d)("b",null,Object(l.get)(e,"additionalData.certSignatories[0].name")),Object(m.d)("br",null),Object(l.get)(e,"additionalData.certSignatories[0].position"),","," ",Object(l.get)(e,"additionalData.certSignatories[0].organisation")))))))}},{id:"media",label:"Media",template:()=>Object(m.d)("div",{className:"embed-responsive embed-responsive-16by9"},Object(m.d)("iframe",{id:"youtube-vid",className:"embed-responsive-item",src:"https://www.youtube.com/embed/oskddwGpwUw?autoplay=1",allowFullScreen:!0}))}];var S=a.p+"24ff1f84878632992098d1ed6cf7c04b.png";const D=Object(r.a)("div",{target:"e1bocjee0"})("background:rgb(255,255,255);width:","21cm",";min-height:","29.7cm",";margin:1cm auto;border:1px #d3d3d3 solid;border-radius:5px;box-shadow:0 0 5px rgba(0,0,0,0.1);@media print{margin:3cm auto;page-break-after:always;}"),C=Object(r.a)("div",{target:"e1bocjee1"})({name:"a9n7s9",styles:"margin-left:auto;margin-right:auto;"});var M={name:"yvjx77",styles:"@page{size:A4;margin:0;}"};const E=({children:e})=>Object(m.d)(C,null,Object(m.d)(j,null),Object(m.d)(m.a,{styles:M}),e);const A=Object(r.a)("div",{target:"ensrldt0"})("height:","0.5cm",";background:rgb(164,139,97);background:radial-gradient(circle,rgba(164,139,97,1) 0%,rgba(116,100,75,1) 100%);color:white;overflow:none;display:flex;justify-content:center;align-items:center;font-size:0.5rem;white-space:nowrap;@media print{white-space:normal;}"),k=Object(r.a)(D,{target:"ensrldt1"})({name:"11q98y6",styles:"background:linear-gradient(176deg,rgba(255,255,255,1) 0%,rgba(240,252,255,1) 100%);"});var R=a.p+"832b50d45890fe3ca65cc89a371bbf27.jpg";var F={name:"g2rzi",styles:"padding-bottom:1cm;padding-top:1cm;"},U={name:"v7v99c",styles:"width:100px;"},Y={name:"3bfp34",styles:"background-color:#0072c6;color:white;height:1cm;margin-left:1cm;margin-right:1cm;padding-left:1cm;padding-right:1cm;display:flex;align-items:center;"};const z=[{id:"testimonial",label:"Testimonial",template:({document:e})=>Object(m.d)(E,null,Object(m.d)(k,null,Object(m.d)("div",{css:Object(m.c)("background-color:white;height:","8cm",";border-bottom:1px solid #999c97;img{margin-top:1cm;height:4cm;}.title{font-size:1.2rem;font-weight:bold;}.sub-title{font-size:1.1rem;font-weight:bold;}}")},Object(m.d)("div",{className:"text-center"},Object(m.d)("img",{src:S})),Object(m.d)("div",{className:"text-center title"},"Ministry of Education"),Object(m.d)("div",{className:"text-center sub-title"},"Republic of Singapore")),Object(m.d)("div",{className:"text-center",css:Object(m.c)("height:","calc(29.7cm - 8cm - 1cm)",";.certificate{font-size:2rem;font-weight:bold;color:#80764f;margin-top:1.5cm;margin-bottom:1.5cm;}.recipient{margin-top:1.5cm;}.institute{margin-bottom:0.75cm;}.year-of-graduation{margin-bottom:1.5cm;}.information{margin-bottom:1.5cm;}.seal img{width:4cm;}.signature img{width:6cm;border-bottom:2px solid #b6b7b2;margin-bottom:0.25cm;}.signature-container{}")},Object(m.d)("div",{className:"certificate"},"School Graduate Certificate"),Object(m.d)("div",{className:"text-uppercase"},"The certificate is awarded to"),Object(m.d)("div",{className:"recipient"},e.recipient.name),Object(m.d)("div",null,"Identification No: ",e.recipient.studentId),Object(m.d)("div",{className:"institute"},e.testimonial.instituteName),Object(m.d)("div",{className:"year-of-graduation"},"Year of Graduation: ",e.testimonial.achievementDate.split("-")[0]),Object(m.d)("div",{className:"information"},Object(m.d)("div",null,"This School Graduation Certificates provides a hlistic qualitative"),Object(m.d)("div",null,"assessment of the student's personal qualities, academic and co-curricular"),Object(m.d)("div",null,"development upon completion of the student's pre-university education.")),Object(m.d)("div",{className:"flex-row d-flex justify-content-around"},Object(m.d)("div",{className:"seal w-50"},Object(m.d)("img",{src:e.referee.seal})),Object(m.d)("div",{className:"signature w-50"},Object(m.d)("div",{className:"signature-container"},Object(m.d)("img",{src:e.referee.signature})),Object(m.d)("div",null,e.referee.designation)))),Object(m.d)("div",{css:Object(m.c)("height:","1cm",";background:rgb(164,139,97);background:radial-gradient(circle,rgba(164,139,97,1) 0%,rgba(116,100,75,1) 100%);")})),Object(m.d)(k,null,Object(m.d)("div",{className:"header d-flex justify-content-between align-items-end",css:Object(m.c)("padding:0 2cm 0.5cm;background-color:white;height:","4.5cm",";.type{font-size:1.2rem;}.institute{font-size:0.9rem;}.title{font-size:1.1rem;font-weight:bold;}.sub-title{font-size:0.9rem;font-weight:bold;}img{width:2.5cm;}")},Object(m.d)("div",null,Object(m.d)("div",{className:"type"},"Testimonial"),Object(m.d)("div",{className:"institute"},e.testimonial.instituteName)),Object(m.d)("div",{className:"text-center"},Object(m.d)("div",null,Object(m.d)("img",{src:S})),Object(m.d)("div",{className:"title"},"Ministry of Education"),Object(m.d)("div",{className:"sub-title"},"Republic of Singapore"))),Object(m.d)(A,null,"MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION"),Object(m.d)("div",{css:Object(m.c)("height:","calc(29.7cm - 4.5cm - 2cm - 2 * 0.5cm)",";padding:1cm 2cm;"),className:"text-justify",dangerouslySetInnerHTML:{__html:e.content}}),Object(m.d)(A,null,"MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION MINISTRY OF EDUCATION"),Object(m.d)("div",{className:"footer",css:Object(m.c)("background-color:white;height:","2cm",";")})))}],_=[{id:"testimonial",label:"Testimonial",template:({document:e})=>Object(m.d)(E,null,Object(m.d)(D,null,Object(m.d)("div",{css:Object(m.c)("background:url(",R,");height:","5cm",";background-size:","21cm"," ","5cm",";")}),Object(m.d)("div",{css:Object(m.c)("padding:1cm 2cm;height:","calc(29.7cm - 5cm - 2cm)",";")},Object(m.d)("p",{css:F},new Date(e.issuedOn).toLocaleDateString(void 0,{year:"numeric",month:"long",day:"numeric"})),Object(m.d)("div",{className:"text-justify",dangerouslySetInnerHTML:{__html:e.content}}),Object(m.d)("div",null,Object(m.d)("div",null,Object(m.d)("img",{css:U,src:e.referee.signature})),Object(m.d)("div",null,e.referee.name),Object(m.d)("div",null,e.referee.designation),Object(m.d)("div",null,"Mobile: ",e.referee.mobile),Object(m.d)("div",null,"Email: ",e.referee.email))),Object(m.d)("div",{css:Object(m.c)("height:","2cm",";")},Object(m.d)("div",{css:Y},e.referee.address))))}];var L=a.p+"e0ed8c80ffffe9324f7897bf53b30a62.svg",q=a.p+"d7815ac578223eb4c71c971b6e8422ad.png";const G={default:x,GOVTECH_DEMO:x,TESTIMONIALS:z,TESTIMONIALS2:_,CERTIFICATE_OF_AWARD:[{id:"certificate-of-award",label:"Certificate Of Award",template:({document:e})=>Object(m.d)(E,null,Object(m.d)(D,null,Object(m.d)("div",{css:Object(m.c)("background:linear-gradient(90deg,rgba(255,255,255,1) 0%,rgba(180,191,245,1) 100%);display:flex;align-items:center;justify-content:space-between;padding:0 2cm;height:","3.5cm",";border-bottom:1px solid #999c97;img{height:2.5cm;}")},Object(m.d)("img",{src:L}),Object(m.d)("img",{src:q})),Object(m.d)("div",{className:"text-center",css:Object(m.c)("height:","calc(29.7cm - 3.5cm - 2cm)",";padding:0 2cm;.certificate{font-size:2rem;font-weight:bold;color:#000000;margin:1.5cm 0;}.recipient{font-size:1.5rem;font-weight:bold;color:#000000;margin:0.5cm 0;}.achievement-area{margin-top:0.5cm;}.institute{margin-bottom:1.5cm;}.information{margin:1.5cm 0;}.signature{line-height:2rem;}.signature img{width:6cm;border-bottom:2px solid #b6b7b2;margin-bottom:0.5cm;}")},Object(m.d)("div",{className:"certificate text-uppercase"},e.name),Object(m.d)("div",null,"is presented to"),Object(m.d)("div",{className:"recipient text-uppercase"},e.recipient.name),Object(m.d)("div",null,"of"),Object(m.d)("div",{className:"achievement-area text-uppercase"},e.award.achievementArea,Object(m.d)("br",null),"in",Object(m.d)("br",null),e.award.instituteName),Object(m.d)("div",{className:"information"},"for demonstrating leadership qualities, service to community and schools,",Object(m.d)("br",null),"excellence in non-academic activities, and",Object(m.d)("br",null),"good conduct"),Object(m.d)("div",{className:"signature"},Object(m.d)("img",{src:e.signature.signature}),Object(m.d)("div",{className:"text-uppercase"},e.signature.name,Object(m.d)("br",null),e.signature.designation))),Object(m.d)("div",{className:"footer",css:Object(m.c)("height:","2cm",";padding:0 2cm;")},"ABC123456")))}]};a(338);n.a.render(Object(m.d)(d.a,{templateRegistry:G}),document.getElementById("root"))},81:function(e,t,a){var c=a(39),i=a(330);"string"==typeof(i=i.__esModule?i.default:i)&&(i=[[e.i,i,""]]);var s={insert:"head",singleton:!1},n=(c(i,s),i.locals?i.locals:{});e.exports=n}});
