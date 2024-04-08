<script setup>
import { ref, onMounted, watch } from 'vue';
import { storeToRefs } from 'pinia';
import { Form, Field, useForm } from 'vee-validate';
import { useChatStore } from '@/stores/chat.store.js';
import { useUserStore } from '@/stores/user.store.js';

const user = useUserStore();
const chatStore = useChatStore();
const MAX_ROWS = 10;
const INPUT_MIN_SIZE = 46;
const CONVERSATION_SIZE_OFFSET = 300;
const INPUT_PADDING = 4;

const { isLoading } = storeToRefs(chatStore);


const props = defineProps({
  promptInput: {
    type: String,
    required: false,
  },
  conversationContainer: {
    type: Object,
    required: false,
  }
});

let { resetForm, handleSubmit, defineField } = useForm({
  initialValues: {
    prompt: ''
  }
});

const [prompt, promptAttrs] = defineField('prompt');
const promptElement = ref(null);
let inputFieldLineHeight = ref(0);
let inputRowNum = ref(0);

watch(() => props.promptInput, (value) => {
  prompt.value = value;
  promptElement.value.$el.focus();
});


watch(isLoading, (value, oldValue) => {
  if (oldValue === true && value === false) {
    setTimeout(() => {
      promptElement.value.$el.focus();
    }, 0);
  }
});

function getInputLineHeight() {
  return parseInt(window.getComputedStyle(promptElement.value.$el).getPropertyValue("line-height"));
}

function getInputRowNumber() {
  return (promptElement.value.$el.value.match(/\n/g) || []).length;
}

function updateConversationContainerSize(heightDifference) {
  let containerSizeOffset = CONVERSATION_SIZE_OFFSET + heightDifference;
  containerSizeOffset = containerSizeOffset >= CONVERSATION_SIZE_OFFSET ? containerSizeOffset : CONVERSATION_SIZE_OFFSET;
  props.conversationContainer.style.maxHeight = `calc(100vh - ${containerSizeOffset}px)`;
}

function resizeTextArea() {
  let currentRows = getInputRowNumber();
  currentRows = currentRows > MAX_ROWS ? MAX_ROWS : currentRows;
  
  if (currentRows != inputRowNum.value) {
    inputRowNum.value = currentRows;
    let newHeight = (currentRows * inputFieldLineHeight.value)+inputFieldLineHeight.value + INPUT_PADDING;
    promptElement.value.$el.style.height = `${newHeight}px`;
    let heightDifference = newHeight - INPUT_MIN_SIZE;
    updateConversationContainerSize(heightDifference);
  }
}

const onSubmit = handleSubmit((values, ctx) => {
  
  if (values.prompt.trim() !== '') {
    const payload = {
      session_id: user.session_id,
      avatar_id: user.avatar.ID,
      msg: values.prompt,
    }
    resetForm();

    setTimeout(() => {
      resizeTextArea();
      chatStore.sendPrompt(payload);
    }, 0);
  }
});

onMounted(() => {
  feather.replace();
  inputFieldLineHeight.value = getInputLineHeight();
  inputRowNum.value = getInputRowNumber();
});

</script>

<template>
  <Form @submit="onSubmit">
    <div :class="{ 'textarea-container': true, 'loading': isLoading, 'form-control': true}" class="form-control">
      <Field
        v-on:input="resizeTextArea"
        @keydown.enter.exact.prevent="onSubmit"
        name="prompt"
        v-model="prompt"
        type="text"
        as="textarea"
        class=""
        rows="2"
        placeholder="Send a message..."
        :disabled="isLoading"
        ref="promptElement"
      ></Field>
      <button class="send-button btn btn-light" :disabled="isLoading">
        <i class="align-middle" data-feather="send"></i>
      </button>
    </div>
  </Form>
  
</template>

<style scoped>
.textarea-container {
  position: relative;
  background-color: transparent;
}

.loading {
  background-color: rgba(255, 255, 255, 0.1);
}

.send-button {
  position: absolute;
  right: 0.9rem;
  bottom: 0.9rem;
  background-color: black;
  color: white;
  border: 1px solid white;
}

.send-button:disabled {
  background-color: rgba(255, 255, 255, 0.1);
  color: black;
}

textarea {
  width: calc(100% - 60px);
  background-color: transparent;
  color: white !important;
  z-index: 999;
  border: none;
  min-height: 46px;
}

textarea:focus {
  background-color: transparent;
  color: white;
  border: none !important;
  outline: none !important;
}

textarea:disabled {
  color: white;
}

</style>
