<script setup>
import {useRoute , useRouter,RouterView} from "vue-router"
import cars from "../data/cars.json"
import ContactView from "./ContactView.vue"

const route = useRoute()
const router = useRouter()

const carId = parseInt(route.params.id)

const car = cars.find(c=>c.id ===carId )

const showContact = () =>{
    if(carId === 4) return;
    return router.push(`/cars/${carId}/contact`)
}
// console.log(route.params)
</script>
<template>

<div v-if="car">
    <h1>Car view</h1>
    <p>{{car.name}}</p>
    <p>{{car.year}}</p>
    <p>{{car.price}}</p>
    <button @click="showContact">Click me for Contact</button>
    <RouterView/>
</div>
<div v-else>
    <h1>CAR NOT FOUND</h1>
</div>

</template>