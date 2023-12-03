<script setup>
import { Form, Field } from 'vee-validate';
import { ref, onMounted } from 'vue';

let showModal = ref(false);

const toggleModal = () => {
  showModal.value = !showModal.value;
}

onMounted(() => {
  feather.replace();
});
</script>

<template>
          
  <div class="container-fluid p-0">
    
    <h1 class="h3 mb-3">Plugins</h1>
    <div class="row">
      <div class="col-12">
        <div class="container">
          <!-- Edit Configure Plugin -->
          <div class="modal" :class="{ 'd-block': showModal, 'show': showModal }" tabindex="-1">
            <div class="modal-dialog modal-dialog-centered modal-lg">
              <div class="modal-content">
                <div class="modal-header">
                  <h5 class="modal-title">Configure Plugin</h5>
                  <button type="button" class="btn-close" @click="showModal = false"></button>
                </div>
                <div class="modal-body container">
                  <div class="row">
                      <div class="col-12">
                        <Field 
                          name="token" 
                          type="text" 
                          as="input" 
                          class="form-control" 
                          rows="8" 
                          placeholder="API Token"
                        ></Field>
                      </div>
                    </div>
                </div>
                <div class="modal-footer">
                  <button type="button" class="btn btn-secondary" @click="showModal = false">Save</button>
                </div>
              </div>
            </div>
          </div>
          
          <div class="row">

            <div class="col-6">
              <div class="card">
                <div class="card-header">
                  <div class="row">
                    <div class="col-6">
                      <h5 class="card-title">Slack</h5>
                    </div>
                    <div class="col-6">
                      <div class="form-check form-switch float-end">
                        <span @click="toggleModal" style="cursor: pointer; margin-right: 5px;">
                          <i data-feather="settings" class="float-end"></i>
                        </span>
                        <label class="form-check-label me-2" for="flexSwitchCheckDefault">Active</label>
                        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                      </div>
                    </div>
                  </div>

                  
                </div>
                <div class="card-body">
                  <div class="row">
                    <div class="col-2 d-flex justify-content-center align-items-center">
                      <i data-feather="slack" class="" style="width: 48px; height: 48px"></i>
                    </div>

                    <div class="col-10">
                      <p>Empower your conversations with your AI Avatar on Slack! This dynamic feature translates easily into an assistant for your team, answering inquiries, simplifying tasks, and streamlining information flow like a pro. All this, right within your existing team chats!</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div class="col-6">
              <div class="card">
                <div class="card-header">
                  <div class="row">
                    <div class="col-6">
                      <h5 class="card-title">Discord</h5>
                    </div>
                    <div class="col-6">
                      <div class="form-check form-switch float-end">
                        <span @click="toggleModal" style="cursor: pointer; margin-right: 5px;">
                          <i data-feather="settings" class="float-end"></i>
                        </span>
                        
                        <label class="form-check-label me-2" for="flexSwitchCheckDefault">Active</label>
                        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault">
                      </div>
                    </div>
                  </div>

                  
                </div>
                <div class="card-body">
                  <div class="row">
                    <div class="col-2 d-flex justify-content-center align-items-center">
                      <img src="../assets/discord-icon.jpg" width="108" height="108">
                    </div>

                    <div class="col-10">
                      <p>Welcome your AI Avatar into your Discord channels, ready to redefine your community interactions! Providing instant responses and meaningful insights, this feature seamlessly merges with your discussions, adding a layer of efficiency to your community engagements!</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            
          </div>


        </div>
      </div>
    </div>
  </div>
      
</template>

<style scoped>
.card {
    min-height: 220px;
  }
</style>