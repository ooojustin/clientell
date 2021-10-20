<template>
    <div>
        <div class="ml-3 my-3" v-if="data">
            Found {{ data.length }} results
        </div>
        <PersonCard v-for="person in data" :data="person" :key="person.ID" />
    </div>
</template>

<script>
import PersonCard from "./PersonCard.vue";

export default {
    name: 'People',
    props: {
        data: Array
    },
    components: {
        PersonCard
    }
}
</script>

<style scoped>
</style>
