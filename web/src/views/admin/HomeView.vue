<script setup>
import PromptInput from '@/components/admin/PromptInput.vue';
import { onMounted, watch, ref, computed } from 'vue';
import { useRoute } from 'vue-router';
import { storeToRefs } from 'pinia';
import { useChatStore } from '@/stores/chat.store';
import { useUserStore } from '@/stores/user.store';
import MarkdownIt from 'markdown-it';
import iconASAI from '../../assets/asai-icon.png';
import iconUser from '../../assets/user-icon.png';

const route = useRoute();

const chatStore = useChatStore();
const { messages } = storeToRefs(chatStore, "messages");

const user = useUserStore();

const conversationContainer = ref(null);
const promptContainer = ref(null);
const promptInput = ref('');

const md = new MarkdownIt({ breaks: true });
const lastMessageIndex = computed(() => messages.value.length - 1)
const lastMessageText = computed(() => messages.value[lastMessageIndex.value]?.content || "")

async function scrollToBottom() {
  requestAnimationFrame(() => {
    if (conversationContainer.value) {
      var offset = promptContainer.value.offsetHeight + 30;
      var scrollTo = conversationContainer.value.scrollHeight + offset
      conversationContainer.value.scrollTop = scrollTo
    }
  });
}

watch(messages, () => {
  scrollToBottom();
});

watch(lastMessageText, () => {
  scrollToBottom();
  feather.replace();
});

onMounted(async () => {
  try {
    if (user.session_id === null) {
      // generate new token and session
      await user.getSessionToken();
    } else {
      // fetch the session messages
      await chatStore.getPrivateSession(route.params.avatar_id, user.session_id);
    }
  }
  catch (error) {
    console.log(error);
  }
  scrollToBottom();
  feather.replace();
  chatStore.connectWebSocket();
});

const regeneratePrompt = function regeneratePrompt(index) {
  const lastHumanMessage = messages.value[index] || {};

  chatStore.sendPrompt({
    session_id: user.session_id,
    avatar_id: user.avatar.ID,
    msg: lastHumanMessage.content || '',
  });
}

const editPrompt = function editPrompt(index) {
  const lastHumanMessage = messages.value[index] || {};
  promptInput.value = lastHumanMessage.content || '';
  scrollToBottom();
}

</script>

<template>
  <div class="container-fluid">
    <h1 class="h3 mb-3">Avatar</h1>
    <div class="row">
      <div class="col-12">
        <div class="card">
          <div class="card-body">
            <div class="container d-flex flex-column">

              <div ref="conversationContainer" class="conversation-container flex-grow-1 overflow-auto">

                <template v-if="messages.length > 0">

                  <div class="conversation-item row" v-for="(message, index) in messages" :key="index">

                    <div class="col-1 d-flex flex-column">
                      <img :src="iconASAI" class="logo" alt="Avatar Icon" width="35" v-if="message.sender === 'ai'" />
                      <img :src="iconUser" class="logo" alt="User Icon" width="35" height="50"
                        v-if="message.sender === 'human'" />
                      <button v-if="message.sender === 'ai' && (index + 1) == messages.length"
                        @click="() => regeneratePrompt(index - 1)" class="retry-button btn mt-2 p-0"
                        :disabled="chatStore.isLoading" title="Regenerate prompt">
                        <i class="align-middle" data-feather="refresh-cw"></i>
                      </button>
                      <button v-if="message.sender === 'human'" @click="() => editPrompt(index)"
                        class="retry-button btn mt-2 p-0" :disabled="chatStore.isLoading" title="Edit last prompt">
                        <i class="align-middle" data-feather="edit"></i>
                      </button>
                    </div>

                    <div class="col-10">
                      <div v-if="message.isLoading">
                        <p><span class="me-3">I'm thinking... </span><span class="spinner mb-2 me-2"><img
                              v-if="!message.isAgentRunnig" src="../../assets/loader.png" alt=""></span></p>
                        <p v-if="message.isAgentRunnig"><span class="me-3 fst-italic">Activating agent: {{
                  message.agentName }}... </span><span class="spinner mb-2 me-2"><img
                              src="../../assets/loader.png" alt=""></span></p>
                      </div>

                      <div v-else class="message-content pe-3" v-html="md.render(message.content.trim())"></div>
                    </div>

                    <hr class="separator opacity-100 mt-3" v-if="messages.length > 1 && index !== messages.length - 1">

                  </div>

                </template>

              </div>

              <div ref="promptContainer" class="prompt-container">
                <hr class="border border-3 opacity-100">
                <PromptInput :promptInput="promptInput" :conversationContainer="conversationContainer" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
main {
  padding-bottom: 0 !important;
}

.container-fluid {
  color: white;
  margin-bottom: 0;
  padding-bottom: 0;
}

.conversation-container {
  max-height: calc(100vh - 300px);
  padding: 1.25rem;
  scrollbar-width: none;
  /* Firefox */
  -ms-overflow-style: none;
  /* IE and Edge */
}

.conversation-container ::-webkit-scrollbar {
  display: none;
  /* Chrome, Safari */
  padding: 1.5rem;
}

.card {
  padding-bottom: 0 !important;
  margin-bottom: 0 !important;
}

.card-body .container {
  min-height: 80vh;
}

.card-body {
  background-color: #19232E !important;
  color: white !important;
  padding-bottom: 0 !important;
  margin-bottom: 0 !important;
}

.retry-button {
  vertical-align: middle;
  border-color: transparent !important;
  width: 35px;
}

.retry-button:disabled {
  opacity: .7;
}

.retry-button .feather {
  /* width: 20px;
  height: 20px; */
  margin: 0 auto;
}

@keyframes rotate {
  0% {
    transform: rotate(0deg);
  }

  25% {
    transform: rotate(90deg);
  }

  50% {
    transform: rotate(180deg);
  }

  75% {
    transform: rotate(270deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

.spinner img {
  display: inline-block;
  vertical-align: middle;
  transform-origin: 50% 50%;
  animation: rotate 0.5s linear infinite;
  height: 18px;
  margin-bottom: 5px;
}

@media (min-width: 992px) {
  .content {
      padding-bottom: 0 !important;
  }
}

</style>
