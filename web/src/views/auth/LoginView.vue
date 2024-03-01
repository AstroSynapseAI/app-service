<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { onMounted } from 'vue';
import { Form, Field, useForm } from 'vee-validate';
import { useAuthStore } from '@/stores/auth.store.js'; 
import { useUserStore } from '@/stores/user.store.js';
const { handleSubmit } = useForm();

const router = useRouter();
const auth = useAuthStore();
const user = useUserStore();

let username = ref('');
let password = ref('');

const submitLogin = handleSubmit(async values => {
  try {
    const loggedIn = await auth.login(username.value, password.value)
    if (loggedIn) {
      await user.getUserAvatar(auth.currentUser.ID);
      if (user.avatar) {
        router.push({name: 'admin', params: { avatar_id: user.avatar.ID }});
        return
      }
      router.push({name: 'create-avatar'});
    }
    else {
      alert('Invalid username or password');
    }
  }
  catch (err) {
    console.log(err);
  }
});


onMounted(() => {
  feather.replace();
});

</script>
<template>
  <div class="container d-flex flex-column vh-100">
    <nav class="navbar navbar-expand-md bg-dark bg-transparent">
      <div class="container-fluid">
        <div class="row w-100">
          <div class="col-6">
            <span class="navbar-brand text-white">AI Avatars Platform</span>
          </div>

          <div class="navbar-menu col-6 d-flex justify-content-end">

            <div class="row">
              <div class="col-auto">
                <router-link :to="{name: 'home'}" class="btn text-white">
                  <i class="align-middle feather-icon" data-feather="home"></i>
                </router-link>
              </div>
            </div>

          </div>
        </div>
      </div>
    </nav>
    <div class="row">

      <div class="col-md-6">
        <img class="logo" src="@/assets/ASAILogotype.svg" alt="">
        <div class="card">
          <div class="card-body">
            
            <Form class="form-control" @submit="submitLogin">
              <Field v-model="username" id="Email" name="Username" type="email" class="email-input d-block" placeholder="Username"></Field>
              <Field v-model="password" id="Password" name="Password" type="password" class="pass-input d-block" placeholder="Password"></Field>
              <button class="send-button btn btn-light"> Login</button>
          </Form>
            
          </div>
        </div>
      </div>
      
      <div class="col-md-6">
        <h3 class="px-3 mb-4 mt-3 mt-md-0"> Asai cloud is currently in <b>closed beta</b>, and access is limited to <b>invite only</b>. Plese send us your email, if you are interested, and we will add you in the next onbaording batch of testers.</h3>
        <Form class="form-control d-flex" action="https://formspree.io/f/xyyqjdgr" method="POST">
          <Field id="waitlist-email" name="WaitList Email" type="email" class="email-input flex-fill mb-0 corner-0" placeholder="Email"></Field>
          <button class="send-button btn btn-light" @click="''">Submit</button>
        </Form>
      </div>
    
    </div>
  </div>
</template>

<style scoped>
a {
  color: white !important;
}

nav {
  margin-top: 50px;
  margin-bottom: 15em;
}

.navbar-brand {
  color: white !important;
}

h1, h2, h3, h4, h5, h6 {
  color: white;
}


.feather-icon {
  width: 24px !important;
  height: 24px !important;
}

.logo {
  width: 100%;
  margin-bottom: 3em;
}
.card {
  background-color: black; 
  border: 1px solid white; 
  border-radius: 0;
  width: 100%;
}
.card-body {
  color: white; /* To make text visible in dark background */
}

.form-control {
  background-color: black;
  border: none;
  width: 100%;
  color: white;
  border-radius: 0;
}

.email-input, .pass-input {
  width: 100%;
  margin-bottom: 2em;
  height: 3em;
  background-color: transparent;
  border: 1px solid white;
  color: white;
  border-radius: 0;
}

.btn {
  background-color: black;
  border-color: white;
  border-radius: 0;
  text-decoration: none;
  color: white;
}

.btn:hover {
  color: black !important;
  background-color: white;
}

@media only screen and (max-width: 576px) {
  nav {
    margin-top: 10px;
    margin-bottom: 2em;
  }
}
</style>