<script setup>
import { ref, onMounted } from 'vue';
import { Form, Field, useForm } from 'vee-validate';
import { useChatStore } from '../stores/chat.store.js';
import { useUserStore } from '../stores/user.store.js';

const user = useUserStore();
const chatStore = useChatStore();
const MAX_ROWS = 10;

let { resetForm, handleSubmit, defineField } = useForm({
  initialValues: {
    prompt: ''
  }
});

const [prompt, promptAttrs] = defineField('prompt');
const promptElement = ref(null);
let inputRowNum = ref(0);

function getInputRowNumber() {
  const lineHeight = parseInt(window.getComputedStyle(promptElement.value.$el).getPropertyValue("line-height"));
  return [Math.floor(promptElement.value.$el.scrollHeight / lineHeight), lineHeight];
}

function resizeTextArea() {
  let [currentRows, lineHeight] = getInputRowNumber();
  const maxRows = MAX_ROWS;
  currentRows = currentRows > maxRows ? maxRows : currentRows;
  if (currentRows != inputRowNum.value) {
    inputRowNum.value = currentRows;
    promptElement.value.$el.style.height = `${currentRows * lineHeight}px`;
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
  inputRowNum.value = getInputRowNumber();
});

</script>

<template>
  <Form @submit="onSubmit">
    <div :class="{ 'textarea-container': true, 'loading': chatStore.isLoading, 'form-control': true}" class="form-control">
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
        :disabled="chatStore.isLoading"
        ref="promptElement"
      ></Field>
      <button class="send-button btn btn-light" :disabled="chatStore.isLoading">
        <i class="align-middle" data-feather="send"></i>
      </button>
    </div>
  </Form>
  
</template>

<style scoped>
.textarea-container {
  position: relative;
  background-color: black;
}

.loading {
  background-color: rgba(255, 255, 255, 0.1);
}

.send-button {
  position: absolute;
  right: 0.5rem;
  bottom: 0.7rem;
  background-color: black;
  color: white;
  border: 1px solid white;
}

.send-button:disabled {
  background-color: rgba(255, 255, 255, 0.1);
  color: black;
}

textarea {
  width: calc(100% - 50px);
  background-color: transparent;
  color: white !important;
  z-index: 999;
  border: none;
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