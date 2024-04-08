<template>
  <div class="about">
    <h1>This is an callback page</h1>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useRouter, useRoute } from 'vue-router';

const route = useRoute()
const router = useRouter()
const serverLoginAPI = "http://localhost:5173/user/login"

onMounted(async() => {
  // 使用 code 和 state 获取 accessToken
  let resp = await fetch(`${serverLoginAPI}?code=${route.query.code}&state=${route.query.state}`, {
    method: "GET",
  })

  let res = await resp.json()

  localStorage.setItem("accessToken", res.accessToken)

  router.replace("/about")
})

</script>
