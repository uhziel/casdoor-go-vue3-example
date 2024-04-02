<template>
  <div class="about">
    <h1>This is an about page</h1>
    <ul>
      <li>name: {{ name }}</li>
      <li>email: {{ email }}</li>
    </ul>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';

const name = ref("")
const email = ref("")

onMounted(async () => {
  const accessToken = localStorage.getItem("accessToken")
  const serverInfoAPI = "http://localhost:3111/api/info"

  const resp = await fetch(serverInfoAPI, {
    headers: {
      "Authorization": `Bearer ${accessToken}`,
    },
  })

  const res = await resp.json()
  name.value = res.name
  email.value = res.email
})

</script>

<style>
@media (min-width: 1024px) {
  .about {
    min-height: 100vh;
    display: flex;
    align-items: center;
  }
}
</style>
