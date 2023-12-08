<script setup>
import axios from 'axios';

const fetchData = () => {
    axios.get('http://localhost:8080/api/data')
        .then(response => {
            // Handle the response data
            console.log(response.data);
        })
        .catch(error => {
            // Handle errors
            console.error(error);
        });
};

onMounted(() => {
    fetchData();
});
</script>

<template>
  <h1>Login</h1>
  <form action="/api/login" method="post">
      <div>
          <label for="username">Username:</label>
          <input type="text" id="username" name="username" required>
      </div>
      <div>
          <label for="password">Password:</label>
          <input type="password" id="password" name="password" required>
      </div>
      <div>
          <button type="submit">Login</button>
      </div>
  </form>
</template>

<style scoped>

</style>
