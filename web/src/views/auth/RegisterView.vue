<script setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router';
import { onMounted } from 'vue';
import { useAuthStore } from '@/stores/auth.store';
import { Form, Field } from 'vee-validate';

const route = useRoute();
const auth = useAuthStore();
const username = ref('');
const password = ref('');
const confirmPassword = ref('');

const confirmedPassword = () => {
  if (password.value === confirmPassword.value) {
    return true
  }
  else {
    return false
  }
};

const register = async () => {
  if (!confirmedPassword()) {
    alert('Passwords do not match.')
  }

  try {
    const loggedIn = await auth.registerInvite({
      username: username.value,
      password: password.value,
      invite_token: route.params.invite_token
    })

    if (loggedIn) {
      window.location.href = '/admin/avatar/create';
    }
  }
  catch (error) {
    console.log(error);
  }
}

onMounted(async () => {
  if (route.params.invite_token) {
    try {
      let user = await auth.getInvitedUser(route.params.invite_token);
      username.value = user.username;
    }
    catch (error) {
      console.log(error);
    }
  }
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
            
            <Form class="form-control" @submit="register">
              <Field v-model="username" id="Email" name="Username" type="email" class="email-input d-block" placeholder="Username"></Field>
              <Field v-model="password" id="Password" name="Password" type="password" class="pass-input d-block" placeholder="Password"></Field>
              <Field v-model="confirmPassword" id="confirmPassword" name="confirmPassword" type="password" class="pass-input d-block" placeholder="Confirm Password"></Field>
              <button class="send-button btn btn-light"> REGISTER </button>
          </Form>
            
          </div>
        </div>
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