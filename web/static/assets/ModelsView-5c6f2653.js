import{u as k}from"./llm.store-e461a386.js";import{u as D}from"./avatar.store-696f12a6.js";import{u as L,I as u,o as g,x as c,y as I,z as M,f as s,K as y,c as i,C as v,A as w,h as x,e as C,L as A,k as b}from"./index-510910de.js";const S={class:"container-fluid p-0"},V=s("h1",{class:"h3 mb-3"},"Models",-1),N={class:"card-header"},B={class:"row"},R={class:"col"},z={class:"card-title"},E={class:"col-auto"},F={key:0,class:"form-check form-switch d-flex align-items-center"},K=["checked","onClick"],T=s("label",{style:{"margin-bottom":"-5px"},for:"flexSwitchCheckDefault"},"Active",-1),U={class:"card-body"},J={__name:"ModelsView",setup($){const r=L(),d=D(),n=u(d,"activeLLMs"),l=k(),h=u(l,"records"),f=a=>{const e=n.value.find(o=>o.llm.ID==a);return e?e.is_active:!1},m=async a=>{const e=n.value.find(t=>t.llm.ID==a);e&&(e.is_active=!e.is_active);const o={is_active:e?e.is_active:!1,avatar_id:r.avatar.ID};try{await l.toggleActiveLLM(a,o)}catch(t){console.log(t)}},_=a=>{const e=n.value.find(o=>o.llm.ID===a);return e?e.ID:null};return g(async()=>{try{await l.getLLMs(),await d.getActiveLLMs(r.avatar.ID)}catch(a){console.log(a)}feather.replace()}),(a,e)=>{const o=y("router-link");return i(),c("div",S,[V,(i(!0),c(I,null,M(h.value,(t,p)=>(i(),c("div",{class:"card",key:"row"+p},[s("div",N,[s("div",B,[s("div",R,[s("h5",z,v(t.name),1)]),s("div",E,[_(t.ID)?(i(),c("div",F,[s("input",{class:"form-check-input me-2",type:"checkbox",id:"flexSwitchCheckDefault",checked:f(t.ID),onClick:j=>m(t.ID)},null,8,K),T])):w("",!0)])])]),s("div",U,[s("p",null,v(t.description),1),s("div",null,[x(o,{to:{name:"model-config",params:{avatar_id:b(r).avatar.ID,model_id:t.ID,active_model_id:_(t.ID)}},class:"btn btn-primary"},{default:C(()=>[A(" Configure ")]),_:2},1032,["to"])])])]))),128))])}}};export{J as default};