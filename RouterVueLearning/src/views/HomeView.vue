<script setup>
    import cars from "../data/cars.json"
    import {RouterLink} from "vue-router"

</script>
<template>
    <div>
        <h1>
            Home View
            <div class="cars">
                <RouterLink :to="`/cars/${car.id}`" v-for="car in cars" :key="car.id" href="">
                    {{ car.name }}
                </RouterLink>
            </div>
        </h1>
    </div>
</template>

<style scoped>
    .cars{
        display: flex;
    }
    .cars a{
        text-decoration: none;
        color: black;
        width: 100%;
        box-shadow: 1px 1px 10px 1px rgb(95, 91, 91);
        margin: 10px;
    }
</style>