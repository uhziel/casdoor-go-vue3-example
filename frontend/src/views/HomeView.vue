<script setup lang="ts">
import { ref } from 'vue';
import { useCasdoor } from 'casdoor-vue-sdk';

const loginURL = ref("")
const registerURL = ref("")

const casdoorSDK = useCasdoor()

loginURL.value = casdoorSDK.getSigninUrl()
registerURL.value = casdoorSDK.getSignupUrl()

</script>

<template>
  <main>
    <h1>首页</h1>
    <a :href="loginURL">登陆</a>
    <a :href="registerURL">注册</a>
  </main>
</template>
