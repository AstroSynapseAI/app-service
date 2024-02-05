<script setup> 
import { onMounted, ref } from 'vue';
import { useUserStore } from '@/stores/user.store';

const user = useUserStore();
const username = ref('');
const newPassword = ref('');
const confirmPassword = ref('');
const firstName = ref('');
const lastName = ref('');
const email = ref('');
const confirmEmail = ref('');

const validateEmail = () => {
  return String(email.value)
    .toLowerCase()
    .match(
      /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    );
};

const confirmedEmail = () => {
  if (email.value === confirmEmail.value) {
    return true
  }
  else {
    return false
  }
};

const isValidEmail = () => {
  var validEmail = false
  validEmail = confirmedEmail()
  validEmail = validateEmail()

  return validEmail
};

const confirmedPassword = () => {
  if (newPassword.value === confirmPassword.value) {
    return true
  }
  else {
    return false
  }
};

const saveUserInfo = async () => {
  try {
    const profileData = {
      username: username.value,
      first_name: firstName.value,
      last_name: lastName.value,
      email: email.value
    }

    if (user.account) {
      profileData.account_id = user.account.ID
    }

    user.saveProfile(user.current.ID, profileData)
  }
  catch (error) {
    console.error(error)
  }
}

const updateEmail = async () => {
  if (!isValidEmail()) {
    alert('Email is not valid or emails do not match.')
  }

  try {
    user.changeEmail(user.current.ID, {
      account_id: user.account.ID,
      email: email.value,
    })
  }
  catch (error) {
    console.error(error)
  }
}

const changePassword = async () => {
  if (!confirmedPassword()) {
    alert('Passwords do not match.')
  }

  try {
    user.changePassword(user.current.ID,{
      password: newPassword.value
    })
  }
  catch (error) {
    console.error(error)
  }
}

onMounted(async () => {
  username.value = user.current.username;
  try {
    await user.getUserAccounts(user.current.ID);
    if (user.account) {
      firstName.value = user.account.first_name;
      lastName.value = user.account.last_name;
      email.value = user.account.email;
    }
  }
  catch (error) {
    console.error(error)
  }
  feather.replace()
})


</script>

<template>
          
  <div class="container-fluid p-0">
    
    <h1 class="h3 mb-3">Account</h1>
    
    <div class="card">

      <div class="card-body">
        <div class="container p-4">

          <h3>User Information</h3>

          <div class="row">
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="text" class="form-control" id="username" placeholder="Username" v-model="username">
                <label for="username">Username</label>
              </div>
            </div>
          </div>

          <div class="row">
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="text" class="form-control" id="firstName" placeholder="First Name" v-model="firstName">
                <label for="firstName">First Name</label>
              </div>
            </div>
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="text" class="form-control" id="lastName" placeholder="Last Name" v-model="lastName">
                <label for="lastName">Last Name</label>
              </div>
            </div>
          </div>

          <div class="row"> 
            <div class="col-12">
              <button class="btn btn-primary float-end" @click="saveUserInfo">Save</button>
            </div>
          </div>

          <hr>

          <!-- <h3>Update Email</h3>

          <div class="row">
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="text" class="form-control" id="email" placeholder="Email" v-model="email">
                <label for="email">Email</label>
              </div>
            </div>
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="text" class="form-control" id="confirmEmail" placeholder="Confirm Email" v-model="confirmEmail">
                <label for="confirmEmail">Confirm Email</label>
              </div>
            </div>
          </div>

        
          <div class="row"> 
            <div class="col-12">
              <button class="btn btn-primary float-end" @click="updateEmail">Update</button>
            </div>
          </div> -->
          
         

          <hr>

          <h3>Change Password</h3>

          <div class="row">
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="password" class="form-control" id="newPassword" placeholder="New Password" v-model="newPassword">
                <label for="newPassword">New Password</label>
              </div>
            </div>
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="password" class="form-control" id="confirmPassword" placeholder="Confirm Password" v-model="confirmPassword">
                <label for="confirmPassword">Confirm Password</label>
              </div>
            </div>
          </div>
        
          <div class="row"> 
            <div class="col-12">
              <button class="btn btn-primary float-end" @click="changePassword">Change</button>
            </div>
            
          </div>

        </div>
      </div>
      
    </div>
  
  </div>
      
</template>