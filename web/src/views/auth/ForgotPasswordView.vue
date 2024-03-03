<script setup>
import { onMounted, ref, reactive } from 'vue';
import { Form, Field, ErrorMessage } from 'vee-validate';
import { useAuthStore } from '@/stores/auth.store.js'; 
import { useToast } from 'vue-toastification';
import * as yup from 'yup';

const auth = useAuthStore();
const toast = useToast();
let email = ref('');

const schema = yup.object({
  Email: yup.string().email().required(),
});

const formState = reactive({
  isSubmitting: false, 
});

const submitPasswordRecovery = async () => {
  formState.isSubmitting = true; 
  try {
    const user = await auth.sendRecoverPasswordLink(email.value)
    toast.success("Email sent!");
  }
  catch (error) {
    toast.error(error)
  }
  finally {
    formState.isSubmitting = false; 
  }
}

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
        <h3 class="px-3 mb-4 mt-3 mt-md-0"> Enter your email and ASAI will send you a link to reset your password</h3>
        <Form class="form-control" @submit="submitPasswordRecovery" :validation-schema="schema">
              <Field v-model="email" name="Email" type="email" class="email-input d-block" placeholder="Email"></Field>
              <ErrorMessage name="Email" />
              <button class="send-button btn btn-light" :disabled="formState.isSubmitting">
                <span v-if="formState.isSubmitting">
                  <span class="loader"></span>
                </span>
                <span v-else>RESET</span>
              </button>
        </Form>
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

.email-input {
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