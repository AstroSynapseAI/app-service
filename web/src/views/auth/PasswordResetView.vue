<script setup>
import { reactive, ref } from 'vue';
import { useRoute } from 'vue-router';
import { onMounted } from 'vue';
import { useAuthStore } from '@/stores/auth.store';
import { Form, Field, ErrorMessage } from 'vee-validate';
import { useToast } from 'vue-toastification';
import * as yup from 'yup';

const toast = useToast();
const schema = yup.object({
  Password: yup.string().required().min(8),
  ConfirmPassword: yup.string().required('Please confirm your password').min(8)
    .oneOf([yup.ref('Password'), null], 'Passwords must match')
});

const route = useRoute();
const auth = useAuthStore();
const password = ref('');
const confirmPassword = ref('');

const formState = reactive({
  isSubmitting: false, 
});

const register = async () => {
  formState.isSubmitting = true; 
  try {
     const loggedIn = await auth.registerInvite({
       username: username.value,
       password: password.value,
       invite_token: route.params.invite_token
     });

     if (loggedIn) {
       window.location.href = '/admin/avatar/create';
     }
  }
  catch (error) {
    toast.error(error)
    formState.isSubmitting = false; 
  }
};
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
            <Form class="form-control" @submit="register" :validation-schema="schema">
<ErrorMessage name="Password" />
              <Field v-model="password" name="Password" type="password" class="pass-input d-block" placeholder="New Password"/>
<ErrorMessage name="ConfirmPassword" />
              <Field v-model="confirmPassword" name="ConfirmPassword" type="password" class="pass-input d-block" placeholder="Confirm Password"/>
              <button class="send-button btn btn-light" :disabled="formState.isSubmitting">
                <span v-if="formState.isSubmitting">
                  <span class="loader"></span>
                </span>
                <span v-else>RESET PASSWORD</span>
              </button>
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

.loader {
    width: 18px;
    height: 18px;
    border: 2px solid #FFF;
    border-bottom-color: transparent;
    border-radius: 50%;
    display: inline-block;
    box-sizing: border-box;
    animation: rotation 1s linear infinite;
    }

    @keyframes rotation {
    0% {
        transform: rotate(0deg);
    }
    100% {
        transform: rotate(360deg);
    }
} 

.send-button {
  min-width: 150px;
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