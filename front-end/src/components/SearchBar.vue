<script setup>
import { ref } from "vue";

const query = ref("");

function search() {
  console.log(query.value);
  fetch("http://localhost:5000/search", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: {
      query: query.value,
    },
  })
    .then((data) => {
      console.log(data);
    })
    .catch((error) => {
      console.error(error);
    });
}
</script>

<template>
  <input v-model="query" type="text" />
  <button @click="search">Search</button>
</template>
