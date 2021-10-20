<template>
    <div>
        <p>Found {{ data.length }} results</p>
        <br />
        <Person v-for="person in data" :data="person" :key="person.ID" />
    </div>
</template>

<script>
import Person from "./Person.vue";

export default {
    name: 'People',
    props: {
        data: Array
    },
    components: {
        Person
    }
}
</script>

<style scoped>
</style>
