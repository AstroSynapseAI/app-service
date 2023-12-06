<script setup>
import Sidebar from '@/components/Sidebar.vue';
import Navbar  from '@/components/Navbar.vue';
</script>

<template>
  <div class="wrapper">
    <Sidebar />
    <div class="main">
      <Navbar />
      <main class="content">
        <router-view />
      </main>
    </div>
  </div>
</template>

<style scoped>

</style>