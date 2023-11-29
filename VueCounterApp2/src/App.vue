<script setup>
let count = 0;

</script>

<template>
  <main>
    <div>
      <h4>The current Count is...</h4>
      <h1>
        {{count}}
      </h1>
      <button>+</button>
      <button>-</button>
    </div>
  </main>

</template>

<style scoped>
main{
  background-color: aliceblue;
  width: 100vw;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
}
div{
  text-align: center;
}
</style>
