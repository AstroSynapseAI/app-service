<script setup>
import { onMounted, ref, watch} from 'vue';
import { useRoute } from 'vue-router';
import { useHomeStore } from '@/stores/home.store.js';

const store = useHomeStore();
const route = useRoute();

let title = ref('');
let content = ref('');

watch(() => route.params.slug, (newSlug) => {
  newSlug = newSlug || 'about';
  const pageContent = store.contents[newSlug];

  if (pageContent) {
    title.value = pageContent.title;
    content.value = pageContent.description;
  }
}, { immediate: true });

onMounted(() => {
  feather.replace();
})


</script>

<template>  
  <div>
    <div class="row">
      <div class="col-8">
        <h1>{{ title }}</h1>
      </div>
      <div class="col-4">
        <router-link :to="{name: 'home'}"><i class="large-icon" data-feather="x"></i></router-link>
      </div>
    </div>

    <div class="row">
      <div class="col-12">
        <div v-html="content"></div>
      </div>
    </div>
  </div>
</template>

<style scoped>
a, h1 {
  color: white;
}

p {
  font-size: 16px;
}

.large-icon {
  float: right;
  width: 48px;
  height: 48px;
}

@media only screen and (max-width: 768px) {
  .large-icon {
    width: 36px;
    height: 36px;
  }
}
</style>