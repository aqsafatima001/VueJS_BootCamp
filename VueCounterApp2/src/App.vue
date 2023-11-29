<script setup>

  import {ref} from "vue"

  let count = ref(0) ;
  const addToCount = ()=>{
    count.value  = count.value+1
  }
  const SubtractFromCount = ()=>{
    count.value  = count.value+1
  }
</script>

<template>
  <main>
    <div>
      <h4>The current Count is...</h4> 
      <h1>
        {{count}}
      </h1>
      <button @click="addToCount()">+</button>
      <button @click="SubtractFromCount()">-</button>
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
