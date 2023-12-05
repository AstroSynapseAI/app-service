<script setup>
import PromptInput from '@/components/PromptInput.vue';
import { useChatStore } from '@/stores/chat.store.js';
import { storeToRefs } from 'pinia';

const chatStore = useChatStore();
const { messages } = storeToRefs(chatStore);

</script>

<template>  
  <div class="container-fluid p-0">
    <h1 class="h3 mb-3">Avatar</h1>
    <div class="row">
      <div class="col-12">
        <div class="card">
          <div class="card-body">
            <div class="container p-4 d-flex flex-column">
              <div class="conversation-container flex-grow-1">
                
                <template v-if="messages.length > 0">
                
                  <div class="conversation-item row" v-for="(message, index) in messages" :key="index">
                  
                    <div class="col-1">
                      <img src="../assets/avatar.png" class="logo" alt="Avatar Icon" width="35" v-if="message.type === 'ai'"/>
                      <img src="../assets/user.svg" class="logo" alt="User Icon" width="35" height="50" v-if="message.type === 'user'"/>
                    </div>
                  
                    <div class="col-11">  
                      <p class="message-content">{{ message.text }}</p>
                    </div>
                    
                    <hr class="separator opacity-100" v-if="messages.length > 1 && index !== messages.length - 1">

                  </div>

                </template>

              </div>
            
              <div class="prompt-container">
                <hr class="border border-3 opacity-100">
                <PromptInput />
              </div>
            </div>

          </div>
        </div>
      </div>

      
    
    </div>
  </div>
</template>

<style scoped>
.container-fluid {
  color: white;
}
.card-body .container {
  min-height: 85vh;
}
</style>