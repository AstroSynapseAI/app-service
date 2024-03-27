import{I as b,B as V,r as i,o as D,c as f,x as h,f as e,L as M,C as y,h as T,k as C,F as I,M as l,T as w,y as S,z as U,N as c,R as N}from"./index-510910de.js";import{u as P}from"./agent.store-3dac9ca8.js";import{u as R}from"./llm.store-e461a386.js";import{u as E}from"./avatar.store-696f12a6.js";const B={class:"h3 mb-3"},F={class:"form-check form-switch float-end me-5"},J=e("label",{class:"form-check-label",for:"flexSwitchCheckDefault"},"Active",-1),O=["checked"],$={class:"form-check form-switch float-end me-3"},z=e("label",{class:"form-check-label",for:"flexSwitchCheckDefault"},"Public",-1),H=["checked"],j={class:"card"},q={class:"card-body"},G={class:"container-fluid"},K={class:"row"},Q={class:"col-6"},W=e("h3",null,"Primer",-1),X={class:"col-6"},Y={class:"row"},Z={class:"col-12"},ee=e("h3",null,"Agent Model",-1),te=e("option",{value:"",disabled:"",selected:""},"Select a LLM",-1),oe=["value"],se={class:"row"},le={clas:"col-12"},ae=e("h3",null,"Email Config",-1),ne={class:"row"},ie={class:"col-6"},ce={class:"form-floating mb-3"},re=e("label",{for:"smtp-server"},"SMTP Server",-1),de={class:"col-4"},pe={class:"form-floating mb-3"},ue=N('<option value="" disabled selected>Select Encryption Type</option><option value="ssl">SSL</option><option value="tls">TLS</option><option value="starttls">STARTTLS</option><option value="ssltls">SSL/TLS</option><option value="none">None</option>',6),ve=[ue],me=e("label",{for:"smtp-encryption"},"Encryption",-1),_e={class:"col-2"},fe={class:"form-floating mb-3"},he=e("label",{for:"smtp-port"},"Port",-1),ge={class:"row"},be={class:"col-6"},ye={class:"form-floating mb-3"},we=e("label",{for:"smtp-username"},"Username",-1),Se={class:"col-6"},ke={class:"form-floating mb-3"},Ae=e("label",{for:"smtp-password"},"Password",-1),xe={class:"row"},Le={class:"col-6"},Ve={class:"form-floating mb-3"},De=e("label",{for:"sender"},"Sender",-1),Me={class:"col-6"},Te={class:"form-floating mb-3"},Ce=e("label",{for:"reply-to"},"Reply To",-1),Ee={__name:"EmailAgent",setup(Ie){const a=E(),_=P(),r=b(_,"record"),n=V(),g=R(),k=b(g,"records"),d=i(""),A=i(""),m=i(""),p=i(!1),u=i(!1),s=i({imap_server:"",smtp_server:"",imap_port:"",smtp_port:"",username:"",password:"",encryption:"",sender:"",reply_to:""}),x=async()=>{const v={ID:parseInt(n.params.active_agent_id),avatar_id:parseInt(n.params.avatar_id),agent_id:parseInt(n.params.agent_id),llm_id:m.value,primer:d.value,is_active:p.value,is_public:u.value,config:JSON.stringify(s.value)};r.value.ID&&(v.ID=r.value.ID);try{await _.saveActiveAgent(v)}catch(t){console.log(t)}};return D(async()=>{try{await g.getLLMs(),await _.getAgent(n.params.agent_id),d.value=r.value.primer,A.value=r.value.name,n.params.active_agent_id&&(await a.getActiveAgent(n.params.agent_id,n.params.avatar_id),m.value=a.activeAgent.llm_id,d.value=a.activeAgent.primer,p.value=a.activeAgent.is_active,u.value=a.activeAgent.is_public,a.activeAgent.config&&(s.value=JSON.parse(a.activeAgent.config)))}catch(v){console.log(v)}feather.replace()}),(v,t)=>(f(),h(S,null,[e("h1",B,[M("Configure: "+y(r.value.name)+" ",1),e("div",F,[J,e("input",{class:"form-check-input",type:"checkbox",id:"flexSwitchCheckDefault",checked:p.value,onClick:t[0]||(t[0]=o=>p.value=!p.value)},null,8,O)]),e("div",$,[z,e("input",{class:"form-check-input",type:"checkbox",id:"flexSwitchCheckDefault",checked:u.value,onClick:t[1]||(t[1]=o=>u.value=!u.value)},null,8,H)])]),e("div",j,[e("div",q,[e("div",G,[e("div",K,[e("div",Q,[W,T(C(I),{modelValue:d.value,"onUpdate:modelValue":t[2]||(t[2]=o=>d.value=o),name:"agent_primer",type:"text",as:"textarea",class:"form-control",rows:"30",placeholder:"How should the agent behave..."},null,8,["modelValue"])]),e("div",X,[e("div",Y,[e("div",Z,[ee,l(e("select",{"onUpdate:modelValue":t[3]||(t[3]=o=>m.value=o),class:"form-select model-select","aria-label":"Select Model"},[te,(f(!0),h(S,null,U(k.value,(o,L)=>(f(),h("option",{value:o.ID,key:L},y(o.name),9,oe))),128))],512),[[w,m.value]])])]),e("div",se,[e("div",le,[ae,e("div",ne,[e("div",ie,[e("div",ce,[l(e("input",{type:"text",class:"form-control",id:"smtp-server",placeholder:"mail.example.com","onUpdate:modelValue":t[4]||(t[4]=o=>s.value.smtp_server=o)},null,512),[[c,s.value.smtp_server]]),re])]),e("div",de,[e("div",pe,[l(e("select",{"onUpdate:modelValue":t[5]||(t[5]=o=>s.value.encryption=o),class:"form-select model-select","aria-label":"Select Model"},ve,512),[[w,s.value.encryption]]),me])]),e("div",_e,[e("div",fe,[l(e("input",{type:"text",class:"form-control",id:"smtp-port",placeholder:"587","onUpdate:modelValue":t[6]||(t[6]=o=>s.value.smtp_port=o)},null,512),[[c,s.value.smtp_port]]),he])])]),e("div",ge,[e("div",be,[e("div",ye,[l(e("input",{type:"text",class:"form-control",id:"smtp-username",placeholder:"username","onUpdate:modelValue":t[7]||(t[7]=o=>s.value.username=o)},null,512),[[c,s.value.username]]),we])]),e("div",Se,[e("div",ke,[l(e("input",{type:"password",class:"form-control",id:"smtp-password",placeholder:"password","onUpdate:modelValue":t[8]||(t[8]=o=>s.value.password=o)},null,512),[[c,s.value.password]]),Ae])])]),e("div",xe,[e("div",Le,[e("div",Ve,[l(e("input",{type:"text",class:"form-control",id:"sender",placeholder:"sender","onUpdate:modelValue":t[9]||(t[9]=o=>s.value.sender=o)},null,512),[[c,s.value.sender]]),De])]),e("div",Me,[e("div",Te,[l(e("input",{type:"text",class:"form-control",id:"reply-to",placeholder:"reply-to","onUpdate:modelValue":t[10]||(t[10]=o=>s.value.reply_to=o)},null,512),[[c,s.value.reply_to]]),Ce])])])])])])]),e("div",{class:"row mt-3"},[e("div",{class:"col-12"},[e("button",{type:"button",class:"btn btn-secondary",onClick:x},"Save")])])])])])],64))}};export{Ee as default};