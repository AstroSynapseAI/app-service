<script setup>
import { onMounted, ref } from 'vue';
import { Form, Field, useForm } from 'vee-validate';
import { useAuthStore } from '@/stores/auth.store.js'; 
import { useToast } from 'vue-toastification';

const auth = useAuthStore();
const { handleSubmit } = useForm();
const toast = useToast();

const isLoading = ref(false); 
let email = ref('');

const submitPasswordRecovery = async () => {
  isLoading.value = true
  try {
    const user = await auth.sendRecoverPasswordLink(email.value)
    toast.success("Email sent!");
  }
  catch (error) {
    toast.error(error)
  }
  finally {
    isLoading.value = false
  }
}

onMounted(() => {
  feather.replace();
});

</script>
<template>
  <div class="adminSpinner" v-if="isLoading"></div>
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
        <h3 class="px-3 mb-4 mt-3 mt-md-0"> Enter your email and ASAI will send you a link to reset your password</h3>
        <Form class="form-control" @submit="submitPasswordRecovery">
              <Field v-model="email" name="Email" type="email" class="email-input d-block" placeholder="Email"></Field>
              <button class="send-button btn btn-light">Reset</button>
        </Form>
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