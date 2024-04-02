<template>
  <div class="about">
    <h1>This is an callback page</h1>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';

const props = defineProps<{
  code: string
  state: string
}>()

const router = useRouter()
const serverLoginAPI = "http://localhost:3111/api/login"

// 使用 code 和 state 获取 accessToken
let resp = await fetch(`serverLoginAPI?code={props.code}&state={props.state}`, {
  method: "POST",
})

let res = await resp.json()

localStorage.setItem("accessToken", res.accessToken)
localStorage.setItem("refreshToken", res.refreshToken)

router.replace("/about")

</script>
