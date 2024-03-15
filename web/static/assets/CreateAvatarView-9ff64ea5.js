import{_ as M,u as x,c as D,a as v,r as N,b as p,t as k,o as C,d as c,e as a,f as s,w as L,g as t,h as F,i,F as h,E as b,j as P,k as q,l as B,m as E,p as T,n as U}from"./index-7a4cb04d.js";import{u as R}from"./llm.store-c02b0d93.js";import{u as $}from"./avatar.store-7fa534a5.js";const n=d=>(T("data-v-4b8bcc89"),d=d(),U(),d),j={class:"container-fluid p-0"},z=n(()=>a("h1",{class:"h3 mb-3"},"Create your AI Avatar",-1)),G={class:"row"},H={class:"col-12"},J={class:"card"},K={class:"card-body"},O={class:"row mb-5"},Q={class:"col-6"},W={class:"form-floating mb-1"},X=n(()=>a("label",{for:"floatingInput"},"Avatar name",-1)),Y={class:"col-6"},Z=n(()=>a("option",{value:"",disabled:"",selected:""},"Select a LLM",-1)),aa=["value"],ea={class:"row mb-5"},ta={class:"col-12"},sa=n(()=>a("h3",null,"Primer",-1)),oa=["disabled"],ra={key:0},la=n(()=>a("span",{class:"loader"},null,-1)),ca=[la],ia={key:1},na={__name:"CreateAvatarView",setup(d){const f=x(),S=D({AvatarName:v().required("The Avatar Name is required"),AvatarLLMID:v().required("The LLM is required"),AvatarPrimer:v()}),l=N({isSubmitting:!1}),A=R(),y=$(),g=F(),u=p(""),m=p(""),_=p(""),w=k(A,"records"),V=async()=>{l.isSubmitting=!0;try{const o={name:u.value,llm:m.value,primer:_.value};await y.saveAvatar(o),await g.replace({name:"personality",params:{avatar_id:y.userAvatar.ID}}),f.success("Avatar saved successfully"),l.isSubmitting=!1,window.location.reload()}catch(o){f.error(o),l.isSubmitting=!1}};return C(async()=>{try{await A.getLLMs()}catch(o){console.log(o)}feather.replace()}),(o,r)=>(i(),c("div",j,[z,a("div",G,[a("div",H,[a("div",J,[a("div",K,[s(t(E),{class:"container p-4",onSubmit:V,"validation-schema":t(S)},{default:L(()=>[a("div",O,[a("div",Q,[a("div",W,[s(t(h),{modelValue:u.value,"onUpdate:modelValue":r[0]||(r[0]=e=>u.value=e),name:"AvatarName",type:"text",class:"form-control",id:"floatingInput",placeholder:"Name your Avatar..."},null,8,["modelValue"]),X]),s(t(b),{name:"AvatarName"})]),a("div",Y,[s(t(h),{name:"AvatarLLMID",as:"select",modelValue:m.value,"onUpdate:modelValue":r[1]||(r[1]=e=>m.value=e),class:"form-select model-select mb-1","aria-label":"Select Model"},{default:L(()=>[Z,(i(!0),c(P,null,q(w.value,(e,I)=>(i(),c("option",{value:e.ID,key:I},B(e.name),9,aa))),128))]),_:1},8,["modelValue"]),s(t(b),{name:"AvatarLLMID"})])]),a("div",ea,[a("div",ta,[sa,s(t(h),{modelValue:_.value,"onUpdate:modelValue":r[2]||(r[2]=e=>_.value=e),name:"AvatarPrimer",type:"text",as:"textarea",class:"form-control mb-4",rows:"8",placeholder:""},null,8,["modelValue"]),s(t(b),{name:"AvatarPrimer"})])]),a("button",{type:"submit",class:"btn btn-secondary",disabled:l.isSubmitting},[l.isSubmitting?(i(),c("span",ra,ca)):(i(),c("span",ia,"Save"))],8,oa)]),_:1},8,["validation-schema"])])])])])]))}},_a=M(na,[["__scopeId","data-v-4b8bcc89"]]);export{_a as default};
