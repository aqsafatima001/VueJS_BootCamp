<script setup>
import {ref} from "vue"
        import {RouterView} from "vue-router"
        import Nav from "./components/Nav.vue"

        const randomstate = ref("")
</script>

<template>
<div>
 
    <Nav />
    <RouterView />
</div>
</template>

<style scoped>

</style>


