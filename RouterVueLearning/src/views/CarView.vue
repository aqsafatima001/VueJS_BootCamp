<script setup>
import {useRoute , RouterView} from "vue-router"
import cars from "../data/cars.json"
import ContactView from "./ContactView.vue"

const route = useRoute()

const car = cars.find(c=>c.id === parseInt(route.params.id))

// console.log(route.params)
</script>
<template>

<div>
    <h1>Car view</h1>
    <p>{{car.name}}</p>
    <p>{{car.year}}</p>
    <p>{{car.price}}</p>
    <RouterView/>
</div>

</template>