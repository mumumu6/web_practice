<template>
    <div>
      <h1>Home Page</h1>
      <button @click="goToJson">Go to JSON Data</button>
    </div>
  </template>
  
  <script>
  export default {
    methods: {
      goToJson() {
        this.$router.push('/json');
      }
    }
  }
  </script>
  