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
        
      </div>
      
      <div class="col-md-6">
       
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