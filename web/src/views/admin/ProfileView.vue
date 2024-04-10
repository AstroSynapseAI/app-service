<script setup>
import { onMounted, ref, computed } from 'vue';
import { useUserStore } from '@/stores/user.store';
import { useToast } from 'vue-toastification';

const user = useUserStore();
const username = ref('');
const newPassword = ref('');
const confirmPassword = ref('');
const firstName = ref('');
const lastName = ref('');
const email = ref('');
const newEmail = ref('');
const confirmEmail = ref('');
const isLoading = ref(false);
const isTyping = ref(false);

const toast = useToast();

const checkPasswordMatch = () => {
  isTyping.value = true;
};

const isSaveButtonDisabled = computed(() => {
  return !username.value.trim() || !firstName.value.trim() || !lastName.value.trim();
});

const passwordBtnDisabled = computed(() => {
  if (!newPassword.value.trim() || !confirmPassword.value.trim() || confirmPassword.value.length < 8) {
    return true;
  }
  if (newPassword.value !== confirmPassword.value) {
    return true;
  }
  return false;
});

const emailBtnDisabled = computed(() => {
  return !isValidEmail()
});

const validateEmail = () => {
  const emailRegex = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  return String(newEmail.value).toLowerCase().match(emailRegex);
};

const confirmedEmail = () => {
  return newEmail.value === confirmEmail.value;
};

const isValidEmail = () => {
  var validEmail = false
  validEmail = confirmedEmail()
  validEmail = validateEmail()

  return validEmail
};


const saveUserInfo = async () => {
  isLoading.value = true
  try {
    const profileData = {
      username: username.value,
      first_name: firstName.value,
      last_name: lastName.value
    }

    if (user.account) {
      profileData.account_id = user.account.ID
    }

    await user.saveProfile(user.current.ID, profileData)
    toast.success("Profile updated successfully!");
  }
  catch (error) {
    toast.error(error)
  }
}

const updateEmail = async () => {
  isLoading.value = true
  if (!isValidEmail()) {
    toast.error("Email is not valid or emails do not match.")
    return
  }
  try {
    await user.changeEmail(user.current.ID, {
      account_id: user.account.ID,
      email: newEmail.value,
    })
    toast.success("Email updated successfully!");
  }
  catch (error) {
    toast.error(error)
  }
}

const changePassword = async () => {
  isLoading.value = true
  try {
    await user.changePassword(user.current.ID, {
      password: newPassword.value
    })
    toast.success("Password changed successfully!");
  }
  catch (error) {
    toast.error(error)
  }
  finally {
    isLoading.value = false
  }
}

onMounted(async () => {
  username.value = user.current.username;
  try {
    await user.getUserAccounts(user.current.ID);
    if (user.account) {
      username.value = user.account.username ?? "";
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
                <input type="text" class="form-control" id="username" placeholder="Username" v-model="username"
                  @input="onUserDataChange">
                <label for="username">Username</label>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="text" class="form-control" id="firstName" placeholder="First Name" v-model="firstName"
                  @input="onUserDataChange">
                <label for="firstName">First Name</label>
              </div>
            </div>
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="text" class="form-control" id="lastName" placeholder="Last Name" v-model="lastName"
                  @input="onUserDataChange">
                <label for="lastName">Last Name</label>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col-12">
              <button class="btn btn-primary float-end" @click="saveUserInfo"
                :disabled="isSaveButtonDisabled">Save</button>
            </div>
          </div>
          <hr>

          <h3>Change Email <span class="text-muted">{{ email }}</span></h3>
          <div class="row">
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="email" class="form-control" id="newEmail" placeholder="New Email" v-model="newEmail"
                  @input="onEmailChange">
                <label for="newEmail">New Email</label>
              </div>
            </div>
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="email" class="form-control" id="confirmEmail" placeholder="Confirm Email"
                  v-model="confirmEmail" @input="onEmailChange">
                <label for="confirmPassword">Confirm Email</label>
                <div v-if="isTyping && newEmail !== confirmEmail">
                  Emails do not match
                </div>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col-12">
              <button class="btn btn-primary float-end" @click="updateEmail" :disabled="emailBtnDisabled">Change
                Email</button>
            </div>
          </div>
          <hr>


          <h3>Change Password</h3>
          <div class="row">
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="password" class="form-control" id="newPassword" placeholder="New Password"
                  v-model="newPassword" @input="checkPasswordMatch">
                <label for="newPassword">New Password</label>
              </div>
            </div>
            <div class="col-6">
              <div class="form-floating mb-3">
                <input type="password" class="form-control" id="confirmPassword" placeholder="Confirm Password"
                  v-model="confirmPassword" @input="checkPasswordMatch">
                <label for="confirmPassword">Confirm Password</label>
                <div v-if="isTyping && newPassword !== confirmPassword">
                  Passwords do not match
                </div>
                <div v-else-if="isTyping && (newPassword.length < 8 || confirmPassword.length < 8)">
                  New password has to have more than 8 characters
                </div>
              </div>
            </div>
          </div>
          <div class="row">
            <div class="col-12">
              <button class="btn btn-primary float-end" @click="changePassword" :disabled="passwordBtnDisabled">Change
                Password</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
